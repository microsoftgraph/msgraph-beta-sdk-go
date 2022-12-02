package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ComplianceEdiscoveryCasesItemReviewSetsItemExportPostRequestBody provides operations to call the export method.
type ComplianceEdiscoveryCasesItemReviewSetsItemExportPostRequestBody struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // The azureBlobContainer property
    azureBlobContainer *string
    // The azureBlobToken property
    azureBlobToken *string
    // The description property
    description *string
    // The outputName property
    outputName *string
}
// NewComplianceEdiscoveryCasesItemReviewSetsItemExportPostRequestBody instantiates a new ComplianceEdiscoveryCasesItemReviewSetsItemExportPostRequestBody and sets the default values.
func NewComplianceEdiscoveryCasesItemReviewSetsItemExportPostRequestBody()(*ComplianceEdiscoveryCasesItemReviewSetsItemExportPostRequestBody) {
    m := &ComplianceEdiscoveryCasesItemReviewSetsItemExportPostRequestBody{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateComplianceEdiscoveryCasesItemReviewSetsItemExportPostRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateComplianceEdiscoveryCasesItemReviewSetsItemExportPostRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewComplianceEdiscoveryCasesItemReviewSetsItemExportPostRequestBody(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *ComplianceEdiscoveryCasesItemReviewSetsItemExportPostRequestBody) GetAdditionalData()(map[string]interface{}) {
    return m.additionalData
}
// GetAzureBlobContainer gets the azureBlobContainer property value. The azureBlobContainer property
func (m *ComplianceEdiscoveryCasesItemReviewSetsItemExportPostRequestBody) GetAzureBlobContainer()(*string) {
    return m.azureBlobContainer
}
// GetAzureBlobToken gets the azureBlobToken property value. The azureBlobToken property
func (m *ComplianceEdiscoveryCasesItemReviewSetsItemExportPostRequestBody) GetAzureBlobToken()(*string) {
    return m.azureBlobToken
}
// GetDescription gets the description property value. The description property
func (m *ComplianceEdiscoveryCasesItemReviewSetsItemExportPostRequestBody) GetDescription()(*string) {
    return m.description
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ComplianceEdiscoveryCasesItemReviewSetsItemExportPostRequestBody) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["azureBlobContainer"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAzureBlobContainer(val)
        }
        return nil
    }
    res["azureBlobToken"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAzureBlobToken(val)
        }
        return nil
    }
    res["description"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDescription(val)
        }
        return nil
    }
    res["outputName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOutputName(val)
        }
        return nil
    }
    return res
}
// GetOutputName gets the outputName property value. The outputName property
func (m *ComplianceEdiscoveryCasesItemReviewSetsItemExportPostRequestBody) GetOutputName()(*string) {
    return m.outputName
}
// Serialize serializes information the current object
func (m *ComplianceEdiscoveryCasesItemReviewSetsItemExportPostRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("azureBlobContainer", m.GetAzureBlobContainer())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("azureBlobToken", m.GetAzureBlobToken())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("description", m.GetDescription())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("outputName", m.GetOutputName())
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
func (m *ComplianceEdiscoveryCasesItemReviewSetsItemExportPostRequestBody) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// SetAzureBlobContainer sets the azureBlobContainer property value. The azureBlobContainer property
func (m *ComplianceEdiscoveryCasesItemReviewSetsItemExportPostRequestBody) SetAzureBlobContainer(value *string)() {
    m.azureBlobContainer = value
}
// SetAzureBlobToken sets the azureBlobToken property value. The azureBlobToken property
func (m *ComplianceEdiscoveryCasesItemReviewSetsItemExportPostRequestBody) SetAzureBlobToken(value *string)() {
    m.azureBlobToken = value
}
// SetDescription sets the description property value. The description property
func (m *ComplianceEdiscoveryCasesItemReviewSetsItemExportPostRequestBody) SetDescription(value *string)() {
    m.description = value
}
// SetOutputName sets the outputName property value. The outputName property
func (m *ComplianceEdiscoveryCasesItemReviewSetsItemExportPostRequestBody) SetOutputName(value *string)() {
    m.outputName = value
}
