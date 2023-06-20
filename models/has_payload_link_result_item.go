package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e "github.com/microsoft/kiota-abstractions-go/store"
)

// HasPayloadLinkResultItem a class containing the result of HasPayloadLinks action.
type HasPayloadLinkResultItem struct {
    // Stores model information.
    backingStore ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore
}
// NewHasPayloadLinkResultItem instantiates a new hasPayloadLinkResultItem and sets the default values.
func NewHasPayloadLinkResultItem()(*HasPayloadLinkResultItem) {
    m := &HasPayloadLinkResultItem{
    }
    m.backingStore = ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStoreFactoryInstance();
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateHasPayloadLinkResultItemFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateHasPayloadLinkResultItemFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewHasPayloadLinkResultItem(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *HasPayloadLinkResultItem) GetAdditionalData()(map[string]any) {
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
func (m *HasPayloadLinkResultItem) GetBackingStore()(ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore) {
    return m.backingStore
}
// GetError gets the error property value. Exception information indicates if check for this item was successful or not.Empty string for no error.
func (m *HasPayloadLinkResultItem) GetError()(*string) {
    val, err := m.GetBackingStore().Get("error")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *HasPayloadLinkResultItem) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["error"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetError(val)
        }
        return nil
    }
    res["hasLink"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetHasLink(val)
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
    res["payloadId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPayloadId(val)
        }
        return nil
    }
    res["sources"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfEnumValues(ParseDeviceAndAppManagementAssignmentSource)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]DeviceAndAppManagementAssignmentSource, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = *(v.(*DeviceAndAppManagementAssignmentSource))
                }
            }
            m.SetSources(res)
        }
        return nil
    }
    return res
}
// GetHasLink gets the hasLink property value. Indicate whether a payload has any link or not.
func (m *HasPayloadLinkResultItem) GetHasLink()(*bool) {
    val, err := m.GetBackingStore().Get("hasLink")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// GetOdataType gets the @odata.type property value. The OdataType property
func (m *HasPayloadLinkResultItem) GetOdataType()(*string) {
    val, err := m.GetBackingStore().Get("odataType")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetPayloadId gets the payloadId property value. Key of the Payload, In the format of Guid.
func (m *HasPayloadLinkResultItem) GetPayloadId()(*string) {
    val, err := m.GetBackingStore().Get("payloadId")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetSources gets the sources property value. The reason where the link comes from.
func (m *HasPayloadLinkResultItem) GetSources()([]DeviceAndAppManagementAssignmentSource) {
    val, err := m.GetBackingStore().Get("sources")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]DeviceAndAppManagementAssignmentSource)
    }
    return nil
}
// Serialize serializes information the current object
func (m *HasPayloadLinkResultItem) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("error", m.GetError())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("hasLink", m.GetHasLink())
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
        err := writer.WriteStringValue("payloadId", m.GetPayloadId())
        if err != nil {
            return err
        }
    }
    if m.GetSources() != nil {
        err := writer.WriteCollectionOfStringValues("sources", SerializeDeviceAndAppManagementAssignmentSource(m.GetSources()))
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
func (m *HasPayloadLinkResultItem) SetAdditionalData(value map[string]any)() {
    err := m.GetBackingStore().Set("additionalData", value)
    if err != nil {
        panic(err)
    }
}
// SetBackingStore sets the backingStore property value. Stores model information.
func (m *HasPayloadLinkResultItem) SetBackingStore(value ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)() {
    m.backingStore = value
}
// SetError sets the error property value. Exception information indicates if check for this item was successful or not.Empty string for no error.
func (m *HasPayloadLinkResultItem) SetError(value *string)() {
    err := m.GetBackingStore().Set("error", value)
    if err != nil {
        panic(err)
    }
}
// SetHasLink sets the hasLink property value. Indicate whether a payload has any link or not.
func (m *HasPayloadLinkResultItem) SetHasLink(value *bool)() {
    err := m.GetBackingStore().Set("hasLink", value)
    if err != nil {
        panic(err)
    }
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *HasPayloadLinkResultItem) SetOdataType(value *string)() {
    err := m.GetBackingStore().Set("odataType", value)
    if err != nil {
        panic(err)
    }
}
// SetPayloadId sets the payloadId property value. Key of the Payload, In the format of Guid.
func (m *HasPayloadLinkResultItem) SetPayloadId(value *string)() {
    err := m.GetBackingStore().Set("payloadId", value)
    if err != nil {
        panic(err)
    }
}
// SetSources sets the sources property value. The reason where the link comes from.
func (m *HasPayloadLinkResultItem) SetSources(value []DeviceAndAppManagementAssignmentSource)() {
    err := m.GetBackingStore().Set("sources", value)
    if err != nil {
        panic(err)
    }
}
// HasPayloadLinkResultItemable 
type HasPayloadLinkResultItemable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackedModel
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetBackingStore()(ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)
    GetError()(*string)
    GetHasLink()(*bool)
    GetOdataType()(*string)
    GetPayloadId()(*string)
    GetSources()([]DeviceAndAppManagementAssignmentSource)
    SetBackingStore(value ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)()
    SetError(value *string)()
    SetHasLink(value *bool)()
    SetOdataType(value *string)()
    SetPayloadId(value *string)()
    SetSources(value []DeviceAndAppManagementAssignmentSource)()
}
