package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// Win32LobAppFileSystemDetection contains file or folder path to detect a Win32 App
type Win32LobAppFileSystemDetection struct {
    Win32LobAppDetection
}
// NewWin32LobAppFileSystemDetection instantiates a new win32LobAppFileSystemDetection and sets the default values.
func NewWin32LobAppFileSystemDetection()(*Win32LobAppFileSystemDetection) {
    m := &Win32LobAppFileSystemDetection{
        Win32LobAppDetection: *NewWin32LobAppDetection(),
    }
    odataTypeValue := "#microsoft.graph.win32LobAppFileSystemDetection"
    m.SetOdataType(&odataTypeValue)
    return m
}
// CreateWin32LobAppFileSystemDetectionFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateWin32LobAppFileSystemDetectionFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewWin32LobAppFileSystemDetection(), nil
}
// GetCheck32BitOn64System gets the check32BitOn64System property value. A value indicating whether this file or folder is for checking 32-bit app on 64-bit system
func (m *Win32LobAppFileSystemDetection) GetCheck32BitOn64System()(*bool) {
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
func (m *Win32LobAppFileSystemDetection) GetDetectionType()(*Win32LobAppFileSystemDetectionType) {
    val, err := m.GetBackingStore().Get("detectionType")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*Win32LobAppFileSystemDetectionType)
    }
    return nil
}
// GetDetectionValue gets the detectionValue property value. The file or folder detection value
func (m *Win32LobAppFileSystemDetection) GetDetectionValue()(*string) {
    val, err := m.GetBackingStore().Get("detectionValue")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *Win32LobAppFileSystemDetection) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Win32LobAppDetection.GetFieldDeserializers()
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
    res["detectionValue"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDetectionValue(val)
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
    res["operator"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseWin32LobAppDetectionOperator)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOperator(val.(*Win32LobAppDetectionOperator))
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
func (m *Win32LobAppFileSystemDetection) GetFileOrFolderName()(*string) {
    val, err := m.GetBackingStore().Get("fileOrFolderName")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetOdataType gets the @odata.type property value. The OdataType property
func (m *Win32LobAppFileSystemDetection) GetOdataType()(*string) {
    val, err := m.GetBackingStore().Get("odataType")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetOperator gets the operator property value. Contains properties for detection operator.
func (m *Win32LobAppFileSystemDetection) GetOperator()(*Win32LobAppDetectionOperator) {
    val, err := m.GetBackingStore().Get("operator")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*Win32LobAppDetectionOperator)
    }
    return nil
}
// GetPath gets the path property value. The file or folder path to detect Win32 Line of Business (LoB) app
func (m *Win32LobAppFileSystemDetection) GetPath()(*string) {
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
func (m *Win32LobAppFileSystemDetection) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Win32LobAppDetection.Serialize(writer)
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
        err = writer.WriteStringValue("detectionValue", m.GetDetectionValue())
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
        err = writer.WriteStringValue("@odata.type", m.GetOdataType())
        if err != nil {
            return err
        }
    }
    if m.GetOperator() != nil {
        cast := (*m.GetOperator()).String()
        err = writer.WriteStringValue("operator", &cast)
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
func (m *Win32LobAppFileSystemDetection) SetCheck32BitOn64System(value *bool)() {
    err := m.GetBackingStore().Set("check32BitOn64System", value)
    if err != nil {
        panic(err)
    }
}
// SetDetectionType sets the detectionType property value. Contains all supported file system detection type.
func (m *Win32LobAppFileSystemDetection) SetDetectionType(value *Win32LobAppFileSystemDetectionType)() {
    err := m.GetBackingStore().Set("detectionType", value)
    if err != nil {
        panic(err)
    }
}
// SetDetectionValue sets the detectionValue property value. The file or folder detection value
func (m *Win32LobAppFileSystemDetection) SetDetectionValue(value *string)() {
    err := m.GetBackingStore().Set("detectionValue", value)
    if err != nil {
        panic(err)
    }
}
// SetFileOrFolderName sets the fileOrFolderName property value. The file or folder name to detect Win32 Line of Business (LoB) app
func (m *Win32LobAppFileSystemDetection) SetFileOrFolderName(value *string)() {
    err := m.GetBackingStore().Set("fileOrFolderName", value)
    if err != nil {
        panic(err)
    }
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *Win32LobAppFileSystemDetection) SetOdataType(value *string)() {
    err := m.GetBackingStore().Set("odataType", value)
    if err != nil {
        panic(err)
    }
}
// SetOperator sets the operator property value. Contains properties for detection operator.
func (m *Win32LobAppFileSystemDetection) SetOperator(value *Win32LobAppDetectionOperator)() {
    err := m.GetBackingStore().Set("operator", value)
    if err != nil {
        panic(err)
    }
}
// SetPath sets the path property value. The file or folder path to detect Win32 Line of Business (LoB) app
func (m *Win32LobAppFileSystemDetection) SetPath(value *string)() {
    err := m.GetBackingStore().Set("path", value)
    if err != nil {
        panic(err)
    }
}
// Win32LobAppFileSystemDetectionable 
type Win32LobAppFileSystemDetectionable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    Win32LobAppDetectionable
    GetCheck32BitOn64System()(*bool)
    GetDetectionType()(*Win32LobAppFileSystemDetectionType)
    GetDetectionValue()(*string)
    GetFileOrFolderName()(*string)
    GetOdataType()(*string)
    GetOperator()(*Win32LobAppDetectionOperator)
    GetPath()(*string)
    SetCheck32BitOn64System(value *bool)()
    SetDetectionType(value *Win32LobAppFileSystemDetectionType)()
    SetDetectionValue(value *string)()
    SetFileOrFolderName(value *string)()
    SetOdataType(value *string)()
    SetOperator(value *Win32LobAppDetectionOperator)()
    SetPath(value *string)()
}
