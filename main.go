package main

import "github.com/myidapp/backend/app"

func main(){
	app := &app.App{}
	app.Initialize()
	app.Run(":8000")
}
