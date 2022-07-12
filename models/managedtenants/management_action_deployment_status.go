package managedtenants

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ManagementActionDeploymentStatus 
type ManagementActionDeploymentStatus struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // The identifier for the management action. Required. Read-only.
    managementActionId *string
    // The management template identifier that was used to generate the management action. Required. Read-only.
    managementTemplateId *string
    // The managementTemplateVersion property
    managementTemplateVersion *int32
    // The status property
    status *ManagementActionStatus
    // The collection of workload action deployment statues for the given management action. Optional.
    workloadActionDeploymentStatuses []WorkloadActionDeploymentStatusable
}
// NewManagementActionDeploymentStatus instantiates a new managementActionDeploymentStatus and sets the default values.
func NewManagementActionDeploymentStatus()(*ManagementActionDeploymentStatus) {
    m := &ManagementActionDeploymentStatus{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateManagementActionDeploymentStatusFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateManagementActionDeploymentStatusFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewManagementActionDeploymentStatus(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *ManagementActionDeploymentStatus) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ManagementActionDeploymentStatus) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["managementActionId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetManagementActionId(val)
        }
        return nil
    }
    res["managementTemplateId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetManagementTemplateId(val)
        }
        return nil
    }
    res["managementTemplateVersion"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetManagementTemplateVersion(val)
        }
        return nil
    }
    res["status"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseManagementActionStatus)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetStatus(val.(*ManagementActionStatus))
        }
        return nil
    }
    res["workloadActionDeploymentStatuses"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateWorkloadActionDeploymentStatusFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]WorkloadActionDeploymentStatusable, len(val))
            for i, v := range val {
                res[i] = v.(WorkloadActionDeploymentStatusable)
            }
            m.SetWorkloadActionDeploymentStatuses(res)
        }
        return nil
    }
    return res
}
// GetManagementActionId gets the managementActionId property value. The identifier for the management action. Required. Read-only.
func (m *ManagementActionDeploymentStatus) GetManagementActionId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.managementActionId
    }
}
// GetManagementTemplateId gets the managementTemplateId property value. The management template identifier that was used to generate the management action. Required. Read-only.
func (m *ManagementActionDeploymentStatus) GetManagementTemplateId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.managementTemplateId
    }
}
// GetManagementTemplateVersion gets the managementTemplateVersion property value. The managementTemplateVersion property
func (m *ManagementActionDeploymentStatus) GetManagementTemplateVersion()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.managementTemplateVersion
    }
}
// GetStatus gets the status property value. The status property
func (m *ManagementActionDeploymentStatus) GetStatus()(*ManagementActionStatus) {
    if m == nil {
        return nil
    } else {
        return m.status
    }
}
// GetWorkloadActionDeploymentStatuses gets the workloadActionDeploymentStatuses property value. The collection of workload action deployment statues for the given management action. Optional.
func (m *ManagementActionDeploymentStatus) GetWorkloadActionDeploymentStatuses()([]WorkloadActionDeploymentStatusable) {
    if m == nil {
        return nil
    } else {
        return m.workloadActionDeploymentStatuses
    }
}
// Serialize serializes information the current object
func (m *ManagementActionDeploymentStatus) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("managementActionId", m.GetManagementActionId())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("managementTemplateId", m.GetManagementTemplateId())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("managementTemplateVersion", m.GetManagementTemplateVersion())
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
    if m.GetWorkloadActionDeploymentStatuses() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetWorkloadActionDeploymentStatuses()))
        for i, v := range m.GetWorkloadActionDeploymentStatuses() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err := writer.WriteCollectionOfObjectValues("workloadActionDeploymentStatuses", cast)
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
// SetAdditionalData sets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *ManagementActionDeploymentStatus) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
// SetManagementActionId sets the managementActionId property value. The identifier for the management action. Required. Read-only.
func (m *ManagementActionDeploymentStatus) SetManagementActionId(value *string)() {
    if m != nil {
        m.managementActionId = value
    }
}
// SetManagementTemplateId sets the managementTemplateId property value. The management template identifier that was used to generate the management action. Required. Read-only.
func (m *ManagementActionDeploymentStatus) SetManagementTemplateId(value *string)() {
    if m != nil {
        m.managementTemplateId = value
    }
}
// SetManagementTemplateVersion sets the managementTemplateVersion property value. The managementTemplateVersion property
func (m *ManagementActionDeploymentStatus) SetManagementTemplateVersion(value *int32)() {
    if m != nil {
        m.managementTemplateVersion = value
    }
}
// SetStatus sets the status property value. The status property
func (m *ManagementActionDeploymentStatus) SetStatus(value *ManagementActionStatus)() {
    if m != nil {
        m.status = value
    }
}
// SetWorkloadActionDeploymentStatuses sets the workloadActionDeploymentStatuses property value. The collection of workload action deployment statues for the given management action. Optional.
func (m *ManagementActionDeploymentStatus) SetWorkloadActionDeploymentStatuses(value []WorkloadActionDeploymentStatusable)() {
    if m != nil {
        m.workloadActionDeploymentStatuses = value
    }
}
