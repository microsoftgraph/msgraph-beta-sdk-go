package drives

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
)

// ItemItemsItemPermissionsItemMicrosoftGraphRevokeGrantsRevokeGrantsPostRequestBody 
type ItemItemsItemPermissionsItemMicrosoftGraphRevokeGrantsRevokeGrantsPostRequestBody struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The grantees property
    grantees []ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DriveRecipientable
}
// NewItemItemsItemPermissionsItemMicrosoftGraphRevokeGrantsRevokeGrantsPostRequestBody instantiates a new ItemItemsItemPermissionsItemMicrosoftGraphRevokeGrantsRevokeGrantsPostRequestBody and sets the default values.
func NewItemItemsItemPermissionsItemMicrosoftGraphRevokeGrantsRevokeGrantsPostRequestBody()(*ItemItemsItemPermissionsItemMicrosoftGraphRevokeGrantsRevokeGrantsPostRequestBody) {
    m := &ItemItemsItemPermissionsItemMicrosoftGraphRevokeGrantsRevokeGrantsPostRequestBody{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateItemItemsItemPermissionsItemMicrosoftGraphRevokeGrantsRevokeGrantsPostRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateItemItemsItemPermissionsItemMicrosoftGraphRevokeGrantsRevokeGrantsPostRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewItemItemsItemPermissionsItemMicrosoftGraphRevokeGrantsRevokeGrantsPostRequestBody(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *ItemItemsItemPermissionsItemMicrosoftGraphRevokeGrantsRevokeGrantsPostRequestBody) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ItemItemsItemPermissionsItemMicrosoftGraphRevokeGrantsRevokeGrantsPostRequestBody) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["grantees"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateDriveRecipientFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DriveRecipientable, len(val))
            for i, v := range val {
                res[i] = v.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DriveRecipientable)
            }
            m.SetGrantees(res)
        }
        return nil
    }
    return res
}
// GetGrantees gets the grantees property value. The grantees property
func (m *ItemItemsItemPermissionsItemMicrosoftGraphRevokeGrantsRevokeGrantsPostRequestBody) GetGrantees()([]ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DriveRecipientable) {
    return m.grantees
}
// Serialize serializes information the current object
func (m *ItemItemsItemPermissionsItemMicrosoftGraphRevokeGrantsRevokeGrantsPostRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    if m.GetGrantees() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetGrantees()))
        for i, v := range m.GetGrantees() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err := writer.WriteCollectionOfObjectValues("grantees", cast)
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
// SetAdditionalData sets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *ItemItemsItemPermissionsItemMicrosoftGraphRevokeGrantsRevokeGrantsPostRequestBody) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetGrantees sets the grantees property value. The grantees property
func (m *ItemItemsItemPermissionsItemMicrosoftGraphRevokeGrantsRevokeGrantsPostRequestBody) SetGrantees(value []ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DriveRecipientable)() {
    m.grantees = value
}
