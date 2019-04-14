package ethereum

import (
	"context"
	"errors"
	"math/big"
	"os"
	"strconv"
	"sync"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/marvinkruse/dit-twitterbot/smartcontracts/ditCoordinator"
	"github.com/marvinkruse/dit-twitterbot/smartcontracts/ditToken"
)

// Mutex to keep the nonces in order
var Mutex = &sync.Mutex{}

// KYCPassed will add an address to the list of allowed accounts to interact
// with the ditCoordinator(s) in the demo and/or live contract.
func KYCPassed(_address string, _live bool) error {
	connection, err := getConnection()
	if err != nil {
		return err
	}

	userAddress := common.HexToAddress(_address)

	ditCoordinatorInstance, err := getDitCoordinatorInstance(connection, _live)
	if err != nil {
		return err
	}

	alreadyPassedKYC, err := ditCoordinatorInstance.PassedKYC(nil, userAddress)
	if err != nil {
		return err
	}

	if alreadyPassedKYC {
		return nil
	}

	auth, err := populateTx(connection)
	if err != nil {
		return err
	}

	_, err = ditCoordinatorInstance.PassKYC(auth, userAddress)
	if err != nil {
		return err
	}

	waitingFor := 0
	passedKYC := false
	for !passedKYC {
		waitingFor += 2
		time.Sleep(2 * time.Second)

		// Checking the KYC status every 2 seconds
		passedKYC, err = ditCoordinatorInstance.PassedKYC(nil, userAddress)
		if err != nil {
			return err
		}
		// If we are waiting for more than 2 minutes, the transaction might have failed
		if waitingFor > 180 {
			return errors.New("Transaction failed")
		}
	}

	return nil
}

// SendDitTokens will send a specified amount of xDit tokens to the target address
func SendDitTokens(_targetAddress string) error {
	connection, err := getConnection()
	if err != nil {
		return err
	}

	targetAddress := common.HexToAddress(_targetAddress)

	intValue, _ := strconv.Atoi(os.Getenv("DIT_TOKEN_AMOUNT"))
	bigIntValue := big.NewInt(int64(intValue))
	bigIntWeiValue := new(big.Int).Mul(bigIntValue, big.NewInt(1000000000000000000))

	ditTokenInstance, err := getDitTokenInstance(connection)
	if err != nil {
		return err
	}

	oldxDitBalance, err := ditTokenInstance.BalanceOf(nil, targetAddress)
	if err != nil {
		return err
	}

	if oldxDitBalance.Cmp(bigIntWeiValue) != -1 {
		return nil
	}

	auth, err := populateTx(connection)
	if err != nil {
		return err
	}

	_, err = ditTokenInstance.Mint(auth, targetAddress, bigIntWeiValue)
	if err != nil {
		return err
	}

	waitingFor := 0
	newxDitBalance := oldxDitBalance
	for newxDitBalance.Cmp(oldxDitBalance) == 0 {
		waitingFor += 2
		time.Sleep(2 * time.Second)

		// Checking the balance of the user every 2 seconds, if it changed, a transaction was executed
		newxDitBalance, err = ditTokenInstance.BalanceOf(nil, targetAddress)
		if err != nil {
			return err
		}
		// If we are waiting for more than 2 minutes, the transaction might have failed
		if waitingFor > 180 {
			return errors.New("Transaction failed")
		}
	}

	return nil
}

