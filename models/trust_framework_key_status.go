package models
import (
    "errors"
)
type TrustFrameworkKeyStatus int

const (
    ENABLED_TRUSTFRAMEWORKKEYSTATUS TrustFrameworkKeyStatus = iota
    DISABLED_TRUSTFRAMEWORKKEYSTATUS
    UNKNOWNFUTUREVALUE_TRUSTFRAMEWORKKEYSTATUS
)

func (i TrustFrameworkKeyStatus) String() string {
    return []string{"enabled", "disabled", "unknownFutureValue"}[i]
}
func ParseTrustFrameworkKeyStatus(v string) (any, error) {
    result := ENABLED_TRUSTFRAMEWORKKEYSTATUS
    switch v {
        case "enabled":
            result = ENABLED_TRUSTFRAMEWORKKEYSTATUS
        case "disabled":
            result = DISABLED_TRUSTFRAMEWORKKEYSTATUS
        case "unknownFutureValue":
            result = UNKNOWNFUTUREVALUE_TRUSTFRAMEWORKKEYSTATUS
        default:
            return 0, errors.New("Unknown TrustFrameworkKeyStatus value: " + v)
    }
    return &result, nil
}
func SerializeTrustFrameworkKeyStatus(values []TrustFrameworkKeyStatus) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i TrustFrameworkKeyStatus) isMultiValue() bool {
    return false
}
