package boot

import (
	"go-gin-mgo/app/user"
	"log"
)

var admin = user.User{Username: "admin"}

func InitServer()  {
	usr, err := user.Dao.Load(admin.Username)

	if err != nil {
		log.Fatal(err)
	}

	//if usr == nil {
	//	user.Dao.Create()
	//}

//	s := database.Session.Clone()
//	defer s.Close()
//	userDao := &user.UsersDAO{s.DB(database.Mongo.Database).C(user.Users)}
//
//
//	var u *user.user
//
//	query := bson.M{"username": bson.ObjectIdHex("admin")}
//
//	u, err := userDao.FindOne(query)
//	if err !=nil { log.Printf(err.Error()) }
//
//	if u == nil {
//
//	}
}
