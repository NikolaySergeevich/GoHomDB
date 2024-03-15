package database

import (
	"time"

	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson/primitive"
)
/*
Структура URL адреса. Будет храниться в mongo
*/
type Link struct {
	ID        primitive.ObjectID `bson:"id"`
	Title     string             `bson:"title,omitempty"`
	URL       string             `bson:"url"`
	Images    []string           `bson:"images"`
	Tags      []string           `bson:"tags"`
	UserID    string             `bson:"userID"`
	CreatedAt time.Time          `bson:"created_at"`
	UpdatedAt time.Time          `bson:"updated_at"`
}
/*
Структур Пользователя. Будет храниться в PostgreSQL
*/
type User struct {
	ID        uuid.UUID `db:"id"`
	Username  string    `db:"username"`
	Password  string    `db:"password"`
	CreatedAt time.Time `db:"created_at"`
	UpdatedAt time.Time `db:"updated_at"`
}

func (u *User) UserString()string{
	var res string
	var timeCreate string
	var timeUpdate string
	timeCreate = timeCreate + u.CreatedAt.Format("Дата создания: 02.01.2006 15 часов 04 минут 05 сек.")
	timeUpdate = timeUpdate + u.CreatedAt.Format("Дата обновления: 02.01.2006 15 часов 04 минут 05 сек.")
	res = res + "ПОльзователь " + u.Username + ":" + u.Password + "\n" + timeCreate + "\n" + timeUpdate + "\n"
	return res
}
