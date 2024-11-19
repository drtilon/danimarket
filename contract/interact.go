package contract

import (
    "context"
    "fmt"
    "log"
    "math/big"

    "github.com/ethereum/go-ethereum"
    "github.com/ethereum/go-ethereum/accounts/abi"
    "github.com/ethereum/go-ethereum/common"
    "github.com/ethereum/go-ethereum/ethclient"
)

func ConnectToEthereum() (*ethclient.Client, error) {
    client, err := ethclient.Dial("https://mainnet.infura.io/v3/YOUR_INFURA_PROJECT_ID")
    if err != nil {
        return nil, err
    }
    return client, nil
}

func InteractWithContract() {
    client, err := ConnectToEthereum()
    if err != nil {
        log.Fatalf("Failed to connect to Ethereum: %v", err)
    }

    contractAddress := common.HexToAddress("0xYourContractAddress")
    query := "function getMarketData() public view returns (string)"

    data, err := client.CallContract(context.Background(), ethereum.CallMsg{
        To:   &contractAddress,
        Data: []byte(query),
    }, nil)
    if err != nil {
        log.Fatalf("Failed to interact with contract: %v", err)
    }

    fmt.Println("Contract result:", string(data))
}

