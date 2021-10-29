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
    switch strings.ToUpper(v) {
        case "SUCCESS":
            return SUCCESS_ANDROIDFORWORKSYNCSTATUS, nil
        case "CREDENTIALSNOTVALID":
            return CREDENTIALSNOTVALID_ANDROIDFORWORKSYNCSTATUS, nil
        case "ANDROIDFORWORKAPIERROR":
            return ANDROIDFORWORKAPIERROR_ANDROIDFORWORKSYNCSTATUS, nil
        case "MANAGEMENTSERVICEERROR":
            return MANAGEMENTSERVICEERROR_ANDROIDFORWORKSYNCSTATUS, nil
        case "UNKNOWNERROR":
            return UNKNOWNERROR_ANDROIDFORWORKSYNCSTATUS, nil
        case "NONE":
            return NONE_ANDROIDFORWORKSYNCSTATUS, nil
    }
    return 0, errors.New("Unknown AndroidForWorkSyncStatus value: " + v)
}
func SerializeAndroidForWorkSyncStatus(values []AndroidForWorkSyncStatus) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
