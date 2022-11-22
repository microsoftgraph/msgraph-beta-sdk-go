package ediscovery

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// CaseExportOperation 
type CaseExportOperation struct {
    CaseOperation
    // The name of the Azure storage location where the export will be stored. This only applies to exports stored in your own Azure storage location.
    azureBlobContainer *string
    // The SAS token for the Azure storage location.  This only applies to exports stored in your own Azure storage location.
    azureBlobToken *string
    // The description provided for the export.
    description *string
    // The options provided for the export. For more details, see reviewSet: export. Possible values are: originalFiles, text, pdfReplacement, fileInfo, tags.
    exportOptions *ExportOptions
    // The options provided that specify the structure of the export. For more details, see reviewSet: export. Possible values are: none, directory, pst.
    exportStructure *ExportFileStructure
    // The outputFolderId property
    outputFolderId *string
    // The name provided for the export.
    outputName *string
    // The review set the content is being exported from.
    reviewSet ReviewSetable
}
// NewCaseExportOperation instantiates a new CaseExportOperation and sets the default values.
func NewCaseExportOperation()(*CaseExportOperation) {
    m := &CaseExportOperation{
        CaseOperation: *NewCaseOperation(),
    }
    return m
}
// CreateCaseExportOperationFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateCaseExportOperationFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewCaseExportOperation(), nil
}
// GetAzureBlobContainer gets the azureBlobContainer property value. The name of the Azure storage location where the export will be stored. This only applies to exports stored in your own Azure storage location.
func (m *CaseExportOperation) GetAzureBlobContainer()(*string) {
    return m.azureBlobContainer
}
// GetAzureBlobToken gets the azureBlobToken property value. The SAS token for the Azure storage location.  This only applies to exports stored in your own Azure storage location.
func (m *CaseExportOperation) GetAzureBlobToken()(*string) {
    return m.azureBlobToken
}
// GetDescription gets the description property value. The description provided for the export.
func (m *CaseExportOperation) GetDescription()(*string) {
    return m.description
}
// GetExportOptions gets the exportOptions property value. The options provided for the export. For more details, see reviewSet: export. Possible values are: originalFiles, text, pdfReplacement, fileInfo, tags.
func (m *CaseExportOperation) GetExportOptions()(*ExportOptions) {
    return m.exportOptions
}
// GetExportStructure gets the exportStructure property value. The options provided that specify the structure of the export. For more details, see reviewSet: export. Possible values are: none, directory, pst.
func (m *CaseExportOperation) GetExportStructure()(*ExportFileStructure) {
    return m.exportStructure
}
// GetFieldDeserializers the deserialization information for the current model
func (m *CaseExportOperation) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.CaseOperation.GetFieldDeserializers()
    res["azureBlobContainer"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetAzureBlobContainer)
    res["azureBlobToken"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetAzureBlobToken)
    res["description"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetDescription)
    res["exportOptions"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetEnumValue(ParseExportOptions , m.SetExportOptions)
    res["exportStructure"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetEnumValue(ParseExportFileStructure , m.SetExportStructure)
    res["outputFolderId"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetOutputFolderId)
    res["outputName"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetOutputName)
    res["reviewSet"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetObjectValue(CreateReviewSetFromDiscriminatorValue , m.SetReviewSet)
    return res
}
// GetOutputFolderId gets the outputFolderId property value. The outputFolderId property
func (m *CaseExportOperation) GetOutputFolderId()(*string) {
    return m.outputFolderId
}
// GetOutputName gets the outputName property value. The name provided for the export.
func (m *CaseExportOperation) GetOutputName()(*string) {
    return m.outputName
}
// GetReviewSet gets the reviewSet property value. The review set the content is being exported from.
func (m *CaseExportOperation) GetReviewSet()(ReviewSetable) {
    return m.reviewSet
}
// Serialize serializes information the current object
func (m *CaseExportOperation) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.CaseOperation.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteStringValue("azureBlobContainer", m.GetAzureBlobContainer())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("azureBlobToken", m.GetAzureBlobToken())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("description", m.GetDescription())
        if err != nil {
            return err
        }
    }
    if m.GetExportOptions() != nil {
        cast := (*m.GetExportOptions()).String()
        err = writer.WriteStringValue("exportOptions", &cast)
        if err != nil {
            return err
        }
    }
    if m.GetExportStructure() != nil {
        cast := (*m.GetExportStructure()).String()
        err = writer.WriteStringValue("exportStructure", &cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("outputFolderId", m.GetOutputFolderId())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("outputName", m.GetOutputName())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("reviewSet", m.GetReviewSet())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAzureBlobContainer sets the azureBlobContainer property value. The name of the Azure storage location where the export will be stored. This only applies to exports stored in your own Azure storage location.
func (m *CaseExportOperation) SetAzureBlobContainer(value *string)() {
    m.azureBlobContainer = value
}
// SetAzureBlobToken sets the azureBlobToken property value. The SAS token for the Azure storage location.  This only applies to exports stored in your own Azure storage location.
func (m *CaseExportOperation) SetAzureBlobToken(value *string)() {
    m.azureBlobToken = value
}
// SetDescription sets the description property value. The description provided for the export.
func (m *CaseExportOperation) SetDescription(value *string)() {
    m.description = value
}
// SetExportOptions sets the exportOptions property value. The options provided for the export. For more details, see reviewSet: export. Possible values are: originalFiles, text, pdfReplacement, fileInfo, tags.
func (m *CaseExportOperation) SetExportOptions(value *ExportOptions)() {
    m.exportOptions = value
}
// SetExportStructure sets the exportStructure property value. The options provided that specify the structure of the export. For more details, see reviewSet: export. Possible values are: none, directory, pst.
func (m *CaseExportOperation) SetExportStructure(value *ExportFileStructure)() {
    m.exportStructure = value
}
// SetOutputFolderId sets the outputFolderId property value. The outputFolderId property
func (m *CaseExportOperation) SetOutputFolderId(value *string)() {
    m.outputFolderId = value
}
// SetOutputName sets the outputName property value. The name provided for the export.
func (m *CaseExportOperation) SetOutputName(value *string)() {
    m.outputName = value
}
// SetReviewSet sets the reviewSet property value. The review set the content is being exported from.
func (m *CaseExportOperation) SetReviewSet(value ReviewSetable)() {
    m.reviewSet = value
}
