package generateapplepushnotificationcertificatesigningrequest

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// GenerateApplePushNotificationCertificateSigningRequestResponse provides operations to call the generateApplePushNotificationCertificateSigningRequest method.
type GenerateApplePushNotificationCertificateSigningRequestResponse struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // The value property
    value *string
}
// NewGenerateApplePushNotificationCertificateSigningRequestResponse instantiates a new generateApplePushNotificationCertificateSigningRequestResponse and sets the default values.
func NewGenerateApplePushNotificationCertificateSigningRequestResponse()(*GenerateApplePushNotificationCertificateSigningRequestResponse) {
    m := &GenerateApplePushNotificationCertificateSigningRequestResponse{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateGenerateApplePushNotificationCertificateSigningRequestResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateGenerateApplePushNotificationCertificateSigningRequestResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewGenerateApplePushNotificationCertificateSigningRequestResponse(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *GenerateApplePushNotificationCertificateSigningRequestResponse) GetAdditionalData()(map[string]interface{}) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
func (m *GenerateApplePushNotificationCertificateSigningRequestResponse) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["value"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetValue)
    return res
}
// GetValue gets the value property value. The value property
func (m *GenerateApplePushNotificationCertificateSigningRequestResponse) GetValue()(*string) {
    return m.value
}
// Serialize serializes information the current object
func (m *GenerateApplePushNotificationCertificateSigningRequestResponse) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("value", m.GetValue())
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
func (m *GenerateApplePushNotificationCertificateSigningRequestResponse) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// SetValue sets the value property value. The value property
func (m *GenerateApplePushNotificationCertificateSigningRequestResponse) SetValue(value *string)() {
    m.value = value
}
