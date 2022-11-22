package models

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// SynchronizationSchema 
type SynchronizationSchema struct {
    Entity
    // Contains the collection of directories and all of their objects.
    directories []DirectoryDefinitionable
    // A collection of synchronization rules configured for the synchronizationJob or synchronizationTemplate.
    synchronizationRules []SynchronizationRuleable
    // The version of the schema, updated automatically with every schema change.
    version *string
}
// NewSynchronizationSchema instantiates a new synchronizationSchema and sets the default values.
func NewSynchronizationSchema()(*SynchronizationSchema) {
    m := &SynchronizationSchema{
        Entity: *NewEntity(),
    }
    return m
}
// CreateSynchronizationSchemaFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateSynchronizationSchemaFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewSynchronizationSchema(), nil
}
// GetDirectories gets the directories property value. Contains the collection of directories and all of their objects.
func (m *SynchronizationSchema) GetDirectories()([]DirectoryDefinitionable) {
    return m.directories
}
// GetFieldDeserializers the deserialization information for the current model
func (m *SynchronizationSchema) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["directories"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfObjectValues(CreateDirectoryDefinitionFromDiscriminatorValue , m.SetDirectories)
    res["synchronizationRules"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfObjectValues(CreateSynchronizationRuleFromDiscriminatorValue , m.SetSynchronizationRules)
    res["version"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetVersion)
    return res
}
// GetSynchronizationRules gets the synchronizationRules property value. A collection of synchronization rules configured for the synchronizationJob or synchronizationTemplate.
func (m *SynchronizationSchema) GetSynchronizationRules()([]SynchronizationRuleable) {
    return m.synchronizationRules
}
// GetVersion gets the version property value. The version of the schema, updated automatically with every schema change.
func (m *SynchronizationSchema) GetVersion()(*string) {
    return m.version
}
// Serialize serializes information the current object
func (m *SynchronizationSchema) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetDirectories() != nil {
        cast := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.CollectionCast[i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable](m.GetDirectories())
        err = writer.WriteCollectionOfObjectValues("directories", cast)
        if err != nil {
            return err
        }
    }
    if m.GetSynchronizationRules() != nil {
        cast := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.CollectionCast[i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable](m.GetSynchronizationRules())
        err = writer.WriteCollectionOfObjectValues("synchronizationRules", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("version", m.GetVersion())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetDirectories sets the directories property value. Contains the collection of directories and all of their objects.
func (m *SynchronizationSchema) SetDirectories(value []DirectoryDefinitionable)() {
    m.directories = value
}
// SetSynchronizationRules sets the synchronizationRules property value. A collection of synchronization rules configured for the synchronizationJob or synchronizationTemplate.
func (m *SynchronizationSchema) SetSynchronizationRules(value []SynchronizationRuleable)() {
    m.synchronizationRules = value
}
// SetVersion sets the version property value. The version of the schema, updated automatically with every schema change.
func (m *SynchronizationSchema) SetVersion(value *string)() {
    m.version = value
}
