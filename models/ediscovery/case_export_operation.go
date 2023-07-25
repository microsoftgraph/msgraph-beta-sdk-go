package ediscovery

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// CaseExportOperation 
type CaseExportOperation struct {
    CaseOperation
}
// NewCaseExportOperation instantiates a new caseExportOperation and sets the default values.
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
    val, err := m.GetBackingStore().Get("azureBlobContainer")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetAzureBlobToken gets the azureBlobToken property value. The SAS token for the Azure storage location.  This only applies to exports stored in your own Azure storage location.
func (m *CaseExportOperation) GetAzureBlobToken()(*string) {
    val, err := m.GetBackingStore().Get("azureBlobToken")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetDescription gets the description property value. The description provided for the export.
func (m *CaseExportOperation) GetDescription()(*string) {
    val, err := m.GetBackingStore().Get("description")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetExportOptions gets the exportOptions property value. The options provided for the export. For more details, see reviewSet: export. Possible values are: originalFiles, text, pdfReplacement, fileInfo, tags.
func (m *CaseExportOperation) GetExportOptions()(*ExportOptions) {
    val, err := m.GetBackingStore().Get("exportOptions")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*ExportOptions)
    }
    return nil
}
// GetExportStructure gets the exportStructure property value. The options provided that specify the structure of the export. For more details, see reviewSet: export. Possible values are: none, directory, pst.
func (m *CaseExportOperation) GetExportStructure()(*ExportFileStructure) {
    val, err := m.GetBackingStore().Get("exportStructure")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*ExportFileStructure)
    }
    return nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *CaseExportOperation) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.CaseOperation.GetFieldDeserializers()
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
        val, err := n.GetEnumValue(ParseExportOptions)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetExportOptions(val.(*ExportOptions))
        }
        return nil
    }
    res["exportStructure"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseExportFileStructure)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetExportStructure(val.(*ExportFileStructure))
        }
        return nil
    }
    res["outputFolderId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOutputFolderId(val)
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
    res["reviewSet"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateReviewSetFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetReviewSet(val.(ReviewSetable))
        }
        return nil
    }
    return res
}
// GetOutputFolderId gets the outputFolderId property value. The outputFolderId property
func (m *CaseExportOperation) GetOutputFolderId()(*string) {
    val, err := m.GetBackingStore().Get("outputFolderId")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetOutputName gets the outputName property value. The name provided for the export.
func (m *CaseExportOperation) GetOutputName()(*string) {
    val, err := m.GetBackingStore().Get("outputName")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetReviewSet gets the reviewSet property value. The review set the content is being exported from.
func (m *CaseExportOperation) GetReviewSet()(ReviewSetable) {
    val, err := m.GetBackingStore().Get("reviewSet")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(ReviewSetable)
    }
    return nil
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
    err := m.GetBackingStore().Set("azureBlobContainer", value)
    if err != nil {
        panic(err)
    }
}
// SetAzureBlobToken sets the azureBlobToken property value. The SAS token for the Azure storage location.  This only applies to exports stored in your own Azure storage location.
func (m *CaseExportOperation) SetAzureBlobToken(value *string)() {
    err := m.GetBackingStore().Set("azureBlobToken", value)
    if err != nil {
        panic(err)
    }
}
// SetDescription sets the description property value. The description provided for the export.
func (m *CaseExportOperation) SetDescription(value *string)() {
    err := m.GetBackingStore().Set("description", value)
    if err != nil {
        panic(err)
    }
}
// SetExportOptions sets the exportOptions property value. The options provided for the export. For more details, see reviewSet: export. Possible values are: originalFiles, text, pdfReplacement, fileInfo, tags.
func (m *CaseExportOperation) SetExportOptions(value *ExportOptions)() {
    err := m.GetBackingStore().Set("exportOptions", value)
    if err != nil {
        panic(err)
    }
}
// SetExportStructure sets the exportStructure property value. The options provided that specify the structure of the export. For more details, see reviewSet: export. Possible values are: none, directory, pst.
func (m *CaseExportOperation) SetExportStructure(value *ExportFileStructure)() {
    err := m.GetBackingStore().Set("exportStructure", value)
    if err != nil {
        panic(err)
    }
}
// SetOutputFolderId sets the outputFolderId property value. The outputFolderId property
func (m *CaseExportOperation) SetOutputFolderId(value *string)() {
    err := m.GetBackingStore().Set("outputFolderId", value)
    if err != nil {
        panic(err)
    }
}
// SetOutputName sets the outputName property value. The name provided for the export.
func (m *CaseExportOperation) SetOutputName(value *string)() {
    err := m.GetBackingStore().Set("outputName", value)
    if err != nil {
        panic(err)
    }
}
// SetReviewSet sets the reviewSet property value. The review set the content is being exported from.
func (m *CaseExportOperation) SetReviewSet(value ReviewSetable)() {
    err := m.GetBackingStore().Set("reviewSet", value)
    if err != nil {
        panic(err)
    }
}
// CaseExportOperationable 
type CaseExportOperationable interface {
    CaseOperationable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAzureBlobContainer()(*string)
    GetAzureBlobToken()(*string)
    GetDescription()(*string)
    GetExportOptions()(*ExportOptions)
    GetExportStructure()(*ExportFileStructure)
    GetOutputFolderId()(*string)
    GetOutputName()(*string)
    GetReviewSet()(ReviewSetable)
    SetAzureBlobContainer(value *string)()
    SetAzureBlobToken(value *string)()
    SetDescription(value *string)()
    SetExportOptions(value *ExportOptions)()
    SetExportStructure(value *ExportFileStructure)()
    SetOutputFolderId(value *string)()
    SetOutputName(value *string)()
    SetReviewSet(value ReviewSetable)()
}
