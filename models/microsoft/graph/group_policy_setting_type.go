package graph
import (
    "strings"
    "errors"
)
type GroupPolicySettingType int

const (
    UNKNOWN_GROUPPOLICYSETTINGTYPE GroupPolicySettingType = iota
    POLICY_GROUPPOLICYSETTINGTYPE
    ACCOUNT_GROUPPOLICYSETTINGTYPE
    SECURITYOPTIONS_GROUPPOLICYSETTINGTYPE
    USERRIGHTSASSIGNMENT_GROUPPOLICYSETTINGTYPE
    AUDITSETTING_GROUPPOLICYSETTINGTYPE
    WINDOWSFIREWALLSETTINGS_GROUPPOLICYSETTINGTYPE
)

func (i GroupPolicySettingType) String() string {
    return []string{"UNKNOWN", "POLICY", "ACCOUNT", "SECURITYOPTIONS", "USERRIGHTSASSIGNMENT", "AUDITSETTING", "WINDOWSFIREWALLSETTINGS"}[i]
}
func ParseGroupPolicySettingType(v string) (interface{}, error) {
    switch strings.ToUpper(v) {
        case "UNKNOWN":
            return UNKNOWN_GROUPPOLICYSETTINGTYPE, nil
        case "POLICY":
            return POLICY_GROUPPOLICYSETTINGTYPE, nil
        case "ACCOUNT":
            return ACCOUNT_GROUPPOLICYSETTINGTYPE, nil
        case "SECURITYOPTIONS":
            return SECURITYOPTIONS_GROUPPOLICYSETTINGTYPE, nil
        case "USERRIGHTSASSIGNMENT":
            return USERRIGHTSASSIGNMENT_GROUPPOLICYSETTINGTYPE, nil
        case "AUDITSETTING":
            return AUDITSETTING_GROUPPOLICYSETTINGTYPE, nil
        case "WINDOWSFIREWALLSETTINGS":
            return WINDOWSFIREWALLSETTINGS_GROUPPOLICYSETTINGTYPE, nil
    }
    return 0, errors.New("Unknown GroupPolicySettingType value: " + v)
}
func SerializeGroupPolicySettingType(values []GroupPolicySettingType) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
