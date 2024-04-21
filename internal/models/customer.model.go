package models

type CustomerData struct {
	Nationality Nationality
	Customer    Customer
	FamMember   []FamilyMember
}

type Nationality struct {
	ID   int    `json:"id" db:"nationality_id"`
	Name string `json:"name" db:"nationality_name"`
	Code string `json:"code" db:"nationality_code"`
}

type Customer struct {
	ID            int    `json:"id" db:"cst_id"`
	NationalityID int    `json:"nationality_id" db:"nationality_id"`
	Name          string `json:"nama" db:"cst_name"`
	DOB           string `json:"tanggal_lahir" db:"cst_dob"`
	PhoneNum      string `json:"telepon" db:"cst_phoneNum"`
	Email         string `json:"email" db:"cst_email"`
}

type FamilyMember struct {
	ID         int    `json:"id" db:"fl_id"`
	CustomerID int    `json:"customer_id" db:"cst_id"`
	Relation   string `json:"relation" db:"fl_relation"`
	Name       string `json:"name" db:"fl_name"`
	DOB        string `json:"dob" db:"fl_dob"`
}
