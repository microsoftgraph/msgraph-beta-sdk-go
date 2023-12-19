package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// SensitivityPolicySettings 
type SensitivityPolicySettings struct {
    Entity
}
// NewSensitivityPolicySettings instantiates a new sensitivityPolicySettings and sets the default values.
func NewSensitivityPolicySettings()(*SensitivityPolicySettings) {
    m := &SensitivityPolicySettings{
        Entity: *NewEntity(),
    }
    return m
}
// CreateSensitivityPolicySettingsFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateSensitivityPolicySettingsFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewSensitivityPolicySettings(), nil
}
// GetApplicableTo gets the applicableTo property value. The applicableTo property
func (m *SensitivityPolicySettings) GetApplicableTo()(*SensitivityPolicySettings_applicableTo) {
    val, err := m.GetBackingStore().Get("applicableTo")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*SensitivityPolicySettings_applicableTo)
    }
    return nil
}
// GetDowngradeSensitivityRequiresJustification gets the downgradeSensitivityRequiresJustification property value. The downgradeSensitivityRequiresJustification property
func (m *SensitivityPolicySettings) GetDowngradeSensitivityRequiresJustification()(*bool) {
    val, err := m.GetBackingStore().Get("downgradeSensitivityRequiresJustification")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *SensitivityPolicySettings) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["applicableTo"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseSensitivityPolicySettings_applicableTo)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetApplicableTo(val.(*SensitivityPolicySettings_applicableTo))
        }
        return nil
    }
    res["downgradeSensitivityRequiresJustification"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDowngradeSensitivityRequiresJustification(val)
        }
        return nil
    }
    res["helpWebUrl"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetHelpWebUrl(val)
        }
        return nil
    }
    res["isMandatory"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsMandatory(val)
        }
        return nil
    }
    return res
}
// GetHelpWebUrl gets the helpWebUrl property value. The helpWebUrl property
func (m *SensitivityPolicySettings) GetHelpWebUrl()(*string) {
    val, err := m.GetBackingStore().Get("helpWebUrl")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetIsMandatory gets the isMandatory property value. The isMandatory property
func (m *SensitivityPolicySettings) GetIsMandatory()(*bool) {
    val, err := m.GetBackingStore().Get("isMandatory")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
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
func (m *SensitivityPolicySettings) SetApplicableTo(value *SensitivityPolicySettings_applicableTo)() {
    err := m.GetBackingStore().Set("applicableTo", value)
    if err != nil {
        panic(err)
    }
}
// SetDowngradeSensitivityRequiresJustification sets the downgradeSensitivityRequiresJustification property value. The downgradeSensitivityRequiresJustification property
func (m *SensitivityPolicySettings) SetDowngradeSensitivityRequiresJustification(value *bool)() {
    err := m.GetBackingStore().Set("downgradeSensitivityRequiresJustification", value)
    if err != nil {
        panic(err)
    }
}
// SetHelpWebUrl sets the helpWebUrl property value. The helpWebUrl property
func (m *SensitivityPolicySettings) SetHelpWebUrl(value *string)() {
    err := m.GetBackingStore().Set("helpWebUrl", value)
    if err != nil {
        panic(err)
    }
}
// SetIsMandatory sets the isMandatory property value. The isMandatory property
func (m *SensitivityPolicySettings) SetIsMandatory(value *bool)() {
    err := m.GetBackingStore().Set("isMandatory", value)
    if err != nil {
        panic(err)
    }
}
// SensitivityPolicySettingsable 
type SensitivityPolicySettingsable interface {
    Entityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetApplicableTo()(*SensitivityPolicySettings_applicableTo)
    GetDowngradeSensitivityRequiresJustification()(*bool)
    GetHelpWebUrl()(*string)
    GetIsMandatory()(*bool)
    SetApplicableTo(value *SensitivityPolicySettings_applicableTo)()
    SetDowngradeSensitivityRequiresJustification(value *bool)()
    SetHelpWebUrl(value *string)()
    SetIsMandatory(value *bool)()
}
