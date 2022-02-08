package graph
import (
    "strings"
    "errors"
)
// 
type CertificateStatus int

const (
    NOTPROVISIONED_CERTIFICATESTATUS CertificateStatus = iota
    PROVISIONED_CERTIFICATESTATUS
)

func (i CertificateStatus) String() string {
    return []string{"NOTPROVISIONED", "PROVISIONED"}[i]
}
func ParseCertificateStatus(v string) (interface{}, error) {
    result := NOTPROVISIONED_CERTIFICATESTATUS
    switch strings.ToUpper(v) {
        case "NOTPROVISIONED":
            result = NOTPROVISIONED_CERTIFICATESTATUS
        case "PROVISIONED":
            result = PROVISIONED_CERTIFICATESTATUS
        default:
            return 0, errors.New("Unknown CertificateStatus value: " + v)
    }
    return &result, nil
}
func SerializeCertificateStatus(values []CertificateStatus) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
