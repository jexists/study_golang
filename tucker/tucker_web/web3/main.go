package main

import (
	"net/http"
	"web3/myapp"
)

func main() {
	// http.ListenAndServe(":3000", mux)
	http.ListenAndServe(":3000", myapp.NewHttpHandler())
}

// golang Setting
// "gopls": {
// 	"experimentalWorkspaceModule": true,
// }
