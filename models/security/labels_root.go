package security

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
)

// LabelsRoot 
type LabelsRoot struct {
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Entity
}
// NewLabelsRoot instantiates a new labelsRoot and sets the default values.
func NewLabelsRoot()(*LabelsRoot) {
    m := &LabelsRoot{
        Entity: *ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.NewEntity(),
    }
    return m
}
// CreateLabelsRootFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateLabelsRootFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewLabelsRoot(), nil
}
// GetAuthorities gets the authorities property value. Specifies the underlying authority that describes the type of content to be retained and its retention schedule.
func (m *LabelsRoot) GetAuthorities()([]AuthorityTemplateable) {
    val, err := m.GetBackingStore().Get("authorities")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]AuthorityTemplateable)
    }
    return nil
}
// GetCategories gets the categories property value. Specifies a group of similar types of content in a particular department.
func (m *LabelsRoot) GetCategories()([]CategoryTemplateable) {
    val, err := m.GetBackingStore().Get("categories")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]CategoryTemplateable)
    }
    return nil
}
// GetCitations gets the citations property value. The specific rule or regulation created by a jurisdiction used to determine whether certain labels and content should be retained or deleted.
func (m *LabelsRoot) GetCitations()([]CitationTemplateable) {
    val, err := m.GetBackingStore().Get("citations")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]CitationTemplateable)
    }
    return nil
}
// GetDepartments gets the departments property value. Specifies the department or business unit of an organization to which a label belongs.
func (m *LabelsRoot) GetDepartments()([]DepartmentTemplateable) {
    val, err := m.GetBackingStore().Get("departments")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]DepartmentTemplateable)
    }
    return nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *LabelsRoot) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["authorities"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateAuthorityTemplateFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]AuthorityTemplateable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(AuthorityTemplateable)
                }
            }
            m.SetAuthorities(res)
        }
        return nil
    }
    res["categories"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateCategoryTemplateFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]CategoryTemplateable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(CategoryTemplateable)
                }
            }
            m.SetCategories(res)
        }
        return nil
    }
    res["citations"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateCitationTemplateFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]CitationTemplateable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(CitationTemplateable)
                }
            }
            m.SetCitations(res)
        }
        return nil
    }
    res["departments"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateDepartmentTemplateFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]DepartmentTemplateable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(DepartmentTemplateable)
                }
            }
            m.SetDepartments(res)
        }
        return nil
    }
    res["filePlanReferences"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateFilePlanReferenceTemplateFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]FilePlanReferenceTemplateable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(FilePlanReferenceTemplateable)
                }
            }
            m.SetFilePlanReferences(res)
        }
        return nil
    }
    res["@odata.type"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOdataType(val)
        }
        return nil
    }
    res["retentionLabels"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateRetentionLabelFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]RetentionLabelable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(RetentionLabelable)
                }
            }
            m.SetRetentionLabels(res)
        }
        return nil
    }
    return res
}
// GetFilePlanReferences gets the filePlanReferences property value. Specifies a unique alpha-numeric identifier for an organization’s retention schedule.
func (m *LabelsRoot) GetFilePlanReferences()([]FilePlanReferenceTemplateable) {
    val, err := m.GetBackingStore().Get("filePlanReferences")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]FilePlanReferenceTemplateable)
    }
    return nil
}
// GetOdataType gets the @odata.type property value. The OdataType property
func (m *LabelsRoot) GetOdataType()(*string) {
    val, err := m.GetBackingStore().Get("odataType")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetRetentionLabels gets the retentionLabels property value. Represents how customers can manage their data, whether and for how long to retain or delete it.
func (m *LabelsRoot) GetRetentionLabels()([]RetentionLabelable) {
    val, err := m.GetBackingStore().Get("retentionLabels")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]RetentionLabelable)
    }
    return nil
}
// Serialize serializes information the current object
func (m *LabelsRoot) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetAuthorities() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetAuthorities()))
        for i, v := range m.GetAuthorities() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err = writer.WriteCollectionOfObjectValues("authorities", cast)
        if err != nil {
            return err
        }
    }
    if m.GetCategories() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetCategories()))
        for i, v := range m.GetCategories() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err = writer.WriteCollectionOfObjectValues("categories", cast)
        if err != nil {
            return err
        }
    }
    if m.GetCitations() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetCitations()))
        for i, v := range m.GetCitations() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err = writer.WriteCollectionOfObjectValues("citations", cast)
        if err != nil {
            return err
        }
    }
    if m.GetDepartments() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetDepartments()))
        for i, v := range m.GetDepartments() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err = writer.WriteCollectionOfObjectValues("departments", cast)
        if err != nil {
            return err
        }
    }
    if m.GetFilePlanReferences() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetFilePlanReferences()))
        for i, v := range m.GetFilePlanReferences() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err = writer.WriteCollectionOfObjectValues("filePlanReferences", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("@odata.type", m.GetOdataType())
        if err != nil {
            return err
        }
    }
    if m.GetRetentionLabels() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetRetentionLabels()))
        for i, v := range m.GetRetentionLabels() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err = writer.WriteCollectionOfObjectValues("retentionLabels", cast)
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAuthorities sets the authorities property value. Specifies the underlying authority that describes the type of content to be retained and its retention schedule.
func (m *LabelsRoot) SetAuthorities(value []AuthorityTemplateable)() {
    err := m.GetBackingStore().Set("authorities", value)
    if err != nil {
        panic(err)
    }
}
// SetCategories sets the categories property value. Specifies a group of similar types of content in a particular department.
func (m *LabelsRoot) SetCategories(value []CategoryTemplateable)() {
    err := m.GetBackingStore().Set("categories", value)
    if err != nil {
        panic(err)
    }
}
// SetCitations sets the citations property value. The specific rule or regulation created by a jurisdiction used to determine whether certain labels and content should be retained or deleted.
func (m *LabelsRoot) SetCitations(value []CitationTemplateable)() {
    err := m.GetBackingStore().Set("citations", value)
    if err != nil {
        panic(err)
    }
}
// SetDepartments sets the departments property value. Specifies the department or business unit of an organization to which a label belongs.
func (m *LabelsRoot) SetDepartments(value []DepartmentTemplateable)() {
    err := m.GetBackingStore().Set("departments", value)
    if err != nil {
        panic(err)
    }
}
// SetFilePlanReferences sets the filePlanReferences property value. Specifies a unique alpha-numeric identifier for an organization’s retention schedule.
func (m *LabelsRoot) SetFilePlanReferences(value []FilePlanReferenceTemplateable)() {
    err := m.GetBackingStore().Set("filePlanReferences", value)
    if err != nil {
        panic(err)
    }
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *LabelsRoot) SetOdataType(value *string)() {
    err := m.GetBackingStore().Set("odataType", value)
    if err != nil {
        panic(err)
    }
}
// SetRetentionLabels sets the retentionLabels property value. Represents how customers can manage their data, whether and for how long to retain or delete it.
func (m *LabelsRoot) SetRetentionLabels(value []RetentionLabelable)() {
    err := m.GetBackingStore().Set("retentionLabels", value)
    if err != nil {
        panic(err)
    }
}
// LabelsRootable 
type LabelsRootable interface {
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Entityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAuthorities()([]AuthorityTemplateable)
    GetCategories()([]CategoryTemplateable)
    GetCitations()([]CitationTemplateable)
    GetDepartments()([]DepartmentTemplateable)
    GetFilePlanReferences()([]FilePlanReferenceTemplateable)
    GetOdataType()(*string)
    GetRetentionLabels()([]RetentionLabelable)
    SetAuthorities(value []AuthorityTemplateable)()
    SetCategories(value []CategoryTemplateable)()
    SetCitations(value []CitationTemplateable)()
    SetDepartments(value []DepartmentTemplateable)()
    SetFilePlanReferences(value []FilePlanReferenceTemplateable)()
    SetOdataType(value *string)()
    SetRetentionLabels(value []RetentionLabelable)()
}
