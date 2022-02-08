package graph
import (
    "strings"
    "errors"
)
// 
type AndroidManagedStoreAccountAppSyncStatus int

const (
    SUCCESS_ANDROIDMANAGEDSTOREACCOUNTAPPSYNCSTATUS AndroidManagedStoreAccountAppSyncStatus = iota
    CREDENTIALSNOTVALID_ANDROIDMANAGEDSTOREACCOUNTAPPSYNCSTATUS
    ANDROIDFORWORKAPIERROR_ANDROIDMANAGEDSTOREACCOUNTAPPSYNCSTATUS
    MANAGEMENTSERVICEERROR_ANDROIDMANAGEDSTOREACCOUNTAPPSYNCSTATUS
    UNKNOWNERROR_ANDROIDMANAGEDSTOREACCOUNTAPPSYNCSTATUS
    NONE_ANDROIDMANAGEDSTOREACCOUNTAPPSYNCSTATUS
)

func (i AndroidManagedStoreAccountAppSyncStatus) String() string {
    return []string{"SUCCESS", "CREDENTIALSNOTVALID", "ANDROIDFORWORKAPIERROR", "MANAGEMENTSERVICEERROR", "UNKNOWNERROR", "NONE"}[i]
}
func ParseAndroidManagedStoreAccountAppSyncStatus(v string) (interface{}, error) {
    result := SUCCESS_ANDROIDMANAGEDSTOREACCOUNTAPPSYNCSTATUS
    switch strings.ToUpper(v) {
        case "SUCCESS":
            result = SUCCESS_ANDROIDMANAGEDSTOREACCOUNTAPPSYNCSTATUS
        case "CREDENTIALSNOTVALID":
            result = CREDENTIALSNOTVALID_ANDROIDMANAGEDSTOREACCOUNTAPPSYNCSTATUS
        case "ANDROIDFORWORKAPIERROR":
            result = ANDROIDFORWORKAPIERROR_ANDROIDMANAGEDSTOREACCOUNTAPPSYNCSTATUS
        case "MANAGEMENTSERVICEERROR":
            result = MANAGEMENTSERVICEERROR_ANDROIDMANAGEDSTOREACCOUNTAPPSYNCSTATUS
        case "UNKNOWNERROR":
            result = UNKNOWNERROR_ANDROIDMANAGEDSTOREACCOUNTAPPSYNCSTATUS
        case "NONE":
            result = NONE_ANDROIDMANAGEDSTOREACCOUNTAPPSYNCSTATUS
        default:
            return 0, errors.New("Unknown AndroidManagedStoreAccountAppSyncStatus value: " + v)
    }
    return &result, nil
}
func SerializeAndroidManagedStoreAccountAppSyncStatus(values []AndroidManagedStoreAccountAppSyncStatus) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
