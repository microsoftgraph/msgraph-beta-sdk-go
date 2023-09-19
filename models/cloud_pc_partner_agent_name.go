package models
import (
    "errors"
)
// 
type CloudPcPartnerAgentName int

const (
    CITRIX_CLOUDPCPARTNERAGENTNAME CloudPcPartnerAgentName = iota
    UNKNOWNFUTUREVALUE_CLOUDPCPARTNERAGENTNAME
    VMWARE_CLOUDPCPARTNERAGENTNAME
)

func (i CloudPcPartnerAgentName) String() string {
    return []string{"citrix", "unknownFutureValue", "vMware"}[i]
}
func ParseCloudPcPartnerAgentName(v string) (any, error) {
    result := CITRIX_CLOUDPCPARTNERAGENTNAME
    switch v {
        case "citrix":
            result = CITRIX_CLOUDPCPARTNERAGENTNAME
        case "unknownFutureValue":
            result = UNKNOWNFUTUREVALUE_CLOUDPCPARTNERAGENTNAME
        case "vMware":
            result = VMWARE_CLOUDPCPARTNERAGENTNAME
        default:
            return 0, errors.New("Unknown CloudPcPartnerAgentName value: " + v)
    }
    return &result, nil
}
func SerializeCloudPcPartnerAgentName(values []CloudPcPartnerAgentName) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i CloudPcPartnerAgentName) isMultiValue() bool {
    return false
}
