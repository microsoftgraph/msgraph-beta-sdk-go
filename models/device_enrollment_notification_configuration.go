package models

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// DeviceEnrollmentNotificationConfiguration 
type DeviceEnrollmentNotificationConfiguration struct {
    DeviceEnrollmentConfiguration
    // Branding Options for the Message Template. Branding is defined in the Intune Admin Console.
    brandingOptions *EnrollmentNotificationBrandingOptions
    // DefaultLocale for the Enrollment Notification
    defaultLocale *string
    // Notification Message Template Id
    notificationMessageTemplateId *string
    // The list of notification data -
    notificationTemplates []string
    // This enum indicates the platform type for which the enrollment restriction applies.
    platformType *EnrollmentRestrictionPlatformType
    // This enum indicates the Template type for which the enrollment notification applies.
    templateType *EnrollmentNotificationTemplateType
}
// NewDeviceEnrollmentNotificationConfiguration instantiates a new DeviceEnrollmentNotificationConfiguration and sets the default values.
func NewDeviceEnrollmentNotificationConfiguration()(*DeviceEnrollmentNotificationConfiguration) {
    m := &DeviceEnrollmentNotificationConfiguration{
        DeviceEnrollmentConfiguration: *NewDeviceEnrollmentConfiguration(),
    }
    odataTypeValue := "#microsoft.graph.deviceEnrollmentNotificationConfiguration";
    m.SetOdataType(&odataTypeValue);
    return m
}
// CreateDeviceEnrollmentNotificationConfigurationFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateDeviceEnrollmentNotificationConfigurationFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewDeviceEnrollmentNotificationConfiguration(), nil
}
// GetBrandingOptions gets the brandingOptions property value. Branding Options for the Message Template. Branding is defined in the Intune Admin Console.
func (m *DeviceEnrollmentNotificationConfiguration) GetBrandingOptions()(*EnrollmentNotificationBrandingOptions) {
    return m.brandingOptions
}
// GetDefaultLocale gets the defaultLocale property value. DefaultLocale for the Enrollment Notification
func (m *DeviceEnrollmentNotificationConfiguration) GetDefaultLocale()(*string) {
    return m.defaultLocale
}
// GetFieldDeserializers the deserialization information for the current model
func (m *DeviceEnrollmentNotificationConfiguration) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.DeviceEnrollmentConfiguration.GetFieldDeserializers()
    res["brandingOptions"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetEnumValue(ParseEnrollmentNotificationBrandingOptions , m.SetBrandingOptions)
    res["defaultLocale"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetDefaultLocale)
    res["notificationMessageTemplateId"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetNotificationMessageTemplateId)
    res["notificationTemplates"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfPrimitiveValues("string" , m.SetNotificationTemplates)
    res["platformType"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetEnumValue(ParseEnrollmentRestrictionPlatformType , m.SetPlatformType)
    res["templateType"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetEnumValue(ParseEnrollmentNotificationTemplateType , m.SetTemplateType)
    return res
}
// GetNotificationMessageTemplateId gets the notificationMessageTemplateId property value. Notification Message Template Id
func (m *DeviceEnrollmentNotificationConfiguration) GetNotificationMessageTemplateId()(*string) {
    return m.notificationMessageTemplateId
}
// GetNotificationTemplates gets the notificationTemplates property value. The list of notification data -
func (m *DeviceEnrollmentNotificationConfiguration) GetNotificationTemplates()([]string) {
    return m.notificationTemplates
}
// GetPlatformType gets the platformType property value. This enum indicates the platform type for which the enrollment restriction applies.
func (m *DeviceEnrollmentNotificationConfiguration) GetPlatformType()(*EnrollmentRestrictionPlatformType) {
    return m.platformType
}
// GetTemplateType gets the templateType property value. This enum indicates the Template type for which the enrollment notification applies.
func (m *DeviceEnrollmentNotificationConfiguration) GetTemplateType()(*EnrollmentNotificationTemplateType) {
    return m.templateType
}
// Serialize serializes information the current object
func (m *DeviceEnrollmentNotificationConfiguration) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.DeviceEnrollmentConfiguration.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetBrandingOptions() != nil {
        cast := (*m.GetBrandingOptions()).String()
        err = writer.WriteStringValue("brandingOptions", &cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("defaultLocale", m.GetDefaultLocale())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("notificationMessageTemplateId", m.GetNotificationMessageTemplateId())
        if err != nil {
            return err
        }
    }
    if m.GetNotificationTemplates() != nil {
        err = writer.WriteCollectionOfStringValues("notificationTemplates", m.GetNotificationTemplates())
        if err != nil {
            return err
        }
    }
    if m.GetPlatformType() != nil {
        cast := (*m.GetPlatformType()).String()
        err = writer.WriteStringValue("platformType", &cast)
        if err != nil {
            return err
        }
    }
    if m.GetTemplateType() != nil {
        cast := (*m.GetTemplateType()).String()
        err = writer.WriteStringValue("templateType", &cast)
        if err != nil {
            return err
        }
    }
    return nil
}
// SetBrandingOptions sets the brandingOptions property value. Branding Options for the Message Template. Branding is defined in the Intune Admin Console.
func (m *DeviceEnrollmentNotificationConfiguration) SetBrandingOptions(value *EnrollmentNotificationBrandingOptions)() {
    m.brandingOptions = value
}
// SetDefaultLocale sets the defaultLocale property value. DefaultLocale for the Enrollment Notification
func (m *DeviceEnrollmentNotificationConfiguration) SetDefaultLocale(value *string)() {
    m.defaultLocale = value
}
// SetNotificationMessageTemplateId sets the notificationMessageTemplateId property value. Notification Message Template Id
func (m *DeviceEnrollmentNotificationConfiguration) SetNotificationMessageTemplateId(value *string)() {
    m.notificationMessageTemplateId = value
}
// SetNotificationTemplates sets the notificationTemplates property value. The list of notification data -
func (m *DeviceEnrollmentNotificationConfiguration) SetNotificationTemplates(value []string)() {
    m.notificationTemplates = value
}
// SetPlatformType sets the platformType property value. This enum indicates the platform type for which the enrollment restriction applies.
func (m *DeviceEnrollmentNotificationConfiguration) SetPlatformType(value *EnrollmentRestrictionPlatformType)() {
    m.platformType = value
}
// SetTemplateType sets the templateType property value. This enum indicates the Template type for which the enrollment notification applies.
func (m *DeviceEnrollmentNotificationConfiguration) SetTemplateType(value *EnrollmentNotificationTemplateType)() {
    m.templateType = value
}
