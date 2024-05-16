package models
type CloudPcPartnerAgentName int

const (
    CITRIX_CLOUDPCPARTNERAGENTNAME CloudPcPartnerAgentName = iota
    UNKNOWNFUTUREVALUE_CLOUDPCPARTNERAGENTNAME
    VMWARE_CLOUDPCPARTNERAGENTNAME
    HP_CLOUDPCPARTNERAGENTNAME
)

func (i CloudPcPartnerAgentName) String() string {
    return []string{"citrix", "unknownFutureValue", "vMware", "hp"}[i]
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
        case "hp":
            result = HP_CLOUDPCPARTNERAGENTNAME
        default:
            return nil, nil
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
