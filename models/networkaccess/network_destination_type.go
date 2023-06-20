package networkaccess
import (
    "errors"
)
// 
type NetworkDestinationType int

const (
    URL_NETWORKDESTINATIONTYPE NetworkDestinationType = iota
    FQDN_NETWORKDESTINATIONTYPE
    IPADDRESS_NETWORKDESTINATIONTYPE
    IPRANGE_NETWORKDESTINATIONTYPE
    IPSUBNET_NETWORKDESTINATIONTYPE
    WEBCATEGORY_NETWORKDESTINATIONTYPE
    UNKNOWNFUTUREVALUE_NETWORKDESTINATIONTYPE
)

func (i NetworkDestinationType) String() string {
    return []string{"url", "fqdn", "ipAddress", "ipRange", "ipSubnet", "webCategory", "unknownFutureValue"}[i]
}
func ParseNetworkDestinationType(v string) (any, error) {
    result := URL_NETWORKDESTINATIONTYPE
    switch v {
        case "url":
            result = URL_NETWORKDESTINATIONTYPE
        case "fqdn":
            result = FQDN_NETWORKDESTINATIONTYPE
        case "ipAddress":
            result = IPADDRESS_NETWORKDESTINATIONTYPE
        case "ipRange":
            result = IPRANGE_NETWORKDESTINATIONTYPE
        case "ipSubnet":
            result = IPSUBNET_NETWORKDESTINATIONTYPE
        case "webCategory":
            result = WEBCATEGORY_NETWORKDESTINATIONTYPE
        case "unknownFutureValue":
            result = UNKNOWNFUTUREVALUE_NETWORKDESTINATIONTYPE
        default:
            return 0, errors.New("Unknown NetworkDestinationType value: " + v)
    }
    return &result, nil
}
func SerializeNetworkDestinationType(values []NetworkDestinationType) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
