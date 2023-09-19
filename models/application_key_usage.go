package models
import (
    "errors"
)
// 
type ApplicationKeyUsage int

const (
    SIGN_APPLICATIONKEYUSAGE ApplicationKeyUsage = iota
    VERIFY_APPLICATIONKEYUSAGE
    UNKNOWNFUTUREVALUE_APPLICATIONKEYUSAGE
)

func (i ApplicationKeyUsage) String() string {
    return []string{"sign", "verify", "unknownFutureValue"}[i]
}
func ParseApplicationKeyUsage(v string) (any, error) {
    result := SIGN_APPLICATIONKEYUSAGE
    switch v {
        case "sign":
            result = SIGN_APPLICATIONKEYUSAGE
        case "verify":
            result = VERIFY_APPLICATIONKEYUSAGE
        case "unknownFutureValue":
            result = UNKNOWNFUTUREVALUE_APPLICATIONKEYUSAGE
        default:
            return 0, errors.New("Unknown ApplicationKeyUsage value: " + v)
    }
    return &result, nil
}
func SerializeApplicationKeyUsage(values []ApplicationKeyUsage) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i ApplicationKeyUsage) isMultiValue() bool {
    return false
}
