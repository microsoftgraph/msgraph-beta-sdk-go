package models
// Certificate Validity Period Options.
type CertificateValidityPeriodScale int

const (
    // Days.
    DAYS_CERTIFICATEVALIDITYPERIODSCALE CertificateValidityPeriodScale = iota
    // Months.
    MONTHS_CERTIFICATEVALIDITYPERIODSCALE
    // Years.
    YEARS_CERTIFICATEVALIDITYPERIODSCALE
)

func (i CertificateValidityPeriodScale) String() string {
    return []string{"days", "months", "years"}[i]
}
func ParseCertificateValidityPeriodScale(v string) (any, error) {
    result := DAYS_CERTIFICATEVALIDITYPERIODSCALE
    switch v {
        case "days":
            result = DAYS_CERTIFICATEVALIDITYPERIODSCALE
        case "months":
            result = MONTHS_CERTIFICATEVALIDITYPERIODSCALE
        case "years":
            result = YEARS_CERTIFICATEVALIDITYPERIODSCALE
        default:
            return nil, nil
    }
    return &result, nil
}
func SerializeCertificateValidityPeriodScale(values []CertificateValidityPeriodScale) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i CertificateValidityPeriodScale) isMultiValue() bool {
    return false
}
