package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// DeviceEnrollmentNotificationConfigurationable 
type DeviceEnrollmentNotificationConfigurationable interface {
    DeviceEnrollmentConfigurationable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetBrandingOptions()(*EnrollmentNotificationBrandingOptions)
    GetDefaultLocale()(*string)
    GetNotificationMessageTemplateId()(*string)
    GetNotificationTemplates()([]string)
    GetPlatformType()(*EnrollmentRestrictionPlatformType)
    GetTemplateType()(*EnrollmentNotificationTemplateType)
    SetBrandingOptions(value *EnrollmentNotificationBrandingOptions)()
    SetDefaultLocale(value *string)()
    SetNotificationMessageTemplateId(value *string)()
    SetNotificationTemplates(value []string)()
    SetPlatformType(value *EnrollmentRestrictionPlatformType)()
    SetTemplateType(value *EnrollmentNotificationTemplateType)()
}
