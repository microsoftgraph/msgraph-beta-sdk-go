package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

type MicrosoftTunnelSite struct {
    Entity
    description *string;
    displayName *string;
    microsoftTunnelConfiguration *MicrosoftTunnelConfiguration;
    microsoftTunnelServers []MicrosoftTunnelServer;
    publicAddress *string;
    roleScopeTagIds []string;
    upgradeAutomatically *bool;
    upgradeAvailable *bool;
    upgradeWindowEndTime *string;
    upgradeWindowStartTime *string;
    upgradeWindowUtcOffsetInMinutes *int32;
}
func NewMicrosoftTunnelSite()(*MicrosoftTunnelSite) {
    m := &MicrosoftTunnelSite{
        Entity: *NewEntity(),
    }
    return m
}
func (m *MicrosoftTunnelSite) GetDescription()(*string) {
    if m == nil {
        return nil
    } else {
        return m.description
    }
}
func (m *MicrosoftTunnelSite) GetDisplayName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.displayName
    }
}
func (m *MicrosoftTunnelSite) GetMicrosoftTunnelConfiguration()(*MicrosoftTunnelConfiguration) {
    if m == nil {
        return nil
    } else {
        return m.microsoftTunnelConfiguration
    }
}
func (m *MicrosoftTunnelSite) GetMicrosoftTunnelServers()([]MicrosoftTunnelServer) {
    if m == nil {
        return nil
    } else {
        return m.microsoftTunnelServers
    }
}
func (m *MicrosoftTunnelSite) GetPublicAddress()(*string) {
    if m == nil {
        return nil
    } else {
        return m.publicAddress
    }
}
func (m *MicrosoftTunnelSite) GetRoleScopeTagIds()([]string) {
    if m == nil {
        return nil
    } else {
        return m.roleScopeTagIds
    }
}
func (m *MicrosoftTunnelSite) GetUpgradeAutomatically()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.upgradeAutomatically
    }
}
func (m *MicrosoftTunnelSite) GetUpgradeAvailable()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.upgradeAvailable
    }
}
func (m *MicrosoftTunnelSite) GetUpgradeWindowEndTime()(*string) {
    if m == nil {
        return nil
    } else {
        return m.upgradeWindowEndTime
    }
}
func (m *MicrosoftTunnelSite) GetUpgradeWindowStartTime()(*string) {
    if m == nil {
        return nil
    } else {
        return m.upgradeWindowStartTime
    }
}
func (m *MicrosoftTunnelSite) GetUpgradeWindowUtcOffsetInMinutes()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.upgradeWindowUtcOffsetInMinutes
    }
}
func (m *MicrosoftTunnelSite) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["description"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetDescription(val)
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
    res["microsoftTunnelConfiguration"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewMicrosoftTunnelConfiguration() })
        if err != nil {
            return err
        }
        m.SetMicrosoftTunnelConfiguration(val.(*MicrosoftTunnelConfiguration))
        return nil
    }
    res["microsoftTunnelServers"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewMicrosoftTunnelServer() })
        if err != nil {
            return err
        }
        res := make([]MicrosoftTunnelServer, len(val))
        for i, v := range val {
            res[i] = *(v.(*MicrosoftTunnelServer))
        }
        m.SetMicrosoftTunnelServers(res)
        return nil
    }
    res["publicAddress"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetPublicAddress(val)
        return nil
    }
    res["roleScopeTagIds"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        res := make([]string, len(val))
        for i, v := range val {
            res[i] = v.(string)
        }
        m.SetRoleScopeTagIds(res)
        return nil
    }
    res["upgradeAutomatically"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        m.SetUpgradeAutomatically(val)
        return nil
    }
    res["upgradeAvailable"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        m.SetUpgradeAvailable(val)
        return nil
    }
    res["upgradeWindowEndTime"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetUpgradeWindowEndTime(val)
        return nil
    }
    res["upgradeWindowStartTime"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetUpgradeWindowStartTime(val)
        return nil
    }
    res["upgradeWindowUtcOffsetInMinutes"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        m.SetUpgradeWindowUtcOffsetInMinutes(val)
        return nil
    }
    return res
}
func (m *MicrosoftTunnelSite) IsNil()(bool) {
    return m == nil
}
func (m *MicrosoftTunnelSite) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
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
        err = writer.WriteObjectValue("microsoftTunnelConfiguration", m.GetMicrosoftTunnelConfiguration())
        if err != nil {
            return err
        }
    }
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetMicrosoftTunnelServers()))
        for i, v := range m.GetMicrosoftTunnelServers() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("microsoftTunnelServers", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("publicAddress", m.GetPublicAddress())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteCollectionOfStringValues("roleScopeTagIds", m.GetRoleScopeTagIds())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("upgradeAutomatically", m.GetUpgradeAutomatically())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("upgradeAvailable", m.GetUpgradeAvailable())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("upgradeWindowEndTime", m.GetUpgradeWindowEndTime())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("upgradeWindowStartTime", m.GetUpgradeWindowStartTime())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("upgradeWindowUtcOffsetInMinutes", m.GetUpgradeWindowUtcOffsetInMinutes())
        if err != nil {
            return err
        }
    }
    return nil
}
func (m *MicrosoftTunnelSite) SetDescription(value *string)() {
    m.description = value
}
func (m *MicrosoftTunnelSite) SetDisplayName(value *string)() {
    m.displayName = value
}
func (m *MicrosoftTunnelSite) SetMicrosoftTunnelConfiguration(value *MicrosoftTunnelConfiguration)() {
    m.microsoftTunnelConfiguration = value
}
func (m *MicrosoftTunnelSite) SetMicrosoftTunnelServers(value []MicrosoftTunnelServer)() {
    m.microsoftTunnelServers = value
}
func (m *MicrosoftTunnelSite) SetPublicAddress(value *string)() {
    m.publicAddress = value
}
func (m *MicrosoftTunnelSite) SetRoleScopeTagIds(value []string)() {
    m.roleScopeTagIds = value
}
func (m *MicrosoftTunnelSite) SetUpgradeAutomatically(value *bool)() {
    m.upgradeAutomatically = value
}
func (m *MicrosoftTunnelSite) SetUpgradeAvailable(value *bool)() {
    m.upgradeAvailable = value
}
func (m *MicrosoftTunnelSite) SetUpgradeWindowEndTime(value *string)() {
    m.upgradeWindowEndTime = value
}
func (m *MicrosoftTunnelSite) SetUpgradeWindowStartTime(value *string)() {
    m.upgradeWindowStartTime = value
}
func (m *MicrosoftTunnelSite) SetUpgradeWindowUtcOffsetInMinutes(value *int32)() {
    m.upgradeWindowUtcOffsetInMinutes = value
}
