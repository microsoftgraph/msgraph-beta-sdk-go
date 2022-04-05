package models
import (
    "strings"
    "errors"
)
// Provides operations to call the getStaffAvailability method.
type BookingsAvailabilityStatus int

const (
    AVAILABLE_BOOKINGSAVAILABILITYSTATUS BookingsAvailabilityStatus = iota
    BUSY_BOOKINGSAVAILABILITYSTATUS
    SLOTSAVAILABLE_BOOKINGSAVAILABILITYSTATUS
    OUTOFOFFICE_BOOKINGSAVAILABILITYSTATUS
    UNKNOWNFUTUREVALUE_BOOKINGSAVAILABILITYSTATUS
)

func (i BookingsAvailabilityStatus) String() string {
    return []string{"AVAILABLE", "BUSY", "SLOTSAVAILABLE", "OUTOFOFFICE", "UNKNOWNFUTUREVALUE"}[i]
}
func ParseBookingsAvailabilityStatus(v string) (interface{}, error) {
    result := AVAILABLE_BOOKINGSAVAILABILITYSTATUS
    switch strings.ToUpper(v) {
        case "AVAILABLE":
            result = AVAILABLE_BOOKINGSAVAILABILITYSTATUS
        case "BUSY":
            result = BUSY_BOOKINGSAVAILABILITYSTATUS
        case "SLOTSAVAILABLE":
            result = SLOTSAVAILABLE_BOOKINGSAVAILABILITYSTATUS
        case "OUTOFOFFICE":
            result = OUTOFOFFICE_BOOKINGSAVAILABILITYSTATUS
        case "UNKNOWNFUTUREVALUE":
            result = UNKNOWNFUTUREVALUE_BOOKINGSAVAILABILITYSTATUS
        default:
            return 0, errors.New("Unknown BookingsAvailabilityStatus value: " + v)
    }
    return &result, nil
}
func SerializeBookingsAvailabilityStatus(values []BookingsAvailabilityStatus) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
