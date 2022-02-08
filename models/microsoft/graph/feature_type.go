package graph
import (
    "strings"
    "errors"
)
// 
type FeatureType int

const (
    REGISTRATION_FEATURETYPE FeatureType = iota
    RESET_FEATURETYPE
    UNKNOWNFUTUREVALUE_FEATURETYPE
)

func (i FeatureType) String() string {
    return []string{"REGISTRATION", "RESET", "UNKNOWNFUTUREVALUE"}[i]
}
func ParseFeatureType(v string) (interface{}, error) {
    result := REGISTRATION_FEATURETYPE
    switch strings.ToUpper(v) {
        case "REGISTRATION":
            result = REGISTRATION_FEATURETYPE
        case "RESET":
            result = RESET_FEATURETYPE
        case "UNKNOWNFUTUREVALUE":
            result = UNKNOWNFUTUREVALUE_FEATURETYPE
        default:
            return 0, errors.New("Unknown FeatureType value: " + v)
    }
    return &result, nil
}
func SerializeFeatureType(values []FeatureType) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
