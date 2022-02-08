package graph
import (
    "strings"
    "errors"
)
// 
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
    result := NONE_CERTIFICATEREVOCATIONSTATUS
    switch strings.ToUpper(v) {
        case "NONE":
            result = NONE_CERTIFICATEREVOCATIONSTATUS
        case "PENDING":
            result = PENDING_CERTIFICATEREVOCATIONSTATUS
        case "ISSUED":
            result = ISSUED_CERTIFICATEREVOCATIONSTATUS
        case "FAILED":
            result = FAILED_CERTIFICATEREVOCATIONSTATUS
        case "REVOKED":
            result = REVOKED_CERTIFICATEREVOCATIONSTATUS
        default:
            return 0, errors.New("Unknown CertificateRevocationStatus value: " + v)
    }
    return &result, nil
}
func SerializeCertificateRevocationStatus(values []CertificateRevocationStatus) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
