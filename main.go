package main

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/a-h/templ"
	"github.com/google/uuid"
	"log"
	"net/http"
	"os"
	"resty.dev/v3"
)

func main() {
	http.Handle("/", templ.Handler(page()))

	// Handle the "/invoice" endpoint to process invoice data.
	// It decodes the JSON request body into an Invoice struct,
	// generates a unique identifier for the invoice, creates an HTML file,
	// and renders the invoice content into the file.
	http.HandleFunc("/invoice", func(w http.ResponseWriter, r *http.Request) {
		var invoiceData Invoice
		err := json.NewDecoder(r.Body).Decode(&invoiceData)

		if err != nil {
			http.Error(w, "Invalid invoice data", http.StatusBadRequest)
			return
		}

		uid := uuid.New()
		invoiceData.InvoiceNumber = uid.String()

		outDir := fmt.Sprintf("out/invoices-%s", uid.String())

		os.MkdirAll(outDir, os.ModePerm)
		htmlFile, err := os.Create(fmt.Sprintf("%s/index.html", outDir))

		if err != nil {
			log.Fatalf("failed to create output file: %v", err)
		}
		defer htmlFile.Close()

		err = invoice(&invoiceData).Render(context.Background(), htmlFile)

		if err != nil {
			log.Fatalf("failed to render invoice: %v", err)
		}

		// make the call to Gotenberg API to convert the HTML file to PDF
		client := resty.New()
		defer client.Close()

		pdfName := fmt.Sprintf("%s/invoice.pdf", outDir)
		gotenbergHost := "http://localhost:9000"

		_, err = client.R().
			SetFile("files", htmlFile.Name()).
			SetSaveResponse(true).
			SetOutputFileName(pdfName).
			Post(fmt.Sprintf("%s/forms/chromium/convert/html", gotenbergHost))

		if err != nil {
			log.Fatalf("failed to convert HTML to PDF: %v", err)
		}

		// return http file location
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]string{"file": pdfName})
	})

	assets := http.FileServer(http.Dir("./assets"))
	http.Handle("/assets/", http.StripPrefix("/assets/", assets))

	out := http.FileServer(http.Dir("./out"))
	http.Handle("/out/", http.StripPrefix("/out/", out))

	fmt.Println("Listening on :3000")
	http.ListenAndServe(":3000", nil)
}
