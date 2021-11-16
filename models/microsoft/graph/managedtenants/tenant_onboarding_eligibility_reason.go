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
    switch strings.ToUpper(v) {
        case "NONE":
            return NONE_TENANTONBOARDINGELIGIBILITYREASON, nil
        case "CONTRACTTYPE":
            return CONTRACTTYPE_TENANTONBOARDINGELIGIBILITYREASON, nil
        case "DELEGATEDADMINPRIVILEGES":
            return DELEGATEDADMINPRIVILEGES_TENANTONBOARDINGELIGIBILITYREASON, nil
        case "USERSCOUNT":
            return USERSCOUNT_TENANTONBOARDINGELIGIBILITYREASON, nil
        case "LICENSE":
            return LICENSE_TENANTONBOARDINGELIGIBILITYREASON, nil
        case "UNKNOWNFUTUREVALUE":
            return UNKNOWNFUTUREVALUE_TENANTONBOARDINGELIGIBILITYREASON, nil
    }
    return 0, errors.New("Unknown TenantOnboardingEligibilityReason value: " + v)
}
func SerializeTenantOnboardingEligibilityReason(values []TenantOnboardingEligibilityReason) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
