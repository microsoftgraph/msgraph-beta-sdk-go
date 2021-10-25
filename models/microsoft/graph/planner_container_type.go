package graph
import (
    "strings"
    "errors"
)
type PlannerContainerType int

const (
    GROUP_PLANNERCONTAINERTYPE PlannerContainerType = iota
    UNKNOWNFUTUREVALUE_PLANNERCONTAINERTYPE
    ROSTER_PLANNERCONTAINERTYPE
)

func (i PlannerContainerType) String() string {
    return []string{"GROUP", "UNKNOWNFUTUREVALUE", "ROSTER"}[i]
}
func ParsePlannerContainerType(v string) (interface{}, error) {
    switch strings.ToUpper(v) {
        case "GROUP":
            return GROUP_PLANNERCONTAINERTYPE, nil
        case "UNKNOWNFUTUREVALUE":
            return UNKNOWNFUTUREVALUE_PLANNERCONTAINERTYPE, nil
        case "ROSTER":
            return ROSTER_PLANNERCONTAINERTYPE, nil
    }
    return 0, errors.New("Unknown PlannerContainerType value: " + v)
}
func SerializePlannerContainerType(values []PlannerContainerType) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
