package graph
import (
    "strings"
    "errors"
)
// 
type SensitivityLabelTarget int

const (
    EMAIL_SENSITIVITYLABELTARGET SensitivityLabelTarget = iota
    SITE_SENSITIVITYLABELTARGET
    UNIFIEDGROUP_SENSITIVITYLABELTARGET
    UNKNOWNFUTUREVALUE_SENSITIVITYLABELTARGET
    TEAMWORK_SENSITIVITYLABELTARGET
)

func (i SensitivityLabelTarget) String() string {
    return []string{"EMAIL", "SITE", "UNIFIEDGROUP", "UNKNOWNFUTUREVALUE", "TEAMWORK"}[i]
}
func ParseSensitivityLabelTarget(v string) (interface{}, error) {
    result := EMAIL_SENSITIVITYLABELTARGET
    switch strings.ToUpper(v) {
        case "EMAIL":
            result = EMAIL_SENSITIVITYLABELTARGET
        case "SITE":
            result = SITE_SENSITIVITYLABELTARGET
        case "UNIFIEDGROUP":
            result = UNIFIEDGROUP_SENSITIVITYLABELTARGET
        case "UNKNOWNFUTUREVALUE":
            result = UNKNOWNFUTUREVALUE_SENSITIVITYLABELTARGET
        case "TEAMWORK":
            result = TEAMWORK_SENSITIVITYLABELTARGET
        default:
            return 0, errors.New("Unknown SensitivityLabelTarget value: " + v)
    }
    return &result, nil
}
func SerializeSensitivityLabelTarget(values []SensitivityLabelTarget) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
