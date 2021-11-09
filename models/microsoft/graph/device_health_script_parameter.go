package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// 
type DeviceHealthScriptParameter struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // Whether Apply DefaultValue When Not Assigned
    applyDefaultValueWhenNotAssigned *bool;
    // The description of the param
    description *string;
    // Whether the param is required
    isRequired *bool;
    // The name of the param
    name *string;
}
// Instantiates a new deviceHealthScriptParameter and sets the default values.
func NewDeviceHealthScriptParameter()(*DeviceHealthScriptParameter) {
    m := &DeviceHealthScriptParameter{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// Gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *DeviceHealthScriptParameter) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// Gets the applyDefaultValueWhenNotAssigned property value. Whether Apply DefaultValue When Not Assigned
func (m *DeviceHealthScriptParameter) GetApplyDefaultValueWhenNotAssigned()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.applyDefaultValueWhenNotAssigned
    }
}
// Gets the description property value. The description of the param
func (m *DeviceHealthScriptParameter) GetDescription()(*string) {
    if m == nil {
        return nil
    } else {
        return m.description
    }
}
// Gets the isRequired property value. Whether the param is required
func (m *DeviceHealthScriptParameter) GetIsRequired()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.isRequired
    }
}
// Gets the name property value. The name of the param
func (m *DeviceHealthScriptParameter) GetName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.name
    }
}
// The deserialization information for the current model
func (m *DeviceHealthScriptParameter) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["applyDefaultValueWhenNotAssigned"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetApplyDefaultValueWhenNotAssigned(val)
        }
        return nil
    }
    res["description"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDescription(val)
        }
        return nil
    }
    res["isRequired"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsRequired(val)
        }
        return nil
    }
    res["name"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetName(val)
        }
        return nil
    }
    return res
}
func (m *DeviceHealthScriptParameter) IsNil()(bool) {
    return m == nil
}
// Serializes information the current object
// Parameters:
//  - writer : Serialization writer to use to serialize this model
func (m *DeviceHealthScriptParameter) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    {
        err := writer.WriteBoolValue("applyDefaultValueWhenNotAssigned", m.GetApplyDefaultValueWhenNotAssigned())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("description", m.GetDescription())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("isRequired", m.GetIsRequired())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("name", m.GetName())
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
func (m *DeviceHealthScriptParameter) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// Sets the applyDefaultValueWhenNotAssigned property value. Whether Apply DefaultValue When Not Assigned
// Parameters:
//  - value : Value to set for the applyDefaultValueWhenNotAssigned property.
func (m *DeviceHealthScriptParameter) SetApplyDefaultValueWhenNotAssigned(value *bool)() {
    m.applyDefaultValueWhenNotAssigned = value
}
// Sets the description property value. The description of the param
// Parameters:
//  - value : Value to set for the description property.
func (m *DeviceHealthScriptParameter) SetDescription(value *string)() {
    m.description = value
}
// Sets the isRequired property value. Whether the param is required
// Parameters:
//  - value : Value to set for the isRequired property.
func (m *DeviceHealthScriptParameter) SetIsRequired(value *bool)() {
    m.isRequired = value
}
// Sets the name property value. The name of the param
// Parameters:
//  - value : Value to set for the name property.
func (m *DeviceHealthScriptParameter) SetName(value *string)() {
    m.name = value
}
