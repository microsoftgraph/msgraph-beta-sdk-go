package graph
import (
    "strings"
    "errors"
)
type PrinterFeedDirection int

const (
    LONGEDGEFIRST_PRINTERFEEDDIRECTION PrinterFeedDirection = iota
    SHORTEDGEFIRST_PRINTERFEEDDIRECTION
)

func (i PrinterFeedDirection) String() string {
    return []string{"LONGEDGEFIRST", "SHORTEDGEFIRST"}[i]
}
func ParsePrinterFeedDirection(v string) (interface{}, error) {
    switch strings.ToUpper(v) {
        case "LONGEDGEFIRST":
            return LONGEDGEFIRST_PRINTERFEEDDIRECTION, nil
        case "SHORTEDGEFIRST":
            return SHORTEDGEFIRST_PRINTERFEEDDIRECTION, nil
    }
    return 0, errors.New("Unknown PrinterFeedDirection value: " + v)
}
func SerializePrinterFeedDirection(values []PrinterFeedDirection) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
