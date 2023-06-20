package models
import (
    "errors"
)
// 
type CloudPcPolicySettingType int

const (
    REGION_CLOUDPCPOLICYSETTINGTYPE CloudPcPolicySettingType = iota
    SINGLESIGNON_CLOUDPCPOLICYSETTINGTYPE
    UNKNOWNFUTUREVALUE_CLOUDPCPOLICYSETTINGTYPE
)

func (i CloudPcPolicySettingType) String() string {
    return []string{"region", "singleSignOn", "unknownFutureValue"}[i]
}
func ParseCloudPcPolicySettingType(v string) (any, error) {
    result := REGION_CLOUDPCPOLICYSETTINGTYPE
    switch v {
        case "region":
            result = REGION_CLOUDPCPOLICYSETTINGTYPE
        case "singleSignOn":
            result = SINGLESIGNON_CLOUDPCPOLICYSETTINGTYPE
        case "unknownFutureValue":
            result = UNKNOWNFUTUREVALUE_CLOUDPCPOLICYSETTINGTYPE
        default:
            return 0, errors.New("Unknown CloudPcPolicySettingType value: " + v)
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
