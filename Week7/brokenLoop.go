// broken loop closure
package main
import "fmt"
func BLmain() {
	var functions []func() // array of functions
	for i := 0; i < 10; i++ {
		// build an array of functions to call later
		functions = append(functions, func() {
			fmt.Println(i) // just print i
		})
	}
	// now lets call all the function in the array in order
	for _, f := range functions {
		f()
	}
}