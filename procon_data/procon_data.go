package procon_data

type Msg struct {
	Jwt string `json:"jwt"`
	Type string `json:"type"`
	Data string	`json:"data"`
}