package graph
import (
    "strings"
    "errors"
)
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
    switch strings.ToUpper(v) {
        case "UNKNOWN":
            return UNKNOWN_DIAMONDMODEL, nil
        case "ADVERSARY":
            return ADVERSARY_DIAMONDMODEL, nil
        case "CAPABILITY":
            return CAPABILITY_DIAMONDMODEL, nil
        case "INFRASTRUCTURE":
            return INFRASTRUCTURE_DIAMONDMODEL, nil
        case "VICTIM":
            return VICTIM_DIAMONDMODEL, nil
        case "UNKNOWNFUTUREVALUE":
            return UNKNOWNFUTUREVALUE_DIAMONDMODEL, nil
    }
    return 0, errors.New("Unknown DiamondModel value: " + v)
}
func SerializeDiamondModel(values []DiamondModel) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
