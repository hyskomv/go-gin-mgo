package user

import (
	"context"
	"github.com/mongodb/mongo-go-driver/bson"
	"github.com/mongodb/mongo-go-driver/bson/primitive"
	"github.com/mongodb/mongo-go-driver/mongo"
)

var Dao *UsersDao

type UsersDao struct {
	*mongo.Collection
}

func (dao *UsersDao) Create(usr *UserCredentials) (*User, error) {
	var user *User

	res, err := dao.InsertOne(context.TODO(), usr)
	if err != nil { return nil, err }

	filter := bson.D{{"_id", res.InsertedID}}
	err = dao.FindOne(context.TODO(), filter).Decode(&user)

	return user, err
}

func (dao *UsersDao) FindByUsername(username string) (*User, error) {
	var user *User

	// find by username
	filter := bson.D{{"username", username}}
	err := dao.FindOne(context.TODO(), filter).Decode(&user)

	return user, err
}

// find by ID
func (dao *UsersDao) FindById(id string) (*User, error) {
	var user *User

	user, err := dao.FindByUsername(id)

	// find by _id
	if err != nil && err == mongo.ErrNoDocuments {
		objId, er := primitive.ObjectIDFromHex(id)
		if er == nil {
			filter := bson.D{{"_id", objId}}
			err = dao.FindOne(context.TODO(), filter).Decode(&user)
		}
	}

	return user, err
}

// find all
func (dao *UsersDao) List(filter interface{}) ([]*User, error) {
	var users []*User

	if filter == nil {
		filter = bson.D{}
	}

	cur, err := dao.Find(context.TODO(), filter)
	defer cur.Close(context.TODO())

	if err != nil {
		return nil, err
	}

	for cur.Next(context.TODO()) {
		var user User

		if err := cur.Decode(&user); err != nil {
			return nil, err
		}

		users = append(users, &user)
	}

	return users, err
}

func (dao *UsersDao) UpdateToken(userId primitive.ObjectID, token *Token) (*User, error) {
	var user *User

	filter := bson.D{{"_id", userId}}
	update := bson.D{
		{"$set", bson.D{
			{"token", token},
		}},
	}

	res, err := Dao.UpdateOne(context.TODO(), filter, update)
	if err == nil {
		filter := bson.D{{"_id", res.UpsertedID}}
		err = dao.FindOne(context.TODO(), filter).Decode(&user)
	}

	return user, err
}

func (dao *UsersDao) UpdatePassword(userId primitive.ObjectID, password []byte) (*User, error) {
	var user *User

	filter := bson.D{{"_id", userId}}
	update := bson.D{
		{"$set", bson.D{
			{"password", password},
		}},
	}

	res, err := Dao.UpdateOne(context.TODO(), filter, update)
	if err == nil {
		filter := bson.D{{"_id", res.UpsertedID}}
		err = dao.FindOne(context.TODO(), filter).Decode(&user)
	}

	return user, err
}
