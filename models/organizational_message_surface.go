package models
import (
    "errors"
)
// Provides operations to manage the collection of accessReview entities.
type OrganizationalMessageSurface int

const (
    // Indicates the message will be displayed on the Window's Action Center
    ACTIONCENTER_ORGANIZATIONALMESSAGESURFACE OrganizationalMessageSurface = iota
    // Indicates the message will be displayed on the Get Started page
    GETSTARTED_ORGANIZATIONALMESSAGESURFACE
    // Indicates the message will be displayed to the Soft Landing which is anchored to the Windows taskbar
    SOFTLANDING_ORGANIZATIONALMESSAGESURFACE
    // UnknownFutureValue, Evolvable enumeration sentinel value. Do not use.
    UNKNOWNFUTUREVALUE_ORGANIZATIONALMESSAGESURFACE
)

func (i OrganizationalMessageSurface) String() string {
    return []string{"actionCenter", "getStarted", "softLanding", "unknownFutureValue"}[i]
}
func ParseOrganizationalMessageSurface(v string) (interface{}, error) {
    result := ACTIONCENTER_ORGANIZATIONALMESSAGESURFACE
    switch v {
        case "actionCenter":
            result = ACTIONCENTER_ORGANIZATIONALMESSAGESURFACE
        case "getStarted":
            result = GETSTARTED_ORGANIZATIONALMESSAGESURFACE
        case "softLanding":
            result = SOFTLANDING_ORGANIZATIONALMESSAGESURFACE
        case "unknownFutureValue":
            result = UNKNOWNFUTUREVALUE_ORGANIZATIONALMESSAGESURFACE
        default:
            return 0, errors.New("Unknown OrganizationalMessageSurface value: " + v)
    }
    return &result, nil
}
func SerializeOrganizationalMessageSurface(values []OrganizationalMessageSurface) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
