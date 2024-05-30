package storage

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e "github.com/microsoft/kiota-abstractions-go/store"
)

type FilestorageDeletedcontainersItemDriveItemsItemPreviewPostRequestBody struct {
    // Stores model information.
    backingStore ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore
}
// NewFilestorageDeletedcontainersItemDriveItemsItemPreviewPostRequestBody instantiates a new FilestorageDeletedcontainersItemDriveItemsItemPreviewPostRequestBody and sets the default values.
func NewFilestorageDeletedcontainersItemDriveItemsItemPreviewPostRequestBody()(*FilestorageDeletedcontainersItemDriveItemsItemPreviewPostRequestBody) {
    m := &FilestorageDeletedcontainersItemDriveItemsItemPreviewPostRequestBody{
    }
    m.backingStore = ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStoreFactoryInstance();
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateFilestorageDeletedcontainersItemDriveItemsItemPreviewPostRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateFilestorageDeletedcontainersItemDriveItemsItemPreviewPostRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewFilestorageDeletedcontainersItemDriveItemsItemPreviewPostRequestBody(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *FilestorageDeletedcontainersItemDriveItemsItemPreviewPostRequestBody) GetAdditionalData()(map[string]any) {
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
// GetAllowEdit gets the allowEdit property value. The allowEdit property
// returns a *bool when successful
func (m *FilestorageDeletedcontainersItemDriveItemsItemPreviewPostRequestBody) GetAllowEdit()(*bool) {
    val, err := m.GetBackingStore().Get("allowEdit")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// GetBackingStore gets the BackingStore property value. Stores model information.
// returns a BackingStore when successful
func (m *FilestorageDeletedcontainersItemDriveItemsItemPreviewPostRequestBody) GetBackingStore()(ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore) {
    return m.backingStore
}
// GetChromeless gets the chromeless property value. The chromeless property
// returns a *bool when successful
func (m *FilestorageDeletedcontainersItemDriveItemsItemPreviewPostRequestBody) GetChromeless()(*bool) {
    val, err := m.GetBackingStore().Get("chromeless")
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
func (m *FilestorageDeletedcontainersItemDriveItemsItemPreviewPostRequestBody) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["allowEdit"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAllowEdit(val)
        }
        return nil
    }
    res["chromeless"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetChromeless(val)
        }
        return nil
    }
    res["page"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPage(val)
        }
        return nil
    }
    res["viewer"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetViewer(val)
        }
        return nil
    }
    res["zoom"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetFloat64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetZoom(val)
        }
        return nil
    }
    return res
}
// GetPage gets the page property value. The page property
// returns a *string when successful
func (m *FilestorageDeletedcontainersItemDriveItemsItemPreviewPostRequestBody) GetPage()(*string) {
    val, err := m.GetBackingStore().Get("page")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetViewer gets the viewer property value. The viewer property
// returns a *string when successful
func (m *FilestorageDeletedcontainersItemDriveItemsItemPreviewPostRequestBody) GetViewer()(*string) {
    val, err := m.GetBackingStore().Get("viewer")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetZoom gets the zoom property value. The zoom property
// returns a *float64 when successful
func (m *FilestorageDeletedcontainersItemDriveItemsItemPreviewPostRequestBody) GetZoom()(*float64) {
    val, err := m.GetBackingStore().Get("zoom")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*float64)
    }
    return nil
}
// Serialize serializes information the current object
func (m *FilestorageDeletedcontainersItemDriveItemsItemPreviewPostRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteBoolValue("allowEdit", m.GetAllowEdit())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("chromeless", m.GetChromeless())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("page", m.GetPage())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("viewer", m.GetViewer())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteFloat64Value("zoom", m.GetZoom())
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
func (m *FilestorageDeletedcontainersItemDriveItemsItemPreviewPostRequestBody) SetAdditionalData(value map[string]any)() {
    err := m.GetBackingStore().Set("additionalData", value)
    if err != nil {
        panic(err)
    }
}
// SetAllowEdit sets the allowEdit property value. The allowEdit property
func (m *FilestorageDeletedcontainersItemDriveItemsItemPreviewPostRequestBody) SetAllowEdit(value *bool)() {
    err := m.GetBackingStore().Set("allowEdit", value)
    if err != nil {
        panic(err)
    }
}
// SetBackingStore sets the BackingStore property value. Stores model information.
func (m *FilestorageDeletedcontainersItemDriveItemsItemPreviewPostRequestBody) SetBackingStore(value ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)() {
    m.backingStore = value
}
// SetChromeless sets the chromeless property value. The chromeless property
func (m *FilestorageDeletedcontainersItemDriveItemsItemPreviewPostRequestBody) SetChromeless(value *bool)() {
    err := m.GetBackingStore().Set("chromeless", value)
    if err != nil {
        panic(err)
    }
}
// SetPage sets the page property value. The page property
func (m *FilestorageDeletedcontainersItemDriveItemsItemPreviewPostRequestBody) SetPage(value *string)() {
    err := m.GetBackingStore().Set("page", value)
    if err != nil {
        panic(err)
    }
}
// SetViewer sets the viewer property value. The viewer property
func (m *FilestorageDeletedcontainersItemDriveItemsItemPreviewPostRequestBody) SetViewer(value *string)() {
    err := m.GetBackingStore().Set("viewer", value)
    if err != nil {
        panic(err)
    }
}
// SetZoom sets the zoom property value. The zoom property
func (m *FilestorageDeletedcontainersItemDriveItemsItemPreviewPostRequestBody) SetZoom(value *float64)() {
    err := m.GetBackingStore().Set("zoom", value)
    if err != nil {
        panic(err)
    }
}
type FilestorageDeletedcontainersItemDriveItemsItemPreviewPostRequestBodyable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackedModel
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAllowEdit()(*bool)
    GetBackingStore()(ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)
    GetChromeless()(*bool)
    GetPage()(*string)
    GetViewer()(*string)
    GetZoom()(*float64)
    SetAllowEdit(value *bool)()
    SetBackingStore(value ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)()
    SetChromeless(value *bool)()
    SetPage(value *string)()
    SetViewer(value *string)()
    SetZoom(value *float64)()
}
