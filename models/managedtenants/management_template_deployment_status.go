package managedtenants
import (
    "strings"
    "errors"
)
// Provides operations to manage the tenantRelationship singleton.
type ManagementTemplateDeploymentStatus int

const (
    UNKNOWN_MANAGEMENTTEMPLATEDEPLOYMENTSTATUS ManagementTemplateDeploymentStatus = iota
    INPROGRESS_MANAGEMENTTEMPLATEDEPLOYMENTSTATUS
    COMPLETED_MANAGEMENTTEMPLATEDEPLOYMENTSTATUS
    FAILED_MANAGEMENTTEMPLATEDEPLOYMENTSTATUS
    INELIGIBLE_MANAGEMENTTEMPLATEDEPLOYMENTSTATUS
    UNKNOWNFUTUREVALUE_MANAGEMENTTEMPLATEDEPLOYMENTSTATUS
)

func (i ManagementTemplateDeploymentStatus) String() string {
    return []string{"UNKNOWN", "INPROGRESS", "COMPLETED", "FAILED", "INELIGIBLE", "UNKNOWNFUTUREVALUE"}[i]
}
func ParseManagementTemplateDeploymentStatus(v string) (interface{}, error) {
    result := UNKNOWN_MANAGEMENTTEMPLATEDEPLOYMENTSTATUS
    switch strings.ToUpper(v) {
        case "UNKNOWN":
            result = UNKNOWN_MANAGEMENTTEMPLATEDEPLOYMENTSTATUS
        case "INPROGRESS":
            result = INPROGRESS_MANAGEMENTTEMPLATEDEPLOYMENTSTATUS
        case "COMPLETED":
            result = COMPLETED_MANAGEMENTTEMPLATEDEPLOYMENTSTATUS
        case "FAILED":
            result = FAILED_MANAGEMENTTEMPLATEDEPLOYMENTSTATUS
        case "INELIGIBLE":
            result = INELIGIBLE_MANAGEMENTTEMPLATEDEPLOYMENTSTATUS
        case "UNKNOWNFUTUREVALUE":
            result = UNKNOWNFUTUREVALUE_MANAGEMENTTEMPLATEDEPLOYMENTSTATUS
        default:
            return 0, errors.New("Unknown ManagementTemplateDeploymentStatus value: " + v)
    }
    return &result, nil
}
func SerializeManagementTemplateDeploymentStatus(values []ManagementTemplateDeploymentStatus) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
