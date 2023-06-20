package networkaccess
import (
    "errors"
)
// 
type NetworkingProtocol int

const (
    IP_NETWORKINGPROTOCOL NetworkingProtocol = iota
    ICMP_NETWORKINGPROTOCOL
    IGMP_NETWORKINGPROTOCOL
    GGP_NETWORKINGPROTOCOL
    IPV4_NETWORKINGPROTOCOL
    TCP_NETWORKINGPROTOCOL
    PUP_NETWORKINGPROTOCOL
    UDP_NETWORKINGPROTOCOL
    IDP_NETWORKINGPROTOCOL
    IPV6_NETWORKINGPROTOCOL
    IPV6ROUTINGHEADER_NETWORKINGPROTOCOL
    IPV6FRAGMENTHEADER_NETWORKINGPROTOCOL
    IPSECENCAPSULATINGSECURITYPAYLOAD_NETWORKINGPROTOCOL
    IPSECAUTHENTICATIONHEADER_NETWORKINGPROTOCOL
    ICMPV6_NETWORKINGPROTOCOL
    IPV6NONEXTHEADER_NETWORKINGPROTOCOL
    IPV6DESTINATIONOPTIONS_NETWORKINGPROTOCOL
    ND_NETWORKINGPROTOCOL
    IPX_NETWORKINGPROTOCOL
    RAW_NETWORKINGPROTOCOL
    SPX_NETWORKINGPROTOCOL
    SPXII_NETWORKINGPROTOCOL
    UNKNOWNFUTUREVALUE_NETWORKINGPROTOCOL
)

func (i NetworkingProtocol) String() string {
    return []string{"ip", "icmp", "igmp", "ggp", "ipv4", "tcp", "pup", "udp", "idp", "ipv6", "ipv6RoutingHeader", "ipv6FragmentHeader", "ipSecEncapsulatingSecurityPayload", "ipSecAuthenticationHeader", "icmpV6", "ipv6NoNextHeader", "ipv6DestinationOptions", "nd", "ipx", "raw", "spx", "spxII", "unknownFutureValue"}[i]
}
func ParseNetworkingProtocol(v string) (any, error) {
    result := IP_NETWORKINGPROTOCOL
    switch v {
        case "ip":
            result = IP_NETWORKINGPROTOCOL
        case "icmp":
            result = ICMP_NETWORKINGPROTOCOL
        case "igmp":
            result = IGMP_NETWORKINGPROTOCOL
        case "ggp":
            result = GGP_NETWORKINGPROTOCOL
        case "ipv4":
            result = IPV4_NETWORKINGPROTOCOL
        case "tcp":
            result = TCP_NETWORKINGPROTOCOL
        case "pup":
            result = PUP_NETWORKINGPROTOCOL
        case "udp":
            result = UDP_NETWORKINGPROTOCOL
        case "idp":
            result = IDP_NETWORKINGPROTOCOL
        case "ipv6":
            result = IPV6_NETWORKINGPROTOCOL
        case "ipv6RoutingHeader":
            result = IPV6ROUTINGHEADER_NETWORKINGPROTOCOL
        case "ipv6FragmentHeader":
            result = IPV6FRAGMENTHEADER_NETWORKINGPROTOCOL
        case "ipSecEncapsulatingSecurityPayload":
            result = IPSECENCAPSULATINGSECURITYPAYLOAD_NETWORKINGPROTOCOL
        case "ipSecAuthenticationHeader":
            result = IPSECAUTHENTICATIONHEADER_NETWORKINGPROTOCOL
        case "icmpV6":
            result = ICMPV6_NETWORKINGPROTOCOL
        case "ipv6NoNextHeader":
            result = IPV6NONEXTHEADER_NETWORKINGPROTOCOL
        case "ipv6DestinationOptions":
            result = IPV6DESTINATIONOPTIONS_NETWORKINGPROTOCOL
        case "nd":
            result = ND_NETWORKINGPROTOCOL
        case "ipx":
            result = IPX_NETWORKINGPROTOCOL
        case "raw":
            result = RAW_NETWORKINGPROTOCOL
        case "spx":
            result = SPX_NETWORKINGPROTOCOL
        case "spxII":
            result = SPXII_NETWORKINGPROTOCOL
        case "unknownFutureValue":
            result = UNKNOWNFUTUREVALUE_NETWORKINGPROTOCOL
        default:
            return 0, errors.New("Unknown NetworkingProtocol value: " + v)
    }
    return &result, nil
}
func SerializeNetworkingProtocol(values []NetworkingProtocol) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
