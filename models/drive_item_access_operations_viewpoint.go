package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e "github.com/microsoft/kiota-abstractions-go/store"
)

type DriveItemAccessOperationsViewpoint struct {
    // Stores model information.
    backingStore ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore
}
// NewDriveItemAccessOperationsViewpoint instantiates a new DriveItemAccessOperationsViewpoint and sets the default values.
func NewDriveItemAccessOperationsViewpoint()(*DriveItemAccessOperationsViewpoint) {
    m := &DriveItemAccessOperationsViewpoint{
    }
    m.backingStore = ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStoreFactoryInstance();
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateDriveItemAccessOperationsViewpointFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateDriveItemAccessOperationsViewpointFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewDriveItemAccessOperationsViewpoint(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *DriveItemAccessOperationsViewpoint) GetAdditionalData()(map[string]any) {
    val , err :=  m.backingStore.Get("additionalData")
    if err != nil {
        panic(err)
    }
    if val == nil {
        var value = make(map[string]any);
        m.SetAdditionalData(value);
    }
    return val.(map[string]any)
}
// GetBackingStore gets the BackingStore property value. Stores model information.
// returns a BackingStore when successful
func (m *DriveItemAccessOperationsViewpoint) GetBackingStore()(ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore) {
    return m.backingStore
}
// GetCanComment gets the canComment property value. Indicates whether the user can comment on this item.
// returns a *bool when successful
func (m *DriveItemAccessOperationsViewpoint) GetCanComment()(*bool) {
    val, err := m.GetBackingStore().Get("canComment")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// GetCanCreateFile gets the canCreateFile property value. Indicates whether the user can create files within this object. Returned only on folders.
// returns a *bool when successful
func (m *DriveItemAccessOperationsViewpoint) GetCanCreateFile()(*bool) {
    val, err := m.GetBackingStore().Get("canCreateFile")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// GetCanCreateFolder gets the canCreateFolder property value. Indicates whether the user can create folders within this object. Returned only on folders.
// returns a *bool when successful
func (m *DriveItemAccessOperationsViewpoint) GetCanCreateFolder()(*bool) {
    val, err := m.GetBackingStore().Get("canCreateFolder")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// GetCanDelete gets the canDelete property value. Indicates whether the user can delete this item.
// returns a *bool when successful
func (m *DriveItemAccessOperationsViewpoint) GetCanDelete()(*bool) {
    val, err := m.GetBackingStore().Get("canDelete")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// GetCanDownload gets the canDownload property value. Indicates whether the user can download this item.
// returns a *bool when successful
func (m *DriveItemAccessOperationsViewpoint) GetCanDownload()(*bool) {
    val, err := m.GetBackingStore().Get("canDownload")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// GetCanRead gets the canRead property value. Indicates whether the user can read this item.
// returns a *bool when successful
func (m *DriveItemAccessOperationsViewpoint) GetCanRead()(*bool) {
    val, err := m.GetBackingStore().Get("canRead")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// GetCanUpdate gets the canUpdate property value. Indicates whether the user can update this item.
// returns a *bool when successful
func (m *DriveItemAccessOperationsViewpoint) GetCanUpdate()(*bool) {
    val, err := m.GetBackingStore().Get("canUpdate")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *DriveItemAccessOperationsViewpoint) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["canComment"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCanComment(val)
        }
        return nil
    }
    res["canCreateFile"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCanCreateFile(val)
        }
        return nil
    }
    res["canCreateFolder"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCanCreateFolder(val)
        }
        return nil
    }
    res["canDelete"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCanDelete(val)
        }
        return nil
    }
    res["canDownload"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCanDownload(val)
        }
        return nil
    }
    res["canRead"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCanRead(val)
        }
        return nil
    }
    res["canUpdate"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCanUpdate(val)
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
    return res
}
// GetOdataType gets the @odata.type property value. The OdataType property
// returns a *string when successful
func (m *DriveItemAccessOperationsViewpoint) GetOdataType()(*string) {
    val, err := m.GetBackingStore().Get("odataType")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// Serialize serializes information the current object
func (m *DriveItemAccessOperationsViewpoint) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteBoolValue("canComment", m.GetCanComment())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("canCreateFile", m.GetCanCreateFile())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("canCreateFolder", m.GetCanCreateFolder())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("canDelete", m.GetCanDelete())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("canDownload", m.GetCanDownload())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("canRead", m.GetCanRead())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("canUpdate", m.GetCanUpdate())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("@odata.type", m.GetOdataType())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteAdditionalData(m.GetAdditionalData())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAdditionalData sets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *DriveItemAccessOperationsViewpoint) SetAdditionalData(value map[string]any)() {
    err := m.GetBackingStore().Set("additionalData", value)
    if err != nil {
        panic(err)
    }
}
// SetBackingStore sets the BackingStore property value. Stores model information.
func (m *DriveItemAccessOperationsViewpoint) SetBackingStore(value ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)() {
    m.backingStore = value
}
// SetCanComment sets the canComment property value. Indicates whether the user can comment on this item.
func (m *DriveItemAccessOperationsViewpoint) SetCanComment(value *bool)() {
    err := m.GetBackingStore().Set("canComment", value)
    if err != nil {
        panic(err)
    }
}
// SetCanCreateFile sets the canCreateFile property value. Indicates whether the user can create files within this object. Returned only on folders.
func (m *DriveItemAccessOperationsViewpoint) SetCanCreateFile(value *bool)() {
    err := m.GetBackingStore().Set("canCreateFile", value)
    if err != nil {
        panic(err)
    }
}
// SetCanCreateFolder sets the canCreateFolder property value. Indicates whether the user can create folders within this object. Returned only on folders.
func (m *DriveItemAccessOperationsViewpoint) SetCanCreateFolder(value *bool)() {
    err := m.GetBackingStore().Set("canCreateFolder", value)
    if err != nil {
        panic(err)
    }
}
// SetCanDelete sets the canDelete property value. Indicates whether the user can delete this item.
func (m *DriveItemAccessOperationsViewpoint) SetCanDelete(value *bool)() {
    err := m.GetBackingStore().Set("canDelete", value)
    if err != nil {
        panic(err)
    }
}
// SetCanDownload sets the canDownload property value. Indicates whether the user can download this item.
func (m *DriveItemAccessOperationsViewpoint) SetCanDownload(value *bool)() {
    err := m.GetBackingStore().Set("canDownload", value)
    if err != nil {
        panic(err)
    }
}
// SetCanRead sets the canRead property value. Indicates whether the user can read this item.
func (m *DriveItemAccessOperationsViewpoint) SetCanRead(value *bool)() {
    err := m.GetBackingStore().Set("canRead", value)
    if err != nil {
        panic(err)
    }
}
// SetCanUpdate sets the canUpdate property value. Indicates whether the user can update this item.
func (m *DriveItemAccessOperationsViewpoint) SetCanUpdate(value *bool)() {
    err := m.GetBackingStore().Set("canUpdate", value)
    if err != nil {
        panic(err)
    }
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *DriveItemAccessOperationsViewpoint) SetOdataType(value *string)() {
    err := m.GetBackingStore().Set("odataType", value)
    if err != nil {
        panic(err)
    }
}
type DriveItemAccessOperationsViewpointable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackedModel
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetBackingStore()(ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)
    GetCanComment()(*bool)
    GetCanCreateFile()(*bool)
    GetCanCreateFolder()(*bool)
    GetCanDelete()(*bool)
    GetCanDownload()(*bool)
    GetCanRead()(*bool)
    GetCanUpdate()(*bool)
    GetOdataType()(*string)
    SetBackingStore(value ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)()
    SetCanComment(value *bool)()
    SetCanCreateFile(value *bool)()
    SetCanCreateFolder(value *bool)()
    SetCanDelete(value *bool)()
    SetCanDownload(value *bool)()
    SetCanRead(value *bool)()
    SetCanUpdate(value *bool)()
    SetOdataType(value *string)()
}
