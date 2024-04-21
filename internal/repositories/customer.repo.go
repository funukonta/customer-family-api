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
	query := `SELECT cst_id,nationality_id,cst_name,cst_dob,cst_phoneNum,cst_email from customer where cst_id=?`
	row := r.DB.QueryRow(query, id)
	cust := models.Customer{}
	err := row.Scan(&cust.ID, &cust.NationalityID, &cust.Name, &cust.DOB, &cust.PhoneNum, &cust.Email)
	if err != nil {
		return nil, err
	}

	query = `select fl_id,cst_id,fl_relation,fl_name,fl_dob from family_list where cst_id=?`
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
