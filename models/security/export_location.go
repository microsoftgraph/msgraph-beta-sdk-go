package security
import (
    "errors"
    "strings"
)
// 
type ExportLocation int

const (
    RESPONSIVELOCATIONS_EXPORTLOCATION ExportLocation = iota
    NONRESPONSIVELOCATIONS_EXPORTLOCATION
    UNKNOWNFUTUREVALUE_EXPORTLOCATION
)

func (i ExportLocation) String() string {
    var values []string
    for p := ExportLocation(1); p <= UNKNOWNFUTUREVALUE_EXPORTLOCATION; p <<= 1 {
        if i&p == p {
            values = append(values, []string{"responsiveLocations", "nonresponsiveLocations", "unknownFutureValue"}[p])
        }
    }
    return strings.Join(values, ",")
}
func ParseExportLocation(v string) (any, error) {
    var result ExportLocation
    values := strings.Split(v, ",")
    for _, str := range values {
        switch str {
            case "responsiveLocations":
                result |= RESPONSIVELOCATIONS_EXPORTLOCATION
            case "nonresponsiveLocations":
                result |= NONRESPONSIVELOCATIONS_EXPORTLOCATION
            case "unknownFutureValue":
                result |= UNKNOWNFUTUREVALUE_EXPORTLOCATION
            default:
                return 0, errors.New("Unknown ExportLocation value: " + v)
        }
    }
    return &result, nil
}
func SerializeExportLocation(values []ExportLocation) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i ExportLocation) isMultiValue() bool {
    return true
}
