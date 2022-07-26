package managedtenants

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
)

// ManagementTemplate provides operations to manage the collection of activityStatistics entities.
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
    odataTypeValue := "#microsoft.graph.managedTenants.managementTemplate";
    m.SetOdataType(&odataTypeValue);
    return m
}
// CreateManagementTemplateFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateManagementTemplateFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewManagementTemplate(), nil
}
// GetCategory gets the category property value. The management category for the management template. Possible values are: custom, devices, identity, unknownFutureValue. Required. Read-only.
func (m *ManagementTemplate) GetCategory()(*ManagementCategory) {
    if m == nil {
        return nil
    } else {
        return m.category
    }
}
// GetCreatedByUserId gets the createdByUserId property value. The createdByUserId property
func (m *ManagementTemplate) GetCreatedByUserId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.createdByUserId
    }
}
// GetCreatedDateTime gets the createdDateTime property value. The createdDateTime property
func (m *ManagementTemplate) GetCreatedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.createdDateTime
    }
}
// GetDescription gets the description property value. The description for the management template. Optional. Read-only.
func (m *ManagementTemplate) GetDescription()(*string) {
    if m == nil {
        return nil
    } else {
        return m.description
    }
}
// GetDisplayName gets the displayName property value. The display name for the management template. Required. Read-only.
func (m *ManagementTemplate) GetDisplayName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.displayName
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ManagementTemplate) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["category"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseManagementCategory)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCategory(val.(*ManagementCategory))
        }
        return nil
    }
    res["createdByUserId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCreatedByUserId(val)
        }
        return nil
    }
    res["createdDateTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCreatedDateTime(val)
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
    res["informationLinks"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateActionUrlFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ActionUrlable, len(val))
            for i, v := range val {
                res[i] = v.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ActionUrlable)
            }
            m.SetInformationLinks(res)
        }
        return nil
    }
    res["lastActionByUserId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLastActionByUserId(val)
        }
        return nil
    }
    res["lastActionDateTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLastActionDateTime(val)
        }
        return nil
    }
    res["managementTemplateCollections"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateManagementTemplateCollectionFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]ManagementTemplateCollectionable, len(val))
            for i, v := range val {
                res[i] = v.(ManagementTemplateCollectionable)
            }
            m.SetManagementTemplateCollections(res)
        }
        return nil
    }
    res["managementTemplateSteps"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateManagementTemplateStepFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]ManagementTemplateStepable, len(val))
            for i, v := range val {
                res[i] = v.(ManagementTemplateStepable)
            }
            m.SetManagementTemplateSteps(res)
        }
        return nil
    }
    res["parameters"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateTemplateParameterFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]TemplateParameterable, len(val))
            for i, v := range val {
                res[i] = v.(TemplateParameterable)
            }
            m.SetParameters(res)
        }
        return nil
    }
    res["priority"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPriority(val)
        }
        return nil
    }
    res["provider"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseManagementProvider)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetProvider(val.(*ManagementProvider))
        }
        return nil
    }
    res["userImpact"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUserImpact(val)
        }
        return nil
    }
    res["version"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetVersion(val)
        }
        return nil
    }
    res["workloadActions"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateWorkloadActionFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]WorkloadActionable, len(val))
            for i, v := range val {
                res[i] = v.(WorkloadActionable)
            }
            m.SetWorkloadActions(res)
        }
        return nil
    }
    return res
}
// GetInformationLinks gets the informationLinks property value. The informationLinks property
func (m *ManagementTemplate) GetInformationLinks()([]ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ActionUrlable) {
    if m == nil {
        return nil
    } else {
        return m.informationLinks
    }
}
// GetLastActionByUserId gets the lastActionByUserId property value. The lastActionByUserId property
func (m *ManagementTemplate) GetLastActionByUserId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.lastActionByUserId
    }
}
// GetLastActionDateTime gets the lastActionDateTime property value. The lastActionDateTime property
func (m *ManagementTemplate) GetLastActionDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.lastActionDateTime
    }
}
// GetManagementTemplateCollections gets the managementTemplateCollections property value. The managementTemplateCollections property
func (m *ManagementTemplate) GetManagementTemplateCollections()([]ManagementTemplateCollectionable) {
    if m == nil {
        return nil
    } else {
        return m.managementTemplateCollections
    }
}
// GetManagementTemplateSteps gets the managementTemplateSteps property value. The managementTemplateSteps property
func (m *ManagementTemplate) GetManagementTemplateSteps()([]ManagementTemplateStepable) {
    if m == nil {
        return nil
    } else {
        return m.managementTemplateSteps
    }
}
// GetParameters gets the parameters property value. The collection of parameters used by the management template. Optional. Read-only.
func (m *ManagementTemplate) GetParameters()([]TemplateParameterable) {
    if m == nil {
        return nil
    } else {
        return m.parameters
    }
}
// GetPriority gets the priority property value. The priority property
func (m *ManagementTemplate) GetPriority()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.priority
    }
}
// GetProvider gets the provider property value. The provider property
func (m *ManagementTemplate) GetProvider()(*ManagementProvider) {
    if m == nil {
        return nil
    } else {
        return m.provider
    }
}
// GetUserImpact gets the userImpact property value. The userImpact property
func (m *ManagementTemplate) GetUserImpact()(*string) {
    if m == nil {
        return nil
    } else {
        return m.userImpact
    }
}
// GetVersion gets the version property value. The version property
func (m *ManagementTemplate) GetVersion()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.version
    }
}
// GetWorkloadActions gets the workloadActions property value. The collection of workload actions associated with the management template. Optional. Read-only.
func (m *ManagementTemplate) GetWorkloadActions()([]WorkloadActionable) {
    if m == nil {
        return nil
    } else {
        return m.workloadActions
    }
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
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetInformationLinks()))
        for i, v := range m.GetInformationLinks() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
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
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetManagementTemplateCollections()))
        for i, v := range m.GetManagementTemplateCollections() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("managementTemplateCollections", cast)
        if err != nil {
            return err
        }
    }
    if m.GetManagementTemplateSteps() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetManagementTemplateSteps()))
        for i, v := range m.GetManagementTemplateSteps() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("managementTemplateSteps", cast)
        if err != nil {
            return err
        }
    }
    if m.GetParameters() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetParameters()))
        for i, v := range m.GetParameters() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
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
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetWorkloadActions()))
        for i, v := range m.GetWorkloadActions() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("workloadActions", cast)
        if err != nil {
            return err
        }
    }
    return nil
}
// SetCategory sets the category property value. The management category for the management template. Possible values are: custom, devices, identity, unknownFutureValue. Required. Read-only.
func (m *ManagementTemplate) SetCategory(value *ManagementCategory)() {
    if m != nil {
        m.category = value
    }
}
// SetCreatedByUserId sets the createdByUserId property value. The createdByUserId property
func (m *ManagementTemplate) SetCreatedByUserId(value *string)() {
    if m != nil {
        m.createdByUserId = value
    }
}
// SetCreatedDateTime sets the createdDateTime property value. The createdDateTime property
func (m *ManagementTemplate) SetCreatedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    if m != nil {
        m.createdDateTime = value
    }
}
// SetDescription sets the description property value. The description for the management template. Optional. Read-only.
func (m *ManagementTemplate) SetDescription(value *string)() {
    if m != nil {
        m.description = value
    }
}
// SetDisplayName sets the displayName property value. The display name for the management template. Required. Read-only.
func (m *ManagementTemplate) SetDisplayName(value *string)() {
    if m != nil {
        m.displayName = value
    }
}
// SetInformationLinks sets the informationLinks property value. The informationLinks property
func (m *ManagementTemplate) SetInformationLinks(value []ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ActionUrlable)() {
    if m != nil {
        m.informationLinks = value
    }
}
// SetLastActionByUserId sets the lastActionByUserId property value. The lastActionByUserId property
func (m *ManagementTemplate) SetLastActionByUserId(value *string)() {
    if m != nil {
        m.lastActionByUserId = value
    }
}
// SetLastActionDateTime sets the lastActionDateTime property value. The lastActionDateTime property
func (m *ManagementTemplate) SetLastActionDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    if m != nil {
        m.lastActionDateTime = value
    }
}
// SetManagementTemplateCollections sets the managementTemplateCollections property value. The managementTemplateCollections property
func (m *ManagementTemplate) SetManagementTemplateCollections(value []ManagementTemplateCollectionable)() {
    if m != nil {
        m.managementTemplateCollections = value
    }
}
// SetManagementTemplateSteps sets the managementTemplateSteps property value. The managementTemplateSteps property
func (m *ManagementTemplate) SetManagementTemplateSteps(value []ManagementTemplateStepable)() {
    if m != nil {
        m.managementTemplateSteps = value
    }
}
// SetParameters sets the parameters property value. The collection of parameters used by the management template. Optional. Read-only.
func (m *ManagementTemplate) SetParameters(value []TemplateParameterable)() {
    if m != nil {
        m.parameters = value
    }
}
// SetPriority sets the priority property value. The priority property
func (m *ManagementTemplate) SetPriority(value *int32)() {
    if m != nil {
        m.priority = value
    }
}
// SetProvider sets the provider property value. The provider property
func (m *ManagementTemplate) SetProvider(value *ManagementProvider)() {
    if m != nil {
        m.provider = value
    }
}
// SetUserImpact sets the userImpact property value. The userImpact property
func (m *ManagementTemplate) SetUserImpact(value *string)() {
    if m != nil {
        m.userImpact = value
    }
}
// SetVersion sets the version property value. The version property
func (m *ManagementTemplate) SetVersion(value *int32)() {
    if m != nil {
        m.version = value
    }
}
// SetWorkloadActions sets the workloadActions property value. The collection of workload actions associated with the management template. Optional. Read-only.
func (m *ManagementTemplate) SetWorkloadActions(value []WorkloadActionable)() {
    if m != nil {
        m.workloadActions = value
    }
}
