package repositories

import (
	"context"
	"customer-data-api/internal/models"
	"database/sql"
)

type CustomerRepo interface {
	Create(data *models.CustomerData) error
	GetAllCustomer() ([]models.Customer, error)
	GetCustomer(id int) (*models.CustomerData, error)
	UpdateCustomer(data *models.CustomerData) error
	DeleteCustomer(data *models.DeleteCustomerReq, custID int) error
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

	row := tx.QueryRow(`INSERT into customer (cst_name,cst_dob,cst_phoneNum,cst_email) values ($1,$2,$3,$4)
	RETURNING cst_id`, data.Customer.Name, data.Customer.DOB, data.Customer.PhoneNum, data.Customer.Email)
	err = row.Scan(&data.Customer.ID)
	if err != nil {
		return err
	}

	for _, fam := range data.FamMember {
		row := tx.QueryRow(`INSERT into family_list (cst_id,fl_relation,fl_name,fl_dob) values ($1,$2,$3,$4)
		RETURNING fl_id`, data.Customer.ID, fam.Relation, fam.Name, fam.DOB)
		err := row.Scan(&fam.ID)
		if err != nil {
			return err
		}
	}

	return tx.Commit()
}

func (r *customerRepo) GetAllCustomer() ([]models.Customer, error) {
	query := `SELECT cst_id,nationality_id,cst_name,cst_dob,cst_phoneNum,cst_email from customer`
	rows, err := r.DB.Query(query)
	if err != nil {
		return nil, err
	}

	customer := []models.Customer{}
	for rows.Next() {
		cust := models.Customer{}
		rows.Scan(&cust.ID, &cust.NationalityID, &cust.Name, &cust.DOB, &cust.PhoneNum, &cust.Email)
		customer = append(customer, cust)
	}

	return customer, nil
}

func (r *customerRepo) GetCustomer(id int) (*models.CustomerData, error) {
	query := `SELECT cst_id,nationality_id,cst_name,cst_dob,cst_phoneNum,cst_email from customer where cst_id=$1`
	row := r.DB.QueryRow(query, id)
	cust := models.Customer{}
	err := row.Scan(&cust.ID, &cust.NationalityID, &cust.Name, &cust.DOB, &cust.PhoneNum, &cust.Email)
	if err != nil {
		return nil, err
	}

	query = `select fl_id,cst_id,fl_relation,fl_name,fl_dob from family_list where cst_id=$1`
	fm := []models.FamilyMember{}
	rows, err := r.DB.Query(query, id)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		fam := models.FamilyMember{}
		err = row.Scan(&fam.ID, &fam.CustomerID, &fam.Relation, &fam.Name, &fam.DOB)
		if err != nil {
			return nil, err
		}
		fm = append(fm, fam)
	}

	custData := &models.CustomerData{
		Customer:  cust,
		FamMember: fm,
	}

	return custData, nil
}

func (r *customerRepo) UpdateCustomer(data *models.CustomerData) error {
	query := `UPDATE customer set cst_name=$1,cst_dob=$2,cst_phoneNum=$3,cst_email=$4 where cst_id=$5
	RETURNING cst_name,cst_dob,cst_phoneNum,cst_email`
	tx, err := r.DB.BeginTx(context.Background(), nil)
	if err != nil {
		return err
	}
	defer tx.Rollback()
	err = tx.QueryRow(query, data.Customer.Name, data.Customer.DOB, data.Customer.PhoneNum, data.Customer.Email, data.Customer.ID).Scan(&data.Customer.Name, &data.Customer.DOB, &data.Customer.PhoneNum, &data.Customer.Email)
	if err != nil {
		return err
	}

	query = `UPDATE family_list cst_id=$1,fl_relation=$2,fl_name=$3,fl_dob=$4 where cst_id=$5
	RETURNING cst_id,fl_relation,fl_name,fl_dob`
	for _, fam := range data.FamMember {
		err = tx.QueryRow(query, fam.CustomerID, fam.Relation, fam.Name, fam.DOB).Scan(&fam.CustomerID, &fam.Relation, &fam.Name, &fam.DOB)
		if err != nil {
			return err
		}
	}
	return tx.Commit()
}

func (r *customerRepo) DeleteCustomer(data *models.DeleteCustomerReq, custID int) error {
	query := `DELETE from customer where cst_id=$1`
	tx, err := r.DB.BeginTx(context.Background(), nil)
	if err != nil {
		return err
	}
	defer tx.Rollback()

	_, err = tx.Exec(query, custID)
	if err != nil {
		return err
	}

	query = `DELETE from family_list fl_id=$1`
	for _, fam := range data.FamId {
		_, err = tx.Exec(query, fam.ID)
		if err != nil {
			return err
		}
	}
	return tx.Commit()
}