// SendXDaiCent will send a cent of xDai to the target address
func SendXDaiCent(_targetAddress string) error {
	if len(_targetAddress) < 40 || len(_targetAddress) > 42 {
		return errors.New("Address has a wrong length")
	}

	connection, err := getConnection()
	if err != nil {
		return err
	}

	weiAmount := big.NewInt(10000000000000000)
	targetAddress := common.HexToAddress(_targetAddress)

	xDaiBalanceTarget, err := GetBalance(_targetAddress)
	if err != nil {
		return err
	}

	if xDaiBalanceTarget.Cmp(weiAmount) != -1 {
		return nil
	}

	ethBalance, err := GetBalance(os.Getenv("ETHEREUM_ADDRESS"))
	if err != nil {
		return err
	}

	// If our account doesn't have enough xDai
	if ethBalance.Cmp(weiAmount) != 1 {
		return errors.New("Bot account doesn't have enough funds")
	}

	// Retrieving the current pending nonce of our address
	pendingNonce, err := connection.PendingNonceAt(context.Background(), common.HexToAddress(os.Getenv("ETHEREUM_ADDRESS")))
	if err != nil {
		return errors.New("Failed to retrieve nonce for ethereum transaction")
	}
	// Retrieving the current non-pending nonce of our address
	nonpendingNonce, err := connection.NonceAt(context.Background(), common.HexToAddress(os.Getenv("ETHEREUM_ADDRESS")), nil)
	if err != nil {
		return errors.New("Failed to retrieve nonce for ethereum transaction")
	}

	// Edge-Case for slow nodes
	nonce := pendingNonce
	if nonpendingNonce > pendingNonce {
		nonce = nonpendingNonce
	}

	// Retrieving the suggested gasprice by the network
	gasPrice, err := connection.SuggestGasPrice(context.Background())
	if err != nil {
		return errors.New("Failed to retrieve the gas-price for ethereum transaction")
	}

	// Minimum gas price is 10 gwei for now, which works best for rinkeby
	// Will be changed later on
	defaultGasPrice := big.NewInt(1000000000)
	if gasPrice.Cmp(defaultGasPrice) != 1 {
		gasPrice = defaultGasPrice
	}

	value := weiAmount
	gasLimit := uint64(21000)

	// Converting the private key string into a private key object
	privateKey, err := crypto.HexToECDSA(os.Getenv("ETHEREUM_PRIVATEKEY"))
	if err != nil {
		return errors.New("Failed to convert ethereum private-key")
	}

	tx := types.NewTransaction(nonce, targetAddress, value, gasLimit, gasPrice, []byte(nil))

	chainID, err := connection.NetworkID(context.Background())
	if err != nil {
		return err
	}

	signedTx, err := types.SignTx(tx, types.NewEIP155Signer(chainID), privateKey)
	if err != nil {
		return err
	}

	err = connection.SendTransaction(context.Background(), signedTx)
	if err != nil {
		return err
	}

	waitingFor := 0
	newxDaiBalanceTarget := xDaiBalanceTarget
	for newxDaiBalanceTarget.Cmp(xDaiBalanceTarget) == 0 {
		waitingFor += 2
		time.Sleep(2 * time.Second)

		// Checking the balance of the user every 2 seconds, if it changed, a transaction was executed
		newxDaiBalanceTarget, err = GetBalance(_targetAddress)
		if err != nil {
			return err
		}
		// If we are waiting for more than 2 minutes, the transaction might have failed
		if waitingFor > 180 {
			return errors.New("Transaction failed")
		}
	}

	return nil
}

// GetBalance will return the balance of a user
func GetBalance(_targetAddress string) (*big.Int, error) {
	if len(_targetAddress) < 40 || len(_targetAddress) > 42 {
		return nil, errors.New("Address has a wrong length")
	}

	connection, err := getConnection()
	if err != nil {
		return nil, err
	}

	ethBalance, err := connection.BalanceAt(context.Background(), common.HexToAddress(_targetAddress), nil)
	if err != nil {
		return nil, errors.New("Failed to retrieve ETH balance")
	}

	return ethBalance, nil
}

