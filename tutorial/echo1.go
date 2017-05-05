// echo1 exibe seus argumentos de linha de comando
package tutorial

import (
	"fmt"
	"os"
	"time"
)

func main() {
	start := time.Now()
	var s, sep string
	for i := 1; i < len(os.Args); i++ {
		s += sep + os.Args[i]
		sep = " "
	}
	fmt.Println(s)
	fmt.Printf("%.6fs seconds.\n", time.Since(start).Seconds())
}
