package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// 
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
// Instantiates a new deviceKey and sets the default values.
func NewDeviceKey()(*DeviceKey) {
    m := &DeviceKey{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// Gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *DeviceKey) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// Gets the deviceId property value. 
func (m *DeviceKey) GetDeviceId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.deviceId
    }
}
// Gets the keyMaterial property value. 
func (m *DeviceKey) GetKeyMaterial()([]byte) {
    if m == nil {
        return nil
    } else {
        return m.keyMaterial
    }
}
// Gets the keyType property value. 
func (m *DeviceKey) GetKeyType()(*string) {
    if m == nil {
        return nil
    } else {
        return m.keyType
    }
}
// The deserialization information for the current model
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
// Serializes information the current object
// Parameters:
//  - writer : Serialization writer to use to serialize this model
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
// Sets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// Parameters:
//  - value : Value to set for the AdditionalData property.
func (m *DeviceKey) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// Sets the deviceId property value. 
// Parameters:
//  - value : Value to set for the deviceId property.
func (m *DeviceKey) SetDeviceId(value *string)() {
    m.deviceId = value
}
// Sets the keyMaterial property value. 
// Parameters:
//  - value : Value to set for the keyMaterial property.
func (m *DeviceKey) SetKeyMaterial(value []byte)() {
    m.keyMaterial = value
}
// Sets the keyType property value. 
// Parameters:
//  - value : Value to set for the keyType property.
func (m *DeviceKey) SetKeyType(value *string)() {
    m.keyType = value
}
