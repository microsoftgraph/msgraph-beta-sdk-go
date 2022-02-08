package graph
import (
    "strings"
    "errors"
)
// 
type RestoreTimeRange int

const (
    BEFORE_RESTORETIMERANGE RestoreTimeRange = iota
    AFTER_RESTORETIMERANGE
    BEFOREORAFTER_RESTORETIMERANGE
    UNKNOWNFUTUREVALUE_RESTORETIMERANGE
)

func (i RestoreTimeRange) String() string {
    return []string{"BEFORE", "AFTER", "BEFOREORAFTER", "UNKNOWNFUTUREVALUE"}[i]
}
func ParseRestoreTimeRange(v string) (interface{}, error) {
    result := BEFORE_RESTORETIMERANGE
    switch strings.ToUpper(v) {
        case "BEFORE":
            result = BEFORE_RESTORETIMERANGE
        case "AFTER":
            result = AFTER_RESTORETIMERANGE
        case "BEFOREORAFTER":
            result = BEFOREORAFTER_RESTORETIMERANGE
        case "UNKNOWNFUTUREVALUE":
            result = UNKNOWNFUTUREVALUE_RESTORETIMERANGE
        default:
            return 0, errors.New("Unknown RestoreTimeRange value: " + v)
    }
    return &result, nil
}
func SerializeRestoreTimeRange(values []RestoreTimeRange) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
