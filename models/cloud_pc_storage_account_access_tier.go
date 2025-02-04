package models
type CloudPcStorageAccountAccessTier int

const (
    HOT_CLOUDPCSTORAGEACCOUNTACCESSTIER CloudPcStorageAccountAccessTier = iota
    COOL_CLOUDPCSTORAGEACCOUNTACCESSTIER
    PREMIUM_CLOUDPCSTORAGEACCOUNTACCESSTIER
    COLD_CLOUDPCSTORAGEACCOUNTACCESSTIER
    UNKNOWNFUTUREVALUE_CLOUDPCSTORAGEACCOUNTACCESSTIER
)

func (i CloudPcStorageAccountAccessTier) String() string {
    return []string{"hot", "cool", "premium", "cold", "unknownFutureValue"}[i]
}
func ParseCloudPcStorageAccountAccessTier(v string) (any, error) {
    result := HOT_CLOUDPCSTORAGEACCOUNTACCESSTIER
    switch v {
        case "hot":
            result = HOT_CLOUDPCSTORAGEACCOUNTACCESSTIER
        case "cool":
            result = COOL_CLOUDPCSTORAGEACCOUNTACCESSTIER
        case "premium":
            result = PREMIUM_CLOUDPCSTORAGEACCOUNTACCESSTIER
        case "cold":
            result = COLD_CLOUDPCSTORAGEACCOUNTACCESSTIER
        case "unknownFutureValue":
            result = UNKNOWNFUTUREVALUE_CLOUDPCSTORAGEACCOUNTACCESSTIER
        default:
            return nil, nil
    }
    return &result, nil
}
func SerializeCloudPcStorageAccountAccessTier(values []CloudPcStorageAccountAccessTier) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i CloudPcStorageAccountAccessTier) isMultiValue() bool {
    return false
}
