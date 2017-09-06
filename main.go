package main

import bittrex "github.com/Toorop/go-bittrex"

type args struct {
	bittrexKey    string
	bittrexSecret string
}

func main() {

	bittrex := bittrex.New(API_KEY, API_SECRET)

}
