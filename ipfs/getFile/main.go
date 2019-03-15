package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

func main() {
	cid := "QmTaBR3VpGt8UBoGRBDvaRCDvnRRCPRHWoNCTsBGVTFgpy" // "testing stuff 1109!"
	endpoint := fmt.Sprintf("http://localhost:8080/ipfs/%s", cid)
	resp, err := http.Get(endpoint)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %s", err)
		os.Exit(1)
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %s", err)
		os.Exit(1)
	}
	fmt.Println(string(body))
}
