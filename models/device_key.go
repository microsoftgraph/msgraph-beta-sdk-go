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
}
// NewDeviceKey instantiates a new deviceKey and sets the default values.
func NewDeviceKey()(*DeviceKey) {
    m := &DeviceKey{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateDeviceKeyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateDeviceKeyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewDeviceKey(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *DeviceKey) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetDeviceId gets the deviceId property value. The deviceId property
func (m *DeviceKey) GetDeviceId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.deviceId
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *DeviceKey) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["deviceId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDeviceId(val)
        }
        return nil
    }
    res["keyMaterial"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetByteArrayValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetKeyMaterial(val)
        }
        return nil
    }
    res["keyType"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetKeyType(val)
        }
        return nil
    }
    return res
}
// GetKeyMaterial gets the keyMaterial property value. The keyMaterial property
func (m *DeviceKey) GetKeyMaterial()([]byte) {
    if m == nil {
        return nil
    } else {
        return m.keyMaterial
    }
}
// GetKeyType gets the keyType property value. The keyType property
func (m *DeviceKey) GetKeyType()(*string) {
    if m == nil {
        return nil
    } else {
        return m.keyType
    }
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
        err := writer.WriteAdditionalData(m.GetAdditionalData())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAdditionalData sets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *DeviceKey) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
// SetDeviceId sets the deviceId property value. The deviceId property
func (m *DeviceKey) SetDeviceId(value *string)() {
    if m != nil {
        m.deviceId = value
    }
}
// SetKeyMaterial sets the keyMaterial property value. The keyMaterial property
func (m *DeviceKey) SetKeyMaterial(value []byte)() {
    if m != nil {
        m.keyMaterial = value
    }
}
// SetKeyType sets the keyType property value. The keyType property
func (m *DeviceKey) SetKeyType(value *string)() {
    if m != nil {
        m.keyType = value
    }
}
