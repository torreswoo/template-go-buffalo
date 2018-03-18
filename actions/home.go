package actions

import (
	"github.com/gobuffalo/buffalo"
	"fmt"
)

// HomeHandler is a default handler to serve up
// a home page.
func HomeHandler(c buffalo.Context) error {

	fmt.Println(" - context | app : ", 				c.Value("app") )
	fmt.Println(" - context | current_path: ", c.Value("current_path") )
	fmt.Println(" - context | routes: ",       c.Value("routes") )

	return c.Render(200, r.HTML("index.html"))
}