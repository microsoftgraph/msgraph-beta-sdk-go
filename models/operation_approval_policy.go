package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// OperationApprovalPolicy the OperationApprovalPolicy entity allows an administrator to configure which operations require admin approval and the set of admins who can perform that approval. Creating a policy enables the multiple admin approval service to catch requests which are targeted by the specific policy type defined.
type OperationApprovalPolicy struct {
    Entity
}
// NewOperationApprovalPolicy instantiates a new operationApprovalPolicy and sets the default values.
func NewOperationApprovalPolicy()(*OperationApprovalPolicy) {
    m := &OperationApprovalPolicy{
        Entity: *NewEntity(),
    }
    return m
}
// CreateOperationApprovalPolicyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateOperationApprovalPolicyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewOperationApprovalPolicy(), nil
}
// GetApproverGroupIds gets the approverGroupIds property value. The Microsoft Entra ID (Azure AD) security group IDs for the approvers for the policy. This property is required when the policy is created, and is defined by the user to define the possible approvers for the policy.
func (m *OperationApprovalPolicy) GetApproverGroupIds()([]string) {
    val, err := m.GetBackingStore().Get("approverGroupIds")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]string)
    }
    return nil
}
// GetDescription gets the description property value. Indicates the description of the policy. Maximum length of the description is 1024 characters. This property is not required, but can be used by the user to describe the policy.
func (m *OperationApprovalPolicy) GetDescription()(*string) {
    val, err := m.GetBackingStore().Get("description")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetDisplayName gets the displayName property value. Indicates the display name of the policy. Maximum length of the display name is 128 characters. This property is required when the policy is created, and is defined by the user to identify the policy.
func (m *OperationApprovalPolicy) GetDisplayName()(*string) {
    val, err := m.GetBackingStore().Get("displayName")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *OperationApprovalPolicy) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["approverGroupIds"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]string, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = *(v.(*string))
                }
            }
            m.SetApproverGroupIds(res)
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
    res["lastModifiedDateTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLastModifiedDateTime(val)
        }
        return nil
    }
    res["policyPlatform"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseOperationApprovalPolicyPlatform)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPolicyPlatform(val.(*OperationApprovalPolicyPlatform))
        }
        return nil
    }
    res["policyType"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseOperationApprovalPolicyType)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPolicyType(val.(*OperationApprovalPolicyType))
        }
        return nil
    }
    return res
}
// GetLastModifiedDateTime gets the lastModifiedDateTime property value. Indicates the last DateTime that the policy was modified. The value cannot be modified and is automatically populated whenever values in the request are updated. For example, when the 'policyType' property changes from apps to scripts. The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 would look like this: '2014-01-01T00:00:00Z'. Returned by default. Read-only. This property is read-only.
func (m *OperationApprovalPolicy) GetLastModifiedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    val, err := m.GetBackingStore().Get("lastModifiedDateTime")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    }
    return nil
}
// GetPolicyPlatform gets the policyPlatform property value. The set of available platforms for the OperationApprovalPolicy. Allows configuration of a policy to specific platform(s) for approval. If no specific platform is required or applicable, the platform is `notApplicable`.
func (m *OperationApprovalPolicy) GetPolicyPlatform()(*OperationApprovalPolicyPlatform) {
    val, err := m.GetBackingStore().Get("policyPlatform")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*OperationApprovalPolicyPlatform)
    }
    return nil
}
// GetPolicyType gets the policyType property value. The set of available policy types that can be configured for approval. There is no default value for this enum, indicating that the policy type must always be chosen.
func (m *OperationApprovalPolicy) GetPolicyType()(*OperationApprovalPolicyType) {
    val, err := m.GetBackingStore().Get("policyType")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*OperationApprovalPolicyType)
    }
    return nil
}
// Serialize serializes information the current object
func (m *OperationApprovalPolicy) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetApproverGroupIds() != nil {
        err = writer.WriteCollectionOfStringValues("approverGroupIds", m.GetApproverGroupIds())
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
    if m.GetPolicyPlatform() != nil {
        cast := (*m.GetPolicyPlatform()).String()
        err = writer.WriteStringValue("policyPlatform", &cast)
        if err != nil {
            return err
        }
    }
    if m.GetPolicyType() != nil {
        cast := (*m.GetPolicyType()).String()
        err = writer.WriteStringValue("policyType", &cast)
        if err != nil {
            return err
        }
    }
    return nil
}
// SetApproverGroupIds sets the approverGroupIds property value. The Microsoft Entra ID (Azure AD) security group IDs for the approvers for the policy. This property is required when the policy is created, and is defined by the user to define the possible approvers for the policy.
func (m *OperationApprovalPolicy) SetApproverGroupIds(value []string)() {
    err := m.GetBackingStore().Set("approverGroupIds", value)
    if err != nil {
        panic(err)
    }
}
// SetDescription sets the description property value. Indicates the description of the policy. Maximum length of the description is 1024 characters. This property is not required, but can be used by the user to describe the policy.
func (m *OperationApprovalPolicy) SetDescription(value *string)() {
    err := m.GetBackingStore().Set("description", value)
    if err != nil {
        panic(err)
    }
}
// SetDisplayName sets the displayName property value. Indicates the display name of the policy. Maximum length of the display name is 128 characters. This property is required when the policy is created, and is defined by the user to identify the policy.
func (m *OperationApprovalPolicy) SetDisplayName(value *string)() {
    err := m.GetBackingStore().Set("displayName", value)
    if err != nil {
        panic(err)
    }
}
// SetLastModifiedDateTime sets the lastModifiedDateTime property value. Indicates the last DateTime that the policy was modified. The value cannot be modified and is automatically populated whenever values in the request are updated. For example, when the 'policyType' property changes from apps to scripts. The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 would look like this: '2014-01-01T00:00:00Z'. Returned by default. Read-only. This property is read-only.
func (m *OperationApprovalPolicy) SetLastModifiedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    err := m.GetBackingStore().Set("lastModifiedDateTime", value)
    if err != nil {
        panic(err)
    }
}
// SetPolicyPlatform sets the policyPlatform property value. The set of available platforms for the OperationApprovalPolicy. Allows configuration of a policy to specific platform(s) for approval. If no specific platform is required or applicable, the platform is `notApplicable`.
func (m *OperationApprovalPolicy) SetPolicyPlatform(value *OperationApprovalPolicyPlatform)() {
    err := m.GetBackingStore().Set("policyPlatform", value)
    if err != nil {
        panic(err)
    }
}
// SetPolicyType sets the policyType property value. The set of available policy types that can be configured for approval. There is no default value for this enum, indicating that the policy type must always be chosen.
func (m *OperationApprovalPolicy) SetPolicyType(value *OperationApprovalPolicyType)() {
    err := m.GetBackingStore().Set("policyType", value)
    if err != nil {
        panic(err)
    }
}
// OperationApprovalPolicyable 
type OperationApprovalPolicyable interface {
    Entityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetApproverGroupIds()([]string)
    GetDescription()(*string)
    GetDisplayName()(*string)
    GetLastModifiedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetPolicyPlatform()(*OperationApprovalPolicyPlatform)
    GetPolicyType()(*OperationApprovalPolicyType)
    SetApproverGroupIds(value []string)()
    SetDescription(value *string)()
    SetDisplayName(value *string)()
    SetLastModifiedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetPolicyPlatform(value *OperationApprovalPolicyPlatform)()
    SetPolicyType(value *OperationApprovalPolicyType)()
}
