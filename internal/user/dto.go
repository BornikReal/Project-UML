package user

var (
	RoleUser        = "user"
	RoleSpecialUser = "user_special"
	RoleModerator   = "moderator"
)

var jwtSecretKey = []byte("very-secret-key")

type CreateUserRequest struct {
	Role               string
	Username           string
	ProfileDescription string
	Avatar             string
	Email              string
	Password           string
}

type ContactInfo struct {
	Username           string
	ProfileDescription string
	Avatar             string
	Email              string
	Password           string
}

type User struct {
	ID              int64
	Role            string
	Ratings         []int64
	Posts           []int64
	Comments        []int64
	PrivateMessages []int64
	BlackList       []int64
	Restrictions    []int64
	ContactInfo     ContactInfo
}

func (u *User) fromModel(model userModel) {
	if u == nil {
		return
	}
	*u = User{
		ID:              model.ID,
		Role:            model.Role,
		Ratings:         model.Ratings,
		Posts:           model.Posts,
		Comments:        model.Comments,
		PrivateMessages: model.PrivateMessages,
		BlackList:       model.BlackList,
		Restrictions:    model.Restrictions,
		ContactInfo: ContactInfo{
			Username:           model.Username,
			ProfileDescription: model.ProfileDescription,
			Avatar:             model.Avatar,
			Email:              model.Email,
			Password:           model.Password,
		},
	}
}

type UpdateUserRequest struct {
	ID                 int64
	Role               *string
	Username           *string
	ProfileDescription *string
	Avatar             *string
	Email              *string
	Password           *string

	AddRatings         []int64
	AddPosts           []int64
	AddComments        []int64
	AddPrivateMessages []int64
	AddBlackList       []int64
	AddRestrictions    []int64
}

type AuthData struct {
	ID   int64
	Role string
}
