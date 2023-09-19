package models
import (
    "errors"
    "strings"
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
    var values []string
    for p := SensitivityLabelTarget(1); p <= UNKNOWNFUTUREVALUE_SENSITIVITYLABELTARGET; p <<= 1 {
        if i&p == p {
            values = append(values, []string{"email", "site", "unifiedGroup", "teamwork", "unknownFutureValue"}[p])
        }
    }
    return strings.Join(values, ",")
}
func ParseSensitivityLabelTarget(v string) (any, error) {
    var result SensitivityLabelTarget
    values := strings.Split(v, ",")
    for _, str := range values {
        switch str {
            case "email":
                result |= EMAIL_SENSITIVITYLABELTARGET
            case "site":
                result |= SITE_SENSITIVITYLABELTARGET
            case "unifiedGroup":
                result |= UNIFIEDGROUP_SENSITIVITYLABELTARGET
            case "teamwork":
                result |= TEAMWORK_SENSITIVITYLABELTARGET
            case "unknownFutureValue":
                result |= UNKNOWNFUTUREVALUE_SENSITIVITYLABELTARGET
            default:
                return 0, errors.New("Unknown SensitivityLabelTarget value: " + v)
        }
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
func (i SensitivityLabelTarget) isMultiValue() bool {
    return true
}
