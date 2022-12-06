package models

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// OnPremisesDirectorySynchronization provides operations to manage the collection of accessReviewDecision entities.
type OnPremisesDirectorySynchronization struct {
    Entity
    // The configuration property
    configuration OnPremisesDirectorySynchronizationConfigurationable
    // The features property
    features OnPremisesDirectorySynchronizationFeatureable
}
// NewOnPremisesDirectorySynchronization instantiates a new onPremisesDirectorySynchronization and sets the default values.
func NewOnPremisesDirectorySynchronization()(*OnPremisesDirectorySynchronization) {
    m := &OnPremisesDirectorySynchronization{
        Entity: *NewEntity(),
    }
    return m
}
// CreateOnPremisesDirectorySynchronizationFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateOnPremisesDirectorySynchronizationFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewOnPremisesDirectorySynchronization(), nil
}
// GetConfiguration gets the configuration property value. The configuration property
func (m *OnPremisesDirectorySynchronization) GetConfiguration()(OnPremisesDirectorySynchronizationConfigurationable) {
    return m.configuration
}
// GetFeatures gets the features property value. The features property
func (m *OnPremisesDirectorySynchronization) GetFeatures()(OnPremisesDirectorySynchronizationFeatureable) {
    return m.features
}
// GetFieldDeserializers the deserialization information for the current model
func (m *OnPremisesDirectorySynchronization) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["configuration"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetObjectValue(CreateOnPremisesDirectorySynchronizationConfigurationFromDiscriminatorValue , m.SetConfiguration)
    res["features"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetObjectValue(CreateOnPremisesDirectorySynchronizationFeatureFromDiscriminatorValue , m.SetFeatures)
    return res
}
// Serialize serializes information the current object
func (m *OnPremisesDirectorySynchronization) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteObjectValue("configuration", m.GetConfiguration())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("features", m.GetFeatures())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetConfiguration sets the configuration property value. The configuration property
func (m *OnPremisesDirectorySynchronization) SetConfiguration(value OnPremisesDirectorySynchronizationConfigurationable)() {
    m.configuration = value
}
// SetFeatures sets the features property value. The features property
func (m *OnPremisesDirectorySynchronization) SetFeatures(value OnPremisesDirectorySynchronizationFeatureable)() {
    m.features = value
}
