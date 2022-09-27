package managedtenants

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
)

// ManagementTemplateStepVersion 
type ManagementTemplateStepVersion struct {
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Entity
    // The acceptedFor property
    acceptedFor ManagementTemplateStepable
    // The contentMarkdown property
    contentMarkdown *string
    // The createdByUserId property
    createdByUserId *string
    // The createdDateTime property
    createdDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // The deployments property
    deployments []ManagementTemplateStepDeploymentable
    // The lastActionByUserId property
    lastActionByUserId *string
    // The lastActionDateTime property
    lastActionDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // The name property
    name *string
    // The templateStep property
    templateStep ManagementTemplateStepable
    // The version property
    version *int32
    // The versionInformation property
    versionInformation *string
}
// NewManagementTemplateStepVersion instantiates a new managementTemplateStepVersion and sets the default values.
func NewManagementTemplateStepVersion()(*ManagementTemplateStepVersion) {
    m := &ManagementTemplateStepVersion{
        Entity: *ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.NewEntity(),
    }
    odataTypeValue := "#microsoft.graph.managedTenants.managementTemplateStepVersion";
    m.SetOdataType(&odataTypeValue);
    return m
}
// CreateManagementTemplateStepVersionFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateManagementTemplateStepVersionFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewManagementTemplateStepVersion(), nil
}
// GetAcceptedFor gets the acceptedFor property value. The acceptedFor property
func (m *ManagementTemplateStepVersion) GetAcceptedFor()(ManagementTemplateStepable) {
    return m.acceptedFor
}
// GetContentMarkdown gets the contentMarkdown property value. The contentMarkdown property
func (m *ManagementTemplateStepVersion) GetContentMarkdown()(*string) {
    return m.contentMarkdown
}
// GetCreatedByUserId gets the createdByUserId property value. The createdByUserId property
func (m *ManagementTemplateStepVersion) GetCreatedByUserId()(*string) {
    return m.createdByUserId
}
// GetCreatedDateTime gets the createdDateTime property value. The createdDateTime property
func (m *ManagementTemplateStepVersion) GetCreatedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.createdDateTime
}
// GetDeployments gets the deployments property value. The deployments property
func (m *ManagementTemplateStepVersion) GetDeployments()([]ManagementTemplateStepDeploymentable) {
    return m.deployments
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ManagementTemplateStepVersion) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["acceptedFor"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetObjectValue(CreateManagementTemplateStepFromDiscriminatorValue , m.SetAcceptedFor)
    res["contentMarkdown"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetContentMarkdown)
    res["createdByUserId"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetCreatedByUserId)
    res["createdDateTime"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetTimeValue(m.SetCreatedDateTime)
    res["deployments"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfObjectValues(CreateManagementTemplateStepDeploymentFromDiscriminatorValue , m.SetDeployments)
    res["lastActionByUserId"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetLastActionByUserId)
    res["lastActionDateTime"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetTimeValue(m.SetLastActionDateTime)
    res["name"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetName)
    res["templateStep"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetObjectValue(CreateManagementTemplateStepFromDiscriminatorValue , m.SetTemplateStep)
    res["version"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetInt32Value(m.SetVersion)
    res["versionInformation"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetVersionInformation)
    return res
}
// GetLastActionByUserId gets the lastActionByUserId property value. The lastActionByUserId property
func (m *ManagementTemplateStepVersion) GetLastActionByUserId()(*string) {
    return m.lastActionByUserId
}
// GetLastActionDateTime gets the lastActionDateTime property value. The lastActionDateTime property
func (m *ManagementTemplateStepVersion) GetLastActionDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.lastActionDateTime
}
// GetName gets the name property value. The name property
func (m *ManagementTemplateStepVersion) GetName()(*string) {
    return m.name
}
// GetTemplateStep gets the templateStep property value. The templateStep property
func (m *ManagementTemplateStepVersion) GetTemplateStep()(ManagementTemplateStepable) {
    return m.templateStep
}
// GetVersion gets the version property value. The version property
func (m *ManagementTemplateStepVersion) GetVersion()(*int32) {
    return m.version
}
// GetVersionInformation gets the versionInformation property value. The versionInformation property
func (m *ManagementTemplateStepVersion) GetVersionInformation()(*string) {
    return m.versionInformation
}
// Serialize serializes information the current object
func (m *ManagementTemplateStepVersion) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteObjectValue("acceptedFor", m.GetAcceptedFor())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("contentMarkdown", m.GetContentMarkdown())
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
    if m.GetDeployments() != nil {
        cast := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.CollectionCast[i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable](m.GetDeployments())
        err = writer.WriteCollectionOfObjectValues("deployments", cast)
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
    {
        err = writer.WriteStringValue("name", m.GetName())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("templateStep", m.GetTemplateStep())
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
    {
        err = writer.WriteStringValue("versionInformation", m.GetVersionInformation())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAcceptedFor sets the acceptedFor property value. The acceptedFor property
func (m *ManagementTemplateStepVersion) SetAcceptedFor(value ManagementTemplateStepable)() {
    m.acceptedFor = value
}
// SetContentMarkdown sets the contentMarkdown property value. The contentMarkdown property
func (m *ManagementTemplateStepVersion) SetContentMarkdown(value *string)() {
    m.contentMarkdown = value
}
// SetCreatedByUserId sets the createdByUserId property value. The createdByUserId property
func (m *ManagementTemplateStepVersion) SetCreatedByUserId(value *string)() {
    m.createdByUserId = value
}
// SetCreatedDateTime sets the createdDateTime property value. The createdDateTime property
func (m *ManagementTemplateStepVersion) SetCreatedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.createdDateTime = value
}
// SetDeployments sets the deployments property value. The deployments property
func (m *ManagementTemplateStepVersion) SetDeployments(value []ManagementTemplateStepDeploymentable)() {
    m.deployments = value
}
// SetLastActionByUserId sets the lastActionByUserId property value. The lastActionByUserId property
func (m *ManagementTemplateStepVersion) SetLastActionByUserId(value *string)() {
    m.lastActionByUserId = value
}
// SetLastActionDateTime sets the lastActionDateTime property value. The lastActionDateTime property
func (m *ManagementTemplateStepVersion) SetLastActionDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.lastActionDateTime = value
}
// SetName sets the name property value. The name property
func (m *ManagementTemplateStepVersion) SetName(value *string)() {
    m.name = value
}
// SetTemplateStep sets the templateStep property value. The templateStep property
func (m *ManagementTemplateStepVersion) SetTemplateStep(value ManagementTemplateStepable)() {
    m.templateStep = value
}
// SetVersion sets the version property value. The version property
func (m *ManagementTemplateStepVersion) SetVersion(value *int32)() {
    m.version = value
}
// SetVersionInformation sets the versionInformation property value. The versionInformation property
func (m *ManagementTemplateStepVersion) SetVersionInformation(value *string)() {
    m.versionInformation = value
}
