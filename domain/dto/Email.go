package dto

type Email struct {
	To           []string `json:"to"`
	Body         string   `json:"body"`
	System       string   `json:"system"`
	Subject      string   `json:"subject"`
	AttachBase64 string   `json:"attach"`
	NameAttach   string   `json:"nameattach"`
}
