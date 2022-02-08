package graph
import (
    "strings"
    "errors"
)
// 
type DiamondModel int

const (
    UNKNOWN_DIAMONDMODEL DiamondModel = iota
    ADVERSARY_DIAMONDMODEL
    CAPABILITY_DIAMONDMODEL
    INFRASTRUCTURE_DIAMONDMODEL
    VICTIM_DIAMONDMODEL
    UNKNOWNFUTUREVALUE_DIAMONDMODEL
)

func (i DiamondModel) String() string {
    return []string{"UNKNOWN", "ADVERSARY", "CAPABILITY", "INFRASTRUCTURE", "VICTIM", "UNKNOWNFUTUREVALUE"}[i]
}
func ParseDiamondModel(v string) (interface{}, error) {
    result := UNKNOWN_DIAMONDMODEL
    switch strings.ToUpper(v) {
        case "UNKNOWN":
            result = UNKNOWN_DIAMONDMODEL
        case "ADVERSARY":
            result = ADVERSARY_DIAMONDMODEL
        case "CAPABILITY":
            result = CAPABILITY_DIAMONDMODEL
        case "INFRASTRUCTURE":
            result = INFRASTRUCTURE_DIAMONDMODEL
        case "VICTIM":
            result = VICTIM_DIAMONDMODEL
        case "UNKNOWNFUTUREVALUE":
            result = UNKNOWNFUTUREVALUE_DIAMONDMODEL
        default:
            return 0, errors.New("Unknown DiamondModel value: " + v)
    }
    return &result, nil
}
func SerializeDiamondModel(values []DiamondModel) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
