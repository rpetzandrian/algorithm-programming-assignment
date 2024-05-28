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
	ChoiceNumber int
	ChoiceText   string
	ChoiceFunc   func(userTypeIndex *int, routeIndex *int, choiceIndex *int)
}

type Route struct {
	RouteId    int
	RouteName  string
	RouteFunc  func(choiceIndex *int)
	ChoiceList [4]Choice
}

type UserType struct {
	UserType  string
	RouteList [10]Route
}

type USER_LIST [100]User
type USER_ADMIN_LIST [5]UserAdmin
type EMAIL_LIST [1024]Email
type USER_TYPE_LIST [2]UserType
