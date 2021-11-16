package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i5c2592132064055aae424492b066923068e6d9a29d4565707b3591c21983fe01 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph/managedtenants"
)

// 
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
    status *i5c2592132064055aae424492b066923068e6d9a29d4565707b3591c21983fe01.ManagementActionStatus;
    // The collection of workload action deployment statues for the given management action. Optional.
    workloadActionDeploymentStatuses []WorkloadActionDeploymentStatus;
}
// Instantiates a new managementActionDeploymentStatus and sets the default values.
func NewManagementActionDeploymentStatus()(*ManagementActionDeploymentStatus) {
    m := &ManagementActionDeploymentStatus{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// Gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *ManagementActionDeploymentStatus) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// Gets the managementActionId property value. The identifier for the management action. Required. Read-only.
func (m *ManagementActionDeploymentStatus) GetManagementActionId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.managementActionId
    }
}
// Gets the managementTemplateId property value. The management template identifier that was used to generate the management action. Required. Read-only.
func (m *ManagementActionDeploymentStatus) GetManagementTemplateId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.managementTemplateId
    }
}
// Gets the managementTemplateVersion property value. 
func (m *ManagementActionDeploymentStatus) GetManagementTemplateVersion()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.managementTemplateVersion
    }
}
// Gets the status property value. The status of the management action. Possible values are: toAddress, completed, error, timeOut, inProgress, planned, resolvedBy3rdParty, resolvedThroughAlternateMitigation, riskAccepted, unknownFutureValue. Required.
func (m *ManagementActionDeploymentStatus) GetStatus()(*i5c2592132064055aae424492b066923068e6d9a29d4565707b3591c21983fe01.ManagementActionStatus) {
    if m == nil {
        return nil
    } else {
        return m.status
    }
}
// Gets the workloadActionDeploymentStatuses property value. The collection of workload action deployment statues for the given management action. Optional.
func (m *ManagementActionDeploymentStatus) GetWorkloadActionDeploymentStatuses()([]WorkloadActionDeploymentStatus) {
    if m == nil {
        return nil
    } else {
        return m.workloadActionDeploymentStatuses
    }
}
// The deserialization information for the current model
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
        val, err := n.GetEnumValue(i5c2592132064055aae424492b066923068e6d9a29d4565707b3591c21983fe01.ParseManagementActionStatus)
        if err != nil {
            return err
        }
        if val != nil {
            cast := val.(i5c2592132064055aae424492b066923068e6d9a29d4565707b3591c21983fe01.ManagementActionStatus)
            m.SetStatus(&cast)
        }
        return nil
    }
    res["workloadActionDeploymentStatuses"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewWorkloadActionDeploymentStatus() })
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]WorkloadActionDeploymentStatus, len(val))
            for i, v := range val {
                res[i] = *(v.(*WorkloadActionDeploymentStatus))
            }
            m.SetWorkloadActionDeploymentStatuses(res)
        }
        return nil
    }
    return res
}
func (m *ManagementActionDeploymentStatus) IsNil()(bool) {
    return m == nil
}
// Serializes information the current object
// Parameters:
//  - writer : Serialization writer to use to serialize this model
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
        cast := m.GetStatus().String()
        err := writer.WriteStringValue("status", &cast)
        if err != nil {
            return err
        }
    }
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetWorkloadActionDeploymentStatuses()))
        for i, v := range m.GetWorkloadActionDeploymentStatuses() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
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
// Sets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// Parameters:
//  - value : Value to set for the AdditionalData property.
func (m *ManagementActionDeploymentStatus) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// Sets the managementActionId property value. The identifier for the management action. Required. Read-only.
// Parameters:
//  - value : Value to set for the managementActionId property.
func (m *ManagementActionDeploymentStatus) SetManagementActionId(value *string)() {
    m.managementActionId = value
}
// Sets the managementTemplateId property value. The management template identifier that was used to generate the management action. Required. Read-only.
// Parameters:
//  - value : Value to set for the managementTemplateId property.
func (m *ManagementActionDeploymentStatus) SetManagementTemplateId(value *string)() {
    m.managementTemplateId = value
}
// Sets the managementTemplateVersion property value. 
// Parameters:
//  - value : Value to set for the managementTemplateVersion property.
func (m *ManagementActionDeploymentStatus) SetManagementTemplateVersion(value *int32)() {
    m.managementTemplateVersion = value
}
// Sets the status property value. The status of the management action. Possible values are: toAddress, completed, error, timeOut, inProgress, planned, resolvedBy3rdParty, resolvedThroughAlternateMitigation, riskAccepted, unknownFutureValue. Required.
// Parameters:
//  - value : Value to set for the status property.
func (m *ManagementActionDeploymentStatus) SetStatus(value *i5c2592132064055aae424492b066923068e6d9a29d4565707b3591c21983fe01.ManagementActionStatus)() {
    m.status = value
}
// Sets the workloadActionDeploymentStatuses property value. The collection of workload action deployment statues for the given management action. Optional.
// Parameters:
//  - value : Value to set for the workloadActionDeploymentStatuses property.
func (m *ManagementActionDeploymentStatus) SetWorkloadActionDeploymentStatuses(value []WorkloadActionDeploymentStatus)() {
    m.workloadActionDeploymentStatuses = value
}
