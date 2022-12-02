package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// SecurityCasesEdiscoveryCasesItemReviewSetsItemQueriesItemExportPostRequestBody provides operations to call the export method.
type SecurityCasesEdiscoveryCasesItemReviewSetsItemQueriesItemExportPostRequestBody struct {
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
// NewSecurityCasesEdiscoveryCasesItemReviewSetsItemQueriesItemExportPostRequestBody instantiates a new SecurityCasesEdiscoveryCasesItemReviewSetsItemQueriesItemExportPostRequestBody and sets the default values.
func NewSecurityCasesEdiscoveryCasesItemReviewSetsItemQueriesItemExportPostRequestBody()(*SecurityCasesEdiscoveryCasesItemReviewSetsItemQueriesItemExportPostRequestBody) {
    m := &SecurityCasesEdiscoveryCasesItemReviewSetsItemQueriesItemExportPostRequestBody{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateSecurityCasesEdiscoveryCasesItemReviewSetsItemQueriesItemExportPostRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateSecurityCasesEdiscoveryCasesItemReviewSetsItemQueriesItemExportPostRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewSecurityCasesEdiscoveryCasesItemReviewSetsItemQueriesItemExportPostRequestBody(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *SecurityCasesEdiscoveryCasesItemReviewSetsItemQueriesItemExportPostRequestBody) GetAdditionalData()(map[string]interface{}) {
    return m.additionalData
}
// GetAzureBlobContainer gets the azureBlobContainer property value. The azureBlobContainer property
func (m *SecurityCasesEdiscoveryCasesItemReviewSetsItemQueriesItemExportPostRequestBody) GetAzureBlobContainer()(*string) {
    return m.azureBlobContainer
}
// GetAzureBlobToken gets the azureBlobToken property value. The azureBlobToken property
func (m *SecurityCasesEdiscoveryCasesItemReviewSetsItemQueriesItemExportPostRequestBody) GetAzureBlobToken()(*string) {
    return m.azureBlobToken
}
// GetDescription gets the description property value. The description property
func (m *SecurityCasesEdiscoveryCasesItemReviewSetsItemQueriesItemExportPostRequestBody) GetDescription()(*string) {
    return m.description
}
// GetFieldDeserializers the deserialization information for the current model
func (m *SecurityCasesEdiscoveryCasesItemReviewSetsItemQueriesItemExportPostRequestBody) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
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
func (m *SecurityCasesEdiscoveryCasesItemReviewSetsItemQueriesItemExportPostRequestBody) GetOutputName()(*string) {
    return m.outputName
}
// Serialize serializes information the current object
func (m *SecurityCasesEdiscoveryCasesItemReviewSetsItemQueriesItemExportPostRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
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
func (m *SecurityCasesEdiscoveryCasesItemReviewSetsItemQueriesItemExportPostRequestBody) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// SetAzureBlobContainer sets the azureBlobContainer property value. The azureBlobContainer property
func (m *SecurityCasesEdiscoveryCasesItemReviewSetsItemQueriesItemExportPostRequestBody) SetAzureBlobContainer(value *string)() {
    m.azureBlobContainer = value
}
// SetAzureBlobToken sets the azureBlobToken property value. The azureBlobToken property
func (m *SecurityCasesEdiscoveryCasesItemReviewSetsItemQueriesItemExportPostRequestBody) SetAzureBlobToken(value *string)() {
    m.azureBlobToken = value
}
// SetDescription sets the description property value. The description property
func (m *SecurityCasesEdiscoveryCasesItemReviewSetsItemQueriesItemExportPostRequestBody) SetDescription(value *string)() {
    m.description = value
}
// SetOutputName sets the outputName property value. The outputName property
func (m *SecurityCasesEdiscoveryCasesItemReviewSetsItemQueriesItemExportPostRequestBody) SetOutputName(value *string)() {
    m.outputName = value
}
