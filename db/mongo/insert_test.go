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
* @create: 2018-10-23 10:18
**/
func TestInsert(t *testing.T) {
	SetDbConfig("mongodb://xxx:xxx@ip:port")
	clazz := class{
		ClassCode:"001",
		ClassName:"初一一班",
		Students: []student{
			{
				StudentCode:"001001",
				StudentName:"张三",
			},
			{
				StudentCode:"001002",
				StudentName:"李四",
			},
			{
				StudentCode:"001003",
				StudentName:"王五",
			},
			{
				StudentCode:"001004",
				StudentName:"赵六",
			},
		},
	}
	clazz2 := class{
		ClassCode:"002",
		ClassName:"初一二班",
		Students: []student{
			{
				StudentCode:"002001",
				StudentName:"张三2",
			},
			{
				StudentCode:"002002",
				StudentName:"李四2",
			},
			{
				StudentCode:"002003",
				StudentName:"王五2",
			},
			{
				StudentCode:"002004",
				StudentName:"赵六2",
			},
		},
	}
	fmt.Println(clazz)
	fmt.Println(clazz2)
	errs := Insert("mongotest", "class",clazz,clazz2)
	fmt.Println(errs)
}
