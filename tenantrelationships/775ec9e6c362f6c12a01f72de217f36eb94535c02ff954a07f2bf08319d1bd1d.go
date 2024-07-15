package tenantrelationships

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i72d786f54cc0bb289c971b085dd642b2fc3af6394328682e69783fd7e229b582 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/managedtenants"
)

type ManagedTenantsTenantGroupsMicrosoftGraphManagedTenantsTenantSearchTenantSearchPostResponse struct {
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.BaseCollectionPaginationCountResponse
}
// NewManagedTenantsTenantGroupsMicrosoftGraphManagedTenantsTenantSearchTenantSearchPostResponse instantiates a new ManagedTenantsTenantGroupsMicrosoftGraphManagedTenantsTenantSearchTenantSearchPostResponse and sets the default values.
func NewManagedTenantsTenantGroupsMicrosoftGraphManagedTenantsTenantSearchTenantSearchPostResponse()(*ManagedTenantsTenantGroupsMicrosoftGraphManagedTenantsTenantSearchTenantSearchPostResponse) {
    m := &ManagedTenantsTenantGroupsMicrosoftGraphManagedTenantsTenantSearchTenantSearchPostResponse{
        BaseCollectionPaginationCountResponse: *ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.NewBaseCollectionPaginationCountResponse(),
    }
    return m
}
// CreateManagedTenantsTenantGroupsMicrosoftGraphManagedTenantsTenantSearchTenantSearchPostResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateManagedTenantsTenantGroupsMicrosoftGraphManagedTenantsTenantSearchTenantSearchPostResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewManagedTenantsTenantGroupsMicrosoftGraphManagedTenantsTenantSearchTenantSearchPostResponse(), nil
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *ManagedTenantsTenantGroupsMicrosoftGraphManagedTenantsTenantSearchTenantSearchPostResponse) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.BaseCollectionPaginationCountResponse.GetFieldDeserializers()
    res["value"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(i72d786f54cc0bb289c971b085dd642b2fc3af6394328682e69783fd7e229b582.CreateTenantGroupFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]i72d786f54cc0bb289c971b085dd642b2fc3af6394328682e69783fd7e229b582.TenantGroupable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(i72d786f54cc0bb289c971b085dd642b2fc3af6394328682e69783fd7e229b582.TenantGroupable)
                }
            }
            m.SetValue(res)
        }
        return nil
    }
    return res
}
// GetValue gets the value property value. The value property
// returns a []TenantGroupable when successful
func (m *ManagedTenantsTenantGroupsMicrosoftGraphManagedTenantsTenantSearchTenantSearchPostResponse) GetValue()([]i72d786f54cc0bb289c971b085dd642b2fc3af6394328682e69783fd7e229b582.TenantGroupable) {
    val, err := m.GetBackingStore().Get("value")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]i72d786f54cc0bb289c971b085dd642b2fc3af6394328682e69783fd7e229b582.TenantGroupable)
    }
    return nil
}
// Serialize serializes information the current object
func (m *ManagedTenantsTenantGroupsMicrosoftGraphManagedTenantsTenantSearchTenantSearchPostResponse) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
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
func (m *ManagedTenantsTenantGroupsMicrosoftGraphManagedTenantsTenantSearchTenantSearchPostResponse) SetValue(value []i72d786f54cc0bb289c971b085dd642b2fc3af6394328682e69783fd7e229b582.TenantGroupable)() {
    err := m.GetBackingStore().Set("value", value)
    if err != nil {
        panic(err)
    }
}
type ManagedTenantsTenantGroupsMicrosoftGraphManagedTenantsTenantSearchTenantSearchPostResponseable interface {
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.BaseCollectionPaginationCountResponseable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetValue()([]i72d786f54cc0bb289c971b085dd642b2fc3af6394328682e69783fd7e229b582.TenantGroupable)
    SetValue(value []i72d786f54cc0bb289c971b085dd642b2fc3af6394328682e69783fd7e229b582.TenantGroupable)()
}
