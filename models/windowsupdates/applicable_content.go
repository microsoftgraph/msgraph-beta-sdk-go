package windowsupdates

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e "github.com/microsoft/kiota-abstractions-go/store"
)

type ApplicableContent struct {
    // Stores model information.
    backingStore ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore
}
// NewApplicableContent instantiates a new ApplicableContent and sets the default values.
func NewApplicableContent()(*ApplicableContent) {
    m := &ApplicableContent{
    }
    m.backingStore = ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStoreFactoryInstance();
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateApplicableContentFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateApplicableContentFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewApplicableContent(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *ApplicableContent) GetAdditionalData()(map[string]any) {
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
func (m *ApplicableContent) GetBackingStore()(ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore) {
    return m.backingStore
}
// GetCatalogEntry gets the catalogEntry property value. Catalog entry for the update or content.
// returns a CatalogEntryable when successful
func (m *ApplicableContent) GetCatalogEntry()(CatalogEntryable) {
    val, err := m.GetBackingStore().Get("catalogEntry")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(CatalogEntryable)
    }
    return nil
}
// GetCatalogEntryId gets the catalogEntryId property value. ID of the catalog entry for the applicable content.
// returns a *string when successful
func (m *ApplicableContent) GetCatalogEntryId()(*string) {
    val, err := m.GetBackingStore().Get("catalogEntryId")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *ApplicableContent) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["catalogEntry"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateCatalogEntryFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCatalogEntry(val.(CatalogEntryable))
        }
        return nil
    }
    res["catalogEntryId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCatalogEntryId(val)
        }
        return nil
    }
    res["matchedDevices"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateApplicableContentDeviceMatchFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]ApplicableContentDeviceMatchable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(ApplicableContentDeviceMatchable)
                }
            }
            m.SetMatchedDevices(res)
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
// GetMatchedDevices gets the matchedDevices property value. Collection of devices and recommendations for applicable catalog content.
// returns a []ApplicableContentDeviceMatchable when successful
func (m *ApplicableContent) GetMatchedDevices()([]ApplicableContentDeviceMatchable) {
    val, err := m.GetBackingStore().Get("matchedDevices")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]ApplicableContentDeviceMatchable)
    }
    return nil
}
// GetOdataType gets the @odata.type property value. The OdataType property
// returns a *string when successful
func (m *ApplicableContent) GetOdataType()(*string) {
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
func (m *ApplicableContent) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteObjectValue("catalogEntry", m.GetCatalogEntry())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("catalogEntryId", m.GetCatalogEntryId())
        if err != nil {
            return err
        }
    }
    if m.GetMatchedDevices() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetMatchedDevices()))
        for i, v := range m.GetMatchedDevices() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err := writer.WriteCollectionOfObjectValues("matchedDevices", cast)
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
func (m *ApplicableContent) SetAdditionalData(value map[string]any)() {
    err := m.GetBackingStore().Set("additionalData", value)
    if err != nil {
        panic(err)
    }
}
// SetBackingStore sets the BackingStore property value. Stores model information.
func (m *ApplicableContent) SetBackingStore(value ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)() {
    m.backingStore = value
}
// SetCatalogEntry sets the catalogEntry property value. Catalog entry for the update or content.
func (m *ApplicableContent) SetCatalogEntry(value CatalogEntryable)() {
    err := m.GetBackingStore().Set("catalogEntry", value)
    if err != nil {
        panic(err)
    }
}
// SetCatalogEntryId sets the catalogEntryId property value. ID of the catalog entry for the applicable content.
func (m *ApplicableContent) SetCatalogEntryId(value *string)() {
    err := m.GetBackingStore().Set("catalogEntryId", value)
    if err != nil {
        panic(err)
    }
}
// SetMatchedDevices sets the matchedDevices property value. Collection of devices and recommendations for applicable catalog content.
func (m *ApplicableContent) SetMatchedDevices(value []ApplicableContentDeviceMatchable)() {
    err := m.GetBackingStore().Set("matchedDevices", value)
    if err != nil {
        panic(err)
    }
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *ApplicableContent) SetOdataType(value *string)() {
    err := m.GetBackingStore().Set("odataType", value)
    if err != nil {
        panic(err)
    }
}
type ApplicableContentable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackedModel
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetBackingStore()(ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)
    GetCatalogEntry()(CatalogEntryable)
    GetCatalogEntryId()(*string)
    GetMatchedDevices()([]ApplicableContentDeviceMatchable)
    GetOdataType()(*string)
    SetBackingStore(value ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)()
    SetCatalogEntry(value CatalogEntryable)()
    SetCatalogEntryId(value *string)()
    SetMatchedDevices(value []ApplicableContentDeviceMatchable)()
    SetOdataType(value *string)()
}
