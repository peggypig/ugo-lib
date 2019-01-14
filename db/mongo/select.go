package mongo

import (
	"github.com/goinggo/mapstructure"
	"github.com/peggypig/ugo-lib/db/common"
	"reflect"
)

/**
*
* @description:
*
* @author: codezhang
*
* @create: 2018-10-22 18:35
**/

func Select(dbName string, collectionName string, condition map[string]interface{}, page common.Page, resultModel interface{}) (results []interface{}, err error) {
	session := GetCopySession()
	defer session.Close()
	var resultMap []map[string]interface{}
	if page.Limit < 0 && page.Offset < 0 {
		err = session.DB(dbName).C(collectionName).Find(condition).All(&resultMap)
	} else if page.Limit < 0 && page.Offset > 0 {
		err = session.DB(dbName).C(collectionName).Find(condition).Skip(page.Offset).All(&resultMap)
	} else if page.Limit > 0 && page.Offset < 0 {
		err = session.DB(dbName).C(collectionName).Find(condition).Limit(page.Limit).All(&resultMap)
	} else {
		err = session.DB(dbName).C(collectionName).Find(condition).Skip(page.Offset).Limit(page.Limit).All(resultMap)
	}
	for _, mapInfo := range resultMap {
		err = mapstructure.Decode(mapInfo, resultModel)
		if err == nil {
			obj := reflect.ValueOf(resultModel).Elem().Interface()
			results = append(results, obj)
		}
	}
	return
}

func Count(dbName string, collectionName string, condition map[string]interface{}) (count int, err error) {
	session := GetCopySession()
	defer session.Close()
	count, err = session.DB(dbName).C(collectionName).Find(condition).Count()
	return
}
