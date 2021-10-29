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
    switch strings.ToUpper(v) {
        case "GUEST":
            return GUEST_BOOKINGSTAFFROLE, nil
        case "ADMINISTRATOR":
            return ADMINISTRATOR_BOOKINGSTAFFROLE, nil
        case "VIEWER":
            return VIEWER_BOOKINGSTAFFROLE, nil
        case "EXTERNALGUEST":
            return EXTERNALGUEST_BOOKINGSTAFFROLE, nil
    }
    return 0, errors.New("Unknown BookingStaffRole value: " + v)
}
func SerializeBookingStaffRole(values []BookingStaffRole) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
