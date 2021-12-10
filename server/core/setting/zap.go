package setting

type Zap struct {
	Level         string // 日志级别
	Format        string // 格式化
	Prefix        string // 前缀
	Director      string // 日志文件夹
	LinkName      string // 软连接名称
	ShowLine      bool   // 是否显示行
	EncodeLevel   string // 格式化
	StacktraceKey string // 栈名
	LogInConsole  bool   // 输出到控制台
	LogMaxSize    int    // 日志滚动大小
	LogMaxAge     int    // 日志存储时间
	LogMaxBackups int    // 日志备份时间
}
