package main

import (
	"os"

	"github.com/silverswords/sand/utils/fileserver"
)

func main() {
	wdir, _ := os.Getwd()
	fileserver.StartFileServer(":9573", wdir)
}
