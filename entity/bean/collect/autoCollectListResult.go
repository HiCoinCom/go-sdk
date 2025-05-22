package collect

/**
[
    {
        "symbol":"ETH",
        "amount":"0.0000111",
        "real_fee":"0",
        "trans_type":10,
        "fee":"0.0002782353",
        "address_to":"0xc70d1eebb7c687ec8d56bead73f104d41e6e0bda",
        "memo": "",
        "created_at":1672304978000,
        "txid":"0x8e6beba81b90835fc9fcd40a2bdca33243c7c3b81ac765c240837d4810874a55",
        "confirmations":0,
        "contract_address":"",
        "sub_wallet_id":123,
        "address_from":"0x5EDc9177997Bf6B4db559A5C184051858B3B3704",
        "fee_symbol":"ETH",
        "updated_at":1672318660000,
        "base_symbol":"",
        "id":242,
        "status":1200
    },
    {
        "symbol":"HECO",
        "amount":"0.0000111",
        "real_fee":"0",
        "trans_type":11,
        "fee":"0.0024040607337957",
        "address_to":"0x0AabA82E4ba9c2Fdf80B9F8E1AcE885f092B64F0",
        "memo": "",
        "created_at":1672304978000,
        "txid":"",
        "confirmations":0,
        "contract_address":"",
        "sub_wallet_id":123975,
        "address_from":"0xE76a3d30dAc35C4b9A17E690cd250d8Bec649b65",
        "fee_symbol":"HECO",
        "updated_at":1672318660000,
        "base_symbol":"HECO",
        "id":315,
        "status":2400
    }
]

*/

// AutoCollectListResult struct
type AutoCollectListResult struct {
	Code string                `json:"code"`
	Msg  string                `json:"msg"`
	Data []AutoCollectListData `json:"data"`
}

// AutoCollectListData struct
type AutoCollectListData struct {
	Symbol          string `json:"symbol"`
	Amount          string `json:"amount"`
	RealFee         string `json:"real_fee"`
	TransType       int64  `json:"trans_type"`
	Fee             string `json:"fee"`
	AddressTo       string `json:"address_to"`
	Memo            string `json:"memo"`
	CreatedAt       int64  `json:"created_at"`
	Txid            string `json:"txid"`
	Confirmations   int64  `json:"confirmations"`
	ContractAddress string `json:"contract_address"`
	SubWalletId     int64  `json:"sub_wallet_id"`
	AddressFrom     string `json:"address_from"`
	FeeSymbol       string `json:"fee_symbol"`
	UpdatedAt       int64  `json:"updated_at"`
	BaseSymbol      string `json:"base_symbol"`
	Id              int64  `json:"id"`
	Status          int64  `json:"status"`
}
