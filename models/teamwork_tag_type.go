package models
import (
    "errors"
)
// Provides operations to manage the collection of accessReview entities.
type TeamworkTagType int

const (
    STANDARD_TEAMWORKTAGTYPE TeamworkTagType = iota
)

func (i TeamworkTagType) String() string {
    return []string{"standard"}[i]
}
func ParseTeamworkTagType(v string) (interface{}, error) {
    result := STANDARD_TEAMWORKTAGTYPE
    switch v {
        case "standard":
            result = STANDARD_TEAMWORKTAGTYPE
        default:
            return 0, errors.New("Unknown TeamworkTagType value: " + v)
    }
    return &result, nil
}
func SerializeTeamworkTagType(values []TeamworkTagType) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
