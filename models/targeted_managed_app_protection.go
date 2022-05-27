package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// TargetedManagedAppProtection policy used to configure detailed management settings targeted to specific security groups
type TargetedManagedAppProtection struct {
    ManagedAppProtection
    // Public Apps selection: group or individual. Possible values are: selectedPublicApps, allCoreMicrosoftApps, allMicrosoftApps, allApps.
    appGroupType *TargetedManagedAppGroupType
    // Navigation property to list of inclusion and exclusion groups to which the policy is deployed.
    assignments []TargetedManagedAppPolicyAssignmentable
    // Indicates if the policy is deployed to any inclusion groups or not.
    isAssigned *bool
    // The intended app management levels for this policy. Possible values are: unspecified, unmanaged, mdm, androidEnterprise.
    targetedAppManagementLevels *AppManagementLevel
}
// NewTargetedManagedAppProtection instantiates a new targetedManagedAppProtection and sets the default values.
func NewTargetedManagedAppProtection()(*TargetedManagedAppProtection) {
    m := &TargetedManagedAppProtection{
        ManagedAppProtection: *NewManagedAppProtection(),
    }
    return m
}
// CreateTargetedManagedAppProtectionFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateTargetedManagedAppProtectionFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
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
                mappingStr := *mappingValue
                switch mappingStr {
                    case "#microsoft.graph.targetedManagedAppProtection":
                        return NewTargetedManagedAppProtection(), nil
                }
            }
        }
    }
    return NewTargetedManagedAppProtection(), nil
}
// GetAppGroupType gets the appGroupType property value. Public Apps selection: group or individual. Possible values are: selectedPublicApps, allCoreMicrosoftApps, allMicrosoftApps, allApps.
func (m *TargetedManagedAppProtection) GetAppGroupType()(*TargetedManagedAppGroupType) {
    if m == nil {
        return nil
    } else {
        return m.appGroupType
    }
}
// GetAssignments gets the assignments property value. Navigation property to list of inclusion and exclusion groups to which the policy is deployed.
func (m *TargetedManagedAppProtection) GetAssignments()([]TargetedManagedAppPolicyAssignmentable) {
    if m == nil {
        return nil
    } else {
        return m.assignments
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *TargetedManagedAppProtection) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.ManagedAppProtection.GetFieldDeserializers()
    res["appGroupType"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseTargetedManagedAppGroupType)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAppGroupType(val.(*TargetedManagedAppGroupType))
        }
        return nil
    }
    res["assignments"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateTargetedManagedAppPolicyAssignmentFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]TargetedManagedAppPolicyAssignmentable, len(val))
            for i, v := range val {
                res[i] = v.(TargetedManagedAppPolicyAssignmentable)
            }
            m.SetAssignments(res)
        }
        return nil
    }
    res["isAssigned"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsAssigned(val)
        }
        return nil
    }
    res["targetedAppManagementLevels"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseAppManagementLevel)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTargetedAppManagementLevels(val.(*AppManagementLevel))
        }
        return nil
    }
    return res
}
// GetIsAssigned gets the isAssigned property value. Indicates if the policy is deployed to any inclusion groups or not.
func (m *TargetedManagedAppProtection) GetIsAssigned()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.isAssigned
    }
}
// GetTargetedAppManagementLevels gets the targetedAppManagementLevels property value. The intended app management levels for this policy. Possible values are: unspecified, unmanaged, mdm, androidEnterprise.
func (m *TargetedManagedAppProtection) GetTargetedAppManagementLevels()(*AppManagementLevel) {
    if m == nil {
        return nil
    } else {
        return m.targetedAppManagementLevels
    }
}
// Serialize serializes information the current object
func (m *TargetedManagedAppProtection) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.ManagedAppProtection.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetAppGroupType() != nil {
        cast := (*m.GetAppGroupType()).String()
        err = writer.WriteStringValue("appGroupType", &cast)
        if err != nil {
            return err
        }
    }
    if m.GetAssignments() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetAssignments()))
        for i, v := range m.GetAssignments() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("assignments", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("isAssigned", m.GetIsAssigned())
        if err != nil {
            return err
        }
    }
    if m.GetTargetedAppManagementLevels() != nil {
        cast := (*m.GetTargetedAppManagementLevels()).String()
        err = writer.WriteStringValue("targetedAppManagementLevels", &cast)
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAppGroupType sets the appGroupType property value. Public Apps selection: group or individual. Possible values are: selectedPublicApps, allCoreMicrosoftApps, allMicrosoftApps, allApps.
func (m *TargetedManagedAppProtection) SetAppGroupType(value *TargetedManagedAppGroupType)() {
    if m != nil {
        m.appGroupType = value
    }
}
// SetAssignments sets the assignments property value. Navigation property to list of inclusion and exclusion groups to which the policy is deployed.
func (m *TargetedManagedAppProtection) SetAssignments(value []TargetedManagedAppPolicyAssignmentable)() {
    if m != nil {
        m.assignments = value
    }
}
// SetIsAssigned sets the isAssigned property value. Indicates if the policy is deployed to any inclusion groups or not.
func (m *TargetedManagedAppProtection) SetIsAssigned(value *bool)() {
    if m != nil {
        m.isAssigned = value
    }
}
// SetTargetedAppManagementLevels sets the targetedAppManagementLevels property value. The intended app management levels for this policy. Possible values are: unspecified, unmanaged, mdm, androidEnterprise.
func (m *TargetedManagedAppProtection) SetTargetedAppManagementLevels(value *AppManagementLevel)() {
    if m != nil {
        m.targetedAppManagementLevels = value
    }
}
