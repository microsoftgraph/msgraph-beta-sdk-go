package security
type ReceiverProtocol int

const (
    FTP_RECEIVERPROTOCOL ReceiverProtocol = iota
    FTPS_RECEIVERPROTOCOL
    SYSLOGUDP_RECEIVERPROTOCOL
    SYSLOGTCP_RECEIVERPROTOCOL
    SYSLOGTLS_RECEIVERPROTOCOL
    UNKNOWNFUTUREVALUE_RECEIVERPROTOCOL
)

func (i ReceiverProtocol) String() string {
    return []string{"ftp", "ftps", "syslogUdp", "syslogTcp", "syslogTls", "unknownFutureValue"}[i]
}
func ParseReceiverProtocol(v string) (any, error) {
    result := FTP_RECEIVERPROTOCOL
    switch v {
        case "ftp":
            result = FTP_RECEIVERPROTOCOL
        case "ftps":
            result = FTPS_RECEIVERPROTOCOL
        case "syslogUdp":
            result = SYSLOGUDP_RECEIVERPROTOCOL
        case "syslogTcp":
            result = SYSLOGTCP_RECEIVERPROTOCOL
        case "syslogTls":
            result = SYSLOGTLS_RECEIVERPROTOCOL
        case "unknownFutureValue":
            result = UNKNOWNFUTUREVALUE_RECEIVERPROTOCOL
        default:
            return nil, nil
    }
    return &result, nil
}
func SerializeReceiverProtocol(values []ReceiverProtocol) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i ReceiverProtocol) isMultiValue() bool {
    return false
}
