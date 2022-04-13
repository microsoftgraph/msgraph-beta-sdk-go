package models
import (
    "strings"
    "errors"
)
// Provides operations to call the getCloudPcReviewStatus method.
type CloudPcUserAccessLevel int

const (
    UNRESTRICTED_CLOUDPCUSERACCESSLEVEL CloudPcUserAccessLevel = iota
    RESTRICTED_CLOUDPCUSERACCESSLEVEL
    UNKNOWNFUTUREVALUE_CLOUDPCUSERACCESSLEVEL
)

func (i CloudPcUserAccessLevel) String() string {
    return []string{"UNRESTRICTED", "RESTRICTED", "UNKNOWNFUTUREVALUE"}[i]
}
func ParseCloudPcUserAccessLevel(v string) (interface{}, error) {
    result := UNRESTRICTED_CLOUDPCUSERACCESSLEVEL
    switch strings.ToUpper(v) {
        case "UNRESTRICTED":
            result = UNRESTRICTED_CLOUDPCUSERACCESSLEVEL
        case "RESTRICTED":
            result = RESTRICTED_CLOUDPCUSERACCESSLEVEL
        case "UNKNOWNFUTUREVALUE":
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
