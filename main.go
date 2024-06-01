package main

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/wealdtech/go-ens/v3"
)

const rpcUrl = "https://eth.llamarpc.com"

func handleRequest(w http.ResponseWriter, req *http.Request) {
	// parse string from request
	// module@version/file.ext
	requestParts := strings.Split(req.RequestURI[1:], "@")
	module := requestParts[0]
	versionAndFile := strings.Split(requestParts[1], "/")
	version := versionAndFile[0]
	// file := versionAndFile[1]

	// ENS stuff
	client, err := ethclient.Dial(rpcUrl)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintln(w, err.Error())
	}

	resolver, err := ens.NewResolver(client, fmt.Sprintf("%s.swarmpm.eth", module))

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintln(w, err.Error())
	}

	swarmContentHash, err := resolver.Text(version)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintln(w, err.Error())
	}

	contentHash, _ := ens.ContenthashToString([]byte(swarmContentHash))

	fmt.Println("contentHash")

	if contentHash == "" {
		http.NotFound(w, req)
	}
}

func main() {

	http.HandleFunc("/", handleRequest)

	http.ListenAndServe(":8090", nil)
}
