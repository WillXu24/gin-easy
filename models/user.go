package models

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"time"
)

/***************************************************** Model **********************************************************/
type User struct {
	ID        primitive.ObjectID `json:"id" bson:"_id"`
	CreatedAt int64              `json:"created_at" bson:"created_at"`
	DeleteAt  int64              `json:"delete_at" bson:"delete_at"`
	// status字段可自定义，默认0-正常用户、1-删除账户
	Status int `json:"status" bson:"status"`

	Username  string `json:"username" bson:"username"`
	Password  string `json:"password" bson:"password"`
	Telephone string `json:"telephone" bson:"telephone"`
	Email     string `json:"email" bson:"email"`
}

/***************************************************** Simple *********************************************************/
// 简单封装层，可根据业务需求复杂度选择是否进行简单封装
func UserUpdateOneOfSet(filter bson.M, update bson.M) (*mongo.UpdateResult, error) {
	res, err := UserUpdateOne(filter, bson.M{"$set": update}, nil)
	return res, err
}

/***************************************************** Basic **********************************************************/
func UserInsertOne(msg User) (*mongo.InsertOneResult, error) {
	res, err := UserColl.InsertOne(context.Background(), bson.M{
		"created_at": time.Now().Unix(),
		"delete_at":  time.Now().Unix(),
		"status":     0,

		"username":  msg.Username,
		"password":  msg.Password,
		"telephone": msg.Telephone,
		"email":     msg.Email,
	})
	return res, err
}

func UserFindOne(filter bson.M) (User, error) {
	var msg User
	err := UserColl.FindOne(context.Background(), filter).Decode(&msg)
	return msg, err
}

func UserFindMany(filter bson.M, options *options.FindOptions) ([]User, error) {
	ctx := context.Background()
	cursor, err := UserColl.Find(ctx, filter, options)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)
	var res []User
	for cursor.Next(ctx) {
		var temp User
		if err := cursor.Decode(&temp); err != nil {
			return nil, err
		}
		res = append(res, temp)
	}
	return res, nil
}

func UserUpdateOne(filter bson.M, update bson.M, options *options.UpdateOptions) (*mongo.UpdateResult, error) {
	res, err := UserColl.UpdateOne(context.Background(), filter, update, options)
	return res, err
}

func UserUpdateMany(filter bson.M, update bson.M, options *options.UpdateOptions) (*mongo.UpdateResult, error) {
	res, err := UserColl.UpdateMany(context.Background(), filter, update, options)
	return res, err
}

func UserDeleteOne(filter bson.M) (*mongo.DeleteResult, error) {
	res, err := UserColl.DeleteOne(context.Background(), filter)
	return res, err
}

func UserDeleteMany(filter bson.M) (*mongo.DeleteResult, error) {
	res, err := UserColl.DeleteMany(context.Background(), filter)
	return res, err
}
