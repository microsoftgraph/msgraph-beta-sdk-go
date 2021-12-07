package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// DeviceHealthScriptParameter 
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
// NewDeviceHealthScriptParameter instantiates a new deviceHealthScriptParameter and sets the default values.
func NewDeviceHealthScriptParameter()(*DeviceHealthScriptParameter) {
    m := &DeviceHealthScriptParameter{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *DeviceHealthScriptParameter) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetApplyDefaultValueWhenNotAssigned gets the applyDefaultValueWhenNotAssigned property value. Whether Apply DefaultValue When Not Assigned
func (m *DeviceHealthScriptParameter) GetApplyDefaultValueWhenNotAssigned()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.applyDefaultValueWhenNotAssigned
    }
}
// GetDescription gets the description property value. The description of the param
func (m *DeviceHealthScriptParameter) GetDescription()(*string) {
    if m == nil {
        return nil
    } else {
        return m.description
    }
}
// GetIsRequired gets the isRequired property value. Whether the param is required
func (m *DeviceHealthScriptParameter) GetIsRequired()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.isRequired
    }
}
// GetName gets the name property value. The name of the param
func (m *DeviceHealthScriptParameter) GetName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.name
    }
}
// GetFieldDeserializers the deserialization information for the current model
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
// Serialize serializes information the current object
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
// SetAdditionalData sets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *DeviceHealthScriptParameter) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
// SetApplyDefaultValueWhenNotAssigned sets the applyDefaultValueWhenNotAssigned property value. Whether Apply DefaultValue When Not Assigned
func (m *DeviceHealthScriptParameter) SetApplyDefaultValueWhenNotAssigned(value *bool)() {
    if m != nil {
        m.applyDefaultValueWhenNotAssigned = value
    }
}
// SetDescription sets the description property value. The description of the param
func (m *DeviceHealthScriptParameter) SetDescription(value *string)() {
    if m != nil {
        m.description = value
    }
}
// SetIsRequired sets the isRequired property value. Whether the param is required
func (m *DeviceHealthScriptParameter) SetIsRequired(value *bool)() {
    if m != nil {
        m.isRequired = value
    }
}
// SetName sets the name property value. The name of the param
func (m *DeviceHealthScriptParameter) SetName(value *string)() {
    if m != nil {
        m.name = value
    }
}
