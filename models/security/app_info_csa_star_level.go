package security
type AppInfoCsaStarLevel int

const (
    SELFASSESSMENT_APPINFOCSASTARLEVEL AppInfoCsaStarLevel = iota
    CERTIFICATION_APPINFOCSASTARLEVEL
    ATTESTATION_APPINFOCSASTARLEVEL
    CSTARASSESSMENT_APPINFOCSASTARLEVEL
    CONTINUOUSMONITORING_APPINFOCSASTARLEVEL
    UNKNOWN_APPINFOCSASTARLEVEL
    UNKNOWNFUTUREVALUE_APPINFOCSASTARLEVEL
)

func (i AppInfoCsaStarLevel) String() string {
    return []string{"selfAssessment", "certification", "attestation", "cStarAssessment", "continuousMonitoring", "unknown", "unknownFutureValue"}[i]
}
func ParseAppInfoCsaStarLevel(v string) (any, error) {
    result := SELFASSESSMENT_APPINFOCSASTARLEVEL
    switch v {
        case "selfAssessment":
            result = SELFASSESSMENT_APPINFOCSASTARLEVEL
        case "certification":
            result = CERTIFICATION_APPINFOCSASTARLEVEL
        case "attestation":
            result = ATTESTATION_APPINFOCSASTARLEVEL
        case "cStarAssessment":
            result = CSTARASSESSMENT_APPINFOCSASTARLEVEL
        case "continuousMonitoring":
            result = CONTINUOUSMONITORING_APPINFOCSASTARLEVEL
        case "unknown":
            result = UNKNOWN_APPINFOCSASTARLEVEL
        case "unknownFutureValue":
            result = UNKNOWNFUTUREVALUE_APPINFOCSASTARLEVEL
        default:
            return nil, nil
    }
    return &result, nil
}
func SerializeAppInfoCsaStarLevel(values []AppInfoCsaStarLevel) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i AppInfoCsaStarLevel) isMultiValue() bool {
    return false
}
