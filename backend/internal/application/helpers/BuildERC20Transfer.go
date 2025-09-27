package helpers

import (
	"math/big"
	"strings"
)

func BuildERC20TransferData(to string, amount *big.Int) (string, error) {
	// ERC20 transfer function selector = first 4 bytes of keccak("transfer(address,uint256)")
	selector := "a9059cbb"

	// strip 0x
	to = strings.TrimPrefix(to, "0x")

	// left pad address to 32 bytes
	paddedTo := strings.Repeat("0", 64-len(to)) + to

	// amount in hex, left padded to 32 bytes
	amtHex := amount.Text(16)
	paddedAmt := strings.Repeat("0", 64-len(amtHex)) + amtHex

	// final data
	return "0x" + selector + paddedTo + paddedAmt, nil
}
