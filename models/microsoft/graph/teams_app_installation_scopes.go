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
    switch strings.ToUpper(v) {
        case "TEAM":
            return TEAM_TEAMSAPPINSTALLATIONSCOPES, nil
        case "GROUPCHAT":
            return GROUPCHAT_TEAMSAPPINSTALLATIONSCOPES, nil
        case "PERSONAL":
            return PERSONAL_TEAMSAPPINSTALLATIONSCOPES, nil
        case "UNKNOWNFUTUREVALUE":
            return UNKNOWNFUTUREVALUE_TEAMSAPPINSTALLATIONSCOPES, nil
    }
    return 0, errors.New("Unknown TeamsAppInstallationScopes value: " + v)
}
func SerializeTeamsAppInstallationScopes(values []TeamsAppInstallationScopes) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
