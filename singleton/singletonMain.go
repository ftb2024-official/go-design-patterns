package singleton

import "fmt"

func SingletonMain() {
	for i := 0; i < 30; i++ {
		go getInstance()
	}

	fmt.Scanln()
}
