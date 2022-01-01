// redis里面有bitmap这种数据类型，想着golang如何实现以下
// 以位为粒度进行数据存储，一个uint32占4字节，也就是32位。 如果你的业务能够以位存储那不就节约了32倍内存
//几行代码就能在golang里面进行位粒度的运算业务了

package funny

import (
	"fmt"
)

//位图
type BitMap struct {
	bits []byte
	max  int
}

//初始化一个BitMap
//一个byte有8位,可代表8个数字,取余后加1为存放最大数所需的容量
func NewBitMap(max int) *BitMap {
	bits := make([]byte, (max>>3)+1)
	return &BitMap{bits: bits, max: max}
}

//添加一个数字到位图
//计算添加数字在数组中的索引index,一个索引可以存放8个数字
//计算存放到索引下的第几个位置,一共0-7个位置
//原索引下的内容与1左移到指定位置后做或运算
func (b *BitMap) Add(num uint) {
	index := num >> 3// index=num/8  相当于计算这个num是落在 max二进制的第几字节的  二进制席位
	pos := num & 0x07   //0x07 二进制  0000 0111,   经&计算无论num有多大，都只有num二进制的前三位才可能存在1，及该操作的值范围刚好是2^3个枚举值，且超过部分是求余规律
	b.bits[index] |= 1 << pos//  1 << pos 这个其实就是2的N次方
}

//判断一个数字是否在位图
//找到数字所在的位置,然后做与运算
func (b *BitMap) IsExist(num uint) bool {
	index := num >> 3
	pos := num & 0x07
	return b.bits[index]&(1<<pos) != 0
}

//删除一个数字在位图
//找到数字所在的位置取反,然后与索引下的数字做与运算
func (b *BitMap) Remove(num uint) {
	index := num >> 3
	pos := num & 0x07
	b.bits[index] = b.bits[index] & ^(1 << pos)
}

//位图的最大数字
func (b *BitMap) Max() int {
	return b.max
}

func (b *BitMap) String() string {
	return fmt.Sprint(b.bits)
}

