package models
import (
    "errors"
    "strings"
)
// 
type TeamsAppInstallationScopes int

const (
    TEAM_TEAMSAPPINSTALLATIONSCOPES TeamsAppInstallationScopes = iota
    GROUPCHAT_TEAMSAPPINSTALLATIONSCOPES
    PERSONAL_TEAMSAPPINSTALLATIONSCOPES
    UNKNOWNFUTUREVALUE_TEAMSAPPINSTALLATIONSCOPES
)

func (i TeamsAppInstallationScopes) String() string {
    var values []string
    for p := TeamsAppInstallationScopes(1); p <= UNKNOWNFUTUREVALUE_TEAMSAPPINSTALLATIONSCOPES; p <<= 1 {
        if i&p == p {
            values = append(values, []string{"team", "groupChat", "personal", "unknownFutureValue"}[p])
        }
    }
    return strings.Join(values, ",")
}
func ParseTeamsAppInstallationScopes(v string) (any, error) {
    var result TeamsAppInstallationScopes
    values := strings.Split(v, ",")
    for _, str := range values {
        switch str {
            case "team":
                result |= TEAM_TEAMSAPPINSTALLATIONSCOPES
            case "groupChat":
                result |= GROUPCHAT_TEAMSAPPINSTALLATIONSCOPES
            case "personal":
                result |= PERSONAL_TEAMSAPPINSTALLATIONSCOPES
            case "unknownFutureValue":
                result |= UNKNOWNFUTUREVALUE_TEAMSAPPINSTALLATIONSCOPES
            default:
                return 0, errors.New("Unknown TeamsAppInstallationScopes value: " + v)
        }
    }
    return &result, nil
}
func SerializeTeamsAppInstallationScopes(values []TeamsAppInstallationScopes) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i TeamsAppInstallationScopes) isMultiValue() bool {
    return true
}
