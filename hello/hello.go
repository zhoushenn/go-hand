package main

import (
	"fmt"
	"greeting"
	"moul.io/banner"
)
import "rsc.io/quote"

func main() {

	fmt.Println(quote.Hello())
	fmt.Println(banner.Inline("hey world."))
	//调用其它模块
	greeting.Hello()

}
