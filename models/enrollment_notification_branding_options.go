package models
import (
    "errors"
)
// Provides operations to manage the collection of accessReviewDecision entities.
type EnrollmentNotificationBrandingOptions int

const (
    // No Branding.
    NONE_ENROLLMENTNOTIFICATIONBRANDINGOPTIONS EnrollmentNotificationBrandingOptions = iota
    // Include Company Logo.
    INCLUDECOMPANYLOGO_ENROLLMENTNOTIFICATIONBRANDINGOPTIONS
    // Include Company Name.
    INCLUDECOMPANYNAME_ENROLLMENTNOTIFICATIONBRANDINGOPTIONS
    // Include Contact Info.
    INCLUDECONTACTINFORMATION_ENROLLMENTNOTIFICATIONBRANDINGOPTIONS
    // Include Company Portal Link.
    INCLUDECOMPANYPORTALLINK_ENROLLMENTNOTIFICATIONBRANDINGOPTIONS
    // Include Device Details.
    INCLUDEDEVICEDETAILS_ENROLLMENTNOTIFICATIONBRANDINGOPTIONS
)

func (i EnrollmentNotificationBrandingOptions) String() string {
    return []string{"none", "includeCompanyLogo", "includeCompanyName", "includeContactInformation", "includeCompanyPortalLink", "includeDeviceDetails"}[i]
}
func ParseEnrollmentNotificationBrandingOptions(v string) (interface{}, error) {
    result := NONE_ENROLLMENTNOTIFICATIONBRANDINGOPTIONS
    switch v {
        case "none":
            result = NONE_ENROLLMENTNOTIFICATIONBRANDINGOPTIONS
        case "includeCompanyLogo":
            result = INCLUDECOMPANYLOGO_ENROLLMENTNOTIFICATIONBRANDINGOPTIONS
        case "includeCompanyName":
            result = INCLUDECOMPANYNAME_ENROLLMENTNOTIFICATIONBRANDINGOPTIONS
        case "includeContactInformation":
            result = INCLUDECONTACTINFORMATION_ENROLLMENTNOTIFICATIONBRANDINGOPTIONS
        case "includeCompanyPortalLink":
            result = INCLUDECOMPANYPORTALLINK_ENROLLMENTNOTIFICATIONBRANDINGOPTIONS
        case "includeDeviceDetails":
            result = INCLUDEDEVICEDETAILS_ENROLLMENTNOTIFICATIONBRANDINGOPTIONS
        default:
            return 0, errors.New("Unknown EnrollmentNotificationBrandingOptions value: " + v)
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
