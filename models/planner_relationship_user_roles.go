package models
import (
    "errors"
)
// 
type PlannerRelationshipUserRoles int

const (
    DEFAULTRULES_PLANNERRELATIONSHIPUSERROLES PlannerRelationshipUserRoles = iota
    GROUPOWNERS_PLANNERRELATIONSHIPUSERROLES
    GROUPMEMBERS_PLANNERRELATIONSHIPUSERROLES
    TASKASSIGNEES_PLANNERRELATIONSHIPUSERROLES
    APPLICATIONS_PLANNERRELATIONSHIPUSERROLES
    UNKNOWNFUTUREVALUE_PLANNERRELATIONSHIPUSERROLES
)

func (i PlannerRelationshipUserRoles) String() string {
    return []string{"defaultRules", "groupOwners", "groupMembers", "taskAssignees", "applications", "unknownFutureValue"}[i]
}
func ParsePlannerRelationshipUserRoles(v string) (any, error) {
    result := DEFAULTRULES_PLANNERRELATIONSHIPUSERROLES
    switch v {
        case "defaultRules":
            result = DEFAULTRULES_PLANNERRELATIONSHIPUSERROLES
        case "groupOwners":
            result = GROUPOWNERS_PLANNERRELATIONSHIPUSERROLES
        case "groupMembers":
            result = GROUPMEMBERS_PLANNERRELATIONSHIPUSERROLES
        case "taskAssignees":
            result = TASKASSIGNEES_PLANNERRELATIONSHIPUSERROLES
        case "applications":
            result = APPLICATIONS_PLANNERRELATIONSHIPUSERROLES
        case "unknownFutureValue":
            result = UNKNOWNFUTUREVALUE_PLANNERRELATIONSHIPUSERROLES
        default:
            return 0, errors.New("Unknown PlannerRelationshipUserRoles value: " + v)
    }
    return &result, nil
}
func SerializePlannerRelationshipUserRoles(values []PlannerRelationshipUserRoles) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i PlannerRelationshipUserRoles) isMultiValue() bool {
    return false
}
