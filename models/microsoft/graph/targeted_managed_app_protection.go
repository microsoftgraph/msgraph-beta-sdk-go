package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// 
type TargetedManagedAppProtection struct {
    ManagedAppProtection
    // Public Apps selection: group or individual. Possible values are: selectedPublicApps, allCoreMicrosoftApps, allMicrosoftApps, allApps.
    appGroupType *TargetedManagedAppGroupType;
    // Navigation property to list of inclusion and exclusion groups to which the policy is deployed.
    assignments []TargetedManagedAppPolicyAssignment;
    // Indicates if the policy is deployed to any inclusion groups or not.
    isAssigned *bool;
    // The intended app management levels for this policy. Possible values are: unspecified, unmanaged, mdm, androidEnterprise.
    targetedAppManagementLevels *AppManagementLevel;
}
// Instantiates a new targetedManagedAppProtection and sets the default values.
func NewTargetedManagedAppProtection()(*TargetedManagedAppProtection) {
    m := &TargetedManagedAppProtection{
        ManagedAppProtection: *NewManagedAppProtection(),
    }
    return m
}
// Gets the appGroupType property value. Public Apps selection: group or individual. Possible values are: selectedPublicApps, allCoreMicrosoftApps, allMicrosoftApps, allApps.
func (m *TargetedManagedAppProtection) GetAppGroupType()(*TargetedManagedAppGroupType) {
    if m == nil {
        return nil
    } else {
        return m.appGroupType
    }
}
// Gets the assignments property value. Navigation property to list of inclusion and exclusion groups to which the policy is deployed.
func (m *TargetedManagedAppProtection) GetAssignments()([]TargetedManagedAppPolicyAssignment) {
    if m == nil {
        return nil
    } else {
        return m.assignments
    }
}
// Gets the isAssigned property value. Indicates if the policy is deployed to any inclusion groups or not.
func (m *TargetedManagedAppProtection) GetIsAssigned()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.isAssigned
    }
}
// Gets the targetedAppManagementLevels property value. The intended app management levels for this policy. Possible values are: unspecified, unmanaged, mdm, androidEnterprise.
func (m *TargetedManagedAppProtection) GetTargetedAppManagementLevels()(*AppManagementLevel) {
    if m == nil {
        return nil
    } else {
        return m.targetedAppManagementLevels
    }
}
// The deserialization information for the current model
func (m *TargetedManagedAppProtection) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.ManagedAppProtection.GetFieldDeserializers()
    res["appGroupType"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParseTargetedManagedAppGroupType)
        if err != nil {
            return err
        }
        cast := val.(TargetedManagedAppGroupType)
        m.SetAppGroupType(&cast)
        return nil
    }
    res["assignments"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewTargetedManagedAppPolicyAssignment() })
        if err != nil {
            return err
        }
        res := make([]TargetedManagedAppPolicyAssignment, len(val))
        for i, v := range val {
            res[i] = *(v.(*TargetedManagedAppPolicyAssignment))
        }
        m.SetAssignments(res)
        return nil
    }
    res["isAssigned"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        m.SetIsAssigned(val)
        return nil
    }
    res["targetedAppManagementLevels"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParseAppManagementLevel)
        if err != nil {
            return err
        }
        cast := val.(AppManagementLevel)
        m.SetTargetedAppManagementLevels(&cast)
        return nil
    }
    return res
}
func (m *TargetedManagedAppProtection) IsNil()(bool) {
    return m == nil
}
// Serializes information the current object
// Parameters:
//  - writer : Serialization writer to use to serialize this model
func (m *TargetedManagedAppProtection) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.ManagedAppProtection.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetAppGroupType() != nil {
        cast := m.GetAppGroupType().String()
        err = writer.WriteStringValue("appGroupType", &cast)
        if err != nil {
            return err
        }
    }
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetAssignments()))
        for i, v := range m.GetAssignments() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
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
        cast := m.GetTargetedAppManagementLevels().String()
        err = writer.WriteStringValue("targetedAppManagementLevels", &cast)
        if err != nil {
            return err
        }
    }
    return nil
}
// Sets the appGroupType property value. Public Apps selection: group or individual. Possible values are: selectedPublicApps, allCoreMicrosoftApps, allMicrosoftApps, allApps.
// Parameters:
//  - value : Value to set for the appGroupType property.
func (m *TargetedManagedAppProtection) SetAppGroupType(value *TargetedManagedAppGroupType)() {
    m.appGroupType = value
}
// Sets the assignments property value. Navigation property to list of inclusion and exclusion groups to which the policy is deployed.
// Parameters:
//  - value : Value to set for the assignments property.
func (m *TargetedManagedAppProtection) SetAssignments(value []TargetedManagedAppPolicyAssignment)() {
    m.assignments = value
}
// Sets the isAssigned property value. Indicates if the policy is deployed to any inclusion groups or not.
// Parameters:
//  - value : Value to set for the isAssigned property.
func (m *TargetedManagedAppProtection) SetIsAssigned(value *bool)() {
    m.isAssigned = value
}
// Sets the targetedAppManagementLevels property value. The intended app management levels for this policy. Possible values are: unspecified, unmanaged, mdm, androidEnterprise.
// Parameters:
//  - value : Value to set for the targetedAppManagementLevels property.
func (m *TargetedManagedAppProtection) SetTargetedAppManagementLevels(value *AppManagementLevel)() {
    m.targetedAppManagementLevels = value
}
