package models
type CloudPcBlobAccessTier int

const (
    HOT_CLOUDPCBLOBACCESSTIER CloudPcBlobAccessTier = iota
    COOL_CLOUDPCBLOBACCESSTIER
    COLD_CLOUDPCBLOBACCESSTIER
    ARCHIVE_CLOUDPCBLOBACCESSTIER
    UNKNOWNFUTUREVALUE_CLOUDPCBLOBACCESSTIER
)

func (i CloudPcBlobAccessTier) String() string {
    return []string{"hot", "cool", "cold", "archive", "unknownFutureValue"}[i]
}
func ParseCloudPcBlobAccessTier(v string) (any, error) {
    result := HOT_CLOUDPCBLOBACCESSTIER
    switch v {
        case "hot":
            result = HOT_CLOUDPCBLOBACCESSTIER
        case "cool":
            result = COOL_CLOUDPCBLOBACCESSTIER
        case "cold":
            result = COLD_CLOUDPCBLOBACCESSTIER
        case "archive":
            result = ARCHIVE_CLOUDPCBLOBACCESSTIER
        case "unknownFutureValue":
            result = UNKNOWNFUTUREVALUE_CLOUDPCBLOBACCESSTIER
        default:
            return nil, nil
    }
    return &result, nil
}
func SerializeCloudPcBlobAccessTier(values []CloudPcBlobAccessTier) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i CloudPcBlobAccessTier) isMultiValue() bool {
    return false
}
