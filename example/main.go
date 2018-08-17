package main

import "github.com/locona/drill"

func main() {
	svc := drill.New("http://localhost:8047")
	svc.Query("SELECT * FROM sys.options")
}
