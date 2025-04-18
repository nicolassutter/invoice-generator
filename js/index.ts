import Alpine from "alpinejs";
import persist from "@alpinejs/persist";
import { z } from "zod";

Alpine.plugin(persist);

declare global {
    interface Window {
        Alpine: typeof Alpine;
    }
}

window.Alpine = Alpine;

const ItemSchema = z.object({
    Description: z.string().min(1),
    Quantity: z.number().min(1),
    Price: z.number().optional().default(0),
});

const InvoiceSchema = z.object({
    FromName: z.string(),
    FromEmail: z.string(),
    FromAddress: z.string(),
    ToName: z.string(),
    ToEmail: z.string(),
    ToAddress: z.string(),
    InvoiceNumber: z.string().optional(),
    InvoiceDate: z.string(),
    DueDate: z.string(),
    Items: z.array(ItemSchema),
    Total: z.string(),
});

type Item = z.infer<typeof ItemSchema>;

Alpine.data("invoiceForm", () => ({
    items: [] as Item[],

    addItem() {
        this.items.push({
            Description: "",
            Quantity: 1,
            Price: 0,
        } satisfies Item);
    },

    removeItem(index: number) {
        this.items.splice(index, 1);
    },

    async submitInvoice(event: Event) {
        const form = event.target as HTMLFormElement;
        const formData = new FormData(form);

        const invoiceData = {
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

        const result = InvoiceSchema.safeParse(invoiceData);

        if (!result.success) {
            console.error(result.error.issues);
            alert("Validation error");
            return;
        }

        const invoice = result.data;

        try {
            const res = await fetch("/invoice", {
                method: "POST",
                headers: {
                    "Content-Type": "application/json",
                    Accept: "application/pdf",
                },
                body: JSON.stringify(invoice),
            });

            const blob = await res.blob();
            const url = URL.createObjectURL(blob);
            window.open(url, "_blank");
        } catch (error) {
            alert("An error occurred while generating the invoice");
        }
    },

    init() {
        this.addItem();
    },
}));

Alpine.start();
