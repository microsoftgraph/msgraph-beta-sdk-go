package compliance

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e "github.com/microsoft/kiota-abstractions-go/store"
    ic154d683aa4025ee28853b9c1a3c35cd1f093a1c4542feba4c07682e2752db13 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/ediscovery"
)

type EdiscoveryCasesItemReviewSetsItemMicrosoftGraphEdiscoveryExportExportPostRequestBody struct {
    // Stores model information.
    backingStore ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore
}
// NewEdiscoveryCasesItemReviewSetsItemMicrosoftGraphEdiscoveryExportExportPostRequestBody instantiates a new EdiscoveryCasesItemReviewSetsItemMicrosoftGraphEdiscoveryExportExportPostRequestBody and sets the default values.
func NewEdiscoveryCasesItemReviewSetsItemMicrosoftGraphEdiscoveryExportExportPostRequestBody()(*EdiscoveryCasesItemReviewSetsItemMicrosoftGraphEdiscoveryExportExportPostRequestBody) {
    m := &EdiscoveryCasesItemReviewSetsItemMicrosoftGraphEdiscoveryExportExportPostRequestBody{
    }
    m.backingStore = ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStoreFactoryInstance();
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateEdiscoveryCasesItemReviewSetsItemMicrosoftGraphEdiscoveryExportExportPostRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateEdiscoveryCasesItemReviewSetsItemMicrosoftGraphEdiscoveryExportExportPostRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewEdiscoveryCasesItemReviewSetsItemMicrosoftGraphEdiscoveryExportExportPostRequestBody(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *EdiscoveryCasesItemReviewSetsItemMicrosoftGraphEdiscoveryExportExportPostRequestBody) GetAdditionalData()(map[string]any) {
    val , err :=  m.backingStore.Get("additionalData")
    if err != nil {
        panic(err)
    }
    if val == nil {
        var value = make(map[string]any);
        m.SetAdditionalData(value);
    }
    return val.(map[string]any)
}
// GetAzureBlobContainer gets the azureBlobContainer property value. The azureBlobContainer property
// returns a *string when successful
func (m *EdiscoveryCasesItemReviewSetsItemMicrosoftGraphEdiscoveryExportExportPostRequestBody) GetAzureBlobContainer()(*string) {
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
// returns a *string when successful
func (m *EdiscoveryCasesItemReviewSetsItemMicrosoftGraphEdiscoveryExportExportPostRequestBody) GetAzureBlobToken()(*string) {
    val, err := m.GetBackingStore().Get("azureBlobToken")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetBackingStore gets the BackingStore property value. Stores model information.
// returns a BackingStore when successful
func (m *EdiscoveryCasesItemReviewSetsItemMicrosoftGraphEdiscoveryExportExportPostRequestBody) GetBackingStore()(ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore) {
    return m.backingStore
}
// GetDescription gets the description property value. The description property
// returns a *string when successful
func (m *EdiscoveryCasesItemReviewSetsItemMicrosoftGraphEdiscoveryExportExportPostRequestBody) GetDescription()(*string) {
    val, err := m.GetBackingStore().Get("description")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetExportOptions gets the exportOptions property value. The exportOptions property
// returns a *ExportOptions when successful
func (m *EdiscoveryCasesItemReviewSetsItemMicrosoftGraphEdiscoveryExportExportPostRequestBody) GetExportOptions()(*ic154d683aa4025ee28853b9c1a3c35cd1f093a1c4542feba4c07682e2752db13.ExportOptions) {
    val, err := m.GetBackingStore().Get("exportOptions")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*ic154d683aa4025ee28853b9c1a3c35cd1f093a1c4542feba4c07682e2752db13.ExportOptions)
    }
    return nil
}
// GetExportStructure gets the exportStructure property value. The exportStructure property
// returns a *ExportFileStructure when successful
func (m *EdiscoveryCasesItemReviewSetsItemMicrosoftGraphEdiscoveryExportExportPostRequestBody) GetExportStructure()(*ic154d683aa4025ee28853b9c1a3c35cd1f093a1c4542feba4c07682e2752db13.ExportFileStructure) {
    val, err := m.GetBackingStore().Get("exportStructure")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*ic154d683aa4025ee28853b9c1a3c35cd1f093a1c4542feba4c07682e2752db13.ExportFileStructure)
    }
    return nil
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
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
// returns a *string when successful
func (m *EdiscoveryCasesItemReviewSetsItemMicrosoftGraphEdiscoveryExportExportPostRequestBody) GetOutputName()(*string) {
    val, err := m.GetBackingStore().Get("outputName")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
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
// SetAdditionalData sets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *EdiscoveryCasesItemReviewSetsItemMicrosoftGraphEdiscoveryExportExportPostRequestBody) SetAdditionalData(value map[string]any)() {
    err := m.GetBackingStore().Set("additionalData", value)
    if err != nil {
        panic(err)
    }
}
// SetAzureBlobContainer sets the azureBlobContainer property value. The azureBlobContainer property
func (m *EdiscoveryCasesItemReviewSetsItemMicrosoftGraphEdiscoveryExportExportPostRequestBody) SetAzureBlobContainer(value *string)() {
    err := m.GetBackingStore().Set("azureBlobContainer", value)
    if err != nil {
        panic(err)
    }
}
// SetAzureBlobToken sets the azureBlobToken property value. The azureBlobToken property
func (m *EdiscoveryCasesItemReviewSetsItemMicrosoftGraphEdiscoveryExportExportPostRequestBody) SetAzureBlobToken(value *string)() {
    err := m.GetBackingStore().Set("azureBlobToken", value)
    if err != nil {
        panic(err)
    }
}
// SetBackingStore sets the BackingStore property value. Stores model information.
func (m *EdiscoveryCasesItemReviewSetsItemMicrosoftGraphEdiscoveryExportExportPostRequestBody) SetBackingStore(value ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)() {
    m.backingStore = value
}
// SetDescription sets the description property value. The description property
func (m *EdiscoveryCasesItemReviewSetsItemMicrosoftGraphEdiscoveryExportExportPostRequestBody) SetDescription(value *string)() {
    err := m.GetBackingStore().Set("description", value)
    if err != nil {
        panic(err)
    }
}
// SetExportOptions sets the exportOptions property value. The exportOptions property
func (m *EdiscoveryCasesItemReviewSetsItemMicrosoftGraphEdiscoveryExportExportPostRequestBody) SetExportOptions(value *ic154d683aa4025ee28853b9c1a3c35cd1f093a1c4542feba4c07682e2752db13.ExportOptions)() {
    err := m.GetBackingStore().Set("exportOptions", value)
    if err != nil {
        panic(err)
    }
}
// SetExportStructure sets the exportStructure property value. The exportStructure property
func (m *EdiscoveryCasesItemReviewSetsItemMicrosoftGraphEdiscoveryExportExportPostRequestBody) SetExportStructure(value *ic154d683aa4025ee28853b9c1a3c35cd1f093a1c4542feba4c07682e2752db13.ExportFileStructure)() {
    err := m.GetBackingStore().Set("exportStructure", value)
    if err != nil {
        panic(err)
    }
}
// SetOutputName sets the outputName property value. The outputName property
func (m *EdiscoveryCasesItemReviewSetsItemMicrosoftGraphEdiscoveryExportExportPostRequestBody) SetOutputName(value *string)() {
    err := m.GetBackingStore().Set("outputName", value)
    if err != nil {
        panic(err)
    }
}
type EdiscoveryCasesItemReviewSetsItemMicrosoftGraphEdiscoveryExportExportPostRequestBodyable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackedModel
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAzureBlobContainer()(*string)
    GetAzureBlobToken()(*string)
    GetBackingStore()(ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)
    GetDescription()(*string)
    GetExportOptions()(*ic154d683aa4025ee28853b9c1a3c35cd1f093a1c4542feba4c07682e2752db13.ExportOptions)
    GetExportStructure()(*ic154d683aa4025ee28853b9c1a3c35cd1f093a1c4542feba4c07682e2752db13.ExportFileStructure)
    GetOutputName()(*string)
    SetAzureBlobContainer(value *string)()
    SetAzureBlobToken(value *string)()
    SetBackingStore(value ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)()
    SetDescription(value *string)()
    SetExportOptions(value *ic154d683aa4025ee28853b9c1a3c35cd1f093a1c4542feba4c07682e2752db13.ExportOptions)()
    SetExportStructure(value *ic154d683aa4025ee28853b9c1a3c35cd1f093a1c4542feba4c07682e2752db13.ExportFileStructure)()
    SetOutputName(value *string)()
}
