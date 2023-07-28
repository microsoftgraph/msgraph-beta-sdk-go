package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// Win32LobAppPowerShellScriptRequirement contains PowerShell script properties to detect a Win32 App
type Win32LobAppPowerShellScriptRequirement struct {
    Win32LobAppRequirement
}
// NewWin32LobAppPowerShellScriptRequirement instantiates a new win32LobAppPowerShellScriptRequirement and sets the default values.
func NewWin32LobAppPowerShellScriptRequirement()(*Win32LobAppPowerShellScriptRequirement) {
    m := &Win32LobAppPowerShellScriptRequirement{
        Win32LobAppRequirement: *NewWin32LobAppRequirement(),
    }
    odataTypeValue := "#microsoft.graph.win32LobAppPowerShellScriptRequirement"
    m.SetOdataType(&odataTypeValue)
    return m
}
// CreateWin32LobAppPowerShellScriptRequirementFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateWin32LobAppPowerShellScriptRequirementFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewWin32LobAppPowerShellScriptRequirement(), nil
}
// GetDetectionType gets the detectionType property value. Contains all supported Powershell Script output detection type.
func (m *Win32LobAppPowerShellScriptRequirement) GetDetectionType()(*Win32LobAppPowerShellScriptDetectionType) {
    val, err := m.GetBackingStore().Get("detectionType")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*Win32LobAppPowerShellScriptDetectionType)
    }
    return nil
}
// GetDisplayName gets the displayName property value. The unique display name for this rule
func (m *Win32LobAppPowerShellScriptRequirement) GetDisplayName()(*string) {
    val, err := m.GetBackingStore().Get("displayName")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetEnforceSignatureCheck gets the enforceSignatureCheck property value. A value indicating whether signature check is enforced
func (m *Win32LobAppPowerShellScriptRequirement) GetEnforceSignatureCheck()(*bool) {
    val, err := m.GetBackingStore().Get("enforceSignatureCheck")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *Win32LobAppPowerShellScriptRequirement) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Win32LobAppRequirement.GetFieldDeserializers()
    res["detectionType"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseWin32LobAppPowerShellScriptDetectionType)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDetectionType(val.(*Win32LobAppPowerShellScriptDetectionType))
        }
        return nil
    }
    res["displayName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDisplayName(val)
        }
        return nil
    }
    res["enforceSignatureCheck"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetEnforceSignatureCheck(val)
        }
        return nil
    }
    res["runAs32Bit"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRunAs32Bit(val)
        }
        return nil
    }
    res["runAsAccount"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseRunAsAccountType)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRunAsAccount(val.(*RunAsAccountType))
        }
        return nil
    }
    res["scriptContent"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetScriptContent(val)
        }
        return nil
    }
    return res
}
// GetRunAs32Bit gets the runAs32Bit property value. A value indicating whether this script should run as 32-bit
func (m *Win32LobAppPowerShellScriptRequirement) GetRunAs32Bit()(*bool) {
    val, err := m.GetBackingStore().Get("runAs32Bit")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// GetRunAsAccount gets the runAsAccount property value. Indicates the type of execution context the app runs in.
func (m *Win32LobAppPowerShellScriptRequirement) GetRunAsAccount()(*RunAsAccountType) {
    val, err := m.GetBackingStore().Get("runAsAccount")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*RunAsAccountType)
    }
    return nil
}
// GetScriptContent gets the scriptContent property value. The base64 encoded script content to detect Win32 Line of Business (LoB) app
func (m *Win32LobAppPowerShellScriptRequirement) GetScriptContent()(*string) {
    val, err := m.GetBackingStore().Get("scriptContent")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// Serialize serializes information the current object
func (m *Win32LobAppPowerShellScriptRequirement) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Win32LobAppRequirement.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetDetectionType() != nil {
        cast := (*m.GetDetectionType()).String()
        err = writer.WriteStringValue("detectionType", &cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("displayName", m.GetDisplayName())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("enforceSignatureCheck", m.GetEnforceSignatureCheck())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("runAs32Bit", m.GetRunAs32Bit())
        if err != nil {
            return err
        }
    }
    if m.GetRunAsAccount() != nil {
        cast := (*m.GetRunAsAccount()).String()
        err = writer.WriteStringValue("runAsAccount", &cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("scriptContent", m.GetScriptContent())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetDetectionType sets the detectionType property value. Contains all supported Powershell Script output detection type.
func (m *Win32LobAppPowerShellScriptRequirement) SetDetectionType(value *Win32LobAppPowerShellScriptDetectionType)() {
    err := m.GetBackingStore().Set("detectionType", value)
    if err != nil {
        panic(err)
    }
}
// SetDisplayName sets the displayName property value. The unique display name for this rule
func (m *Win32LobAppPowerShellScriptRequirement) SetDisplayName(value *string)() {
    err := m.GetBackingStore().Set("displayName", value)
    if err != nil {
        panic(err)
    }
}
// SetEnforceSignatureCheck sets the enforceSignatureCheck property value. A value indicating whether signature check is enforced
func (m *Win32LobAppPowerShellScriptRequirement) SetEnforceSignatureCheck(value *bool)() {
    err := m.GetBackingStore().Set("enforceSignatureCheck", value)
    if err != nil {
        panic(err)
    }
}
// SetRunAs32Bit sets the runAs32Bit property value. A value indicating whether this script should run as 32-bit
func (m *Win32LobAppPowerShellScriptRequirement) SetRunAs32Bit(value *bool)() {
    err := m.GetBackingStore().Set("runAs32Bit", value)
    if err != nil {
        panic(err)
    }
}
// SetRunAsAccount sets the runAsAccount property value. Indicates the type of execution context the app runs in.
func (m *Win32LobAppPowerShellScriptRequirement) SetRunAsAccount(value *RunAsAccountType)() {
    err := m.GetBackingStore().Set("runAsAccount", value)
    if err != nil {
        panic(err)
    }
}
// SetScriptContent sets the scriptContent property value. The base64 encoded script content to detect Win32 Line of Business (LoB) app
func (m *Win32LobAppPowerShellScriptRequirement) SetScriptContent(value *string)() {
    err := m.GetBackingStore().Set("scriptContent", value)
    if err != nil {
        panic(err)
    }
}
// Win32LobAppPowerShellScriptRequirementable 
type Win32LobAppPowerShellScriptRequirementable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    Win32LobAppRequirementable
    GetDetectionType()(*Win32LobAppPowerShellScriptDetectionType)
    GetDisplayName()(*string)
    GetEnforceSignatureCheck()(*bool)
    GetRunAs32Bit()(*bool)
    GetRunAsAccount()(*RunAsAccountType)
    GetScriptContent()(*string)
    SetDetectionType(value *Win32LobAppPowerShellScriptDetectionType)()
    SetDisplayName(value *string)()
    SetEnforceSignatureCheck(value *bool)()
    SetRunAs32Bit(value *bool)()
    SetRunAsAccount(value *RunAsAccountType)()
    SetScriptContent(value *string)()
}
