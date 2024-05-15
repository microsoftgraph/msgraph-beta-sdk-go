package models
// The scheduling priority options for downloading and preparing the requested mac OS update
type MacOSPriority int

const (
    // Indicates low scheduling priority for downloading and preparing the requested update
    LOW_MACOSPRIORITY MacOSPriority = iota
    // Indicates high scheduling priority for downloading and preparing the requested update
    HIGH_MACOSPRIORITY
    // Evolvable enumeration sentinel value. Do not use.
    UNKNOWNFUTUREVALUE_MACOSPRIORITY
)

func (i MacOSPriority) String() string {
    return []string{"low", "high", "unknownFutureValue"}[i]
}
func ParseMacOSPriority(v string) (any, error) {
    result := LOW_MACOSPRIORITY
    switch v {
        case "low":
            result = LOW_MACOSPRIORITY
        case "high":
            result = HIGH_MACOSPRIORITY
        case "unknownFutureValue":
            result = UNKNOWNFUTUREVALUE_MACOSPRIORITY
        default:
            return nil, nil
    }
    return &result, nil
}
func SerializeMacOSPriority(values []MacOSPriority) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i MacOSPriority) isMultiValue() bool {
    return false
}
