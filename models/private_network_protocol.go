package models
import (
    "errors"
    "strings"
)
// 
type PrivateNetworkProtocol int

const (
    TCP_PRIVATENETWORKPROTOCOL PrivateNetworkProtocol = iota
    UDP_PRIVATENETWORKPROTOCOL
    UNKNOWNFUTUREVALUE_PRIVATENETWORKPROTOCOL
)

func (i PrivateNetworkProtocol) String() string {
    var values []string
    for p := PrivateNetworkProtocol(1); p <= UNKNOWNFUTUREVALUE_PRIVATENETWORKPROTOCOL; p <<= 1 {
        if i&p == p {
            values = append(values, []string{"tcp", "udp", "unknownFutureValue"}[p])
        }
    }
    return strings.Join(values, ",")
}
func ParsePrivateNetworkProtocol(v string) (any, error) {
    var result PrivateNetworkProtocol
    values := strings.Split(v, ",")
    for _, str := range values {
        switch str {
            case "tcp":
                result |= TCP_PRIVATENETWORKPROTOCOL
            case "udp":
                result |= UDP_PRIVATENETWORKPROTOCOL
            case "unknownFutureValue":
                result |= UNKNOWNFUTUREVALUE_PRIVATENETWORKPROTOCOL
            default:
                return 0, errors.New("Unknown PrivateNetworkProtocol value: " + v)
        }
    }
    return &result, nil
}
func SerializePrivateNetworkProtocol(values []PrivateNetworkProtocol) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i PrivateNetworkProtocol) isMultiValue() bool {
    return true
}
