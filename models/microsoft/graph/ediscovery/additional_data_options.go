package ediscovery
import (
    "strings"
    "errors"
)
// 
type AdditionalDataOptions int

const (
    ALLVERSIONS_ADDITIONALDATAOPTIONS AdditionalDataOptions = iota
    LINKEDFILES_ADDITIONALDATAOPTIONS
    UNKNOWNFUTUREVALUE_ADDITIONALDATAOPTIONS
)

func (i AdditionalDataOptions) String() string {
    return []string{"ALLVERSIONS", "LINKEDFILES", "UNKNOWNFUTUREVALUE"}[i]
}
func ParseAdditionalDataOptions(v string) (interface{}, error) {
    result := ALLVERSIONS_ADDITIONALDATAOPTIONS
    switch strings.ToUpper(v) {
        case "ALLVERSIONS":
            result = ALLVERSIONS_ADDITIONALDATAOPTIONS
        case "LINKEDFILES":
            result = LINKEDFILES_ADDITIONALDATAOPTIONS
        case "UNKNOWNFUTUREVALUE":
            result = UNKNOWNFUTUREVALUE_ADDITIONALDATAOPTIONS
        default:
            return 0, errors.New("Unknown AdditionalDataOptions value: " + v)
    }
    return &result, nil
}
func SerializeAdditionalDataOptions(values []AdditionalDataOptions) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
