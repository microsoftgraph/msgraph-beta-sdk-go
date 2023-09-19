package models
import (
    "errors"
    "strings"
)
// 
type MlClassificationMatchTolerance int

const (
    EXACT_MLCLASSIFICATIONMATCHTOLERANCE MlClassificationMatchTolerance = iota
    NEAR_MLCLASSIFICATIONMATCHTOLERANCE
)

func (i MlClassificationMatchTolerance) String() string {
    var values []string
    for p := MlClassificationMatchTolerance(1); p <= NEAR_MLCLASSIFICATIONMATCHTOLERANCE; p <<= 1 {
        if i&p == p {
            values = append(values, []string{"exact", "near"}[p])
        }
    }
    return strings.Join(values, ",")
}
func ParseMlClassificationMatchTolerance(v string) (any, error) {
    var result MlClassificationMatchTolerance
    values := strings.Split(v, ",")
    for _, str := range values {
        switch str {
            case "exact":
                result |= EXACT_MLCLASSIFICATIONMATCHTOLERANCE
            case "near":
                result |= NEAR_MLCLASSIFICATIONMATCHTOLERANCE
            default:
                return 0, errors.New("Unknown MlClassificationMatchTolerance value: " + v)
        }
    }
    return &result, nil
}
func SerializeMlClassificationMatchTolerance(values []MlClassificationMatchTolerance) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i MlClassificationMatchTolerance) isMultiValue() bool {
    return true
}
