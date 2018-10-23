package mongo

import (
	"gopkg.in/mgo.v2"
	"time"
	"log"
)

/**
*
* @description:
*
* @author: codezhang
*
* @create: 2018-10-22 18:38
**/

// dataSource=mongodb://username:password@db1ip:db1port[,db2ip:db2port]
var dataSource string
var session *mgo.Session

func SetDbConfig(DataSource string) {
	dataSource = DataSource
}

// 获取数据库会话指针
func getSession() *mgo.Session {
	if session == nil {
		conn()
	}
	return session
}

// 获取数据库会话复制指针
func GetCopySession() *mgo.Session {
	if session == nil {
		getSession()
	}
	return session.Copy()
}

func conn() {
	defer func() {
		err := recover()
		if err != nil {
			go connCheck()
		}
	}()
	sessionTemp, err := mgo.Dial(dataSource)
	if err == nil {
		session = sessionTemp
		session.SetMode(mgo.Eventual, true)
	}else {
		log.Println(err.Error())
	}
	go connCheck()
}

func connCheck() {
	for {
		time.Sleep(10 * time.Second)
		if session != nil {
			err := session.Ping()
			if err != nil {
				go conn()
				break
			} else {
				go conn()
				break
			}
		} else {
			go conn()
			break
		}
	}
}
