package security
type AppInfoPciDssVersion int

const (
    V1_APPINFOPCIDSSVERSION AppInfoPciDssVersion = iota
    V2_APPINFOPCIDSSVERSION
    V3_APPINFOPCIDSSVERSION
    V3_1_APPINFOPCIDSSVERSION
    V3_2_APPINFOPCIDSSVERSION
    V3_2_1_APPINFOPCIDSSVERSION
    NOTSUPPORTED_APPINFOPCIDSSVERSION
    UNKNOWN_APPINFOPCIDSSVERSION
    UNKNOWNFUTUREVALUE_APPINFOPCIDSSVERSION
    V4_APPINFOPCIDSSVERSION
)

func (i AppInfoPciDssVersion) String() string {
    return []string{"v1", "v2", "v3", "v3_1", "v3_2", "v3_2_1", "notSupported", "unknown", "unknownFutureValue", "v4"}[i]
}
func ParseAppInfoPciDssVersion(v string) (any, error) {
    result := V1_APPINFOPCIDSSVERSION
    switch v {
        case "v1":
            result = V1_APPINFOPCIDSSVERSION
        case "v2":
            result = V2_APPINFOPCIDSSVERSION
        case "v3":
            result = V3_APPINFOPCIDSSVERSION
        case "v3_1":
            result = V3_1_APPINFOPCIDSSVERSION
        case "v3_2":
            result = V3_2_APPINFOPCIDSSVERSION
        case "v3_2_1":
            result = V3_2_1_APPINFOPCIDSSVERSION
        case "notSupported":
            result = NOTSUPPORTED_APPINFOPCIDSSVERSION
        case "unknown":
            result = UNKNOWN_APPINFOPCIDSSVERSION
        case "unknownFutureValue":
            result = UNKNOWNFUTUREVALUE_APPINFOPCIDSSVERSION
        case "v4":
            result = V4_APPINFOPCIDSSVERSION
        default:
            return nil, nil
    }
    return &result, nil
}
func SerializeAppInfoPciDssVersion(values []AppInfoPciDssVersion) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i AppInfoPciDssVersion) isMultiValue() bool {
    return false
}
