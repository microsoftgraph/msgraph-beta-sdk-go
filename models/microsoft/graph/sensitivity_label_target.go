package graph
import (
    "strings"
    "errors"
)
type SensitivityLabelTarget int

const (
    EMAIL_SENSITIVITYLABELTARGET SensitivityLabelTarget = iota
    SITE_SENSITIVITYLABELTARGET
    UNIFIEDGROUP_SENSITIVITYLABELTARGET
    UNKNOWNFUTUREVALUE_SENSITIVITYLABELTARGET
)

func (i SensitivityLabelTarget) String() string {
    return []string{"EMAIL", "SITE", "UNIFIEDGROUP", "UNKNOWNFUTUREVALUE"}[i]
}
func ParseSensitivityLabelTarget(v string) (interface{}, error) {
    switch strings.ToUpper(v) {
        case "EMAIL":
            return EMAIL_SENSITIVITYLABELTARGET, nil
        case "SITE":
            return SITE_SENSITIVITYLABELTARGET, nil
        case "UNIFIEDGROUP":
            return UNIFIEDGROUP_SENSITIVITYLABELTARGET, nil
        case "UNKNOWNFUTUREVALUE":
            return UNKNOWNFUTUREVALUE_SENSITIVITYLABELTARGET, nil
    }
    return 0, errors.New("Unknown SensitivityLabelTarget value: " + v)
}
func SerializeSensitivityLabelTarget(values []SensitivityLabelTarget) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
