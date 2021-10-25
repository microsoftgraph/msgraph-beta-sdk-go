package graph
import (
    "strings"
    "errors"
)
type CertificateRevocationStatus int

const (
    NONE_CERTIFICATEREVOCATIONSTATUS CertificateRevocationStatus = iota
    PENDING_CERTIFICATEREVOCATIONSTATUS
    ISSUED_CERTIFICATEREVOCATIONSTATUS
    FAILED_CERTIFICATEREVOCATIONSTATUS
    REVOKED_CERTIFICATEREVOCATIONSTATUS
)

func (i CertificateRevocationStatus) String() string {
    return []string{"NONE", "PENDING", "ISSUED", "FAILED", "REVOKED"}[i]
}
func ParseCertificateRevocationStatus(v string) (interface{}, error) {
    switch strings.ToUpper(v) {
        case "NONE":
            return NONE_CERTIFICATEREVOCATIONSTATUS, nil
        case "PENDING":
            return PENDING_CERTIFICATEREVOCATIONSTATUS, nil
        case "ISSUED":
            return ISSUED_CERTIFICATEREVOCATIONSTATUS, nil
        case "FAILED":
            return FAILED_CERTIFICATEREVOCATIONSTATUS, nil
        case "REVOKED":
            return REVOKED_CERTIFICATEREVOCATIONSTATUS, nil
    }
    return 0, errors.New("Unknown CertificateRevocationStatus value: " + v)
}
func SerializeCertificateRevocationStatus(values []CertificateRevocationStatus) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
