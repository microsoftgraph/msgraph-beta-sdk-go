package security

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// EdiscoverySearchExportOperation 
type EdiscoverySearchExportOperation struct {
    CaseOperation
}
// NewEdiscoverySearchExportOperation instantiates a new ediscoverySearchExportOperation and sets the default values.
func NewEdiscoverySearchExportOperation()(*EdiscoverySearchExportOperation) {
    m := &EdiscoverySearchExportOperation{
        CaseOperation: *NewCaseOperation(),
    }
    return m
}
// CreateEdiscoverySearchExportOperationFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateEdiscoverySearchExportOperationFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewEdiscoverySearchExportOperation(), nil
}
// GetAdditionalOptions gets the additionalOptions property value. The additionalOptions property
func (m *EdiscoverySearchExportOperation) GetAdditionalOptions()(*EdiscoverySearchExportOperation_additionalOptions) {
    val, err := m.GetBackingStore().Get("additionalOptions")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*EdiscoverySearchExportOperation_additionalOptions)
    }
    return nil
}
// GetDescription gets the description property value. The description property
func (m *EdiscoverySearchExportOperation) GetDescription()(*string) {
    val, err := m.GetBackingStore().Get("description")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetDisplayName gets the displayName property value. The displayName property
func (m *EdiscoverySearchExportOperation) GetDisplayName()(*string) {
    val, err := m.GetBackingStore().Get("displayName")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetExportCriteria gets the exportCriteria property value. The exportCriteria property
func (m *EdiscoverySearchExportOperation) GetExportCriteria()(*EdiscoverySearchExportOperation_exportCriteria) {
    val, err := m.GetBackingStore().Get("exportCriteria")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*EdiscoverySearchExportOperation_exportCriteria)
    }
    return nil
}
// GetExportFileMetadata gets the exportFileMetadata property value. The exportFileMetadata property
func (m *EdiscoverySearchExportOperation) GetExportFileMetadata()([]ExportFileMetadataable) {
    val, err := m.GetBackingStore().Get("exportFileMetadata")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]ExportFileMetadataable)
    }
    return nil
}
// GetExportFormat gets the exportFormat property value. The exportFormat property
func (m *EdiscoverySearchExportOperation) GetExportFormat()(*EdiscoverySearchExportOperation_exportFormat) {
    val, err := m.GetBackingStore().Get("exportFormat")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*EdiscoverySearchExportOperation_exportFormat)
    }
    return nil
}
// GetExportLocation gets the exportLocation property value. The exportLocation property
func (m *EdiscoverySearchExportOperation) GetExportLocation()(*EdiscoverySearchExportOperation_exportLocation) {
    val, err := m.GetBackingStore().Get("exportLocation")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*EdiscoverySearchExportOperation_exportLocation)
    }
    return nil
}
// GetExportSingleItems gets the exportSingleItems property value. The exportSingleItems property
func (m *EdiscoverySearchExportOperation) GetExportSingleItems()(*bool) {
    val, err := m.GetBackingStore().Get("exportSingleItems")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *EdiscoverySearchExportOperation) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.CaseOperation.GetFieldDeserializers()
    res["additionalOptions"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseEdiscoverySearchExportOperation_additionalOptions)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAdditionalOptions(val.(*EdiscoverySearchExportOperation_additionalOptions))
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
    res["displayName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDisplayName(val)
        }
        return nil
    }
    res["exportCriteria"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseEdiscoverySearchExportOperation_exportCriteria)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetExportCriteria(val.(*EdiscoverySearchExportOperation_exportCriteria))
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
    res["exportFormat"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseEdiscoverySearchExportOperation_exportFormat)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetExportFormat(val.(*EdiscoverySearchExportOperation_exportFormat))
        }
        return nil
    }
    res["exportLocation"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseEdiscoverySearchExportOperation_exportLocation)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetExportLocation(val.(*EdiscoverySearchExportOperation_exportLocation))
        }
        return nil
    }
    res["exportSingleItems"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetExportSingleItems(val)
        }
        return nil
    }
    res["search"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateEdiscoverySearchFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSearch(val.(EdiscoverySearchable))
        }
        return nil
    }
    return res
}
// GetSearch gets the search property value. The search property
func (m *EdiscoverySearchExportOperation) GetSearch()(EdiscoverySearchable) {
    val, err := m.GetBackingStore().Get("search")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(EdiscoverySearchable)
    }
    return nil
}
// Serialize serializes information the current object
func (m *EdiscoverySearchExportOperation) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.CaseOperation.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetAdditionalOptions() != nil {
        cast := (*m.GetAdditionalOptions()).String()
        err = writer.WriteStringValue("additionalOptions", &cast)
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
        err = writer.WriteStringValue("displayName", m.GetDisplayName())
        if err != nil {
            return err
        }
    }
    if m.GetExportCriteria() != nil {
        cast := (*m.GetExportCriteria()).String()
        err = writer.WriteStringValue("exportCriteria", &cast)
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
    if m.GetExportFormat() != nil {
        cast := (*m.GetExportFormat()).String()
        err = writer.WriteStringValue("exportFormat", &cast)
        if err != nil {
            return err
        }
    }
    if m.GetExportLocation() != nil {
        cast := (*m.GetExportLocation()).String()
        err = writer.WriteStringValue("exportLocation", &cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("exportSingleItems", m.GetExportSingleItems())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("search", m.GetSearch())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAdditionalOptions sets the additionalOptions property value. The additionalOptions property
func (m *EdiscoverySearchExportOperation) SetAdditionalOptions(value *EdiscoverySearchExportOperation_additionalOptions)() {
    err := m.GetBackingStore().Set("additionalOptions", value)
    if err != nil {
        panic(err)
    }
}
// SetDescription sets the description property value. The description property
func (m *EdiscoverySearchExportOperation) SetDescription(value *string)() {
    err := m.GetBackingStore().Set("description", value)
    if err != nil {
        panic(err)
    }
}
// SetDisplayName sets the displayName property value. The displayName property
func (m *EdiscoverySearchExportOperation) SetDisplayName(value *string)() {
    err := m.GetBackingStore().Set("displayName", value)
    if err != nil {
        panic(err)
    }
}
// SetExportCriteria sets the exportCriteria property value. The exportCriteria property
func (m *EdiscoverySearchExportOperation) SetExportCriteria(value *EdiscoverySearchExportOperation_exportCriteria)() {
    err := m.GetBackingStore().Set("exportCriteria", value)
    if err != nil {
        panic(err)
    }
}
// SetExportFileMetadata sets the exportFileMetadata property value. The exportFileMetadata property
func (m *EdiscoverySearchExportOperation) SetExportFileMetadata(value []ExportFileMetadataable)() {
    err := m.GetBackingStore().Set("exportFileMetadata", value)
    if err != nil {
        panic(err)
    }
}
// SetExportFormat sets the exportFormat property value. The exportFormat property
func (m *EdiscoverySearchExportOperation) SetExportFormat(value *EdiscoverySearchExportOperation_exportFormat)() {
    err := m.GetBackingStore().Set("exportFormat", value)
    if err != nil {
        panic(err)
    }
}
// SetExportLocation sets the exportLocation property value. The exportLocation property
func (m *EdiscoverySearchExportOperation) SetExportLocation(value *EdiscoverySearchExportOperation_exportLocation)() {
    err := m.GetBackingStore().Set("exportLocation", value)
    if err != nil {
        panic(err)
    }
}
// SetExportSingleItems sets the exportSingleItems property value. The exportSingleItems property
func (m *EdiscoverySearchExportOperation) SetExportSingleItems(value *bool)() {
    err := m.GetBackingStore().Set("exportSingleItems", value)
    if err != nil {
        panic(err)
    }
}
// SetSearch sets the search property value. The search property
func (m *EdiscoverySearchExportOperation) SetSearch(value EdiscoverySearchable)() {
    err := m.GetBackingStore().Set("search", value)
    if err != nil {
        panic(err)
    }
}
// EdiscoverySearchExportOperationable 
type EdiscoverySearchExportOperationable interface {
    CaseOperationable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAdditionalOptions()(*EdiscoverySearchExportOperation_additionalOptions)
    GetDescription()(*string)
    GetDisplayName()(*string)
    GetExportCriteria()(*EdiscoverySearchExportOperation_exportCriteria)
    GetExportFileMetadata()([]ExportFileMetadataable)
    GetExportFormat()(*EdiscoverySearchExportOperation_exportFormat)
    GetExportLocation()(*EdiscoverySearchExportOperation_exportLocation)
    GetExportSingleItems()(*bool)
    GetSearch()(EdiscoverySearchable)
    SetAdditionalOptions(value *EdiscoverySearchExportOperation_additionalOptions)()
    SetDescription(value *string)()
    SetDisplayName(value *string)()
    SetExportCriteria(value *EdiscoverySearchExportOperation_exportCriteria)()
    SetExportFileMetadata(value []ExportFileMetadataable)()
    SetExportFormat(value *EdiscoverySearchExportOperation_exportFormat)()
    SetExportLocation(value *EdiscoverySearchExportOperation_exportLocation)()
    SetExportSingleItems(value *bool)()
    SetSearch(value EdiscoverySearchable)()
}
