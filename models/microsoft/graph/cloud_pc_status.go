package graph
import (
    "strings"
    "errors"
)
// 
type CloudPcStatus int

const (
    NOTPROVISIONED_CLOUDPCSTATUS CloudPcStatus = iota
    PROVISIONING_CLOUDPCSTATUS
    PROVISIONED_CLOUDPCSTATUS
    INGRACEPERIOD_CLOUDPCSTATUS
    DEPROVISIONING_CLOUDPCSTATUS
    FAILED_CLOUDPCSTATUS
    PROVISIONEDWITHWARNINGS_CLOUDPCSTATUS
    RESIZING_CLOUDPCSTATUS
    RESTORING_CLOUDPCSTATUS
    PENDINGPROVISION_CLOUDPCSTATUS
    UNKNOWNFUTUREVALUE_CLOUDPCSTATUS
)

func (i CloudPcStatus) String() string {
    return []string{"NOTPROVISIONED", "PROVISIONING", "PROVISIONED", "INGRACEPERIOD", "DEPROVISIONING", "FAILED", "PROVISIONEDWITHWARNINGS", "RESIZING", "RESTORING", "PENDINGPROVISION", "UNKNOWNFUTUREVALUE"}[i]
}
func ParseCloudPcStatus(v string) (interface{}, error) {
    result := NOTPROVISIONED_CLOUDPCSTATUS
    switch strings.ToUpper(v) {
        case "NOTPROVISIONED":
            result = NOTPROVISIONED_CLOUDPCSTATUS
        case "PROVISIONING":
            result = PROVISIONING_CLOUDPCSTATUS
        case "PROVISIONED":
            result = PROVISIONED_CLOUDPCSTATUS
        case "INGRACEPERIOD":
            result = INGRACEPERIOD_CLOUDPCSTATUS
        case "DEPROVISIONING":
            result = DEPROVISIONING_CLOUDPCSTATUS
        case "FAILED":
            result = FAILED_CLOUDPCSTATUS
        case "PROVISIONEDWITHWARNINGS":
            result = PROVISIONEDWITHWARNINGS_CLOUDPCSTATUS
        case "RESIZING":
            result = RESIZING_CLOUDPCSTATUS
        case "RESTORING":
            result = RESTORING_CLOUDPCSTATUS
        case "PENDINGPROVISION":
            result = PENDINGPROVISION_CLOUDPCSTATUS
        case "UNKNOWNFUTUREVALUE":
            result = UNKNOWNFUTUREVALUE_CLOUDPCSTATUS
        default:
            return 0, errors.New("Unknown CloudPcStatus value: " + v)
    }
    return &result, nil
}
func SerializeCloudPcStatus(values []CloudPcStatus) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
