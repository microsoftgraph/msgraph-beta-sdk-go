package models
import (
    "errors"
)
// 
type TokenProtectionStatus int

const (
    NONE_TOKENPROTECTIONSTATUS TokenProtectionStatus = iota
    BOUND_TOKENPROTECTIONSTATUS
    UNBOUND_TOKENPROTECTIONSTATUS
    UNKNOWNFUTUREVALUE_TOKENPROTECTIONSTATUS
)

func (i TokenProtectionStatus) String() string {
    return []string{"none", "bound", "unbound", "unknownFutureValue"}[i]
}
func ParseTokenProtectionStatus(v string) (any, error) {
    result := NONE_TOKENPROTECTIONSTATUS
    switch v {
        case "none":
            result = NONE_TOKENPROTECTIONSTATUS
        case "bound":
            result = BOUND_TOKENPROTECTIONSTATUS
        case "unbound":
            result = UNBOUND_TOKENPROTECTIONSTATUS
        case "unknownFutureValue":
            result = UNKNOWNFUTUREVALUE_TOKENPROTECTIONSTATUS
        default:
            return 0, errors.New("Unknown TokenProtectionStatus value: " + v)
    }
    return &result, nil
}
func SerializeTokenProtectionStatus(values []TokenProtectionStatus) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i TokenProtectionStatus) isMultiValue() bool {
    return false
}
