package graph
import (
    "strings"
    "errors"
)
// 
type PrinterFeedDirection int

const (
    LONGEDGEFIRST_PRINTERFEEDDIRECTION PrinterFeedDirection = iota
    SHORTEDGEFIRST_PRINTERFEEDDIRECTION
)

func (i PrinterFeedDirection) String() string {
    return []string{"LONGEDGEFIRST", "SHORTEDGEFIRST"}[i]
}
func ParsePrinterFeedDirection(v string) (interface{}, error) {
    result := LONGEDGEFIRST_PRINTERFEEDDIRECTION
    switch strings.ToUpper(v) {
        case "LONGEDGEFIRST":
            result = LONGEDGEFIRST_PRINTERFEEDDIRECTION
        case "SHORTEDGEFIRST":
            result = SHORTEDGEFIRST_PRINTERFEEDDIRECTION
        default:
            return 0, errors.New("Unknown PrinterFeedDirection value: " + v)
    }
    return &result, nil
}
func SerializePrinterFeedDirection(values []PrinterFeedDirection) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
