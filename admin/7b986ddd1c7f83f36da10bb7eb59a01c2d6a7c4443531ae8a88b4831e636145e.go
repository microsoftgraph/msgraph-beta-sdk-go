// Code generated by Microsoft Kiota - DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package admin

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i17376df570f19ff3c32da2d66a677d31250ed0ff64059351645f48a152316b3c "github.com/microsoftgraph/msgraph-beta-sdk-go/models/windowsupdates"
)

type WindowsUpdatesProductsMicrosoftGraphWindowsUpdatesFindByCatalogIdWithCatalogIDFindByCatalogIdWithCatalogIDGetResponse struct {
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.BaseCollectionPaginationCountResponse
}
// NewWindowsUpdatesProductsMicrosoftGraphWindowsUpdatesFindByCatalogIdWithCatalogIDFindByCatalogIdWithCatalogIDGetResponse instantiates a new WindowsUpdatesProductsMicrosoftGraphWindowsUpdatesFindByCatalogIdWithCatalogIDFindByCatalogIdWithCatalogIDGetResponse and sets the default values.
func NewWindowsUpdatesProductsMicrosoftGraphWindowsUpdatesFindByCatalogIdWithCatalogIDFindByCatalogIdWithCatalogIDGetResponse()(*WindowsUpdatesProductsMicrosoftGraphWindowsUpdatesFindByCatalogIdWithCatalogIDFindByCatalogIdWithCatalogIDGetResponse) {
    m := &WindowsUpdatesProductsMicrosoftGraphWindowsUpdatesFindByCatalogIdWithCatalogIDFindByCatalogIdWithCatalogIDGetResponse{
        BaseCollectionPaginationCountResponse: *ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.NewBaseCollectionPaginationCountResponse(),
    }
    return m
}
// CreateWindowsUpdatesProductsMicrosoftGraphWindowsUpdatesFindByCatalogIdWithCatalogIDFindByCatalogIdWithCatalogIDGetResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateWindowsUpdatesProductsMicrosoftGraphWindowsUpdatesFindByCatalogIdWithCatalogIDFindByCatalogIdWithCatalogIDGetResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewWindowsUpdatesProductsMicrosoftGraphWindowsUpdatesFindByCatalogIdWithCatalogIDFindByCatalogIdWithCatalogIDGetResponse(), nil
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *WindowsUpdatesProductsMicrosoftGraphWindowsUpdatesFindByCatalogIdWithCatalogIDFindByCatalogIdWithCatalogIDGetResponse) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.BaseCollectionPaginationCountResponse.GetFieldDeserializers()
    res["value"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(i17376df570f19ff3c32da2d66a677d31250ed0ff64059351645f48a152316b3c.CreateProductFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]i17376df570f19ff3c32da2d66a677d31250ed0ff64059351645f48a152316b3c.Productable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(i17376df570f19ff3c32da2d66a677d31250ed0ff64059351645f48a152316b3c.Productable)
                }
            }
            m.SetValue(res)
        }
        return nil
    }
    return res
}
// GetValue gets the value property value. The value property
// returns a []Productable when successful
func (m *WindowsUpdatesProductsMicrosoftGraphWindowsUpdatesFindByCatalogIdWithCatalogIDFindByCatalogIdWithCatalogIDGetResponse) GetValue()([]i17376df570f19ff3c32da2d66a677d31250ed0ff64059351645f48a152316b3c.Productable) {
    val, err := m.GetBackingStore().Get("value")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]i17376df570f19ff3c32da2d66a677d31250ed0ff64059351645f48a152316b3c.Productable)
    }
    return nil
}
// Serialize serializes information the current object
func (m *WindowsUpdatesProductsMicrosoftGraphWindowsUpdatesFindByCatalogIdWithCatalogIDFindByCatalogIdWithCatalogIDGetResponse) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.BaseCollectionPaginationCountResponse.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetValue() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetValue()))
        for i, v := range m.GetValue() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err = writer.WriteCollectionOfObjectValues("value", cast)
        if err != nil {
            return err
        }
    }
    return nil
}
// SetValue sets the value property value. The value property
func (m *WindowsUpdatesProductsMicrosoftGraphWindowsUpdatesFindByCatalogIdWithCatalogIDFindByCatalogIdWithCatalogIDGetResponse) SetValue(value []i17376df570f19ff3c32da2d66a677d31250ed0ff64059351645f48a152316b3c.Productable)() {
    err := m.GetBackingStore().Set("value", value)
    if err != nil {
        panic(err)
    }
}
type WindowsUpdatesProductsMicrosoftGraphWindowsUpdatesFindByCatalogIdWithCatalogIDFindByCatalogIdWithCatalogIDGetResponseable interface {
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.BaseCollectionPaginationCountResponseable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetValue()([]i17376df570f19ff3c32da2d66a677d31250ed0ff64059351645f48a152316b3c.Productable)
    SetValue(value []i17376df570f19ff3c32da2d66a677d31250ed0ff64059351645f48a152316b3c.Productable)()
}
