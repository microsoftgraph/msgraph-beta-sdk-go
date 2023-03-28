package security
import (
    "errors"
)
// 
type IndicatorSource int

const (
    MICROSOFTDEFENDERTHREATINTELLIGENCE_INDICATORSOURCE IndicatorSource = iota
    OPENSOURCEINTELLIGENCE_INDICATORSOURCE
    PUBLIC_INDICATORSOURCE
    UNKNOWNFUTUREVALUE_INDICATORSOURCE
)

func (i IndicatorSource) String() string {
    return []string{"microsoftDefenderThreatIntelligence", "openSourceIntelligence", "public", "unknownFutureValue"}[i]
}
func ParseIndicatorSource(v string) (any, error) {
    result := MICROSOFTDEFENDERTHREATINTELLIGENCE_INDICATORSOURCE
    switch v {
        case "microsoftDefenderThreatIntelligence":
            result = MICROSOFTDEFENDERTHREATINTELLIGENCE_INDICATORSOURCE
        case "openSourceIntelligence":
            result = OPENSOURCEINTELLIGENCE_INDICATORSOURCE
        case "public":
            result = PUBLIC_INDICATORSOURCE
        case "unknownFutureValue":
            result = UNKNOWNFUTUREVALUE_INDICATORSOURCE
        default:
            return 0, errors.New("Unknown IndicatorSource value: " + v)
    }
    return &result, nil
}
func SerializeIndicatorSource(values []IndicatorSource) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
