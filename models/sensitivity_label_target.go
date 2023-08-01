package models
import (
    "errors"
)
// 
type SensitivityLabelTarget int

const (
    EMAIL_SENSITIVITYLABELTARGET SensitivityLabelTarget = iota
    SITE_SENSITIVITYLABELTARGET
    UNIFIEDGROUP_SENSITIVITYLABELTARGET
    TEAMWORK_SENSITIVITYLABELTARGET
    UNKNOWNFUTUREVALUE_SENSITIVITYLABELTARGET
)

func (i SensitivityLabelTarget) String() string {
    return []string{"email", "site", "unifiedGroup", "teamwork", "unknownFutureValue"}[i]
}
func ParseSensitivityLabelTarget(v string) (any, error) {
    result := EMAIL_SENSITIVITYLABELTARGET
    switch v {
        case "email":
            result = EMAIL_SENSITIVITYLABELTARGET
        case "site":
            result = SITE_SENSITIVITYLABELTARGET
        case "unifiedGroup":
            result = UNIFIEDGROUP_SENSITIVITYLABELTARGET
        case "teamwork":
            result = TEAMWORK_SENSITIVITYLABELTARGET
        case "unknownFutureValue":
            result = UNKNOWNFUTUREVALUE_SENSITIVITYLABELTARGET
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
