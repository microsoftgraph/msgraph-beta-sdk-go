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
    switch strings.ToUpper(v) {
        case "UNKNOWN":
            return UNKNOWN_TEAMWORKSUPPORTEDCLIENT, nil
        case "SKYPEDEFAULTANDTEAMS":
            return SKYPEDEFAULTANDTEAMS_TEAMWORKSUPPORTEDCLIENT, nil
        case "TEAMSDEFAULTANDSKYPE":
            return TEAMSDEFAULTANDSKYPE_TEAMWORKSUPPORTEDCLIENT, nil
        case "SKYPEONLY":
            return SKYPEONLY_TEAMWORKSUPPORTEDCLIENT, nil
        case "TEAMSONLY":
            return TEAMSONLY_TEAMWORKSUPPORTEDCLIENT, nil
        case "UNKNOWNFUTUREVALUE":
            return UNKNOWNFUTUREVALUE_TEAMWORKSUPPORTEDCLIENT, nil
    }
    return 0, errors.New("Unknown TeamworkSupportedClient value: " + v)
}
func SerializeTeamworkSupportedClient(values []TeamworkSupportedClient) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
