package models
// CertificateStore types
type CertificateStore int

const (
    USER_CERTIFICATESTORE CertificateStore = iota
    MACHINE_CERTIFICATESTORE
)

func (i CertificateStore) String() string {
    return []string{"user", "machine"}[i]
}
func ParseCertificateStore(v string) (any, error) {
    result := USER_CERTIFICATESTORE
    switch v {
        case "user":
            result = USER_CERTIFICATESTORE
        case "machine":
            result = MACHINE_CERTIFICATESTORE
        default:
            return nil, nil
    }
    return &result, nil
}
func SerializeCertificateStore(values []CertificateStore) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i CertificateStore) isMultiValue() bool {
    return false
}
