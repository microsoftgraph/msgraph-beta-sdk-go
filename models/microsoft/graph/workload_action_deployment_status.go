package graph

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i5c2592132064055aae424492b066923068e6d9a29d4565707b3591c21983fe01 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph/managedtenants"
)

// 
type WorkloadActionDeploymentStatus struct {
    // The unique identifier for the workload action. Required. Read-only.
    actionId *string;
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // The identifier of any policy that was created by applying the workload action. Optional. Read-only.
    deployedPolicyId *string;
    // The detailed information for exceptions that occur when deploying the workload action. Optional. Required.
    error *GenericError;
    // 
    excludeGroups []string;
    // 
    includeAllUsers *bool;
    // 
    includeGroups []string;
    // The date and time the workload action was last deployed. Optional.
    lastDeploymentDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time;
    // The status of the workload action deployment. Possible values are: toAddress, completed, error, timeOut, inProgress, unknownFutureValue. Required. Read-only.
    status *i5c2592132064055aae424492b066923068e6d9a29d4565707b3591c21983fe01.WorkloadActionStatus;
}
// Instantiates a new workloadActionDeploymentStatus and sets the default values.
func NewWorkloadActionDeploymentStatus()(*WorkloadActionDeploymentStatus) {
    m := &WorkloadActionDeploymentStatus{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// Gets the actionId property value. The unique identifier for the workload action. Required. Read-only.
func (m *WorkloadActionDeploymentStatus) GetActionId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.actionId
    }
}
// Gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *WorkloadActionDeploymentStatus) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// Gets the deployedPolicyId property value. The identifier of any policy that was created by applying the workload action. Optional. Read-only.
func (m *WorkloadActionDeploymentStatus) GetDeployedPolicyId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.deployedPolicyId
    }
}
// Gets the error property value. The detailed information for exceptions that occur when deploying the workload action. Optional. Required.
func (m *WorkloadActionDeploymentStatus) GetError()(*GenericError) {
    if m == nil {
        return nil
    } else {
        return m.error
    }
}
// Gets the excludeGroups property value. 
func (m *WorkloadActionDeploymentStatus) GetExcludeGroups()([]string) {
    if m == nil {
        return nil
    } else {
        return m.excludeGroups
    }
}
// Gets the includeAllUsers property value. 
func (m *WorkloadActionDeploymentStatus) GetIncludeAllUsers()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.includeAllUsers
    }
}
// Gets the includeGroups property value. 
func (m *WorkloadActionDeploymentStatus) GetIncludeGroups()([]string) {
    if m == nil {
        return nil
    } else {
        return m.includeGroups
    }
}
// Gets the lastDeploymentDateTime property value. The date and time the workload action was last deployed. Optional.
func (m *WorkloadActionDeploymentStatus) GetLastDeploymentDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.lastDeploymentDateTime
    }
}
// Gets the status property value. The status of the workload action deployment. Possible values are: toAddress, completed, error, timeOut, inProgress, unknownFutureValue. Required. Read-only.
func (m *WorkloadActionDeploymentStatus) GetStatus()(*i5c2592132064055aae424492b066923068e6d9a29d4565707b3591c21983fe01.WorkloadActionStatus) {
    if m == nil {
        return nil
    } else {
        return m.status
    }
}
// The deserialization information for the current model
func (m *WorkloadActionDeploymentStatus) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["actionId"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetActionId(val)
        }
        return nil
    }
    res["deployedPolicyId"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDeployedPolicyId(val)
        }
        return nil
    }
    res["error"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewGenericError() })
        if err != nil {
            return err
        }
        if val != nil {
            m.SetError(val.(*GenericError))
        }
        return nil
    }
    res["excludeGroups"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]string, len(val))
            for i, v := range val {
                res[i] = *(v.(*string))
            }
            m.SetExcludeGroups(res)
        }
        return nil
    }
    res["includeAllUsers"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIncludeAllUsers(val)
        }
        return nil
    }
    res["includeGroups"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]string, len(val))
            for i, v := range val {
                res[i] = *(v.(*string))
            }
            m.SetIncludeGroups(res)
        }
        return nil
    }
    res["lastDeploymentDateTime"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLastDeploymentDateTime(val)
        }
        return nil
    }
    res["status"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(i5c2592132064055aae424492b066923068e6d9a29d4565707b3591c21983fe01.ParseWorkloadActionStatus)
        if err != nil {
            return err
        }
        if val != nil {
            cast := val.(i5c2592132064055aae424492b066923068e6d9a29d4565707b3591c21983fe01.WorkloadActionStatus)
            m.SetStatus(&cast)
        }
        return nil
    }
    return res
}
func (m *WorkloadActionDeploymentStatus) IsNil()(bool) {
    return m == nil
}
// Serializes information the current object
// Parameters:
//  - writer : Serialization writer to use to serialize this model
func (m *WorkloadActionDeploymentStatus) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
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
    {
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
    {
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
    if m.GetStatus() != nil {
        cast := m.GetStatus().String()
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
// Sets the actionId property value. The unique identifier for the workload action. Required. Read-only.
// Parameters:
//  - value : Value to set for the actionId property.
func (m *WorkloadActionDeploymentStatus) SetActionId(value *string)() {
    m.actionId = value
}
// Sets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// Parameters:
//  - value : Value to set for the AdditionalData property.
func (m *WorkloadActionDeploymentStatus) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// Sets the deployedPolicyId property value. The identifier of any policy that was created by applying the workload action. Optional. Read-only.
// Parameters:
//  - value : Value to set for the deployedPolicyId property.
func (m *WorkloadActionDeploymentStatus) SetDeployedPolicyId(value *string)() {
    m.deployedPolicyId = value
}
// Sets the error property value. The detailed information for exceptions that occur when deploying the workload action. Optional. Required.
// Parameters:
//  - value : Value to set for the error property.
func (m *WorkloadActionDeploymentStatus) SetError(value *GenericError)() {
    m.error = value
}
// Sets the excludeGroups property value. 
// Parameters:
//  - value : Value to set for the excludeGroups property.
func (m *WorkloadActionDeploymentStatus) SetExcludeGroups(value []string)() {
    m.excludeGroups = value
}
// Sets the includeAllUsers property value. 
// Parameters:
//  - value : Value to set for the includeAllUsers property.
func (m *WorkloadActionDeploymentStatus) SetIncludeAllUsers(value *bool)() {
    m.includeAllUsers = value
}
// Sets the includeGroups property value. 
// Parameters:
//  - value : Value to set for the includeGroups property.
func (m *WorkloadActionDeploymentStatus) SetIncludeGroups(value []string)() {
    m.includeGroups = value
}
// Sets the lastDeploymentDateTime property value. The date and time the workload action was last deployed. Optional.
// Parameters:
//  - value : Value to set for the lastDeploymentDateTime property.
func (m *WorkloadActionDeploymentStatus) SetLastDeploymentDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.lastDeploymentDateTime = value
}
// Sets the status property value. The status of the workload action deployment. Possible values are: toAddress, completed, error, timeOut, inProgress, unknownFutureValue. Required. Read-only.
// Parameters:
//  - value : Value to set for the status property.
func (m *WorkloadActionDeploymentStatus) SetStatus(value *i5c2592132064055aae424492b066923068e6d9a29d4565707b3591c21983fe01.WorkloadActionStatus)() {
    m.status = value
}
