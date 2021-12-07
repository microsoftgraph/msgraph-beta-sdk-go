package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// DeviceKey 
type DeviceKey struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // 
    deviceId *string;
    // 
    keyMaterial []byte;
    // 
    keyType *string;
}
// NewDeviceKey instantiates a new deviceKey and sets the default values.
func NewDeviceKey()(*DeviceKey) {
    m := &DeviceKey{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *DeviceKey) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetDeviceId gets the deviceId property value. 
func (m *DeviceKey) GetDeviceId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.deviceId
    }
}
// GetKeyMaterial gets the keyMaterial property value. 
func (m *DeviceKey) GetKeyMaterial()([]byte) {
    if m == nil {
        return nil
    } else {
        return m.keyMaterial
    }
}
// GetKeyType gets the keyType property value. 
func (m *DeviceKey) GetKeyType()(*string) {
    if m == nil {
        return nil
    } else {
        return m.keyType
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *DeviceKey) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["deviceId"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDeviceId(val)
        }
        return nil
    }
    res["keyMaterial"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetByteArrayValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetKeyMaterial(val)
        }
        return nil
    }
    res["keyType"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
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
func (m *DeviceKey) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
func (m *DeviceKey) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
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
// SetDeviceId sets the deviceId property value. 
func (m *DeviceKey) SetDeviceId(value *string)() {
    if m != nil {
        m.deviceId = value
    }
}
// SetKeyMaterial sets the keyMaterial property value. 
func (m *DeviceKey) SetKeyMaterial(value []byte)() {
    if m != nil {
        m.keyMaterial = value
    }
}
// SetKeyType sets the keyType property value. 
func (m *DeviceKey) SetKeyType(value *string)() {
    if m != nil {
        m.keyType = value
    }
}
