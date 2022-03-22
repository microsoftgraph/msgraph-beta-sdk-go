package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// GovernanceRoleAssignmentRequestStatus 
type GovernanceRoleAssignmentRequestStatus struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // The status of the role assignment request. The value can be InProgress or Closed.
    status *string;
    // The details of the status of the role assignment request. It represents the evaluation results of different rules.
    statusDetails []KeyValueable;
    // The sub status of the role assignment request. The values can be Accepted, PendingEvaluation, Granted, Denied, PendingProvisioning, Provisioned, PendingRevocation, Revoked, Canceled, Failed, PendingApprovalProvisioning, PendingApproval, FailedAsResourceIsLocked, PendingAdminDecision, AdminApproved, AdminDenied, TimedOut, and ProvisioningStarted.
    subStatus *string;
}
// NewGovernanceRoleAssignmentRequestStatus instantiates a new governanceRoleAssignmentRequestStatus and sets the default values.
func NewGovernanceRoleAssignmentRequestStatus()(*GovernanceRoleAssignmentRequestStatus) {
    m := &GovernanceRoleAssignmentRequestStatus{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateGovernanceRoleAssignmentRequestStatusFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateGovernanceRoleAssignmentRequestStatusFromDiscriminatorValue(parseNode i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, error) {
    return NewGovernanceRoleAssignmentRequestStatus(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *GovernanceRoleAssignmentRequestStatus) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *GovernanceRoleAssignmentRequestStatus) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["status"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetStatus(val)
        }
        return nil
    }
    res["statusDetails"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateKeyValueFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]KeyValueable, len(val))
            for i, v := range val {
                res[i] = v.(KeyValueable)
            }
            m.SetStatusDetails(res)
        }
        return nil
    }
    res["subStatus"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSubStatus(val)
        }
        return nil
    }
    return res
}
// GetStatus gets the status property value. The status of the role assignment request. The value can be InProgress or Closed.
func (m *GovernanceRoleAssignmentRequestStatus) GetStatus()(*string) {
    if m == nil {
        return nil
    } else {
        return m.status
    }
}
// GetStatusDetails gets the statusDetails property value. The details of the status of the role assignment request. It represents the evaluation results of different rules.
func (m *GovernanceRoleAssignmentRequestStatus) GetStatusDetails()([]KeyValueable) {
    if m == nil {
        return nil
    } else {
        return m.statusDetails
    }
}
// GetSubStatus gets the subStatus property value. The sub status of the role assignment request. The values can be Accepted, PendingEvaluation, Granted, Denied, PendingProvisioning, Provisioned, PendingRevocation, Revoked, Canceled, Failed, PendingApprovalProvisioning, PendingApproval, FailedAsResourceIsLocked, PendingAdminDecision, AdminApproved, AdminDenied, TimedOut, and ProvisioningStarted.
func (m *GovernanceRoleAssignmentRequestStatus) GetSubStatus()(*string) {
    if m == nil {
        return nil
    } else {
        return m.subStatus
    }
}
// Serialize serializes information the current object
func (m *GovernanceRoleAssignmentRequestStatus) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("status", m.GetStatus())
        if err != nil {
            return err
        }
    }
    if m.GetStatusDetails() != nil {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetStatusDetails()))
        for i, v := range m.GetStatusDetails() {
            cast[i] = v.(i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable)
        }
        err := writer.WriteCollectionOfObjectValues("statusDetails", cast)
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("subStatus", m.GetSubStatus())
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
func (m *GovernanceRoleAssignmentRequestStatus) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
// SetStatus sets the status property value. The status of the role assignment request. The value can be InProgress or Closed.
func (m *GovernanceRoleAssignmentRequestStatus) SetStatus(value *string)() {
    if m != nil {
        m.status = value
    }
}
// SetStatusDetails sets the statusDetails property value. The details of the status of the role assignment request. It represents the evaluation results of different rules.
func (m *GovernanceRoleAssignmentRequestStatus) SetStatusDetails(value []KeyValueable)() {
    if m != nil {
        m.statusDetails = value
    }
}
// SetSubStatus sets the subStatus property value. The sub status of the role assignment request. The values can be Accepted, PendingEvaluation, Granted, Denied, PendingProvisioning, Provisioned, PendingRevocation, Revoked, Canceled, Failed, PendingApprovalProvisioning, PendingApproval, FailedAsResourceIsLocked, PendingAdminDecision, AdminApproved, AdminDenied, TimedOut, and ProvisioningStarted.
func (m *GovernanceRoleAssignmentRequestStatus) SetSubStatus(value *string)() {
    if m != nil {
        m.subStatus = value
    }
}
