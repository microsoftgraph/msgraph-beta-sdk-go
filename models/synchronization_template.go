package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// SynchronizationTemplate provides operations to manage the collection of accessReviewDecision entities.
type SynchronizationTemplate struct {
    Entity
    // Identifier of the application this template belongs to.
    applicationId *string
    // true if this template is recommended to be the default for the application.
    default_escaped *bool
    // Description of the template.
    description *string
    // true if this template should appear in the collection of templates available for the application instance (service principal).
    discoverable *bool
    // One of the well-known factory tags supported by the synchronization engine. The factoryTag tells the synchronization engine which implementation to use when processing jobs based on this template.
    factoryTag *string
    // Additional extension properties. Unless mentioned explicitly, metadata values should not be changed.
    metadata []MetadataEntryable
    // Default synchronization schema for the jobs based on this template.
    schema SynchronizationSchemaable
}
// NewSynchronizationTemplate instantiates a new synchronizationTemplate and sets the default values.
func NewSynchronizationTemplate()(*SynchronizationTemplate) {
    m := &SynchronizationTemplate{
        Entity: *NewEntity(),
    }
    odataTypeValue := "#microsoft.graph.synchronizationTemplate";
    m.SetOdataType(&odataTypeValue);
    return m
}
// CreateSynchronizationTemplateFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateSynchronizationTemplateFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewSynchronizationTemplate(), nil
}
// GetApplicationId gets the applicationId property value. Identifier of the application this template belongs to.
func (m *SynchronizationTemplate) GetApplicationId()(*string) {
    return m.applicationId
}
// GetDefault gets the default property value. true if this template is recommended to be the default for the application.
func (m *SynchronizationTemplate) GetDefault()(*bool) {
    return m.default_escaped
}
// GetDescription gets the description property value. Description of the template.
func (m *SynchronizationTemplate) GetDescription()(*string) {
    return m.description
}
// GetDiscoverable gets the discoverable property value. true if this template should appear in the collection of templates available for the application instance (service principal).
func (m *SynchronizationTemplate) GetDiscoverable()(*bool) {
    return m.discoverable
}
// GetFactoryTag gets the factoryTag property value. One of the well-known factory tags supported by the synchronization engine. The factoryTag tells the synchronization engine which implementation to use when processing jobs based on this template.
func (m *SynchronizationTemplate) GetFactoryTag()(*string) {
    return m.factoryTag
}
// GetFieldDeserializers the deserialization information for the current model
func (m *SynchronizationTemplate) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["applicationId"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetApplicationId)
    res["default"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetBoolValue(m.SetDefault)
    res["description"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetDescription)
    res["discoverable"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetBoolValue(m.SetDiscoverable)
    res["factoryTag"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetFactoryTag)
    res["metadata"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfObjectValues(CreateMetadataEntryFromDiscriminatorValue , m.SetMetadata)
    res["schema"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetObjectValue(CreateSynchronizationSchemaFromDiscriminatorValue , m.SetSchema)
    return res
}
// GetMetadata gets the metadata property value. Additional extension properties. Unless mentioned explicitly, metadata values should not be changed.
func (m *SynchronizationTemplate) GetMetadata()([]MetadataEntryable) {
    return m.metadata
}
// GetSchema gets the schema property value. Default synchronization schema for the jobs based on this template.
func (m *SynchronizationTemplate) GetSchema()(SynchronizationSchemaable) {
    return m.schema
}
// Serialize serializes information the current object
func (m *SynchronizationTemplate) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
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
        err = writer.WriteBoolValue("default", m.GetDefault())
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
    if m.GetMetadata() != nil {
        cast := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.CollectionCast[i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable](m.GetMetadata())
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
// SetApplicationId sets the applicationId property value. Identifier of the application this template belongs to.
func (m *SynchronizationTemplate) SetApplicationId(value *string)() {
    m.applicationId = value
}
// SetDefault sets the default property value. true if this template is recommended to be the default for the application.
func (m *SynchronizationTemplate) SetDefault(value *bool)() {
    m.default_escaped = value
}
// SetDescription sets the description property value. Description of the template.
func (m *SynchronizationTemplate) SetDescription(value *string)() {
    m.description = value
}
// SetDiscoverable sets the discoverable property value. true if this template should appear in the collection of templates available for the application instance (service principal).
func (m *SynchronizationTemplate) SetDiscoverable(value *bool)() {
    m.discoverable = value
}
// SetFactoryTag sets the factoryTag property value. One of the well-known factory tags supported by the synchronization engine. The factoryTag tells the synchronization engine which implementation to use when processing jobs based on this template.
func (m *SynchronizationTemplate) SetFactoryTag(value *string)() {
    m.factoryTag = value
}
// SetMetadata sets the metadata property value. Additional extension properties. Unless mentioned explicitly, metadata values should not be changed.
func (m *SynchronizationTemplate) SetMetadata(value []MetadataEntryable)() {
    m.metadata = value
}
// SetSchema sets the schema property value. Default synchronization schema for the jobs based on this template.
func (m *SynchronizationTemplate) SetSchema(value SynchronizationSchemaable)() {
    m.schema = value
}
