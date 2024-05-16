package models
import (
    "math"
    "strings"
)
type SensitivityLabelTarget int

const (
    EMAIL_SENSITIVITYLABELTARGET = 1
    SITE_SENSITIVITYLABELTARGET = 2
    UNIFIEDGROUP_SENSITIVITYLABELTARGET = 4
    TEAMWORK_SENSITIVITYLABELTARGET = 8
    UNKNOWNFUTUREVALUE_SENSITIVITYLABELTARGET = 16
)

func (i SensitivityLabelTarget) String() string {
    var values []string
    options := []string{"email", "site", "unifiedGroup", "teamwork", "unknownFutureValue"}
    for p := 0; p < 5; p++ {
        mantis := SensitivityLabelTarget(int(math.Pow(2, float64(p))))
        if i&mantis == mantis {
            values = append(values, options[p])
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
                return nil, nil
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
