package viewUser

type User struct {
	Id       string `json:"id"`
	Admin    int    `json:"admin"`
	Name     string `json:"name"`
	NickName string `json:"nickname"`
	Mail     string `json:"email"`
	Password string `json:"-"`
}
