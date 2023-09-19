package models
import (
    "errors"
)
// Possible values for App update on Windows10 recurrence.
type Windows10AppsUpdateRecurrence int

const (
    // Default value, specifies a single occurence.
    NONE_WINDOWS10APPSUPDATERECURRENCE Windows10AppsUpdateRecurrence = iota
    // Daily.
    DAILY_WINDOWS10APPSUPDATERECURRENCE
    // Weekly.
    WEEKLY_WINDOWS10APPSUPDATERECURRENCE
    // Monthly.
    MONTHLY_WINDOWS10APPSUPDATERECURRENCE
)

func (i Windows10AppsUpdateRecurrence) String() string {
    return []string{"none", "daily", "weekly", "monthly"}[i]
}
func ParseWindows10AppsUpdateRecurrence(v string) (any, error) {
    result := NONE_WINDOWS10APPSUPDATERECURRENCE
    switch v {
        case "none":
            result = NONE_WINDOWS10APPSUPDATERECURRENCE
        case "daily":
            result = DAILY_WINDOWS10APPSUPDATERECURRENCE
        case "weekly":
            result = WEEKLY_WINDOWS10APPSUPDATERECURRENCE
        case "monthly":
            result = MONTHLY_WINDOWS10APPSUPDATERECURRENCE
        default:
            return 0, errors.New("Unknown Windows10AppsUpdateRecurrence value: " + v)
    }
    return &result, nil
}
func SerializeWindows10AppsUpdateRecurrence(values []Windows10AppsUpdateRecurrence) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i Windows10AppsUpdateRecurrence) isMultiValue() bool {
    return false
}
