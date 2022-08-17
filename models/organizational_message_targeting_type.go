package models
import (
    "errors"
)
// Provides operations to manage the collection of accessReviewDecision entities.
type OrganizationalMessageTargetingType int

const (
    // Indicates that client devices are targeted by their AAD group
    AADGROUP_ORGANIZATIONALMESSAGETARGETINGTYPE OrganizationalMessageTargetingType = iota
)

func (i OrganizationalMessageTargetingType) String() string {
    return []string{"aadGroup"}[i]
}
func ParseOrganizationalMessageTargetingType(v string) (interface{}, error) {
    result := AADGROUP_ORGANIZATIONALMESSAGETARGETINGTYPE
    switch v {
        case "aadGroup":
            result = AADGROUP_ORGANIZATIONALMESSAGETARGETINGTYPE
        default:
            return 0, errors.New("Unknown OrganizationalMessageTargetingType value: " + v)
    }
    return &result, nil
}
func SerializeOrganizationalMessageTargetingType(values []OrganizationalMessageTargetingType) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
