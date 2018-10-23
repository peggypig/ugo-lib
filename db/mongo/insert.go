package mongo

/**
*
* @description:
*
* @author: codezhang
*
* @create: 2018-10-23 10:00
**/

func Insert(dbName string, collectionName string, value ...interface{}) (errs []error) {
	if len(value) > 0 {
		copySession := GetCopySession()
		defer copySession.Close()
		err := copySession.DB(dbName).C(collectionName).Insert(value...)
		if err != nil {
			errs = append(errs, err)
		}
	}
	return
}
