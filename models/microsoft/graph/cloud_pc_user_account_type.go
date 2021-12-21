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
    switch strings.ToUpper(v) {
        case "STANDARDUSER":
            return STANDARDUSER_CLOUDPCUSERACCOUNTTYPE, nil
        case "ADMINISTRATOR":
            return ADMINISTRATOR_CLOUDPCUSERACCOUNTTYPE, nil
        case "UNKNOWNFUTUREVALUE":
            return UNKNOWNFUTUREVALUE_CLOUDPCUSERACCOUNTTYPE, nil
    }
    return 0, errors.New("Unknown CloudPcUserAccountType value: " + v)
}
func SerializeCloudPcUserAccountType(values []CloudPcUserAccountType) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
