package managedtenants
import (
    "strings"
    "errors"
)
// 
type TenantOnboardingEligibilityReason int

const (
    NONE_TENANTONBOARDINGELIGIBILITYREASON TenantOnboardingEligibilityReason = iota
    CONTRACTTYPE_TENANTONBOARDINGELIGIBILITYREASON
    DELEGATEDADMINPRIVILEGES_TENANTONBOARDINGELIGIBILITYREASON
    USERSCOUNT_TENANTONBOARDINGELIGIBILITYREASON
    LICENSE_TENANTONBOARDINGELIGIBILITYREASON
    UNKNOWNFUTUREVALUE_TENANTONBOARDINGELIGIBILITYREASON
)

func (i TenantOnboardingEligibilityReason) String() string {
    return []string{"NONE", "CONTRACTTYPE", "DELEGATEDADMINPRIVILEGES", "USERSCOUNT", "LICENSE", "UNKNOWNFUTUREVALUE"}[i]
}
func ParseTenantOnboardingEligibilityReason(v string) (interface{}, error) {
    result := NONE_TENANTONBOARDINGELIGIBILITYREASON
    switch strings.ToUpper(v) {
        case "NONE":
            result = NONE_TENANTONBOARDINGELIGIBILITYREASON
        case "CONTRACTTYPE":
            result = CONTRACTTYPE_TENANTONBOARDINGELIGIBILITYREASON
        case "DELEGATEDADMINPRIVILEGES":
            result = DELEGATEDADMINPRIVILEGES_TENANTONBOARDINGELIGIBILITYREASON
        case "USERSCOUNT":
            result = USERSCOUNT_TENANTONBOARDINGELIGIBILITYREASON
        case "LICENSE":
            result = LICENSE_TENANTONBOARDINGELIGIBILITYREASON
        case "UNKNOWNFUTUREVALUE":
            result = UNKNOWNFUTUREVALUE_TENANTONBOARDINGELIGIBILITYREASON
        default:
            return 0, errors.New("Unknown TenantOnboardingEligibilityReason value: " + v)
    }
    return &result, nil
}
func SerializeTenantOnboardingEligibilityReason(values []TenantOnboardingEligibilityReason) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
