package main

import (
	"bytes"
	"context"
	"fmt"
	"github.com/a-h/templ"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/adaptor"
	"github.com/google/uuid"
)

type Item struct {
	Description string
	Quantity    int
	Price       float64
}

type Invoice struct {
	FromName      string
	FromEmail     string
	FromAddress   string
	ToName        string
	ToEmail       string
	ToAddress     string
	InvoiceNumber string
	InvoiceDate   string
	DueDate       string
	Items         []Item
	Total         string
}

func main() {
	app := fiber.New()

	app.Static("/assets", "./assets")

	app.Get("/", adaptor.HTTPHandler(templ.Handler(page())))

	// Handle the "/invoice" endpoint to process invoice data.
	// It decodes the JSON request body into an Invoice struct,
	// generates a unique identifier for the invoice, creates an HTML file,
	// and renders the invoice content into the file.
	app.Post("/invoice", func(c *fiber.Ctx) error {
		var invoiceData Invoice

		if err := c.BodyParser(&invoiceData); err != nil {
			return err
		}

		uid := uuid.New()
		invoiceData.InvoiceNumber = uid.String()

		var builder bytes.Buffer
		err := invoice(&invoiceData).Render(context.Background(), &builder)
		html := builder.Bytes()

		if err != nil {
			return fiber.NewError(fiber.StatusInternalServerError, "Failed to render invoice")
		}

		// make the call to Gotenberg API to convert the HTML file to PDF
		gotenbergHost := "http://localhost:9000"
		gotenbergUrl := fmt.Sprintf("%s/forms/chromium/convert/html", gotenbergHost)

		agent := fiber.Post(gotenbergUrl)

		agent.FileData(&fiber.FormFile{
			Fieldname: "files",
			Name:      "index.html",
			Content:   html,
		}).MultipartForm(nil)

		_, resp, responseErrors := agent.Bytes()

		if len(responseErrors) > 0 {
			return fiber.NewError(fiber.StatusInternalServerError, "Failed to convert HTML to PDF")
		}

		c.Response().Header.Set("Content-Type", "application/pdf")
		return c.Send(resp)
	})

	app.Listen(":3000")
}
