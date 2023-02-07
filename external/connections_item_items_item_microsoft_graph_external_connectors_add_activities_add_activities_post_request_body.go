package external

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    ie98116770ca9f5eee835504331ccb9976e822c2f776cca356ee95c843b4cce86 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/externalconnectors"
)

// ConnectionsItemItemsItemMicrosoftGraphExternalConnectorsAddActivitiesAddActivitiesPostRequestBody 
type ConnectionsItemItemsItemMicrosoftGraphExternalConnectorsAddActivitiesAddActivitiesPostRequestBody struct {
    // The activities property
    activities []ie98116770ca9f5eee835504331ccb9976e822c2f776cca356ee95c843b4cce86.ExternalActivityable
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
}
// NewConnectionsItemItemsItemMicrosoftGraphExternalConnectorsAddActivitiesAddActivitiesPostRequestBody instantiates a new ConnectionsItemItemsItemMicrosoftGraphExternalConnectorsAddActivitiesAddActivitiesPostRequestBody and sets the default values.
func NewConnectionsItemItemsItemMicrosoftGraphExternalConnectorsAddActivitiesAddActivitiesPostRequestBody()(*ConnectionsItemItemsItemMicrosoftGraphExternalConnectorsAddActivitiesAddActivitiesPostRequestBody) {
    m := &ConnectionsItemItemsItemMicrosoftGraphExternalConnectorsAddActivitiesAddActivitiesPostRequestBody{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateConnectionsItemItemsItemMicrosoftGraphExternalConnectorsAddActivitiesAddActivitiesPostRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateConnectionsItemItemsItemMicrosoftGraphExternalConnectorsAddActivitiesAddActivitiesPostRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewConnectionsItemItemsItemMicrosoftGraphExternalConnectorsAddActivitiesAddActivitiesPostRequestBody(), nil
}
// GetActivities gets the activities property value. The activities property
func (m *ConnectionsItemItemsItemMicrosoftGraphExternalConnectorsAddActivitiesAddActivitiesPostRequestBody) GetActivities()([]ie98116770ca9f5eee835504331ccb9976e822c2f776cca356ee95c843b4cce86.ExternalActivityable) {
    return m.activities
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *ConnectionsItemItemsItemMicrosoftGraphExternalConnectorsAddActivitiesAddActivitiesPostRequestBody) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ConnectionsItemItemsItemMicrosoftGraphExternalConnectorsAddActivitiesAddActivitiesPostRequestBody) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["activities"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(ie98116770ca9f5eee835504331ccb9976e822c2f776cca356ee95c843b4cce86.CreateExternalActivityFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]ie98116770ca9f5eee835504331ccb9976e822c2f776cca356ee95c843b4cce86.ExternalActivityable, len(val))
            for i, v := range val {
                res[i] = v.(ie98116770ca9f5eee835504331ccb9976e822c2f776cca356ee95c843b4cce86.ExternalActivityable)
            }
            m.SetActivities(res)
        }
        return nil
    }
    return res
}
// Serialize serializes information the current object
func (m *ConnectionsItemItemsItemMicrosoftGraphExternalConnectorsAddActivitiesAddActivitiesPostRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    if m.GetActivities() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetActivities()))
        for i, v := range m.GetActivities() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err := writer.WriteCollectionOfObjectValues("activities", cast)
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
// SetActivities sets the activities property value. The activities property
func (m *ConnectionsItemItemsItemMicrosoftGraphExternalConnectorsAddActivitiesAddActivitiesPostRequestBody) SetActivities(value []ie98116770ca9f5eee835504331ccb9976e822c2f776cca356ee95c843b4cce86.ExternalActivityable)() {
    m.activities = value
}
// SetAdditionalData sets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *ConnectionsItemItemsItemMicrosoftGraphExternalConnectorsAddActivitiesAddActivitiesPostRequestBody) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
