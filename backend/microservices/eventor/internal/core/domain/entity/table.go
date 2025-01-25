package entity

type UUID string

type Age struct {
	Number uint16
	Unit   uint8
}

type TableConfig struct {
	CryptoCurrency string `json:"cryptocurrency"`
	MinAmount      string `json:"min_amount"`
	Age            Age    `json:"age"`
}

type TableRecord struct {
	BlockId        int     `json:"block_id"`
	OutputTotalUsd float64 `json:"output_total_usd"`
	Hash           string  `json:"hash"`
	Time           string  `json:"time"`
	Res            string  `json:"res"`
	Token          string  `json:"token"`
	Type           string  `json:"type"`
}
