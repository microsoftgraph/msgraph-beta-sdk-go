package models
import (
    "errors"
)
type RestoreTimeRange int

const (
    BEFORE_RESTORETIMERANGE RestoreTimeRange = iota
    AFTER_RESTORETIMERANGE
    BEFOREORAFTER_RESTORETIMERANGE
    UNKNOWNFUTUREVALUE_RESTORETIMERANGE
)

func (i RestoreTimeRange) String() string {
    return []string{"before", "after", "beforeOrAfter", "unknownFutureValue"}[i]
}
func ParseRestoreTimeRange(v string) (any, error) {
    result := BEFORE_RESTORETIMERANGE
    switch v {
        case "before":
            result = BEFORE_RESTORETIMERANGE
        case "after":
            result = AFTER_RESTORETIMERANGE
        case "beforeOrAfter":
            result = BEFOREORAFTER_RESTORETIMERANGE
        case "unknownFutureValue":
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
func (i RestoreTimeRange) isMultiValue() bool {
    return false
}
