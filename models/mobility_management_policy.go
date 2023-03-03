package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// MobilityManagementPolicy 
type MobilityManagementPolicy struct {
    Entity
}
// NewMobilityManagementPolicy instantiates a new mobilityManagementPolicy and sets the default values.
func NewMobilityManagementPolicy()(*MobilityManagementPolicy) {
    m := &MobilityManagementPolicy{
        Entity: *NewEntity(),
    }
    return m
}
// CreateMobilityManagementPolicyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateMobilityManagementPolicyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewMobilityManagementPolicy(), nil
}
// GetAppliesTo gets the appliesTo property value. Indicates the user scope of the mobility management policy. Possible values are: none, all, selected.
func (m *MobilityManagementPolicy) GetAppliesTo()(*PolicyScope) {
    val, err := m.GetBackingStore().Get("appliesTo")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*PolicyScope)
    }
    return nil
}
// GetComplianceUrl gets the complianceUrl property value. Compliance URL of the mobility management application.
func (m *MobilityManagementPolicy) GetComplianceUrl()(*string) {
    val, err := m.GetBackingStore().Get("complianceUrl")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetDescription gets the description property value. Description of the mobility management application.
func (m *MobilityManagementPolicy) GetDescription()(*string) {
    val, err := m.GetBackingStore().Get("description")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetDiscoveryUrl gets the discoveryUrl property value. Discovery URL of the mobility management application.
func (m *MobilityManagementPolicy) GetDiscoveryUrl()(*string) {
    val, err := m.GetBackingStore().Get("discoveryUrl")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetDisplayName gets the displayName property value. Display name of the mobility management application.
func (m *MobilityManagementPolicy) GetDisplayName()(*string) {
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
func (m *MobilityManagementPolicy) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["appliesTo"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParsePolicyScope)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAppliesTo(val.(*PolicyScope))
        }
        return nil
    }
    res["complianceUrl"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetComplianceUrl(val)
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
    res["discoveryUrl"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDiscoveryUrl(val)
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
    res["includedGroups"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateGroupFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]Groupable, len(val))
            for i, v := range val {
                res[i] = v.(Groupable)
            }
            m.SetIncludedGroups(res)
        }
        return nil
    }
    res["isValid"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsValid(val)
        }
        return nil
    }
    res["termsOfUseUrl"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
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
// GetIncludedGroups gets the includedGroups property value. Azure AD groups under the scope of the mobility management application if appliesTo is selected
func (m *MobilityManagementPolicy) GetIncludedGroups()([]Groupable) {
    val, err := m.GetBackingStore().Get("includedGroups")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]Groupable)
    }
    return nil
}
// GetIsValid gets the isValid property value. Whether policy is valid. Invalid policies may not be updated and should be deleted.
func (m *MobilityManagementPolicy) GetIsValid()(*bool) {
    val, err := m.GetBackingStore().Get("isValid")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// GetTermsOfUseUrl gets the termsOfUseUrl property value. Terms of Use URL of the mobility management application.
func (m *MobilityManagementPolicy) GetTermsOfUseUrl()(*string) {
    val, err := m.GetBackingStore().Get("termsOfUseUrl")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// Serialize serializes information the current object
func (m *MobilityManagementPolicy) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
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
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetIncludedGroups()))
        for i, v := range m.GetIncludedGroups() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
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
    err := m.GetBackingStore().Set("appliesTo", value)
    if err != nil {
        panic(err)
    }
}
// SetComplianceUrl sets the complianceUrl property value. Compliance URL of the mobility management application.
func (m *MobilityManagementPolicy) SetComplianceUrl(value *string)() {
    err := m.GetBackingStore().Set("complianceUrl", value)
    if err != nil {
        panic(err)
    }
}
// SetDescription sets the description property value. Description of the mobility management application.
func (m *MobilityManagementPolicy) SetDescription(value *string)() {
    err := m.GetBackingStore().Set("description", value)
    if err != nil {
        panic(err)
    }
}
// SetDiscoveryUrl sets the discoveryUrl property value. Discovery URL of the mobility management application.
func (m *MobilityManagementPolicy) SetDiscoveryUrl(value *string)() {
    err := m.GetBackingStore().Set("discoveryUrl", value)
    if err != nil {
        panic(err)
    }
}
// SetDisplayName sets the displayName property value. Display name of the mobility management application.
func (m *MobilityManagementPolicy) SetDisplayName(value *string)() {
    err := m.GetBackingStore().Set("displayName", value)
    if err != nil {
        panic(err)
    }
}
// SetIncludedGroups sets the includedGroups property value. Azure AD groups under the scope of the mobility management application if appliesTo is selected
func (m *MobilityManagementPolicy) SetIncludedGroups(value []Groupable)() {
    err := m.GetBackingStore().Set("includedGroups", value)
    if err != nil {
        panic(err)
    }
}
// SetIsValid sets the isValid property value. Whether policy is valid. Invalid policies may not be updated and should be deleted.
func (m *MobilityManagementPolicy) SetIsValid(value *bool)() {
    err := m.GetBackingStore().Set("isValid", value)
    if err != nil {
        panic(err)
    }
}
// SetTermsOfUseUrl sets the termsOfUseUrl property value. Terms of Use URL of the mobility management application.
func (m *MobilityManagementPolicy) SetTermsOfUseUrl(value *string)() {
    err := m.GetBackingStore().Set("termsOfUseUrl", value)
    if err != nil {
        panic(err)
    }
}
// MobilityManagementPolicyable 
type MobilityManagementPolicyable interface {
    Entityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAppliesTo()(*PolicyScope)
    GetComplianceUrl()(*string)
    GetDescription()(*string)
    GetDiscoveryUrl()(*string)
    GetDisplayName()(*string)
    GetIncludedGroups()([]Groupable)
    GetIsValid()(*bool)
    GetTermsOfUseUrl()(*string)
    SetAppliesTo(value *PolicyScope)()
    SetComplianceUrl(value *string)()
    SetDescription(value *string)()
    SetDiscoveryUrl(value *string)()
    SetDisplayName(value *string)()
    SetIncludedGroups(value []Groupable)()
    SetIsValid(value *bool)()
    SetTermsOfUseUrl(value *string)()
}
