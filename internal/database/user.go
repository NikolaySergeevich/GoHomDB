package database

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	ID        uuid.UUID `db:"id"`
	Username  string    `db:"username"`
	Password  string    `db:"password"`
	CreatedAt time.Time `db:"created_at"`
	UpdatedAt time.Time `db:"updated_at"`
}

type CreateUserReq struct {
	ID       uuid.UUID
	Username string
	Password string
}

type FindUserCriteria struct {
	ID       *uuid.UUID
	Username *string
}

func (u *User) UserString()string{
	var res string
	var timeCreate string
	var timeUpdate string
	timeCreate = timeCreate + u.CreatedAt.Format("Дата создания: 02.01.2006 15 часов 04 минут 05 сек.")
	timeUpdate = timeUpdate + u.UpdatedAt.Format("Дата обновления: 02.01.2006 15 часов 04 минут 05 сек.")
	res = res + "ПОльзователь " + u.Username + ":" + u.Password + "\n" + timeCreate + "\n" + timeUpdate + "\n"
	return res
}