package mongo

import (
	"fmt"
	"github.com/peggypig/ugo-lib/db/common"
	"reflect"
	"testing"
)

/**
*
* @description:
*
* @author: codezhang
*
* @create: 2018-10-22 19:21
**/

func TestSelect(t *testing.T) {
	SetDbConfig("mongodb://xxx:xxx@ip:port")
	results, _ := Select("mongotest", "class", map[string]interface{}{
		"Students.StudentName": "张三",
	},
		common.Page{-1, -1}, &class{})
	var clazzs []class
	for _, result := range results {
		if value, ok := result.(class); ok {
			fmt.Println(reflect.TypeOf(value))
			clazzs = append(clazzs, value)
		}
	}
	fmt.Println(clazzs)
}

func TestCount(t *testing.T) {
	SetDbConfig("mongodb://xxx:xxx@ip:port")
	count, _ := Count("mongotest", "class", nil)
	fmt.Println(count)
}
