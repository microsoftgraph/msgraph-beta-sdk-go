package security

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// EdiscoveryExportOperation 
type EdiscoveryExportOperation struct {
    CaseOperation
}
// NewEdiscoveryExportOperation instantiates a new ediscoveryExportOperation and sets the default values.
func NewEdiscoveryExportOperation()(*EdiscoveryExportOperation) {
    m := &EdiscoveryExportOperation{
        CaseOperation: *NewCaseOperation(),
    }
    return m
}
// CreateEdiscoveryExportOperationFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateEdiscoveryExportOperationFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewEdiscoveryExportOperation(), nil
}
// GetAzureBlobContainer gets the azureBlobContainer property value. The azureBlobContainer property
func (m *EdiscoveryExportOperation) GetAzureBlobContainer()(*string) {
    val, err := m.GetBackingStore().Get("azureBlobContainer")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetAzureBlobToken gets the azureBlobToken property value. The azureBlobToken property
func (m *EdiscoveryExportOperation) GetAzureBlobToken()(*string) {
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
func (m *EdiscoveryExportOperation) GetDescription()(*string) {
    val, err := m.GetBackingStore().Get("description")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetExportFileMetadata gets the exportFileMetadata property value. The exportFileMetadata property
func (m *EdiscoveryExportOperation) GetExportFileMetadata()([]ExportFileMetadataable) {
    val, err := m.GetBackingStore().Get("exportFileMetadata")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]ExportFileMetadataable)
    }
    return nil
}
// GetExportOptions gets the exportOptions property value. The options provided for the export. For more information, see reviewSet: export. Possible values are: originalFiles, text, pdfReplacement, fileInfo, tags. The fileInfo member is deprecated and will stop returning data on April 30, 2023. Going forward, the summary and load file are always included.
func (m *EdiscoveryExportOperation) GetExportOptions()(*ExportOptions) {
    val, err := m.GetBackingStore().Get("exportOptions")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*ExportOptions)
    }
    return nil
}
// GetExportStructure gets the exportStructure property value. The options provided that specify the structure of the export. For more information, see reviewSet: export. Possible values are: none, directory, pst.
func (m *EdiscoveryExportOperation) GetExportStructure()(*ExportFileStructure) {
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
func (m *EdiscoveryExportOperation) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
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
    res["exportFileMetadata"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateExportFileMetadataFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]ExportFileMetadataable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(ExportFileMetadataable)
                }
            }
            m.SetExportFileMetadata(res)
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
        val, err := n.GetObjectValue(CreateEdiscoveryReviewSetFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetReviewSet(val.(EdiscoveryReviewSetable))
        }
        return nil
    }
    res["reviewSetQuery"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateEdiscoveryReviewSetQueryFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetReviewSetQuery(val.(EdiscoveryReviewSetQueryable))
        }
        return nil
    }
    return res
}
// GetOutputFolderId gets the outputFolderId property value. The outputFolderId property
func (m *EdiscoveryExportOperation) GetOutputFolderId()(*string) {
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
func (m *EdiscoveryExportOperation) GetOutputName()(*string) {
    val, err := m.GetBackingStore().Get("outputName")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetReviewSet gets the reviewSet property value. Review set from where documents are exported.
func (m *EdiscoveryExportOperation) GetReviewSet()(EdiscoveryReviewSetable) {
    val, err := m.GetBackingStore().Get("reviewSet")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(EdiscoveryReviewSetable)
    }
    return nil
}
// GetReviewSetQuery gets the reviewSetQuery property value. The review set query that is used to filter the documents for export.
func (m *EdiscoveryExportOperation) GetReviewSetQuery()(EdiscoveryReviewSetQueryable) {
    val, err := m.GetBackingStore().Get("reviewSetQuery")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(EdiscoveryReviewSetQueryable)
    }
    return nil
}
// Serialize serializes information the current object
func (m *EdiscoveryExportOperation) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
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
    if m.GetExportFileMetadata() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetExportFileMetadata()))
        for i, v := range m.GetExportFileMetadata() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err = writer.WriteCollectionOfObjectValues("exportFileMetadata", cast)
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
    {
        err = writer.WriteObjectValue("reviewSetQuery", m.GetReviewSetQuery())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAzureBlobContainer sets the azureBlobContainer property value. The azureBlobContainer property
func (m *EdiscoveryExportOperation) SetAzureBlobContainer(value *string)() {
    err := m.GetBackingStore().Set("azureBlobContainer", value)
    if err != nil {
        panic(err)
    }
}
// SetAzureBlobToken sets the azureBlobToken property value. The azureBlobToken property
func (m *EdiscoveryExportOperation) SetAzureBlobToken(value *string)() {
    err := m.GetBackingStore().Set("azureBlobToken", value)
    if err != nil {
        panic(err)
    }
}
// SetDescription sets the description property value. The description provided for the export.
func (m *EdiscoveryExportOperation) SetDescription(value *string)() {
    err := m.GetBackingStore().Set("description", value)
    if err != nil {
        panic(err)
    }
}
// SetExportFileMetadata sets the exportFileMetadata property value. The exportFileMetadata property
func (m *EdiscoveryExportOperation) SetExportFileMetadata(value []ExportFileMetadataable)() {
    err := m.GetBackingStore().Set("exportFileMetadata", value)
    if err != nil {
        panic(err)
    }
}
// SetExportOptions sets the exportOptions property value. The options provided for the export. For more information, see reviewSet: export. Possible values are: originalFiles, text, pdfReplacement, fileInfo, tags. The fileInfo member is deprecated and will stop returning data on April 30, 2023. Going forward, the summary and load file are always included.
func (m *EdiscoveryExportOperation) SetExportOptions(value *ExportOptions)() {
    err := m.GetBackingStore().Set("exportOptions", value)
    if err != nil {
        panic(err)
    }
}
// SetExportStructure sets the exportStructure property value. The options provided that specify the structure of the export. For more information, see reviewSet: export. Possible values are: none, directory, pst.
func (m *EdiscoveryExportOperation) SetExportStructure(value *ExportFileStructure)() {
    err := m.GetBackingStore().Set("exportStructure", value)
    if err != nil {
        panic(err)
    }
}
// SetOutputFolderId sets the outputFolderId property value. The outputFolderId property
func (m *EdiscoveryExportOperation) SetOutputFolderId(value *string)() {
    err := m.GetBackingStore().Set("outputFolderId", value)
    if err != nil {
        panic(err)
    }
}
// SetOutputName sets the outputName property value. The name provided for the export.
func (m *EdiscoveryExportOperation) SetOutputName(value *string)() {
    err := m.GetBackingStore().Set("outputName", value)
    if err != nil {
        panic(err)
    }
}
// SetReviewSet sets the reviewSet property value. Review set from where documents are exported.
func (m *EdiscoveryExportOperation) SetReviewSet(value EdiscoveryReviewSetable)() {
    err := m.GetBackingStore().Set("reviewSet", value)
    if err != nil {
        panic(err)
    }
}
// SetReviewSetQuery sets the reviewSetQuery property value. The review set query that is used to filter the documents for export.
func (m *EdiscoveryExportOperation) SetReviewSetQuery(value EdiscoveryReviewSetQueryable)() {
    err := m.GetBackingStore().Set("reviewSetQuery", value)
    if err != nil {
        panic(err)
    }
}
// EdiscoveryExportOperationable 
type EdiscoveryExportOperationable interface {
    CaseOperationable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAzureBlobContainer()(*string)
    GetAzureBlobToken()(*string)
    GetDescription()(*string)
    GetExportFileMetadata()([]ExportFileMetadataable)
    GetExportOptions()(*ExportOptions)
    GetExportStructure()(*ExportFileStructure)
    GetOutputFolderId()(*string)
    GetOutputName()(*string)
    GetReviewSet()(EdiscoveryReviewSetable)
    GetReviewSetQuery()(EdiscoveryReviewSetQueryable)
    SetAzureBlobContainer(value *string)()
    SetAzureBlobToken(value *string)()
    SetDescription(value *string)()
    SetExportFileMetadata(value []ExportFileMetadataable)()
    SetExportOptions(value *ExportOptions)()
    SetExportStructure(value *ExportFileStructure)()
    SetOutputFolderId(value *string)()
    SetOutputName(value *string)()
    SetReviewSet(value EdiscoveryReviewSetable)()
    SetReviewSetQuery(value EdiscoveryReviewSetQueryable)()
}