// populateTX will set the necessary values for a ethereum transaction
// amount of gas, gasprice, nonce, sign this with the private key
func populateTx(_connection *ethclient.Client) (*bind.TransactOpts, error) {
	// Converting the private key string into a private key object
	privateKey, err := crypto.HexToECDSA(os.Getenv("ETHEREUM_PRIVATEKEY"))
	if err != nil {
		return nil, errors.New("Failed to convert ethereum private-key")
	}

	// Retrieving the current pending nonce of our address
	pendingNonce, err := _connection.PendingNonceAt(context.Background(), common.HexToAddress(os.Getenv("ETHEREUM_ADDRESS")))
	if err != nil {
		return nil, errors.New("Failed to retrieve nonce for ethereum transaction")
	}
	// Retrieving the current non-pending nonce of our address
	nonPendingNonce, err := _connection.NonceAt(context.Background(), common.HexToAddress(os.Getenv("ETHEREUM_ADDRESS")), nil)
	if err != nil {
		return nil, errors.New("Failed to retrieve nonce for ethereum transaction")
	}

	// Edge-Case for slow nodes
	nonce := pendingNonce
	if nonPendingNonce > pendingNonce {
		nonce = nonPendingNonce
	}

	// Retrieving the suggested gasprice by the network
	gasPrice, err := _connection.SuggestGasPrice(context.Background())
	if err != nil {
		return nil, errors.New("Failed to retrieve the gas-price for ethereum transaction")
	}

	// Minimum gas price is 10 gwei for now, which works best for rinkeby
	// Will be changed later on
	defaultGasPrice := big.NewInt(1000000000)
	if gasPrice.Cmp(defaultGasPrice) != 1 {
		gasPrice = defaultGasPrice
	}
	// Setting the values into the transaction-options object
	auth := bind.NewKeyedTransactor(privateKey)
	auth.Nonce = big.NewInt(int64(nonce))

	auth.Value = big.NewInt(0)
	auth.GasLimit = uint64(6000000)
	auth.GasPrice = gasPrice

	return auth, nil
}

func weiToEther(wei *big.Int) *big.Int {
	return new(big.Int).Div(wei, big.NewInt(1000000000000000000))
}

func etherToWei(ether *big.Int) *big.Int {
	return new(big.Int).Mul(ether, big.NewInt(1000000000000000000))
}

// getDitTokenInstance will return an instance of the deployed ditToken contract
func getDitTokenInstance(_connection *ethclient.Client) (*ditToken.MintableERC20, error) {
	ditTokenAddress := common.HexToAddress(os.Getenv("CONTRACT_DIT_TOKEN"))

	// Create a new instance of the ditToken contract to access it
	ditTokenInstance, err := ditToken.NewMintableERC20(ditTokenAddress, _connection)
	if err != nil {
		return nil, errors.New("Failed to find ditToken at provided address")
	}

	return ditTokenInstance, nil
}

// getDitCoordinatorInstance will return an instance of the deployed ditCoordinator contract
func getDitCoordinatorInstance(_connection *ethclient.Client, _live bool) (*ditCoordinator.DitCoordinator, error) {
	var ditCoordinatorAddress common.Address
	if _live {
		ditCoordinatorAddress = common.HexToAddress(os.Getenv("CONTRACT_DIT_COORDINATOR_LIVE"))
	} else {
		ditCoordinatorAddress = common.HexToAddress(os.Getenv("CONTRACT_DIT_COORDINATOR_DEMO"))
	}

	// Create a new instance of the ditToken contract to access it
	ditCoordinatorInstance, err := ditCoordinator.NewDitCoordinator(ditCoordinatorAddress, _connection)
	if err != nil {
		return nil, errors.New("Failed to find ditCoordinator at provided address")
	}

	return ditCoordinatorInstance, nil
}

// getConnection will return a connection to the ethereum blockchain
func getConnection() (*ethclient.Client, error) {
	// Connecting to rinkeby via infura
	connection, err := ethclient.Dial(os.Getenv("ETHEREUM_RPC"))
	if err != nil {
		return nil, errors.New("Failed to connect to the ethereum network")
	}

	return connection, nil
}
