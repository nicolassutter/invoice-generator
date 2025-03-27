// Code generated by templ - DO NOT EDIT.

// templ: version: v0.3.857
package main

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import templruntime "github.com/a-h/templ/runtime"

import "fmt"

func page() templ.Component {
	return templruntime.GeneratedTemplate(func(templ_7745c5c3_Input templruntime.GeneratedComponentInput) (templ_7745c5c3_Err error) {
		templ_7745c5c3_W, ctx := templ_7745c5c3_Input.Writer, templ_7745c5c3_Input.Context
		if templ_7745c5c3_CtxErr := ctx.Err(); templ_7745c5c3_CtxErr != nil {
			return templ_7745c5c3_CtxErr
		}
		templ_7745c5c3_Buffer, templ_7745c5c3_IsBuffer := templruntime.GetBuffer(templ_7745c5c3_W)
		if !templ_7745c5c3_IsBuffer {
			defer func() {
				templ_7745c5c3_BufErr := templruntime.ReleaseBuffer(templ_7745c5c3_Buffer)
				if templ_7745c5c3_Err == nil {
					templ_7745c5c3_Err = templ_7745c5c3_BufErr
				}
			}()
		}
		ctx = templ.InitializeContext(ctx)
		templ_7745c5c3_Var1 := templ.GetChildren(ctx)
		if templ_7745c5c3_Var1 == nil {
			templ_7745c5c3_Var1 = templ.NopComponent
		}
		ctx = templ.ClearChildren(ctx)
		templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 1, "<!doctype html><html lang=\"en\" x-data=\"{ darkMode: $persist(false), lightTheme: &#39;winter&#39; }\" data-theme=\"winter\" :data-theme=\"darkMode ? &#39;dark&#39; : lightTheme\" class=\"bg-base-200\"><head><meta charset=\"UTF-8\"><meta name=\"viewport\" content=\"width=device-width, initial-scale=1.0\"><title>Invoice Maker</title><link rel=\"stylesheet\" href=\"/assets/styles.css\"><script type=\"module\" src=\"/assets/index.js\"></script></head><body class=\"min-h-screen\"><div class=\"container mx-auto p-8\"><!-- Dark mode toggle --><div class=\"flex justify-end mb-6\"><label class=\"swap swap-rotate\"><input type=\"checkbox\" @click=\"darkMode = !darkMode\"> <svg class=\"swap-on fill-current w-8 h-8\" xmlns=\"http://www.w3.org/2000/svg\" viewBox=\"0 0 24 24\"><path d=\"M5.64,17l-.71.71a1,1,0,0,0,0,1.41,1,1,0,0,0,1.41,0l.71-.71A1,1,0,0,0,5.64,17ZM5,12a1,1,0,0,0-1-1H3a1,1,0,0,0,0,2H4A1,1,0,0,0,5,12Zm7-7a1,1,0,0,0,1-1V3a1,1,0,0,0-2,0V4A1,1,0,0,0,12,5ZM5.64,7.05a1,1,0,0,0,.7.29,1,1,0,0,0,.71-.29,1,1,0,0,0,0-1.41l-.71-.71A1,1,0,0,0,4.93,6.34Zm12,.29a1,1,0,0,0,.7-.29l.71-.71a1,1,0,1,0-1.41-1.41L17,5.64a1,1,0,0,0,0,1.41A1,1,0,0,0,17.66,7.34ZM21,11H20a1,1,0,0,0,0,2h1a1,1,0,0,0,0-2Zm-9,8a1,1,0,0,0-1,1v1a1,1,0,0,0,2,0V20A1,1,0,0,0,12,19ZM18.36,17A1,1,0,0,0,17,18.36l.71.71a1,1,0,0,0,1.41,0,1,1,0,0,0,0-1.41ZM12,6.5A5.5,5.5,0,1,0,17.5,12,5.51,5.51,0,0,0,12,6.5Zm0,9A3.5,3.5,0,1,1,15.5,12,3.5,3.5,0,0,1,12,15.5Z\"></path></svg> <svg class=\"swap-off fill-current w-8 h-8\" xmlns=\"http://www.w3.org/2000/svg\" viewBox=\"0 0 24 24\"><path d=\"M21.64,13a1,1,0,0,0-1.05-.14,8.05,8.05,0,0,1-3.37.73A8.15,8.15,0,0,1,9.08,5.49a8.59,8.59,0,0,1,.25-2A1,1,0,0,0,8,2.36,10.14,10.14,0,1,0,22,14.05,1,1,0,0,0,21.64,13Zm-9.5,6.69A8.14,8.14,0,0,1,7.08,5.22v.27A10.15,10.15,0,0,0,17.22,15.63a9.79,9.79,0,0,0,2.1-.22A8.11,8.11,0,0,1,12.14,19.73Z\"></path></svg></label></div><!-- Invoice Form --><div class=\"card bg-base-100 shadow-xl\"><div class=\"card-body\"><h1 class=\"card-title text-3xl mb-8\">Invoice Maker</h1><form x-data=\"invoiceForm\" @submit.prevent=\"submitInvoice\" id=\"invoiceForm\" class=\"space-y-6\"><!-- Bill From --><div class=\"space-y-4\"><h2 class=\"text-xl font-semibold\">Bill From</h2><div class=\"grid grid-cols-1 md:grid-cols-2 gap-4\"><input type=\"text\" name=\"fromName\" placeholder=\"Your Name/Company\" class=\"input input-bordered w-full\"> <input type=\"email\" name=\"fromEmail\" placeholder=\"Your Email\" class=\"input input-bordered w-full\"> <textarea name=\"fromAddress\" placeholder=\"Your Address\" class=\"textarea textarea-bordered w-full md:col-span-2\" rows=\"3\"></textarea></div></div><!-- Bill To --><div class=\"space-y-4\"><h2 class=\"text-xl font-semibold\">Bill To</h2><div class=\"grid grid-cols-1 md:grid-cols-2 gap-4\"><input type=\"text\" name=\"toName\" placeholder=\"Client Name/Company\" class=\"input input-bordered w-full\"> <input type=\"email\" name=\"toEmail\" placeholder=\"Client Email\" class=\"input input-bordered w-full\"> <textarea name=\"toAddress\" placeholder=\"Client Address\" class=\"textarea textarea-bordered w-full md:col-span-2\" rows=\"3\"></textarea></div></div><!-- Invoice Details --><div class=\"space-y-4\"><h2 class=\"text-xl font-semibold\">Invoice Details</h2><div class=\"grid grid-cols-1 md:grid-cols-3 gap-4\"><input type=\"text\" name=\"invoiceNumber\" placeholder=\"Invoice Number\" class=\"input input-bordered w-full\"> <input type=\"date\" name=\"invoiceDate\" class=\"input input-bordered w-full\"> <input type=\"date\" name=\"dueDate\" class=\"input input-bordered w-full\"></div></div><!-- Items --><div class=\"space-y-4\"><h2 class=\"text-xl font-semibold\">Items</h2><div class=\"space-y-4\"><template x-for=\"(item, index) in items\" :key=\"index\"><div class=\"grid grid-cols-12 gap-4\"><input type=\"text\" x-model=\"item.Description\" placeholder=\"Description\" class=\"input input-bordered col-span-5\"> <input type=\"number\" x-model=\"item.Quantity\" placeholder=\"Qty\" class=\"input input-bordered col-span-2\"> <input type=\"number\" x-model=\"item.Price\" placeholder=\"Price\" class=\"input input-bordered col-span-2\"><div class=\"col-span-2 flex items-center\"><span x-text=\"(item.Quantity * item.Price || 0).toFixed(2)\" class=\"text-lg font-semibold\"></span></div><button type=\"button\" @click=\"() =&gt; removeItem(index)\" class=\"btn btn-circle btn-error btn-sm col-span-1\" aria-label=\"Remove Item\"><span aria-hidden=\"true\">x</span></button></div></template><button type=\"button\" @click=\"addItem\" class=\"btn btn-secondary w-full\">Add Item</button></div></div><!-- Submit Button --><button type=\"submit\" class=\"btn btn-primary w-full\">Generate Invoice</button></form></div></div></div></body></html>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		return nil
	})
}

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

