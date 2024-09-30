//go:generate go-enum --marshal --nocase --noprefix --lower
package payment


/*
ENUM(

	DEPOSIT = "deposit"
	WITHDRAWAL = "withdrawal"

)
*/
type PaymentType string