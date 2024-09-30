//go:generate go-enum --marshal --nocase --noprefix --lower
package payment

/*
ENUM(

	PROCESSING = "PROCESSING"
	COMPLETED = "COMPLETED"
	FAILED = "FAILED"

)
*/
type PaymentStatus string
