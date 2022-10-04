package models
import (
    "errors"
)
// Provides operations to manage the collection of activityStatistics entities.
type OrganizationalMessageFrequency int

const (
    // Indicates that the message will be displayed once a week
    WEEKLYONCE_ORGANIZATIONALMESSAGEFREQUENCY OrganizationalMessageFrequency = iota
    // Indicates that the message will be displayed once a month
    MONTHLYONCE_ORGANIZATIONALMESSAGEFREQUENCY
    // Indicates that the message will be displayed twice a month
    MONTHLYTWICE_ORGANIZATIONALMESSAGEFREQUENCY
    // UnknownFutureValue, Evolvable enumeration sentinel value. Do not use.
    UNKNOWNFUTUREVALUE_ORGANIZATIONALMESSAGEFREQUENCY
)

func (i OrganizationalMessageFrequency) String() string {
    return []string{"weeklyOnce", "monthlyOnce", "monthlyTwice", "unknownFutureValue"}[i]
}
func ParseOrganizationalMessageFrequency(v string) (interface{}, error) {
    result := WEEKLYONCE_ORGANIZATIONALMESSAGEFREQUENCY
    switch v {
        case "weeklyOnce":
            result = WEEKLYONCE_ORGANIZATIONALMESSAGEFREQUENCY
        case "monthlyOnce":
            result = MONTHLYONCE_ORGANIZATIONALMESSAGEFREQUENCY
        case "monthlyTwice":
            result = MONTHLYTWICE_ORGANIZATIONALMESSAGEFREQUENCY
        case "unknownFutureValue":
            result = UNKNOWNFUTUREVALUE_ORGANIZATIONALMESSAGEFREQUENCY
        default:
            return 0, errors.New("Unknown OrganizationalMessageFrequency value: " + v)
    }
    return &result, nil
}
func SerializeOrganizationalMessageFrequency(values []OrganizationalMessageFrequency) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
