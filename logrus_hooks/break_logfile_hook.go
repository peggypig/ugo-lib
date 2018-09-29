package logrus_hooks

import (
	"github.com/gorhill/cronexpr"
	"github.com/sirupsen/logrus"
	"os"
	"sync"
	"time"
)

/**
* @author: codezhang
*
* @create: 2018-07-13 17:24
**/

// 分割文件的HOOK

type BreakLogfileHook struct {
	BaseFilename  string               //基础日志文件名称
	MaxAge        time.Duration        //最大保存时间
	DeleteFile    bool                 //设置是否删除文件
	CronExpr      string               //切割周期表达式
	lastSplitTime time.Time            //上次分割时间
	file          *os.File             //文件
	logFiles      map[string]time.Time //记录本次启动生成的日志文件信息  key 文件名，包括路径；value 生成时间+保存时间
	flag          bool                 //用于标记协程是否打开
	notice        chan int             //用于通知打开新的分割文件的协程
	mutexFile     sync.RWMutex         //读写锁
	mutexFire     sync.Mutex           // fire锁
}

func getNewLogfileName(hook *BreakLogfileHook) string {
	now := time.Now().Local()
	//每一次获取新的日志文件名就认为是一次分割
	hook.lastSplitTime = now
	fileName := hook.BaseFilename + POINT + now.Format(TIME_FORMAT)
	if hook.logFiles == nil {
		hook.logFiles = make(map[string]time.Time)
	}
	hook.mutexFile.Lock()
	// 记录本次文件
	hook.logFiles[fileName] = now.Add(hook.MaxAge)
	hook.mutexFile.Unlock()
	return fileName
}

func openFile(hook *BreakLogfileHook, entry *logrus.Entry) (err error) {
	tempFile, err := os.OpenFile(getNewLogfileName(hook), os.O_CREATE|os.O_WRONLY|os.O_APPEND, FILE_MODAL)
	if err != nil {
		//出错就假定分割和删除日志的进程已经打开
		hook.flag = true
		//出错时设置logrus的输出为std
		entry.Logger.Out = os.Stdout
	} else {
		err = nil
		hook.file = tempFile
		entry.Logger.Out = hook.file
		hook.notice <- NOTICE
	}
	return err
}

func (hook *BreakLogfileHook) Fire(entry *logrus.Entry) (err error) {
	hook.mutexFire.Lock()
	defer hook.mutexFire.Unlock()
	if hook.file == nil { //如果文件还没有打开
		hook.flag = false
		hook.notice = make(chan int, 1)
		//打开文件
		err = openFile(hook, entry)
	}
	if err == nil { //不出错才尝试打开日志文件处理协程
		if hook.flag == false {
			defer func(h *BreakLogfileHook) {
				if recover() != nil {
					h.flag = false
				}
			}(hook)
			hook.flag = true
			go func(h *BreakLogfileHook, en *logrus.Entry) {
				for {
					<-h.notice
					now := time.Now().Local()
					nextTime := cronexpr.MustParse(h.CronExpr).Next(now)
					timer := time.NewTimer(nextTime.Sub(now))
					<-timer.C
					//首先关闭文件
					h.file.Close()
					//新打开文件
					openFile(h, en)
				}
			}(hook, entry)

			go func(h *BreakLogfileHook, en *logrus.Entry) {
				for {
					if h.DeleteFile == true {
						//检查日志文件
						var dels []string
						hook.mutexFile.Lock()
						for name, value := range h.logFiles {
							if value.Before(time.Now().Local()) {
								dels = append(dels, name)
							}
						}
						for i := 0; i < len(dels); i++ {
							errRemove := os.Remove(dels[i])
							if errRemove == nil {
								delete(h.logFiles, dels[i])
							}
						}
						hook.mutexFile.Unlock()
						time.Sleep(time.Second)
					} else {
						break
					}
				}
			}(hook, entry)

		}
	}
	return err
}

func (hook *BreakLogfileHook) Levels() []logrus.Level {
	return logrus.AllLevels
}
