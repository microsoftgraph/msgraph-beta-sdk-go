package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// ResponsiblePolicy 
type ResponsiblePolicy struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // 
    id *string;
    // 
    name *string;
}
// NewResponsiblePolicy instantiates a new responsiblePolicy and sets the default values.
func NewResponsiblePolicy()(*ResponsiblePolicy) {
    m := &ResponsiblePolicy{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *ResponsiblePolicy) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetId gets the id property value. 
func (m *ResponsiblePolicy) GetId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.id
    }
}
// GetName gets the name property value. 
func (m *ResponsiblePolicy) GetName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.name
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ResponsiblePolicy) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["id"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetId(val)
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
func (m *ResponsiblePolicy) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
func (m *ResponsiblePolicy) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("id", m.GetId())
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
func (m *ResponsiblePolicy) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
// SetId sets the id property value. 
func (m *ResponsiblePolicy) SetId(value *string)() {
    if m != nil {
        m.id = value
    }
}
// SetName sets the name property value. 
func (m *ResponsiblePolicy) SetName(value *string)() {
    if m != nil {
        m.name = value
    }
}
