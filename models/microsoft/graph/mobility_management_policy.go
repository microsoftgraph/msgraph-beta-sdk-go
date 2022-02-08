package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// MobilityManagementPolicy 
type MobilityManagementPolicy struct {
    Entity
    // Indicates the user scope of the mobility management policy. Possible values are: none, all, selected.
    appliesTo *PolicyScope;
    // Compliance URL of the mobility management application.
    complianceUrl *string;
    // Description of the mobility management application.
    description *string;
    // Discovery URL of the mobility management application.
    discoveryUrl *string;
    // Display name of the mobility management application.
    displayName *string;
    // Azure AD groups under the scope of the mobility management application if appliesTo is selected
    includedGroups []Group;
    // Whether policy is valid. Invalid policies may not be updated and should be deleted.
    isValid *bool;
    // Terms of Use URL of the mobility management application.
    termsOfUseUrl *string;
}
// NewMobilityManagementPolicy instantiates a new mobilityManagementPolicy and sets the default values.
func NewMobilityManagementPolicy()(*MobilityManagementPolicy) {
    m := &MobilityManagementPolicy{
        Entity: *NewEntity(),
    }
    return m
}
// GetAppliesTo gets the appliesTo property value. Indicates the user scope of the mobility management policy. Possible values are: none, all, selected.
func (m *MobilityManagementPolicy) GetAppliesTo()(*PolicyScope) {
    if m == nil {
        return nil
    } else {
        return m.appliesTo
    }
}
// GetComplianceUrl gets the complianceUrl property value. Compliance URL of the mobility management application.
func (m *MobilityManagementPolicy) GetComplianceUrl()(*string) {
    if m == nil {
        return nil
    } else {
        return m.complianceUrl
    }
}
// GetDescription gets the description property value. Description of the mobility management application.
func (m *MobilityManagementPolicy) GetDescription()(*string) {
    if m == nil {
        return nil
    } else {
        return m.description
    }
}
// GetDiscoveryUrl gets the discoveryUrl property value. Discovery URL of the mobility management application.
func (m *MobilityManagementPolicy) GetDiscoveryUrl()(*string) {
    if m == nil {
        return nil
    } else {
        return m.discoveryUrl
    }
}
// GetDisplayName gets the displayName property value. Display name of the mobility management application.
func (m *MobilityManagementPolicy) GetDisplayName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.displayName
    }
}
// GetIncludedGroups gets the includedGroups property value. Azure AD groups under the scope of the mobility management application if appliesTo is selected
func (m *MobilityManagementPolicy) GetIncludedGroups()([]Group) {
    if m == nil {
        return nil
    } else {
        return m.includedGroups
    }
}
// GetIsValid gets the isValid property value. Whether policy is valid. Invalid policies may not be updated and should be deleted.
func (m *MobilityManagementPolicy) GetIsValid()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.isValid
    }
}
// GetTermsOfUseUrl gets the termsOfUseUrl property value. Terms of Use URL of the mobility management application.
func (m *MobilityManagementPolicy) GetTermsOfUseUrl()(*string) {
    if m == nil {
        return nil
    } else {
        return m.termsOfUseUrl
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *MobilityManagementPolicy) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["appliesTo"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParsePolicyScope)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAppliesTo(val.(*PolicyScope))
        }
        return nil
    }
    res["complianceUrl"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetComplianceUrl(val)
        }
        return nil
    }
    res["description"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDescription(val)
        }
        return nil
    }
    res["discoveryUrl"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDiscoveryUrl(val)
        }
        return nil
    }
    res["displayName"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDisplayName(val)
        }
        return nil
    }
    res["includedGroups"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewGroup() })
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]Group, len(val))
            for i, v := range val {
                res[i] = *(v.(*Group))
            }
            m.SetIncludedGroups(res)
        }
        return nil
    }
    res["isValid"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsValid(val)
        }
        return nil
    }
    res["termsOfUseUrl"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTermsOfUseUrl(val)
        }
        return nil
    }
    return res
}
func (m *MobilityManagementPolicy) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
func (m *MobilityManagementPolicy) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetAppliesTo() != nil {
        cast := (*m.GetAppliesTo()).String()
        err = writer.WriteStringValue("appliesTo", &cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("complianceUrl", m.GetComplianceUrl())
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
        err = writer.WriteStringValue("discoveryUrl", m.GetDiscoveryUrl())
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
    if m.GetIncludedGroups() != nil {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetIncludedGroups()))
        for i, v := range m.GetIncludedGroups() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("includedGroups", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("isValid", m.GetIsValid())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("termsOfUseUrl", m.GetTermsOfUseUrl())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAppliesTo sets the appliesTo property value. Indicates the user scope of the mobility management policy. Possible values are: none, all, selected.
func (m *MobilityManagementPolicy) SetAppliesTo(value *PolicyScope)() {
    if m != nil {
        m.appliesTo = value
    }
}
// SetComplianceUrl sets the complianceUrl property value. Compliance URL of the mobility management application.
func (m *MobilityManagementPolicy) SetComplianceUrl(value *string)() {
    if m != nil {
        m.complianceUrl = value
    }
}
// SetDescription sets the description property value. Description of the mobility management application.
func (m *MobilityManagementPolicy) SetDescription(value *string)() {
    if m != nil {
        m.description = value
    }
}
// SetDiscoveryUrl sets the discoveryUrl property value. Discovery URL of the mobility management application.
func (m *MobilityManagementPolicy) SetDiscoveryUrl(value *string)() {
    if m != nil {
        m.discoveryUrl = value
    }
}
// SetDisplayName sets the displayName property value. Display name of the mobility management application.
func (m *MobilityManagementPolicy) SetDisplayName(value *string)() {
    if m != nil {
        m.displayName = value
    }
}
// SetIncludedGroups sets the includedGroups property value. Azure AD groups under the scope of the mobility management application if appliesTo is selected
func (m *MobilityManagementPolicy) SetIncludedGroups(value []Group)() {
    if m != nil {
        m.includedGroups = value
    }
}
// SetIsValid sets the isValid property value. Whether policy is valid. Invalid policies may not be updated and should be deleted.
func (m *MobilityManagementPolicy) SetIsValid(value *bool)() {
    if m != nil {
        m.isValid = value
    }
}
// SetTermsOfUseUrl sets the termsOfUseUrl property value. Terms of Use URL of the mobility management application.
func (m *MobilityManagementPolicy) SetTermsOfUseUrl(value *string)() {
    if m != nil {
        m.termsOfUseUrl = value
    }
}
