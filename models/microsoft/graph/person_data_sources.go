package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

type PersonDataSources struct {
    additionalData map[string]interface{};
    type_escaped []string;
}
func NewPersonDataSources()(*PersonDataSources) {
    m := &PersonDataSources{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
func (m *PersonDataSources) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
func (m *PersonDataSources) GetType_escaped()([]string) {
    if m == nil {
        return nil
    } else {
        return m.type_escaped
    }
}
func (m *PersonDataSources) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["type_escaped"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        res := make([]string, len(val))
        for i, v := range val {
            res[i] = v.(string)
        }
        m.SetType_escaped(res)
        return nil
    }
    return res
}
func (m *PersonDataSources) IsNil()(bool) {
    return m == nil
}
func (m *PersonDataSources) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    {
        err := writer.WriteCollectionOfStringValues("type_escaped", m.GetType_escaped())
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
func (m *PersonDataSources) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
func (m *PersonDataSources) SetType_escaped(value []string)() {
    m.type_escaped = value
}
