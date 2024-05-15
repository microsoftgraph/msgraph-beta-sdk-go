package models
import (
    "math"
    "strings"
)
type PrivateNetworkProtocol int

const (
    TCP_PRIVATENETWORKPROTOCOL = 1
    UDP_PRIVATENETWORKPROTOCOL = 2
    UNKNOWNFUTUREVALUE_PRIVATENETWORKPROTOCOL = 4
)

func (i PrivateNetworkProtocol) String() string {
    var values []string
    options := []string{"tcp", "udp", "unknownFutureValue"}
    for p := 0; p < 3; p++ {
        mantis := PrivateNetworkProtocol(int(math.Pow(2, float64(p))))
        if i&mantis == mantis {
            values = append(values, options[p])
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
                return nil, nil
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
