package managedtenants
import (
    "strings"
    "errors"
)
// 
type TenantOnboardingStatus int

const (
    INELIGIBLE_TENANTONBOARDINGSTATUS TenantOnboardingStatus = iota
    INPROCESS_TENANTONBOARDINGSTATUS
    ACTIVE_TENANTONBOARDINGSTATUS
    INACTIVE_TENANTONBOARDINGSTATUS
    UNKNOWNFUTUREVALUE_TENANTONBOARDINGSTATUS
)

func (i TenantOnboardingStatus) String() string {
    return []string{"INELIGIBLE", "INPROCESS", "ACTIVE", "INACTIVE", "UNKNOWNFUTUREVALUE"}[i]
}
func ParseTenantOnboardingStatus(v string) (interface{}, error) {
    result := INELIGIBLE_TENANTONBOARDINGSTATUS
    switch strings.ToUpper(v) {
        case "INELIGIBLE":
            result = INELIGIBLE_TENANTONBOARDINGSTATUS
        case "INPROCESS":
            result = INPROCESS_TENANTONBOARDINGSTATUS
        case "ACTIVE":
            result = ACTIVE_TENANTONBOARDINGSTATUS
        case "INACTIVE":
            result = INACTIVE_TENANTONBOARDINGSTATUS
        case "UNKNOWNFUTUREVALUE":
            result = UNKNOWNFUTUREVALUE_TENANTONBOARDINGSTATUS
        default:
            return 0, errors.New("Unknown TenantOnboardingStatus value: " + v)
    }
    return &result, nil
}
func SerializeTenantOnboardingStatus(values []TenantOnboardingStatus) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
