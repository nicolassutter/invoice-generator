import Alpine from "alpinejs";
import persist from "@alpinejs/persist";

Alpine.plugin(persist);

declare global {
  interface Window {
    Alpine: typeof Alpine;
  }
}

window.Alpine = Alpine;

type Item = {
  Description: string;
  Quantity: number;
  Price: number;
};

type Invoice = {
  FromName: string;
  FromEmail: string;
  FromAddress: string;
  ToName: string;
  ToEmail: string;
  ToAddress: string;
  InvoiceNumber?: string;
  InvoiceDate: string;
  DueDate: string;
  Items: Item[];
  Total: string;
};

Alpine.data("invoiceForm", () => ({
  items: [] as Item[],

  addItem() {
    this.items.push({
      Description: "",
      Quantity: 0,
      Price: 0,
    } satisfies Item);
  },

  async submitInvoice(event: Event) {
    const form = event.target as HTMLFormElement;
    const formData = new FormData(form);

    const invoice: Invoice = {
      DueDate: formData.get("dueDate") as string,
      FromAddress: formData.get("fromAddress") as string,
      FromEmail: formData.get("fromEmail") as string,
      FromName: formData.get("fromName") as string,
      InvoiceDate: formData.get("invoiceDate") as string,
      ToAddress: formData.get("toAddress") as string,
      ToEmail: formData.get("toEmail") as string,
      ToName: formData.get("toName") as string,
      Items: this.items.map((item) => ({
        Description: item.Description,
        Quantity: Number(item.Quantity),
        Price: Number(item.Price),
      })),
      Total: this.items
        .reduce((acc, item) => acc + item.Price * item.Quantity, 0)
        .toFixed(2),
    };

    try {
      const res = await fetch("/invoice", {
        method: "POST",
        headers: {
          "Content-Type": "application/json",
        },
        body: JSON.stringify(invoice),
      });
      const json: { file: string } = await res.json();
      const { file } = json;

      const fileUrl = new URL(window.location.href);
      fileUrl.pathname = file;
      window.open(fileUrl, "_blank");
    } catch (error) {
      alert("An error occurred while generating the invoice");
    }
  },

  init() {
    this.addItem();
  },
}));

Alpine.start();
