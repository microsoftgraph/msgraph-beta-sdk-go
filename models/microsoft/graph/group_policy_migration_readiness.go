package graph
import (
    "strings"
    "errors"
)
// 
type GroupPolicyMigrationReadiness int

const (
    NONE_GROUPPOLICYMIGRATIONREADINESS GroupPolicyMigrationReadiness = iota
    PARTIAL_GROUPPOLICYMIGRATIONREADINESS
    COMPLETE_GROUPPOLICYMIGRATIONREADINESS
    ERROR_GROUPPOLICYMIGRATIONREADINESS
    NOTAPPLICABLE_GROUPPOLICYMIGRATIONREADINESS
)

func (i GroupPolicyMigrationReadiness) String() string {
    return []string{"NONE", "PARTIAL", "COMPLETE", "ERROR", "NOTAPPLICABLE"}[i]
}
func ParseGroupPolicyMigrationReadiness(v string) (interface{}, error) {
    result := NONE_GROUPPOLICYMIGRATIONREADINESS
    switch strings.ToUpper(v) {
        case "NONE":
            result = NONE_GROUPPOLICYMIGRATIONREADINESS
        case "PARTIAL":
            result = PARTIAL_GROUPPOLICYMIGRATIONREADINESS
        case "COMPLETE":
            result = COMPLETE_GROUPPOLICYMIGRATIONREADINESS
        case "ERROR":
            result = ERROR_GROUPPOLICYMIGRATIONREADINESS
        case "NOTAPPLICABLE":
            result = NOTAPPLICABLE_GROUPPOLICYMIGRATIONREADINESS
        default:
            return 0, errors.New("Unknown GroupPolicyMigrationReadiness value: " + v)
    }
    return &result, nil
}
func SerializeGroupPolicyMigrationReadiness(values []GroupPolicyMigrationReadiness) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
