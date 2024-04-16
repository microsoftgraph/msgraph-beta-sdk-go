package industrydata
import (
    "errors"
)
type FilterOptions int

const (
    ORGEXTERNALID_FILTEROPTIONS FilterOptions = iota
    UNKNOWNFUTUREVALUE_FILTEROPTIONS
)

func (i FilterOptions) String() string {
    return []string{"orgExternalId", "unknownFutureValue"}[i]
}
func ParseFilterOptions(v string) (any, error) {
    result := ORGEXTERNALID_FILTEROPTIONS
    switch v {
        case "orgExternalId":
            result = ORGEXTERNALID_FILTEROPTIONS
        case "unknownFutureValue":
            result = UNKNOWNFUTUREVALUE_FILTEROPTIONS
        default:
            return 0, errors.New("Unknown FilterOptions value: " + v)
    }
    return &result, nil
}
func SerializeFilterOptions(values []FilterOptions) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i FilterOptions) isMultiValue() bool {
    return false
}
