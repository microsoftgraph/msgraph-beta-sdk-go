package graph
import (
    "strings"
    "errors"
)
// 
type TeamworkSoftwareFreshness int

const (
    UNKNOWN_TEAMWORKSOFTWAREFRESHNESS TeamworkSoftwareFreshness = iota
    LATEST_TEAMWORKSOFTWAREFRESHNESS
    UPDATEAVAILABLE_TEAMWORKSOFTWAREFRESHNESS
    UNKNOWNFUTUREVALUE_TEAMWORKSOFTWAREFRESHNESS
)

func (i TeamworkSoftwareFreshness) String() string {
    return []string{"UNKNOWN", "LATEST", "UPDATEAVAILABLE", "UNKNOWNFUTUREVALUE"}[i]
}
func ParseTeamworkSoftwareFreshness(v string) (interface{}, error) {
    result := UNKNOWN_TEAMWORKSOFTWAREFRESHNESS
    switch strings.ToUpper(v) {
        case "UNKNOWN":
            result = UNKNOWN_TEAMWORKSOFTWAREFRESHNESS
        case "LATEST":
            result = LATEST_TEAMWORKSOFTWAREFRESHNESS
        case "UPDATEAVAILABLE":
            result = UPDATEAVAILABLE_TEAMWORKSOFTWAREFRESHNESS
        case "UNKNOWNFUTUREVALUE":
            result = UNKNOWNFUTUREVALUE_TEAMWORKSOFTWAREFRESHNESS
        default:
            return 0, errors.New("Unknown TeamworkSoftwareFreshness value: " + v)
    }
    return &result, nil
}
func SerializeTeamworkSoftwareFreshness(values []TeamworkSoftwareFreshness) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
