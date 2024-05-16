package models
// A list of possible System Management Mode levels for a device. System Management Mode levels is determined by report sent from Microsoft Azure Attestation service. Only specific hardwares support System Management Mode. Windows 11 devices will have values "notApplicable", "level1", "level2" or "level3". Windows 10 devices will have value "notApplicable".
type SystemManagementModeLevel int

const (
    // Indicates that the device does not have Firmware protection (System Management Mode) enabled.
    NOTAPPLICABLE_SYSTEMMANAGEMENTMODELEVEL SystemManagementModeLevel = iota
    // Indicates that deny System Management Mode (SMM) read/write access to OS and Virtualization-based security (VBS) memory. The benefit is that by design System Management Mode (SMM) cannot modify security of or exfiltrate secrets from the OS (including Virtualization-based security).
    LEVEL1_SYSTEMMANAGEMENTMODELEVEL
    // Indicates that in addition to the System Management Mode (SMM) Level 1 protections, this level prevents System Management Mode (SMM) from tampering with Input-Output Memory Management Unit (IOMMU) config. The benefit is that by design System Management Mode (SMM) cannot disable Virtualization-based security (VBS) and kernel Direct memory access (DMA) protections.
    LEVEL2_SYSTEMMANAGEMENTMODELEVEL
    // Indicates that in addition to the System Management Mode (SMM) Level 2 protections, this level reduces System Management Mode (SMM) save state capabilities. The benefit is that by design System Management Mode (SMM) cannot exploit save state to modify security of or exfiltrate secrets from the OS (including Virtualization-based security).
    LEVEL3_SYSTEMMANAGEMENTMODELEVEL
    // Evolvable enumeration sentinel value. Do not use.
    UNKNOWNFUTUREVALUE_SYSTEMMANAGEMENTMODELEVEL
)

func (i SystemManagementModeLevel) String() string {
    return []string{"notApplicable", "level1", "level2", "level3", "unknownFutureValue"}[i]
}
func ParseSystemManagementModeLevel(v string) (any, error) {
    result := NOTAPPLICABLE_SYSTEMMANAGEMENTMODELEVEL
    switch v {
        case "notApplicable":
            result = NOTAPPLICABLE_SYSTEMMANAGEMENTMODELEVEL
        case "level1":
            result = LEVEL1_SYSTEMMANAGEMENTMODELEVEL
        case "level2":
            result = LEVEL2_SYSTEMMANAGEMENTMODELEVEL
        case "level3":
            result = LEVEL3_SYSTEMMANAGEMENTMODELEVEL
        case "unknownFutureValue":
            result = UNKNOWNFUTUREVALUE_SYSTEMMANAGEMENTMODELEVEL
        default:
            return nil, nil
    }
    return &result, nil
}
func SerializeSystemManagementModeLevel(values []SystemManagementModeLevel) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i SystemManagementModeLevel) isMultiValue() bool {
    return false
}
