package models
import (
    "math"
    "strings"
)
type MlClassificationMatchTolerance int

const (
    EXACT_MLCLASSIFICATIONMATCHTOLERANCE = 1
    NEAR_MLCLASSIFICATIONMATCHTOLERANCE = 2
)

func (i MlClassificationMatchTolerance) String() string {
    var values []string
    options := []string{"exact", "near"}
    for p := 0; p < 2; p++ {
        mantis := MlClassificationMatchTolerance(int(math.Pow(2, float64(p))))
        if i&mantis == mantis {
            values = append(values, options[p])
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
                return nil, nil
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
