package models
// Possible values for the Certificate Destination Store.
type CertificateDestinationStore int

const (
    // Computer Certificate Store - Root.
    COMPUTERCERTSTOREROOT_CERTIFICATEDESTINATIONSTORE CertificateDestinationStore = iota
    // Computer Certificate Store - Intermediate.
    COMPUTERCERTSTOREINTERMEDIATE_CERTIFICATEDESTINATIONSTORE
    // User Certificate Store - Intermediate.
    USERCERTSTOREINTERMEDIATE_CERTIFICATEDESTINATIONSTORE
)

func (i CertificateDestinationStore) String() string {
    return []string{"computerCertStoreRoot", "computerCertStoreIntermediate", "userCertStoreIntermediate"}[i]
}
func ParseCertificateDestinationStore(v string) (any, error) {
    result := COMPUTERCERTSTOREROOT_CERTIFICATEDESTINATIONSTORE
    switch v {
        case "computerCertStoreRoot":
            result = COMPUTERCERTSTOREROOT_CERTIFICATEDESTINATIONSTORE
        case "computerCertStoreIntermediate":
            result = COMPUTERCERTSTOREINTERMEDIATE_CERTIFICATEDESTINATIONSTORE
        case "userCertStoreIntermediate":
            result = USERCERTSTOREINTERMEDIATE_CERTIFICATEDESTINATIONSTORE
        default:
            return nil, nil
    }
    return &result, nil
}
func SerializeCertificateDestinationStore(values []CertificateDestinationStore) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i CertificateDestinationStore) isMultiValue() bool {
    return false
}
