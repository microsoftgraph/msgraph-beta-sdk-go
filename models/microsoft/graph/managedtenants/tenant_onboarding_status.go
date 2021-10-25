package managedtenants
import (
    "strings"
    "errors"
)
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
    switch strings.ToUpper(v) {
        case "INELIGIBLE":
            return INELIGIBLE_TENANTONBOARDINGSTATUS, nil
        case "INPROCESS":
            return INPROCESS_TENANTONBOARDINGSTATUS, nil
        case "ACTIVE":
            return ACTIVE_TENANTONBOARDINGSTATUS, nil
        case "INACTIVE":
            return INACTIVE_TENANTONBOARDINGSTATUS, nil
        case "UNKNOWNFUTUREVALUE":
            return UNKNOWNFUTUREVALUE_TENANTONBOARDINGSTATUS, nil
    }
    return 0, errors.New("Unknown TenantOnboardingStatus value: " + v)
}
func SerializeTenantOnboardingStatus(values []TenantOnboardingStatus) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
