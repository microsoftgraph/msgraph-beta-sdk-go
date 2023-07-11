package models

import (
    i561e97a8befe7661a44c8f54600992b4207a3a0cf6770e5559949bc276de2e22 "github.com/google/uuid"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// DeviceEnrollmentNotificationConfiguration enrollment Notification Configuration which is used to send notification
type DeviceEnrollmentNotificationConfiguration struct {
    DeviceEnrollmentConfiguration
}
// NewDeviceEnrollmentNotificationConfiguration instantiates a new deviceEnrollmentNotificationConfiguration and sets the default values.
func NewDeviceEnrollmentNotificationConfiguration()(*DeviceEnrollmentNotificationConfiguration) {
    m := &DeviceEnrollmentNotificationConfiguration{
        DeviceEnrollmentConfiguration: *NewDeviceEnrollmentConfiguration(),
    }
    odataTypeValue := "#microsoft.graph.deviceEnrollmentNotificationConfiguration"
    m.SetOdataType(&odataTypeValue)
    return m
}
// CreateDeviceEnrollmentNotificationConfigurationFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateDeviceEnrollmentNotificationConfigurationFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewDeviceEnrollmentNotificationConfiguration(), nil
}
// GetBrandingOptions gets the brandingOptions property value. Branding Options for the Message Template. Branding is defined in the Intune Admin Console.
func (m *DeviceEnrollmentNotificationConfiguration) GetBrandingOptions()(*EnrollmentNotificationBrandingOptions) {
    val, err := m.GetBackingStore().Get("brandingOptions")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*EnrollmentNotificationBrandingOptions)
    }
    return nil
}
// GetDefaultLocale gets the defaultLocale property value. DefaultLocale for the Enrollment Notification
func (m *DeviceEnrollmentNotificationConfiguration) GetDefaultLocale()(*string) {
    val, err := m.GetBackingStore().Get("defaultLocale")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
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
        val, err := n.GetUUIDValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetNotificationMessageTemplateId(val)
        }
        return nil
    }
    res["notificationTemplates"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]string, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = *(v.(*string))
                }
            }
            m.SetNotificationTemplates(res)
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
func (m *DeviceEnrollmentNotificationConfiguration) GetNotificationMessageTemplateId()(*i561e97a8befe7661a44c8f54600992b4207a3a0cf6770e5559949bc276de2e22.UUID) {
    val, err := m.GetBackingStore().Get("notificationMessageTemplateId")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*i561e97a8befe7661a44c8f54600992b4207a3a0cf6770e5559949bc276de2e22.UUID)
    }
    return nil
}
// GetNotificationTemplates gets the notificationTemplates property value. The list of notification data -
func (m *DeviceEnrollmentNotificationConfiguration) GetNotificationTemplates()([]string) {
    val, err := m.GetBackingStore().Get("notificationTemplates")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]string)
    }
    return nil
}
// GetPlatformType gets the platformType property value. This enum indicates the platform type for which the enrollment restriction applies.
func (m *DeviceEnrollmentNotificationConfiguration) GetPlatformType()(*EnrollmentRestrictionPlatformType) {
    val, err := m.GetBackingStore().Get("platformType")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*EnrollmentRestrictionPlatformType)
    }
    return nil
}
// GetTemplateType gets the templateType property value. This enum indicates the Template type for which the enrollment notification applies.
func (m *DeviceEnrollmentNotificationConfiguration) GetTemplateType()(*EnrollmentNotificationTemplateType) {
    val, err := m.GetBackingStore().Get("templateType")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*EnrollmentNotificationTemplateType)
    }
    return nil
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
        err = writer.WriteUUIDValue("notificationMessageTemplateId", m.GetNotificationMessageTemplateId())
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
    err := m.GetBackingStore().Set("brandingOptions", value)
    if err != nil {
        panic(err)
    }
}
// SetDefaultLocale sets the defaultLocale property value. DefaultLocale for the Enrollment Notification
func (m *DeviceEnrollmentNotificationConfiguration) SetDefaultLocale(value *string)() {
    err := m.GetBackingStore().Set("defaultLocale", value)
    if err != nil {
        panic(err)
    }
}
// SetNotificationMessageTemplateId sets the notificationMessageTemplateId property value. Notification Message Template Id
func (m *DeviceEnrollmentNotificationConfiguration) SetNotificationMessageTemplateId(value *i561e97a8befe7661a44c8f54600992b4207a3a0cf6770e5559949bc276de2e22.UUID)() {
    err := m.GetBackingStore().Set("notificationMessageTemplateId", value)
    if err != nil {
        panic(err)
    }
}
// SetNotificationTemplates sets the notificationTemplates property value. The list of notification data -
func (m *DeviceEnrollmentNotificationConfiguration) SetNotificationTemplates(value []string)() {
    err := m.GetBackingStore().Set("notificationTemplates", value)
    if err != nil {
        panic(err)
    }
}
// SetPlatformType sets the platformType property value. This enum indicates the platform type for which the enrollment restriction applies.
func (m *DeviceEnrollmentNotificationConfiguration) SetPlatformType(value *EnrollmentRestrictionPlatformType)() {
    err := m.GetBackingStore().Set("platformType", value)
    if err != nil {
        panic(err)
    }
}
// SetTemplateType sets the templateType property value. This enum indicates the Template type for which the enrollment notification applies.
func (m *DeviceEnrollmentNotificationConfiguration) SetTemplateType(value *EnrollmentNotificationTemplateType)() {
    err := m.GetBackingStore().Set("templateType", value)
    if err != nil {
        panic(err)
    }
}
// DeviceEnrollmentNotificationConfigurationable 
type DeviceEnrollmentNotificationConfigurationable interface {
    DeviceEnrollmentConfigurationable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetBrandingOptions()(*EnrollmentNotificationBrandingOptions)
    GetDefaultLocale()(*string)
    GetNotificationMessageTemplateId()(*i561e97a8befe7661a44c8f54600992b4207a3a0cf6770e5559949bc276de2e22.UUID)
    GetNotificationTemplates()([]string)
    GetPlatformType()(*EnrollmentRestrictionPlatformType)
    GetTemplateType()(*EnrollmentNotificationTemplateType)
    SetBrandingOptions(value *EnrollmentNotificationBrandingOptions)()
    SetDefaultLocale(value *string)()
    SetNotificationMessageTemplateId(value *i561e97a8befe7661a44c8f54600992b4207a3a0cf6770e5559949bc276de2e22.UUID)()
    SetNotificationTemplates(value []string)()
    SetPlatformType(value *EnrollmentRestrictionPlatformType)()
    SetTemplateType(value *EnrollmentNotificationTemplateType)()
}
