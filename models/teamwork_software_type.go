package models
import (
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
    return []string{"adminAgent", "operatingSystem", "teamsClient", "firmware", "partnerAgent", "companyPortal", "unknownFutureValue"}[i]
}
func ParseTeamworkSoftwareType(v string) (any, error) {
    result := ADMINAGENT_TEAMWORKSOFTWARETYPE
    switch v {
        case "adminAgent":
            result = ADMINAGENT_TEAMWORKSOFTWARETYPE
        case "operatingSystem":
            result = OPERATINGSYSTEM_TEAMWORKSOFTWARETYPE
        case "teamsClient":
            result = TEAMSCLIENT_TEAMWORKSOFTWARETYPE
        case "firmware":
            result = FIRMWARE_TEAMWORKSOFTWARETYPE
        case "partnerAgent":
            result = PARTNERAGENT_TEAMWORKSOFTWARETYPE
        case "companyPortal":
            result = COMPANYPORTAL_TEAMWORKSOFTWARETYPE
        case "unknownFutureValue":
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
func (i TeamworkSoftwareType) isMultiValue() bool {
    return false
}
