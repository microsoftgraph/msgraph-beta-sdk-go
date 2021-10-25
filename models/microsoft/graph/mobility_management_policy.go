package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

type MobilityManagementPolicy struct {
    Entity
    appliesTo *PolicyScope;
    complianceUrl *string;
    description *string;
    discoveryUrl *string;
    displayName *string;
    includedGroups []Group;
    isValid *bool;
    termsOfUseUrl *string;
}
func NewMobilityManagementPolicy()(*MobilityManagementPolicy) {
    m := &MobilityManagementPolicy{
        Entity: *NewEntity(),
    }
    return m
}
func (m *MobilityManagementPolicy) GetAppliesTo()(*PolicyScope) {
    if m == nil {
        return nil
    } else {
        return m.appliesTo
    }
}
func (m *MobilityManagementPolicy) GetComplianceUrl()(*string) {
    if m == nil {
        return nil
    } else {
        return m.complianceUrl
    }
}
func (m *MobilityManagementPolicy) GetDescription()(*string) {
    if m == nil {
        return nil
    } else {
        return m.description
    }
}
func (m *MobilityManagementPolicy) GetDiscoveryUrl()(*string) {
    if m == nil {
        return nil
    } else {
        return m.discoveryUrl
    }
}
func (m *MobilityManagementPolicy) GetDisplayName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.displayName
    }
}
func (m *MobilityManagementPolicy) GetIncludedGroups()([]Group) {
    if m == nil {
        return nil
    } else {
        return m.includedGroups
    }
}
func (m *MobilityManagementPolicy) GetIsValid()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.isValid
    }
}
func (m *MobilityManagementPolicy) GetTermsOfUseUrl()(*string) {
    if m == nil {
        return nil
    } else {
        return m.termsOfUseUrl
    }
}
func (m *MobilityManagementPolicy) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["appliesTo"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParsePolicyScope)
        if err != nil {
            return err
        }
        cast := val.(PolicyScope)
        m.SetAppliesTo(&cast)
        return nil
    }
    res["complianceUrl"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetComplianceUrl(val)
        return nil
    }
    res["description"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetDescription(val)
        return nil
    }
    res["discoveryUrl"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetDiscoveryUrl(val)
        return nil
    }
    res["displayName"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetDisplayName(val)
        return nil
    }
    res["includedGroups"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewGroup() })
        if err != nil {
            return err
        }
        res := make([]Group, len(val))
        for i, v := range val {
            res[i] = *(v.(*Group))
        }
        m.SetIncludedGroups(res)
        return nil
    }
    res["isValid"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        m.SetIsValid(val)
        return nil
    }
    res["termsOfUseUrl"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetTermsOfUseUrl(val)
        return nil
    }
    return res
}
func (m *MobilityManagementPolicy) IsNil()(bool) {
    return m == nil
}
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
func (m *MobilityManagementPolicy) SetAppliesTo(value *PolicyScope)() {
    m.appliesTo = value
}
func (m *MobilityManagementPolicy) SetComplianceUrl(value *string)() {
    m.complianceUrl = value
}
func (m *MobilityManagementPolicy) SetDescription(value *string)() {
    m.description = value
}
func (m *MobilityManagementPolicy) SetDiscoveryUrl(value *string)() {
    m.discoveryUrl = value
}
func (m *MobilityManagementPolicy) SetDisplayName(value *string)() {
    m.displayName = value
}
func (m *MobilityManagementPolicy) SetIncludedGroups(value []Group)() {
    m.includedGroups = value
}
func (m *MobilityManagementPolicy) SetIsValid(value *bool)() {
    m.isValid = value
}
func (m *MobilityManagementPolicy) SetTermsOfUseUrl(value *string)() {
    m.termsOfUseUrl = value
}
