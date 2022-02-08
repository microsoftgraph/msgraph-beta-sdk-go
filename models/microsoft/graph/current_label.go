package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// CurrentLabel 
type CurrentLabel struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // 
    applicationMode *ApplicationMode;
    // 
    id *string;
}
// NewCurrentLabel instantiates a new currentLabel and sets the default values.
func NewCurrentLabel()(*CurrentLabel) {
    m := &CurrentLabel{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *CurrentLabel) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetApplicationMode gets the applicationMode property value. 
func (m *CurrentLabel) GetApplicationMode()(*ApplicationMode) {
    if m == nil {
        return nil
    } else {
        return m.applicationMode
    }
}
// GetId gets the id property value. 
func (m *CurrentLabel) GetId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.id
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *CurrentLabel) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["applicationMode"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParseApplicationMode)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetApplicationMode(val.(*ApplicationMode))
        }
        return nil
    }
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
    return res
}
func (m *CurrentLabel) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
func (m *CurrentLabel) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    if m.GetApplicationMode() != nil {
        cast := (*m.GetApplicationMode()).String()
        err := writer.WriteStringValue("applicationMode", &cast)
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("id", m.GetId())
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
func (m *CurrentLabel) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
// SetApplicationMode sets the applicationMode property value. 
func (m *CurrentLabel) SetApplicationMode(value *ApplicationMode)() {
    if m != nil {
        m.applicationMode = value
    }
}
// SetId sets the id property value. 
func (m *CurrentLabel) SetId(value *string)() {
    if m != nil {
        m.id = value
    }
}
