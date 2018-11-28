package model

import (
	"crypto/rand"
	"crypto/sha1"
	"fmt"
	"log"
	"time"

	"github.com/nomnom0452/potatolshop/data"
)

type User struct {
	Id       int    `json:"ID"`
	Name     string `json:"Name"`
	Age      int    `json:"Age"`
	Email    string `json:"Email"`
	Password string `json:"Password"`
}

func (u *User) CreateSession() (session Session, err error) {
	uuid := createUUID()

	_, err = data.Db.Exec("INSERT INTO sessions(Uuid, Email, UserId, CreateAt) VALUES(?, ?, ?, ?)", uuid, u.Email, u.Id, time.Now())

	if err != nil {
		fmt.Println("Fail to insert session to presistent ", err)
		return
	}

	session, err = SessionGET(uuid)

	return
}

func (u *User) Compare(password string) bool {
	if u.Password != Encrypt(password) {
		return false
	}
	return true
}

func CustByEmail(email string) (user User, err error) {
	row := data.Db.QueryRow("SELECT * FROM customers WHERE email = ?", email)

	err = row.Scan(&user.Id, &user.Name, &user.Age, &user.Email, &user.Password)

	return
}

func CreateCust(u User) (err error) {
	Encrypt(u.Password)
	_, err = data.Db.Exec("INSERT INTO customers(Name, Age, Email, Password) VALUES(?, ?, ?, ?)",
		u.Name, u.Age, u.Email, Encrypt(u.Password))

	return
}

func createUUID() (uuid string) {
	u := new([16]byte)
	_, err := rand.Read(u[:])
	if err != nil {
		log.Fatalln("Cannot generate UUID", err)
	}

	// 0x40 is reserved variant from RFC 4122
	u[8] = (u[8] | 0x40) & 0x7F
	// Set the four most significant bits (bits 12 through 15) of the
	// time_hi_and_version field to the 4-bit version number.
	u[6] = (u[6] & 0xF) | (0x4 << 4)
	uuid = fmt.Sprintf("%x-%x-%x-%x-%x", u[0:4], u[4:6], u[6:8], u[8:10], u[10:])
	return
}

func Encrypt(plaintext string) (cryptext string) {
	cryptext = fmt.Sprintf("%x", sha1.Sum([]byte(plaintext)))
	return
}

type Session struct {
	Id       int       `json:"ID"`
	Uuid     string    `json: "Uuid"`
	Email    string    `json: "Email"`
	UserId   int       `json: "UserId"`
	CreateAt time.Time `json: "CreateAt"`
}

func SessionGET(uuid string) (session Session, err error) {
	row := data.Db.QueryRow("SELECT * FROM sessions WHERE Uuid = ?", uuid)

	err = row.Scan(&session.Id, &session.Uuid, &session.Email, &session.UserId, &session.CreateAt)

	return
}
