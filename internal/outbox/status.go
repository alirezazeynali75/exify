//go:generate go-enum --marshal --nocase --noprefix --lower
package outbox

/*
ENUM(

	READY = "READY"
	SENDING = "SENDING"
	SENT = "SENT"

)
*/
type OutboxStatus string
