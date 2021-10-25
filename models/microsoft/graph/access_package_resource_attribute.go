package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

type AccessPackageResourceAttribute struct {
    additionalData map[string]interface{};
    attributeDestination *AccessPackageResourceAttributeDestination;
    attributeName *string;
    attributeSource *AccessPackageResourceAttributeSource;
    id *string;
}
func NewAccessPackageResourceAttribute()(*AccessPackageResourceAttribute) {
    m := &AccessPackageResourceAttribute{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
func (m *AccessPackageResourceAttribute) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
func (m *AccessPackageResourceAttribute) GetAttributeDestination()(*AccessPackageResourceAttributeDestination) {
    if m == nil {
        return nil
    } else {
        return m.attributeDestination
    }
}
func (m *AccessPackageResourceAttribute) GetAttributeName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.attributeName
    }
}
func (m *AccessPackageResourceAttribute) GetAttributeSource()(*AccessPackageResourceAttributeSource) {
    if m == nil {
        return nil
    } else {
        return m.attributeSource
    }
}
func (m *AccessPackageResourceAttribute) GetId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.id
    }
}
func (m *AccessPackageResourceAttribute) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["attributeDestination"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewAccessPackageResourceAttributeDestination() })
        if err != nil {
            return err
        }
        m.SetAttributeDestination(val.(*AccessPackageResourceAttributeDestination))
        return nil
    }
    res["attributeName"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetAttributeName(val)
        return nil
    }
    res["attributeSource"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewAccessPackageResourceAttributeSource() })
        if err != nil {
            return err
        }
        m.SetAttributeSource(val.(*AccessPackageResourceAttributeSource))
        return nil
    }
    res["id"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetId(val)
        return nil
    }
    return res
}
func (m *AccessPackageResourceAttribute) IsNil()(bool) {
    return m == nil
}
func (m *AccessPackageResourceAttribute) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    {
        err := writer.WriteObjectValue("attributeDestination", m.GetAttributeDestination())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("attributeName", m.GetAttributeName())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("attributeSource", m.GetAttributeSource())
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
func (m *AccessPackageResourceAttribute) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
func (m *AccessPackageResourceAttribute) SetAttributeDestination(value *AccessPackageResourceAttributeDestination)() {
    m.attributeDestination = value
}
func (m *AccessPackageResourceAttribute) SetAttributeName(value *string)() {
    m.attributeName = value
}
func (m *AccessPackageResourceAttribute) SetAttributeSource(value *AccessPackageResourceAttributeSource)() {
    m.attributeSource = value
}
func (m *AccessPackageResourceAttribute) SetId(value *string)() {
    m.id = value
}
