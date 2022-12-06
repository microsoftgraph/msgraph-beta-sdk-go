package managedtenants

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
)

// WorkloadActionDeploymentStatus 
type WorkloadActionDeploymentStatus struct {
    // The unique identifier for the workload action. Required. Read-only.
    actionId *string
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // The identifier of any policy that was created by applying the workload action. Optional. Read-only.
    deployedPolicyId *string
    // The detailed information for exceptions that occur when deploying the workload action. Optional. Required.
    error ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.GenericErrorable
    // The excludeGroups property
    excludeGroups []string
    // The includeAllUsers property
    includeAllUsers *bool
    // The includeGroups property
    includeGroups []string
    // The date and time the workload action was last deployed. Optional.
    lastDeploymentDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // The OdataType property
    odataType *string
    // The status property
    status *WorkloadActionStatus
}
// NewWorkloadActionDeploymentStatus instantiates a new workloadActionDeploymentStatus and sets the default values.
func NewWorkloadActionDeploymentStatus()(*WorkloadActionDeploymentStatus) {
    m := &WorkloadActionDeploymentStatus{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateWorkloadActionDeploymentStatusFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateWorkloadActionDeploymentStatusFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewWorkloadActionDeploymentStatus(), nil
}
// GetActionId gets the actionId property value. The unique identifier for the workload action. Required. Read-only.
func (m *WorkloadActionDeploymentStatus) GetActionId()(*string) {
    return m.actionId
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *WorkloadActionDeploymentStatus) GetAdditionalData()(map[string]interface{}) {
    return m.additionalData
}
// GetDeployedPolicyId gets the deployedPolicyId property value. The identifier of any policy that was created by applying the workload action. Optional. Read-only.
func (m *WorkloadActionDeploymentStatus) GetDeployedPolicyId()(*string) {
    return m.deployedPolicyId
}
// GetError gets the error property value. The detailed information for exceptions that occur when deploying the workload action. Optional. Required.
func (m *WorkloadActionDeploymentStatus) GetError()(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.GenericErrorable) {
    return m.error
}
// GetExcludeGroups gets the excludeGroups property value. The excludeGroups property
func (m *WorkloadActionDeploymentStatus) GetExcludeGroups()([]string) {
    return m.excludeGroups
}
// GetFieldDeserializers the deserialization information for the current model
func (m *WorkloadActionDeploymentStatus) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["actionId"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetActionId)
    res["deployedPolicyId"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetDeployedPolicyId)
    res["error"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetObjectValue(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateGenericErrorFromDiscriminatorValue , m.SetError)
    res["excludeGroups"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfPrimitiveValues("string" , m.SetExcludeGroups)
    res["includeAllUsers"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetBoolValue(m.SetIncludeAllUsers)
    res["includeGroups"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfPrimitiveValues("string" , m.SetIncludeGroups)
    res["lastDeploymentDateTime"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetTimeValue(m.SetLastDeploymentDateTime)
    res["@odata.type"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetOdataType)
    res["status"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetEnumValue(ParseWorkloadActionStatus , m.SetStatus)
    return res
}
// GetIncludeAllUsers gets the includeAllUsers property value. The includeAllUsers property
func (m *WorkloadActionDeploymentStatus) GetIncludeAllUsers()(*bool) {
    return m.includeAllUsers
}
// GetIncludeGroups gets the includeGroups property value. The includeGroups property
func (m *WorkloadActionDeploymentStatus) GetIncludeGroups()([]string) {
    return m.includeGroups
}
// GetLastDeploymentDateTime gets the lastDeploymentDateTime property value. The date and time the workload action was last deployed. Optional.
func (m *WorkloadActionDeploymentStatus) GetLastDeploymentDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.lastDeploymentDateTime
}
// GetOdataType gets the @odata.type property value. The OdataType property
func (m *WorkloadActionDeploymentStatus) GetOdataType()(*string) {
    return m.odataType
}
// GetStatus gets the status property value. The status property
func (m *WorkloadActionDeploymentStatus) GetStatus()(*WorkloadActionStatus) {
    return m.status
}
// Serialize serializes information the current object
func (m *WorkloadActionDeploymentStatus) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("actionId", m.GetActionId())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("deployedPolicyId", m.GetDeployedPolicyId())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("error", m.GetError())
        if err != nil {
            return err
        }
    }
    if m.GetExcludeGroups() != nil {
        err := writer.WriteCollectionOfStringValues("excludeGroups", m.GetExcludeGroups())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("includeAllUsers", m.GetIncludeAllUsers())
        if err != nil {
            return err
        }
    }
    if m.GetIncludeGroups() != nil {
        err := writer.WriteCollectionOfStringValues("includeGroups", m.GetIncludeGroups())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteTimeValue("lastDeploymentDateTime", m.GetLastDeploymentDateTime())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("@odata.type", m.GetOdataType())
        if err != nil {
            return err
        }
    }
    if m.GetStatus() != nil {
        cast := (*m.GetStatus()).String()
        err := writer.WriteStringValue("status", &cast)
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
// SetActionId sets the actionId property value. The unique identifier for the workload action. Required. Read-only.
func (m *WorkloadActionDeploymentStatus) SetActionId(value *string)() {
    m.actionId = value
}
// SetAdditionalData sets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *WorkloadActionDeploymentStatus) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// SetDeployedPolicyId sets the deployedPolicyId property value. The identifier of any policy that was created by applying the workload action. Optional. Read-only.
func (m *WorkloadActionDeploymentStatus) SetDeployedPolicyId(value *string)() {
    m.deployedPolicyId = value
}
// SetError sets the error property value. The detailed information for exceptions that occur when deploying the workload action. Optional. Required.
func (m *WorkloadActionDeploymentStatus) SetError(value ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.GenericErrorable)() {
    m.error = value
}
// SetExcludeGroups sets the excludeGroups property value. The excludeGroups property
func (m *WorkloadActionDeploymentStatus) SetExcludeGroups(value []string)() {
    m.excludeGroups = value
}
// SetIncludeAllUsers sets the includeAllUsers property value. The includeAllUsers property
func (m *WorkloadActionDeploymentStatus) SetIncludeAllUsers(value *bool)() {
    m.includeAllUsers = value
}
// SetIncludeGroups sets the includeGroups property value. The includeGroups property
func (m *WorkloadActionDeploymentStatus) SetIncludeGroups(value []string)() {
    m.includeGroups = value
}
// SetLastDeploymentDateTime sets the lastDeploymentDateTime property value. The date and time the workload action was last deployed. Optional.
func (m *WorkloadActionDeploymentStatus) SetLastDeploymentDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.lastDeploymentDateTime = value
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *WorkloadActionDeploymentStatus) SetOdataType(value *string)() {
    m.odataType = value
}
// SetStatus sets the status property value. The status property
func (m *WorkloadActionDeploymentStatus) SetStatus(value *WorkloadActionStatus)() {
    m.status = value
}
