package btc

import "github.com/btcsuite/btcd/rpcclient"

func getClientFor(chainnet string) (*rpcclient.Client, error) {
	switch chainnet {
	case ChainSignet:
		return rpcclient.New(&rpcclient.ConnConfig{
			Host:         "115.29.163.193:38332",
			User:         "auth",
			Pass:         "bitcoin-b2dd077",
			HTTPPostMode: true,
			DisableTLS:   true,
		}, nil)

	case ChainMainnet, ChainBitcoin:
		return rpcclient.New(&rpcclient.ConnConfig{
			Host:         "115.29.163.193:8332",
			User:         "auth",
			Pass:         "bitcoin-b2dd077",
			HTTPPostMode: true,
			DisableTLS:   true,
		}, nil)
	}

	return nil, ErrUnsupportedChain
}
