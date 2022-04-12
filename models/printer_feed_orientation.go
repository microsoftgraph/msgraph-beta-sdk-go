package models
import (
    "strings"
    "errors"
)
// Provides operations to manage the deviceManagement singleton.
type PrinterFeedOrientation int

const (
    LONGEDGEFIRST_PRINTERFEEDORIENTATION PrinterFeedOrientation = iota
    SHORTEDGEFIRST_PRINTERFEEDORIENTATION
)

func (i PrinterFeedOrientation) String() string {
    return []string{"LONGEDGEFIRST", "SHORTEDGEFIRST"}[i]
}
func ParsePrinterFeedOrientation(v string) (interface{}, error) {
    result := LONGEDGEFIRST_PRINTERFEEDORIENTATION
    switch strings.ToUpper(v) {
        case "LONGEDGEFIRST":
            result = LONGEDGEFIRST_PRINTERFEEDORIENTATION
        case "SHORTEDGEFIRST":
            result = SHORTEDGEFIRST_PRINTERFEEDORIENTATION
        default:
            return 0, errors.New("Unknown PrinterFeedOrientation value: " + v)
    }
    return &result, nil
}
func SerializePrinterFeedOrientation(values []PrinterFeedOrientation) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
