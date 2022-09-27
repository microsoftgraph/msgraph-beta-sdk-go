package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// DeviceKey 
type DeviceKey struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // The deviceId property
    deviceId *string
    // The keyMaterial property
    keyMaterial []byte
    // The keyType property
    keyType *string
    // The OdataType property
    odataType *string
}
// NewDeviceKey instantiates a new deviceKey and sets the default values.
func NewDeviceKey()(*DeviceKey) {
    m := &DeviceKey{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    odataTypeValue := "#microsoft.graph.deviceKey";
    m.SetOdataType(&odataTypeValue);
    return m
}
// CreateDeviceKeyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateDeviceKeyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewDeviceKey(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *DeviceKey) GetAdditionalData()(map[string]interface{}) {
    return m.additionalData
}
// GetDeviceId gets the deviceId property value. The deviceId property
func (m *DeviceKey) GetDeviceId()(*string) {
    return m.deviceId
}
// GetFieldDeserializers the deserialization information for the current model
func (m *DeviceKey) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["deviceId"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetDeviceId)
    res["keyMaterial"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetByteArrayValue(m.SetKeyMaterial)
    res["keyType"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetKeyType)
    res["@odata.type"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetOdataType)
    return res
}
// GetKeyMaterial gets the keyMaterial property value. The keyMaterial property
func (m *DeviceKey) GetKeyMaterial()([]byte) {
    return m.keyMaterial
}
// GetKeyType gets the keyType property value. The keyType property
func (m *DeviceKey) GetKeyType()(*string) {
    return m.keyType
}
// GetOdataType gets the @odata.type property value. The OdataType property
func (m *DeviceKey) GetOdataType()(*string) {
    return m.odataType
}
// Serialize serializes information the current object
func (m *DeviceKey) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("deviceId", m.GetDeviceId())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteByteArrayValue("keyMaterial", m.GetKeyMaterial())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("keyType", m.GetKeyType())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("@odata.type", m.GetOdataType())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteAdditionalData(m.GetAdditionalData())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAdditionalData sets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *DeviceKey) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// SetDeviceId sets the deviceId property value. The deviceId property
func (m *DeviceKey) SetDeviceId(value *string)() {
    m.deviceId = value
}
// SetKeyMaterial sets the keyMaterial property value. The keyMaterial property
func (m *DeviceKey) SetKeyMaterial(value []byte)() {
    m.keyMaterial = value
}
// SetKeyType sets the keyType property value. The keyType property
func (m *DeviceKey) SetKeyType(value *string)() {
    m.keyType = value
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *DeviceKey) SetOdataType(value *string)() {
    m.odataType = value
}
