package models
import (
    "errors"
)
// Provides operations to manage the collection of accessReviewDecision entities.
type OrganizationalMessageLogoType int

const (
    // Indicates that logo is a png file
    PNG_ORGANIZATIONALMESSAGELOGOTYPE OrganizationalMessageLogoType = iota
)

func (i OrganizationalMessageLogoType) String() string {
    return []string{"png"}[i]
}
func ParseOrganizationalMessageLogoType(v string) (interface{}, error) {
    result := PNG_ORGANIZATIONALMESSAGELOGOTYPE
    switch v {
        case "png":
            result = PNG_ORGANIZATIONALMESSAGELOGOTYPE
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
