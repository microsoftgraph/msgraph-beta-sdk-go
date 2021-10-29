package graph
import (
    "strings"
    "errors"
)
// 
type AndroidForWorkBindStatus int

const (
    NOTBOUND_ANDROIDFORWORKBINDSTATUS AndroidForWorkBindStatus = iota
    BOUND_ANDROIDFORWORKBINDSTATUS
    BOUNDANDVALIDATED_ANDROIDFORWORKBINDSTATUS
    UNBINDING_ANDROIDFORWORKBINDSTATUS
)

func (i AndroidForWorkBindStatus) String() string {
    return []string{"NOTBOUND", "BOUND", "BOUNDANDVALIDATED", "UNBINDING"}[i]
}
func ParseAndroidForWorkBindStatus(v string) (interface{}, error) {
    switch strings.ToUpper(v) {
        case "NOTBOUND":
            return NOTBOUND_ANDROIDFORWORKBINDSTATUS, nil
        case "BOUND":
            return BOUND_ANDROIDFORWORKBINDSTATUS, nil
        case "BOUNDANDVALIDATED":
            return BOUNDANDVALIDATED_ANDROIDFORWORKBINDSTATUS, nil
        case "UNBINDING":
            return UNBINDING_ANDROIDFORWORKBINDSTATUS, nil
    }
    return 0, errors.New("Unknown AndroidForWorkBindStatus value: " + v)
}
func SerializeAndroidForWorkBindStatus(values []AndroidForWorkBindStatus) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
