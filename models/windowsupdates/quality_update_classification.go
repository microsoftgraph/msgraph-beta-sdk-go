package windowsupdates
import (
    "errors"
)
// 
type QualityUpdateClassification int

const (
    ALL_QUALITYUPDATECLASSIFICATION QualityUpdateClassification = iota
    SECURITY_QUALITYUPDATECLASSIFICATION
    NONSECURITY_QUALITYUPDATECLASSIFICATION
    UNKNOWNFUTUREVALUE_QUALITYUPDATECLASSIFICATION
)

func (i QualityUpdateClassification) String() string {
    return []string{"all", "security", "nonSecurity", "unknownFutureValue"}[i]
}
func ParseQualityUpdateClassification(v string) (any, error) {
    result := ALL_QUALITYUPDATECLASSIFICATION
    switch v {
        case "all":
            result = ALL_QUALITYUPDATECLASSIFICATION
        case "security":
            result = SECURITY_QUALITYUPDATECLASSIFICATION
        case "nonSecurity":
            result = NONSECURITY_QUALITYUPDATECLASSIFICATION
        case "unknownFutureValue":
            result = UNKNOWNFUTUREVALUE_QUALITYUPDATECLASSIFICATION
        default:
            return 0, errors.New("Unknown QualityUpdateClassification value: " + v)
    }
    return &result, nil
}
func SerializeQualityUpdateClassification(values []QualityUpdateClassification) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i QualityUpdateClassification) isMultiValue() bool {
    return false
}
