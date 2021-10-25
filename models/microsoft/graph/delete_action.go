package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

type DeleteAction struct {
    additionalData map[string]interface{};
    name *string;
    objectType *string;
}
func NewDeleteAction()(*DeleteAction) {
    m := &DeleteAction{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
func (m *DeleteAction) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
func (m *DeleteAction) GetName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.name
    }
}
func (m *DeleteAction) GetObjectType()(*string) {
    if m == nil {
        return nil
    } else {
        return m.objectType
    }
}
func (m *DeleteAction) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["name"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetName(val)
        return nil
    }
    res["objectType"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetObjectType(val)
        return nil
    }
    return res
}
func (m *DeleteAction) IsNil()(bool) {
    return m == nil
}
func (m *DeleteAction) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("name", m.GetName())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("objectType", m.GetObjectType())
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
func (m *DeleteAction) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
func (m *DeleteAction) SetName(value *string)() {
    m.name = value
}
func (m *DeleteAction) SetObjectType(value *string)() {
    m.objectType = value
}
