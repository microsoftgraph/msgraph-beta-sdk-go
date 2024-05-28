package models
type FileStorageContainerOwnershipType int

const (
    TENANTOWNED_FILESTORAGECONTAINEROWNERSHIPTYPE FileStorageContainerOwnershipType = iota
    USEROWNED_FILESTORAGECONTAINEROWNERSHIPTYPE
    UNKNOWNFUTUREVALUE_FILESTORAGECONTAINEROWNERSHIPTYPE
)

func (i FileStorageContainerOwnershipType) String() string {
    return []string{"tenantOwned", "userOwned", "unknownFutureValue"}[i]
}
func ParseFileStorageContainerOwnershipType(v string) (any, error) {
    result := TENANTOWNED_FILESTORAGECONTAINEROWNERSHIPTYPE
    switch v {
        case "tenantOwned":
            result = TENANTOWNED_FILESTORAGECONTAINEROWNERSHIPTYPE
        case "userOwned":
            result = USEROWNED_FILESTORAGECONTAINEROWNERSHIPTYPE
        case "unknownFutureValue":
            result = UNKNOWNFUTUREVALUE_FILESTORAGECONTAINEROWNERSHIPTYPE
        default:
            return nil, nil
    }
    return &result, nil
}
func SerializeFileStorageContainerOwnershipType(values []FileStorageContainerOwnershipType) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i FileStorageContainerOwnershipType) isMultiValue() bool {
    return false
}
