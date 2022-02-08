package termstore
import (
    "strings"
    "errors"
)
// 
type RelationType int

const (
    PIN_RELATIONTYPE RelationType = iota
    REUSE_RELATIONTYPE
)

func (i RelationType) String() string {
    return []string{"PIN", "REUSE"}[i]
}
func ParseRelationType(v string) (interface{}, error) {
    result := PIN_RELATIONTYPE
    switch strings.ToUpper(v) {
        case "PIN":
            result = PIN_RELATIONTYPE
        case "REUSE":
            result = REUSE_RELATIONTYPE
        default:
            return 0, errors.New("Unknown RelationType value: " + v)
    }
    return &result, nil
}
func SerializeRelationType(values []RelationType) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
