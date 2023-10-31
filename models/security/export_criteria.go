package security
import (
    "errors"
    "strings"
)
// 
type ExportCriteria int

const (
    SEARCHHITS_EXPORTCRITERIA ExportCriteria = iota
    PARTIALLYINDEXED_EXPORTCRITERIA
    UNKNOWNFUTUREVALUE_EXPORTCRITERIA
)

func (i ExportCriteria) String() string {
    var values []string
    for p := ExportCriteria(1); p <= UNKNOWNFUTUREVALUE_EXPORTCRITERIA; p <<= 1 {
        if i&p == p {
            values = append(values, []string{"searchHits", "partiallyIndexed", "unknownFutureValue"}[p])
        }
    }
    return strings.Join(values, ",")
}
func ParseExportCriteria(v string) (any, error) {
    var result ExportCriteria
    values := strings.Split(v, ",")
    for _, str := range values {
        switch str {
            case "searchHits":
                result |= SEARCHHITS_EXPORTCRITERIA
            case "partiallyIndexed":
                result |= PARTIALLYINDEXED_EXPORTCRITERIA
            case "unknownFutureValue":
                result |= UNKNOWNFUTUREVALUE_EXPORTCRITERIA
            default:
                return 0, errors.New("Unknown ExportCriteria value: " + v)
        }
    }
    return &result, nil
}
func SerializeExportCriteria(values []ExportCriteria) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i ExportCriteria) isMultiValue() bool {
    return true
}
