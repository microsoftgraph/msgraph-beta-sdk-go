package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// DeviceEnrollmentNotificationConfiguration 
type DeviceEnrollmentNotificationConfiguration struct {
    DeviceEnrollmentConfiguration
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // Branding Options for the Enrollment Notification. Possible values are: none, includeCompanyLogo, includeCompanyName, includeContactInformation, includeCompanyPortalLink, includeDeviceDetails.
    brandingOptions *EnrollmentNotificationBrandingOptions
    // DefaultLocale for the Enrollment Notification
    defaultLocale *string
    // Notification Message Template Id
    notificationMessageTemplateId *string
    // Platform type of the Enrollment Notification. Possible values are: allPlatforms, ios, windows, windowsPhone, android, androidForWork, mac.
    platformType *EnrollmentRestrictionPlatformType
    // Template type of the Enrollment Notification. Possible values are: email, push, unknownFutureValue.
    templateType *EnrollmentNotificationTemplateType
}
// NewDeviceEnrollmentNotificationConfiguration instantiates a new DeviceEnrollmentNotificationConfiguration and sets the default values.
func NewDeviceEnrollmentNotificationConfiguration()(*DeviceEnrollmentNotificationConfiguration) {
    m := &DeviceEnrollmentNotificationConfiguration{
        DeviceEnrollmentConfiguration: *NewDeviceEnrollmentConfiguration(),
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateDeviceEnrollmentNotificationConfigurationFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateDeviceEnrollmentNotificationConfigurationFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewDeviceEnrollmentNotificationConfiguration(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *DeviceEnrollmentNotificationConfiguration) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetBrandingOptions gets the brandingOptions property value. Branding Options for the Enrollment Notification. Possible values are: none, includeCompanyLogo, includeCompanyName, includeContactInformation, includeCompanyPortalLink, includeDeviceDetails.
func (m *DeviceEnrollmentNotificationConfiguration) GetBrandingOptions()(*EnrollmentNotificationBrandingOptions) {
    if m == nil {
        return nil
    } else {
        return m.brandingOptions
    }
}
// GetDefaultLocale gets the defaultLocale property value. DefaultLocale for the Enrollment Notification
func (m *DeviceEnrollmentNotificationConfiguration) GetDefaultLocale()(*string) {
    if m == nil {
        return nil
    } else {
        return m.defaultLocale
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *DeviceEnrollmentNotificationConfiguration) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.DeviceEnrollmentConfiguration.GetFieldDeserializers()
    res["brandingOptions"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseEnrollmentNotificationBrandingOptions)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetBrandingOptions(val.(*EnrollmentNotificationBrandingOptions))
        }
        return nil
    }
    res["defaultLocale"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDefaultLocale(val)
        }
        return nil
    }
    res["notificationMessageTemplateId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetNotificationMessageTemplateId(val)
        }
        return nil
    }
    res["platformType"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseEnrollmentRestrictionPlatformType)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPlatformType(val.(*EnrollmentRestrictionPlatformType))
        }
        return nil
    }
    res["templateType"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseEnrollmentNotificationTemplateType)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTemplateType(val.(*EnrollmentNotificationTemplateType))
        }
        return nil
    }
    return res
}
// GetNotificationMessageTemplateId gets the notificationMessageTemplateId property value. Notification Message Template Id
func (m *DeviceEnrollmentNotificationConfiguration) GetNotificationMessageTemplateId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.notificationMessageTemplateId
    }
}
// GetPlatformType gets the platformType property value. Platform type of the Enrollment Notification. Possible values are: allPlatforms, ios, windows, windowsPhone, android, androidForWork, mac.
func (m *DeviceEnrollmentNotificationConfiguration) GetPlatformType()(*EnrollmentRestrictionPlatformType) {
    if m == nil {
        return nil
    } else {
        return m.platformType
    }
}
// GetTemplateType gets the templateType property value. Template type of the Enrollment Notification. Possible values are: email, push, unknownFutureValue.
func (m *DeviceEnrollmentNotificationConfiguration) GetTemplateType()(*EnrollmentNotificationTemplateType) {
    if m == nil {
        return nil
    } else {
        return m.templateType
    }
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
    {
        err = writer.WriteAdditionalData(m.GetAdditionalData())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAdditionalData sets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *DeviceEnrollmentNotificationConfiguration) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
// SetBrandingOptions sets the brandingOptions property value. Branding Options for the Enrollment Notification. Possible values are: none, includeCompanyLogo, includeCompanyName, includeContactInformation, includeCompanyPortalLink, includeDeviceDetails.
func (m *DeviceEnrollmentNotificationConfiguration) SetBrandingOptions(value *EnrollmentNotificationBrandingOptions)() {
    if m != nil {
        m.brandingOptions = value
    }
}
// SetDefaultLocale sets the defaultLocale property value. DefaultLocale for the Enrollment Notification
func (m *DeviceEnrollmentNotificationConfiguration) SetDefaultLocale(value *string)() {
    if m != nil {
        m.defaultLocale = value
    }
}
// SetNotificationMessageTemplateId sets the notificationMessageTemplateId property value. Notification Message Template Id
func (m *DeviceEnrollmentNotificationConfiguration) SetNotificationMessageTemplateId(value *string)() {
    if m != nil {
        m.notificationMessageTemplateId = value
    }
}
// SetPlatformType sets the platformType property value. Platform type of the Enrollment Notification. Possible values are: allPlatforms, ios, windows, windowsPhone, android, androidForWork, mac.
func (m *DeviceEnrollmentNotificationConfiguration) SetPlatformType(value *EnrollmentRestrictionPlatformType)() {
    if m != nil {
        m.platformType = value
    }
}
// SetTemplateType sets the templateType property value. Template type of the Enrollment Notification. Possible values are: email, push, unknownFutureValue.
func (m *DeviceEnrollmentNotificationConfiguration) SetTemplateType(value *EnrollmentNotificationTemplateType)() {
    if m != nil {
        m.templateType = value
    }
}
