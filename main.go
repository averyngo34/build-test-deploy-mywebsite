package main

// import used by the project
import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/handlebars"
)

/* reIndex - renders the index.hbs page
** Input = fiber content
** Output = passes map to layout/main to be rendered / error

*/
func renderIndex(c *fiber.Ctx) error {
	return c.Render("index", fiber.Map{
		"Title":       "Welcome to Avery's World",
		"Japan-Description": "Let's visit Japan",
		"DR-Description": "Let's visit Dominican Republic",
		"Thailand-Description": "Let's visit Thailand",
		"Taiwan-Description": "Let's visit Taiwan",
	}, "layouts/main")
}

/* setupRoutes - abstracted function for all application routes
** Input = pointer to fiber App struct
** Output = null
*/
func setupRoutes(app *fiber.App) {
	app.Get("/", renderIndex) // home route
}

/* main - main function executing the server and handler bar templating
** Input = null
** Output = null / listens on port 3001 for the app server
*/
func main() {
	engine := handlebars.New("./views", ".hbs")
	app := fiber.New(fiber.Config{ // fiber server app
		Views: engine,
	})
	setupRoutes(app)
	app.Listen(":3000")
}
