package model

//Local representa a localidade
type Local struct {
	Pais string `json:"pais" db:"COUNTRY"`
	//Cidade           sql.NullString `json:"cidade" db:"CITY"`
	Cidade           string `json:"cidade" db:"CITY"`
	CodigoTelefonico int    `json:"codTelefone" db:"TELCODE"`
}
