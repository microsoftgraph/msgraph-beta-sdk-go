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
    switch strings.ToUpper(v) {
        case "TWOSIDEDLONGEDGE":
            return TWOSIDEDLONGEDGE_PRINTDUPLEXCONFIGURATION, nil
        case "TWOSIDEDSHORTEDGE":
            return TWOSIDEDSHORTEDGE_PRINTDUPLEXCONFIGURATION, nil
        case "ONESIDED":
            return ONESIDED_PRINTDUPLEXCONFIGURATION, nil
    }
    return 0, errors.New("Unknown PrintDuplexConfiguration value: " + v)
}
func SerializePrintDuplexConfiguration(values []PrintDuplexConfiguration) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
