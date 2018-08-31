package validation

import (
	"fmt"
	"testing"
	)

/**
*
* @description: 
*
* @author: codezhang
*
* @create: 2018-08-30 19:20
**/
type Person struct {
	Mail  string  `valid:"Mail();Require()"`
	Age   int     `valid:"Range(10,20)"`
	Money float32 `valid:"Range(10.1,20.2)"`
}

func Test_Validator(t *testing.T) {
	errs := Validator(Person{
		Mail:  "79911@qq.com",
		Age:   15,
		Money: 17,
	})
	fmt.Println("errs:", errs)

}
