package funny

import (
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

