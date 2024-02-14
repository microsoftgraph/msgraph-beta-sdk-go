package models
import (
    "errors"
    "math"
    "strings"
)
type TeamsAppInstallationScopes int

const (
    TEAM_TEAMSAPPINSTALLATIONSCOPES = 1
    GROUPCHAT_TEAMSAPPINSTALLATIONSCOPES = 2
    PERSONAL_TEAMSAPPINSTALLATIONSCOPES = 4
    UNKNOWNFUTUREVALUE_TEAMSAPPINSTALLATIONSCOPES = 8
)

func (i TeamsAppInstallationScopes) String() string {
    var values []string
    options := []string{"team", "groupChat", "personal", "unknownFutureValue"}
    for p := 0; p < 4; p++ {
        mantis := TeamsAppInstallationScopes(int(math.Pow(2, float64(p))))
        if i&mantis == mantis {
            values = append(values, options[p])
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
