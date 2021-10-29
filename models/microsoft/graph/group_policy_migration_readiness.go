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
    switch strings.ToUpper(v) {
        case "NONE":
            return NONE_GROUPPOLICYMIGRATIONREADINESS, nil
        case "PARTIAL":
            return PARTIAL_GROUPPOLICYMIGRATIONREADINESS, nil
        case "COMPLETE":
            return COMPLETE_GROUPPOLICYMIGRATIONREADINESS, nil
        case "ERROR":
            return ERROR_GROUPPOLICYMIGRATIONREADINESS, nil
        case "NOTAPPLICABLE":
            return NOTAPPLICABLE_GROUPPOLICYMIGRATIONREADINESS, nil
    }
    return 0, errors.New("Unknown GroupPolicyMigrationReadiness value: " + v)
}
func SerializeGroupPolicyMigrationReadiness(values []GroupPolicyMigrationReadiness) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
