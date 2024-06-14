package models

//  "github.com/volatiletech/sqlboiler/v4/queries/qm"

type Order struct {
	id           int
	brand        string
	color        string
	total_amount int
	time_stamp   string
}

// func InitModels(db  *sql.DB) error{

// }
