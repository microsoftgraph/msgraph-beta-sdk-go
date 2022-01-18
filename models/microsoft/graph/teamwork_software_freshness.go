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
    switch strings.ToUpper(v) {
        case "UNKNOWN":
            return UNKNOWN_TEAMWORKSOFTWAREFRESHNESS, nil
        case "LATEST":
            return LATEST_TEAMWORKSOFTWAREFRESHNESS, nil
        case "UPDATEAVAILABLE":
            return UPDATEAVAILABLE_TEAMWORKSOFTWAREFRESHNESS, nil
        case "UNKNOWNFUTUREVALUE":
            return UNKNOWNFUTUREVALUE_TEAMWORKSOFTWAREFRESHNESS, nil
    }
    return 0, errors.New("Unknown TeamworkSoftwareFreshness value: " + v)
}
func SerializeTeamworkSoftwareFreshness(values []TeamworkSoftwareFreshness) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
