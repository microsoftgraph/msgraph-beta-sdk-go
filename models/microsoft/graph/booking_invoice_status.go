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
    switch strings.ToUpper(v) {
        case "DRAFT":
            return DRAFT_BOOKINGINVOICESTATUS, nil
        case "REVIEWING":
            return REVIEWING_BOOKINGINVOICESTATUS, nil
        case "OPEN":
            return OPEN_BOOKINGINVOICESTATUS, nil
        case "CANCELED":
            return CANCELED_BOOKINGINVOICESTATUS, nil
        case "PAID":
            return PAID_BOOKINGINVOICESTATUS, nil
        case "CORRECTIVE":
            return CORRECTIVE_BOOKINGINVOICESTATUS, nil
    }
    return 0, errors.New("Unknown BookingInvoiceStatus value: " + v)
}
func SerializeBookingInvoiceStatus(values []BookingInvoiceStatus) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
