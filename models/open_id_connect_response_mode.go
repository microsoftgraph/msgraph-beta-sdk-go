package models
type OpenIdConnectResponseMode int

const (
    FORM_POST_OPENIDCONNECTRESPONSEMODE OpenIdConnectResponseMode = iota
    QUERY_OPENIDCONNECTRESPONSEMODE
    UNKNOWNFUTUREVALUE_OPENIDCONNECTRESPONSEMODE
)

func (i OpenIdConnectResponseMode) String() string {
    return []string{"form_post", "query", "unknownFutureValue"}[i]
}
func ParseOpenIdConnectResponseMode(v string) (any, error) {
    result := FORM_POST_OPENIDCONNECTRESPONSEMODE
    switch v {
        case "form_post":
            result = FORM_POST_OPENIDCONNECTRESPONSEMODE
        case "query":
            result = QUERY_OPENIDCONNECTRESPONSEMODE
        case "unknownFutureValue":
            result = UNKNOWNFUTUREVALUE_OPENIDCONNECTRESPONSEMODE
        default:
            return nil, nil
    }
    return &result, nil
}
func SerializeOpenIdConnectResponseMode(values []OpenIdConnectResponseMode) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i OpenIdConnectResponseMode) isMultiValue() bool {
    return false
}
