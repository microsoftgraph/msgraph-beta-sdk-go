package tenantrelationships

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i72d786f54cc0bb289c971b085dd642b2fc3af6394328682e69783fd7e229b582 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/managedtenants"
)

// ManagedTenantsTenantGroupsMicrosoftGraphManagedTenantsTenantSearchTenantSearchResponse 
type ManagedTenantsTenantGroupsMicrosoftGraphManagedTenantsTenantSearchTenantSearchResponse struct {
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.BaseCollectionPaginationCountResponse
    // The value property
    value []i72d786f54cc0bb289c971b085dd642b2fc3af6394328682e69783fd7e229b582.TenantGroupable
}
// NewManagedTenantsTenantGroupsMicrosoftGraphManagedTenantsTenantSearchTenantSearchResponse instantiates a new ManagedTenantsTenantGroupsMicrosoftGraphManagedTenantsTenantSearchTenantSearchResponse and sets the default values.
func NewManagedTenantsTenantGroupsMicrosoftGraphManagedTenantsTenantSearchTenantSearchResponse()(*ManagedTenantsTenantGroupsMicrosoftGraphManagedTenantsTenantSearchTenantSearchResponse) {
    m := &ManagedTenantsTenantGroupsMicrosoftGraphManagedTenantsTenantSearchTenantSearchResponse{
        BaseCollectionPaginationCountResponse: *ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.NewBaseCollectionPaginationCountResponse(),
    }
    return m
}
// CreateManagedTenantsTenantGroupsMicrosoftGraphManagedTenantsTenantSearchTenantSearchResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateManagedTenantsTenantGroupsMicrosoftGraphManagedTenantsTenantSearchTenantSearchResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewManagedTenantsTenantGroupsMicrosoftGraphManagedTenantsTenantSearchTenantSearchResponse(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ManagedTenantsTenantGroupsMicrosoftGraphManagedTenantsTenantSearchTenantSearchResponse) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.BaseCollectionPaginationCountResponse.GetFieldDeserializers()
    res["value"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(i72d786f54cc0bb289c971b085dd642b2fc3af6394328682e69783fd7e229b582.CreateTenantGroupFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]i72d786f54cc0bb289c971b085dd642b2fc3af6394328682e69783fd7e229b582.TenantGroupable, len(val))
            for i, v := range val {
                res[i] = v.(i72d786f54cc0bb289c971b085dd642b2fc3af6394328682e69783fd7e229b582.TenantGroupable)
            }
            m.SetValue(res)
        }
        return nil
    }
    return res
}
// GetValue gets the value property value. The value property
func (m *ManagedTenantsTenantGroupsMicrosoftGraphManagedTenantsTenantSearchTenantSearchResponse) GetValue()([]i72d786f54cc0bb289c971b085dd642b2fc3af6394328682e69783fd7e229b582.TenantGroupable) {
    return m.value
}
// Serialize serializes information the current object
func (m *ManagedTenantsTenantGroupsMicrosoftGraphManagedTenantsTenantSearchTenantSearchResponse) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
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
func (m *ManagedTenantsTenantGroupsMicrosoftGraphManagedTenantsTenantSearchTenantSearchResponse) SetValue(value []i72d786f54cc0bb289c971b085dd642b2fc3af6394328682e69783fd7e229b582.TenantGroupable)() {
    m.value = value
}
