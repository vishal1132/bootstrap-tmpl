package model

// type usertype string

// const (
// 	Recruiter usertype = "r"
// 	Seeker    usertype = "s"
// 	Admin     usertype = "a"
// )

type status string

const (
	Active     status = "a"
	Inactive   status = "i"
	Onboarding status = "o"
	// UnVerified status = "u"
)

type User struct {
	ID          string `json:"id" gorm:"column:id"`
	Email       string `json:"email" gorm:"column:email"`
	FirstName   string `json:"first_name" gorm:"column:first_name"`
	LastName    string `json:"last_name" gorm:"column:last_name"`
	Phone       string `json:"phone" gorm:"column:phone"`
	Status      string `json:"status" gorm:"column:status"`
	CreatedAt   int64  `json:"created_at" gorm:"column:created_at"`
	CountryCode int64  `json:"country_code" gorm:"column:country_code"`
	CompanyID   string `json:"company_id" gorm:"column:company_id"`
	UpdatedAt   int64  `json:"updated_at" gorm:"column:updated_at"`
	ExtraData   []byte `json:"extra_data" gorm:"column:extra_data"`
}

type UserUpdateRequest struct {
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	Phone       string `json:"phone"`
	CountryCode int64  `json:"country_code"`
}

func FillModel(user *User, req *UserUpdateRequest) {
	if user.FirstName != "" {
		user.FirstName = req.FirstName
	}
	if user.LastName != "" {
		user.LastName = req.LastName
	}
	if user.Phone != "" {
		user.Phone = req.Phone
	}
	if user.CountryCode != 0 {
		user.CountryCode = req.CountryCode
	}
}
