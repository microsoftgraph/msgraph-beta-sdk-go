package networkaccess
type HttpMethod int

const (
    GET_HTTPMETHOD HttpMethod = iota
    POST_HTTPMETHOD
    PUT_HTTPMETHOD
    DELETE_HTTPMETHOD
    HEAD_HTTPMETHOD
    OPTIONS_HTTPMETHOD
    CONNECT_HTTPMETHOD
    PATCH_HTTPMETHOD
    TRACE_HTTPMETHOD
    UNKNOWNFUTUREVALUE_HTTPMETHOD
)

func (i HttpMethod) String() string {
    return []string{"get", "post", "put", "delete", "head", "options", "connect", "patch", "trace", "unknownFutureValue"}[i]
}
func ParseHttpMethod(v string) (any, error) {
    result := GET_HTTPMETHOD
    switch v {
        case "get":
            result = GET_HTTPMETHOD
        case "post":
            result = POST_HTTPMETHOD
        case "put":
            result = PUT_HTTPMETHOD
        case "delete":
            result = DELETE_HTTPMETHOD
        case "head":
            result = HEAD_HTTPMETHOD
        case "options":
            result = OPTIONS_HTTPMETHOD
        case "connect":
            result = CONNECT_HTTPMETHOD
        case "patch":
            result = PATCH_HTTPMETHOD
        case "trace":
            result = TRACE_HTTPMETHOD
        case "unknownFutureValue":
            result = UNKNOWNFUTUREVALUE_HTTPMETHOD
        default:
            return nil, nil
    }
    return &result, nil
}
func SerializeHttpMethod(values []HttpMethod) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i HttpMethod) isMultiValue() bool {
    return false
}
