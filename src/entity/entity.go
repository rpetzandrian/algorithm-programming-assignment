package entity

type User struct {
	Id         int
	Name       string
	Email      string
	Password   string
	IsVerified bool
}

type UserAdmin struct {
	Id       int
	Name     string
	Email    string
	Password string
}

type Email struct {
	Id        int
	From      int
	To        int
	Body      string
	IsRead    bool
	Timestamp string
}

type USER_LIST [100]User
type USER_ADMIN_LIST [5]UserAdmin
type EMAIL_LIST [255]Email
