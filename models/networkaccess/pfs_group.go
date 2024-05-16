package networkaccess
type PfsGroup int

const (
    NONE_PFSGROUP PfsGroup = iota
    PFS1_PFSGROUP
    PFS2_PFSGROUP
    PFS14_PFSGROUP
    PFS24_PFSGROUP
    PFS2048_PFSGROUP
    PFSMM_PFSGROUP
    ECP256_PFSGROUP
    ECP384_PFSGROUP
    UNKNOWNFUTUREVALUE_PFSGROUP
)

func (i PfsGroup) String() string {
    return []string{"none", "pfs1", "pfs2", "pfs14", "pfs24", "pfs2048", "pfsmm", "ecp256", "ecp384", "unknownFutureValue"}[i]
}
func ParsePfsGroup(v string) (any, error) {
    result := NONE_PFSGROUP
    switch v {
        case "none":
            result = NONE_PFSGROUP
        case "pfs1":
            result = PFS1_PFSGROUP
        case "pfs2":
            result = PFS2_PFSGROUP
        case "pfs14":
            result = PFS14_PFSGROUP
        case "pfs24":
            result = PFS24_PFSGROUP
        case "pfs2048":
            result = PFS2048_PFSGROUP
        case "pfsmm":
            result = PFSMM_PFSGROUP
        case "ecp256":
            result = ECP256_PFSGROUP
        case "ecp384":
            result = ECP384_PFSGROUP
        case "unknownFutureValue":
            result = UNKNOWNFUTUREVALUE_PFSGROUP
        default:
            return nil, nil
    }
    return &result, nil
}
func SerializePfsGroup(values []PfsGroup) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i PfsGroup) isMultiValue() bool {
    return false
}
