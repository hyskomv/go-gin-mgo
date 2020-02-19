package user

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/mongodb/mongo-go-driver/bson/primitive"
	"github.com/mongodb/mongo-go-driver/mongo"
	"go-gin-mgo/database"
	"golang.org/x/crypto/bcrypt"
)

const (
	collection = "users"
)

type UserCredentials struct {
	Username string `form:"username" json:"username" binding:"required"`
	Password string `form:"password" json:"password" binding:"required"`
	Admin bool `form:"admin" json:"admin"`
}

type Token struct {
	AccessToken string `json:"accessToken" bson:"accessToken"`
	ExpiresIn int64 `json:"expiresIn" bson:"expiresIn"`
}

type User struct {
	Id primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	Username string `json:"username" bson:"username"`
	Password []byte `json:"-" bson:"password"`
	Token Token `json:"token" bson:"token"`
	Admin bool `json:"admin" bson:"admin"`
	CreatedAt primitive.DateTime `json:"createdAt" bson:"createdAt"`
	UpdatedAt primitive.DateTime `json:"updatedAt" bson:"updatedAt"`
	dao *UsersDao `json:"-" bson:"-"`
}

func init()  {
	database.Model(&User{})
}

func (u *User) SignIn(password string) (*User, error) {
	var user *User

	//shaPass := sha512.Sum512([]byte(password))
	err := bcrypt.CompareHashAndPassword(u.Password, []byte(password))

	return user, err
}

func (u *User) SignOut() error {
	_, err := Dao.UpdateToken(u.Id, nil)
	return err
}

func (u *User) UpdatePassword(password string) error {
	pass, err := u.encryptPassword(password)
	if err != nil { return err }

	_, err = Dao.UpdatePassword(u.Id, pass)

	return err
}

func (u *User) encryptPassword(password string) ([]byte, error) {
	return bcrypt.GenerateFromPassword([]byte(password), 64)
}

func (u *User) makeToken(secret string) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)
	return token.SignedString([]byte(secret))
}

func (u *User) InitModel(d *mongo.Database)  {
	// init DAO
	Dao = &UsersDao{d.Collection(collection)}
}
