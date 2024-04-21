package models

type CreateCustomerReq struct {
	NationalityID int            `json:"nationality_id" db:"nationality_id"`
	Name          string         `json:"nama" db:"cst_name"`
	DOB           string         `json:"tanggal_lahir" db:"cst_dob"`
	PhoneNum      string         `json:"telepon" db:"cst_phoneNum"`
	Email         string         `json:"email" db:"cst_email"`
	FamMember     []FamilyMember `json:"keluarga"`
}

type UpdateCustomerReq struct {
	NationalityID int            `json:"nationality_id" db:"nationality_id"`
	Name          string         `json:"nama" db:"cst_name"`
	DOB           string         `json:"tanggal_lahir" db:"cst_dob"`
	PhoneNum      string         `json:"telepon" db:"cst_phoneNum"`
	Email         string         `json:"email" db:"cst_email"`
	FamMember     []FamilyMember `json:"keluarga"`
}

type DeleteCustomerReq struct {
	FamId []FamilyIdDelete `json:"keluarga"`
}

type FamilyIdDelete struct {
	ID int `json:"id"`
}

// type Customer struct {
// 	ID            int    `json:"id" db:"cst_id"`
// 	NationalityID int    `json:"nationality_id" db:"nationality_id"`
// 	Name          string `json:"nama" db:"cst_name"`
// 	DOB           string `json:"tanggal_lahir" db:"cst_dob"`
// 	PhoneNum      string `json:"telepon" db:"cst_phoneNum"`
// 	Email         string `json:"email" db:"cst_email"`
// }

// type FamilyMember struct {
// 	ID         int    `json:"id" db:"fl_id"`
// 	CustomerID int    `json:"customer_id" db:"cst_id"`
// 	Relation   string `json:"relation" db:"fl_relation"`
// 	Name       string `json:"name" db:"fl_name"`
// 	DOB        string `json:"dob" db:"fl_dob"`
// }
