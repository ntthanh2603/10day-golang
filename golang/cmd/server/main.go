package main

import "github.com/ntthanh2603/10day-golang.git/golang/internal/routers"

func main() {
	r := routers.NewRouter()

	r.Run(":8080") // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
