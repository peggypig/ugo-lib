package mongo

import (
	"testing"
	"fmt"
)

/**
*
* @description: 
*
* @author: codezhang
*
* @create: 2018-10-23 11:18
**/

func TestDelete(t *testing.T) {
	SetDbConfig("mongodb://xxx:xxx@ip:port")
	//clazz := class{
	//	ClassCode:"003",
	//	ClassName:"初一三班",
	//	Students: []student{
	//		{
	//			StudentCode:"001001",
	//			StudentName:"张三",
	//		},
	//		{
	//			StudentCode:"001002",
	//			StudentName:"李四",
	//		},
	//		{
	//			StudentCode:"001003",
	//			StudentName:"王五",
	//		},
	//		{
	//			StudentCode:"001004",
	//			StudentName:"赵六",
	//		},
	//	},
	//}
	//fmt.Println(Insert("mongotest","class",clazz))
	fmt.Println(Delete("mongotest","class", map[string]interface{}{
		"ClassCode":"003",
	}))
}