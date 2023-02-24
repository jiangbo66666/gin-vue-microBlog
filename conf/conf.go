package conf

// host: 127.0.0.1:3306
// port: 3306
// userName: root
// passWord: imm62611
// timeOut: 1000
// 读取ymal要映射到最外层，最外层也要在项目下面
type Config struct {
	DataBase DataBase `ymal:"dataBase"`
}

type DataBase struct {
	Host     string `ymal:"host"`
	Port     string `ymal:"port"`
	UserName string `ymal:"userName"`
	PassWord string `ymal:"passWord"`
	TimeOut  uint   `ymal:"timeOut"`
}
