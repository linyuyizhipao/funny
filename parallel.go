//编排 goroutinue

package funny

import (
	"errors"
	"fmt"
	"runtime"
	"sync"
)

// Parallel 并发执行
func Parallel(fns ...func()) func() {
	var wg sync.WaitGroup
	return func() {
		wg.Add(len(fns))
		for _, fn := range fns {
			go try(fn, wg.Done)
		}
		wg.Wait()
	}
}

// Serial 串行
func Serial(fns ...func()) func() {
	return func() {
		for _, fn := range fns {
			try(fn,nil)
		}
	}
}

//并发执行多个fn，之哟啊有一个错误立马退出
func GoFns(fns ...func()error)(err error)  {
	exitCh := make(chan error)
	var once sync.Once
	exitFunc := func(err error) {
		if err != nil {
			once.Do(func() {
				exitCh <- err
			})
		}
	}
	for _, fn := range fns {
		fn:=fn
		go func() {
			if err:=fn();err!=nil{
				exitFunc(err)
			}
		}()
	}
	err = <-exitCh
	return
}


func try(fn func(), cleaner func()) (err error) {
	if cleaner != nil {
		defer cleaner()
	}
	defer func() {
		_, file, line, _ := runtime.Caller(2)  //抛出错误代码的调用stack往上推2级
		if rErr := recover(); rErr != nil {
			if _, ok := rErr.(error); ok {
				err = errors.New(fmt.Sprintf("%s:%d,err:%v", file, line,rErr.(error)))
			} else {
				err = fmt.Errorf("%+v", rErr)
			}
		}
	}()

	fn()

	return nil
}
