package graph
import (
    "strings"
    "errors"
)
// 
type BookingStaffRole int

const (
    GUEST_BOOKINGSTAFFROLE BookingStaffRole = iota
    ADMINISTRATOR_BOOKINGSTAFFROLE
    VIEWER_BOOKINGSTAFFROLE
    EXTERNALGUEST_BOOKINGSTAFFROLE
)

func (i BookingStaffRole) String() string {
    return []string{"GUEST", "ADMINISTRATOR", "VIEWER", "EXTERNALGUEST"}[i]
}
func ParseBookingStaffRole(v string) (interface{}, error) {
    result := GUEST_BOOKINGSTAFFROLE
    switch strings.ToUpper(v) {
        case "GUEST":
            result = GUEST_BOOKINGSTAFFROLE
        case "ADMINISTRATOR":
            result = ADMINISTRATOR_BOOKINGSTAFFROLE
        case "VIEWER":
            result = VIEWER_BOOKINGSTAFFROLE
        case "EXTERNALGUEST":
            result = EXTERNALGUEST_BOOKINGSTAFFROLE
        default:
            return 0, errors.New("Unknown BookingStaffRole value: " + v)
    }
    return &result, nil
}
func SerializeBookingStaffRole(values []BookingStaffRole) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
