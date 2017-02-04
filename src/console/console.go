package console

import (
	"flag"
	"fmt"
	"path/filepath"
	"runtime"

	"github.com/golang/glog"
)

/*

   -logtostderr=false	  //日志写入标准错误而不是文件.
   -alsologtostderr=false //日志写入标准错误以及文件.
   -stderrthreshold=ERROR // Info , Warning  , Error , 设置级别以上的才会输出
   -log_dir=""			  // 日志输出目录
   -log_backtrace_at=""	  // 追踪的this.Debugtrace() , 已经实现动态追踪
   -v=0					  // 日志等级, 小于等于此等级的日志都会被显示 , 具体等级可以使用glog.V(1).Infoln("xxxxx") , glog.V(2).Infoln("xxxxx") , ...

   -vmodule=""
	   The syntax of the argument is a comma-separated list of pattern=N,
	   where pattern is a literal file name (minus the ".go" suffix) or
	   "glob" pattern and N is a V level. For instance,
	   -vmodule=gopher*=3
	   sets the V level to 3 in all Go files whose names begin "gopher".
*/
func init() {
	flag.Set("log_dir", "../logs")       // 设置日志存储目录
	flag.Set("v", "3")                   // 设置日志数字等级	与日志级别不一样 , 小于等于此等级的日志都会被显示
	flag.Set("logtostderr", "true")      // 日志写入标准错误而不是文件	, false 只写入到文件中 , true 标准输出不写文件
	flag.Set("alsologtostderr", "false") // 设置日志输出格式 , true终端和文件都输出 ,   false 只输出到文件中
	flag.Set("stderrthreshold", "Info")  // Info , Warning  , Error , Fatal , 设置级别以上的才会输出
	flag.Parse()
}

// 动态追踪trace
func Debugtrace() bool {
	pc, file, line, ok := runtime.Caller(0) // line 为当前行号
	if !ok {
		glog.Error("runtime.Caller() error.")
	}
	basename := filepath.Base(file) // 获取文件名 line +1 行
	f := runtime.FuncForPC(pc)      // line + 2 行
	flag.Set("log_backtrace_at", fmt.Sprintf("%s:%d", basename, line+8))
	flag.Parse() //使设置的变量生效
	glog.Info("\n------------   " + f.Name() + " --------------\n")
	return true
}
