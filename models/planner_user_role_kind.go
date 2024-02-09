package models
import (
    "errors"
)
type PlannerUserRoleKind int

const (
    RELATIONSHIP_PLANNERUSERROLEKIND PlannerUserRoleKind = iota
    UNKNOWNFUTUREVALUE_PLANNERUSERROLEKIND
)

func (i PlannerUserRoleKind) String() string {
    return []string{"relationship", "unknownFutureValue"}[i]
}
func ParsePlannerUserRoleKind(v string) (any, error) {
    result := RELATIONSHIP_PLANNERUSERROLEKIND
    switch v {
        case "relationship":
            result = RELATIONSHIP_PLANNERUSERROLEKIND
        case "unknownFutureValue":
            result = UNKNOWNFUTUREVALUE_PLANNERUSERROLEKIND
        default:
            return 0, errors.New("Unknown PlannerUserRoleKind value: " + v)
    }
    return &result, nil
}
func SerializePlannerUserRoleKind(values []PlannerUserRoleKind) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i PlannerUserRoleKind) isMultiValue() bool {
    return false
}
