package config

import (
	"fmt"
	"strings"

	"github.com/ethereum/go-ethereum/crypto"
	"github.com/quytm/blockchain-engineer-interview/genomicdao-server/pkg/utils"
	"github.com/spf13/viper"
)

type Config struct {
	UserPrivateKey string `mapstructure:"user_private_key"`

	UseNetwork                  string     `mapstructure:"use_network"`
	NetworkLocal                NetworkCfg `mapstructure:"network_local"`
	NetworkAvalancheFujiTestnet NetworkCfg `mapstructure:"network_avalanche_fuji_testnet"`
}

func (c Config) GetUserPublicKey() ([]byte, error) {
	// Convert the private key hex string to an ECDSA private key
	ecdsaPrivateKey, err := utils.HexToECDSAPrivateKey(c.UserPrivateKey)
	if err != nil {
		fmt.Printf("HexToECDSAPrivateKey error : %v", err)
		return nil, err
	}

	// Generate user pubkey from private key
	return crypto.FromECDSAPub(&ecdsaPrivateKey.PublicKey), nil
}

func (c Config) GetNetworkCfg() NetworkCfg {
	switch c.UseNetwork {
	case "network_avalanche_fuji_testnet":
		return c.NetworkAvalancheFujiTestnet
	default:
		return c.NetworkLocal
	}
}

// Load config from config.yml
func Load(paths ...string) Config {
	vip := viper.New()

	vip.SetConfigName("config")
	vip.SetConfigType("yml")
	if len(paths) == 0 {
		vip.AddConfigPath(".") // ROOT
	} else {
		vip.AddConfigPath(paths[0])
	}

	vip.SetEnvPrefix("docker")
	vip.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	vip.AutomaticEnv()

	err := vip.ReadInConfig()
	if err != nil {
		panic(err)
	}

	// workaround https://github.com/spf13/viper/issues/188#issuecomment-399518663
	// to allow read from environment variables when Unmarshal
	for _, key := range vip.AllKeys() {
		val := vip.Get(key)
		vip.Set(key, val)
	}

	fmt.Println("===== Config file used:", vip.ConfigFileUsed())

	cfg := Config{}
	err = vip.Unmarshal(&cfg)
	if err != nil {
		panic(err)
	}

	return cfg
}
