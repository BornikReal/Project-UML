package user

type userModel struct {
	ID                 int64   `db:"id"`
	Role               string  `db:"role"`
	Ratings            []int64 `db:"ratings"`
	Posts              []int64 `db:"posts"`
	Comments           []int64 `db:"comments"`
	PrivateMessages    []int64 `db:"private_messages"`
	BlackList          []int64 `db:"black_list"`
	Restrictions       []int64 `db:"restrictions"`
	Username           string  `db:"username"`
	ProfileDescription string  `db:"profile_description"`
	Avatar             string  `db:"avatar"`
	Email              string  `db:"email"`
	Password           string  `db:"password"`
}

type authDataModel struct {
	ID   int64  `db:"id"`
	Role string `db:"role"`
}
