package users

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ItemMicrosoftGraphInvalidateAllRefreshTokensInvalidateAllRefreshTokensResponse 
type ItemMicrosoftGraphInvalidateAllRefreshTokensInvalidateAllRefreshTokensResponse struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The value property
    value *bool
}
// NewItemMicrosoftGraphInvalidateAllRefreshTokensInvalidateAllRefreshTokensResponse instantiates a new ItemMicrosoftGraphInvalidateAllRefreshTokensInvalidateAllRefreshTokensResponse and sets the default values.
func NewItemMicrosoftGraphInvalidateAllRefreshTokensInvalidateAllRefreshTokensResponse()(*ItemMicrosoftGraphInvalidateAllRefreshTokensInvalidateAllRefreshTokensResponse) {
    m := &ItemMicrosoftGraphInvalidateAllRefreshTokensInvalidateAllRefreshTokensResponse{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateItemMicrosoftGraphInvalidateAllRefreshTokensInvalidateAllRefreshTokensResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateItemMicrosoftGraphInvalidateAllRefreshTokensInvalidateAllRefreshTokensResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewItemMicrosoftGraphInvalidateAllRefreshTokensInvalidateAllRefreshTokensResponse(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *ItemMicrosoftGraphInvalidateAllRefreshTokensInvalidateAllRefreshTokensResponse) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ItemMicrosoftGraphInvalidateAllRefreshTokensInvalidateAllRefreshTokensResponse) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["value"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetValue(val)
        }
        return nil
    }
    return res
}
// GetValue gets the value property value. The value property
func (m *ItemMicrosoftGraphInvalidateAllRefreshTokensInvalidateAllRefreshTokensResponse) GetValue()(*bool) {
    return m.value
}
// Serialize serializes information the current object
func (m *ItemMicrosoftGraphInvalidateAllRefreshTokensInvalidateAllRefreshTokensResponse) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteBoolValue("value", m.GetValue())
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
func (m *ItemMicrosoftGraphInvalidateAllRefreshTokensInvalidateAllRefreshTokensResponse) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetValue sets the value property value. The value property
func (m *ItemMicrosoftGraphInvalidateAllRefreshTokensInvalidateAllRefreshTokensResponse) SetValue(value *bool)() {
    m.value = value
}
