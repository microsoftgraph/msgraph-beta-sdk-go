package users

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
)

// ItemPresenceMicrosoftGraphSetStatusMessageSetStatusMessagePostRequestBody 
type ItemPresenceMicrosoftGraphSetStatusMessageSetStatusMessagePostRequestBody struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The statusMessage property
    statusMessage ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.PresenceStatusMessageable
}
// NewItemPresenceMicrosoftGraphSetStatusMessageSetStatusMessagePostRequestBody instantiates a new ItemPresenceMicrosoftGraphSetStatusMessageSetStatusMessagePostRequestBody and sets the default values.
func NewItemPresenceMicrosoftGraphSetStatusMessageSetStatusMessagePostRequestBody()(*ItemPresenceMicrosoftGraphSetStatusMessageSetStatusMessagePostRequestBody) {
    m := &ItemPresenceMicrosoftGraphSetStatusMessageSetStatusMessagePostRequestBody{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateItemPresenceMicrosoftGraphSetStatusMessageSetStatusMessagePostRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateItemPresenceMicrosoftGraphSetStatusMessageSetStatusMessagePostRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewItemPresenceMicrosoftGraphSetStatusMessageSetStatusMessagePostRequestBody(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *ItemPresenceMicrosoftGraphSetStatusMessageSetStatusMessagePostRequestBody) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ItemPresenceMicrosoftGraphSetStatusMessageSetStatusMessagePostRequestBody) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["statusMessage"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreatePresenceStatusMessageFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetStatusMessage(val.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.PresenceStatusMessageable))
        }
        return nil
    }
    return res
}
// GetStatusMessage gets the statusMessage property value. The statusMessage property
func (m *ItemPresenceMicrosoftGraphSetStatusMessageSetStatusMessagePostRequestBody) GetStatusMessage()(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.PresenceStatusMessageable) {
    return m.statusMessage
}
// Serialize serializes information the current object
func (m *ItemPresenceMicrosoftGraphSetStatusMessageSetStatusMessagePostRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteObjectValue("statusMessage", m.GetStatusMessage())
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
func (m *ItemPresenceMicrosoftGraphSetStatusMessageSetStatusMessagePostRequestBody) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetStatusMessage sets the statusMessage property value. The statusMessage property
func (m *ItemPresenceMicrosoftGraphSetStatusMessageSetStatusMessagePostRequestBody) SetStatusMessage(value ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.PresenceStatusMessageable)() {
    m.statusMessage = value
}
