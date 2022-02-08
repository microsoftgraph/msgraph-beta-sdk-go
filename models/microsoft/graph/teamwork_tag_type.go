package graph
import (
    "strings"
    "errors"
)
// 
type TeamworkTagType int

const (
    STANDARD_TEAMWORKTAGTYPE TeamworkTagType = iota
)

func (i TeamworkTagType) String() string {
    return []string{"STANDARD"}[i]
}
func ParseTeamworkTagType(v string) (interface{}, error) {
    result := STANDARD_TEAMWORKTAGTYPE
    switch strings.ToUpper(v) {
        case "STANDARD":
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
