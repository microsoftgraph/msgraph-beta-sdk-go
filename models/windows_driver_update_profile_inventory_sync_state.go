package models
import (
    "strings"
    "errors"
)
// Provides operations to manage the deviceManagement singleton.
type WindowsDriverUpdateProfileInventorySyncState int

const (
    PENDING_WINDOWSDRIVERUPDATEPROFILEINVENTORYSYNCSTATE WindowsDriverUpdateProfileInventorySyncState = iota
    SUCCESS_WINDOWSDRIVERUPDATEPROFILEINVENTORYSYNCSTATE
    FAILURE_WINDOWSDRIVERUPDATEPROFILEINVENTORYSYNCSTATE
)

func (i WindowsDriverUpdateProfileInventorySyncState) String() string {
    return []string{"PENDING", "SUCCESS", "FAILURE"}[i]
}
func ParseWindowsDriverUpdateProfileInventorySyncState(v string) (interface{}, error) {
    result := PENDING_WINDOWSDRIVERUPDATEPROFILEINVENTORYSYNCSTATE
    switch strings.ToUpper(v) {
        case "PENDING":
            result = PENDING_WINDOWSDRIVERUPDATEPROFILEINVENTORYSYNCSTATE
        case "SUCCESS":
            result = SUCCESS_WINDOWSDRIVERUPDATEPROFILEINVENTORYSYNCSTATE
        case "FAILURE":
            result = FAILURE_WINDOWSDRIVERUPDATEPROFILEINVENTORYSYNCSTATE
        default:
            return 0, errors.New("Unknown WindowsDriverUpdateProfileInventorySyncState value: " + v)
    }
    return &result, nil
}
func SerializeWindowsDriverUpdateProfileInventorySyncState(values []WindowsDriverUpdateProfileInventorySyncState) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
