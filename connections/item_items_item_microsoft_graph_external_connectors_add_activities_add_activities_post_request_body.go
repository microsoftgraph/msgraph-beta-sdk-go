package connections

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e "github.com/microsoft/kiota-abstractions-go/store"
    ie98116770ca9f5eee835504331ccb9976e822c2f776cca356ee95c843b4cce86 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/externalconnectors"
)

// ItemItemsItemMicrosoftGraphExternalConnectorsAddActivitiesAddActivitiesPostRequestBody 
type ItemItemsItemMicrosoftGraphExternalConnectorsAddActivitiesAddActivitiesPostRequestBody struct {
    // Stores model information.
    backingStore ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore
}
// NewItemItemsItemMicrosoftGraphExternalConnectorsAddActivitiesAddActivitiesPostRequestBody instantiates a new ItemItemsItemMicrosoftGraphExternalConnectorsAddActivitiesAddActivitiesPostRequestBody and sets the default values.
func NewItemItemsItemMicrosoftGraphExternalConnectorsAddActivitiesAddActivitiesPostRequestBody()(*ItemItemsItemMicrosoftGraphExternalConnectorsAddActivitiesAddActivitiesPostRequestBody) {
    m := &ItemItemsItemMicrosoftGraphExternalConnectorsAddActivitiesAddActivitiesPostRequestBody{
    }
    m.backingStore = ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStoreFactoryInstance();
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateItemItemsItemMicrosoftGraphExternalConnectorsAddActivitiesAddActivitiesPostRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateItemItemsItemMicrosoftGraphExternalConnectorsAddActivitiesAddActivitiesPostRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewItemItemsItemMicrosoftGraphExternalConnectorsAddActivitiesAddActivitiesPostRequestBody(), nil
}
// GetActivities gets the activities property value. The activities property
func (m *ItemItemsItemMicrosoftGraphExternalConnectorsAddActivitiesAddActivitiesPostRequestBody) GetActivities()([]ie98116770ca9f5eee835504331ccb9976e822c2f776cca356ee95c843b4cce86.ExternalActivityable) {
    val, err := m.GetBackingStore().Get("activities")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]ie98116770ca9f5eee835504331ccb9976e822c2f776cca356ee95c843b4cce86.ExternalActivityable)
    }
    return nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *ItemItemsItemMicrosoftGraphExternalConnectorsAddActivitiesAddActivitiesPostRequestBody) GetAdditionalData()(map[string]any) {
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
func (m *ItemItemsItemMicrosoftGraphExternalConnectorsAddActivitiesAddActivitiesPostRequestBody) GetBackingStore()(ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore) {
    return m.backingStore
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ItemItemsItemMicrosoftGraphExternalConnectorsAddActivitiesAddActivitiesPostRequestBody) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["activities"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(ie98116770ca9f5eee835504331ccb9976e822c2f776cca356ee95c843b4cce86.CreateExternalActivityFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]ie98116770ca9f5eee835504331ccb9976e822c2f776cca356ee95c843b4cce86.ExternalActivityable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(ie98116770ca9f5eee835504331ccb9976e822c2f776cca356ee95c843b4cce86.ExternalActivityable)
                }
            }
            m.SetActivities(res)
        }
        return nil
    }
    return res
}
// Serialize serializes information the current object
func (m *ItemItemsItemMicrosoftGraphExternalConnectorsAddActivitiesAddActivitiesPostRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    if m.GetActivities() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetActivities()))
        for i, v := range m.GetActivities() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
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
func (m *ItemItemsItemMicrosoftGraphExternalConnectorsAddActivitiesAddActivitiesPostRequestBody) SetActivities(value []ie98116770ca9f5eee835504331ccb9976e822c2f776cca356ee95c843b4cce86.ExternalActivityable)() {
    err := m.GetBackingStore().Set("activities", value)
    if err != nil {
        panic(err)
    }
}
// SetAdditionalData sets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *ItemItemsItemMicrosoftGraphExternalConnectorsAddActivitiesAddActivitiesPostRequestBody) SetAdditionalData(value map[string]any)() {
    err := m.GetBackingStore().Set("additionalData", value)
    if err != nil {
        panic(err)
    }
}
// SetBackingStore sets the BackingStore property value. Stores model information.
func (m *ItemItemsItemMicrosoftGraphExternalConnectorsAddActivitiesAddActivitiesPostRequestBody) SetBackingStore(value ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)() {
    m.backingStore = value
}
// ItemItemsItemMicrosoftGraphExternalConnectorsAddActivitiesAddActivitiesPostRequestBodyable 
type ItemItemsItemMicrosoftGraphExternalConnectorsAddActivitiesAddActivitiesPostRequestBodyable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackedModel
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetActivities()([]ie98116770ca9f5eee835504331ccb9976e822c2f776cca356ee95c843b4cce86.ExternalActivityable)
    GetBackingStore()(ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)
    SetActivities(value []ie98116770ca9f5eee835504331ccb9976e822c2f776cca356ee95c843b4cce86.ExternalActivityable)()
    SetBackingStore(value ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)()
}
