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
    result := NOTBOUND_ANDROIDFORWORKBINDSTATUS
    switch strings.ToUpper(v) {
        case "NOTBOUND":
            result = NOTBOUND_ANDROIDFORWORKBINDSTATUS
        case "BOUND":
            result = BOUND_ANDROIDFORWORKBINDSTATUS
        case "BOUNDANDVALIDATED":
            result = BOUNDANDVALIDATED_ANDROIDFORWORKBINDSTATUS
        case "UNBINDING":
            result = UNBINDING_ANDROIDFORWORKBINDSTATUS
        default:
            return 0, errors.New("Unknown AndroidForWorkBindStatus value: " + v)
    }
    return &result, nil
}
func SerializeAndroidForWorkBindStatus(values []AndroidForWorkBindStatus) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
