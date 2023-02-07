package teamtemplatedefinition

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ItemTeamDefinitionPermissionGrantsMicrosoftGraphGetUserOwnedObjectsGetUserOwnedObjectsPostRequestBody 
type ItemTeamDefinitionPermissionGrantsMicrosoftGraphGetUserOwnedObjectsGetUserOwnedObjectsPostRequestBody struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The type property
    typeEscaped *string
    // The userId property
    userId *string
}
// NewItemTeamDefinitionPermissionGrantsMicrosoftGraphGetUserOwnedObjectsGetUserOwnedObjectsPostRequestBody instantiates a new ItemTeamDefinitionPermissionGrantsMicrosoftGraphGetUserOwnedObjectsGetUserOwnedObjectsPostRequestBody and sets the default values.
func NewItemTeamDefinitionPermissionGrantsMicrosoftGraphGetUserOwnedObjectsGetUserOwnedObjectsPostRequestBody()(*ItemTeamDefinitionPermissionGrantsMicrosoftGraphGetUserOwnedObjectsGetUserOwnedObjectsPostRequestBody) {
    m := &ItemTeamDefinitionPermissionGrantsMicrosoftGraphGetUserOwnedObjectsGetUserOwnedObjectsPostRequestBody{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateItemTeamDefinitionPermissionGrantsMicrosoftGraphGetUserOwnedObjectsGetUserOwnedObjectsPostRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateItemTeamDefinitionPermissionGrantsMicrosoftGraphGetUserOwnedObjectsGetUserOwnedObjectsPostRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewItemTeamDefinitionPermissionGrantsMicrosoftGraphGetUserOwnedObjectsGetUserOwnedObjectsPostRequestBody(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *ItemTeamDefinitionPermissionGrantsMicrosoftGraphGetUserOwnedObjectsGetUserOwnedObjectsPostRequestBody) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ItemTeamDefinitionPermissionGrantsMicrosoftGraphGetUserOwnedObjectsGetUserOwnedObjectsPostRequestBody) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
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
func (m *ItemTeamDefinitionPermissionGrantsMicrosoftGraphGetUserOwnedObjectsGetUserOwnedObjectsPostRequestBody) GetType()(*string) {
    return m.typeEscaped
}
// GetUserId gets the userId property value. The userId property
func (m *ItemTeamDefinitionPermissionGrantsMicrosoftGraphGetUserOwnedObjectsGetUserOwnedObjectsPostRequestBody) GetUserId()(*string) {
    return m.userId
}
// Serialize serializes information the current object
func (m *ItemTeamDefinitionPermissionGrantsMicrosoftGraphGetUserOwnedObjectsGetUserOwnedObjectsPostRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
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
func (m *ItemTeamDefinitionPermissionGrantsMicrosoftGraphGetUserOwnedObjectsGetUserOwnedObjectsPostRequestBody) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetType sets the type property value. The type property
func (m *ItemTeamDefinitionPermissionGrantsMicrosoftGraphGetUserOwnedObjectsGetUserOwnedObjectsPostRequestBody) SetType(value *string)() {
    m.typeEscaped = value
}
// SetUserId sets the userId property value. The userId property
func (m *ItemTeamDefinitionPermissionGrantsMicrosoftGraphGetUserOwnedObjectsGetUserOwnedObjectsPostRequestBody) SetUserId(value *string)() {
    m.userId = value
}
