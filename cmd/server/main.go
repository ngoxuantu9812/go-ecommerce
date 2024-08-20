package server

import (
	"go-ecomm/internal/routers"
)

func main() {
	r := routers.NewRouter()
	err := r.Run(":8080")
	if err != nil {
		return
	}
}
