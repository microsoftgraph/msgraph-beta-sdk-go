package models

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// AllowedValue provides operations to manage the collection of accessReviewDecision entities.
type AllowedValue struct {
    Entity
    // Indicates whether the predefined value is active or deactivated. If set to false, this predefined value cannot be assigned to any additional supported directory objects.
    isActive *bool
}
// NewAllowedValue instantiates a new allowedValue and sets the default values.
func NewAllowedValue()(*AllowedValue) {
    m := &AllowedValue{
        Entity: *NewEntity(),
    }
    odataTypeValue := "#microsoft.graph.allowedValue";
    m.SetOdataType(&odataTypeValue);
    return m
}
// CreateAllowedValueFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateAllowedValueFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewAllowedValue(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *AllowedValue) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["isActive"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetBoolValue(m.SetIsActive)
    return res
}
// GetIsActive gets the isActive property value. Indicates whether the predefined value is active or deactivated. If set to false, this predefined value cannot be assigned to any additional supported directory objects.
func (m *AllowedValue) GetIsActive()(*bool) {
    return m.isActive
}
// Serialize serializes information the current object
func (m *AllowedValue) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteBoolValue("isActive", m.GetIsActive())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetIsActive sets the isActive property value. Indicates whether the predefined value is active or deactivated. If set to false, this predefined value cannot be assigned to any additional supported directory objects.
func (m *AllowedValue) SetIsActive(value *bool)() {
    m.isActive = value
}
