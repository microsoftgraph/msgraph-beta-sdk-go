package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// 
type SynchronizationTemplate struct {
    Entity
    // Identifier of the application this template belongs to.
    applicationId *string;
    // true if this template is recommended to be the default for the application.
    default_escaped *bool;
    // Description of the template.
    description *string;
    // true if this template should appear in the collection of templates available for the application instance (service principal).
    discoverable *bool;
    // One of the well-known factory tags supported by the synchronization engine. The factoryTag tells the synchronization engine which implementation to use when processing jobs based on this template.
    factoryTag *string;
    // Additional extension properties. Unless mentioned explicitly, metadata values should not be changed.
    metadata []MetadataEntry;
    // Default synchronization schema for the jobs based on this template.
    schema *SynchronizationSchema;
}
// Instantiates a new synchronizationTemplate and sets the default values.
func NewSynchronizationTemplate()(*SynchronizationTemplate) {
    m := &SynchronizationTemplate{
        Entity: *NewEntity(),
    }
    return m
}
// Gets the applicationId property value. Identifier of the application this template belongs to.
func (m *SynchronizationTemplate) GetApplicationId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.applicationId
    }
}
// Gets the default_escaped property value. true if this template is recommended to be the default for the application.
func (m *SynchronizationTemplate) GetDefault_escaped()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.default_escaped
    }
}
// Gets the description property value. Description of the template.
func (m *SynchronizationTemplate) GetDescription()(*string) {
    if m == nil {
        return nil
    } else {
        return m.description
    }
}
// Gets the discoverable property value. true if this template should appear in the collection of templates available for the application instance (service principal).
func (m *SynchronizationTemplate) GetDiscoverable()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.discoverable
    }
}
// Gets the factoryTag property value. One of the well-known factory tags supported by the synchronization engine. The factoryTag tells the synchronization engine which implementation to use when processing jobs based on this template.
func (m *SynchronizationTemplate) GetFactoryTag()(*string) {
    if m == nil {
        return nil
    } else {
        return m.factoryTag
    }
}
// Gets the metadata property value. Additional extension properties. Unless mentioned explicitly, metadata values should not be changed.
func (m *SynchronizationTemplate) GetMetadata()([]MetadataEntry) {
    if m == nil {
        return nil
    } else {
        return m.metadata
    }
}
// Gets the schema property value. Default synchronization schema for the jobs based on this template.
func (m *SynchronizationTemplate) GetSchema()(*SynchronizationSchema) {
    if m == nil {
        return nil
    } else {
        return m.schema
    }
}
// The deserialization information for the current model
func (m *SynchronizationTemplate) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["applicationId"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetApplicationId(val)
        }
        return nil
    }
    res["default_escaped"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDefault_escaped(val)
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
    res["discoverable"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDiscoverable(val)
        }
        return nil
    }
    res["factoryTag"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetFactoryTag(val)
        }
        return nil
    }
    res["metadata"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewMetadataEntry() })
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]MetadataEntry, len(val))
            for i, v := range val {
                res[i] = *(v.(*MetadataEntry))
            }
            m.SetMetadata(res)
        }
        return nil
    }
    res["schema"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewSynchronizationSchema() })
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSchema(val.(*SynchronizationSchema))
        }
        return nil
    }
    return res
}
func (m *SynchronizationTemplate) IsNil()(bool) {
    return m == nil
}
// Serializes information the current object
// Parameters:
//  - writer : Serialization writer to use to serialize this model
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
        err = writer.WriteBoolValue("default_escaped", m.GetDefault_escaped())
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
// Sets the applicationId property value. Identifier of the application this template belongs to.
// Parameters:
//  - value : Value to set for the applicationId property.
func (m *SynchronizationTemplate) SetApplicationId(value *string)() {
    m.applicationId = value
}
// Sets the default_escaped property value. true if this template is recommended to be the default for the application.
// Parameters:
//  - value : Value to set for the default_escaped property.
func (m *SynchronizationTemplate) SetDefault_escaped(value *bool)() {
    m.default_escaped = value
}
// Sets the description property value. Description of the template.
// Parameters:
//  - value : Value to set for the description property.
func (m *SynchronizationTemplate) SetDescription(value *string)() {
    m.description = value
}
// Sets the discoverable property value. true if this template should appear in the collection of templates available for the application instance (service principal).
// Parameters:
//  - value : Value to set for the discoverable property.
func (m *SynchronizationTemplate) SetDiscoverable(value *bool)() {
    m.discoverable = value
}
// Sets the factoryTag property value. One of the well-known factory tags supported by the synchronization engine. The factoryTag tells the synchronization engine which implementation to use when processing jobs based on this template.
// Parameters:
//  - value : Value to set for the factoryTag property.
func (m *SynchronizationTemplate) SetFactoryTag(value *string)() {
    m.factoryTag = value
}
// Sets the metadata property value. Additional extension properties. Unless mentioned explicitly, metadata values should not be changed.
// Parameters:
//  - value : Value to set for the metadata property.
func (m *SynchronizationTemplate) SetMetadata(value []MetadataEntry)() {
    m.metadata = value
}
// Sets the schema property value. Default synchronization schema for the jobs based on this template.
// Parameters:
//  - value : Value to set for the schema property.
func (m *SynchronizationTemplate) SetSchema(value *SynchronizationSchema)() {
    m.schema = value
}
