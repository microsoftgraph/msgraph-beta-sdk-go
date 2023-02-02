package serviceprincipals

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ItemMicrosoftGraphGetPasswordSingleSignOnCredentialsGetPasswordSingleSignOnCredentialsPostRequestBody 
type ItemMicrosoftGraphGetPasswordSingleSignOnCredentialsGetPasswordSingleSignOnCredentialsPostRequestBody struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The id property
    id *string
}
// NewItemMicrosoftGraphGetPasswordSingleSignOnCredentialsGetPasswordSingleSignOnCredentialsPostRequestBody instantiates a new ItemMicrosoftGraphGetPasswordSingleSignOnCredentialsGetPasswordSingleSignOnCredentialsPostRequestBody and sets the default values.
func NewItemMicrosoftGraphGetPasswordSingleSignOnCredentialsGetPasswordSingleSignOnCredentialsPostRequestBody()(*ItemMicrosoftGraphGetPasswordSingleSignOnCredentialsGetPasswordSingleSignOnCredentialsPostRequestBody) {
    m := &ItemMicrosoftGraphGetPasswordSingleSignOnCredentialsGetPasswordSingleSignOnCredentialsPostRequestBody{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateItemMicrosoftGraphGetPasswordSingleSignOnCredentialsGetPasswordSingleSignOnCredentialsPostRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateItemMicrosoftGraphGetPasswordSingleSignOnCredentialsGetPasswordSingleSignOnCredentialsPostRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewItemMicrosoftGraphGetPasswordSingleSignOnCredentialsGetPasswordSingleSignOnCredentialsPostRequestBody(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *ItemMicrosoftGraphGetPasswordSingleSignOnCredentialsGetPasswordSingleSignOnCredentialsPostRequestBody) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ItemMicrosoftGraphGetPasswordSingleSignOnCredentialsGetPasswordSingleSignOnCredentialsPostRequestBody) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["id"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetId(val)
        }
        return nil
    }
    return res
}
// GetId gets the id property value. The id property
func (m *ItemMicrosoftGraphGetPasswordSingleSignOnCredentialsGetPasswordSingleSignOnCredentialsPostRequestBody) GetId()(*string) {
    return m.id
}
// Serialize serializes information the current object
func (m *ItemMicrosoftGraphGetPasswordSingleSignOnCredentialsGetPasswordSingleSignOnCredentialsPostRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("id", m.GetId())
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
func (m *ItemMicrosoftGraphGetPasswordSingleSignOnCredentialsGetPasswordSingleSignOnCredentialsPostRequestBody) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetId sets the id property value. The id property
func (m *ItemMicrosoftGraphGetPasswordSingleSignOnCredentialsGetPasswordSingleSignOnCredentialsPostRequestBody) SetId(value *string)() {
    m.id = value
}
