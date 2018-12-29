package mongo

/**
*
* @description:
*
* @author: codezhang
*
* @create: 2018-10-23 10:20
**/
type student struct {
	StudentCode string `bson:"StudentCode"`
	StudentName string `bson:"StudentName"`
}

type class struct {
	ClassCode string    `bson:"ClassCode"`
	ClassName string    `bson:"ClassName"`
	Students  []student `bson:"Students"`
}

