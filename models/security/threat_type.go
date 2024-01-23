package security
import (
    "errors"
)
// 
type ThreatType int

const (
    UNKNOWN_THREATTYPE ThreatType = iota
    SPAM_THREATTYPE
    MALWARE_THREATTYPE
    PHISHING_THREATTYPE
    NONE_THREATTYPE
    UNKNOWNFUTUREVALUE_THREATTYPE
)

func (i ThreatType) String() string {
    return []string{"unknown", "spam", "malware", "phishing", "none", "unknownFutureValue"}[i]
}
func ParseThreatType(v string) (any, error) {
    result := UNKNOWN_THREATTYPE
    switch v {
        case "unknown":
            result = UNKNOWN_THREATTYPE
        case "spam":
            result = SPAM_THREATTYPE
        case "malware":
            result = MALWARE_THREATTYPE
        case "phishing":
            result = PHISHING_THREATTYPE
        case "none":
            result = NONE_THREATTYPE
        case "unknownFutureValue":
            result = UNKNOWNFUTUREVALUE_THREATTYPE
        default:
            return 0, errors.New("Unknown ThreatType value: " + v)
    }
    return &result, nil
}
func SerializeThreatType(values []ThreatType) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i ThreatType) isMultiValue() bool {
    return false
}
