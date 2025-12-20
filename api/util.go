package main

import (
	"encoding/base64"
	"encoding/json"
)

type Pagination struct {
	Offset string `json:"o"`
}

func Encode(cursor string) string {
	if cursor == "" {
		return ""
	}

	p := &Pagination{Offset: cursor}
	b, _ := json.Marshal(p)
	return base64.URLEncoding.EncodeToString(b)
}

func Decode(token string) (string, error) {
	b, err := base64.URLEncoding.DecodeString(token)
	if err != nil {
		return "", err
	}

	var p Pagination
	if err := json.Unmarshal(b, &p); err != nil {
		return "", err
	}

	return p.Offset, nil
}
