package models
import (
    "errors"
)
// 
type PlannerPlanAccessLevel int

const (
    READACCESS_PLANNERPLANACCESSLEVEL PlannerPlanAccessLevel = iota
    READWRITEACCESS_PLANNERPLANACCESSLEVEL
    FULLACCESS_PLANNERPLANACCESSLEVEL
    UNKNOWNFUTUREVALUE_PLANNERPLANACCESSLEVEL
)

func (i PlannerPlanAccessLevel) String() string {
    return []string{"readAccess", "readWriteAccess", "fullAccess", "unknownFutureValue"}[i]
}
func ParsePlannerPlanAccessLevel(v string) (any, error) {
    result := READACCESS_PLANNERPLANACCESSLEVEL
    switch v {
        case "readAccess":
            result = READACCESS_PLANNERPLANACCESSLEVEL
        case "readWriteAccess":
            result = READWRITEACCESS_PLANNERPLANACCESSLEVEL
        case "fullAccess":
            result = FULLACCESS_PLANNERPLANACCESSLEVEL
        case "unknownFutureValue":
            result = UNKNOWNFUTUREVALUE_PLANNERPLANACCESSLEVEL
        default:
            return 0, errors.New("Unknown PlannerPlanAccessLevel value: " + v)
    }
    return &result, nil
}
func SerializePlannerPlanAccessLevel(values []PlannerPlanAccessLevel) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i PlannerPlanAccessLevel) isMultiValue() bool {
    return false
}
