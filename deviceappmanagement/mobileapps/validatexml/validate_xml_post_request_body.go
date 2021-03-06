package validatexml

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ValidateXmlPostRequestBody provides operations to call the validateXml method.
type ValidateXmlPostRequestBody struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // The officeConfigurationXml property
    officeConfigurationXml []byte
}
// NewValidateXmlPostRequestBody instantiates a new validateXmlPostRequestBody and sets the default values.
func NewValidateXmlPostRequestBody()(*ValidateXmlPostRequestBody) {
    m := &ValidateXmlPostRequestBody{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateValidateXmlPostRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateValidateXmlPostRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewValidateXmlPostRequestBody(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *ValidateXmlPostRequestBody) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ValidateXmlPostRequestBody) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
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
func (m *ValidateXmlPostRequestBody) GetOfficeConfigurationXml()([]byte) {
    if m == nil {
        return nil
    } else {
        return m.officeConfigurationXml
    }
}
// Serialize serializes information the current object
func (m *ValidateXmlPostRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
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
func (m *ValidateXmlPostRequestBody) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
// SetOfficeConfigurationXml sets the officeConfigurationXml property value. The officeConfigurationXml property
func (m *ValidateXmlPostRequestBody) SetOfficeConfigurationXml(value []byte)() {
    if m != nil {
        m.officeConfigurationXml = value
    }
}
