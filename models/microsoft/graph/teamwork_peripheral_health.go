package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// TeamworkPeripheralHealth 
type TeamworkPeripheralHealth struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // The connected state and time since the peripheral device was connected.
    connection *TeamworkConnection;
    // True if the peripheral is optional. Used for health computation.
    isOptional *bool;
    // 
    peripheral *TeamworkPeripheral;
}
// NewTeamworkPeripheralHealth instantiates a new teamworkPeripheralHealth and sets the default values.
func NewTeamworkPeripheralHealth()(*TeamworkPeripheralHealth) {
    m := &TeamworkPeripheralHealth{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *TeamworkPeripheralHealth) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetConnection gets the connection property value. The connected state and time since the peripheral device was connected.
func (m *TeamworkPeripheralHealth) GetConnection()(*TeamworkConnection) {
    if m == nil {
        return nil
    } else {
        return m.connection
    }
}
// GetIsOptional gets the isOptional property value. True if the peripheral is optional. Used for health computation.
func (m *TeamworkPeripheralHealth) GetIsOptional()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.isOptional
    }
}
// GetPeripheral gets the peripheral property value. 
func (m *TeamworkPeripheralHealth) GetPeripheral()(*TeamworkPeripheral) {
    if m == nil {
        return nil
    } else {
        return m.peripheral
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *TeamworkPeripheralHealth) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["connection"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewTeamworkConnection() })
        if err != nil {
            return err
        }
        if val != nil {
            m.SetConnection(val.(*TeamworkConnection))
        }
        return nil
    }
    res["isOptional"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsOptional(val)
        }
        return nil
    }
    res["peripheral"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewTeamworkPeripheral() })
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPeripheral(val.(*TeamworkPeripheral))
        }
        return nil
    }
    return res
}
func (m *TeamworkPeripheralHealth) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
func (m *TeamworkPeripheralHealth) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    {
        err := writer.WriteObjectValue("connection", m.GetConnection())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("isOptional", m.GetIsOptional())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("peripheral", m.GetPeripheral())
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
func (m *TeamworkPeripheralHealth) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
// SetConnection sets the connection property value. The connected state and time since the peripheral device was connected.
func (m *TeamworkPeripheralHealth) SetConnection(value *TeamworkConnection)() {
    if m != nil {
        m.connection = value
    }
}
// SetIsOptional sets the isOptional property value. True if the peripheral is optional. Used for health computation.
func (m *TeamworkPeripheralHealth) SetIsOptional(value *bool)() {
    if m != nil {
        m.isOptional = value
    }
}
// SetPeripheral sets the peripheral property value. 
func (m *TeamworkPeripheralHealth) SetPeripheral(value *TeamworkPeripheral)() {
    if m != nil {
        m.peripheral = value
    }
}
