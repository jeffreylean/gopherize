package main

import (
	"fmt"

	"github.com/gookit/color"
)

func main() {
	fmt.Println("Welcome to gopherize!!")
	gopherize := `
  ________              .__                 .__               
 /  _____/  ____ ______ |  |__   ___________|__|_______ ____  
/   \  ___ /  _ \\____ \|  |  \_/ __ \_  __ \  \___   // __ \ 
\    \_\  (  <_> )  |_> >   Y  \  ___/|  | \/  |/    /\  ___/ 
 \______  /\____/|   __/|___|  /\___  >__|  |__/_____ \\___  >
        \/       |__|        \/     \/               \/    \/ 
    `
	color.Blueln(gopherize)
}
