package models
type PrivateNetworkDestinationType int

const (
    IPADDRESS_PRIVATENETWORKDESTINATIONTYPE PrivateNetworkDestinationType = iota
    IPRANGE_PRIVATENETWORKDESTINATIONTYPE
    IPRANGECIDR_PRIVATENETWORKDESTINATIONTYPE
    FQDN_PRIVATENETWORKDESTINATIONTYPE
    DNSSUFFIX_PRIVATENETWORKDESTINATIONTYPE
    UNKNOWNFUTUREVALUE_PRIVATENETWORKDESTINATIONTYPE
)

func (i PrivateNetworkDestinationType) String() string {
    return []string{"ipAddress", "ipRange", "ipRangeCidr", "fqdn", "dnsSuffix", "unknownFutureValue"}[i]
}
func ParsePrivateNetworkDestinationType(v string) (any, error) {
    result := IPADDRESS_PRIVATENETWORKDESTINATIONTYPE
    switch v {
        case "ipAddress":
            result = IPADDRESS_PRIVATENETWORKDESTINATIONTYPE
        case "ipRange":
            result = IPRANGE_PRIVATENETWORKDESTINATIONTYPE
        case "ipRangeCidr":
            result = IPRANGECIDR_PRIVATENETWORKDESTINATIONTYPE
        case "fqdn":
            result = FQDN_PRIVATENETWORKDESTINATIONTYPE
        case "dnsSuffix":
            result = DNSSUFFIX_PRIVATENETWORKDESTINATIONTYPE
        case "unknownFutureValue":
            result = UNKNOWNFUTUREVALUE_PRIVATENETWORKDESTINATIONTYPE
        default:
            return nil, nil
    }
    return &result, nil
}
func SerializePrivateNetworkDestinationType(values []PrivateNetworkDestinationType) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i PrivateNetworkDestinationType) isMultiValue() bool {
    return false
}
