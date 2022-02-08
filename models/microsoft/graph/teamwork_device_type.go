package graph
import (
    "strings"
    "errors"
)
// 
type TeamworkDeviceType int

const (
    UNKNOWN_TEAMWORKDEVICETYPE TeamworkDeviceType = iota
    IPPHONE_TEAMWORKDEVICETYPE
    TEAMSROOM_TEAMWORKDEVICETYPE
    SURFACEHUB_TEAMWORKDEVICETYPE
    COLLABORATIONBAR_TEAMWORKDEVICETYPE
    TEAMSDISPLAY_TEAMWORKDEVICETYPE
    TOUCHCONSOLE_TEAMWORKDEVICETYPE
    LOWCOSTPHONE_TEAMWORKDEVICETYPE
    TEAMSPANEL_TEAMWORKDEVICETYPE
    SIP_TEAMWORKDEVICETYPE
    UNKNOWNFUTUREVALUE_TEAMWORKDEVICETYPE
)

func (i TeamworkDeviceType) String() string {
    return []string{"UNKNOWN", "IPPHONE", "TEAMSROOM", "SURFACEHUB", "COLLABORATIONBAR", "TEAMSDISPLAY", "TOUCHCONSOLE", "LOWCOSTPHONE", "TEAMSPANEL", "SIP", "UNKNOWNFUTUREVALUE"}[i]
}
func ParseTeamworkDeviceType(v string) (interface{}, error) {
    result := UNKNOWN_TEAMWORKDEVICETYPE
    switch strings.ToUpper(v) {
        case "UNKNOWN":
            result = UNKNOWN_TEAMWORKDEVICETYPE
        case "IPPHONE":
            result = IPPHONE_TEAMWORKDEVICETYPE
        case "TEAMSROOM":
            result = TEAMSROOM_TEAMWORKDEVICETYPE
        case "SURFACEHUB":
            result = SURFACEHUB_TEAMWORKDEVICETYPE
        case "COLLABORATIONBAR":
            result = COLLABORATIONBAR_TEAMWORKDEVICETYPE
        case "TEAMSDISPLAY":
            result = TEAMSDISPLAY_TEAMWORKDEVICETYPE
        case "TOUCHCONSOLE":
            result = TOUCHCONSOLE_TEAMWORKDEVICETYPE
        case "LOWCOSTPHONE":
            result = LOWCOSTPHONE_TEAMWORKDEVICETYPE
        case "TEAMSPANEL":
            result = TEAMSPANEL_TEAMWORKDEVICETYPE
        case "SIP":
            result = SIP_TEAMWORKDEVICETYPE
        case "UNKNOWNFUTUREVALUE":
            result = UNKNOWNFUTUREVALUE_TEAMWORKDEVICETYPE
        default:
            return 0, errors.New("Unknown TeamworkDeviceType value: " + v)
    }
    return &result, nil
}
func SerializeTeamworkDeviceType(values []TeamworkDeviceType) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
