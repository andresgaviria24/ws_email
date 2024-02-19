package dto

type Email struct {
	To           []string `json:"to"`
	Body         string   `json:"body"`
	System       string   `json:"system"`
	Subject      string   `json:"subject"`
	AttachBase64 string   `json:"attach"`
	NameAttach   string   `json:"nameattach"`
}

type EmailBrevo struct {
	Sender      InfoEmail    `json:"sender"`
	To          []InfoEmail  `json:"to"`
	Subject     string       `json:"subject"`
	HTMLContent string       `json:"htmlContent"`
	Attachment  []Attachment `json:"attachment"`
}

type InfoEmail struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

type Attachment struct {
	Content string `json:"content"`
	Name    string `json:"name"`
}
