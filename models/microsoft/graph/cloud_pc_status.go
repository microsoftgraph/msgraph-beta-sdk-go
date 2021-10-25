package graph
import (
    "strings"
    "errors"
)
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
    UNKNOWNFUTUREVALUE_CLOUDPCSTATUS
)

func (i CloudPcStatus) String() string {
    return []string{"NOTPROVISIONED", "PROVISIONING", "PROVISIONED", "INGRACEPERIOD", "DEPROVISIONING", "FAILED", "PROVISIONEDWITHWARNINGS", "RESIZING", "UNKNOWNFUTUREVALUE"}[i]
}
func ParseCloudPcStatus(v string) (interface{}, error) {
    switch strings.ToUpper(v) {
        case "NOTPROVISIONED":
            return NOTPROVISIONED_CLOUDPCSTATUS, nil
        case "PROVISIONING":
            return PROVISIONING_CLOUDPCSTATUS, nil
        case "PROVISIONED":
            return PROVISIONED_CLOUDPCSTATUS, nil
        case "INGRACEPERIOD":
            return INGRACEPERIOD_CLOUDPCSTATUS, nil
        case "DEPROVISIONING":
            return DEPROVISIONING_CLOUDPCSTATUS, nil
        case "FAILED":
            return FAILED_CLOUDPCSTATUS, nil
        case "PROVISIONEDWITHWARNINGS":
            return PROVISIONEDWITHWARNINGS_CLOUDPCSTATUS, nil
        case "RESIZING":
            return RESIZING_CLOUDPCSTATUS, nil
        case "UNKNOWNFUTUREVALUE":
            return UNKNOWNFUTUREVALUE_CLOUDPCSTATUS, nil
    }
    return 0, errors.New("Unknown CloudPcStatus value: " + v)
}
func SerializeCloudPcStatus(values []CloudPcStatus) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
