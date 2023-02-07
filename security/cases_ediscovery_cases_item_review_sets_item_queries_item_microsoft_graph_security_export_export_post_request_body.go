package security

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    i084fa7ab3bba802bf5cc3b408e230cc64c167a57976e0d42c37e17154afd5b78 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/security"
)

// CasesEdiscoveryCasesItemReviewSetsItemQueriesItemMicrosoftGraphSecurityExportExportPostRequestBody 
type CasesEdiscoveryCasesItemReviewSetsItemQueriesItemMicrosoftGraphSecurityExportExportPostRequestBody struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The azureBlobContainer property
    azureBlobContainer *string
    // The azureBlobToken property
    azureBlobToken *string
    // The description property
    description *string
    // The exportOptions property
    exportOptions *i084fa7ab3bba802bf5cc3b408e230cc64c167a57976e0d42c37e17154afd5b78.ExportOptions
    // The exportStructure property
    exportStructure *i084fa7ab3bba802bf5cc3b408e230cc64c167a57976e0d42c37e17154afd5b78.ExportFileStructure
    // The outputName property
    outputName *string
}
// NewCasesEdiscoveryCasesItemReviewSetsItemQueriesItemMicrosoftGraphSecurityExportExportPostRequestBody instantiates a new CasesEdiscoveryCasesItemReviewSetsItemQueriesItemMicrosoftGraphSecurityExportExportPostRequestBody and sets the default values.
func NewCasesEdiscoveryCasesItemReviewSetsItemQueriesItemMicrosoftGraphSecurityExportExportPostRequestBody()(*CasesEdiscoveryCasesItemReviewSetsItemQueriesItemMicrosoftGraphSecurityExportExportPostRequestBody) {
    m := &CasesEdiscoveryCasesItemReviewSetsItemQueriesItemMicrosoftGraphSecurityExportExportPostRequestBody{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateCasesEdiscoveryCasesItemReviewSetsItemQueriesItemMicrosoftGraphSecurityExportExportPostRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateCasesEdiscoveryCasesItemReviewSetsItemQueriesItemMicrosoftGraphSecurityExportExportPostRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewCasesEdiscoveryCasesItemReviewSetsItemQueriesItemMicrosoftGraphSecurityExportExportPostRequestBody(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *CasesEdiscoveryCasesItemReviewSetsItemQueriesItemMicrosoftGraphSecurityExportExportPostRequestBody) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetAzureBlobContainer gets the azureBlobContainer property value. The azureBlobContainer property
func (m *CasesEdiscoveryCasesItemReviewSetsItemQueriesItemMicrosoftGraphSecurityExportExportPostRequestBody) GetAzureBlobContainer()(*string) {
    return m.azureBlobContainer
}
// GetAzureBlobToken gets the azureBlobToken property value. The azureBlobToken property
func (m *CasesEdiscoveryCasesItemReviewSetsItemQueriesItemMicrosoftGraphSecurityExportExportPostRequestBody) GetAzureBlobToken()(*string) {
    return m.azureBlobToken
}
// GetDescription gets the description property value. The description property
func (m *CasesEdiscoveryCasesItemReviewSetsItemQueriesItemMicrosoftGraphSecurityExportExportPostRequestBody) GetDescription()(*string) {
    return m.description
}
// GetExportOptions gets the exportOptions property value. The exportOptions property
func (m *CasesEdiscoveryCasesItemReviewSetsItemQueriesItemMicrosoftGraphSecurityExportExportPostRequestBody) GetExportOptions()(*i084fa7ab3bba802bf5cc3b408e230cc64c167a57976e0d42c37e17154afd5b78.ExportOptions) {
    return m.exportOptions
}
// GetExportStructure gets the exportStructure property value. The exportStructure property
func (m *CasesEdiscoveryCasesItemReviewSetsItemQueriesItemMicrosoftGraphSecurityExportExportPostRequestBody) GetExportStructure()(*i084fa7ab3bba802bf5cc3b408e230cc64c167a57976e0d42c37e17154afd5b78.ExportFileStructure) {
    return m.exportStructure
}
// GetFieldDeserializers the deserialization information for the current model
func (m *CasesEdiscoveryCasesItemReviewSetsItemQueriesItemMicrosoftGraphSecurityExportExportPostRequestBody) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
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
    res["exportOptions"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(i084fa7ab3bba802bf5cc3b408e230cc64c167a57976e0d42c37e17154afd5b78.ParseExportOptions)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetExportOptions(val.(*i084fa7ab3bba802bf5cc3b408e230cc64c167a57976e0d42c37e17154afd5b78.ExportOptions))
        }
        return nil
    }
    res["exportStructure"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(i084fa7ab3bba802bf5cc3b408e230cc64c167a57976e0d42c37e17154afd5b78.ParseExportFileStructure)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetExportStructure(val.(*i084fa7ab3bba802bf5cc3b408e230cc64c167a57976e0d42c37e17154afd5b78.ExportFileStructure))
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
func (m *CasesEdiscoveryCasesItemReviewSetsItemQueriesItemMicrosoftGraphSecurityExportExportPostRequestBody) GetOutputName()(*string) {
    return m.outputName
}
// Serialize serializes information the current object
func (m *CasesEdiscoveryCasesItemReviewSetsItemQueriesItemMicrosoftGraphSecurityExportExportPostRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
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
    if m.GetExportOptions() != nil {
        cast := (*m.GetExportOptions()).String()
        err := writer.WriteStringValue("exportOptions", &cast)
        if err != nil {
            return err
        }
    }
    if m.GetExportStructure() != nil {
        cast := (*m.GetExportStructure()).String()
        err := writer.WriteStringValue("exportStructure", &cast)
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
func (m *CasesEdiscoveryCasesItemReviewSetsItemQueriesItemMicrosoftGraphSecurityExportExportPostRequestBody) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetAzureBlobContainer sets the azureBlobContainer property value. The azureBlobContainer property
func (m *CasesEdiscoveryCasesItemReviewSetsItemQueriesItemMicrosoftGraphSecurityExportExportPostRequestBody) SetAzureBlobContainer(value *string)() {
    m.azureBlobContainer = value
}
// SetAzureBlobToken sets the azureBlobToken property value. The azureBlobToken property
func (m *CasesEdiscoveryCasesItemReviewSetsItemQueriesItemMicrosoftGraphSecurityExportExportPostRequestBody) SetAzureBlobToken(value *string)() {
    m.azureBlobToken = value
}
// SetDescription sets the description property value. The description property
func (m *CasesEdiscoveryCasesItemReviewSetsItemQueriesItemMicrosoftGraphSecurityExportExportPostRequestBody) SetDescription(value *string)() {
    m.description = value
}
// SetExportOptions sets the exportOptions property value. The exportOptions property
func (m *CasesEdiscoveryCasesItemReviewSetsItemQueriesItemMicrosoftGraphSecurityExportExportPostRequestBody) SetExportOptions(value *i084fa7ab3bba802bf5cc3b408e230cc64c167a57976e0d42c37e17154afd5b78.ExportOptions)() {
    m.exportOptions = value
}
// SetExportStructure sets the exportStructure property value. The exportStructure property
func (m *CasesEdiscoveryCasesItemReviewSetsItemQueriesItemMicrosoftGraphSecurityExportExportPostRequestBody) SetExportStructure(value *i084fa7ab3bba802bf5cc3b408e230cc64c167a57976e0d42c37e17154afd5b78.ExportFileStructure)() {
    m.exportStructure = value
}
// SetOutputName sets the outputName property value. The outputName property
func (m *CasesEdiscoveryCasesItemReviewSetsItemQueriesItemMicrosoftGraphSecurityExportExportPostRequestBody) SetOutputName(value *string)() {
    m.outputName = value
}
