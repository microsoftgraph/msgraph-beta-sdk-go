package models
type PlannerContainerType int

const (
    GROUP_PLANNERCONTAINERTYPE PlannerContainerType = iota
    UNKNOWNFUTUREVALUE_PLANNERCONTAINERTYPE
    ROSTER_PLANNERCONTAINERTYPE
    PROJECT_PLANNERCONTAINERTYPE
    DRIVEITEM_PLANNERCONTAINERTYPE
    USER_PLANNERCONTAINERTYPE
)

func (i PlannerContainerType) String() string {
    return []string{"group", "unknownFutureValue", "roster", "project", "driveItem", "user"}[i]
}
func ParsePlannerContainerType(v string) (any, error) {
    result := GROUP_PLANNERCONTAINERTYPE
    switch v {
        case "group":
            result = GROUP_PLANNERCONTAINERTYPE
        case "unknownFutureValue":
            result = UNKNOWNFUTUREVALUE_PLANNERCONTAINERTYPE
        case "roster":
            result = ROSTER_PLANNERCONTAINERTYPE
        case "project":
            result = PROJECT_PLANNERCONTAINERTYPE
        case "driveItem":
            result = DRIVEITEM_PLANNERCONTAINERTYPE
        case "user":
            result = USER_PLANNERCONTAINERTYPE
        default:
            return nil, nil
    }
    return &result, nil
}
func SerializePlannerContainerType(values []PlannerContainerType) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i PlannerContainerType) isMultiValue() bool {
    return false
}
