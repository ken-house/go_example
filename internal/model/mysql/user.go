package mysql

type User struct {
	Id       int    `json:"id" xorm:"not null pk autoincr INT(10)"`
	Username string `json:"username" xorm:"default '' VARCHAR(100)"`
	Password string `json:"password" xorm:"default '' VARCHAR(30)"`
	Gender   int    `json:"gender" xorm:"default 0 SMALLINT(5)"`
}
