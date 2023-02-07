package governanceresources

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// MicrosoftGraphRegisterRegisterPostRequestBody 
type MicrosoftGraphRegisterRegisterPostRequestBody struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The externalId property
    externalId *string
}
// NewMicrosoftGraphRegisterRegisterPostRequestBody instantiates a new MicrosoftGraphRegisterRegisterPostRequestBody and sets the default values.
func NewMicrosoftGraphRegisterRegisterPostRequestBody()(*MicrosoftGraphRegisterRegisterPostRequestBody) {
    m := &MicrosoftGraphRegisterRegisterPostRequestBody{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateMicrosoftGraphRegisterRegisterPostRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateMicrosoftGraphRegisterRegisterPostRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewMicrosoftGraphRegisterRegisterPostRequestBody(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *MicrosoftGraphRegisterRegisterPostRequestBody) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetExternalId gets the externalId property value. The externalId property
func (m *MicrosoftGraphRegisterRegisterPostRequestBody) GetExternalId()(*string) {
    return m.externalId
}
// GetFieldDeserializers the deserialization information for the current model
func (m *MicrosoftGraphRegisterRegisterPostRequestBody) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["externalId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetExternalId(val)
        }
        return nil
    }
    return res
}
// Serialize serializes information the current object
func (m *MicrosoftGraphRegisterRegisterPostRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("externalId", m.GetExternalId())
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
func (m *MicrosoftGraphRegisterRegisterPostRequestBody) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetExternalId sets the externalId property value. The externalId property
func (m *MicrosoftGraphRegisterRegisterPostRequestBody) SetExternalId(value *string)() {
    m.externalId = value
}
