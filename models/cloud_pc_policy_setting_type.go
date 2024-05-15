package models
import (
    "math"
    "strings"
)
type CloudPcPolicySettingType int

const (
    REGION_CLOUDPCPOLICYSETTINGTYPE = 1
    SINGLESIGNON_CLOUDPCPOLICYSETTINGTYPE = 2
    UNKNOWNFUTUREVALUE_CLOUDPCPOLICYSETTINGTYPE = 4
)

func (i CloudPcPolicySettingType) String() string {
    var values []string
    options := []string{"region", "singleSignOn", "unknownFutureValue"}
    for p := 0; p < 3; p++ {
        mantis := CloudPcPolicySettingType(int(math.Pow(2, float64(p))))
        if i&mantis == mantis {
            values = append(values, options[p])
        }
    }
    return strings.Join(values, ",")
}
func ParseCloudPcPolicySettingType(v string) (any, error) {
    var result CloudPcPolicySettingType
    values := strings.Split(v, ",")
    for _, str := range values {
        switch str {
            case "region":
                result |= REGION_CLOUDPCPOLICYSETTINGTYPE
            case "singleSignOn":
                result |= SINGLESIGNON_CLOUDPCPOLICYSETTINGTYPE
            case "unknownFutureValue":
                result |= UNKNOWNFUTUREVALUE_CLOUDPCPOLICYSETTINGTYPE
            default:
                return nil, nil
        }
    }
    return &result, nil
}
func SerializeCloudPcPolicySettingType(values []CloudPcPolicySettingType) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i CloudPcPolicySettingType) isMultiValue() bool {
    return true
}
