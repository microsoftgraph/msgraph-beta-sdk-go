package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// 
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
// Instantiates a new mobilityManagementPolicy and sets the default values.
func NewMobilityManagementPolicy()(*MobilityManagementPolicy) {
    m := &MobilityManagementPolicy{
        Entity: *NewEntity(),
    }
    return m
}
// Gets the appliesTo property value. Indicates the user scope of the mobility management policy. Possible values are: none, all, selected.
func (m *MobilityManagementPolicy) GetAppliesTo()(*PolicyScope) {
    if m == nil {
        return nil
    } else {
        return m.appliesTo
    }
}
// Gets the complianceUrl property value. Compliance URL of the mobility management application.
func (m *MobilityManagementPolicy) GetComplianceUrl()(*string) {
    if m == nil {
        return nil
    } else {
        return m.complianceUrl
    }
}
// Gets the description property value. Description of the mobility management application.
func (m *MobilityManagementPolicy) GetDescription()(*string) {
    if m == nil {
        return nil
    } else {
        return m.description
    }
}
// Gets the discoveryUrl property value. Discovery URL of the mobility management application.
func (m *MobilityManagementPolicy) GetDiscoveryUrl()(*string) {
    if m == nil {
        return nil
    } else {
        return m.discoveryUrl
    }
}
// Gets the displayName property value. Display name of the mobility management application.
func (m *MobilityManagementPolicy) GetDisplayName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.displayName
    }
}
// Gets the includedGroups property value. Azure AD groups under the scope of the mobility management application if appliesTo is selected
func (m *MobilityManagementPolicy) GetIncludedGroups()([]Group) {
    if m == nil {
        return nil
    } else {
        return m.includedGroups
    }
}
// Gets the isValid property value. Whether policy is valid. Invalid policies may not be updated and should be deleted.
func (m *MobilityManagementPolicy) GetIsValid()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.isValid
    }
}
// Gets the termsOfUseUrl property value. Terms of Use URL of the mobility management application.
func (m *MobilityManagementPolicy) GetTermsOfUseUrl()(*string) {
    if m == nil {
        return nil
    } else {
        return m.termsOfUseUrl
    }
}
// The deserialization information for the current model
func (m *MobilityManagementPolicy) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["appliesTo"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParsePolicyScope)
        if err != nil {
            return err
        }
        if val != nil {
            cast := val.(PolicyScope)
            m.SetAppliesTo(&cast)
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
// Serializes information the current object
// Parameters:
//  - writer : Serialization writer to use to serialize this model
func (m *MobilityManagementPolicy) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetAppliesTo() != nil {
        cast := m.GetAppliesTo().String()
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
    {
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
// Sets the appliesTo property value. Indicates the user scope of the mobility management policy. Possible values are: none, all, selected.
// Parameters:
//  - value : Value to set for the appliesTo property.
func (m *MobilityManagementPolicy) SetAppliesTo(value *PolicyScope)() {
    m.appliesTo = value
}
// Sets the complianceUrl property value. Compliance URL of the mobility management application.
// Parameters:
//  - value : Value to set for the complianceUrl property.
func (m *MobilityManagementPolicy) SetComplianceUrl(value *string)() {
    m.complianceUrl = value
}
// Sets the description property value. Description of the mobility management application.
// Parameters:
//  - value : Value to set for the description property.
func (m *MobilityManagementPolicy) SetDescription(value *string)() {
    m.description = value
}
// Sets the discoveryUrl property value. Discovery URL of the mobility management application.
// Parameters:
//  - value : Value to set for the discoveryUrl property.
func (m *MobilityManagementPolicy) SetDiscoveryUrl(value *string)() {
    m.discoveryUrl = value
}
// Sets the displayName property value. Display name of the mobility management application.
// Parameters:
//  - value : Value to set for the displayName property.
func (m *MobilityManagementPolicy) SetDisplayName(value *string)() {
    m.displayName = value
}
// Sets the includedGroups property value. Azure AD groups under the scope of the mobility management application if appliesTo is selected
// Parameters:
//  - value : Value to set for the includedGroups property.
func (m *MobilityManagementPolicy) SetIncludedGroups(value []Group)() {
    m.includedGroups = value
}
// Sets the isValid property value. Whether policy is valid. Invalid policies may not be updated and should be deleted.
// Parameters:
//  - value : Value to set for the isValid property.
func (m *MobilityManagementPolicy) SetIsValid(value *bool)() {
    m.isValid = value
}
// Sets the termsOfUseUrl property value. Terms of Use URL of the mobility management application.
// Parameters:
//  - value : Value to set for the termsOfUseUrl property.
func (m *MobilityManagementPolicy) SetTermsOfUseUrl(value *string)() {
    m.termsOfUseUrl = value
}
