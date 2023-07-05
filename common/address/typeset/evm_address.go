package typeset

import (
	"regexp"

	"github.com/consolelabs/mochi-typeset/common/chain/typeset"
)

type AccountAddress struct {
	hash  string
	chain string
}

// TODO: add test
// FromHash returns an AccountAddress from a hash
func FromHash(hash string) *AccountAddress {
	if hash == "" {
		return nil
	}

	// check if evm address
	if evm := evmCheck(hash); evm {
		return &AccountAddress{
			hash:  hash,
			chain: typeset.CHAIN_TYPE_EVM,
		}
	}

	if solana := solanaCheck(hash); solana {
		return &AccountAddress{
			hash:  hash,
			chain: typeset.CHAIN_TYPE_SOLANA,
		}
	}

	if ronin := roninCheck(hash); ronin {
		return &AccountAddress{
			hash:  hash,
			chain: typeset.CHAIN_TYPE_RONIN,
		}
	}

	if bitcoin := bitcoinCheck(hash); bitcoin {
		return &AccountAddress{
			hash:  hash,
			chain: typeset.CHAIN_TYPE_BITCOIN,
		}
	}

	return nil
}

func (a *AccountAddress) Hash() string {
	return a.hash
}

func (a *AccountAddress) Chain() string {
	return a.chain
}

func evmCheck(hash string) bool {
	re := regexp.MustCompile("^0x[0-9a-fA-F]{40}$")
	return re.MatchString(hash)
}

func solanaCheck(hash string) bool {
	return false
}

func roninCheck(hash string) bool {
	return false
}

func bitcoinCheck(hash string) bool {
	return false
}
