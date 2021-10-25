package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

type AndroidForWorkAppConfigurationSchema struct {
    Entity
    exampleJson []byte;
    schemaItems []AndroidForWorkAppConfigurationSchemaItem;
}
func NewAndroidForWorkAppConfigurationSchema()(*AndroidForWorkAppConfigurationSchema) {
    m := &AndroidForWorkAppConfigurationSchema{
        Entity: *NewEntity(),
    }
    return m
}
func (m *AndroidForWorkAppConfigurationSchema) GetExampleJson()([]byte) {
    if m == nil {
        return nil
    } else {
        return m.exampleJson
    }
}
func (m *AndroidForWorkAppConfigurationSchema) GetSchemaItems()([]AndroidForWorkAppConfigurationSchemaItem) {
    if m == nil {
        return nil
    } else {
        return m.schemaItems
    }
}
func (m *AndroidForWorkAppConfigurationSchema) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["exampleJson"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetByteArrayValue()
        if err != nil {
            return err
        }
        m.SetExampleJson(val)
        return nil
    }
    res["schemaItems"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewAndroidForWorkAppConfigurationSchemaItem() })
        if err != nil {
            return err
        }
        res := make([]AndroidForWorkAppConfigurationSchemaItem, len(val))
        for i, v := range val {
            res[i] = *(v.(*AndroidForWorkAppConfigurationSchemaItem))
        }
        m.SetSchemaItems(res)
        return nil
    }
    return res
}
func (m *AndroidForWorkAppConfigurationSchema) IsNil()(bool) {
    return m == nil
}
func (m *AndroidForWorkAppConfigurationSchema) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteByteArrayValue("exampleJson", m.GetExampleJson())
        if err != nil {
            return err
        }
    }
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetSchemaItems()))
        for i, v := range m.GetSchemaItems() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("schemaItems", cast)
        if err != nil {
            return err
        }
    }
    return nil
}
func (m *AndroidForWorkAppConfigurationSchema) SetExampleJson(value []byte)() {
    m.exampleJson = value
}
func (m *AndroidForWorkAppConfigurationSchema) SetSchemaItems(value []AndroidForWorkAppConfigurationSchemaItem)() {
    m.schemaItems = value
}
