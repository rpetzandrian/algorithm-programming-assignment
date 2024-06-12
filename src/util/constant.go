package util

type MenuConfig struct {
	Name                             string
	RouteIdx, ChoiceIdx, UserTypeIdx int
}

const (
	ADMIN_TYPE = "ADMIN_TYPE"
	USER_TYPE  = "USER_TYPE"

	ADMIN_AUTH_MENU      = "ADMIN_AUTH_MENU"
	ADMIN_REGISTER_MENU  = "ADMIN_REGISTER_MENU"
	ADMIN_LOGIN_MENU     = "ADMIN_LOGIN_MENU"
	ADMIN_AUTH_MENU_BACK = "BACK"
	ADMIN_APPROVAL_MENU  = "ADMIN_APPROVAL_MENU"
	ADMIN_APPROVE_MENU   = "ADMIN_APPROVE_MENU"
	ADMIN_LOGOUT         = "ADMIN_LOGOUT"

	USER_AUTH_MENU          = "USER_AUTH_MENU"
	USER_AUTH_REGISTER_MENU = "USER_AUTH_REGISTER_MENU"
	USER_AUTH_LOGIN_MENU    = "USER_AUTH_LOGIN_MENU"
	USER_AUTH_BACK_MENU     = "USER_AUTH_BACK_MENU"

	USER_SUB_MENU            = "USER_SUB_MENU"
	USER_SUB_MENU_SEND_EMAIL = "USER_SUB_MENU_SEND_EMAIL"
	USER_SUB_MENU_INBOX      = "USER_SUB_MENU_INBOX"
	USER_SUB_MENU_LOGOUT     = "USER_SUB_MENU_LOGOUT"
	USER_SUB_MENU_EMAIL_LIST = "USER_SUB_MENU_EMAIL_LIST"

	BACK   = "BACK"
	LOGOUT = "LOGOUT"

	PRINT_STATUS_SUCCESS = "PRINT_STATUS_SUCCESS"
	PRINT_STATUS_ERROR   = "PRINT_STATUS_ERROR"
	PRINT_STATUS_NOTHING = "PRINT_STATUS_NOTHING"
)

func GetMenuConfig(menu string) MenuConfig {
	switch menu {
	// ADMIN ROUTE CONFIG
	case ADMIN_AUTH_MENU:
		return MenuConfig{
			Name:        "ADMIN_AUTH_MENU",
			RouteIdx:    0,
			ChoiceIdx:   -1,
			UserTypeIdx: 0,
		}
	case ADMIN_REGISTER_MENU:
		return MenuConfig{
			Name:        "ADMIN_REGISTER_MENU",
			RouteIdx:    0,
			ChoiceIdx:   1,
			UserTypeIdx: 0,
		}
	case ADMIN_LOGIN_MENU:
		return MenuConfig{
			Name:        "ADMIN_LOGIN_MENU",
			RouteIdx:    0,
			ChoiceIdx:   2,
			UserTypeIdx: 0,
		}
	case ADMIN_AUTH_MENU_BACK:
		return MenuConfig{
			Name:        "ADMIN_AUTH_MENU_BACK",
			RouteIdx:    0,
			ChoiceIdx:   3,
			UserTypeIdx: 0,
		}
	case ADMIN_APPROVAL_MENU:
		return MenuConfig{
			Name:        "ADMIN_APPROVAL_MENU",
			RouteIdx:    1,
			ChoiceIdx:   -1,
			UserTypeIdx: 0,
		}
	case ADMIN_APPROVE_MENU:
		return MenuConfig{
			Name:        "ADMIN_APPROVE_MENU",
			RouteIdx:    1,
			ChoiceIdx:   0,
			UserTypeIdx: 0,
		}
	case ADMIN_LOGOUT:
		return MenuConfig{
			Name:        "ADMIN_LOGOUT",
			RouteIdx:    1,
			ChoiceIdx:   1,
			UserTypeIdx: 0,
		}

	// USER ROUTE CONFIG
	case USER_AUTH_MENU:
		return MenuConfig{
			Name:        "USER_AUTH_MENU",
			RouteIdx:    0,
			ChoiceIdx:   -1,
			UserTypeIdx: 1,
		}
	case USER_AUTH_REGISTER_MENU:
		return MenuConfig{
			Name:        "USER_AUTH_REGISTER_MENU",
			RouteIdx:    0,
			ChoiceIdx:   0,
			UserTypeIdx: 1,
		}
	case USER_AUTH_LOGIN_MENU:
		return MenuConfig{
			Name:        "USER_AUTH_LOGIN_MENU",
			RouteIdx:    0,
			ChoiceIdx:   1,
			UserTypeIdx: 1,
		}
	case USER_SUB_MENU:
		return MenuConfig{
			Name:        "USER_SUB_MENU",
			RouteIdx:    1,
			ChoiceIdx:   -1,
			UserTypeIdx: 1,
		}
	case USER_SUB_MENU_SEND_EMAIL:
		return MenuConfig{
			Name:        "USER_SUB_MENU_SEND_EMAIL",
			RouteIdx:    1,
			ChoiceIdx:   0,
			UserTypeIdx: 1,
		}
	case USER_SUB_MENU_INBOX:
		return MenuConfig{
			Name:        "USER_SUB_MENU_INBOX",
			RouteIdx:    1,
			ChoiceIdx:   1,
			UserTypeIdx: 1,
		}
	case USER_SUB_MENU_LOGOUT:
		return MenuConfig{
			Name:        "USER_SUB_MENU_LOGOUT",
			RouteIdx:    1,
			ChoiceIdx:   2,
			UserTypeIdx: 1,
		}
	case USER_SUB_MENU_EMAIL_LIST:
		return MenuConfig{
			Name:        "USER_SUB_MENU_EMAIL_LIST",
			RouteIdx:    1,
			ChoiceIdx:   3,
			UserTypeIdx: 1,
		}
	}

	return MenuConfig{}
}
