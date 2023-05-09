package external

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    ie98116770ca9f5eee835504331ccb9976e822c2f776cca356ee95c843b4cce86 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/externalconnectors"
)

// ConnectionsItemItemsItemMicrosoftGraphExternalConnectorsAddActivitiesAddActivitiesResponse 
type ConnectionsItemItemsItemMicrosoftGraphExternalConnectorsAddActivitiesAddActivitiesResponse struct {
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.BaseCollectionPaginationCountResponse
}
// NewConnectionsItemItemsItemMicrosoftGraphExternalConnectorsAddActivitiesAddActivitiesResponse instantiates a new ConnectionsItemItemsItemMicrosoftGraphExternalConnectorsAddActivitiesAddActivitiesResponse and sets the default values.
func NewConnectionsItemItemsItemMicrosoftGraphExternalConnectorsAddActivitiesAddActivitiesResponse()(*ConnectionsItemItemsItemMicrosoftGraphExternalConnectorsAddActivitiesAddActivitiesResponse) {
    m := &ConnectionsItemItemsItemMicrosoftGraphExternalConnectorsAddActivitiesAddActivitiesResponse{
        BaseCollectionPaginationCountResponse: *ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.NewBaseCollectionPaginationCountResponse(),
    }
    return m
}
// CreateConnectionsItemItemsItemMicrosoftGraphExternalConnectorsAddActivitiesAddActivitiesResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateConnectionsItemItemsItemMicrosoftGraphExternalConnectorsAddActivitiesAddActivitiesResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewConnectionsItemItemsItemMicrosoftGraphExternalConnectorsAddActivitiesAddActivitiesResponse(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ConnectionsItemItemsItemMicrosoftGraphExternalConnectorsAddActivitiesAddActivitiesResponse) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.BaseCollectionPaginationCountResponse.GetFieldDeserializers()
    res["value"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(ie98116770ca9f5eee835504331ccb9976e822c2f776cca356ee95c843b4cce86.CreateExternalActivityResultFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]ie98116770ca9f5eee835504331ccb9976e822c2f776cca356ee95c843b4cce86.ExternalActivityResultable, len(val))
            for i, v := range val {
                res[i] = v.(ie98116770ca9f5eee835504331ccb9976e822c2f776cca356ee95c843b4cce86.ExternalActivityResultable)
            }
            m.SetValue(res)
        }
        return nil
    }
    return res
}
// GetValue gets the value property value. The value property
func (m *ConnectionsItemItemsItemMicrosoftGraphExternalConnectorsAddActivitiesAddActivitiesResponse) GetValue()([]ie98116770ca9f5eee835504331ccb9976e822c2f776cca356ee95c843b4cce86.ExternalActivityResultable) {
    val, err := m.GetBackingStore().Get("value")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]ie98116770ca9f5eee835504331ccb9976e822c2f776cca356ee95c843b4cce86.ExternalActivityResultable)
    }
    return nil
}
// Serialize serializes information the current object
func (m *ConnectionsItemItemsItemMicrosoftGraphExternalConnectorsAddActivitiesAddActivitiesResponse) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
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
func (m *ConnectionsItemItemsItemMicrosoftGraphExternalConnectorsAddActivitiesAddActivitiesResponse) SetValue(value []ie98116770ca9f5eee835504331ccb9976e822c2f776cca356ee95c843b4cce86.ExternalActivityResultable)() {
    err := m.GetBackingStore().Set("value", value)
    if err != nil {
        panic(err)
    }
}
// ConnectionsItemItemsItemMicrosoftGraphExternalConnectorsAddActivitiesAddActivitiesResponseable 
type ConnectionsItemItemsItemMicrosoftGraphExternalConnectorsAddActivitiesAddActivitiesResponseable interface {
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.BaseCollectionPaginationCountResponseable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetValue()([]ie98116770ca9f5eee835504331ccb9976e822c2f776cca356ee95c843b4cce86.ExternalActivityResultable)
    SetValue(value []ie98116770ca9f5eee835504331ccb9976e822c2f776cca356ee95c843b4cce86.ExternalActivityResultable)()
}
