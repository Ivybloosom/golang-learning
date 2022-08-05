// 功能说明
/*
* @Author: Ivy
* @Date: 2022/7/21 12:26 AM
 */
package model_etc

import "github.com/globalsign/mgo/bson"

type User struct {
	ID   bson.ObjectId `bson:"_id"`
	Name string        `bson:"name"`
}
