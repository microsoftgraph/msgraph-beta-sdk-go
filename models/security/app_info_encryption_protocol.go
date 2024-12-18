package security
type AppInfoEncryptionProtocol int

const (
    TLS1_0_APPINFOENCRYPTIONPROTOCOL AppInfoEncryptionProtocol = iota
    TLS1_1_APPINFOENCRYPTIONPROTOCOL
    TLS1_2_APPINFOENCRYPTIONPROTOCOL
    TLS1_3_APPINFOENCRYPTIONPROTOCOL
    NOTAPPLICABLE_APPINFOENCRYPTIONPROTOCOL
    NOTSUPPORTED_APPINFOENCRYPTIONPROTOCOL
    UNKNOWN_APPINFOENCRYPTIONPROTOCOL
    UNKNOWNFUTUREVALUE_APPINFOENCRYPTIONPROTOCOL
    SSL3_APPINFOENCRYPTIONPROTOCOL
)

func (i AppInfoEncryptionProtocol) String() string {
    return []string{"tls1_0", "tls1_1", "tls1_2", "tls1_3", "notApplicable", "notSupported", "unknown", "unknownFutureValue", "ssl3"}[i]
}
func ParseAppInfoEncryptionProtocol(v string) (any, error) {
    result := TLS1_0_APPINFOENCRYPTIONPROTOCOL
    switch v {
        case "tls1_0":
            result = TLS1_0_APPINFOENCRYPTIONPROTOCOL
        case "tls1_1":
            result = TLS1_1_APPINFOENCRYPTIONPROTOCOL
        case "tls1_2":
            result = TLS1_2_APPINFOENCRYPTIONPROTOCOL
        case "tls1_3":
            result = TLS1_3_APPINFOENCRYPTIONPROTOCOL
        case "notApplicable":
            result = NOTAPPLICABLE_APPINFOENCRYPTIONPROTOCOL
        case "notSupported":
            result = NOTSUPPORTED_APPINFOENCRYPTIONPROTOCOL
        case "unknown":
            result = UNKNOWN_APPINFOENCRYPTIONPROTOCOL
        case "unknownFutureValue":
            result = UNKNOWNFUTUREVALUE_APPINFOENCRYPTIONPROTOCOL
        case "ssl3":
            result = SSL3_APPINFOENCRYPTIONPROTOCOL
        default:
            return nil, nil
    }
    return &result, nil
}
func SerializeAppInfoEncryptionProtocol(values []AppInfoEncryptionProtocol) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i AppInfoEncryptionProtocol) isMultiValue() bool {
    return false
}
