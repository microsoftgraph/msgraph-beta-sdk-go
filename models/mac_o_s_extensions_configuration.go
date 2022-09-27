package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// MacOSExtensionsConfiguration 
type MacOSExtensionsConfiguration struct {
    DeviceConfiguration
    // All kernel extensions validly signed by the team identifiers in this list will be allowed to load.
    kernelExtensionAllowedTeamIdentifiers []string
    // If set to true, users can approve additional kernel extensions not explicitly allowed by configurations profiles.
    kernelExtensionOverridesAllowed *bool
    // A list of kernel extensions that will be allowed to load. . This collection can contain a maximum of 500 elements.
    kernelExtensionsAllowed []MacOSKernelExtensionable
    // Gets or sets a list of allowed macOS system extensions. This collection can contain a maximum of 500 elements.
    systemExtensionsAllowed []MacOSSystemExtensionable
    // Gets or sets a list of allowed team identifiers. Any system extension signed with any of the specified team identifiers will be approved.
    systemExtensionsAllowedTeamIdentifiers []string
    // Gets or sets a list of allowed macOS system extension types. This collection can contain a maximum of 500 elements.
    systemExtensionsAllowedTypes []MacOSSystemExtensionTypeMappingable
    // Gets or sets whether to allow the user to approve additional system extensions not explicitly allowed by configuration profiles.
    systemExtensionsBlockOverride *bool
}
// NewMacOSExtensionsConfiguration instantiates a new MacOSExtensionsConfiguration and sets the default values.
func NewMacOSExtensionsConfiguration()(*MacOSExtensionsConfiguration) {
    m := &MacOSExtensionsConfiguration{
        DeviceConfiguration: *NewDeviceConfiguration(),
    }
    odataTypeValue := "#microsoft.graph.macOSExtensionsConfiguration";
    m.SetOdataType(&odataTypeValue);
    return m
}
// CreateMacOSExtensionsConfigurationFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateMacOSExtensionsConfigurationFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewMacOSExtensionsConfiguration(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *MacOSExtensionsConfiguration) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.DeviceConfiguration.GetFieldDeserializers()
    res["kernelExtensionAllowedTeamIdentifiers"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfPrimitiveValues("string" , m.SetKernelExtensionAllowedTeamIdentifiers)
    res["kernelExtensionOverridesAllowed"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetBoolValue(m.SetKernelExtensionOverridesAllowed)
    res["kernelExtensionsAllowed"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfObjectValues(CreateMacOSKernelExtensionFromDiscriminatorValue , m.SetKernelExtensionsAllowed)
    res["systemExtensionsAllowed"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfObjectValues(CreateMacOSSystemExtensionFromDiscriminatorValue , m.SetSystemExtensionsAllowed)
    res["systemExtensionsAllowedTeamIdentifiers"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfPrimitiveValues("string" , m.SetSystemExtensionsAllowedTeamIdentifiers)
    res["systemExtensionsAllowedTypes"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfObjectValues(CreateMacOSSystemExtensionTypeMappingFromDiscriminatorValue , m.SetSystemExtensionsAllowedTypes)
    res["systemExtensionsBlockOverride"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetBoolValue(m.SetSystemExtensionsBlockOverride)
    return res
}
// GetKernelExtensionAllowedTeamIdentifiers gets the kernelExtensionAllowedTeamIdentifiers property value. All kernel extensions validly signed by the team identifiers in this list will be allowed to load.
func (m *MacOSExtensionsConfiguration) GetKernelExtensionAllowedTeamIdentifiers()([]string) {
    return m.kernelExtensionAllowedTeamIdentifiers
}
// GetKernelExtensionOverridesAllowed gets the kernelExtensionOverridesAllowed property value. If set to true, users can approve additional kernel extensions not explicitly allowed by configurations profiles.
func (m *MacOSExtensionsConfiguration) GetKernelExtensionOverridesAllowed()(*bool) {
    return m.kernelExtensionOverridesAllowed
}
// GetKernelExtensionsAllowed gets the kernelExtensionsAllowed property value. A list of kernel extensions that will be allowed to load. . This collection can contain a maximum of 500 elements.
func (m *MacOSExtensionsConfiguration) GetKernelExtensionsAllowed()([]MacOSKernelExtensionable) {
    return m.kernelExtensionsAllowed
}
// GetSystemExtensionsAllowed gets the systemExtensionsAllowed property value. Gets or sets a list of allowed macOS system extensions. This collection can contain a maximum of 500 elements.
func (m *MacOSExtensionsConfiguration) GetSystemExtensionsAllowed()([]MacOSSystemExtensionable) {
    return m.systemExtensionsAllowed
}
// GetSystemExtensionsAllowedTeamIdentifiers gets the systemExtensionsAllowedTeamIdentifiers property value. Gets or sets a list of allowed team identifiers. Any system extension signed with any of the specified team identifiers will be approved.
func (m *MacOSExtensionsConfiguration) GetSystemExtensionsAllowedTeamIdentifiers()([]string) {
    return m.systemExtensionsAllowedTeamIdentifiers
}
// GetSystemExtensionsAllowedTypes gets the systemExtensionsAllowedTypes property value. Gets or sets a list of allowed macOS system extension types. This collection can contain a maximum of 500 elements.
func (m *MacOSExtensionsConfiguration) GetSystemExtensionsAllowedTypes()([]MacOSSystemExtensionTypeMappingable) {
    return m.systemExtensionsAllowedTypes
}
// GetSystemExtensionsBlockOverride gets the systemExtensionsBlockOverride property value. Gets or sets whether to allow the user to approve additional system extensions not explicitly allowed by configuration profiles.
func (m *MacOSExtensionsConfiguration) GetSystemExtensionsBlockOverride()(*bool) {
    return m.systemExtensionsBlockOverride
}
// Serialize serializes information the current object
func (m *MacOSExtensionsConfiguration) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.DeviceConfiguration.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetKernelExtensionAllowedTeamIdentifiers() != nil {
        err = writer.WriteCollectionOfStringValues("kernelExtensionAllowedTeamIdentifiers", m.GetKernelExtensionAllowedTeamIdentifiers())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("kernelExtensionOverridesAllowed", m.GetKernelExtensionOverridesAllowed())
        if err != nil {
            return err
        }
    }
    if m.GetKernelExtensionsAllowed() != nil {
        cast := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.CollectionCast[i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable](m.GetKernelExtensionsAllowed())
        err = writer.WriteCollectionOfObjectValues("kernelExtensionsAllowed", cast)
        if err != nil {
            return err
        }
    }
    if m.GetSystemExtensionsAllowed() != nil {
        cast := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.CollectionCast[i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable](m.GetSystemExtensionsAllowed())
        err = writer.WriteCollectionOfObjectValues("systemExtensionsAllowed", cast)
        if err != nil {
            return err
        }
    }
    if m.GetSystemExtensionsAllowedTeamIdentifiers() != nil {
        err = writer.WriteCollectionOfStringValues("systemExtensionsAllowedTeamIdentifiers", m.GetSystemExtensionsAllowedTeamIdentifiers())
        if err != nil {
            return err
        }
    }
    if m.GetSystemExtensionsAllowedTypes() != nil {
        cast := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.CollectionCast[i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable](m.GetSystemExtensionsAllowedTypes())
        err = writer.WriteCollectionOfObjectValues("systemExtensionsAllowedTypes", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("systemExtensionsBlockOverride", m.GetSystemExtensionsBlockOverride())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetKernelExtensionAllowedTeamIdentifiers sets the kernelExtensionAllowedTeamIdentifiers property value. All kernel extensions validly signed by the team identifiers in this list will be allowed to load.
func (m *MacOSExtensionsConfiguration) SetKernelExtensionAllowedTeamIdentifiers(value []string)() {
    m.kernelExtensionAllowedTeamIdentifiers = value
}
// SetKernelExtensionOverridesAllowed sets the kernelExtensionOverridesAllowed property value. If set to true, users can approve additional kernel extensions not explicitly allowed by configurations profiles.
func (m *MacOSExtensionsConfiguration) SetKernelExtensionOverridesAllowed(value *bool)() {
    m.kernelExtensionOverridesAllowed = value
}
// SetKernelExtensionsAllowed sets the kernelExtensionsAllowed property value. A list of kernel extensions that will be allowed to load. . This collection can contain a maximum of 500 elements.
func (m *MacOSExtensionsConfiguration) SetKernelExtensionsAllowed(value []MacOSKernelExtensionable)() {
    m.kernelExtensionsAllowed = value
}
// SetSystemExtensionsAllowed sets the systemExtensionsAllowed property value. Gets or sets a list of allowed macOS system extensions. This collection can contain a maximum of 500 elements.
func (m *MacOSExtensionsConfiguration) SetSystemExtensionsAllowed(value []MacOSSystemExtensionable)() {
    m.systemExtensionsAllowed = value
}
// SetSystemExtensionsAllowedTeamIdentifiers sets the systemExtensionsAllowedTeamIdentifiers property value. Gets or sets a list of allowed team identifiers. Any system extension signed with any of the specified team identifiers will be approved.
func (m *MacOSExtensionsConfiguration) SetSystemExtensionsAllowedTeamIdentifiers(value []string)() {
    m.systemExtensionsAllowedTeamIdentifiers = value
}
// SetSystemExtensionsAllowedTypes sets the systemExtensionsAllowedTypes property value. Gets or sets a list of allowed macOS system extension types. This collection can contain a maximum of 500 elements.
func (m *MacOSExtensionsConfiguration) SetSystemExtensionsAllowedTypes(value []MacOSSystemExtensionTypeMappingable)() {
    m.systemExtensionsAllowedTypes = value
}
// SetSystemExtensionsBlockOverride sets the systemExtensionsBlockOverride property value. Gets or sets whether to allow the user to approve additional system extensions not explicitly allowed by configuration profiles.
func (m *MacOSExtensionsConfiguration) SetSystemExtensionsBlockOverride(value *bool)() {
    m.systemExtensionsBlockOverride = value
}
