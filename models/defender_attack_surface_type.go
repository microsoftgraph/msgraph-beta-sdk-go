package models
// Possible values of Defender Attack Surface Reduction Rules
type DefenderAttackSurfaceType int

const (
    // Default, which disables attack surface reduction rule.
    USERDEFINED_DEFENDERATTACKSURFACETYPE DefenderAttackSurfaceType = iota
    // Enable the attack surface reduction rule.
    BLOCK_DEFENDERATTACKSURFACETYPE
    // Evaluate how the ASR rule would impact your organization if enabled. Does not change functionality but generate logs.
    AUDITMODE_DEFENDERATTACKSURFACETYPE
    // Warning message to end user with ability to bypass block from attack surface reduction rule.
    WARN_DEFENDERATTACKSURFACETYPE
    // Disable the attack surface reduction rule
    DISABLE_DEFENDERATTACKSURFACETYPE
)

func (i DefenderAttackSurfaceType) String() string {
    return []string{"userDefined", "block", "auditMode", "warn", "disable"}[i]
}
func ParseDefenderAttackSurfaceType(v string) (any, error) {
    result := USERDEFINED_DEFENDERATTACKSURFACETYPE
    switch v {
        case "userDefined":
            result = USERDEFINED_DEFENDERATTACKSURFACETYPE
        case "block":
            result = BLOCK_DEFENDERATTACKSURFACETYPE
        case "auditMode":
            result = AUDITMODE_DEFENDERATTACKSURFACETYPE
        case "warn":
            result = WARN_DEFENDERATTACKSURFACETYPE
        case "disable":
            result = DISABLE_DEFENDERATTACKSURFACETYPE
        default:
            return nil, nil
    }
    return &result, nil
}
func SerializeDefenderAttackSurfaceType(values []DefenderAttackSurfaceType) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i DefenderAttackSurfaceType) isMultiValue() bool {
    return false
}
