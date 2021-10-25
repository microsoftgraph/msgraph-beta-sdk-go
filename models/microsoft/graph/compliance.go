package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

type Compliance struct {
    additionalData map[string]interface{};
    ediscovery *Ediscoveryroot;
}
func NewCompliance()(*Compliance) {
    m := &Compliance{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
func (m *Compliance) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
func (m *Compliance) GetEdiscovery()(*Ediscoveryroot) {
    if m == nil {
        return nil
    } else {
        return m.ediscovery
    }
}
func (m *Compliance) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["ediscovery"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewEdiscoveryroot() })
        if err != nil {
            return err
        }
        m.SetEdiscovery(val.(*Ediscoveryroot))
        return nil
    }
    return res
}
func (m *Compliance) IsNil()(bool) {
    return m == nil
}
func (m *Compliance) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    {
        err := writer.WriteObjectValue("ediscovery", m.GetEdiscovery())
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
func (m *Compliance) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
func (m *Compliance) SetEdiscovery(value *Ediscoveryroot)() {
    m.ediscovery = value
}
