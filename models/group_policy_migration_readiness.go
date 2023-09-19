package models
import (
    "errors"
)
// Indicates if the Group Policy Object file is covered and ready for Intune migration.
type GroupPolicyMigrationReadiness int

const (
    // No Intune coverage
    NONE_GROUPPOLICYMIGRATIONREADINESS GroupPolicyMigrationReadiness = iota
    // Partial Intune coverage
    PARTIAL_GROUPPOLICYMIGRATIONREADINESS
    // Complete Intune coverage
    COMPLETE_GROUPPOLICYMIGRATIONREADINESS
    // Error when analyzing coverage
    ERROR_GROUPPOLICYMIGRATIONREADINESS
    // No Group Policy settings in GPO
    NOTAPPLICABLE_GROUPPOLICYMIGRATIONREADINESS
)

func (i GroupPolicyMigrationReadiness) String() string {
    return []string{"none", "partial", "complete", "error", "notApplicable"}[i]
}
func ParseGroupPolicyMigrationReadiness(v string) (any, error) {
    result := NONE_GROUPPOLICYMIGRATIONREADINESS
    switch v {
        case "none":
            result = NONE_GROUPPOLICYMIGRATIONREADINESS
        case "partial":
            result = PARTIAL_GROUPPOLICYMIGRATIONREADINESS
        case "complete":
            result = COMPLETE_GROUPPOLICYMIGRATIONREADINESS
        case "error":
            result = ERROR_GROUPPOLICYMIGRATIONREADINESS
        case "notApplicable":
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
func (i GroupPolicyMigrationReadiness) isMultiValue() bool {
    return false
}
