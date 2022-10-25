package models

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// BusinessFlowTemplate provides operations to manage the collection of accessReviewDecision entities.
type BusinessFlowTemplate struct {
    Entity
    // The name of the business flow template
    displayName *string
}
// NewBusinessFlowTemplate instantiates a new businessFlowTemplate and sets the default values.
func NewBusinessFlowTemplate()(*BusinessFlowTemplate) {
    m := &BusinessFlowTemplate{
        Entity: *NewEntity(),
    }
    odataTypeValue := "#microsoft.graph.businessFlowTemplate";
    m.SetOdataType(&odataTypeValue);
    return m
}
// CreateBusinessFlowTemplateFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateBusinessFlowTemplateFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewBusinessFlowTemplate(), nil
}
// GetDisplayName gets the displayName property value. The name of the business flow template
func (m *BusinessFlowTemplate) GetDisplayName()(*string) {
    return m.displayName
}
// GetFieldDeserializers the deserialization information for the current model
func (m *BusinessFlowTemplate) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["displayName"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetDisplayName)
    return res
}
// Serialize serializes information the current object
func (m *BusinessFlowTemplate) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteStringValue("displayName", m.GetDisplayName())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetDisplayName sets the displayName property value. The name of the business flow template
func (m *BusinessFlowTemplate) SetDisplayName(value *string)() {
    m.displayName = value
}
