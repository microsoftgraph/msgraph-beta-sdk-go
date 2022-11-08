package models

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// OfficeClientConfiguration 
type OfficeClientConfiguration struct {
    Entity
    // The list of group assignments for the policy.
    assignments []OfficeClientConfigurationAssignmentable
    // List of office Client check-in status.
    checkinStatuses []OfficeClientCheckinStatusable
    // Not yet documented
    description *string
    // Admin provided description of the office client configuration policy.
    displayName *string
    // Policy settings JSON string in binary format, these values cannot be changed by the user.
    policyPayload []byte
    // Priority value should be unique value for each policy under a tenant and will be used for conflict resolution, lower values mean priority is high.
    priority *int32
    // User check-in summary for the policy.
    userCheckinSummary OfficeUserCheckinSummaryable
    // Preference settings JSON string in binary format, these values can be overridden by the user.
    userPreferencePayload []byte
}
// NewOfficeClientConfiguration instantiates a new OfficeClientConfiguration and sets the default values.
func NewOfficeClientConfiguration()(*OfficeClientConfiguration) {
    m := &OfficeClientConfiguration{
        Entity: *NewEntity(),
    }
    odataTypeValue := "#microsoft.graph.officeClientConfiguration";
    m.SetOdataType(&odataTypeValue);
    return m
}
// CreateOfficeClientConfigurationFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateOfficeClientConfigurationFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    if parseNode != nil {
        mappingValueNode, err := parseNode.GetChildNode("@odata.type")
        if err != nil {
            return nil, err
        }
        if mappingValueNode != nil {
            mappingValue, err := mappingValueNode.GetStringValue()
            if err != nil {
                return nil, err
            }
            if mappingValue != nil {
                switch *mappingValue {
                    case "#microsoft.graph.windowsOfficeClientConfiguration":
                        return NewWindowsOfficeClientConfiguration(), nil
                    case "#microsoft.graph.windowsOfficeClientSecurityConfiguration":
                        return NewWindowsOfficeClientSecurityConfiguration(), nil
                }
            }
        }
    }
    return NewOfficeClientConfiguration(), nil
}
// GetAssignments gets the assignments property value. The list of group assignments for the policy.
func (m *OfficeClientConfiguration) GetAssignments()([]OfficeClientConfigurationAssignmentable) {
    return m.assignments
}
// GetCheckinStatuses gets the checkinStatuses property value. List of office Client check-in status.
func (m *OfficeClientConfiguration) GetCheckinStatuses()([]OfficeClientCheckinStatusable) {
    return m.checkinStatuses
}
// GetDescription gets the description property value. Not yet documented
func (m *OfficeClientConfiguration) GetDescription()(*string) {
    return m.description
}
// GetDisplayName gets the displayName property value. Admin provided description of the office client configuration policy.
func (m *OfficeClientConfiguration) GetDisplayName()(*string) {
    return m.displayName
}
// GetFieldDeserializers the deserialization information for the current model
func (m *OfficeClientConfiguration) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["assignments"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfObjectValues(CreateOfficeClientConfigurationAssignmentFromDiscriminatorValue , m.SetAssignments)
    res["checkinStatuses"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfObjectValues(CreateOfficeClientCheckinStatusFromDiscriminatorValue , m.SetCheckinStatuses)
    res["description"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetDescription)
    res["displayName"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetDisplayName)
    res["policyPayload"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetByteArrayValue(m.SetPolicyPayload)
    res["priority"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetInt32Value(m.SetPriority)
    res["userCheckinSummary"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetObjectValue(CreateOfficeUserCheckinSummaryFromDiscriminatorValue , m.SetUserCheckinSummary)
    res["userPreferencePayload"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetByteArrayValue(m.SetUserPreferencePayload)
    return res
}
// GetPolicyPayload gets the policyPayload property value. Policy settings JSON string in binary format, these values cannot be changed by the user.
func (m *OfficeClientConfiguration) GetPolicyPayload()([]byte) {
    return m.policyPayload
}
// GetPriority gets the priority property value. Priority value should be unique value for each policy under a tenant and will be used for conflict resolution, lower values mean priority is high.
func (m *OfficeClientConfiguration) GetPriority()(*int32) {
    return m.priority
}
// GetUserCheckinSummary gets the userCheckinSummary property value. User check-in summary for the policy.
func (m *OfficeClientConfiguration) GetUserCheckinSummary()(OfficeUserCheckinSummaryable) {
    return m.userCheckinSummary
}
// GetUserPreferencePayload gets the userPreferencePayload property value. Preference settings JSON string in binary format, these values can be overridden by the user.
func (m *OfficeClientConfiguration) GetUserPreferencePayload()([]byte) {
    return m.userPreferencePayload
}
// Serialize serializes information the current object
func (m *OfficeClientConfiguration) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetAssignments() != nil {
        cast := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.CollectionCast[i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable](m.GetAssignments())
        err = writer.WriteCollectionOfObjectValues("assignments", cast)
        if err != nil {
            return err
        }
    }
    if m.GetCheckinStatuses() != nil {
        cast := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.CollectionCast[i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable](m.GetCheckinStatuses())
        err = writer.WriteCollectionOfObjectValues("checkinStatuses", cast)
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
    {
        err = writer.WriteByteArrayValue("policyPayload", m.GetPolicyPayload())
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
    {
        err = writer.WriteObjectValue("userCheckinSummary", m.GetUserCheckinSummary())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteByteArrayValue("userPreferencePayload", m.GetUserPreferencePayload())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAssignments sets the assignments property value. The list of group assignments for the policy.
func (m *OfficeClientConfiguration) SetAssignments(value []OfficeClientConfigurationAssignmentable)() {
    m.assignments = value
}
// SetCheckinStatuses sets the checkinStatuses property value. List of office Client check-in status.
func (m *OfficeClientConfiguration) SetCheckinStatuses(value []OfficeClientCheckinStatusable)() {
    m.checkinStatuses = value
}
// SetDescription sets the description property value. Not yet documented
func (m *OfficeClientConfiguration) SetDescription(value *string)() {
    m.description = value
}
// SetDisplayName sets the displayName property value. Admin provided description of the office client configuration policy.
func (m *OfficeClientConfiguration) SetDisplayName(value *string)() {
    m.displayName = value
}
// SetPolicyPayload sets the policyPayload property value. Policy settings JSON string in binary format, these values cannot be changed by the user.
func (m *OfficeClientConfiguration) SetPolicyPayload(value []byte)() {
    m.policyPayload = value
}
// SetPriority sets the priority property value. Priority value should be unique value for each policy under a tenant and will be used for conflict resolution, lower values mean priority is high.
func (m *OfficeClientConfiguration) SetPriority(value *int32)() {
    m.priority = value
}
// SetUserCheckinSummary sets the userCheckinSummary property value. User check-in summary for the policy.
func (m *OfficeClientConfiguration) SetUserCheckinSummary(value OfficeUserCheckinSummaryable)() {
    m.userCheckinSummary = value
}
// SetUserPreferencePayload sets the userPreferencePayload property value. Preference settings JSON string in binary format, these values can be overridden by the user.
func (m *OfficeClientConfiguration) SetUserPreferencePayload(value []byte)() {
    m.userPreferencePayload = value
}
