package security

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
)

// InformationProtection 
type InformationProtection struct {
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Entity
    // Read the Microsoft Purview Information Protection policy settings for the user or organization.
    labelPolicySettings InformationProtectionPolicySettingable
    // Read the Microsoft Purview Information Protection labels for the user or organization.
    sensitivityLabels []SensitivityLabelable
}
// NewInformationProtection instantiates a new informationProtection and sets the default values.
func NewInformationProtection()(*InformationProtection) {
    m := &InformationProtection{
        Entity: *ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.NewEntity(),
    }
    odataTypeValue := "#microsoft.graph.security.informationProtection";
    m.SetOdataType(&odataTypeValue);
    return m
}
// CreateInformationProtectionFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateInformationProtectionFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewInformationProtection(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *InformationProtection) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["labelPolicySettings"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetObjectValue(CreateInformationProtectionPolicySettingFromDiscriminatorValue , m.SetLabelPolicySettings)
    res["sensitivityLabels"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfObjectValues(CreateSensitivityLabelFromDiscriminatorValue , m.SetSensitivityLabels)
    return res
}
// GetLabelPolicySettings gets the labelPolicySettings property value. Read the Microsoft Purview Information Protection policy settings for the user or organization.
func (m *InformationProtection) GetLabelPolicySettings()(InformationProtectionPolicySettingable) {
    return m.labelPolicySettings
}
// GetSensitivityLabels gets the sensitivityLabels property value. Read the Microsoft Purview Information Protection labels for the user or organization.
func (m *InformationProtection) GetSensitivityLabels()([]SensitivityLabelable) {
    return m.sensitivityLabels
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
        cast := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.CollectionCast[i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable](m.GetSensitivityLabels())
        err = writer.WriteCollectionOfObjectValues("sensitivityLabels", cast)
        if err != nil {
            return err
        }
    }
    return nil
}
// SetLabelPolicySettings sets the labelPolicySettings property value. Read the Microsoft Purview Information Protection policy settings for the user or organization.
func (m *InformationProtection) SetLabelPolicySettings(value InformationProtectionPolicySettingable)() {
    m.labelPolicySettings = value
}
// SetSensitivityLabels sets the sensitivityLabels property value. Read the Microsoft Purview Information Protection labels for the user or organization.
func (m *InformationProtection) SetSensitivityLabels(value []SensitivityLabelable)() {
    m.sensitivityLabels = value
}
