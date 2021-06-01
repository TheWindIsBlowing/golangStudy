package main

// 工程目录要在 GOPATH 中进行配置
// 并且要按默认的路径进行创建 bin pkg src 导入的包才能找到 src 下放go工程

import(
	"5_func_init/pkg1"
	pack2 "5_func_init/pkg2" // 导入pgk2这个包并起一个别名pack2
	

	_ "5_func_init/pkg3" // 导入pgk3这个包并起一个匿名，此时无法使用这个包，但是导入时会执行该包的init方法
	// . "5_func_init/pkg3" // 导入pgk3这个包 导入到当前包的作用域，可直接使用被导入的包中的方法  不用pkg3.API来调用
)
// 引入包时 将先调用引入包的init 方法

func main() {
	pkg1.TestPkg1()
	// pkg2.TestPkg2()
	pack2.TestPkg2()
}