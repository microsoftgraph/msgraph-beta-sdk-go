package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// AndroidOmaCpConfiguration 
type AndroidOmaCpConfiguration struct {
    DeviceConfiguration
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // Configuration XML that will be applied to the device. When it is read, it only provides a placeholder string since the original data is encrypted and stored.
    configurationXml []byte
}
// NewAndroidOmaCpConfiguration instantiates a new AndroidOmaCpConfiguration and sets the default values.
func NewAndroidOmaCpConfiguration()(*AndroidOmaCpConfiguration) {
    m := &AndroidOmaCpConfiguration{
        DeviceConfiguration: *NewDeviceConfiguration(),
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateAndroidOmaCpConfigurationFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateAndroidOmaCpConfigurationFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewAndroidOmaCpConfiguration(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *AndroidOmaCpConfiguration) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetConfigurationXml gets the configurationXml property value. Configuration XML that will be applied to the device. When it is read, it only provides a placeholder string since the original data is encrypted and stored.
func (m *AndroidOmaCpConfiguration) GetConfigurationXml()([]byte) {
    if m == nil {
        return nil
    } else {
        return m.configurationXml
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *AndroidOmaCpConfiguration) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.DeviceConfiguration.GetFieldDeserializers()
    res["configurationXml"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetByteArrayValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetConfigurationXml(val)
        }
        return nil
    }
    return res
}
// Serialize serializes information the current object
func (m *AndroidOmaCpConfiguration) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.DeviceConfiguration.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteByteArrayValue("configurationXml", m.GetConfigurationXml())
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
func (m *AndroidOmaCpConfiguration) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
// SetConfigurationXml sets the configurationXml property value. Configuration XML that will be applied to the device. When it is read, it only provides a placeholder string since the original data is encrypted and stored.
func (m *AndroidOmaCpConfiguration) SetConfigurationXml(value []byte)() {
    if m != nil {
        m.configurationXml = value
    }
}
