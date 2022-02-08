package graph
import (
    "strings"
    "errors"
)
// 
type AndroidManagedStoreAccountBindStatus int

const (
    NOTBOUND_ANDROIDMANAGEDSTOREACCOUNTBINDSTATUS AndroidManagedStoreAccountBindStatus = iota
    BOUND_ANDROIDMANAGEDSTOREACCOUNTBINDSTATUS
    BOUNDANDVALIDATED_ANDROIDMANAGEDSTOREACCOUNTBINDSTATUS
    UNBINDING_ANDROIDMANAGEDSTOREACCOUNTBINDSTATUS
)

func (i AndroidManagedStoreAccountBindStatus) String() string {
    return []string{"NOTBOUND", "BOUND", "BOUNDANDVALIDATED", "UNBINDING"}[i]
}
func ParseAndroidManagedStoreAccountBindStatus(v string) (interface{}, error) {
    result := NOTBOUND_ANDROIDMANAGEDSTOREACCOUNTBINDSTATUS
    switch strings.ToUpper(v) {
        case "NOTBOUND":
            result = NOTBOUND_ANDROIDMANAGEDSTOREACCOUNTBINDSTATUS
        case "BOUND":
            result = BOUND_ANDROIDMANAGEDSTOREACCOUNTBINDSTATUS
        case "BOUNDANDVALIDATED":
            result = BOUNDANDVALIDATED_ANDROIDMANAGEDSTOREACCOUNTBINDSTATUS
        case "UNBINDING":
            result = UNBINDING_ANDROIDMANAGEDSTOREACCOUNTBINDSTATUS
        default:
            return 0, errors.New("Unknown AndroidManagedStoreAccountBindStatus value: " + v)
    }
    return &result, nil
}
func SerializeAndroidManagedStoreAccountBindStatus(values []AndroidManagedStoreAccountBindStatus) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
