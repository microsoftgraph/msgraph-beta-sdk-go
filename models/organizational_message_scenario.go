package models
import (
    "errors"
)
// Provides operations to manage the collection of activityStatistics entities.
type OrganizationalMessageScenario int

const (
    // Indicates onboarding scenario
    ONBOARDING_ORGANIZATIONALMESSAGESCENARIO OrganizationalMessageScenario = iota
    // Indicates lifecycle scenario
    LIFECYCLE_ORGANIZATIONALMESSAGESCENARIO
)

func (i OrganizationalMessageScenario) String() string {
    return []string{"onboarding", "lifecycle"}[i]
}
func ParseOrganizationalMessageScenario(v string) (interface{}, error) {
    result := ONBOARDING_ORGANIZATIONALMESSAGESCENARIO
    switch v {
        case "onboarding":
            result = ONBOARDING_ORGANIZATIONALMESSAGESCENARIO
        case "lifecycle":
            result = LIFECYCLE_ORGANIZATIONALMESSAGESCENARIO
        default:
            return 0, errors.New("Unknown OrganizationalMessageScenario value: " + v)
    }
    return &result, nil
}
func SerializeOrganizationalMessageScenario(values []OrganizationalMessageScenario) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
