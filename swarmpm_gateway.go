package main

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/wealdtech/go-ens/v3"
)

const RPC = "https://eth.llamarpc.com"

func handleRequest(w http.ResponseWriter, req *http.Request) {
	// parse string from request
	// module@version/file.ext
	requestParts := strings.Split(req.RequestURI[1:], "@")
	module := requestParts[0]
	versionAndFile := strings.Split(requestParts[1], "/")

	// no error handling yet
	client, _ := ethclient.Dial(RPC)
	address, _ := ens.Resolve(client, module)

	fmt.Fprintln(w, address)
	//just printing parased module version and file
	fmt.Fprintln(w, versionAndFile[0])
	fmt.Fprintln(w, versionAndFile[1])

}

func main() {

	http.HandleFunc("/", handleRequest)

	http.ListenAndServe(":8090", nil)
}
