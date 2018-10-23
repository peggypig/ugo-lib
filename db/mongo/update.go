package mongo

/**
*
* @description: 
*
* @author: codezhang
*
* @create: 2018-10-23 10:09
**/


func Update(dbName string, collectionName string, condition map[string]interface{},value interface{}) (err error) {
	copySession := GetCopySession()
	defer copySession.Close()
	err = copySession.DB(dbName).C(collectionName).Update(condition,value)
	return
}

 