// Code generated by Microsoft Kiota - DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package models
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
            return nil, nil
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
