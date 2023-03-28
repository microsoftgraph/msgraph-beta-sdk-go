package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e "github.com/microsoft/kiota-abstractions-go/store"
)

// Admin 
type Admin struct {
    // Stores model information.
    backingStore ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore
}
// NewAdmin instantiates a new Admin and sets the default values.
func NewAdmin()(*Admin) {
    m := &Admin{
    }
    m.backingStore = ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStoreFactoryInstance();
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateAdminFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateAdminFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewAdmin(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *Admin) GetAdditionalData()(map[string]any) {
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
// GetBackingStore gets the backingStore property value. Stores model information.
func (m *Admin) GetBackingStore()(ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore) {
    return m.backingStore
}
// GetEdge gets the edge property value. A container for Microsoft Edge resources. Read-only.
func (m *Admin) GetEdge()(Edgeable) {
    val, err := m.GetBackingStore().Get("edge")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(Edgeable)
    }
    return nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *Admin) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["edge"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateEdgeFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetEdge(val.(Edgeable))
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
    res["reportSettings"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateAdminReportSettingsFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetReportSettings(val.(AdminReportSettingsable))
        }
        return nil
    }
    res["serviceAnnouncement"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateServiceAnnouncementFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetServiceAnnouncement(val.(ServiceAnnouncementable))
        }
        return nil
    }
    res["windows"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateAdminWindowsFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetWindows(val.(AdminWindowsable))
        }
        return nil
    }
    return res
}
// GetOdataType gets the @odata.type property value. The OdataType property
func (m *Admin) GetOdataType()(*string) {
    val, err := m.GetBackingStore().Get("odataType")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetReportSettings gets the reportSettings property value. A container for administrative resources to manage reports.
func (m *Admin) GetReportSettings()(AdminReportSettingsable) {
    val, err := m.GetBackingStore().Get("reportSettings")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(AdminReportSettingsable)
    }
    return nil
}
// GetServiceAnnouncement gets the serviceAnnouncement property value. A container for service communications resources. Read-only.
func (m *Admin) GetServiceAnnouncement()(ServiceAnnouncementable) {
    val, err := m.GetBackingStore().Get("serviceAnnouncement")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(ServiceAnnouncementable)
    }
    return nil
}
// GetWindows gets the windows property value. A container for all Windows administrator functionalities. Read-only.
func (m *Admin) GetWindows()(AdminWindowsable) {
    val, err := m.GetBackingStore().Get("windows")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(AdminWindowsable)
    }
    return nil
}
// Serialize serializes information the current object
func (m *Admin) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteObjectValue("edge", m.GetEdge())
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
        err := writer.WriteObjectValue("reportSettings", m.GetReportSettings())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("serviceAnnouncement", m.GetServiceAnnouncement())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("windows", m.GetWindows())
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
// SetAdditionalData sets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *Admin) SetAdditionalData(value map[string]any)() {
    err := m.GetBackingStore().Set("additionalData", value)
    if err != nil {
        panic(err)
    }
}
// SetBackingStore sets the backingStore property value. Stores model information.
func (m *Admin) SetBackingStore(value ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)() {
    m.backingStore = value
}
// SetEdge sets the edge property value. A container for Microsoft Edge resources. Read-only.
func (m *Admin) SetEdge(value Edgeable)() {
    err := m.GetBackingStore().Set("edge", value)
    if err != nil {
        panic(err)
    }
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *Admin) SetOdataType(value *string)() {
    err := m.GetBackingStore().Set("odataType", value)
    if err != nil {
        panic(err)
    }
}
// SetReportSettings sets the reportSettings property value. A container for administrative resources to manage reports.
func (m *Admin) SetReportSettings(value AdminReportSettingsable)() {
    err := m.GetBackingStore().Set("reportSettings", value)
    if err != nil {
        panic(err)
    }
}
// SetServiceAnnouncement sets the serviceAnnouncement property value. A container for service communications resources. Read-only.
func (m *Admin) SetServiceAnnouncement(value ServiceAnnouncementable)() {
    err := m.GetBackingStore().Set("serviceAnnouncement", value)
    if err != nil {
        panic(err)
    }
}
// SetWindows sets the windows property value. A container for all Windows administrator functionalities. Read-only.
func (m *Admin) SetWindows(value AdminWindowsable)() {
    err := m.GetBackingStore().Set("windows", value)
    if err != nil {
        panic(err)
    }
}
// Adminable 
type Adminable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackedModel
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetBackingStore()(ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)
    GetEdge()(Edgeable)
    GetOdataType()(*string)
    GetReportSettings()(AdminReportSettingsable)
    GetServiceAnnouncement()(ServiceAnnouncementable)
    GetWindows()(AdminWindowsable)
    SetBackingStore(value ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)()
    SetEdge(value Edgeable)()
    SetOdataType(value *string)()
    SetReportSettings(value AdminReportSettingsable)()
    SetServiceAnnouncement(value ServiceAnnouncementable)()
    SetWindows(value AdminWindowsable)()
}
// AdminWindows 
type AdminWindows struct {
    Entity
}
// NewAdminWindows instantiates a new AdminWindows and sets the default values.
func NewAdminWindows()(*AdminWindows) {
    m := &AdminWindows{
        Entity: *NewEntity(),
    }
    return m
}
// CreateAdminWindowsFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateAdminWindowsFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewAdminWindows(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *AdminWindows) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["updates"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateAdminWindowsUpdatesFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUpdates(val.(AdminWindowsUpdatesable))
        }
        return nil
    }
    return res
}
// GetUpdates gets the updates property value. Entity that acts as a container for all Windows Update for Business deployment service functionalities. Read-only.
func (m *AdminWindows) GetUpdates()(AdminWindowsUpdatesable) {
    val, err := m.GetBackingStore().Get("updates")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(AdminWindowsUpdatesable)
    }
    return nil
}
// Serialize serializes information the current object
func (m *AdminWindows) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteObjectValue("updates", m.GetUpdates())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetUpdates sets the updates property value. Entity that acts as a container for all Windows Update for Business deployment service functionalities. Read-only.
func (m *AdminWindows) SetUpdates(value AdminWindowsUpdatesable)() {
    err := m.GetBackingStore().Set("updates", value)
    if err != nil {
        panic(err)
    }
}
// AdminWindowsable 
type AdminWindowsable interface {
    Entityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetUpdates()(AdminWindowsUpdatesable)
    SetUpdates(value AdminWindowsUpdatesable)()
}
