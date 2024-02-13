package models
import (
    "errors"
)
type X509CertificateIssuerHintsState int

const (
    DISABLED_X509CERTIFICATEISSUERHINTSSTATE X509CertificateIssuerHintsState = iota
    ENABLED_X509CERTIFICATEISSUERHINTSSTATE
    UNKNOWNFUTUREVALUE_X509CERTIFICATEISSUERHINTSSTATE
)

func (i X509CertificateIssuerHintsState) String() string {
    return []string{"disabled", "enabled", "unknownFutureValue"}[i]
}
func ParseX509CertificateIssuerHintsState(v string) (any, error) {
    result := DISABLED_X509CERTIFICATEISSUERHINTSSTATE
    switch v {
        case "disabled":
            result = DISABLED_X509CERTIFICATEISSUERHINTSSTATE
        case "enabled":
            result = ENABLED_X509CERTIFICATEISSUERHINTSSTATE
        case "unknownFutureValue":
            result = UNKNOWNFUTUREVALUE_X509CERTIFICATEISSUERHINTSSTATE
        default:
            return 0, errors.New("Unknown X509CertificateIssuerHintsState value: " + v)
    }
    return &result, nil
}
func SerializeX509CertificateIssuerHintsState(values []X509CertificateIssuerHintsState) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i X509CertificateIssuerHintsState) isMultiValue() bool {
    return false
}
