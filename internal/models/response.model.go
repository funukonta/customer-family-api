package models

type CreateCustomerRes struct {
	NationalityID int            `json:"nationality_id" db:"nationality_id"`
	ID            int            `json:"id"`
	Name          string         `json:"nama" db:"cst_name"`
	DOB           string         `json:"tanggal_lahir" db:"cst_dob"`
	PhoneNum      string         `json:"telepon" db:"cst_phoneNum"`
	Email         string         `json:"email" db:"cst_email"`
	FamMember     []FamilyMember `json:"keluarga"`
}

type GetAllCustomerRes struct {
	Data []Customer `json:"data"`
}
