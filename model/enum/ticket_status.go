package enum

type TicketStatus string

const (
	TicketStatusValid   TicketStatus = "valid"
	TicketStatusUsed    TicketStatus = "used"
	TicketStatusExpired TicketStatus = "expired"
)
