package models
// Windows Operating System is available in different editions, which have a specific set of features available. This enum type defines the corresponding edition.
type ManagedDeviceWindowsOperatingSystemEditionType int

const (
    // Default. Indicates Professional Operating System edition used for the managed device.
    PROFESSIONAL_MANAGEDDEVICEWINDOWSOPERATINGSYSTEMEDITIONTYPE ManagedDeviceWindowsOperatingSystemEditionType = iota
    // Indicates Professional N Operating System edition used for the managed device.
    PROFESSIONALN_MANAGEDDEVICEWINDOWSOPERATINGSYSTEMEDITIONTYPE
    // Indicates Enterprise Operating System edition used for the managed device.
    ENTERPRISE_MANAGEDDEVICEWINDOWSOPERATINGSYSTEMEDITIONTYPE
    // Indicates Enterprise N Operating System edition used for the managed device.
    ENTERPRISEN_MANAGEDDEVICEWINDOWSOPERATINGSYSTEMEDITIONTYPE
    // Indicates Education Operating System edition used for the managed device.
    EDUCATION_MANAGEDDEVICEWINDOWSOPERATINGSYSTEMEDITIONTYPE
    // Indicates Education N Operating System edition used for the managed device.
    EDUCATIONN_MANAGEDDEVICEWINDOWSOPERATINGSYSTEMEDITIONTYPE
    // Indicates Pro Education Operating System edition used for the managed device.
    PROEDUCATION_MANAGEDDEVICEWINDOWSOPERATINGSYSTEMEDITIONTYPE
    // Indicates Pro Education N Operating System edition used for the managed device.
    PROEDUCATIONN_MANAGEDDEVICEWINDOWSOPERATINGSYSTEMEDITIONTYPE
    // Indicates Pro Workstation Operating System edition used for the managed device.
    PROWORKSTATION_MANAGEDDEVICEWINDOWSOPERATINGSYSTEMEDITIONTYPE
    // Indicates Pro Workstation N Operating System edition used for the managed device.
    PROWORKSTATIONN_MANAGEDDEVICEWINDOWSOPERATINGSYSTEMEDITIONTYPE
    // Evolvable enumeration sentinel value. Do not use.
    UNKNOWNFUTUREVALUE_MANAGEDDEVICEWINDOWSOPERATINGSYSTEMEDITIONTYPE
)

func (i ManagedDeviceWindowsOperatingSystemEditionType) String() string {
    return []string{"professional", "professionalN", "enterprise", "enterpriseN", "education", "educationN", "proEducation", "proEducationN", "proWorkstation", "proWorkstationN", "unknownFutureValue"}[i]
}
func ParseManagedDeviceWindowsOperatingSystemEditionType(v string) (any, error) {
    result := PROFESSIONAL_MANAGEDDEVICEWINDOWSOPERATINGSYSTEMEDITIONTYPE
    switch v {
        case "professional":
            result = PROFESSIONAL_MANAGEDDEVICEWINDOWSOPERATINGSYSTEMEDITIONTYPE
        case "professionalN":
            result = PROFESSIONALN_MANAGEDDEVICEWINDOWSOPERATINGSYSTEMEDITIONTYPE
        case "enterprise":
            result = ENTERPRISE_MANAGEDDEVICEWINDOWSOPERATINGSYSTEMEDITIONTYPE
        case "enterpriseN":
            result = ENTERPRISEN_MANAGEDDEVICEWINDOWSOPERATINGSYSTEMEDITIONTYPE
        case "education":
            result = EDUCATION_MANAGEDDEVICEWINDOWSOPERATINGSYSTEMEDITIONTYPE
        case "educationN":
            result = EDUCATIONN_MANAGEDDEVICEWINDOWSOPERATINGSYSTEMEDITIONTYPE
        case "proEducation":
            result = PROEDUCATION_MANAGEDDEVICEWINDOWSOPERATINGSYSTEMEDITIONTYPE
        case "proEducationN":
            result = PROEDUCATIONN_MANAGEDDEVICEWINDOWSOPERATINGSYSTEMEDITIONTYPE
        case "proWorkstation":
            result = PROWORKSTATION_MANAGEDDEVICEWINDOWSOPERATINGSYSTEMEDITIONTYPE
        case "proWorkstationN":
            result = PROWORKSTATIONN_MANAGEDDEVICEWINDOWSOPERATINGSYSTEMEDITIONTYPE
        case "unknownFutureValue":
            result = UNKNOWNFUTUREVALUE_MANAGEDDEVICEWINDOWSOPERATINGSYSTEMEDITIONTYPE
        default:
            return nil, nil
    }
    return &result, nil
}
func SerializeManagedDeviceWindowsOperatingSystemEditionType(values []ManagedDeviceWindowsOperatingSystemEditionType) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i ManagedDeviceWindowsOperatingSystemEditionType) isMultiValue() bool {
    return false
}
