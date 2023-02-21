package main

import (
	"errors"
	"fmt"
	"net/http"
	"os"
	"path/filepath"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/hyperledger/fabric-sdk-go/pkg/core/config"
	"github.com/hyperledger/fabric-sdk-go/pkg/gateway"
)

type FunctionBody struct {
	Name string   `json:"name"`
	Args []string `json:"args"`
}
type Err struct {
	Error string `json:"error"`
}

const (
	CHANNEL_NAME  = "blockchain2023"
	CONTRACT_NAME = "test"
)

func main() {

	os.RemoveAll("keystore")
	os.RemoveAll("wallet")
	c := gin.Default()
	c.Use(cors.Default())
	os.Setenv("DISCOVERY_AS_LOCALHOST", "true")

	c.POST("/transact", func(ctx *gin.Context) {
		var body FunctionBody
		ctx.BindJSON(&body)
		contact, gw := GetContract()
		defer gw.Close()
		json, err := contact.SubmitTransaction(body.Name, body.Args...)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, &Err{err.Error()})
			return
		}
		ctx.PureJSON(http.StatusOK, string(json))
	})
	c.POST("/call", func(ctx *gin.Context) {
		var body FunctionBody
		ctx.BindJSON(&body)
		contact, gw := GetContract()
		defer gw.Close()
		json, err := contact.EvaluateTransaction(body.Name, body.Args...)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, &Err{err.Error()})
			return
		}
		ctx.PureJSON(http.StatusOK, string(json))
	})
	c.Run("0.0.0.0:1110")
}

func GetContract() (*gateway.Contract, *gateway.Gateway) {
	os.Setenv("DISCOVERY_AS_LOCALHOST", "true")
	wallet, err := gateway.NewFileSystemWallet("wallet")
	if err != nil {
		fmt.Printf("Failed to create wallet: %s\n", err)
		os.Exit(1)
	}

	if !wallet.Exists("appUser") {
		err = populateWallet(wallet)
		if err != nil {
			fmt.Printf("Failed to populate wallet contents: %s\n", err)
			os.Exit(1)
		}
	}

	ccpPath := filepath.Join(
		"test-network",
		"organizations",
		"peerOrganizations",
		"org1.example.com",
		"connection-org1.yaml",
	)

	gw, err := gateway.Connect(
		gateway.WithConfig(config.FromFile(filepath.Clean(ccpPath))),
		gateway.WithIdentity(wallet, "appUser"),
	)
	if err != nil {
		fmt.Printf("Failed to connect to gateway: %s\n", err)
		os.Exit(1)
	}
	network, err := gw.GetNetwork(CHANNEL_NAME)
	if err != nil {
		fmt.Printf("Failed to get network: %s\n", err)
		os.Exit(1)
	}
	return network.GetContract(CONTRACT_NAME), gw
}

func populateWallet(wallet *gateway.Wallet) error {
	credPath := filepath.Join(
		"test-network",
		"organizations",
		"peerOrganizations",
		"org1.example.com",
		"users",
		"User1@org1.example.com",
		"msp",
	)

	certPath := filepath.Join(credPath, "signcerts", "cert.pem")
	// read the certificate pem
	cert, err := os.ReadFile(filepath.Clean(certPath))
	if err != nil {
		return err
	}

	keyDir := filepath.Join(credPath, "keystore")
	// there's a single file in this dir containing the private key
	files, err := os.ReadDir(keyDir)
	if err != nil {
		return err
	}
	if len(files) != 1 {
		return errors.New("keystore folder should have contain one file")
	}
	keyPath := filepath.Join(keyDir, files[0].Name())
	key, err := os.ReadFile(filepath.Clean(keyPath))
	if err != nil {
		return err
	}

	identity := gateway.NewX509Identity("Org1MSP", string(cert), string(key))

	err = wallet.Put("appUser", identity)
	if err != nil {
		return err
	}
	return nil
}
