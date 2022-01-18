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
    switch strings.ToUpper(v) {
        case "BEFORE":
            return BEFORE_RESTORETIMERANGE, nil
        case "AFTER":
            return AFTER_RESTORETIMERANGE, nil
        case "BEFOREORAFTER":
            return BEFOREORAFTER_RESTORETIMERANGE, nil
        case "UNKNOWNFUTUREVALUE":
            return UNKNOWNFUTUREVALUE_RESTORETIMERANGE, nil
    }
    return 0, errors.New("Unknown RestoreTimeRange value: " + v)
}
func SerializeRestoreTimeRange(values []RestoreTimeRange) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
