// Code generated by Microsoft Kiota - DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package networkaccess
type SecurityRuleStatus int

const (
    ENABLED_SECURITYRULESTATUS SecurityRuleStatus = iota
    DISABLED_SECURITYRULESTATUS
    REPORTONLY_SECURITYRULESTATUS
    UNKNOWNFUTUREVALUE_SECURITYRULESTATUS
)

func (i SecurityRuleStatus) String() string {
    return []string{"enabled", "disabled", "reportOnly", "unknownFutureValue"}[i]
}
func ParseSecurityRuleStatus(v string) (any, error) {
    result := ENABLED_SECURITYRULESTATUS
    switch v {
        case "enabled":
            result = ENABLED_SECURITYRULESTATUS
        case "disabled":
            result = DISABLED_SECURITYRULESTATUS
        case "reportOnly":
            result = REPORTONLY_SECURITYRULESTATUS
        case "unknownFutureValue":
            result = UNKNOWNFUTUREVALUE_SECURITYRULESTATUS
        default:
            return nil, nil
    }
    return &result, nil
}
func SerializeSecurityRuleStatus(values []SecurityRuleStatus) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i SecurityRuleStatus) isMultiValue() bool {
    return false
}
