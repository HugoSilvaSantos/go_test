
package main
import "fmt"
import "runtime"

func main() {
   fmt.Println("Hello, World!")
   fmt.Printf("OS: %s\nArchitecture: %s\n", runtime.GOOS, runtime.GOARCH)
}