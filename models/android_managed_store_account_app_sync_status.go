package models
// Sync status of the tenant with the Google EMM API
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
    return []string{"success", "credentialsNotValid", "androidForWorkApiError", "managementServiceError", "unknownError", "none"}[i]
}
func ParseAndroidManagedStoreAccountAppSyncStatus(v string) (any, error) {
    result := SUCCESS_ANDROIDMANAGEDSTOREACCOUNTAPPSYNCSTATUS
    switch v {
        case "success":
            result = SUCCESS_ANDROIDMANAGEDSTOREACCOUNTAPPSYNCSTATUS
        case "credentialsNotValid":
            result = CREDENTIALSNOTVALID_ANDROIDMANAGEDSTOREACCOUNTAPPSYNCSTATUS
        case "androidForWorkApiError":
            result = ANDROIDFORWORKAPIERROR_ANDROIDMANAGEDSTOREACCOUNTAPPSYNCSTATUS
        case "managementServiceError":
            result = MANAGEMENTSERVICEERROR_ANDROIDMANAGEDSTOREACCOUNTAPPSYNCSTATUS
        case "unknownError":
            result = UNKNOWNERROR_ANDROIDMANAGEDSTOREACCOUNTAPPSYNCSTATUS
        case "none":
            result = NONE_ANDROIDMANAGEDSTOREACCOUNTAPPSYNCSTATUS
        default:
            return nil, nil
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
func (i AndroidManagedStoreAccountAppSyncStatus) isMultiValue() bool {
    return false
}
