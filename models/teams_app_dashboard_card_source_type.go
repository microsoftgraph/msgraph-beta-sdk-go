package models
import (
    "errors"
)
type TeamsAppDashboardCardSourceType int

const (
    BOT_TEAMSAPPDASHBOARDCARDSOURCETYPE TeamsAppDashboardCardSourceType = iota
    UNKNOWNFUTUREVALUE_TEAMSAPPDASHBOARDCARDSOURCETYPE
)

func (i TeamsAppDashboardCardSourceType) String() string {
    return []string{"bot", "unknownFutureValue"}[i]
}
func ParseTeamsAppDashboardCardSourceType(v string) (any, error) {
    result := BOT_TEAMSAPPDASHBOARDCARDSOURCETYPE
    switch v {
        case "bot":
            result = BOT_TEAMSAPPDASHBOARDCARDSOURCETYPE
        case "unknownFutureValue":
            result = UNKNOWNFUTUREVALUE_TEAMSAPPDASHBOARDCARDSOURCETYPE
        default:
            return 0, errors.New("Unknown TeamsAppDashboardCardSourceType value: " + v)
    }
    return &result, nil
}
func SerializeTeamsAppDashboardCardSourceType(values []TeamsAppDashboardCardSourceType) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i TeamsAppDashboardCardSourceType) isMultiValue() bool {
    return false
}
