package main

import (
	"context"
	"fmt"
	"github.com/a-h/templ"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/adaptor"
	"github.com/google/uuid"
	"os"
	"resty.dev/v3"
)

func main() {
	app := fiber.New()

	app.Static("/assets", "./assets")
	app.Static("/out", "./out")

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

		outDir := fmt.Sprintf("out/invoices-%s", uid.String())

		os.MkdirAll(outDir, os.ModePerm)
		htmlFile, err := os.Create(fmt.Sprintf("%s/index.html", outDir))

		if err != nil {
			return fiber.NewError(fiber.StatusInternalServerError, "Failed to create output file")
		}
		defer htmlFile.Close()

		err = invoice(&invoiceData).Render(context.Background(), htmlFile)

		if err != nil {
			return fiber.NewError(fiber.StatusInternalServerError, "Failed to render invoice")
		}

		// make the call to Gotenberg API to convert the HTML file to PDF
		client := resty.New()
		defer client.Close()

		gotenbergHost := "http://localhost:9000"

		resp, err := client.R().
			SetFile("files", htmlFile.Name()).
			Post(fmt.Sprintf("%s/forms/chromium/convert/html", gotenbergHost))

		if err != nil {
			return fiber.NewError(fiber.StatusInternalServerError, "Failed to convert HTML to PDF")
		}

		pdfBytes := resp.Bytes()
		c.Response().Header.Set("Content-Type", "application/pdf")
		return c.Send(pdfBytes)
	})

	app.Listen(":3000")
}
