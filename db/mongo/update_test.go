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
* @create: 2018-10-23 11:09
**/

func TestUpdate(t *testing.T) {
	SetDbConfig("mongodb://xxx:xxx@ip:port")
	clazz := class{
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
			{
				StudentCode:"002005",
				StudentName:"钱七",
			},
		},
	}
	fmt.Println(clazz)
	errs := Update("mongotest", "class", map[string]interface{}{
		"ClassCode":"002",
	},clazz)
	fmt.Println(errs)
}