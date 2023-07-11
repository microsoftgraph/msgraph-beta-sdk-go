package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// Win32LobAppFileSystemRequirement base class to detect a Win32 App
type Win32LobAppFileSystemRequirement struct {
    Win32LobAppRequirement
    // The OdataType property
    OdataType *string
}
// NewWin32LobAppFileSystemRequirement instantiates a new win32LobAppFileSystemRequirement and sets the default values.
func NewWin32LobAppFileSystemRequirement()(*Win32LobAppFileSystemRequirement) {
    m := &Win32LobAppFileSystemRequirement{
        Win32LobAppRequirement: *NewWin32LobAppRequirement(),
    }
    odataTypeValue := "#microsoft.graph.win32LobAppFileSystemRequirement"
    m.SetOdataType(&odataTypeValue)
    return m
}
// CreateWin32LobAppFileSystemRequirementFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateWin32LobAppFileSystemRequirementFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewWin32LobAppFileSystemRequirement(), nil
}
// GetCheck32BitOn64System gets the check32BitOn64System property value. A value indicating whether this file or folder is for checking 32-bit app on 64-bit system
func (m *Win32LobAppFileSystemRequirement) GetCheck32BitOn64System()(*bool) {
    val, err := m.GetBackingStore().Get("check32BitOn64System")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// GetDetectionType gets the detectionType property value. Contains all supported file system detection type.
func (m *Win32LobAppFileSystemRequirement) GetDetectionType()(*Win32LobAppFileSystemDetectionType) {
    val, err := m.GetBackingStore().Get("detectionType")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*Win32LobAppFileSystemDetectionType)
    }
    return nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *Win32LobAppFileSystemRequirement) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
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
        val, err := n.GetEnumValue(ParseWin32LobAppFileSystemDetectionType)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDetectionType(val.(*Win32LobAppFileSystemDetectionType))
        }
        return nil
    }
    res["fileOrFolderName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetFileOrFolderName(val)
        }
        return nil
    }
    res["path"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPath(val)
        }
        return nil
    }
    return res
}
// GetFileOrFolderName gets the fileOrFolderName property value. The file or folder name to detect Win32 Line of Business (LoB) app
func (m *Win32LobAppFileSystemRequirement) GetFileOrFolderName()(*string) {
    val, err := m.GetBackingStore().Get("fileOrFolderName")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetPath gets the path property value. The file or folder path to detect Win32 Line of Business (LoB) app
func (m *Win32LobAppFileSystemRequirement) GetPath()(*string) {
    val, err := m.GetBackingStore().Get("path")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// Serialize serializes information the current object
func (m *Win32LobAppFileSystemRequirement) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
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
        err = writer.WriteStringValue("fileOrFolderName", m.GetFileOrFolderName())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("path", m.GetPath())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetCheck32BitOn64System sets the check32BitOn64System property value. A value indicating whether this file or folder is for checking 32-bit app on 64-bit system
func (m *Win32LobAppFileSystemRequirement) SetCheck32BitOn64System(value *bool)() {
    err := m.GetBackingStore().Set("check32BitOn64System", value)
    if err != nil {
        panic(err)
    }
}
// SetDetectionType sets the detectionType property value. Contains all supported file system detection type.
func (m *Win32LobAppFileSystemRequirement) SetDetectionType(value *Win32LobAppFileSystemDetectionType)() {
    err := m.GetBackingStore().Set("detectionType", value)
    if err != nil {
        panic(err)
    }
}
// SetFileOrFolderName sets the fileOrFolderName property value. The file or folder name to detect Win32 Line of Business (LoB) app
func (m *Win32LobAppFileSystemRequirement) SetFileOrFolderName(value *string)() {
    err := m.GetBackingStore().Set("fileOrFolderName", value)
    if err != nil {
        panic(err)
    }
}
// SetPath sets the path property value. The file or folder path to detect Win32 Line of Business (LoB) app
func (m *Win32LobAppFileSystemRequirement) SetPath(value *string)() {
    err := m.GetBackingStore().Set("path", value)
    if err != nil {
        panic(err)
    }
}
// Win32LobAppFileSystemRequirementable 
type Win32LobAppFileSystemRequirementable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    Win32LobAppRequirementable
    GetCheck32BitOn64System()(*bool)
    GetDetectionType()(*Win32LobAppFileSystemDetectionType)
    GetFileOrFolderName()(*string)
    GetPath()(*string)
    SetCheck32BitOn64System(value *bool)()
    SetDetectionType(value *Win32LobAppFileSystemDetectionType)()
    SetFileOrFolderName(value *string)()
    SetPath(value *string)()
}
