package wallet

/**
{
    "open_main_chain":[
        {
            "coin_net":"BTC",
            "symbol":"BTC",
            "symbol_alias":"BTC",
            "support_acceleration" : false
        },
        {
            "coin_net":"ETH",
            "symbol":"ETH",
            "symbol_alias":"ETH",
            "support_acceleration" : true
        }
    ],
    "support_main_chain":[
        {
            "coin_net":"BTC",
            "if_open_chain":true,
            "symbol":"BTC",
            "symbol_alias":"BTC",
            "support_acceleration" : false
        },
        {
            "coin_net":"ETH",
            "if_open_chain":false,
            "symbol":"ETH",
            "symbol_alias":"ETH",
            "support_acceleration" : true
        }
    ]
}
*/

type OpenCoinResult struct {
	Code string       `json:"code"`
	Data OpenCoinData `json:"data"`
	Msg  string       `json:"msg"`
}

type OpenCoinData struct {
	OpenMainChain    []OpenMainChain    `json:"open_main_chain"`
	SupportMainChain []SupportMainChain `json:"support_main_chain"`
}

type OpenMainChain struct {
	CoinNet             string `json:"coin_net"`
	Symbol              string `json:"symbol"`
	SymbolAlias         string `json:"symbol_alias"`
	SupportAcceleration bool   `json:"support_acceleration"`
}

type SupportMainChain struct {
	CoinNet             string `json:"coin_net"`
	IfOpenChain         bool   `json:"if_open_chain"`
	Symbol              string `json:"symbol"`
	SymbolAlias         string `json:"symbol_alias"`
	SupportAcceleration bool   `json:"support_acceleration"`
}
