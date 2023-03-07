package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// MacOSExtensionsConfiguration 
type MacOSExtensionsConfiguration struct {
    DeviceConfiguration
}
// NewMacOSExtensionsConfiguration instantiates a new MacOSExtensionsConfiguration and sets the default values.
func NewMacOSExtensionsConfiguration()(*MacOSExtensionsConfiguration) {
    m := &MacOSExtensionsConfiguration{
        DeviceConfiguration: *NewDeviceConfiguration(),
    }
    odataTypeValue := "#microsoft.graph.macOSExtensionsConfiguration"
    m.SetOdataType(&odataTypeValue)
    return m
}
// CreateMacOSExtensionsConfigurationFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateMacOSExtensionsConfigurationFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewMacOSExtensionsConfiguration(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *MacOSExtensionsConfiguration) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.DeviceConfiguration.GetFieldDeserializers()
    res["kernelExtensionAllowedTeamIdentifiers"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]string, len(val))
            for i, v := range val {
                res[i] = *(v.(*string))
            }
            m.SetKernelExtensionAllowedTeamIdentifiers(res)
        }
        return nil
    }
    res["kernelExtensionOverridesAllowed"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetKernelExtensionOverridesAllowed(val)
        }
        return nil
    }
    res["kernelExtensionsAllowed"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateMacOSKernelExtensionFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]MacOSKernelExtensionable, len(val))
            for i, v := range val {
                res[i] = v.(MacOSKernelExtensionable)
            }
            m.SetKernelExtensionsAllowed(res)
        }
        return nil
    }
    res["systemExtensionsAllowed"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateMacOSSystemExtensionFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]MacOSSystemExtensionable, len(val))
            for i, v := range val {
                res[i] = v.(MacOSSystemExtensionable)
            }
            m.SetSystemExtensionsAllowed(res)
        }
        return nil
    }
    res["systemExtensionsAllowedTeamIdentifiers"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]string, len(val))
            for i, v := range val {
                res[i] = *(v.(*string))
            }
            m.SetSystemExtensionsAllowedTeamIdentifiers(res)
        }
        return nil
    }
    res["systemExtensionsAllowedTypes"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateMacOSSystemExtensionTypeMappingFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]MacOSSystemExtensionTypeMappingable, len(val))
            for i, v := range val {
                res[i] = v.(MacOSSystemExtensionTypeMappingable)
            }
            m.SetSystemExtensionsAllowedTypes(res)
        }
        return nil
    }
    res["systemExtensionsBlockOverride"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSystemExtensionsBlockOverride(val)
        }
        return nil
    }
    return res
}
// GetKernelExtensionAllowedTeamIdentifiers gets the kernelExtensionAllowedTeamIdentifiers property value. All kernel extensions validly signed by the team identifiers in this list will be allowed to load.
func (m *MacOSExtensionsConfiguration) GetKernelExtensionAllowedTeamIdentifiers()([]string) {
    val, err := m.GetBackingStore().Get("kernelExtensionAllowedTeamIdentifiers")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]string)
    }
    return nil
}
// GetKernelExtensionOverridesAllowed gets the kernelExtensionOverridesAllowed property value. If set to true, users can approve additional kernel extensions not explicitly allowed by configurations profiles.
func (m *MacOSExtensionsConfiguration) GetKernelExtensionOverridesAllowed()(*bool) {
    val, err := m.GetBackingStore().Get("kernelExtensionOverridesAllowed")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// GetKernelExtensionsAllowed gets the kernelExtensionsAllowed property value. A list of kernel extensions that will be allowed to load. . This collection can contain a maximum of 500 elements.
func (m *MacOSExtensionsConfiguration) GetKernelExtensionsAllowed()([]MacOSKernelExtensionable) {
    val, err := m.GetBackingStore().Get("kernelExtensionsAllowed")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]MacOSKernelExtensionable)
    }
    return nil
}
// GetSystemExtensionsAllowed gets the systemExtensionsAllowed property value. Gets or sets a list of allowed macOS system extensions. This collection can contain a maximum of 500 elements.
func (m *MacOSExtensionsConfiguration) GetSystemExtensionsAllowed()([]MacOSSystemExtensionable) {
    val, err := m.GetBackingStore().Get("systemExtensionsAllowed")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]MacOSSystemExtensionable)
    }
    return nil
}
// GetSystemExtensionsAllowedTeamIdentifiers gets the systemExtensionsAllowedTeamIdentifiers property value. Gets or sets a list of allowed team identifiers. Any system extension signed with any of the specified team identifiers will be approved.
func (m *MacOSExtensionsConfiguration) GetSystemExtensionsAllowedTeamIdentifiers()([]string) {
    val, err := m.GetBackingStore().Get("systemExtensionsAllowedTeamIdentifiers")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]string)
    }
    return nil
}
// GetSystemExtensionsAllowedTypes gets the systemExtensionsAllowedTypes property value. Gets or sets a list of allowed macOS system extension types. This collection can contain a maximum of 500 elements.
func (m *MacOSExtensionsConfiguration) GetSystemExtensionsAllowedTypes()([]MacOSSystemExtensionTypeMappingable) {
    val, err := m.GetBackingStore().Get("systemExtensionsAllowedTypes")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]MacOSSystemExtensionTypeMappingable)
    }
    return nil
}
// GetSystemExtensionsBlockOverride gets the systemExtensionsBlockOverride property value. Gets or sets whether to allow the user to approve additional system extensions not explicitly allowed by configuration profiles.
func (m *MacOSExtensionsConfiguration) GetSystemExtensionsBlockOverride()(*bool) {
    val, err := m.GetBackingStore().Get("systemExtensionsBlockOverride")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
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
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetKernelExtensionsAllowed()))
        for i, v := range m.GetKernelExtensionsAllowed() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("kernelExtensionsAllowed", cast)
        if err != nil {
            return err
        }
    }
    if m.GetSystemExtensionsAllowed() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetSystemExtensionsAllowed()))
        for i, v := range m.GetSystemExtensionsAllowed() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
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
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetSystemExtensionsAllowedTypes()))
        for i, v := range m.GetSystemExtensionsAllowedTypes() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
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
    err := m.GetBackingStore().Set("kernelExtensionAllowedTeamIdentifiers", value)
    if err != nil {
        panic(err)
    }
}
// SetKernelExtensionOverridesAllowed sets the kernelExtensionOverridesAllowed property value. If set to true, users can approve additional kernel extensions not explicitly allowed by configurations profiles.
func (m *MacOSExtensionsConfiguration) SetKernelExtensionOverridesAllowed(value *bool)() {
    err := m.GetBackingStore().Set("kernelExtensionOverridesAllowed", value)
    if err != nil {
        panic(err)
    }
}
// SetKernelExtensionsAllowed sets the kernelExtensionsAllowed property value. A list of kernel extensions that will be allowed to load. . This collection can contain a maximum of 500 elements.
func (m *MacOSExtensionsConfiguration) SetKernelExtensionsAllowed(value []MacOSKernelExtensionable)() {
    err := m.GetBackingStore().Set("kernelExtensionsAllowed", value)
    if err != nil {
        panic(err)
    }
}
// SetSystemExtensionsAllowed sets the systemExtensionsAllowed property value. Gets or sets a list of allowed macOS system extensions. This collection can contain a maximum of 500 elements.
func (m *MacOSExtensionsConfiguration) SetSystemExtensionsAllowed(value []MacOSSystemExtensionable)() {
    err := m.GetBackingStore().Set("systemExtensionsAllowed", value)
    if err != nil {
        panic(err)
    }
}
// SetSystemExtensionsAllowedTeamIdentifiers sets the systemExtensionsAllowedTeamIdentifiers property value. Gets or sets a list of allowed team identifiers. Any system extension signed with any of the specified team identifiers will be approved.
func (m *MacOSExtensionsConfiguration) SetSystemExtensionsAllowedTeamIdentifiers(value []string)() {
    err := m.GetBackingStore().Set("systemExtensionsAllowedTeamIdentifiers", value)
    if err != nil {
        panic(err)
    }
}
// SetSystemExtensionsAllowedTypes sets the systemExtensionsAllowedTypes property value. Gets or sets a list of allowed macOS system extension types. This collection can contain a maximum of 500 elements.
func (m *MacOSExtensionsConfiguration) SetSystemExtensionsAllowedTypes(value []MacOSSystemExtensionTypeMappingable)() {
    err := m.GetBackingStore().Set("systemExtensionsAllowedTypes", value)
    if err != nil {
        panic(err)
    }
}
// SetSystemExtensionsBlockOverride sets the systemExtensionsBlockOverride property value. Gets or sets whether to allow the user to approve additional system extensions not explicitly allowed by configuration profiles.
func (m *MacOSExtensionsConfiguration) SetSystemExtensionsBlockOverride(value *bool)() {
    err := m.GetBackingStore().Set("systemExtensionsBlockOverride", value)
    if err != nil {
        panic(err)
    }
}
// MacOSExtensionsConfigurationable 
type MacOSExtensionsConfigurationable interface {
    DeviceConfigurationable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetKernelExtensionAllowedTeamIdentifiers()([]string)
    GetKernelExtensionOverridesAllowed()(*bool)
    GetKernelExtensionsAllowed()([]MacOSKernelExtensionable)
    GetSystemExtensionsAllowed()([]MacOSSystemExtensionable)
    GetSystemExtensionsAllowedTeamIdentifiers()([]string)
    GetSystemExtensionsAllowedTypes()([]MacOSSystemExtensionTypeMappingable)
    GetSystemExtensionsBlockOverride()(*bool)
    SetKernelExtensionAllowedTeamIdentifiers(value []string)()
    SetKernelExtensionOverridesAllowed(value *bool)()
    SetKernelExtensionsAllowed(value []MacOSKernelExtensionable)()
    SetSystemExtensionsAllowed(value []MacOSSystemExtensionable)()
    SetSystemExtensionsAllowedTeamIdentifiers(value []string)()
    SetSystemExtensionsAllowedTypes(value []MacOSSystemExtensionTypeMappingable)()
    SetSystemExtensionsBlockOverride(value *bool)()
}
