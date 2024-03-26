package models
import (
    "errors"
)
type InsiderRiskLevel int

const (
    NONE_INSIDERRISKLEVEL InsiderRiskLevel = iota
    MINOR_INSIDERRISKLEVEL
    MODERATE_INSIDERRISKLEVEL
    ELEVATED_INSIDERRISKLEVEL
    UNKNOWNFUTUREVALUE_INSIDERRISKLEVEL
)

func (i InsiderRiskLevel) String() string {
    return []string{"none", "minor", "moderate", "elevated", "unknownFutureValue"}[i]
}
func ParseInsiderRiskLevel(v string) (any, error) {
    result := NONE_INSIDERRISKLEVEL
    switch v {
        case "none":
            result = NONE_INSIDERRISKLEVEL
        case "minor":
            result = MINOR_INSIDERRISKLEVEL
        case "moderate":
            result = MODERATE_INSIDERRISKLEVEL
        case "elevated":
            result = ELEVATED_INSIDERRISKLEVEL
        case "unknownFutureValue":
            result = UNKNOWNFUTUREVALUE_INSIDERRISKLEVEL
        default:
            return 0, errors.New("Unknown InsiderRiskLevel value: " + v)
    }
    return &result, nil
}
func SerializeInsiderRiskLevel(values []InsiderRiskLevel) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i InsiderRiskLevel) isMultiValue() bool {
    return false
}
