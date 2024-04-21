package repositories

import (
	"context"
	"customer-data-api/internal/models"
	"database/sql"
)

type CustomerRepo interface {
	Create(data *models.CustomerData) error
}

type customerRepo struct {
	*sql.DB
}

func NewCustomerRepo(db *sql.DB) CustomerRepo {
	return &customerRepo{db}
}

func (r *customerRepo) Create(data *models.CustomerData) error {
	tx, err := r.DB.BeginTx(context.Background(), nil)
	if err != nil {
		return err
	}
	defer tx.Rollback()

	row := tx.QueryRow(`INSERT into customer (cst_name,cst_dob,cst_phoneNum,cst_email) values (?,?,?,?)
	RETURNING cst_id`, data.Customer.Name, data.Customer.DOB, data.Customer.PhoneNum, data.Customer.Email)
	err = row.Scan(&data.Customer.ID)
	if err != nil {
		return err
	}

	for _, fam := range data.FamMember {
		row := tx.QueryRow(`INSERT into family_list (cst_id,fl_relation,fl_name,fl_dob) values (?,?,?,?)
		RETURNING fl_id`, fam.CustomerID, fam.Relation, fam.Name, fam.DOB)
		err := row.Scan(&fam.ID)
		if err != nil {
			return err
		}
	}

	return nil
}
