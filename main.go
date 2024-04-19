// ⚡️ Fiber is an Express inspired web framework written in Go with ☕️
// 🤖 Github Repository: https://github.com/gofiber/fiber
// 📌 API Documentation: https://docs.gofiber.io

package main

import (
	"log"
	"net/http"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html/v2"
)

func main() {
	engine := html.New("./templates", ".html")

	// Fiber instance
	app := fiber.New(fiber.Config{
		Views: engine,
	})

	// Routes
	app.Get("/todos", listTODOs)
	app.Post("/add-todo", addTODOs)
	app.Delete("/delete-todo/:id", deleteTODOs)
	app.Static("/", "./frontend/build")

	// Start server
	log.Fatal(app.Listen(":3000"))
}

// Handler
func listTODOs(c *fiber.Ctx) error {
	log.Println("listTODOs", "Total TODOs: ", len(TODOs))
	return c.Render("todo", fiber.Map{
		"list": TODOs,
	})
	// Example using plain text
	// var resultBuffer strings.Builder
	// resultBuffer.WriteString("<table>\n<tr><th>ID</th><th>Title</th></tr>\n")
	// for _, todo := range TODOs {
	// 	resultBuffer.WriteString(fmt.Sprintf("<tr><td>%d</td><td>%s</td></tr>\n", todo.ID, todo.Title))
	// }
	// resultBuffer.WriteString("</table>")
	// return c.SendString(resultBuffer.String())
}

func addTODOs(c *fiber.Ctx) error {
	CreateTODO(c.FormValue("title", "empty"))
	log.Printf("Added TODO: %+v\n", TODOs[len(TODOs)-1])
	return c.Redirect("/todos", http.StatusFound)
}

func deleteTODOs(c *fiber.Ctx) error {
	log.Printf("Deleted TODO puro: %+v\n", c.Params("id"))
	id, err := strconv.ParseUint(c.Params("id"), 10, 32)
	if err != nil {
		return c.Status(http.StatusBadRequest).SendString("invalid id")
	}
	log.Printf("Deleted TODO uint: %+v\n", id)
	log.Println("Total TODOs before delete: ", len(TODOs))
	DeleteTODO(uint32(id))
	log.Println("Total TODOs after delete: ", len(TODOs))
	return listTODOs(c)
}
