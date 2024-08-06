package networkaccess
type DeviceVendor int

const (
    BARRACUDANETWORKS_DEVICEVENDOR DeviceVendor = iota
    CHECKPOINT_DEVICEVENDOR
    CISCOMERAKI_DEVICEVENDOR
    CITRIX_DEVICEVENDOR
    FORTINET_DEVICEVENDOR
    HPEARUBA_DEVICEVENDOR
    NETFOUNDRY_DEVICEVENDOR
    NUAGE_DEVICEVENDOR
    OPENSYSTEMS_DEVICEVENDOR
    PALOALTONETWORKS_DEVICEVENDOR
    RIVERBEDTECHNOLOGY_DEVICEVENDOR
    SILVERPEAK_DEVICEVENDOR
    VMWARESDWAN_DEVICEVENDOR
    VERSA_DEVICEVENDOR
    OTHER_DEVICEVENDOR
    CISCOCATALYST_DEVICEVENDOR
    UNKNOWNFUTUREVALUE_DEVICEVENDOR
)

func (i DeviceVendor) String() string {
    return []string{"barracudaNetworks", "checkPoint", "ciscoMeraki", "citrix", "fortinet", "hpeAruba", "netFoundry", "nuage", "openSystems", "paloAltoNetworks", "riverbedTechnology", "silverPeak", "vmWareSdWan", "versa", "other", "ciscoCatalyst", "unknownFutureValue"}[i]
}
func ParseDeviceVendor(v string) (any, error) {
    result := BARRACUDANETWORKS_DEVICEVENDOR
    switch v {
        case "barracudaNetworks":
            result = BARRACUDANETWORKS_DEVICEVENDOR
        case "checkPoint":
            result = CHECKPOINT_DEVICEVENDOR
        case "ciscoMeraki":
            result = CISCOMERAKI_DEVICEVENDOR
        case "citrix":
            result = CITRIX_DEVICEVENDOR
        case "fortinet":
            result = FORTINET_DEVICEVENDOR
        case "hpeAruba":
            result = HPEARUBA_DEVICEVENDOR
        case "netFoundry":
            result = NETFOUNDRY_DEVICEVENDOR
        case "nuage":
            result = NUAGE_DEVICEVENDOR
        case "openSystems":
            result = OPENSYSTEMS_DEVICEVENDOR
        case "paloAltoNetworks":
            result = PALOALTONETWORKS_DEVICEVENDOR
        case "riverbedTechnology":
            result = RIVERBEDTECHNOLOGY_DEVICEVENDOR
        case "silverPeak":
            result = SILVERPEAK_DEVICEVENDOR
        case "vmWareSdWan":
            result = VMWARESDWAN_DEVICEVENDOR
        case "versa":
            result = VERSA_DEVICEVENDOR
        case "other":
            result = OTHER_DEVICEVENDOR
        case "ciscoCatalyst":
            result = CISCOCATALYST_DEVICEVENDOR
        case "unknownFutureValue":
            result = UNKNOWNFUTUREVALUE_DEVICEVENDOR
        default:
            return nil, nil
    }
    return &result, nil
}
func SerializeDeviceVendor(values []DeviceVendor) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i DeviceVendor) isMultiValue() bool {
    return false
}
