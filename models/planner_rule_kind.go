package models
type PlannerRuleKind int

const (
    TASKRULE_PLANNERRULEKIND PlannerRuleKind = iota
    BUCKETRULE_PLANNERRULEKIND
    PLANRULE_PLANNERRULEKIND
    UNKNOWNFUTUREVALUE_PLANNERRULEKIND
)

func (i PlannerRuleKind) String() string {
    return []string{"taskRule", "bucketRule", "planRule", "unknownFutureValue"}[i]
}
func ParsePlannerRuleKind(v string) (any, error) {
    result := TASKRULE_PLANNERRULEKIND
    switch v {
        case "taskRule":
            result = TASKRULE_PLANNERRULEKIND
        case "bucketRule":
            result = BUCKETRULE_PLANNERRULEKIND
        case "planRule":
            result = PLANRULE_PLANNERRULEKIND
        case "unknownFutureValue":
            result = UNKNOWNFUTUREVALUE_PLANNERRULEKIND
        default:
            return nil, nil
    }
    return &result, nil
}
func SerializePlannerRuleKind(values []PlannerRuleKind) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i PlannerRuleKind) isMultiValue() bool {
    return false
}
