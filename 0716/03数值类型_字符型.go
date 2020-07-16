package main
import (
	"fmt"
	"unsafe"
)

func main() {
	//1.byte
	//占用1个字节8个比特与uint8没有本质区别，表示ASCII表中的一个字符
	var a byte = 'a'
	var b byte = 98
	fmt.Printf("%c , %c \n",a,b)
	//2.rune
	//占用4个字节 存储一个unicode 与uint32本质相同
	var c rune = '码'
	fmt.Printf("%c \n",c)
	//验证
	fmt.Println(unsafe.Sizeof(a),unsafe.Sizeof(c))
	//3.字符串
	//go字符串不可改变,使用utf-8编码，英文占一个字节，中文占三个字节，因为go只保留了一个内存的指针
	//3.1 直接建
	var str01 string = "Daily Go 每日"
	fmt.Println(str01)
	//3.2 从byte数组转
	var str02 = [5]byte{104, 101, 108, 108, 111}
	fmt.Printf("%s",str02)
	//``也可以作为字符串括起来的
}
