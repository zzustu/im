package data

type ImCfg struct {
	ImServer ImServer `ini:"server"`   // 服务器配置
	ImDB     ImDB     `ini:"database"` // 数据库连接配置
	ImBolt   ImBolt   `ini:"bolt"`     // BoltDB配置
	ImLog    ImLog    `ini:"log"`      // 日志配置
}

type ImServer struct {
	Port       int    `ini:"port"`       // http端口号
	ApiPrefix  string `ini:"apiprefix"`  // 后端URL访问前缀
	HtmlPrefix string `ini:"htmlprefix"` // 前端URL访问前缀
	HtmlPath   string `ini:"htmlpath"`   // HTML页面存放位置
	Log        string `ini:"log"`        // 路由日志输出位置
	Mode       string `ini:"mode"`       // gin mode
}

type ImDB struct {
	Host     string `ini:"host"`     // 数据库地址
	Port     int    `ini:"port"`     // 数据库端口号
	Username string `ini:"username"` // 数据库用户名
	Password string `ini:"password"` // 数据库密码
	Schema   string `ini:"schema"`   // 数据库库
	ShowSQL  bool   `ini:"showsql"`  // 是否打印SQL
}

type ImBolt struct {
	Path   string `ini:"path"`   // 持久化文件存放位置
	Bucket string `ini:"bucket"` // bucket名字
}

type ImLog struct {
	Path   string `ini:"path"`   // 日志文件存放位置
	Prefix string `ini:"prefix"` // 日志输出前缀
}
