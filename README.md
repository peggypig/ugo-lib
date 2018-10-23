# 项目介绍
ugo-lib是go语言的“轮子库”，包括一些常用的功能性组件，会不定期的更新项目内容。
***
## logrus_hooks 组件
>>这个组件基于logrus日志框架，提供了一些常用的与日志相关的功能性钩子。
>>目前包括：

>>### BreakLogfileHook
>>>>```Go
>>>>    BaseFilename      string          //基础日志文件名称
>>>>    MaxAge            time.Duration   //最大保存时间
>>>>    DeleteFile        bool            //设置是否删除文件
>>>>    CronExpr          string          //切割周期表达式
>>>>```
>>>>#### 日志文件切割:<br>
>>>> 可以自定义日志文件的名称，但是实际日志文件名称会在BaseFilename的基础上加上时间后缀。<br>
可以设置日志文件的最大保存时间，但是这个设置仅在DeleteFile（默认是false）为true，即定删除日志文件的情况下生效。<br>
可以设置日志文件的切割周期表达式（ps:使用表达式的想法来自于crontab，[关于CronExpr的编写详见](https://github.com/gorhill/cronexpr)）<br>
>>>>#### QuickStart
>>>>```Go
>>>>	log := logrus.New()
>>>>	log.Hooks.Add(&BreakLogfileHook{
>>>>		BaseFilename: "./test.log",      // 设置日志文件基础名称  包括路径  手动建立文件夹路径
>>>>		MaxAge:       25 * time.Second,  // 设置日志文件最大存在时长
>>>>		CronExpr:     "*/5 * * * * * *", // cron表达式，每五秒钟切割一次
>>>>		DeleteFile:   true,              // 设置是否删除日志文件 默认false不删除，不删除的情况下，MaxAge的值不生效
>>>>	})
>>>>	go func() {
>>>>		for {
>>>>			log.Info("aaa")
>>>>			time.Sleep(time.Second)
>>>>		}
>>>>	}()
>>>>	time.Sleep(time.Second * 2)
>>>>```

>>### FieldHook
>>>>```Go
>>>>	HookLevels []logrus.Level //添加字段的日志等级，默认为全部
>>>>	Fields     logrus.Fields
>>>>```
>>>>#### 日志自定义字段:<br>
>>>>根据日志的不同等级（默认是全部）设置一些全局的日志字段，比如，每条日志需要增加请求API的用户ID，或每条记录想增加请求到达的时间etc.<br>
>>>>#### QuickStart
>>>>```Go
>>>>	log := logrus.New()
>>>>	for i := 0; i < 10; i++ {
>>>>		log.Info("aaa")
>>>>	}
>>>>	log.Hooks.Add(&FieldHook{
>>>>		HookLevels: logrus.AllLevels,   //设置hook的生效日志等级
>>>>		Fields: map[string]interface{}{ //添加Key:value
>>>>			"UserName": "codezhang",
>>>>			"Action":   "test",
>>>>		},
>>>>	})
>>>>	for i := 0; i < 10; i++ {
>>>>		log.Info("aaa")
>>>>	}
>>>>```

>>### MailHook
>>>>```Go
>>>>	Sender     string         //发送者
>>>>	SenderPass string         //授权码
>>>>	SmtpServer string         //邮件服务器地址
>>>>	SmtpPort   int            //邮件服务器端口
>>>>	Receivers  string         //接受者 ,以';'间隔
>>>>	Title      string         //邮件主题
>>>>	Content    interface{}    //邮件内容
>>>>	HookLevels []logrus.Level //发送邮件的日志等级
>>>>	MailType   MailType       //邮件类型
>>>>```
>>>>#### 日志邮件报警:<br>
>>>>设置邮件相关配置，并且可以设置报警的日志等级。
>>>>#### QuickStart
>>>>```Go
>>>>	log := logrus.New()
>>>>	log.Hooks.Add(&MailHook{
>>>>		Receivers:  "*****@qq.com;*************@qq.com", //设置接受邮件，多个邮件使用;间隔
>>>>		Sender:     "**********@qq.com",                 //设置发送邮件
>>>>		SenderPass: "****************",                  //设置授权码
>>>>		SmtpServer: "smtp.qq.com",                       //smtp服务器地址
>>>>		SmtpPort:   587,                                 //smtp服务器端口
>>>>		Title:      "日志报错",                            //设置邮件主题
>>>>		MailType:   MAIL_TYPE_TEXT,                      //设置邮件内容类型
>>>>	})
>>>>	log.Error("aaa")
>>>>	time.Sleep(time.Second * 10)
>>>>```
***
## validation 组件
>>>> 这个组件是用于结构体“简单”属性（基础数据类型属性）的值校验，参数校验。<br>
<em>高度可扩展:</em>使用插件的形式，不满足使用需求时，只需要自定义func，并且注册到validator_plugin中即可。
>>>>```Go
>>>>type validatorFunc func(rule string ,msg string,fieldName string, value reflect.Value,kind reflect.Kind)(error)
>>>>```
>>>>其中rule是定义的规则，比如正则，最大最小值等等，msg是错误提示消息，fieldName是当前正在校验的属性的字段名，value是当前属性的值，kind是当前属性的类型。<br>
校验器格式：
>>>>```Go
>>>>    Validator(rules):errorInfo
>>>>```
>>>>其中，Validator是校验器名称，如Mail,Range等等；rules是可选的校验规则参数，如正则表达式；errorInfo是可选的错误提示消息。

>>>>#### 内置校验器：
>>>>|    校验器                  | 说明 |
>>>>| --------------------------| --- |
>>>>| Mail(\[*\])\[:errorInfo\] |  邮箱校验，*是可选的自定义邮箱正则 |
>>>>| Range(min,max)            |  数值类型的取值范围校验，min、max的数据类型需要和字段的数据类型一致 |
>>>>| MaxLength(int)            | string类型的最大长度校验 |
>>>>| MinLength(int)            | string类型的最小长度校验 |
>>>>| Length(int)               | string类型的固定长度校验 |
>>>>| Regexp(*)                 | string类型的正则校验，*是正则表达式|
>>>>| Require()                 | 非nil和字符串非空串校验|

>>>>#### QuickStart：
>>>>```Go
>>>> // 定义一个Person结构体
>>>> type Person struct {
>>>>   	Mail  string  `valid:"Mail();Require()"`
>>>>   	Age   int     `valid:"Range(10,20)"`
>>>>   	Money float32 `valid:"Range(10.1,20.2)"`
>>>>   }
>>>>```
>>>>```Go
>>>>errs := Validator(Person{
>>>>    	Mail:  "79911@qq.com",
>>>>    	Age:   15,
>>>>    	Money: 17,
>>>>    })
>>>>    fmt.Println("errs:", errs)
>>>>```

>>>>#### 自定义校验器插件：
>>>>自定义一个int属性非0判断校验器IsZero：
>>>>1. 校验插件
>>>>```Go
>>>>/**
>>>> rule是自定义规则，可选
>>>> msg是自定义错误消息
>>>> fieldName当前校验的字段名称
>>>> value 当前校验字段的值
>>>> */
>>>>func isZeroValidatorPlugin(rule string, msg string, fieldName string, value reflect.Value, kind reflect.Kind) error {
>>>>	var err error = nil
>>>>	if len(msg) <= 0 {
>>>>    	// 使用默认的消息
>>>>    	msg = "value of "+fieldName+" can not be 0"
>>>>    }
>>>>    if kind == reflect.int {
>>>>    	// 只针对int
>>>>    	if len(value.Int())== 0 {
>>>>    		err = errors.New(msg)
>>>>    	}
>>>>    }else {
>>>>    	panic("valid:IsZero() tag for field with int type,but filed named "+fieldName+" is "+kind.String())
>>>>  	}
>>>>	return err
>>>>}
>>>>```
>>>>2. 注册校验插件
>>>>```Go
>>>> RegisterValidatorPlugin("IsZero",isZeroValidatorPlugin)
>>>>```
***
## db 组件
>>>>#### mongo:<br>
>>>>针对mongoDb封装的curd操作，Select、Count、Insert、Update、Delete。<br>
>>>>demo详见db/mongo组件下的*_test.go。
