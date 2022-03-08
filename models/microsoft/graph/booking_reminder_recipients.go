package graph
import (
    "strings"
    "errors"
)
// Provides operations to manage the collection of bookingBusiness entities.
type BookingReminderRecipients int

const (
    ALLATTENDEES_BOOKINGREMINDERRECIPIENTS BookingReminderRecipients = iota
    STAFF_BOOKINGREMINDERRECIPIENTS
    CUSTOMER_BOOKINGREMINDERRECIPIENTS
)

func (i BookingReminderRecipients) String() string {
    return []string{"ALLATTENDEES", "STAFF", "CUSTOMER"}[i]
}
func ParseBookingReminderRecipients(v string) (interface{}, error) {
    result := ALLATTENDEES_BOOKINGREMINDERRECIPIENTS
    switch strings.ToUpper(v) {
        case "ALLATTENDEES":
            result = ALLATTENDEES_BOOKINGREMINDERRECIPIENTS
        case "STAFF":
            result = STAFF_BOOKINGREMINDERRECIPIENTS
        case "CUSTOMER":
            result = CUSTOMER_BOOKINGREMINDERRECIPIENTS
        default:
            return 0, errors.New("Unknown BookingReminderRecipients value: " + v)
    }
    return &result, nil
}
func SerializeBookingReminderRecipients(values []BookingReminderRecipients) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
