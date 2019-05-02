package db

import "github.com/weifuchuan/GKD/kit"

type Account struct {
	Id       int64  `json:"id"`
	Username string `json:"username" xorm:"not null VARCHAR(100)"`
	Password string `json:"password" xorm:"not null VARCHAR(100)"`
	Avatar   string `json:"avatar" xorm:"null TEXT"`

	// 0 1 2 3 4 5 6 7
	// 1                用户
	//   1              员工
	//     1            ggu
	Status int `json:"status" xorm:"not null"`
}

func (this *Account) IsUser() bool {
	if kit.BitAt(this.Status, 0) == 1 {
		return true
	} else {
		return false
	}
}

type Order struct {
	Id int64 `json:"id"`
}

type Car struct {
	Id int64 `json:"id"`
}
