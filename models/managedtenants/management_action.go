package managedtenants

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
)

// ManagementAction provides operations to manage the collection of accessReviewDecision entities.
type ManagementAction struct {
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Entity
    // The category property
    category *ManagementCategory
    // The description for the management action. Optional. Read-only.
    description *string
    // The display name for the management action. Optional. Read-only.
    displayName *string
    // The reference for the management template used to generate the management action. Required. Read-only.
    referenceTemplateId *string
    // The referenceTemplateVersion property
    referenceTemplateVersion *int32
    // The collection of workload actions associated with the management action. Required. Read-only.
    workloadActions []WorkloadActionable
}
// NewManagementAction instantiates a new managementAction and sets the default values.
func NewManagementAction()(*ManagementAction) {
    m := &ManagementAction{
        Entity: *ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.NewEntity(),
    }
    return m
}
// CreateManagementActionFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateManagementActionFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewManagementAction(), nil
}
// GetCategory gets the category property value. The category property
func (m *ManagementAction) GetCategory()(*ManagementCategory) {
    return m.category
}
// GetDescription gets the description property value. The description for the management action. Optional. Read-only.
func (m *ManagementAction) GetDescription()(*string) {
    return m.description
}
// GetDisplayName gets the displayName property value. The display name for the management action. Optional. Read-only.
func (m *ManagementAction) GetDisplayName()(*string) {
    return m.displayName
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ManagementAction) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["category"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetEnumValue(ParseManagementCategory , m.SetCategory)
    res["description"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetDescription)
    res["displayName"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetDisplayName)
    res["referenceTemplateId"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetReferenceTemplateId)
    res["referenceTemplateVersion"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetInt32Value(m.SetReferenceTemplateVersion)
    res["workloadActions"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfObjectValues(CreateWorkloadActionFromDiscriminatorValue , m.SetWorkloadActions)
    return res
}
// GetReferenceTemplateId gets the referenceTemplateId property value. The reference for the management template used to generate the management action. Required. Read-only.
func (m *ManagementAction) GetReferenceTemplateId()(*string) {
    return m.referenceTemplateId
}
// GetReferenceTemplateVersion gets the referenceTemplateVersion property value. The referenceTemplateVersion property
func (m *ManagementAction) GetReferenceTemplateVersion()(*int32) {
    return m.referenceTemplateVersion
}
// GetWorkloadActions gets the workloadActions property value. The collection of workload actions associated with the management action. Required. Read-only.
func (m *ManagementAction) GetWorkloadActions()([]WorkloadActionable) {
    return m.workloadActions
}
// Serialize serializes information the current object
func (m *ManagementAction) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetCategory() != nil {
        cast := (*m.GetCategory()).String()
        err = writer.WriteStringValue("category", &cast)
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
    {
        err = writer.WriteStringValue("referenceTemplateId", m.GetReferenceTemplateId())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("referenceTemplateVersion", m.GetReferenceTemplateVersion())
        if err != nil {
            return err
        }
    }
    if m.GetWorkloadActions() != nil {
        cast := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.CollectionCast[i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable](m.GetWorkloadActions())
        err = writer.WriteCollectionOfObjectValues("workloadActions", cast)
        if err != nil {
            return err
        }
    }
    return nil
}
// SetCategory sets the category property value. The category property
func (m *ManagementAction) SetCategory(value *ManagementCategory)() {
    m.category = value
}
// SetDescription sets the description property value. The description for the management action. Optional. Read-only.
func (m *ManagementAction) SetDescription(value *string)() {
    m.description = value
}
// SetDisplayName sets the displayName property value. The display name for the management action. Optional. Read-only.
func (m *ManagementAction) SetDisplayName(value *string)() {
    m.displayName = value
}
// SetReferenceTemplateId sets the referenceTemplateId property value. The reference for the management template used to generate the management action. Required. Read-only.
func (m *ManagementAction) SetReferenceTemplateId(value *string)() {
    m.referenceTemplateId = value
}
// SetReferenceTemplateVersion sets the referenceTemplateVersion property value. The referenceTemplateVersion property
func (m *ManagementAction) SetReferenceTemplateVersion(value *int32)() {
    m.referenceTemplateVersion = value
}
// SetWorkloadActions sets the workloadActions property value. The collection of workload actions associated with the management action. Required. Read-only.
func (m *ManagementAction) SetWorkloadActions(value []WorkloadActionable)() {
    m.workloadActions = value
}
