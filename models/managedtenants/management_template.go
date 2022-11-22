package managedtenants

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
)

// ManagementTemplate 
type ManagementTemplate struct {
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Entity
    // The management category for the management template. Possible values are: custom, devices, identity, unknownFutureValue. Required. Read-only.
    category *ManagementCategory
    // The createdByUserId property
    createdByUserId *string
    // The createdDateTime property
    createdDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // The description for the management template. Optional. Read-only.
    description *string
    // The display name for the management template. Required. Read-only.
    displayName *string
    // The informationLinks property
    informationLinks []ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ActionUrlable
    // The lastActionByUserId property
    lastActionByUserId *string
    // The lastActionDateTime property
    lastActionDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // The managementTemplateCollections property
    managementTemplateCollections []ManagementTemplateCollectionable
    // The managementTemplateSteps property
    managementTemplateSteps []ManagementTemplateStepable
    // The collection of parameters used by the management template. Optional. Read-only.
    parameters []TemplateParameterable
    // The priority property
    priority *int32
    // The provider property
    provider *ManagementProvider
    // The userImpact property
    userImpact *string
    // The version property
    version *int32
    // The collection of workload actions associated with the management template. Optional. Read-only.
    workloadActions []WorkloadActionable
}
// NewManagementTemplate instantiates a new managementTemplate and sets the default values.
func NewManagementTemplate()(*ManagementTemplate) {
    m := &ManagementTemplate{
        Entity: *ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.NewEntity(),
    }
    return m
}
// CreateManagementTemplateFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateManagementTemplateFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewManagementTemplate(), nil
}
// GetCategory gets the category property value. The management category for the management template. Possible values are: custom, devices, identity, unknownFutureValue. Required. Read-only.
func (m *ManagementTemplate) GetCategory()(*ManagementCategory) {
    return m.category
}
// GetCreatedByUserId gets the createdByUserId property value. The createdByUserId property
func (m *ManagementTemplate) GetCreatedByUserId()(*string) {
    return m.createdByUserId
}
// GetCreatedDateTime gets the createdDateTime property value. The createdDateTime property
func (m *ManagementTemplate) GetCreatedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.createdDateTime
}
// GetDescription gets the description property value. The description for the management template. Optional. Read-only.
func (m *ManagementTemplate) GetDescription()(*string) {
    return m.description
}
// GetDisplayName gets the displayName property value. The display name for the management template. Required. Read-only.
func (m *ManagementTemplate) GetDisplayName()(*string) {
    return m.displayName
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ManagementTemplate) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["category"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetEnumValue(ParseManagementCategory , m.SetCategory)
    res["createdByUserId"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetCreatedByUserId)
    res["createdDateTime"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetTimeValue(m.SetCreatedDateTime)
    res["description"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetDescription)
    res["displayName"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetDisplayName)
    res["informationLinks"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfObjectValues(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateActionUrlFromDiscriminatorValue , m.SetInformationLinks)
    res["lastActionByUserId"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetLastActionByUserId)
    res["lastActionDateTime"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetTimeValue(m.SetLastActionDateTime)
    res["managementTemplateCollections"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfObjectValues(CreateManagementTemplateCollectionFromDiscriminatorValue , m.SetManagementTemplateCollections)
    res["managementTemplateSteps"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfObjectValues(CreateManagementTemplateStepFromDiscriminatorValue , m.SetManagementTemplateSteps)
    res["parameters"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfObjectValues(CreateTemplateParameterFromDiscriminatorValue , m.SetParameters)
    res["priority"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetInt32Value(m.SetPriority)
    res["provider"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetEnumValue(ParseManagementProvider , m.SetProvider)
    res["userImpact"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetUserImpact)
    res["version"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetInt32Value(m.SetVersion)
    res["workloadActions"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfObjectValues(CreateWorkloadActionFromDiscriminatorValue , m.SetWorkloadActions)
    return res
}
// GetInformationLinks gets the informationLinks property value. The informationLinks property
func (m *ManagementTemplate) GetInformationLinks()([]ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ActionUrlable) {
    return m.informationLinks
}
// GetLastActionByUserId gets the lastActionByUserId property value. The lastActionByUserId property
func (m *ManagementTemplate) GetLastActionByUserId()(*string) {
    return m.lastActionByUserId
}
// GetLastActionDateTime gets the lastActionDateTime property value. The lastActionDateTime property
func (m *ManagementTemplate) GetLastActionDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.lastActionDateTime
}
// GetManagementTemplateCollections gets the managementTemplateCollections property value. The managementTemplateCollections property
func (m *ManagementTemplate) GetManagementTemplateCollections()([]ManagementTemplateCollectionable) {
    return m.managementTemplateCollections
}
// GetManagementTemplateSteps gets the managementTemplateSteps property value. The managementTemplateSteps property
func (m *ManagementTemplate) GetManagementTemplateSteps()([]ManagementTemplateStepable) {
    return m.managementTemplateSteps
}
// GetParameters gets the parameters property value. The collection of parameters used by the management template. Optional. Read-only.
func (m *ManagementTemplate) GetParameters()([]TemplateParameterable) {
    return m.parameters
}
// GetPriority gets the priority property value. The priority property
func (m *ManagementTemplate) GetPriority()(*int32) {
    return m.priority
}
// GetProvider gets the provider property value. The provider property
func (m *ManagementTemplate) GetProvider()(*ManagementProvider) {
    return m.provider
}
// GetUserImpact gets the userImpact property value. The userImpact property
func (m *ManagementTemplate) GetUserImpact()(*string) {
    return m.userImpact
}
// GetVersion gets the version property value. The version property
func (m *ManagementTemplate) GetVersion()(*int32) {
    return m.version
}
// GetWorkloadActions gets the workloadActions property value. The collection of workload actions associated with the management template. Optional. Read-only.
func (m *ManagementTemplate) GetWorkloadActions()([]WorkloadActionable) {
    return m.workloadActions
}
// Serialize serializes information the current object
func (m *ManagementTemplate) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
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
        err = writer.WriteStringValue("createdByUserId", m.GetCreatedByUserId())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteTimeValue("createdDateTime", m.GetCreatedDateTime())
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
    if m.GetInformationLinks() != nil {
        cast := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.CollectionCast[i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable](m.GetInformationLinks())
        err = writer.WriteCollectionOfObjectValues("informationLinks", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("lastActionByUserId", m.GetLastActionByUserId())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteTimeValue("lastActionDateTime", m.GetLastActionDateTime())
        if err != nil {
            return err
        }
    }
    if m.GetManagementTemplateCollections() != nil {
        cast := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.CollectionCast[i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable](m.GetManagementTemplateCollections())
        err = writer.WriteCollectionOfObjectValues("managementTemplateCollections", cast)
        if err != nil {
            return err
        }
    }
    if m.GetManagementTemplateSteps() != nil {
        cast := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.CollectionCast[i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable](m.GetManagementTemplateSteps())
        err = writer.WriteCollectionOfObjectValues("managementTemplateSteps", cast)
        if err != nil {
            return err
        }
    }
    if m.GetParameters() != nil {
        cast := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.CollectionCast[i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable](m.GetParameters())
        err = writer.WriteCollectionOfObjectValues("parameters", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("priority", m.GetPriority())
        if err != nil {
            return err
        }
    }
    if m.GetProvider() != nil {
        cast := (*m.GetProvider()).String()
        err = writer.WriteStringValue("provider", &cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("userImpact", m.GetUserImpact())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("version", m.GetVersion())
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
// SetCategory sets the category property value. The management category for the management template. Possible values are: custom, devices, identity, unknownFutureValue. Required. Read-only.
func (m *ManagementTemplate) SetCategory(value *ManagementCategory)() {
    m.category = value
}
// SetCreatedByUserId sets the createdByUserId property value. The createdByUserId property
func (m *ManagementTemplate) SetCreatedByUserId(value *string)() {
    m.createdByUserId = value
}
// SetCreatedDateTime sets the createdDateTime property value. The createdDateTime property
func (m *ManagementTemplate) SetCreatedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.createdDateTime = value
}
// SetDescription sets the description property value. The description for the management template. Optional. Read-only.
func (m *ManagementTemplate) SetDescription(value *string)() {
    m.description = value
}
// SetDisplayName sets the displayName property value. The display name for the management template. Required. Read-only.
func (m *ManagementTemplate) SetDisplayName(value *string)() {
    m.displayName = value
}
// SetInformationLinks sets the informationLinks property value. The informationLinks property
func (m *ManagementTemplate) SetInformationLinks(value []ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ActionUrlable)() {
    m.informationLinks = value
}
// SetLastActionByUserId sets the lastActionByUserId property value. The lastActionByUserId property
func (m *ManagementTemplate) SetLastActionByUserId(value *string)() {
    m.lastActionByUserId = value
}
// SetLastActionDateTime sets the lastActionDateTime property value. The lastActionDateTime property
func (m *ManagementTemplate) SetLastActionDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.lastActionDateTime = value
}
// SetManagementTemplateCollections sets the managementTemplateCollections property value. The managementTemplateCollections property
func (m *ManagementTemplate) SetManagementTemplateCollections(value []ManagementTemplateCollectionable)() {
    m.managementTemplateCollections = value
}
// SetManagementTemplateSteps sets the managementTemplateSteps property value. The managementTemplateSteps property
func (m *ManagementTemplate) SetManagementTemplateSteps(value []ManagementTemplateStepable)() {
    m.managementTemplateSteps = value
}
// SetParameters sets the parameters property value. The collection of parameters used by the management template. Optional. Read-only.
func (m *ManagementTemplate) SetParameters(value []TemplateParameterable)() {
    m.parameters = value
}
// SetPriority sets the priority property value. The priority property
func (m *ManagementTemplate) SetPriority(value *int32)() {
    m.priority = value
}
// SetProvider sets the provider property value. The provider property
func (m *ManagementTemplate) SetProvider(value *ManagementProvider)() {
    m.provider = value
}
// SetUserImpact sets the userImpact property value. The userImpact property
func (m *ManagementTemplate) SetUserImpact(value *string)() {
    m.userImpact = value
}
// SetVersion sets the version property value. The version property
func (m *ManagementTemplate) SetVersion(value *int32)() {
    m.version = value
}
// SetWorkloadActions sets the workloadActions property value. The collection of workload actions associated with the management template. Optional. Read-only.
func (m *ManagementTemplate) SetWorkloadActions(value []WorkloadActionable)() {
    m.workloadActions = value
}
