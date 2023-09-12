package models
import (
    "errors"
)
// 
type CloudPcUserAccountType int

const (
    STANDARDUSER_CLOUDPCUSERACCOUNTTYPE CloudPcUserAccountType = iota
    ADMINISTRATOR_CLOUDPCUSERACCOUNTTYPE
    UNKNOWNFUTUREVALUE_CLOUDPCUSERACCOUNTTYPE
)

func (i CloudPcUserAccountType) String() string {
    return []string{"standardUser", "administrator", "unknownFutureValue"}[i]
}
func ParseCloudPcUserAccountType(v string) (any, error) {
    result := STANDARDUSER_CLOUDPCUSERACCOUNTTYPE
    switch v {
        case "standardUser":
            result = STANDARDUSER_CLOUDPCUSERACCOUNTTYPE
        case "administrator":
            result = ADMINISTRATOR_CLOUDPCUSERACCOUNTTYPE
        case "unknownFutureValue":
            result = UNKNOWNFUTUREVALUE_CLOUDPCUSERACCOUNTTYPE
        default:
            return 0, errors.New("Unknown CloudPcUserAccountType value: " + v)
    }
    return &result, nil
}
func SerializeCloudPcUserAccountType(values []CloudPcUserAccountType) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i CloudPcUserAccountType) isMultiValue() bool {
    return false
}
