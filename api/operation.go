package main

type Operation struct {
	ID       string         `json:"id"`
	Done     bool           `json:"done"`
	Metadata map[string]any `json:"metadata"`
	Response any            `json:"response"`
	Error    any            `json:"error"`
}

type OperationMetadata struct {
	Progress int    `json:"progress"`
	Status   string `json:"status"`
}

