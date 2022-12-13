package models

import (
    i2bacd9b8d8db2e77ee2b5c5ccb19d679c36f920b8fee9d786a0adafff458afcd "github.com/google/UUID"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// DeviceEnrollmentNotificationConfigurationable 
type DeviceEnrollmentNotificationConfigurationable interface {
    DeviceEnrollmentConfigurationable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetBrandingOptions()(*EnrollmentNotificationBrandingOptions)
    GetDefaultLocale()(*string)
    GetNotificationMessageTemplateId()(*UUID)
    GetNotificationTemplates()([]string)
    GetPlatformType()(*EnrollmentRestrictionPlatformType)
    GetTemplateType()(*EnrollmentNotificationTemplateType)
    SetBrandingOptions(value *EnrollmentNotificationBrandingOptions)()
    SetDefaultLocale(value *string)()
    SetNotificationMessageTemplateId(value *UUID)()
    SetNotificationTemplates(value []string)()
    SetPlatformType(value *EnrollmentRestrictionPlatformType)()
    SetTemplateType(value *EnrollmentNotificationTemplateType)()
}
