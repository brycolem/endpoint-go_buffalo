package models

type Application struct {
	ID        int    `json:"id" db:"id"`
	Employer  string `json:"employer" db:"employer"`
	Title     string `json:"title" db:"title"`
	Link      string `json:"link" db:"link"`
	CompanyID int    `json:"companyId" db:"company_id"`
	Notes     []Note `json:"notes" has_many:"notes"`
}
