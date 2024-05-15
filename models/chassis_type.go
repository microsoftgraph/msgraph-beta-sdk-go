package models
// Chassis type.
type ChassisType int

const (
    // Unknown.
    UNKNOWN_CHASSISTYPE ChassisType = iota
    // Desktop.
    DESKTOP_CHASSISTYPE
    // Laptop.
    LAPTOP_CHASSISTYPE
    // Workstation.
    WORKSWORKSTATION_CHASSISTYPE
    // Enterprise server.
    ENTERPRISESERVER_CHASSISTYPE
    // Phone.
    PHONE_CHASSISTYPE
    // Mobile tablet.
    TABLET_CHASSISTYPE
    // Other mobile.
    MOBILEOTHER_CHASSISTYPE
    // Unknown mobile.
    MOBILEUNKNOWN_CHASSISTYPE
)

func (i ChassisType) String() string {
    return []string{"unknown", "desktop", "laptop", "worksWorkstation", "enterpriseServer", "phone", "tablet", "mobileOther", "mobileUnknown"}[i]
}
func ParseChassisType(v string) (any, error) {
    result := UNKNOWN_CHASSISTYPE
    switch v {
        case "unknown":
            result = UNKNOWN_CHASSISTYPE
        case "desktop":
            result = DESKTOP_CHASSISTYPE
        case "laptop":
            result = LAPTOP_CHASSISTYPE
        case "worksWorkstation":
            result = WORKSWORKSTATION_CHASSISTYPE
        case "enterpriseServer":
            result = ENTERPRISESERVER_CHASSISTYPE
        case "phone":
            result = PHONE_CHASSISTYPE
        case "tablet":
            result = TABLET_CHASSISTYPE
        case "mobileOther":
            result = MOBILEOTHER_CHASSISTYPE
        case "mobileUnknown":
            result = MOBILEUNKNOWN_CHASSISTYPE
        default:
            return nil, nil
    }
    return &result, nil
}
func SerializeChassisType(values []ChassisType) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i ChassisType) isMultiValue() bool {
    return false
}
