package funny

import (
	"fmt"
	"testing"
	"time"
)

type page struct {
	Name string
}
func(p *page)newObject()(o Object){
	return &page{Name: "dsdsds"}
}

func TestTimeout(t *testing.T) {
	pp:=&page{}
	np:=NewPool(10,pp)
	for i:=0;i<12;i++{
		go func() {
			o,err:=np.Get(time.Second)
			if err!=nil{
				t.Error("get error")
			}
			if pag,ok:=o.(*page);ok{
				t.Log(pag.Name)
				time.Sleep(time.Millisecond)
				np.Close(pag)
			}
		}()
	}
	time.Sleep(time.Second*10)

}

func TestNoSource(t *testing.T) {
	pp:=&page{}
	np:=NewPool(10,pp)
	for i:=0;i<12;i++{
		o,err:=np.Get(time.Millisecond)
		if err!=nil{
			t.Error("get error")
		}
		if pag,ok:=o.(*page);ok{
			t.Log(pag.Name)
		}
	}

}

type ssss struct {
	Fg string
}

func TestNoSource2(t *testing.T) {
	s:=fg()

	fmt.Printf("%p\n%p\n",&s,&s)


}

func fg()chan ssss  {
	var ty chan ssss
	ty = make(chan ssss)

	go func() {
		close(ty)

		close(ty)
		fmt.Printf("sdsdsdsdddsdd")
	}()
	time.Sleep(time.Second)
	gh := make(chan ssss,1)
	hj := gh
	hj2 := gh
	fmt.Printf("%p\n%p\n%p\n%p\n",&gh,&hj,&ty,&hj2)
	return gh
}