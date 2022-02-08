package graph
import (
    "strings"
    "errors"
)
// 
type TeamworkSupportedClient int

const (
    UNKNOWN_TEAMWORKSUPPORTEDCLIENT TeamworkSupportedClient = iota
    SKYPEDEFAULTANDTEAMS_TEAMWORKSUPPORTEDCLIENT
    TEAMSDEFAULTANDSKYPE_TEAMWORKSUPPORTEDCLIENT
    SKYPEONLY_TEAMWORKSUPPORTEDCLIENT
    TEAMSONLY_TEAMWORKSUPPORTEDCLIENT
    UNKNOWNFUTUREVALUE_TEAMWORKSUPPORTEDCLIENT
)

func (i TeamworkSupportedClient) String() string {
    return []string{"UNKNOWN", "SKYPEDEFAULTANDTEAMS", "TEAMSDEFAULTANDSKYPE", "SKYPEONLY", "TEAMSONLY", "UNKNOWNFUTUREVALUE"}[i]
}
func ParseTeamworkSupportedClient(v string) (interface{}, error) {
    result := UNKNOWN_TEAMWORKSUPPORTEDCLIENT
    switch strings.ToUpper(v) {
        case "UNKNOWN":
            result = UNKNOWN_TEAMWORKSUPPORTEDCLIENT
        case "SKYPEDEFAULTANDTEAMS":
            result = SKYPEDEFAULTANDTEAMS_TEAMWORKSUPPORTEDCLIENT
        case "TEAMSDEFAULTANDSKYPE":
            result = TEAMSDEFAULTANDSKYPE_TEAMWORKSUPPORTEDCLIENT
        case "SKYPEONLY":
            result = SKYPEONLY_TEAMWORKSUPPORTEDCLIENT
        case "TEAMSONLY":
            result = TEAMSONLY_TEAMWORKSUPPORTEDCLIENT
        case "UNKNOWNFUTUREVALUE":
            result = UNKNOWNFUTUREVALUE_TEAMWORKSUPPORTEDCLIENT
        default:
            return 0, errors.New("Unknown TeamworkSupportedClient value: " + v)
    }
    return &result, nil
}
func SerializeTeamworkSupportedClient(values []TeamworkSupportedClient) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