func invoice(invoice *Invoice) templ.Component {
	return templruntime.GeneratedTemplate(func(templ_7745c5c3_Input templruntime.GeneratedComponentInput) (templ_7745c5c3_Err error) {
		templ_7745c5c3_W, ctx := templ_7745c5c3_Input.Writer, templ_7745c5c3_Input.Context
		if templ_7745c5c3_CtxErr := ctx.Err(); templ_7745c5c3_CtxErr != nil {
			return templ_7745c5c3_CtxErr
		}
		templ_7745c5c3_Buffer, templ_7745c5c3_IsBuffer := templruntime.GetBuffer(templ_7745c5c3_W)
		if !templ_7745c5c3_IsBuffer {
			defer func() {
				templ_7745c5c3_BufErr := templruntime.ReleaseBuffer(templ_7745c5c3_Buffer)
				if templ_7745c5c3_Err == nil {
					templ_7745c5c3_Err = templ_7745c5c3_BufErr
				}
			}()
		}
		ctx = templ.InitializeContext(ctx)
		templ_7745c5c3_Var2 := templ.GetChildren(ctx)
		if templ_7745c5c3_Var2 == nil {
			templ_7745c5c3_Var2 = templ.NopComponent
		}
		ctx = templ.ClearChildren(ctx)
		templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 2, "<!doctype html><html lang=\"en\"><head><meta charset=\"UTF-8\"><meta name=\"viewport\" content=\"width=device-width, initial-scale=1.0\"><title>Invoice</title><script defer src=\"https://unpkg.com/alpinejs@3.x.x/dist/cdn.min.js\"></script><style>\n    body {\n      font-family: Arial, sans-serif;\n      line-height: 1.6;\n      margin: 0;\n      padding: 40px;\n    }\n\n    .invoice-header {\n      display: flex;\n      justify-content: space-between;\n      margin-bottom: 40px;\n    }\n\n    .invoice-title {\n      font-size: 2em;\n      color: #2563eb;\n      margin: 0;\n    }\n\n    .invoice-details {\n      text-align: right;\n    }\n\n    .bill-section {\n      display: flex;\n      justify-content: space-between;\n      margin-bottom: 40px;\n    }\n\n    .bill-box {\n      width: 45%;\n    }\n\n    .bill-box h2 {\n      color: #4b5563;\n      border-bottom: 2px solid #e5e7eb;\n      padding-bottom: 10px;\n    }\n\n    .items-table {\n      width: 100%;\n      border-collapse: collapse;\n      margin-bottom: 40px;\n    }\n\n    .items-table th {\n      background-color: #f3f4f6;\n      padding: 12px;\n      text-align: left;\n      border-bottom: 2px solid #e5e7eb;\n    }\n\n    .items-table td {\n      padding: 12px;\n      border-bottom: 1px solid #e5e7eb;\n    }\n\n    .items-table tr:last-child td {\n      border-bottom: none;\n    }\n\n    .total-section {\n      text-align: right;\n      margin-top: 20px;\n    }\n\n    .total-row {\n      display: flex;\n      justify-content: flex-end;\n      margin-bottom: 10px;\n    }\n\n    .total-label {\n      width: 150px;\n      text-align: right;\n      margin-right: 20px;\n    }\n\n    .total-amount {\n      width: 150px;\n      text-align: right;\n    }\n\n    .grand-total {\n      font-size: 1.2em;\n      font-weight: bold;\n      color: #2563eb;\n      border-top: 2px solid #e5e7eb;\n      padding-top: 10px;\n    }\n  </style></head><body><div class=\"invoice-header\"><div><h1 class=\"invoice-title\">INVOICE</h1><div id=\"from-details\"><p><strong>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		var templ_7745c5c3_Var3 string
		templ_7745c5c3_Var3, templ_7745c5c3_Err = templ.JoinStringErrs(invoice.FromName)
		if templ_7745c5c3_Err != nil {
			return templ.Error{Err: templ_7745c5c3_Err, FileName: `page.templ`, Line: 312, Col: 35}
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var3))
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 3, "</strong></p><p>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		var templ_7745c5c3_Var4 string
		templ_7745c5c3_Var4, templ_7745c5c3_Err = templ.JoinStringErrs(invoice.FromEmail)
		if templ_7745c5c3_Err != nil {
			return templ.Error{Err: templ_7745c5c3_Err, FileName: `page.templ`, Line: 313, Col: 28}
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var4))
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 4, "</p><p>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		var templ_7745c5c3_Var5 string
		templ_7745c5c3_Var5, templ_7745c5c3_Err = templ.JoinStringErrs(invoice.FromAddress)
		if templ_7745c5c3_Err != nil {
			return templ.Error{Err: templ_7745c5c3_Err, FileName: `page.templ`, Line: 314, Col: 30}
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var5))
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 5, "</p></div></div><div class=\"invoice-details\"><p><strong>Invoice #:</strong> <span>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		var templ_7745c5c3_Var6 string
		templ_7745c5c3_Var6, templ_7745c5c3_Err = templ.JoinStringErrs(invoice.InvoiceNumber)
		if templ_7745c5c3_Err != nil {
			return templ.Error{Err: templ_7745c5c3_Err, FileName: `page.templ`, Line: 318, Col: 65}
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var6))
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 6, "</span></p><p><strong>Date:</strong> <span>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		var templ_7745c5c3_Var7 string
		templ_7745c5c3_Var7, templ_7745c5c3_Err = templ.JoinStringErrs(invoice.InvoiceDate)
		if templ_7745c5c3_Err != nil {
			return templ.Error{Err: templ_7745c5c3_Err, FileName: `page.templ`, Line: 319, Col: 58}
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var7))
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 7, "</span></p><p><strong>Due Date:</strong> <span>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		var templ_7745c5c3_Var8 string
		templ_7745c5c3_Var8, templ_7745c5c3_Err = templ.JoinStringErrs(invoice.DueDate)
		if templ_7745c5c3_Err != nil {
			return templ.Error{Err: templ_7745c5c3_Err, FileName: `page.templ`, Line: 320, Col: 58}
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var8))
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 8, "</span></p></div></div><div class=\"bill-section\"><div class=\"bill-box\"><h2>Bill To</h2><div id=\"to-details\"><p><strong>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		var templ_7745c5c3_Var9 string
		templ_7745c5c3_Var9, templ_7745c5c3_Err = templ.JoinStringErrs(invoice.ToName)
		if templ_7745c5c3_Err != nil {
			return templ.Error{Err: templ_7745c5c3_Err, FileName: `page.templ`, Line: 327, Col: 33}
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var9))
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 9, "</strong></p><p>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		var templ_7745c5c3_Var10 string
		templ_7745c5c3_Var10, templ_7745c5c3_Err = templ.JoinStringErrs(invoice.ToEmail)
		if templ_7745c5c3_Err != nil {
			return templ.Error{Err: templ_7745c5c3_Err, FileName: `page.templ`, Line: 328, Col: 26}
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var10))
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 10, "</p><p>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		var templ_7745c5c3_Var11 string
		templ_7745c5c3_Var11, templ_7745c5c3_Err = templ.JoinStringErrs(invoice.ToAddress)
		if templ_7745c5c3_Err != nil {
			return templ.Error{Err: templ_7745c5c3_Err, FileName: `page.templ`, Line: 329, Col: 28}
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var11))
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 11, "</p></div></div></div><table class=\"items-table\"><thead><tr><th>Description</th><th>Quantity</th><th>Price</th><th>Amount</th></tr></thead> <tbody>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		for _, item := range invoice.Items {
			templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 12, "<tr><td>")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			var templ_7745c5c3_Var12 string
			templ_7745c5c3_Var12, templ_7745c5c3_Err = templ.JoinStringErrs(item.Description)
			if templ_7745c5c3_Err != nil {
				return templ.Error{Err: templ_7745c5c3_Err, FileName: `page.templ`, Line: 345, Col: 29}
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var12))
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 13, "</td><td>")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			var templ_7745c5c3_Var13 string
			templ_7745c5c3_Var13, templ_7745c5c3_Err = templ.JoinStringErrs(fmt.Sprintf("%d", item.Quantity))
			if templ_7745c5c3_Err != nil {
				return templ.Error{Err: templ_7745c5c3_Err, FileName: `page.templ`, Line: 346, Col: 45}
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var13))
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 14, "</td><td>")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			var templ_7745c5c3_Var14 string
			templ_7745c5c3_Var14, templ_7745c5c3_Err = templ.JoinStringErrs(fmt.Sprintf("%f", item.Price))
			if templ_7745c5c3_Err != nil {
				return templ.Error{Err: templ_7745c5c3_Err, FileName: `page.templ`, Line: 347, Col: 42}
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var14))
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 15, "</td><td>")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			var templ_7745c5c3_Var15 string
			templ_7745c5c3_Var15, templ_7745c5c3_Err = templ.JoinStringErrs(fmt.Sprintf("%f", float64(item.Quantity)*item.Price))
			if templ_7745c5c3_Err != nil {
				return templ.Error{Err: templ_7745c5c3_Err, FileName: `page.templ`, Line: 348, Col: 67}
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var15))
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 16, "</td></tr>")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
		}
		templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 17, "</tbody></table><div class=\"total-section\"><div class=\"total-row\"><div class=\"total-label\">Tax (0%):</div><div class=\"total-amount\">$0.00</div></div><div class=\"total-row grand-total\"><div class=\"total-label\">Total:</div><div class=\"total-amount\">")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		var templ_7745c5c3_Var16 string
		templ_7745c5c3_Var16, templ_7745c5c3_Err = templ.JoinStringErrs(invoice.Total)
		if templ_7745c5c3_Err != nil {
			return templ.Error{Err: templ_7745c5c3_Err, FileName: `page.templ`, Line: 360, Col: 46}
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var16))
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 18, "</div></div></div></body></html>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		return nil
	})
}

var _ = templruntime.GeneratedTemplate
