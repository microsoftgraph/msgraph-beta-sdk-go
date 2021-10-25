package graph
import (
    "strings"
    "errors"
)
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
    switch strings.ToUpper(v) {
        case "SUCCESS":
            return SUCCESS_ANDROIDMANAGEDSTOREACCOUNTAPPSYNCSTATUS, nil
        case "CREDENTIALSNOTVALID":
            return CREDENTIALSNOTVALID_ANDROIDMANAGEDSTOREACCOUNTAPPSYNCSTATUS, nil
        case "ANDROIDFORWORKAPIERROR":
            return ANDROIDFORWORKAPIERROR_ANDROIDMANAGEDSTOREACCOUNTAPPSYNCSTATUS, nil
        case "MANAGEMENTSERVICEERROR":
            return MANAGEMENTSERVICEERROR_ANDROIDMANAGEDSTOREACCOUNTAPPSYNCSTATUS, nil
        case "UNKNOWNERROR":
            return UNKNOWNERROR_ANDROIDMANAGEDSTOREACCOUNTAPPSYNCSTATUS, nil
        case "NONE":
            return NONE_ANDROIDMANAGEDSTOREACCOUNTAPPSYNCSTATUS, nil
    }
    return 0, errors.New("Unknown AndroidManagedStoreAccountAppSyncStatus value: " + v)
}
func SerializeAndroidManagedStoreAccountAppSyncStatus(values []AndroidManagedStoreAccountAppSyncStatus) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
