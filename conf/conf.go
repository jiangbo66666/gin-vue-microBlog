package conf

// host: 127.0.0.1:3306
// port: 3306
// userName: root
// passWord: imm62611
// timeOut: 1000
// 读取yaml要映射到最外层，最外层也要在项目下面
type Config struct {
	DataBase DataBase `yaml:"dataBase"`
}

type DataBase struct {
	Host     string `yaml:"host"`
	Port     string `yaml:"port"`
	UserName string `yaml:"userName"`
	PassWord string `yaml:"passWord"`
	User     string `yaml:"user"`
	TimeOut  int    `yaml:"timeOut"`
	MaxConn  int
	MaxOpen  int
}
