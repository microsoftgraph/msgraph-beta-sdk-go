package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// IosEduDeviceConfiguration 
type IosEduDeviceConfiguration struct {
    DeviceConfiguration
}
// NewIosEduDeviceConfiguration instantiates a new IosEduDeviceConfiguration and sets the default values.
func NewIosEduDeviceConfiguration()(*IosEduDeviceConfiguration) {
    m := &IosEduDeviceConfiguration{
        DeviceConfiguration: *NewDeviceConfiguration(),
    }
    odataTypeValue := "#microsoft.graph.iosEduDeviceConfiguration"
    m.SetOdataType(&odataTypeValue)
    return m
}
// CreateIosEduDeviceConfigurationFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateIosEduDeviceConfigurationFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewIosEduDeviceConfiguration(), nil
}
// GetDeviceCertificateSettings gets the deviceCertificateSettings property value. The Trusted Root and PFX certificates for Device
func (m *IosEduDeviceConfiguration) GetDeviceCertificateSettings()(IosEduCertificateSettingsable) {
    val, err := m.GetBackingStore().Get("deviceCertificateSettings")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(IosEduCertificateSettingsable)
    }
    return nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *IosEduDeviceConfiguration) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.DeviceConfiguration.GetFieldDeserializers()
    res["deviceCertificateSettings"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateIosEduCertificateSettingsFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDeviceCertificateSettings(val.(IosEduCertificateSettingsable))
        }
        return nil
    }
    res["studentCertificateSettings"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateIosEduCertificateSettingsFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetStudentCertificateSettings(val.(IosEduCertificateSettingsable))
        }
        return nil
    }
    res["teacherCertificateSettings"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateIosEduCertificateSettingsFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTeacherCertificateSettings(val.(IosEduCertificateSettingsable))
        }
        return nil
    }
    return res
}
// GetStudentCertificateSettings gets the studentCertificateSettings property value. The Trusted Root and PFX certificates for Student
func (m *IosEduDeviceConfiguration) GetStudentCertificateSettings()(IosEduCertificateSettingsable) {
    val, err := m.GetBackingStore().Get("studentCertificateSettings")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(IosEduCertificateSettingsable)
    }
    return nil
}
// GetTeacherCertificateSettings gets the teacherCertificateSettings property value. Trusted Root and PFX certificates for iOS EDU.
func (m *IosEduDeviceConfiguration) GetTeacherCertificateSettings()(IosEduCertificateSettingsable) {
    val, err := m.GetBackingStore().Get("teacherCertificateSettings")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(IosEduCertificateSettingsable)
    }
    return nil
}
// Serialize serializes information the current object
func (m *IosEduDeviceConfiguration) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.DeviceConfiguration.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteObjectValue("deviceCertificateSettings", m.GetDeviceCertificateSettings())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("studentCertificateSettings", m.GetStudentCertificateSettings())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("teacherCertificateSettings", m.GetTeacherCertificateSettings())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetDeviceCertificateSettings sets the deviceCertificateSettings property value. The Trusted Root and PFX certificates for Device
func (m *IosEduDeviceConfiguration) SetDeviceCertificateSettings(value IosEduCertificateSettingsable)() {
    err := m.GetBackingStore().Set("deviceCertificateSettings", value)
    if err != nil {
        panic(err)
    }
}
// SetStudentCertificateSettings sets the studentCertificateSettings property value. The Trusted Root and PFX certificates for Student
func (m *IosEduDeviceConfiguration) SetStudentCertificateSettings(value IosEduCertificateSettingsable)() {
    err := m.GetBackingStore().Set("studentCertificateSettings", value)
    if err != nil {
        panic(err)
    }
}
// SetTeacherCertificateSettings sets the teacherCertificateSettings property value. Trusted Root and PFX certificates for iOS EDU.
func (m *IosEduDeviceConfiguration) SetTeacherCertificateSettings(value IosEduCertificateSettingsable)() {
    err := m.GetBackingStore().Set("teacherCertificateSettings", value)
    if err != nil {
        panic(err)
    }
}
// IosEduDeviceConfigurationable 
type IosEduDeviceConfigurationable interface {
    DeviceConfigurationable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetDeviceCertificateSettings()(IosEduCertificateSettingsable)
    GetStudentCertificateSettings()(IosEduCertificateSettingsable)
    GetTeacherCertificateSettings()(IosEduCertificateSettingsable)
    SetDeviceCertificateSettings(value IosEduCertificateSettingsable)()
    SetStudentCertificateSettings(value IosEduCertificateSettingsable)()
    SetTeacherCertificateSettings(value IosEduCertificateSettingsable)()
}
