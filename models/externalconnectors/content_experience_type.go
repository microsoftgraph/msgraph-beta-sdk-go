package externalconnectors
import (
    "errors"
    "math"
    "strings"
)
// 
type ContentExperienceType int

const (
    SEARCH_CONTENTEXPERIENCETYPE = 1
    COMPLIANCE_CONTENTEXPERIENCETYPE = 2
    UNKNOWNFUTUREVALUE_CONTENTEXPERIENCETYPE = 4
)

func (i ContentExperienceType) String() string {
    var values []string
    options := []string{"search", "compliance", "unknownFutureValue"}
    for p := 0; p < 3; p++ {
        mantis := ContentExperienceType(int(math.Pow(2, float64(p))))
        if i&mantis == mantis {
            values = append(values, options[p])
        }
    }
    return strings.Join(values, ",")
}
func ParseContentExperienceType(v string) (any, error) {
    var result ContentExperienceType
    values := strings.Split(v, ",")
    for _, str := range values {
        switch str {
            case "search":
                result |= SEARCH_CONTENTEXPERIENCETYPE
            case "compliance":
                result |= COMPLIANCE_CONTENTEXPERIENCETYPE
            case "unknownFutureValue":
                result |= UNKNOWNFUTUREVALUE_CONTENTEXPERIENCETYPE
            default:
                return 0, errors.New("Unknown ContentExperienceType value: " + v)
        }
    }
    return &result, nil
}
func SerializeContentExperienceType(values []ContentExperienceType) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i ContentExperienceType) isMultiValue() bool {
    return true
}
