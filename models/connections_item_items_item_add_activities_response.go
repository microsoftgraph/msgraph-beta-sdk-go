package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ConnectionsItemItemsItemAddActivitiesResponse provides operations to call the addActivities method.
type ConnectionsItemItemsItemAddActivitiesResponse struct {
    BaseCollectionPaginationCountResponse
}
// NewConnectionsItemItemsItemAddActivitiesResponse instantiates a new ConnectionsItemItemsItemAddActivitiesResponse and sets the default values.
func NewConnectionsItemItemsItemAddActivitiesResponse()(*ConnectionsItemItemsItemAddActivitiesResponse) {
    m := &ConnectionsItemItemsItemAddActivitiesResponse{
        BaseCollectionPaginationCountResponse: *NewBaseCollectionPaginationCountResponse(),
    }
    return m
}
// CreateConnectionsItemItemsItemAddActivitiesResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateConnectionsItemItemsItemAddActivitiesResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewConnectionsItemItemsItemAddActivitiesResponse(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ConnectionsItemItemsItemAddActivitiesResponse) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.BaseCollectionPaginationCountResponse.GetFieldDeserializers()
    return res
}
// Serialize serializes information the current object
func (m *ConnectionsItemItemsItemAddActivitiesResponse) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.BaseCollectionPaginationCountResponse.Serialize(writer)
    if err != nil {
        return err
    }
    return nil
}
