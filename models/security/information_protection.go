package security

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
)

// InformationProtection 
type InformationProtection struct {
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Entity
}
// NewInformationProtection instantiates a new informationProtection and sets the default values.
func NewInformationProtection()(*InformationProtection) {
    m := &InformationProtection{
        Entity: *ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.NewEntity(),
    }
    return m
}
// CreateInformationProtectionFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateInformationProtectionFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewInformationProtection(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *InformationProtection) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["labelPolicySettings"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateInformationProtectionPolicySettingFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLabelPolicySettings(val.(InformationProtectionPolicySettingable))
        }
        return nil
    }
    res["sensitivityLabels"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateSensitivityLabelFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]SensitivityLabelable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(SensitivityLabelable)
                }
            }
            m.SetSensitivityLabels(res)
        }
        return nil
    }
    return res
}
// GetLabelPolicySettings gets the labelPolicySettings property value. Read the Microsoft Purview Information Protection policy settings for the user or organization.
func (m *InformationProtection) GetLabelPolicySettings()(InformationProtectionPolicySettingable) {
    val, err := m.GetBackingStore().Get("labelPolicySettings")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(InformationProtectionPolicySettingable)
    }
    return nil
}
// GetSensitivityLabels gets the sensitivityLabels property value. Read the Microsoft Purview Information Protection labels for the user or organization.
func (m *InformationProtection) GetSensitivityLabels()([]SensitivityLabelable) {
    val, err := m.GetBackingStore().Get("sensitivityLabels")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]SensitivityLabelable)
    }
    return nil
}
// Serialize serializes information the current object
func (m *InformationProtection) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteObjectValue("labelPolicySettings", m.GetLabelPolicySettings())
        if err != nil {
            return err
        }
    }
    if m.GetSensitivityLabels() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetSensitivityLabels()))
        for i, v := range m.GetSensitivityLabels() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err = writer.WriteCollectionOfObjectValues("sensitivityLabels", cast)
        if err != nil {
            return err
        }
    }
    return nil
}
// SetLabelPolicySettings sets the labelPolicySettings property value. Read the Microsoft Purview Information Protection policy settings for the user or organization.
func (m *InformationProtection) SetLabelPolicySettings(value InformationProtectionPolicySettingable)() {
    err := m.GetBackingStore().Set("labelPolicySettings", value)
    if err != nil {
        panic(err)
    }
}
// SetSensitivityLabels sets the sensitivityLabels property value. Read the Microsoft Purview Information Protection labels for the user or organization.
func (m *InformationProtection) SetSensitivityLabels(value []SensitivityLabelable)() {
    err := m.GetBackingStore().Set("sensitivityLabels", value)
    if err != nil {
        panic(err)
    }
}
// InformationProtectionable 
type InformationProtectionable interface {
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Entityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetLabelPolicySettings()(InformationProtectionPolicySettingable)
    GetSensitivityLabels()([]SensitivityLabelable)
    SetLabelPolicySettings(value InformationProtectionPolicySettingable)()
    SetSensitivityLabels(value []SensitivityLabelable)()
}
