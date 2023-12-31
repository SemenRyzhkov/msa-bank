package models

import (
	"encoding/json"

)
type Message struct {
	Schema struct {
		Type   string `json:"type"`
		Fields []struct {
			Type   string `json:"type"`
			Fields []struct {
				Type     string `json:"type"`
				Optional bool   `json:"optional"`
				Name     string `json:"name,omitempty"`
				Version  int    `json:"version,omitempty"`
				Field    string `json:"field"`
			} `json:"fields,omitempty"`
			Optional bool   `json:"optional"`
			Name     string `json:"name,omitempty"`
			Field    string `json:"field"`
		} `json:"fields"`
		Optional bool   `json:"optional"`
		Name     string `json:"name"`
	} `json:"schema"`
	Payload struct {
		Before interface{} `json:"before"`
		After  json.RawMessage
		Source struct {
			Version   string      `json:"version"`
			Connector string      `json:"connector"`
			Name      string      `json:"name"`
			TsMs      int64       `json:"ts_ms"`
			Snapshot  string      `json:"snapshot"`
			Db        string      `json:"db"`
			Schema    string      `json:"schema"`
			Table     string      `json:"table"`
			TxID      int         `json:"txId"`
			Lsn       int         `json:"lsn"`
			Xmin      interface{} `json:"xmin"`
		} `json:"source"`
		Op          string      `json:"op"`
		TsMs        int64       `json:"ts_ms"`
		Transaction interface{} `json:"transaction"`
	} `json:"payload"`
}
