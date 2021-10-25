package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

type SynchronizationTemplate struct {
    Entity
    applicationId *string;
    default_escpaped *bool;
    description *string;
    discoverable *bool;
    factoryTag *string;
    metadata []MetadataEntry;
    schema *SynchronizationSchema;
}
func NewSynchronizationTemplate()(*SynchronizationTemplate) {
    m := &SynchronizationTemplate{
        Entity: *NewEntity(),
    }
    return m
}
func (m *SynchronizationTemplate) GetApplicationId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.applicationId
    }
}
func (m *SynchronizationTemplate) GetDefault_escpaped()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.default_escpaped
    }
}
func (m *SynchronizationTemplate) GetDescription()(*string) {
    if m == nil {
        return nil
    } else {
        return m.description
    }
}
func (m *SynchronizationTemplate) GetDiscoverable()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.discoverable
    }
}
func (m *SynchronizationTemplate) GetFactoryTag()(*string) {
    if m == nil {
        return nil
    } else {
        return m.factoryTag
    }
}
func (m *SynchronizationTemplate) GetMetadata()([]MetadataEntry) {
    if m == nil {
        return nil
    } else {
        return m.metadata
    }
}
func (m *SynchronizationTemplate) GetSchema()(*SynchronizationSchema) {
    if m == nil {
        return nil
    } else {
        return m.schema
    }
}
func (m *SynchronizationTemplate) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["applicationId"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetApplicationId(val)
        return nil
    }
    res["default_escpaped"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        m.SetDefault_escpaped(val)
        return nil
    }
    res["description"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetDescription(val)
        return nil
    }
    res["discoverable"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        m.SetDiscoverable(val)
        return nil
    }
    res["factoryTag"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetFactoryTag(val)
        return nil
    }
    res["metadata"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewMetadataEntry() })
        if err != nil {
            return err
        }
        res := make([]MetadataEntry, len(val))
        for i, v := range val {
            res[i] = *(v.(*MetadataEntry))
        }
        m.SetMetadata(res)
        return nil
    }
    res["schema"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewSynchronizationSchema() })
        if err != nil {
            return err
        }
        m.SetSchema(val.(*SynchronizationSchema))
        return nil
    }
    return res
}
func (m *SynchronizationTemplate) IsNil()(bool) {
    return m == nil
}
func (m *SynchronizationTemplate) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteStringValue("applicationId", m.GetApplicationId())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("default_escpaped", m.GetDefault_escpaped())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("description", m.GetDescription())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("discoverable", m.GetDiscoverable())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("factoryTag", m.GetFactoryTag())
        if err != nil {
            return err
        }
    }
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetMetadata()))
        for i, v := range m.GetMetadata() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("metadata", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("schema", m.GetSchema())
        if err != nil {
            return err
        }
    }
    return nil
}
func (m *SynchronizationTemplate) SetApplicationId(value *string)() {
    m.applicationId = value
}
func (m *SynchronizationTemplate) SetDefault_escpaped(value *bool)() {
    m.default_escpaped = value
}
func (m *SynchronizationTemplate) SetDescription(value *string)() {
    m.description = value
}
func (m *SynchronizationTemplate) SetDiscoverable(value *bool)() {
    m.discoverable = value
}
func (m *SynchronizationTemplate) SetFactoryTag(value *string)() {
    m.factoryTag = value
}
func (m *SynchronizationTemplate) SetMetadata(value []MetadataEntry)() {
    m.metadata = value
}
func (m *SynchronizationTemplate) SetSchema(value *SynchronizationSchema)() {
    m.schema = value
}
