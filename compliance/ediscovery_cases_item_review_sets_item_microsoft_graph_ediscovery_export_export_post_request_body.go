package compliance

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    ic154d683aa4025ee28853b9c1a3c35cd1f093a1c4542feba4c07682e2752db13 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/ediscovery"
)

// EdiscoveryCasesItemReviewSetsItemMicrosoftGraphEdiscoveryExportExportPostRequestBody 
type EdiscoveryCasesItemReviewSetsItemMicrosoftGraphEdiscoveryExportExportPostRequestBody struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The azureBlobContainer property
    azureBlobContainer *string
    // The azureBlobToken property
    azureBlobToken *string
    // The description property
    description *string
    // The exportOptions property
    exportOptions *ic154d683aa4025ee28853b9c1a3c35cd1f093a1c4542feba4c07682e2752db13.ExportOptions
    // The exportStructure property
    exportStructure *ic154d683aa4025ee28853b9c1a3c35cd1f093a1c4542feba4c07682e2752db13.ExportFileStructure
    // The outputName property
    outputName *string
}
// NewEdiscoveryCasesItemReviewSetsItemMicrosoftGraphEdiscoveryExportExportPostRequestBody instantiates a new EdiscoveryCasesItemReviewSetsItemMicrosoftGraphEdiscoveryExportExportPostRequestBody and sets the default values.
func NewEdiscoveryCasesItemReviewSetsItemMicrosoftGraphEdiscoveryExportExportPostRequestBody()(*EdiscoveryCasesItemReviewSetsItemMicrosoftGraphEdiscoveryExportExportPostRequestBody) {
    m := &EdiscoveryCasesItemReviewSetsItemMicrosoftGraphEdiscoveryExportExportPostRequestBody{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateEdiscoveryCasesItemReviewSetsItemMicrosoftGraphEdiscoveryExportExportPostRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateEdiscoveryCasesItemReviewSetsItemMicrosoftGraphEdiscoveryExportExportPostRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewEdiscoveryCasesItemReviewSetsItemMicrosoftGraphEdiscoveryExportExportPostRequestBody(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *EdiscoveryCasesItemReviewSetsItemMicrosoftGraphEdiscoveryExportExportPostRequestBody) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetAzureBlobContainer gets the azureBlobContainer property value. The azureBlobContainer property
func (m *EdiscoveryCasesItemReviewSetsItemMicrosoftGraphEdiscoveryExportExportPostRequestBody) GetAzureBlobContainer()(*string) {
    return m.azureBlobContainer
}
// GetAzureBlobToken gets the azureBlobToken property value. The azureBlobToken property
func (m *EdiscoveryCasesItemReviewSetsItemMicrosoftGraphEdiscoveryExportExportPostRequestBody) GetAzureBlobToken()(*string) {
    return m.azureBlobToken
}
// GetDescription gets the description property value. The description property
func (m *EdiscoveryCasesItemReviewSetsItemMicrosoftGraphEdiscoveryExportExportPostRequestBody) GetDescription()(*string) {
    return m.description
}
// GetExportOptions gets the exportOptions property value. The exportOptions property
func (m *EdiscoveryCasesItemReviewSetsItemMicrosoftGraphEdiscoveryExportExportPostRequestBody) GetExportOptions()(*ic154d683aa4025ee28853b9c1a3c35cd1f093a1c4542feba4c07682e2752db13.ExportOptions) {
    return m.exportOptions
}
// GetExportStructure gets the exportStructure property value. The exportStructure property
func (m *EdiscoveryCasesItemReviewSetsItemMicrosoftGraphEdiscoveryExportExportPostRequestBody) GetExportStructure()(*ic154d683aa4025ee28853b9c1a3c35cd1f093a1c4542feba4c07682e2752db13.ExportFileStructure) {
    return m.exportStructure
}
// GetFieldDeserializers the deserialization information for the current model
func (m *EdiscoveryCasesItemReviewSetsItemMicrosoftGraphEdiscoveryExportExportPostRequestBody) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
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
        val, err := n.GetEnumValue(ic154d683aa4025ee28853b9c1a3c35cd1f093a1c4542feba4c07682e2752db13.ParseExportOptions)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetExportOptions(val.(*ic154d683aa4025ee28853b9c1a3c35cd1f093a1c4542feba4c07682e2752db13.ExportOptions))
        }
        return nil
    }
    res["exportStructure"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ic154d683aa4025ee28853b9c1a3c35cd1f093a1c4542feba4c07682e2752db13.ParseExportFileStructure)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetExportStructure(val.(*ic154d683aa4025ee28853b9c1a3c35cd1f093a1c4542feba4c07682e2752db13.ExportFileStructure))
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
func (m *EdiscoveryCasesItemReviewSetsItemMicrosoftGraphEdiscoveryExportExportPostRequestBody) GetOutputName()(*string) {
    return m.outputName
}
// Serialize serializes information the current object
func (m *EdiscoveryCasesItemReviewSetsItemMicrosoftGraphEdiscoveryExportExportPostRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
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
func (m *EdiscoveryCasesItemReviewSetsItemMicrosoftGraphEdiscoveryExportExportPostRequestBody) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetAzureBlobContainer sets the azureBlobContainer property value. The azureBlobContainer property
func (m *EdiscoveryCasesItemReviewSetsItemMicrosoftGraphEdiscoveryExportExportPostRequestBody) SetAzureBlobContainer(value *string)() {
    m.azureBlobContainer = value
}
// SetAzureBlobToken sets the azureBlobToken property value. The azureBlobToken property
func (m *EdiscoveryCasesItemReviewSetsItemMicrosoftGraphEdiscoveryExportExportPostRequestBody) SetAzureBlobToken(value *string)() {
    m.azureBlobToken = value
}
// SetDescription sets the description property value. The description property
func (m *EdiscoveryCasesItemReviewSetsItemMicrosoftGraphEdiscoveryExportExportPostRequestBody) SetDescription(value *string)() {
    m.description = value
}
// SetExportOptions sets the exportOptions property value. The exportOptions property
func (m *EdiscoveryCasesItemReviewSetsItemMicrosoftGraphEdiscoveryExportExportPostRequestBody) SetExportOptions(value *ic154d683aa4025ee28853b9c1a3c35cd1f093a1c4542feba4c07682e2752db13.ExportOptions)() {
    m.exportOptions = value
}
// SetExportStructure sets the exportStructure property value. The exportStructure property
func (m *EdiscoveryCasesItemReviewSetsItemMicrosoftGraphEdiscoveryExportExportPostRequestBody) SetExportStructure(value *ic154d683aa4025ee28853b9c1a3c35cd1f093a1c4542feba4c07682e2752db13.ExportFileStructure)() {
    m.exportStructure = value
}
// SetOutputName sets the outputName property value. The outputName property
func (m *EdiscoveryCasesItemReviewSetsItemMicrosoftGraphEdiscoveryExportExportPostRequestBody) SetOutputName(value *string)() {
    m.outputName = value
}
