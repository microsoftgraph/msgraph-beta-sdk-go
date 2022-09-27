package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// DataLossPreventionPolicy provides operations to manage the collection of accessReviewDecision entities.
type DataLossPreventionPolicy struct {
    Entity
    // The name property
    name *string
}
// NewDataLossPreventionPolicy instantiates a new dataLossPreventionPolicy and sets the default values.
func NewDataLossPreventionPolicy()(*DataLossPreventionPolicy) {
    m := &DataLossPreventionPolicy{
        Entity: *NewEntity(),
    }
    odataTypeValue := "#microsoft.graph.dataLossPreventionPolicy";
    m.SetOdataType(&odataTypeValue);
    return m
}
// CreateDataLossPreventionPolicyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateDataLossPreventionPolicyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewDataLossPreventionPolicy(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *DataLossPreventionPolicy) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["name"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetName)
    return res
}
// GetName gets the name property value. The name property
func (m *DataLossPreventionPolicy) GetName()(*string) {
    return m.name
}
// Serialize serializes information the current object
func (m *DataLossPreventionPolicy) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteStringValue("name", m.GetName())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetName sets the name property value. The name property
func (m *DataLossPreventionPolicy) SetName(value *string)() {
    m.name = value
}
