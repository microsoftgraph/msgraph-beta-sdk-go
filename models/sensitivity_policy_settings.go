package models

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// SensitivityPolicySettings 
type SensitivityPolicySettings struct {
    Entity
    // The applicableTo property
    applicableTo *SensitivityLabelTarget
    // The downgradeSensitivityRequiresJustification property
    downgradeSensitivityRequiresJustification *bool
    // The helpWebUrl property
    helpWebUrl *string
    // The isMandatory property
    isMandatory *bool
}
// NewSensitivityPolicySettings instantiates a new sensitivityPolicySettings and sets the default values.
func NewSensitivityPolicySettings()(*SensitivityPolicySettings) {
    m := &SensitivityPolicySettings{
        Entity: *NewEntity(),
    }
    odataTypeValue := "#microsoft.graph.sensitivityPolicySettings";
    m.SetOdataType(&odataTypeValue);
    return m
}
// CreateSensitivityPolicySettingsFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateSensitivityPolicySettingsFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewSensitivityPolicySettings(), nil
}
// GetApplicableTo gets the applicableTo property value. The applicableTo property
func (m *SensitivityPolicySettings) GetApplicableTo()(*SensitivityLabelTarget) {
    return m.applicableTo
}
// GetDowngradeSensitivityRequiresJustification gets the downgradeSensitivityRequiresJustification property value. The downgradeSensitivityRequiresJustification property
func (m *SensitivityPolicySettings) GetDowngradeSensitivityRequiresJustification()(*bool) {
    return m.downgradeSensitivityRequiresJustification
}
// GetFieldDeserializers the deserialization information for the current model
func (m *SensitivityPolicySettings) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["applicableTo"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetEnumValue(ParseSensitivityLabelTarget , m.SetApplicableTo)
    res["downgradeSensitivityRequiresJustification"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetBoolValue(m.SetDowngradeSensitivityRequiresJustification)
    res["helpWebUrl"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetHelpWebUrl)
    res["isMandatory"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetBoolValue(m.SetIsMandatory)
    return res
}
// GetHelpWebUrl gets the helpWebUrl property value. The helpWebUrl property
func (m *SensitivityPolicySettings) GetHelpWebUrl()(*string) {
    return m.helpWebUrl
}
// GetIsMandatory gets the isMandatory property value. The isMandatory property
func (m *SensitivityPolicySettings) GetIsMandatory()(*bool) {
    return m.isMandatory
}
// Serialize serializes information the current object
func (m *SensitivityPolicySettings) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetApplicableTo() != nil {
        cast := (*m.GetApplicableTo()).String()
        err = writer.WriteStringValue("applicableTo", &cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("downgradeSensitivityRequiresJustification", m.GetDowngradeSensitivityRequiresJustification())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("helpWebUrl", m.GetHelpWebUrl())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("isMandatory", m.GetIsMandatory())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetApplicableTo sets the applicableTo property value. The applicableTo property
func (m *SensitivityPolicySettings) SetApplicableTo(value *SensitivityLabelTarget)() {
    m.applicableTo = value
}
// SetDowngradeSensitivityRequiresJustification sets the downgradeSensitivityRequiresJustification property value. The downgradeSensitivityRequiresJustification property
func (m *SensitivityPolicySettings) SetDowngradeSensitivityRequiresJustification(value *bool)() {
    m.downgradeSensitivityRequiresJustification = value
}
// SetHelpWebUrl sets the helpWebUrl property value. The helpWebUrl property
func (m *SensitivityPolicySettings) SetHelpWebUrl(value *string)() {
    m.helpWebUrl = value
}
// SetIsMandatory sets the isMandatory property value. The isMandatory property
func (m *SensitivityPolicySettings) SetIsMandatory(value *bool)() {
    m.isMandatory = value
}
