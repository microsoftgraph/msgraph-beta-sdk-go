package security

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// RegistryValueEvidence 
type RegistryValueEvidence struct {
    AlertEvidence
    // Registry hive of the key that the recorded action was applied to.
    registryHive *string
    // Registry key that the recorded action was applied to.
    registryKey *string
    // Data of the registry value that the recorded action was applied to.
    registryValue *string
    // Name of the registry value that the recorded action was applied to.
    registryValueName *string
    // Data type, such as binary or string, of the registry value that the recorded action was applied to.
    registryValueType *string
}
// NewRegistryValueEvidence instantiates a new RegistryValueEvidence and sets the default values.
func NewRegistryValueEvidence()(*RegistryValueEvidence) {
    m := &RegistryValueEvidence{
        AlertEvidence: *NewAlertEvidence(),
    }
    return m
}
// CreateRegistryValueEvidenceFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateRegistryValueEvidenceFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewRegistryValueEvidence(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *RegistryValueEvidence) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.AlertEvidence.GetFieldDeserializers()
    res["registryHive"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetRegistryHive)
    res["registryKey"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetRegistryKey)
    res["registryValue"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetRegistryValue)
    res["registryValueName"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetRegistryValueName)
    res["registryValueType"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetRegistryValueType)
    return res
}
// GetRegistryHive gets the registryHive property value. Registry hive of the key that the recorded action was applied to.
func (m *RegistryValueEvidence) GetRegistryHive()(*string) {
    return m.registryHive
}
// GetRegistryKey gets the registryKey property value. Registry key that the recorded action was applied to.
func (m *RegistryValueEvidence) GetRegistryKey()(*string) {
    return m.registryKey
}
// GetRegistryValue gets the registryValue property value. Data of the registry value that the recorded action was applied to.
func (m *RegistryValueEvidence) GetRegistryValue()(*string) {
    return m.registryValue
}
// GetRegistryValueName gets the registryValueName property value. Name of the registry value that the recorded action was applied to.
func (m *RegistryValueEvidence) GetRegistryValueName()(*string) {
    return m.registryValueName
}
// GetRegistryValueType gets the registryValueType property value. Data type, such as binary or string, of the registry value that the recorded action was applied to.
func (m *RegistryValueEvidence) GetRegistryValueType()(*string) {
    return m.registryValueType
}
// Serialize serializes information the current object
func (m *RegistryValueEvidence) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
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
    {
        err = writer.WriteStringValue("registryValue", m.GetRegistryValue())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("registryValueName", m.GetRegistryValueName())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("registryValueType", m.GetRegistryValueType())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetRegistryHive sets the registryHive property value. Registry hive of the key that the recorded action was applied to.
func (m *RegistryValueEvidence) SetRegistryHive(value *string)() {
    m.registryHive = value
}
// SetRegistryKey sets the registryKey property value. Registry key that the recorded action was applied to.
func (m *RegistryValueEvidence) SetRegistryKey(value *string)() {
    m.registryKey = value
}
// SetRegistryValue sets the registryValue property value. Data of the registry value that the recorded action was applied to.
func (m *RegistryValueEvidence) SetRegistryValue(value *string)() {
    m.registryValue = value
}
// SetRegistryValueName sets the registryValueName property value. Name of the registry value that the recorded action was applied to.
func (m *RegistryValueEvidence) SetRegistryValueName(value *string)() {
    m.registryValueName = value
}
// SetRegistryValueType sets the registryValueType property value. Data type, such as binary or string, of the registry value that the recorded action was applied to.
func (m *RegistryValueEvidence) SetRegistryValueType(value *string)() {
    m.registryValueType = value
}
