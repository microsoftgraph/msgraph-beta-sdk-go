package users

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    i084fa7ab3bba802bf5cc3b408e230cc64c167a57976e0d42c37e17154afd5b78 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/security"
)

// ItemSecurityInformationProtectionSensitivityLabelsMicrosoftGraphSecurityExtractContentLabelExtractContentLabelPostRequestBody 
type ItemSecurityInformationProtectionSensitivityLabelsMicrosoftGraphSecurityExtractContentLabelExtractContentLabelPostRequestBody struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The contentInfo property
    contentInfo i084fa7ab3bba802bf5cc3b408e230cc64c167a57976e0d42c37e17154afd5b78.ContentInfoable
}
// NewItemSecurityInformationProtectionSensitivityLabelsMicrosoftGraphSecurityExtractContentLabelExtractContentLabelPostRequestBody instantiates a new ItemSecurityInformationProtectionSensitivityLabelsMicrosoftGraphSecurityExtractContentLabelExtractContentLabelPostRequestBody and sets the default values.
func NewItemSecurityInformationProtectionSensitivityLabelsMicrosoftGraphSecurityExtractContentLabelExtractContentLabelPostRequestBody()(*ItemSecurityInformationProtectionSensitivityLabelsMicrosoftGraphSecurityExtractContentLabelExtractContentLabelPostRequestBody) {
    m := &ItemSecurityInformationProtectionSensitivityLabelsMicrosoftGraphSecurityExtractContentLabelExtractContentLabelPostRequestBody{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateItemSecurityInformationProtectionSensitivityLabelsMicrosoftGraphSecurityExtractContentLabelExtractContentLabelPostRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateItemSecurityInformationProtectionSensitivityLabelsMicrosoftGraphSecurityExtractContentLabelExtractContentLabelPostRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewItemSecurityInformationProtectionSensitivityLabelsMicrosoftGraphSecurityExtractContentLabelExtractContentLabelPostRequestBody(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *ItemSecurityInformationProtectionSensitivityLabelsMicrosoftGraphSecurityExtractContentLabelExtractContentLabelPostRequestBody) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetContentInfo gets the contentInfo property value. The contentInfo property
func (m *ItemSecurityInformationProtectionSensitivityLabelsMicrosoftGraphSecurityExtractContentLabelExtractContentLabelPostRequestBody) GetContentInfo()(i084fa7ab3bba802bf5cc3b408e230cc64c167a57976e0d42c37e17154afd5b78.ContentInfoable) {
    return m.contentInfo
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ItemSecurityInformationProtectionSensitivityLabelsMicrosoftGraphSecurityExtractContentLabelExtractContentLabelPostRequestBody) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["contentInfo"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(i084fa7ab3bba802bf5cc3b408e230cc64c167a57976e0d42c37e17154afd5b78.CreateContentInfoFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetContentInfo(val.(i084fa7ab3bba802bf5cc3b408e230cc64c167a57976e0d42c37e17154afd5b78.ContentInfoable))
        }
        return nil
    }
    return res
}
// Serialize serializes information the current object
func (m *ItemSecurityInformationProtectionSensitivityLabelsMicrosoftGraphSecurityExtractContentLabelExtractContentLabelPostRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteObjectValue("contentInfo", m.GetContentInfo())
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
func (m *ItemSecurityInformationProtectionSensitivityLabelsMicrosoftGraphSecurityExtractContentLabelExtractContentLabelPostRequestBody) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetContentInfo sets the contentInfo property value. The contentInfo property
func (m *ItemSecurityInformationProtectionSensitivityLabelsMicrosoftGraphSecurityExtractContentLabelExtractContentLabelPostRequestBody) SetContentInfo(value i084fa7ab3bba802bf5cc3b408e230cc64c167a57976e0d42c37e17154afd5b78.ContentInfoable)() {
    m.contentInfo = value
}
