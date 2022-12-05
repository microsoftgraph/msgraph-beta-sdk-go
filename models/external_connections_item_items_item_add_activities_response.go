package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ExternalConnectionsItemItemsItemAddActivitiesResponse provides operations to call the addActivities method.
type ExternalConnectionsItemItemsItemAddActivitiesResponse struct {
    BaseCollectionPaginationCountResponse
}
// NewExternalConnectionsItemItemsItemAddActivitiesResponse instantiates a new ExternalConnectionsItemItemsItemAddActivitiesResponse and sets the default values.
func NewExternalConnectionsItemItemsItemAddActivitiesResponse()(*ExternalConnectionsItemItemsItemAddActivitiesResponse) {
    m := &ExternalConnectionsItemItemsItemAddActivitiesResponse{
        BaseCollectionPaginationCountResponse: *NewBaseCollectionPaginationCountResponse(),
    }
    return m
}
// CreateExternalConnectionsItemItemsItemAddActivitiesResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateExternalConnectionsItemItemsItemAddActivitiesResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewExternalConnectionsItemItemsItemAddActivitiesResponse(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ExternalConnectionsItemItemsItemAddActivitiesResponse) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.BaseCollectionPaginationCountResponse.GetFieldDeserializers()
    return res
}
// Serialize serializes information the current object
func (m *ExternalConnectionsItemItemsItemAddActivitiesResponse) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.BaseCollectionPaginationCountResponse.Serialize(writer)
    if err != nil {
        return err
    }
    return nil
}
