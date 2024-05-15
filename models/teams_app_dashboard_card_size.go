package models
type TeamsAppDashboardCardSize int

const (
    MEDIUM_TEAMSAPPDASHBOARDCARDSIZE TeamsAppDashboardCardSize = iota
    LARGE_TEAMSAPPDASHBOARDCARDSIZE
    UNKNOWNFUTUREVALUE_TEAMSAPPDASHBOARDCARDSIZE
)

func (i TeamsAppDashboardCardSize) String() string {
    return []string{"medium", "large", "unknownFutureValue"}[i]
}
func ParseTeamsAppDashboardCardSize(v string) (any, error) {
    result := MEDIUM_TEAMSAPPDASHBOARDCARDSIZE
    switch v {
        case "medium":
            result = MEDIUM_TEAMSAPPDASHBOARDCARDSIZE
        case "large":
            result = LARGE_TEAMSAPPDASHBOARDCARDSIZE
        case "unknownFutureValue":
            result = UNKNOWNFUTUREVALUE_TEAMSAPPDASHBOARDCARDSIZE
        default:
            return nil, nil
    }
    return &result, nil
}
func SerializeTeamsAppDashboardCardSize(values []TeamsAppDashboardCardSize) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i TeamsAppDashboardCardSize) isMultiValue() bool {
    return false
}
