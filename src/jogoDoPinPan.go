// brincadeira do pin e pan.
//era exibir pin para os numeros divisivei

package main

import "fmt"

func main() {

	for i := 1; i <= 100; i++ {

		if i%3 == 0 {
			if i%5 == 0 {
				fmt.Println("Pin Pan")
			} else {
				fmt.Println("Pin")
			}

		} else if i%5 == 0 {
			fmt.Println("Pan")
		} else {
			fmt.Println(i)
		}
	}
}
