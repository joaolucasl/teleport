package database

import (
	"encoding/json"
	"github.com/pagarme/teleport/batcher/ddldiff"
)

type Ddl struct {
	PreSchemas []Schema `json:"pre"`
	PostSchemas []Schema `json:"post"`
}

func NewDdl(data []byte) *Ddl {
	var ddl Ddl
	err := json.Unmarshal(data, &ddl)

	if err != nil {
		panic(err)
	}

	return &ddl
}

func (d *Ddl) schemaToDiffable(schema []Schema) []ddldiff.Diffable {
	diff := make([]ddldiff.Diffable, 0)

	for i, _ := range schema {
		diff = append(diff, &schema[i])
	}

	return diff
}

func (d *Ddl) Diff() []ddldiff.Action {
	return ddldiff.Diff(
		d.schemaToDiffable(d.PreSchemas),
		d.schemaToDiffable(d.PostSchemas),
	)
}