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
    switch strings.ToUpper(v) {
        case "REGISTRATION":
            return REGISTRATION_FEATURETYPE, nil
        case "RESET":
            return RESET_FEATURETYPE, nil
        case "UNKNOWNFUTUREVALUE":
            return UNKNOWNFUTUREVALUE_FEATURETYPE, nil
    }
    return 0, errors.New("Unknown FeatureType value: " + v)
}
func SerializeFeatureType(values []FeatureType) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
