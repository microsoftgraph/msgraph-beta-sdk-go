package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// IosEduDeviceConfiguration 
type IosEduDeviceConfiguration struct {
    DeviceConfiguration
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // The Trusted Root and PFX certificates for Device
    deviceCertificateSettings IosEduCertificateSettingsable
    // The Trusted Root and PFX certificates for Student
    studentCertificateSettings IosEduCertificateSettingsable
    // Trusted Root and PFX certificates for iOS EDU.
    teacherCertificateSettings IosEduCertificateSettingsable
}
// NewIosEduDeviceConfiguration instantiates a new IosEduDeviceConfiguration and sets the default values.
func NewIosEduDeviceConfiguration()(*IosEduDeviceConfiguration) {
    m := &IosEduDeviceConfiguration{
        DeviceConfiguration: *NewDeviceConfiguration(),
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateIosEduDeviceConfigurationFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateIosEduDeviceConfigurationFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewIosEduDeviceConfiguration(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *IosEduDeviceConfiguration) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetDeviceCertificateSettings gets the deviceCertificateSettings property value. The Trusted Root and PFX certificates for Device
func (m *IosEduDeviceConfiguration) GetDeviceCertificateSettings()(IosEduCertificateSettingsable) {
    if m == nil {
        return nil
    } else {
        return m.deviceCertificateSettings
    }
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
    if m == nil {
        return nil
    } else {
        return m.studentCertificateSettings
    }
}
// GetTeacherCertificateSettings gets the teacherCertificateSettings property value. Trusted Root and PFX certificates for iOS EDU.
func (m *IosEduDeviceConfiguration) GetTeacherCertificateSettings()(IosEduCertificateSettingsable) {
    if m == nil {
        return nil
    } else {
        return m.teacherCertificateSettings
    }
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
    {
        err = writer.WriteAdditionalData(m.GetAdditionalData())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAdditionalData sets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *IosEduDeviceConfiguration) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
// SetDeviceCertificateSettings sets the deviceCertificateSettings property value. The Trusted Root and PFX certificates for Device
func (m *IosEduDeviceConfiguration) SetDeviceCertificateSettings(value IosEduCertificateSettingsable)() {
    if m != nil {
        m.deviceCertificateSettings = value
    }
}
// SetStudentCertificateSettings sets the studentCertificateSettings property value. The Trusted Root and PFX certificates for Student
func (m *IosEduDeviceConfiguration) SetStudentCertificateSettings(value IosEduCertificateSettingsable)() {
    if m != nil {
        m.studentCertificateSettings = value
    }
}
// SetTeacherCertificateSettings sets the teacherCertificateSettings property value. Trusted Root and PFX certificates for iOS EDU.
func (m *IosEduDeviceConfiguration) SetTeacherCertificateSettings(value IosEduCertificateSettingsable)() {
    if m != nil {
        m.teacherCertificateSettings = value
    }
}
