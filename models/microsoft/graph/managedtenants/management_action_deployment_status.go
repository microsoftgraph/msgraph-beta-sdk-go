package managedtenants

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// ManagementActionDeploymentStatus provides operations to manage the tenantRelationship singleton.
type ManagementActionDeploymentStatus struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // The identifier for the management action. Required. Read-only.
    managementActionId *string;
    // The management template identifier that was used to generate the management action. Required. Read-only.
    managementTemplateId *string;
    // 
    managementTemplateVersion *int32;
    // The status of the management action. Possible values are: toAddress, completed, error, timeOut, inProgress, planned, resolvedBy3rdParty, resolvedThroughAlternateMitigation, riskAccepted, unknownFutureValue. Required.
    status *ManagementActionStatus;
    // The collection of workload action deployment statues for the given management action. Optional.
    workloadActionDeploymentStatuses []WorkloadActionDeploymentStatusable;
}
// NewManagementActionDeploymentStatus instantiates a new managementActionDeploymentStatus and sets the default values.
func NewManagementActionDeploymentStatus()(*ManagementActionDeploymentStatus) {
    m := &ManagementActionDeploymentStatus{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateManagementActionDeploymentStatusFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateManagementActionDeploymentStatusFromDiscriminatorValue(parseNode i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, error) {
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
func (m *ManagementActionDeploymentStatus) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["managementActionId"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetManagementActionId(val)
        }
        return nil
    }
    res["managementTemplateId"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetManagementTemplateId(val)
        }
        return nil
    }
    res["managementTemplateVersion"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetManagementTemplateVersion(val)
        }
        return nil
    }
    res["status"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParseManagementActionStatus)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetStatus(val.(*ManagementActionStatus))
        }
        return nil
    }
    res["workloadActionDeploymentStatuses"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
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
// GetManagementTemplateVersion gets the managementTemplateVersion property value. 
func (m *ManagementActionDeploymentStatus) GetManagementTemplateVersion()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.managementTemplateVersion
    }
}
// GetStatus gets the status property value. The status of the management action. Possible values are: toAddress, completed, error, timeOut, inProgress, planned, resolvedBy3rdParty, resolvedThroughAlternateMitigation, riskAccepted, unknownFutureValue. Required.
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
func (m *ManagementActionDeploymentStatus) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
func (m *ManagementActionDeploymentStatus) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
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
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetWorkloadActionDeploymentStatuses()))
        for i, v := range m.GetWorkloadActionDeploymentStatuses() {
            cast[i] = v.(i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable)
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
// SetManagementTemplateVersion sets the managementTemplateVersion property value. 
func (m *ManagementActionDeploymentStatus) SetManagementTemplateVersion(value *int32)() {
    if m != nil {
        m.managementTemplateVersion = value
    }
}
// SetStatus sets the status property value. The status of the management action. Possible values are: toAddress, completed, error, timeOut, inProgress, planned, resolvedBy3rdParty, resolvedThroughAlternateMitigation, riskAccepted, unknownFutureValue. Required.
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