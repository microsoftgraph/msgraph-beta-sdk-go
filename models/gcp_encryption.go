package models
import (
    "errors"
)
// 
type GcpEncryption int

const (
    GOOGLE_GCPENCRYPTION GcpEncryption = iota
    CUSTOMER_GCPENCRYPTION
    UNKNOWNFUTUREVALUE_GCPENCRYPTION
)

func (i GcpEncryption) String() string {
    return []string{"google", "customer", "unknownFutureValue"}[i]
}
func ParseGcpEncryption(v string) (any, error) {
    result := GOOGLE_GCPENCRYPTION
    switch v {
        case "google":
            result = GOOGLE_GCPENCRYPTION
        case "customer":
            result = CUSTOMER_GCPENCRYPTION
        case "unknownFutureValue":
            result = UNKNOWNFUTUREVALUE_GCPENCRYPTION
        default:
            return 0, errors.New("Unknown GcpEncryption value: " + v)
    }
    return &result, nil
}
func SerializeGcpEncryption(values []GcpEncryption) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i GcpEncryption) isMultiValue() bool {
    return false
}
