package models
type UriUsageType int

const (
    REDIRECTURI_URIUSAGETYPE UriUsageType = iota
    IDENTIFIERURI_URIUSAGETYPE
    LOGINURL_URIUSAGETYPE
    LOGOUTURL_URIUSAGETYPE
    UNKNOWNFUTUREVALUE_URIUSAGETYPE
)

func (i UriUsageType) String() string {
    return []string{"redirectUri", "identifierUri", "loginUrl", "logoutUrl", "unknownFutureValue"}[i]
}
func ParseUriUsageType(v string) (any, error) {
    result := REDIRECTURI_URIUSAGETYPE
    switch v {
        case "redirectUri":
            result = REDIRECTURI_URIUSAGETYPE
        case "identifierUri":
            result = IDENTIFIERURI_URIUSAGETYPE
        case "loginUrl":
            result = LOGINURL_URIUSAGETYPE
        case "logoutUrl":
            result = LOGOUTURL_URIUSAGETYPE
        case "unknownFutureValue":
            result = UNKNOWNFUTUREVALUE_URIUSAGETYPE
        default:
            return nil, nil
    }
    return &result, nil
}
func SerializeUriUsageType(values []UriUsageType) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i UriUsageType) isMultiValue() bool {
    return false
}
