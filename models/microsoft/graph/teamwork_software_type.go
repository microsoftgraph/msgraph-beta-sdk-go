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
    switch strings.ToUpper(v) {
        case "ADMINAGENT":
            return ADMINAGENT_TEAMWORKSOFTWARETYPE, nil
        case "OPERATINGSYSTEM":
            return OPERATINGSYSTEM_TEAMWORKSOFTWARETYPE, nil
        case "TEAMSCLIENT":
            return TEAMSCLIENT_TEAMWORKSOFTWARETYPE, nil
        case "FIRMWARE":
            return FIRMWARE_TEAMWORKSOFTWARETYPE, nil
        case "PARTNERAGENT":
            return PARTNERAGENT_TEAMWORKSOFTWARETYPE, nil
        case "COMPANYPORTAL":
            return COMPANYPORTAL_TEAMWORKSOFTWARETYPE, nil
        case "UNKNOWNFUTUREVALUE":
            return UNKNOWNFUTUREVALUE_TEAMWORKSOFTWARETYPE, nil
    }
    return 0, errors.New("Unknown TeamworkSoftwareType value: " + v)
}
func SerializeTeamworkSoftwareType(values []TeamworkSoftwareType) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
