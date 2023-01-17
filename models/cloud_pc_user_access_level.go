package models
import (
    "errors"
)
// 
type CloudPcUserAccessLevel int

const (
    UNRESTRICTED_CLOUDPCUSERACCESSLEVEL CloudPcUserAccessLevel = iota
    RESTRICTED_CLOUDPCUSERACCESSLEVEL
    UNKNOWNFUTUREVALUE_CLOUDPCUSERACCESSLEVEL
)

func (i CloudPcUserAccessLevel) String() string {
    return []string{"unrestricted", "restricted", "unknownFutureValue"}[i]
}
func ParseCloudPcUserAccessLevel(v string) (any, error) {
    result := UNRESTRICTED_CLOUDPCUSERACCESSLEVEL
    switch v {
        case "unrestricted":
            result = UNRESTRICTED_CLOUDPCUSERACCESSLEVEL
        case "restricted":
            result = RESTRICTED_CLOUDPCUSERACCESSLEVEL
        case "unknownFutureValue":
            result = UNKNOWNFUTUREVALUE_CLOUDPCUSERACCESSLEVEL
        default:
            return 0, errors.New("Unknown CloudPcUserAccessLevel value: " + v)
    }
    return &result, nil
}
func SerializeCloudPcUserAccessLevel(values []CloudPcUserAccessLevel) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
