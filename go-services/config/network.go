package config

type NetworkConfig struct {
	ChainID uint
	URL     string
}

var networkConfig = map[string]*NetworkConfig{
	"localnet": {
		ChainID: 31415926,
		URL:     "http://localhost:8545",
	},
	"calibrationnet": {
		ChainID: 314159,
		URL:     "https://api.calibration.node.glif.io/rpc/v1",
	},
	"filecoinmainnet": {
		ChainID: 314,
		URL:     "https://api.node.glif.io",
	},
}

func GetNetworkConfig(network string) *NetworkConfig {
	return networkConfig[network]
}
