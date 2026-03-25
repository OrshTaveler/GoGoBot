package main

import (
	"fmt"
	"gogobot/internal/api"
	"gogobot/internal/session"
	"gogobot/internal/utils"
)

func main() {
	fmt.Println(`
   ____        ____       ____        _   
  / ___| ___  / ___| ___ | __ )  ___ | |_ 
 | |  _ / _ \| |  _ / _ \|  _ \ / _ \| __|
 | |_| | (_) | |_| | (_) | |_) | (_) | |_ 
  \____|\___/ \____|\___/|____/ \___/ \__|
  _________________________________________                                         
	`, utils.GenerateAuthURL())

	var session session.Session

	api.Router(session)

	fmt.Println(session.Users[0].Username)

}
