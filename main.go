package main

import "one-sentence/initRouter"

func main() {
	router := initRouter.SetupRouter()
	_ = router.Run(":1010")
}