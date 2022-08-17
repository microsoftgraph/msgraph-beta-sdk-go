package models
import (
    "errors"
)
// Provides operations to manage the collection of activityStatistics entities.
type OrganizationalMessageStatus int

const (
    // Indicates that the message has been scheduled for a future date
    SCHEDULED_ORGANIZATIONALMESSAGESTATUS OrganizationalMessageStatus = iota
    // Indicates that the message is currently live and being presented to clients
    ACTIVE_ORGANIZATIONALMESSAGESTATUS
    // Indicates that the message has been displayed to users already and is no longer active
    COMPLETED_ORGANIZATIONALMESSAGESTATUS
    // Indicates that the message has been cancelled and will not be shown
    CANCELLED_ORGANIZATIONALMESSAGESTATUS
)

func (i OrganizationalMessageStatus) String() string {
    return []string{"scheduled", "active", "completed", "cancelled"}[i]
}
func ParseOrganizationalMessageStatus(v string) (interface{}, error) {
    result := SCHEDULED_ORGANIZATIONALMESSAGESTATUS
    switch v {
        case "scheduled":
            result = SCHEDULED_ORGANIZATIONALMESSAGESTATUS
        case "active":
            result = ACTIVE_ORGANIZATIONALMESSAGESTATUS
        case "completed":
            result = COMPLETED_ORGANIZATIONALMESSAGESTATUS
        case "cancelled":
            result = CANCELLED_ORGANIZATIONALMESSAGESTATUS
        default:
            return 0, errors.New("Unknown OrganizationalMessageStatus value: " + v)
    }
    return &result, nil
}
func SerializeOrganizationalMessageStatus(values []OrganizationalMessageStatus) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
