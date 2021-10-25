package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

type ReferencedObject struct {
    additionalData map[string]interface{};
    referencedObjectName *string;
    referencedProperty *string;
}
func NewReferencedObject()(*ReferencedObject) {
    m := &ReferencedObject{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
func (m *ReferencedObject) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
func (m *ReferencedObject) GetReferencedObjectName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.referencedObjectName
    }
}
func (m *ReferencedObject) GetReferencedProperty()(*string) {
    if m == nil {
        return nil
    } else {
        return m.referencedProperty
    }
}
func (m *ReferencedObject) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["referencedObjectName"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetReferencedObjectName(val)
        return nil
    }
    res["referencedProperty"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetReferencedProperty(val)
        return nil
    }
    return res
}
func (m *ReferencedObject) IsNil()(bool) {
    return m == nil
}
func (m *ReferencedObject) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("referencedObjectName", m.GetReferencedObjectName())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("referencedProperty", m.GetReferencedProperty())
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
func (m *ReferencedObject) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
func (m *ReferencedObject) SetReferencedObjectName(value *string)() {
    m.referencedObjectName = value
}
func (m *ReferencedObject) SetReferencedProperty(value *string)() {
    m.referencedProperty = value
}
