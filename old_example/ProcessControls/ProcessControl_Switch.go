// 流程控制之switch
package ProcessControls

import (
	"fmt"
)

func main() {
	switch a := 1; {
	case a >= 0:
		fmt.Println("a >= 0")
	case a >= 1:
		fmt.Println("a >= 1")
	default:
		fmt.Println("default")
	}
}
