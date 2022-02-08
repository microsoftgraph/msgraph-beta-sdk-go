package graph
import (
    "strings"
    "errors"
)
// 
type BookingInvoiceStatus int

const (
    DRAFT_BOOKINGINVOICESTATUS BookingInvoiceStatus = iota
    REVIEWING_BOOKINGINVOICESTATUS
    OPEN_BOOKINGINVOICESTATUS
    CANCELED_BOOKINGINVOICESTATUS
    PAID_BOOKINGINVOICESTATUS
    CORRECTIVE_BOOKINGINVOICESTATUS
)

func (i BookingInvoiceStatus) String() string {
    return []string{"DRAFT", "REVIEWING", "OPEN", "CANCELED", "PAID", "CORRECTIVE"}[i]
}
func ParseBookingInvoiceStatus(v string) (interface{}, error) {
    result := DRAFT_BOOKINGINVOICESTATUS
    switch strings.ToUpper(v) {
        case "DRAFT":
            result = DRAFT_BOOKINGINVOICESTATUS
        case "REVIEWING":
            result = REVIEWING_BOOKINGINVOICESTATUS
        case "OPEN":
            result = OPEN_BOOKINGINVOICESTATUS
        case "CANCELED":
            result = CANCELED_BOOKINGINVOICESTATUS
        case "PAID":
            result = PAID_BOOKINGINVOICESTATUS
        case "CORRECTIVE":
            result = CORRECTIVE_BOOKINGINVOICESTATUS
        default:
            return 0, errors.New("Unknown BookingInvoiceStatus value: " + v)
    }
    return &result, nil
}
func SerializeBookingInvoiceStatus(values []BookingInvoiceStatus) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
