package graph
import (
    "strings"
    "errors"
)
// 
type TeamworkSoftwareType int

const (
    ADMINAGENT_TEAMWORKSOFTWARETYPE TeamworkSoftwareType = iota
    OPERATINGSYSTEM_TEAMWORKSOFTWARETYPE
    TEAMSCLIENT_TEAMWORKSOFTWARETYPE
    FIRMWARE_TEAMWORKSOFTWARETYPE
    PARTNERAGENT_TEAMWORKSOFTWARETYPE
    COMPANYPORTAL_TEAMWORKSOFTWARETYPE
    UNKNOWNFUTUREVALUE_TEAMWORKSOFTWARETYPE
)

func (i TeamworkSoftwareType) String() string {
    return []string{"ADMINAGENT", "OPERATINGSYSTEM", "TEAMSCLIENT", "FIRMWARE", "PARTNERAGENT", "COMPANYPORTAL", "UNKNOWNFUTUREVALUE"}[i]
}
func ParseTeamworkSoftwareType(v string) (interface{}, error) {
    result := ADMINAGENT_TEAMWORKSOFTWARETYPE
    switch strings.ToUpper(v) {
        case "ADMINAGENT":
            result = ADMINAGENT_TEAMWORKSOFTWARETYPE
        case "OPERATINGSYSTEM":
            result = OPERATINGSYSTEM_TEAMWORKSOFTWARETYPE
        case "TEAMSCLIENT":
            result = TEAMSCLIENT_TEAMWORKSOFTWARETYPE
        case "FIRMWARE":
            result = FIRMWARE_TEAMWORKSOFTWARETYPE
        case "PARTNERAGENT":
            result = PARTNERAGENT_TEAMWORKSOFTWARETYPE
        case "COMPANYPORTAL":
            result = COMPANYPORTAL_TEAMWORKSOFTWARETYPE
        case "UNKNOWNFUTUREVALUE":
            result = UNKNOWNFUTUREVALUE_TEAMWORKSOFTWARETYPE
        default:
            return 0, errors.New("Unknown TeamworkSoftwareType value: " + v)
    }
    return &result, nil
}
func SerializeTeamworkSoftwareType(values []TeamworkSoftwareType) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
