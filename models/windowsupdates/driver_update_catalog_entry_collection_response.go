package windowsupdates

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
)

// DriverUpdateCatalogEntryCollectionResponse 
type DriverUpdateCatalogEntryCollectionResponse struct {
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.BaseCollectionPaginationCountResponse
}
// NewDriverUpdateCatalogEntryCollectionResponse instantiates a new DriverUpdateCatalogEntryCollectionResponse and sets the default values.
func NewDriverUpdateCatalogEntryCollectionResponse()(*DriverUpdateCatalogEntryCollectionResponse) {
    m := &DriverUpdateCatalogEntryCollectionResponse{
        BaseCollectionPaginationCountResponse: *ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.NewBaseCollectionPaginationCountResponse(),
    }
    return m
}
// CreateDriverUpdateCatalogEntryCollectionResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateDriverUpdateCatalogEntryCollectionResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewDriverUpdateCatalogEntryCollectionResponse(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *DriverUpdateCatalogEntryCollectionResponse) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.BaseCollectionPaginationCountResponse.GetFieldDeserializers()
    res["value"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateDriverUpdateCatalogEntryFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]DriverUpdateCatalogEntryable, len(val))
            for i, v := range val {
                res[i] = v.(DriverUpdateCatalogEntryable)
            }
            m.SetValue(res)
        }
        return nil
    }
    return res
}
// GetValue gets the value property value. The value property
func (m *DriverUpdateCatalogEntryCollectionResponse) GetValue()([]DriverUpdateCatalogEntryable) {
    val, err := m.GetBackingStore().Get("value")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]DriverUpdateCatalogEntryable)
    }
    return nil
}
// Serialize serializes information the current object
func (m *DriverUpdateCatalogEntryCollectionResponse) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.BaseCollectionPaginationCountResponse.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetValue() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetValue()))
        for i, v := range m.GetValue() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("value", cast)
        if err != nil {
            return err
        }
    }
    return nil
}
// SetValue sets the value property value. The value property
func (m *DriverUpdateCatalogEntryCollectionResponse) SetValue(value []DriverUpdateCatalogEntryable)() {
    err := m.GetBackingStore().Set("value", value)
    if err != nil {
        panic(err)
    }
}
// DriverUpdateCatalogEntryCollectionResponseable 
type DriverUpdateCatalogEntryCollectionResponseable interface {
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.BaseCollectionPaginationCountResponseable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetValue()([]DriverUpdateCatalogEntryable)
    SetValue(value []DriverUpdateCatalogEntryable)()
}
