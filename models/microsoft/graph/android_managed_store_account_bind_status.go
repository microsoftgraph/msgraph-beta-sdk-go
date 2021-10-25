package graph
import (
    "strings"
    "errors"
)
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
    switch strings.ToUpper(v) {
        case "NOTBOUND":
            return NOTBOUND_ANDROIDMANAGEDSTOREACCOUNTBINDSTATUS, nil
        case "BOUND":
            return BOUND_ANDROIDMANAGEDSTOREACCOUNTBINDSTATUS, nil
        case "BOUNDANDVALIDATED":
            return BOUNDANDVALIDATED_ANDROIDMANAGEDSTOREACCOUNTBINDSTATUS, nil
        case "UNBINDING":
            return UNBINDING_ANDROIDMANAGEDSTOREACCOUNTBINDSTATUS, nil
    }
    return 0, errors.New("Unknown AndroidManagedStoreAccountBindStatus value: " + v)
}
func SerializeAndroidManagedStoreAccountBindStatus(values []AndroidManagedStoreAccountBindStatus) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
