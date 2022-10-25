package models

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// GovernancePolicyTemplate provides operations to manage the collection of accessReviewDecision entities.
type GovernancePolicyTemplate struct {
    Entity
    // The displayName property
    displayName *string
    // The policy property
    policy GovernancePolicyable
    // The settings property
    settings BusinessFlowSettingsable
}
// NewGovernancePolicyTemplate instantiates a new governancePolicyTemplate and sets the default values.
func NewGovernancePolicyTemplate()(*GovernancePolicyTemplate) {
    m := &GovernancePolicyTemplate{
        Entity: *NewEntity(),
    }
    odataTypeValue := "#microsoft.graph.governancePolicyTemplate";
    m.SetOdataType(&odataTypeValue);
    return m
}
// CreateGovernancePolicyTemplateFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateGovernancePolicyTemplateFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewGovernancePolicyTemplate(), nil
}
// GetDisplayName gets the displayName property value. The displayName property
func (m *GovernancePolicyTemplate) GetDisplayName()(*string) {
    return m.displayName
}
// GetFieldDeserializers the deserialization information for the current model
func (m *GovernancePolicyTemplate) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["displayName"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetDisplayName)
    res["policy"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetObjectValue(CreateGovernancePolicyFromDiscriminatorValue , m.SetPolicy)
    res["settings"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetObjectValue(CreateBusinessFlowSettingsFromDiscriminatorValue , m.SetSettings)
    return res
}
// GetPolicy gets the policy property value. The policy property
func (m *GovernancePolicyTemplate) GetPolicy()(GovernancePolicyable) {
    return m.policy
}
// GetSettings gets the settings property value. The settings property
func (m *GovernancePolicyTemplate) GetSettings()(BusinessFlowSettingsable) {
    return m.settings
}
// Serialize serializes information the current object
func (m *GovernancePolicyTemplate) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
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
    {
        err = writer.WriteObjectValue("policy", m.GetPolicy())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("settings", m.GetSettings())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetDisplayName sets the displayName property value. The displayName property
func (m *GovernancePolicyTemplate) SetDisplayName(value *string)() {
    m.displayName = value
}
// SetPolicy sets the policy property value. The policy property
func (m *GovernancePolicyTemplate) SetPolicy(value GovernancePolicyable)() {
    m.policy = value
}
// SetSettings sets the settings property value. The settings property
func (m *GovernancePolicyTemplate) SetSettings(value BusinessFlowSettingsable)() {
    m.settings = value
}
