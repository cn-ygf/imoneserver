package member

// 会员
type Member struct {
	Id       int64  `json:"id"`
	Email    string `json:"email"`
	Phone    string `json:"phone"`
	Password string `json:"password"`
	Salt     string `json:"salt"`
	Name     string `json:"name"`
	Nick     string `json:"nick"`
	Head     string `json:"head"`
	Regdate  int64  `json:"regdate"`
	Current  string `json:"current"`
	State    int64  `json:"state"`
}

// TableName return table name
func (*Member) TableName() string {
	return "member"
}
