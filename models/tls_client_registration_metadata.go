package models
type TlsClientRegistrationMetadata int

const (
    TLS_CLIENT_AUTH_SUBJECT_DN_TLSCLIENTREGISTRATIONMETADATA TlsClientRegistrationMetadata = iota
    TLS_CLIENT_AUTH_SAN_DNS_TLSCLIENTREGISTRATIONMETADATA
    TLS_CLIENT_AUTH_SAN_URI_TLSCLIENTREGISTRATIONMETADATA
    TLS_CLIENT_AUTH_SAN_IP_TLSCLIENTREGISTRATIONMETADATA
    TLS_CLIENT_AUTH_SAN_EMAIL_TLSCLIENTREGISTRATIONMETADATA
    UNKNOWNFUTUREVALUE_TLSCLIENTREGISTRATIONMETADATA
)

func (i TlsClientRegistrationMetadata) String() string {
    return []string{"tls_client_auth_subject_dn", "tls_client_auth_san_dns", "tls_client_auth_san_uri", "tls_client_auth_san_ip", "tls_client_auth_san_email", "unknownFutureValue"}[i]
}
func ParseTlsClientRegistrationMetadata(v string) (any, error) {
    result := TLS_CLIENT_AUTH_SUBJECT_DN_TLSCLIENTREGISTRATIONMETADATA
    switch v {
        case "tls_client_auth_subject_dn":
            result = TLS_CLIENT_AUTH_SUBJECT_DN_TLSCLIENTREGISTRATIONMETADATA
        case "tls_client_auth_san_dns":
            result = TLS_CLIENT_AUTH_SAN_DNS_TLSCLIENTREGISTRATIONMETADATA
        case "tls_client_auth_san_uri":
            result = TLS_CLIENT_AUTH_SAN_URI_TLSCLIENTREGISTRATIONMETADATA
        case "tls_client_auth_san_ip":
            result = TLS_CLIENT_AUTH_SAN_IP_TLSCLIENTREGISTRATIONMETADATA
        case "tls_client_auth_san_email":
            result = TLS_CLIENT_AUTH_SAN_EMAIL_TLSCLIENTREGISTRATIONMETADATA
        case "unknownFutureValue":
            result = UNKNOWNFUTUREVALUE_TLSCLIENTREGISTRATIONMETADATA
        default:
            return nil, nil
    }
    return &result, nil
}
func SerializeTlsClientRegistrationMetadata(values []TlsClientRegistrationMetadata) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i TlsClientRegistrationMetadata) isMultiValue() bool {
    return false
}
