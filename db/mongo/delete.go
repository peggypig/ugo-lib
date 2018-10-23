package mongo

/**
*
* @description: 
*
* @author: codezhang
*
* @create: 2018-10-23 10:12
**/

func Delete(dbName string, collectionName string, condition map[string]interface{}) (err error) {
	copySession := GetCopySession()
	defer copySession.Close()
	err = copySession.DB(dbName).C(collectionName).Remove(condition)
	return
}
