package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// DirectoryDeletedItemsGetUserOwnedObjectsPostRequestBody provides operations to call the getUserOwnedObjects method.
type DirectoryDeletedItemsGetUserOwnedObjectsPostRequestBody struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // The type property
    type_escaped *string
    // The userId property
    userId *string
}
// NewDirectoryDeletedItemsGetUserOwnedObjectsPostRequestBody instantiates a new DirectoryDeletedItemsGetUserOwnedObjectsPostRequestBody and sets the default values.
func NewDirectoryDeletedItemsGetUserOwnedObjectsPostRequestBody()(*DirectoryDeletedItemsGetUserOwnedObjectsPostRequestBody) {
    m := &DirectoryDeletedItemsGetUserOwnedObjectsPostRequestBody{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateDirectoryDeletedItemsGetUserOwnedObjectsPostRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateDirectoryDeletedItemsGetUserOwnedObjectsPostRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewDirectoryDeletedItemsGetUserOwnedObjectsPostRequestBody(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *DirectoryDeletedItemsGetUserOwnedObjectsPostRequestBody) GetAdditionalData()(map[string]interface{}) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
func (m *DirectoryDeletedItemsGetUserOwnedObjectsPostRequestBody) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["type"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetType(val)
        }
        return nil
    }
    res["userId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUserId(val)
        }
        return nil
    }
    return res
}
// GetType gets the type property value. The type property
func (m *DirectoryDeletedItemsGetUserOwnedObjectsPostRequestBody) GetType()(*string) {
    return m.type_escaped
}
// GetUserId gets the userId property value. The userId property
func (m *DirectoryDeletedItemsGetUserOwnedObjectsPostRequestBody) GetUserId()(*string) {
    return m.userId
}
// Serialize serializes information the current object
func (m *DirectoryDeletedItemsGetUserOwnedObjectsPostRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("type", m.GetType())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("userId", m.GetUserId())
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
func (m *DirectoryDeletedItemsGetUserOwnedObjectsPostRequestBody) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// SetType sets the type property value. The type property
func (m *DirectoryDeletedItemsGetUserOwnedObjectsPostRequestBody) SetType(value *string)() {
    m.type_escaped = value
}
// SetUserId sets the userId property value. The userId property
func (m *DirectoryDeletedItemsGetUserOwnedObjectsPostRequestBody) SetUserId(value *string)() {
    m.userId = value
}
