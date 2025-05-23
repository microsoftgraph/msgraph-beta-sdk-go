// Code generated by Microsoft Kiota - DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package models
// Possible values of Defender PUA Protection
type DefenderProtectionType int

const (
    // Device default value, no intent.
    USERDEFINED_DEFENDERPROTECTIONTYPE DefenderProtectionType = iota
    // Block functionality.
    ENABLE_DEFENDERPROTECTIONTYPE
    // Allow functionality but generate logs.
    AUDITMODE_DEFENDERPROTECTIONTYPE
    // Warning message to end user with ability to bypass block from attack surface reduction rule.
    WARN_DEFENDERPROTECTIONTYPE
    // Not configured.
    NOTCONFIGURED_DEFENDERPROTECTIONTYPE
)

func (i DefenderProtectionType) String() string {
    return []string{"userDefined", "enable", "auditMode", "warn", "notConfigured"}[i]
}
func ParseDefenderProtectionType(v string) (any, error) {
    result := USERDEFINED_DEFENDERPROTECTIONTYPE
    switch v {
        case "userDefined":
            result = USERDEFINED_DEFENDERPROTECTIONTYPE
        case "enable":
            result = ENABLE_DEFENDERPROTECTIONTYPE
        case "auditMode":
            result = AUDITMODE_DEFENDERPROTECTIONTYPE
        case "warn":
            result = WARN_DEFENDERPROTECTIONTYPE
        case "notConfigured":
            result = NOTCONFIGURED_DEFENDERPROTECTIONTYPE
        default:
            return nil, nil
    }
    return &result, nil
}
func SerializeDefenderProtectionType(values []DefenderProtectionType) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i DefenderProtectionType) isMultiValue() bool {
    return false
}
