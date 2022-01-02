package funny

import "sync"

// 返回一个只会执行一次的单例函数
func SingleFn(fn func())(singleFun func())  {
	once:=sync.Once{}
	singleFun = func() {
		once.Do(func() {
			fn()
		})
	}
	return
}