package config

type LogConfig struct {
	Output string // 日志路径
}

type AuthConfig struct {
	TokenSecret      string // token 秘钥
	AdminTokenSecret string // 管理端 token 秘钥
}

type DBConfig struct {
	Host     string
	User     string
	Password string
	Port     uint32
	Dbname   string
}

type EmailConfig struct {
	Host     string
	Port     int
	Secure   bool
	User     string
	Password string
	From     string // 发件人邮箱
}

type Result struct {
	Env           string      //  环境
	IsDev         bool        //  是否是开发环境
	IsProd        bool        //  是否是生产环境
	Port          int         //  端口
	Log           LogConfig   //  日志
	Auth          AuthConfig  //  用户认证
	Db            DBConfig    //  数据库配置
	Email         EmailConfig //  邮箱配置
	FileUploadDir string      `mapstructure:"file_upload_dir"` //  文件上传路径
}
