package models
import (
    "errors"
)
// Provides operations to manage the collection of accessReviewDecision entities.
type OrganizationalMessageLogoType int

const (
    // Indicates that logo is a png file
    PNG_ORGANIZATIONALMESSAGELOGOTYPE OrganizationalMessageLogoType = iota
    // UnknownFutureValue, Evolvable enumeration sentinel value. Do not use.
    UNKNOWNFUTUREVALUE_ORGANIZATIONALMESSAGELOGOTYPE
)

func (i OrganizationalMessageLogoType) String() string {
    return []string{"png", "unknownFutureValue"}[i]
}
func ParseOrganizationalMessageLogoType(v string) (interface{}, error) {
    result := PNG_ORGANIZATIONALMESSAGELOGOTYPE
    switch v {
        case "png":
            result = PNG_ORGANIZATIONALMESSAGELOGOTYPE
        case "unknownFutureValue":
            result = UNKNOWNFUTUREVALUE_ORGANIZATIONALMESSAGELOGOTYPE
        default:
            return 0, errors.New("Unknown OrganizationalMessageLogoType value: " + v)
    }
    return &result, nil
}
func SerializeOrganizationalMessageLogoType(values []OrganizationalMessageLogoType) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
