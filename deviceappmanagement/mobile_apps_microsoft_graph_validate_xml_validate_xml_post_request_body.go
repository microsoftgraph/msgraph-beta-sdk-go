package deviceappmanagement

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// MobileAppsMicrosoftGraphValidateXmlValidateXmlPostRequestBody 
type MobileAppsMicrosoftGraphValidateXmlValidateXmlPostRequestBody struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The officeConfigurationXml property
    officeConfigurationXml []byte
}
// NewMobileAppsMicrosoftGraphValidateXmlValidateXmlPostRequestBody instantiates a new MobileAppsMicrosoftGraphValidateXmlValidateXmlPostRequestBody and sets the default values.
func NewMobileAppsMicrosoftGraphValidateXmlValidateXmlPostRequestBody()(*MobileAppsMicrosoftGraphValidateXmlValidateXmlPostRequestBody) {
    m := &MobileAppsMicrosoftGraphValidateXmlValidateXmlPostRequestBody{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateMobileAppsMicrosoftGraphValidateXmlValidateXmlPostRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateMobileAppsMicrosoftGraphValidateXmlValidateXmlPostRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewMobileAppsMicrosoftGraphValidateXmlValidateXmlPostRequestBody(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *MobileAppsMicrosoftGraphValidateXmlValidateXmlPostRequestBody) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
func (m *MobileAppsMicrosoftGraphValidateXmlValidateXmlPostRequestBody) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["officeConfigurationXml"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetByteArrayValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOfficeConfigurationXml(val)
        }
        return nil
    }
    return res
}
// GetOfficeConfigurationXml gets the officeConfigurationXml property value. The officeConfigurationXml property
func (m *MobileAppsMicrosoftGraphValidateXmlValidateXmlPostRequestBody) GetOfficeConfigurationXml()([]byte) {
    return m.officeConfigurationXml
}
// Serialize serializes information the current object
func (m *MobileAppsMicrosoftGraphValidateXmlValidateXmlPostRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteByteArrayValue("officeConfigurationXml", m.GetOfficeConfigurationXml())
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
func (m *MobileAppsMicrosoftGraphValidateXmlValidateXmlPostRequestBody) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetOfficeConfigurationXml sets the officeConfigurationXml property value. The officeConfigurationXml property
func (m *MobileAppsMicrosoftGraphValidateXmlValidateXmlPostRequestBody) SetOfficeConfigurationXml(value []byte)() {
    m.officeConfigurationXml = value
}
