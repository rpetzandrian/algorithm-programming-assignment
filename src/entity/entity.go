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

type Choice struct {
	ChoiceText string
	ChoiceFunc func(userTypeIndex *int, routeIndex *int, choiceIndex *int)
}

type Route struct {
	RouteName  string
	RouteFunc  func(printStatus *string, printText *string, choiceIndex *int)
	ChoiceList [4]Choice
}

type UserType struct {
	UserType  string
	RouteList [10]Route
}

type USER_LIST [100]User
type USER_ADMIN_LIST [5]UserAdmin
type EMAIL_LIST [1024]Email

/*
*
- This is a list of user types
 0. For Admin
 1. For User
*/
type USER_TYPE_LIST [2]UserType

/**
* This is a struct for logged user
* Role 0 | 1 => 0 for admin and 1 for user
 */
type LoggedUser struct {
	Id    int
	Name  string
	Email string
	Role  int
}
