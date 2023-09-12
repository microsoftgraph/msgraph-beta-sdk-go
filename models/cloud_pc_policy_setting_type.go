package models
import (
    "errors"
    "strings"
)
// 
type CloudPcPolicySettingType int

const (
    REGION_CLOUDPCPOLICYSETTINGTYPE CloudPcPolicySettingType = iota
    SINGLESIGNON_CLOUDPCPOLICYSETTINGTYPE
    UNKNOWNFUTUREVALUE_CLOUDPCPOLICYSETTINGTYPE
)

func (i CloudPcPolicySettingType) String() string {
    var values []string
    for p := CloudPcPolicySettingType(1); p <= UNKNOWNFUTUREVALUE_CLOUDPCPOLICYSETTINGTYPE; p <<= 1 {
        if i&p == p {
            values = append(values, []string{"region", "singleSignOn", "unknownFutureValue"}[p])
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
                return 0, errors.New("Unknown CloudPcPolicySettingType value: " + v)
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
