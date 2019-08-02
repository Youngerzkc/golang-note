/* package main

import (
	"fmt"
)

func main() {
	fmt.Println("")

	// 前景 背景 颜色
	// ---------------------------------------
	// 30  40  黑色
	// 31  41  红色
	// 32  42  绿色
	// 33  43  黄色
	// 34  44  蓝色
	// 35  45  紫红色
	// 36  46  青蓝色
	// 37  47  白色
	//
	// 代码 意义
	// -------------------------
	//  0  终端默认设置
	//  1  高亮显示
	//  4  使用下划线
	//  5  闪烁
	//  7  反白显示
	//  8  不可见

	conf := 0  // 配置、终端默认设置
	bg := 0    // 背景色、终端默认设置
	text := 31 // 前景色、红色
	fmt.Printf("%c[%d;%d;%dm%s%c[0m\n", 0x1B, conf, bg, text, "testPrintColor", 0x1B)

	// for b := 40; b <= 47; b++ { // 背景色彩 = 40-47
	// 	for f := 30; f <= 37; f++ { // 前景色彩 = 30-37
	// 		for d := range []int{0, 1, 4, 5, 7, 8} { // 显示方式 = 0,1,4,5,7,8
	// 			fmt.Printf(" %c[%d;%d;%dm%s(f=%d,b=%d,d=%d)%c[0m ", 0x1B, d, b, f, "", f, b, d, 0x1B)
	// 			if d == 0 {
	// 				break
	// 			}
	// 		}
	// 		fmt.Println("")
	// 		if f == 31 {
	// 			break
	// 		}
	// 	}
	// 	fmt.Println("")
	// 	if b == 40 {
	// 		break
	// 	}
	// }
	red := string([]byte{27, 91, 57, 55, 59, 52, 49, 109})
	fmt.Printf("%s-%s\n", red, "debug")
}
*/
// green   = string([]byte{27, 91, 57, 55, 59, 52, 50, 109})
// white   = string([]byte{27, 91, 57, 48, 59, 52, 55, 109})
// yellow  = string([]byte{27, 91, 57, 55, 59, 52, 51, 109})
// red     = string([]byte{27, 91, 57, 55, 59, 52, 49, 109})
// blue    = string([]byte{27, 91, 57, 55, 59, 52, 52, 109})
// magenta = string([]byte{27, 91, 57, 55, 59, 52, 53, 109})
// cyan    = string([]byte{27, 91, 57, 55, 59, 52, 54, 109})
// reset   = string([]byte{27, 91, 48, 109})
package main

import "os"

// import "fmt"

func main() {
	os.Stdout.WriteString("hello, world\n")
	f, _ := os.OpenFile("test", os.O_CREATE|os.O_WRONLY, 0666)
	red := string([]byte{27, 91, 57, 55, 59, 52, 49, 109})
	reset := string([]byte{27, 91, 48, 109}) //着色尾端
	defer f.Close()
	f.WriteString(red + "hello, world  a file\n" + reset + "beijing")
}
