package main

import (
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/iotaledger/iota.go/api"
	"github.com/iotaledger/iota.go/bundle"
	"github.com/iotaledger/iota.go/pow"
	"github.com/iotaledger/iota.go/trinary"
)

const seed = trinary.Trytes("JBN9ZRCOH9YRUGSWIQNZWAIFEZUBDUGTFPVRKXWPAUCEQQFS9NHPQLXCKZKRHVCCUZNF9CZZWKXRZVCWQ")
const address = trinary.Trytes("XBN9ZRCOH9YRUGSWIQNZWAIFEZUBDUGTFPVRKXWPAUCEQQFS9NHPQLXCKZKRHVCCUZNF9CZZWKXRZVCWQMZOCAHYPD")

func main() {
	// compose a new API instance, we provide no PoW function so this uses remote PoW
	_api, _ := api.ComposeAPI(api.HTTPClientSettings{URI: "http://localhost:14265"})
	message := ""

	transfer := bundle.Transfer{
		Address: address,
		Value:   0,
		Message: message,
	}
	num, _ := strconv.Atoi(os.Args[1])

	var transfers bundle.Transfers
	for i := 0; i < num; i++ {
		transfers = append(transfers, transfer)
	}

	// Use the default options
	prepTransferOpts := api.PrepareTransfersOptions{}

	trytes, _ := _api.PrepareTransfers(seed, transfers, prepTransferOpts)

	_, powImpl := pow.GetFastestProofOfWorkImpl()
	trunk := trinary.Hash("IZDLTHKCRVYXJVGKHJG9DGNOSWSHNKDLCOSBFYCQEIWDNZJWYKHFENJNQCEJMBUNCETAMPYPZFWVZ9999")
	branch := trunk
	start := time.Now()
	_, _ = pow.DoPoW(trunk, branch, trytes, 14, powImpl)
	fmt.Printf("Took %f seconds\n", time.Since(start).Seconds())

}
