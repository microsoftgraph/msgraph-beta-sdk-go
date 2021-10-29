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
    switch strings.ToUpper(v) {
        case "NOTPROVISIONED":
            return NOTPROVISIONED_CERTIFICATESTATUS, nil
        case "PROVISIONED":
            return PROVISIONED_CERTIFICATESTATUS, nil
    }
    return 0, errors.New("Unknown CertificateStatus value: " + v)
}
func SerializeCertificateStatus(values []CertificateStatus) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
