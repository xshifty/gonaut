package main

import (
	. "github.com/xshifty/gonaut"
	"github.com/xshifty/gonaut/examples/welcome/controllers"
)

func main() {
	NewBootstrap(func(b *Bootstrap) {
		// b.WithControllers("./examples/welcome/controllers")
		b.RegisterController(controllers.Home{})
	}).Run(42069)
}
