package graph
import (
    "strings"
    "errors"
)
// 
type PrintDuplexConfiguration int

const (
    TWOSIDEDLONGEDGE_PRINTDUPLEXCONFIGURATION PrintDuplexConfiguration = iota
    TWOSIDEDSHORTEDGE_PRINTDUPLEXCONFIGURATION
    ONESIDED_PRINTDUPLEXCONFIGURATION
)

func (i PrintDuplexConfiguration) String() string {
    return []string{"TWOSIDEDLONGEDGE", "TWOSIDEDSHORTEDGE", "ONESIDED"}[i]
}
func ParsePrintDuplexConfiguration(v string) (interface{}, error) {
    result := TWOSIDEDLONGEDGE_PRINTDUPLEXCONFIGURATION
    switch strings.ToUpper(v) {
        case "TWOSIDEDLONGEDGE":
            result = TWOSIDEDLONGEDGE_PRINTDUPLEXCONFIGURATION
        case "TWOSIDEDSHORTEDGE":
            result = TWOSIDEDSHORTEDGE_PRINTDUPLEXCONFIGURATION
        case "ONESIDED":
            result = ONESIDED_PRINTDUPLEXCONFIGURATION
        default:
            return 0, errors.New("Unknown PrintDuplexConfiguration value: " + v)
    }
    return &result, nil
}
func SerializePrintDuplexConfiguration(values []PrintDuplexConfiguration) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
