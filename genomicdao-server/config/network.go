package config

type NetworkCfg struct {
	Url                       string           `mapstructure:"url"`
	UrlWs                     string           `mapstructure:"url_ws"`
	Controller                SmartContractCfg `mapstructure:"controller"`
	PostCovidStrokePrevention SmartContractCfg `mapstructure:"post_covid_stroke_prevention"`
	GeneNFT                   SmartContractCfg `mapstructure:"gene_nft"`
}

type SmartContractCfg struct {
	Address string `mapstructure:"address"`
}
