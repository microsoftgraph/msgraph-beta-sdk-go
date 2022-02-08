package graph
import (
    "strings"
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
    return []string{"STANDARDUSER", "ADMINISTRATOR", "UNKNOWNFUTUREVALUE"}[i]
}
func ParseCloudPcUserAccountType(v string) (interface{}, error) {
    result := STANDARDUSER_CLOUDPCUSERACCOUNTTYPE
    switch strings.ToUpper(v) {
        case "STANDARDUSER":
            result = STANDARDUSER_CLOUDPCUSERACCOUNTTYPE
        case "ADMINISTRATOR":
            result = ADMINISTRATOR_CLOUDPCUSERACCOUNTTYPE
        case "UNKNOWNFUTUREVALUE":
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
