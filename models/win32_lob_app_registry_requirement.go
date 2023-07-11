package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// Win32LobAppRegistryRequirement base class to detect a Win32 App
type Win32LobAppRegistryRequirement struct {
    Win32LobAppRequirement
}
// NewWin32LobAppRegistryRequirement instantiates a new win32LobAppRegistryRequirement and sets the default values.
func NewWin32LobAppRegistryRequirement()(*Win32LobAppRegistryRequirement) {
    m := &Win32LobAppRegistryRequirement{
        Win32LobAppRequirement: *NewWin32LobAppRequirement(),
    }
    odataTypeValue := "#microsoft.graph.win32LobAppRegistryRequirement"
    m.SetOdataType(&odataTypeValue)
    return m
}
// CreateWin32LobAppRegistryRequirementFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateWin32LobAppRegistryRequirementFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewWin32LobAppRegistryRequirement(), nil
}
// GetCheck32BitOn64System gets the check32BitOn64System property value. A value indicating whether this registry path is for checking 32-bit app on 64-bit system
func (m *Win32LobAppRegistryRequirement) GetCheck32BitOn64System()(*bool) {
    val, err := m.GetBackingStore().Get("check32BitOn64System")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// GetDetectionType gets the detectionType property value. Contains all supported registry data detection type.
func (m *Win32LobAppRegistryRequirement) GetDetectionType()(*Win32LobAppRegistryDetectionType) {
    val, err := m.GetBackingStore().Get("detectionType")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*Win32LobAppRegistryDetectionType)
    }
    return nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *Win32LobAppRegistryRequirement) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Win32LobAppRequirement.GetFieldDeserializers()
    res["check32BitOn64System"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCheck32BitOn64System(val)
        }
        return nil
    }
    res["detectionType"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseWin32LobAppRegistryDetectionType)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDetectionType(val.(*Win32LobAppRegistryDetectionType))
        }
        return nil
    }
    res["keyPath"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetKeyPath(val)
        }
        return nil
    }
    res["@odata.type"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOdataType(val)
        }
        return nil
    }
    res["valueName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetValueName(val)
        }
        return nil
    }
    return res
}
// GetKeyPath gets the keyPath property value. The registry key path to detect Win32 Line of Business (LoB) app
func (m *Win32LobAppRegistryRequirement) GetKeyPath()(*string) {
    val, err := m.GetBackingStore().Get("keyPath")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetOdataType gets the @odata.type property value. The OdataType property
func (m *Win32LobAppRegistryRequirement) GetOdataType()(*string) {
    val, err := m.GetBackingStore().Get("odataType")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetValueName gets the valueName property value. The registry value name
func (m *Win32LobAppRegistryRequirement) GetValueName()(*string) {
    val, err := m.GetBackingStore().Get("valueName")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// Serialize serializes information the current object
func (m *Win32LobAppRegistryRequirement) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Win32LobAppRequirement.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteBoolValue("check32BitOn64System", m.GetCheck32BitOn64System())
        if err != nil {
            return err
        }
    }
    if m.GetDetectionType() != nil {
        cast := (*m.GetDetectionType()).String()
        err = writer.WriteStringValue("detectionType", &cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("keyPath", m.GetKeyPath())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("@odata.type", m.GetOdataType())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("valueName", m.GetValueName())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetCheck32BitOn64System sets the check32BitOn64System property value. A value indicating whether this registry path is for checking 32-bit app on 64-bit system
func (m *Win32LobAppRegistryRequirement) SetCheck32BitOn64System(value *bool)() {
    err := m.GetBackingStore().Set("check32BitOn64System", value)
    if err != nil {
        panic(err)
    }
}
// SetDetectionType sets the detectionType property value. Contains all supported registry data detection type.
func (m *Win32LobAppRegistryRequirement) SetDetectionType(value *Win32LobAppRegistryDetectionType)() {
    err := m.GetBackingStore().Set("detectionType", value)
    if err != nil {
        panic(err)
    }
}
// SetKeyPath sets the keyPath property value. The registry key path to detect Win32 Line of Business (LoB) app
func (m *Win32LobAppRegistryRequirement) SetKeyPath(value *string)() {
    err := m.GetBackingStore().Set("keyPath", value)
    if err != nil {
        panic(err)
    }
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *Win32LobAppRegistryRequirement) SetOdataType(value *string)() {
    err := m.GetBackingStore().Set("odataType", value)
    if err != nil {
        panic(err)
    }
}
// SetValueName sets the valueName property value. The registry value name
func (m *Win32LobAppRegistryRequirement) SetValueName(value *string)() {
    err := m.GetBackingStore().Set("valueName", value)
    if err != nil {
        panic(err)
    }
}
// Win32LobAppRegistryRequirementable 
type Win32LobAppRegistryRequirementable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    Win32LobAppRequirementable
    GetCheck32BitOn64System()(*bool)
    GetDetectionType()(*Win32LobAppRegistryDetectionType)
    GetKeyPath()(*string)
    GetOdataType()(*string)
    GetValueName()(*string)
    SetCheck32BitOn64System(value *bool)()
    SetDetectionType(value *Win32LobAppRegistryDetectionType)()
    SetKeyPath(value *string)()
    SetOdataType(value *string)()
    SetValueName(value *string)()
}
