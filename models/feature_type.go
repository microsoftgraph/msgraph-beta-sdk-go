package models
import (
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
    return []string{"registration", "reset", "unknownFutureValue"}[i]
}
func ParseFeatureType(v string) (any, error) {
    result := REGISTRATION_FEATURETYPE
    switch v {
        case "registration":
            result = REGISTRATION_FEATURETYPE
        case "reset":
            result = RESET_FEATURETYPE
        case "unknownFutureValue":
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
func (i FeatureType) isMultiValue() bool {
    return false
}
