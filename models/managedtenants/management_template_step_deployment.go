package managedtenants

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
)

// ManagementTemplateStepDeployment provides operations to manage the collection of accessReviewDecision entities.
type ManagementTemplateStepDeployment struct {
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Entity
    // The createdByUserId property
    createdByUserId *string
    // The createdDateTime property
    createdDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // The error property
    error GraphAPIErrorDetailsable
    // The lastActionByUserId property
    lastActionByUserId *string
    // The lastActionDateTime property
    lastActionDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // The status property
    status *ManagementTemplateDeploymentStatus
    // The templateStepVersion property
    templateStepVersion ManagementTemplateStepVersionable
    // The tenantId property
    tenantId *string
}
// NewManagementTemplateStepDeployment instantiates a new managementTemplateStepDeployment and sets the default values.
func NewManagementTemplateStepDeployment()(*ManagementTemplateStepDeployment) {
    m := &ManagementTemplateStepDeployment{
        Entity: *ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.NewEntity(),
    }
    odataTypeValue := "#microsoft.graph.managedTenants.managementTemplateStepDeployment";
    m.SetOdataType(&odataTypeValue);
    return m
}
// CreateManagementTemplateStepDeploymentFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateManagementTemplateStepDeploymentFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewManagementTemplateStepDeployment(), nil
}
// GetCreatedByUserId gets the createdByUserId property value. The createdByUserId property
func (m *ManagementTemplateStepDeployment) GetCreatedByUserId()(*string) {
    return m.createdByUserId
}
// GetCreatedDateTime gets the createdDateTime property value. The createdDateTime property
func (m *ManagementTemplateStepDeployment) GetCreatedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.createdDateTime
}
// GetError gets the error property value. The error property
func (m *ManagementTemplateStepDeployment) GetError()(GraphAPIErrorDetailsable) {
    return m.error
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ManagementTemplateStepDeployment) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["createdByUserId"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetCreatedByUserId)
    res["createdDateTime"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetTimeValue(m.SetCreatedDateTime)
    res["error"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetObjectValue(CreateGraphAPIErrorDetailsFromDiscriminatorValue , m.SetError)
    res["lastActionByUserId"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetLastActionByUserId)
    res["lastActionDateTime"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetTimeValue(m.SetLastActionDateTime)
    res["status"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetEnumValue(ParseManagementTemplateDeploymentStatus , m.SetStatus)
    res["templateStepVersion"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetObjectValue(CreateManagementTemplateStepVersionFromDiscriminatorValue , m.SetTemplateStepVersion)
    res["tenantId"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetTenantId)
    return res
}
// GetLastActionByUserId gets the lastActionByUserId property value. The lastActionByUserId property
func (m *ManagementTemplateStepDeployment) GetLastActionByUserId()(*string) {
    return m.lastActionByUserId
}
// GetLastActionDateTime gets the lastActionDateTime property value. The lastActionDateTime property
func (m *ManagementTemplateStepDeployment) GetLastActionDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.lastActionDateTime
}
// GetStatus gets the status property value. The status property
func (m *ManagementTemplateStepDeployment) GetStatus()(*ManagementTemplateDeploymentStatus) {
    return m.status
}
// GetTemplateStepVersion gets the templateStepVersion property value. The templateStepVersion property
func (m *ManagementTemplateStepDeployment) GetTemplateStepVersion()(ManagementTemplateStepVersionable) {
    return m.templateStepVersion
}
// GetTenantId gets the tenantId property value. The tenantId property
func (m *ManagementTemplateStepDeployment) GetTenantId()(*string) {
    return m.tenantId
}
// Serialize serializes information the current object
func (m *ManagementTemplateStepDeployment) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
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
        err = writer.WriteObjectValue("error", m.GetError())
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
    if m.GetStatus() != nil {
        cast := (*m.GetStatus()).String()
        err = writer.WriteStringValue("status", &cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("templateStepVersion", m.GetTemplateStepVersion())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("tenantId", m.GetTenantId())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetCreatedByUserId sets the createdByUserId property value. The createdByUserId property
func (m *ManagementTemplateStepDeployment) SetCreatedByUserId(value *string)() {
    m.createdByUserId = value
}
// SetCreatedDateTime sets the createdDateTime property value. The createdDateTime property
func (m *ManagementTemplateStepDeployment) SetCreatedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.createdDateTime = value
}
// SetError sets the error property value. The error property
func (m *ManagementTemplateStepDeployment) SetError(value GraphAPIErrorDetailsable)() {
    m.error = value
}
// SetLastActionByUserId sets the lastActionByUserId property value. The lastActionByUserId property
func (m *ManagementTemplateStepDeployment) SetLastActionByUserId(value *string)() {
    m.lastActionByUserId = value
}
// SetLastActionDateTime sets the lastActionDateTime property value. The lastActionDateTime property
func (m *ManagementTemplateStepDeployment) SetLastActionDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.lastActionDateTime = value
}
// SetStatus sets the status property value. The status property
func (m *ManagementTemplateStepDeployment) SetStatus(value *ManagementTemplateDeploymentStatus)() {
    m.status = value
}
// SetTemplateStepVersion sets the templateStepVersion property value. The templateStepVersion property
func (m *ManagementTemplateStepDeployment) SetTemplateStepVersion(value ManagementTemplateStepVersionable)() {
    m.templateStepVersion = value
}
// SetTenantId sets the tenantId property value. The tenantId property
func (m *ManagementTemplateStepDeployment) SetTenantId(value *string)() {
    m.tenantId = value
}
