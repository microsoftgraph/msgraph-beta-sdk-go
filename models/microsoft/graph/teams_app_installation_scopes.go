package graph
import (
    "strings"
    "errors"
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
    return []string{"TEAM", "GROUPCHAT", "PERSONAL", "UNKNOWNFUTUREVALUE"}[i]
}
func ParseTeamsAppInstallationScopes(v string) (interface{}, error) {
    result := TEAM_TEAMSAPPINSTALLATIONSCOPES
    switch strings.ToUpper(v) {
        case "TEAM":
            result = TEAM_TEAMSAPPINSTALLATIONSCOPES
        case "GROUPCHAT":
            result = GROUPCHAT_TEAMSAPPINSTALLATIONSCOPES
        case "PERSONAL":
            result = PERSONAL_TEAMSAPPINSTALLATIONSCOPES
        case "UNKNOWNFUTUREVALUE":
            result = UNKNOWNFUTUREVALUE_TEAMSAPPINSTALLATIONSCOPES
        default:
            return 0, errors.New("Unknown TeamsAppInstallationScopes value: " + v)
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
