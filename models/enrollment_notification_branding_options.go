package models
import (
    "errors"
    "math"
    "strings"
)
// Branding Options for the Message Template. Branding is defined in the Intune Admin Console.
type EnrollmentNotificationBrandingOptions int

const (
    // Indicates that the template has no branding.
    NONE_ENROLLMENTNOTIFICATIONBRANDINGOPTIONS = 1
    // Indicates that the Company Logo is included in the notification.
    INCLUDECOMPANYLOGO_ENROLLMENTNOTIFICATIONBRANDINGOPTIONS = 2
    // Indicates that the Company Name is included in the notification.
    INCLUDECOMPANYNAME_ENROLLMENTNOTIFICATIONBRANDINGOPTIONS = 4
    // Indicates that the Contact Information is included in the notification.
    INCLUDECONTACTINFORMATION_ENROLLMENTNOTIFICATIONBRANDINGOPTIONS = 8
    // Indicates that the Company Portal Link is included in the notification.
    INCLUDECOMPANYPORTALLINK_ENROLLMENTNOTIFICATIONBRANDINGOPTIONS = 16
    // Indicates that the DeviceDetails is included in the notification.
    INCLUDEDEVICEDETAILS_ENROLLMENTNOTIFICATIONBRANDINGOPTIONS = 32
    // unknownFutureValue for evolvable enums pattern.
    UNKNOWNFUTUREVALUE_ENROLLMENTNOTIFICATIONBRANDINGOPTIONS = 64
)

func (i EnrollmentNotificationBrandingOptions) String() string {
    var values []string
    options := []string{"none", "includeCompanyLogo", "includeCompanyName", "includeContactInformation", "includeCompanyPortalLink", "includeDeviceDetails", "unknownFutureValue"}
    for p := 0; p < 7; p++ {
        mantis := EnrollmentNotificationBrandingOptions(int(math.Pow(2, float64(p))))
        if i&mantis == mantis {
            values = append(values, options[p])
        }
    }
    return strings.Join(values, ",")
}
func ParseEnrollmentNotificationBrandingOptions(v string) (any, error) {
    var result EnrollmentNotificationBrandingOptions
    values := strings.Split(v, ",")
    for _, str := range values {
        switch str {
            case "none":
                result |= NONE_ENROLLMENTNOTIFICATIONBRANDINGOPTIONS
            case "includeCompanyLogo":
                result |= INCLUDECOMPANYLOGO_ENROLLMENTNOTIFICATIONBRANDINGOPTIONS
            case "includeCompanyName":
                result |= INCLUDECOMPANYNAME_ENROLLMENTNOTIFICATIONBRANDINGOPTIONS
            case "includeContactInformation":
                result |= INCLUDECONTACTINFORMATION_ENROLLMENTNOTIFICATIONBRANDINGOPTIONS
            case "includeCompanyPortalLink":
                result |= INCLUDECOMPANYPORTALLINK_ENROLLMENTNOTIFICATIONBRANDINGOPTIONS
            case "includeDeviceDetails":
                result |= INCLUDEDEVICEDETAILS_ENROLLMENTNOTIFICATIONBRANDINGOPTIONS
            case "unknownFutureValue":
                result |= UNKNOWNFUTUREVALUE_ENROLLMENTNOTIFICATIONBRANDINGOPTIONS
            default:
                return 0, errors.New("Unknown EnrollmentNotificationBrandingOptions value: " + v)
        }
    }
    return &result, nil
}
func SerializeEnrollmentNotificationBrandingOptions(values []EnrollmentNotificationBrandingOptions) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i EnrollmentNotificationBrandingOptions) isMultiValue() bool {
    return true
}
