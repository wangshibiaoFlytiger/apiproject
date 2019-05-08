package env

/**
系统环境变量
*/

var SysEnv Env

type Env struct {
	//运行环境:dev,test,pro
	Profile string
}
