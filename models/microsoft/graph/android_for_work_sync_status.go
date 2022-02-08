package graph
import (
    "strings"
    "errors"
)
// 
type AndroidForWorkSyncStatus int

const (
    SUCCESS_ANDROIDFORWORKSYNCSTATUS AndroidForWorkSyncStatus = iota
    CREDENTIALSNOTVALID_ANDROIDFORWORKSYNCSTATUS
    ANDROIDFORWORKAPIERROR_ANDROIDFORWORKSYNCSTATUS
    MANAGEMENTSERVICEERROR_ANDROIDFORWORKSYNCSTATUS
    UNKNOWNERROR_ANDROIDFORWORKSYNCSTATUS
    NONE_ANDROIDFORWORKSYNCSTATUS
)

func (i AndroidForWorkSyncStatus) String() string {
    return []string{"SUCCESS", "CREDENTIALSNOTVALID", "ANDROIDFORWORKAPIERROR", "MANAGEMENTSERVICEERROR", "UNKNOWNERROR", "NONE"}[i]
}
func ParseAndroidForWorkSyncStatus(v string) (interface{}, error) {
    result := SUCCESS_ANDROIDFORWORKSYNCSTATUS
    switch strings.ToUpper(v) {
        case "SUCCESS":
            result = SUCCESS_ANDROIDFORWORKSYNCSTATUS
        case "CREDENTIALSNOTVALID":
            result = CREDENTIALSNOTVALID_ANDROIDFORWORKSYNCSTATUS
        case "ANDROIDFORWORKAPIERROR":
            result = ANDROIDFORWORKAPIERROR_ANDROIDFORWORKSYNCSTATUS
        case "MANAGEMENTSERVICEERROR":
            result = MANAGEMENTSERVICEERROR_ANDROIDFORWORKSYNCSTATUS
        case "UNKNOWNERROR":
            result = UNKNOWNERROR_ANDROIDFORWORKSYNCSTATUS
        case "NONE":
            result = NONE_ANDROIDFORWORKSYNCSTATUS
        default:
            return 0, errors.New("Unknown AndroidForWorkSyncStatus value: " + v)
    }
    return &result, nil
}
func SerializeAndroidForWorkSyncStatus(values []AndroidForWorkSyncStatus) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
