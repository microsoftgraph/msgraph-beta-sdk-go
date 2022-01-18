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
    switch strings.ToUpper(v) {
        case "UNKNOWN":
            return UNKNOWN_TEAMWORKDEVICETYPE, nil
        case "IPPHONE":
            return IPPHONE_TEAMWORKDEVICETYPE, nil
        case "TEAMSROOM":
            return TEAMSROOM_TEAMWORKDEVICETYPE, nil
        case "SURFACEHUB":
            return SURFACEHUB_TEAMWORKDEVICETYPE, nil
        case "COLLABORATIONBAR":
            return COLLABORATIONBAR_TEAMWORKDEVICETYPE, nil
        case "TEAMSDISPLAY":
            return TEAMSDISPLAY_TEAMWORKDEVICETYPE, nil
        case "TOUCHCONSOLE":
            return TOUCHCONSOLE_TEAMWORKDEVICETYPE, nil
        case "LOWCOSTPHONE":
            return LOWCOSTPHONE_TEAMWORKDEVICETYPE, nil
        case "TEAMSPANEL":
            return TEAMSPANEL_TEAMWORKDEVICETYPE, nil
        case "SIP":
            return SIP_TEAMWORKDEVICETYPE, nil
        case "UNKNOWNFUTUREVALUE":
            return UNKNOWNFUTUREVALUE_TEAMWORKDEVICETYPE, nil
    }
    return 0, errors.New("Unknown TeamworkDeviceType value: " + v)
}
func SerializeTeamworkDeviceType(values []TeamworkDeviceType) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
