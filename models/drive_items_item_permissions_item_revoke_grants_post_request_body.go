package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// DriveItemsItemPermissionsItemRevokeGrantsPostRequestBody provides operations to call the revokeGrants method.
type DriveItemsItemPermissionsItemRevokeGrantsPostRequestBody struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // The grantees property
    grantees []DriveRecipientable
}
// NewDriveItemsItemPermissionsItemRevokeGrantsPostRequestBody instantiates a new DriveItemsItemPermissionsItemRevokeGrantsPostRequestBody and sets the default values.
func NewDriveItemsItemPermissionsItemRevokeGrantsPostRequestBody()(*DriveItemsItemPermissionsItemRevokeGrantsPostRequestBody) {
    m := &DriveItemsItemPermissionsItemRevokeGrantsPostRequestBody{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateDriveItemsItemPermissionsItemRevokeGrantsPostRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateDriveItemsItemPermissionsItemRevokeGrantsPostRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewDriveItemsItemPermissionsItemRevokeGrantsPostRequestBody(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *DriveItemsItemPermissionsItemRevokeGrantsPostRequestBody) GetAdditionalData()(map[string]interface{}) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
func (m *DriveItemsItemPermissionsItemRevokeGrantsPostRequestBody) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["grantees"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateDriveRecipientFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]DriveRecipientable, len(val))
            for i, v := range val {
                res[i] = v.(DriveRecipientable)
            }
            m.SetGrantees(res)
        }
        return nil
    }
    return res
}
// GetGrantees gets the grantees property value. The grantees property
func (m *DriveItemsItemPermissionsItemRevokeGrantsPostRequestBody) GetGrantees()([]DriveRecipientable) {
    return m.grantees
}
// Serialize serializes information the current object
func (m *DriveItemsItemPermissionsItemRevokeGrantsPostRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
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
func (m *DriveItemsItemPermissionsItemRevokeGrantsPostRequestBody) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// SetGrantees sets the grantees property value. The grantees property
func (m *DriveItemsItemPermissionsItemRevokeGrantsPostRequestBody) SetGrantees(value []DriveRecipientable)() {
    m.grantees = value
}
