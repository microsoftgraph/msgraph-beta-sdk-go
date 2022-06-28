package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// CaseExportOperation 
type CaseExportOperation struct {
    CaseOperation
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // The name of the Azure storage location where the export will be stored. This only applies to exports stored in your own Azure storage location.
    azureBlobContainer *string
    // The SAS token for the Azure storage location.  This only applies to exports stored in your own Azure storage location.
    azureBlobToken *string
    // The description provided for the export.
    description *string
    // The outputFolderId property
    outputFolderId *string
    // The name provided for the export.
    outputName *string
}
// NewCaseExportOperation instantiates a new CaseExportOperation and sets the default values.
func NewCaseExportOperation()(*CaseExportOperation) {
    m := &CaseExportOperation{
        CaseOperation: *NewCaseOperation(),
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateCaseExportOperationFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateCaseExportOperationFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewCaseExportOperation(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *CaseExportOperation) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetAzureBlobContainer gets the azureBlobContainer property value. The name of the Azure storage location where the export will be stored. This only applies to exports stored in your own Azure storage location.
func (m *CaseExportOperation) GetAzureBlobContainer()(*string) {
    if m == nil {
        return nil
    } else {
        return m.azureBlobContainer
    }
}
// GetAzureBlobToken gets the azureBlobToken property value. The SAS token for the Azure storage location.  This only applies to exports stored in your own Azure storage location.
func (m *CaseExportOperation) GetAzureBlobToken()(*string) {
    if m == nil {
        return nil
    } else {
        return m.azureBlobToken
    }
}
// GetDescription gets the description property value. The description provided for the export.
func (m *CaseExportOperation) GetDescription()(*string) {
    if m == nil {
        return nil
    } else {
        return m.description
    }
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
    return res
}
// GetOutputFolderId gets the outputFolderId property value. The outputFolderId property
func (m *CaseExportOperation) GetOutputFolderId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.outputFolderId
    }
}
// GetOutputName gets the outputName property value. The name provided for the export.
func (m *CaseExportOperation) GetOutputName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.outputName
    }
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
        err = writer.WriteAdditionalData(m.GetAdditionalData())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAdditionalData sets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *CaseExportOperation) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
// SetAzureBlobContainer sets the azureBlobContainer property value. The name of the Azure storage location where the export will be stored. This only applies to exports stored in your own Azure storage location.
func (m *CaseExportOperation) SetAzureBlobContainer(value *string)() {
    if m != nil {
        m.azureBlobContainer = value
    }
}
// SetAzureBlobToken sets the azureBlobToken property value. The SAS token for the Azure storage location.  This only applies to exports stored in your own Azure storage location.
func (m *CaseExportOperation) SetAzureBlobToken(value *string)() {
    if m != nil {
        m.azureBlobToken = value
    }
}
// SetDescription sets the description property value. The description provided for the export.
func (m *CaseExportOperation) SetDescription(value *string)() {
    if m != nil {
        m.description = value
    }
}
// SetOutputFolderId sets the outputFolderId property value. The outputFolderId property
func (m *CaseExportOperation) SetOutputFolderId(value *string)() {
    if m != nil {
        m.outputFolderId = value
    }
}
// SetOutputName sets the outputName property value. The name provided for the export.
func (m *CaseExportOperation) SetOutputName(value *string)() {
    if m != nil {
        m.outputName = value
    }
}
