package graph
import (
    "strings"
    "errors"
)
// 
type MlClassificationMatchTolerance int

const (
    EXACT_MLCLASSIFICATIONMATCHTOLERANCE MlClassificationMatchTolerance = iota
    NEAR_MLCLASSIFICATIONMATCHTOLERANCE
)

func (i MlClassificationMatchTolerance) String() string {
    return []string{"EXACT", "NEAR"}[i]
}
func ParseMlClassificationMatchTolerance(v string) (interface{}, error) {
    result := EXACT_MLCLASSIFICATIONMATCHTOLERANCE
    switch strings.ToUpper(v) {
        case "EXACT":
            result = EXACT_MLCLASSIFICATIONMATCHTOLERANCE
        case "NEAR":
            result = NEAR_MLCLASSIFICATIONMATCHTOLERANCE
        default:
            return 0, errors.New("Unknown MlClassificationMatchTolerance value: " + v)
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
