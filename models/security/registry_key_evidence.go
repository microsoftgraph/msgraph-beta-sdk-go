package security

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// RegistryKeyEvidence 
type RegistryKeyEvidence struct {
    AlertEvidence
    // Registry hive of the key that the recorded action was applied to.
    registryHive *string
    // Registry key that the recorded action was applied to.
    registryKey *string
}
// NewRegistryKeyEvidence instantiates a new RegistryKeyEvidence and sets the default values.
func NewRegistryKeyEvidence()(*RegistryKeyEvidence) {
    m := &RegistryKeyEvidence{
        AlertEvidence: *NewAlertEvidence(),
    }
    odataTypeValue := "#microsoft.graph.security.registryKeyEvidence";
    m.SetOdataType(&odataTypeValue);
    return m
}
// CreateRegistryKeyEvidenceFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateRegistryKeyEvidenceFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewRegistryKeyEvidence(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *RegistryKeyEvidence) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.AlertEvidence.GetFieldDeserializers()
    res["registryHive"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetRegistryHive)
    res["registryKey"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetRegistryKey)
    return res
}
// GetRegistryHive gets the registryHive property value. Registry hive of the key that the recorded action was applied to.
func (m *RegistryKeyEvidence) GetRegistryHive()(*string) {
    return m.registryHive
}
// GetRegistryKey gets the registryKey property value. Registry key that the recorded action was applied to.
func (m *RegistryKeyEvidence) GetRegistryKey()(*string) {
    return m.registryKey
}
// Serialize serializes information the current object
func (m *RegistryKeyEvidence) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.AlertEvidence.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteStringValue("registryHive", m.GetRegistryHive())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("registryKey", m.GetRegistryKey())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetRegistryHive sets the registryHive property value. Registry hive of the key that the recorded action was applied to.
func (m *RegistryKeyEvidence) SetRegistryHive(value *string)() {
    m.registryHive = value
}
// SetRegistryKey sets the registryKey property value. Registry key that the recorded action was applied to.
func (m *RegistryKeyEvidence) SetRegistryKey(value *string)() {
    m.registryKey = value
}
