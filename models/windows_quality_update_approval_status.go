// Code generated by Microsoft Kiota - DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package models
// Enum to describe policy's approval status for catalogitems 
type WindowsQualityUpdateApprovalStatus int

const (
    // unknown status for corresponding catalog item
    UNKNOWN_WINDOWSQUALITYUPDATEAPPROVALSTATUS WindowsQualityUpdateApprovalStatus = iota
    // approved for corresponding catalog item
    APPROVED_WINDOWSQUALITYUPDATEAPPROVALSTATUS
    // suspended for corresponding catalog item
    SUSPENDED_WINDOWSQUALITYUPDATEAPPROVALSTATUS
    // Evolvable enumeration sentinel value. Do not use.
    UNKNOWNFUTUREVALUE_WINDOWSQUALITYUPDATEAPPROVALSTATUS
)

func (i WindowsQualityUpdateApprovalStatus) String() string {
    return []string{"unknown", "approved", "suspended", "unknownFutureValue"}[i]
}
func ParseWindowsQualityUpdateApprovalStatus(v string) (any, error) {
    result := UNKNOWN_WINDOWSQUALITYUPDATEAPPROVALSTATUS
    switch v {
        case "unknown":
            result = UNKNOWN_WINDOWSQUALITYUPDATEAPPROVALSTATUS
        case "approved":
            result = APPROVED_WINDOWSQUALITYUPDATEAPPROVALSTATUS
        case "suspended":
            result = SUSPENDED_WINDOWSQUALITYUPDATEAPPROVALSTATUS
        case "unknownFutureValue":
            result = UNKNOWNFUTUREVALUE_WINDOWSQUALITYUPDATEAPPROVALSTATUS
        default:
            return nil, nil
    }
    return &result, nil
}
func SerializeWindowsQualityUpdateApprovalStatus(values []WindowsQualityUpdateApprovalStatus) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i WindowsQualityUpdateApprovalStatus) isMultiValue() bool {
    return false
}
