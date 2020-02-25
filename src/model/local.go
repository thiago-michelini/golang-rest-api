package model

import "database/sql"

//Local representa a localidade
type Local struct {
	Pais             string         `json:"pais" db:"country"`
	Cidade           sql.NullString `json:"cidade" db:"city"`
	CodigoTelefonico int            `json:"codTelefone" db:"telcode"`
}
