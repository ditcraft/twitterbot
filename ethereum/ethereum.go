package ethereum

import (
	"context"
	"errors"
	"math/big"
	"os"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
)

// SendEther will send a specified amount of ETH (in wei) to the target address
func SendEther(_targetAddress string, _amount int64) error {
	if len(_targetAddress) < 40 || len(_targetAddress) > 42 {
		return errors.New("Address has a wrong length")
	}
	if _amount <= 0 {
		return errors.New("Faulty amount to be sent")
	}

	connection, err := getConnection()
	if err != nil {
		return err
	}

	weiAmount := etherToWei(big.NewInt(_amount))
	targetAddress := common.HexToAddress(_targetAddress)

	ethBalance, err := GetBalance(os.Getenv("ETHEREUM_ADDRESS"))
	if err != nil {
		return err
	}

	// If our account doesn't have enough ether
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
	defaultGasPrice := big.NewInt(10000000000)
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
	nonpendingNonce, err := _connection.NonceAt(context.Background(), common.HexToAddress(os.Getenv("ETHEREUM_ADDRESS")), nil)
	if err != nil {
		return nil, errors.New("Failed to retrieve nonce for ethereum transaction")
	}

	// Edge-Case for slow nodes
	nonce := pendingNonce
	if nonpendingNonce > pendingNonce {
		nonce = nonpendingNonce
	}

	// Retrieving the suggested gasprice by the network
	gasPrice, err := _connection.SuggestGasPrice(context.Background())
	if err != nil {
		return nil, errors.New("Failed to retrieve the gas-price for ethereum transaction")
	}

	// Minimum gas price is 10 gwei for now, which works best for rinkeby
	// Will be changed later on
	defaultGasPrice := big.NewInt(10000000000)
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

// getConnection will return a connection to the ethereum blockchain
func getConnection() (*ethclient.Client, error) {
	// Connecting to rinkeby via infura
	connection, err := ethclient.Dial(os.Getenv("ETHEREUM_RPC"))
	if err != nil {
		return nil, errors.New("Failed to connect to the ethereum network")
	}

	return connection, nil
}
