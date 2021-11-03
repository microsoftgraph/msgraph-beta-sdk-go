package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// 
type MicrosoftTunnelSite struct {
    Entity
    // The MicrosoftTunnelSite's description
    description *string;
    // The MicrosoftTunnelSite's display name
    displayName *string;
    // The MicrosoftTunnelSite's Internal Network Access Probe URL
    internalNetworkProbeUrl *string;
    // The MicrosoftTunnelConfiguration that has been applied to this MicrosoftTunnelSite
    microsoftTunnelConfiguration *MicrosoftTunnelConfiguration;
    // A list of MicrosoftTunnelServers that are registered to this MicrosoftTunnelSite
    microsoftTunnelServers []MicrosoftTunnelServer;
    // The MicrosoftTunnelSite's public domain name or IP address
    publicAddress *string;
    // List of Scope Tags for this Entity instance.
    roleScopeTagIds []string;
    // The site's automatic upgrade setting. True for automatic upgrades, false for manual control
    upgradeAutomatically *bool;
    // True if an upgrade is available
    upgradeAvailable *bool;
    // The site's upgrade window end time of day
    upgradeWindowEndTime *string;
    // The site's upgrade window start time of day
    upgradeWindowStartTime *string;
    // The site's timezone represented as a minute offset from UTC
    upgradeWindowUtcOffsetInMinutes *int32;
}
// Instantiates a new microsoftTunnelSite and sets the default values.
func NewMicrosoftTunnelSite()(*MicrosoftTunnelSite) {
    m := &MicrosoftTunnelSite{
        Entity: *NewEntity(),
    }
    return m
}
// Gets the description property value. The MicrosoftTunnelSite's description
func (m *MicrosoftTunnelSite) GetDescription()(*string) {
    if m == nil {
        return nil
    } else {
        return m.description
    }
}
// Gets the displayName property value. The MicrosoftTunnelSite's display name
func (m *MicrosoftTunnelSite) GetDisplayName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.displayName
    }
}
// Gets the internalNetworkProbeUrl property value. The MicrosoftTunnelSite's Internal Network Access Probe URL
func (m *MicrosoftTunnelSite) GetInternalNetworkProbeUrl()(*string) {
    if m == nil {
        return nil
    } else {
        return m.internalNetworkProbeUrl
    }
}
// Gets the microsoftTunnelConfiguration property value. The MicrosoftTunnelConfiguration that has been applied to this MicrosoftTunnelSite
func (m *MicrosoftTunnelSite) GetMicrosoftTunnelConfiguration()(*MicrosoftTunnelConfiguration) {
    if m == nil {
        return nil
    } else {
        return m.microsoftTunnelConfiguration
    }
}
// Gets the microsoftTunnelServers property value. A list of MicrosoftTunnelServers that are registered to this MicrosoftTunnelSite
func (m *MicrosoftTunnelSite) GetMicrosoftTunnelServers()([]MicrosoftTunnelServer) {
    if m == nil {
        return nil
    } else {
        return m.microsoftTunnelServers
    }
}
// Gets the publicAddress property value. The MicrosoftTunnelSite's public domain name or IP address
func (m *MicrosoftTunnelSite) GetPublicAddress()(*string) {
    if m == nil {
        return nil
    } else {
        return m.publicAddress
    }
}
// Gets the roleScopeTagIds property value. List of Scope Tags for this Entity instance.
func (m *MicrosoftTunnelSite) GetRoleScopeTagIds()([]string) {
    if m == nil {
        return nil
    } else {
        return m.roleScopeTagIds
    }
}
// Gets the upgradeAutomatically property value. The site's automatic upgrade setting. True for automatic upgrades, false for manual control
func (m *MicrosoftTunnelSite) GetUpgradeAutomatically()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.upgradeAutomatically
    }
}
// Gets the upgradeAvailable property value. True if an upgrade is available
func (m *MicrosoftTunnelSite) GetUpgradeAvailable()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.upgradeAvailable
    }
}
// Gets the upgradeWindowEndTime property value. The site's upgrade window end time of day
func (m *MicrosoftTunnelSite) GetUpgradeWindowEndTime()(*string) {
    if m == nil {
        return nil
    } else {
        return m.upgradeWindowEndTime
    }
}
// Gets the upgradeWindowStartTime property value. The site's upgrade window start time of day
func (m *MicrosoftTunnelSite) GetUpgradeWindowStartTime()(*string) {
    if m == nil {
        return nil
    } else {
        return m.upgradeWindowStartTime
    }
}
// Gets the upgradeWindowUtcOffsetInMinutes property value. The site's timezone represented as a minute offset from UTC
func (m *MicrosoftTunnelSite) GetUpgradeWindowUtcOffsetInMinutes()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.upgradeWindowUtcOffsetInMinutes
    }
}
// The deserialization information for the current model
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
    res["internalNetworkProbeUrl"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetInternalNetworkProbeUrl(val)
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
            res[i] = *(v.(*string))
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
// Serializes information the current object
// Parameters:
//  - writer : Serialization writer to use to serialize this model
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
        err = writer.WriteStringValue("internalNetworkProbeUrl", m.GetInternalNetworkProbeUrl())
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
// Sets the description property value. The MicrosoftTunnelSite's description
// Parameters:
//  - value : Value to set for the description property.
func (m *MicrosoftTunnelSite) SetDescription(value *string)() {
    m.description = value
}
// Sets the displayName property value. The MicrosoftTunnelSite's display name
// Parameters:
//  - value : Value to set for the displayName property.
func (m *MicrosoftTunnelSite) SetDisplayName(value *string)() {
    m.displayName = value
}
// Sets the internalNetworkProbeUrl property value. The MicrosoftTunnelSite's Internal Network Access Probe URL
// Parameters:
//  - value : Value to set for the internalNetworkProbeUrl property.
func (m *MicrosoftTunnelSite) SetInternalNetworkProbeUrl(value *string)() {
    m.internalNetworkProbeUrl = value
}
// Sets the microsoftTunnelConfiguration property value. The MicrosoftTunnelConfiguration that has been applied to this MicrosoftTunnelSite
// Parameters:
//  - value : Value to set for the microsoftTunnelConfiguration property.
func (m *MicrosoftTunnelSite) SetMicrosoftTunnelConfiguration(value *MicrosoftTunnelConfiguration)() {
    m.microsoftTunnelConfiguration = value
}
// Sets the microsoftTunnelServers property value. A list of MicrosoftTunnelServers that are registered to this MicrosoftTunnelSite
// Parameters:
//  - value : Value to set for the microsoftTunnelServers property.
func (m *MicrosoftTunnelSite) SetMicrosoftTunnelServers(value []MicrosoftTunnelServer)() {
    m.microsoftTunnelServers = value
}
// Sets the publicAddress property value. The MicrosoftTunnelSite's public domain name or IP address
// Parameters:
//  - value : Value to set for the publicAddress property.
func (m *MicrosoftTunnelSite) SetPublicAddress(value *string)() {
    m.publicAddress = value
}
// Sets the roleScopeTagIds property value. List of Scope Tags for this Entity instance.
// Parameters:
//  - value : Value to set for the roleScopeTagIds property.
func (m *MicrosoftTunnelSite) SetRoleScopeTagIds(value []string)() {
    m.roleScopeTagIds = value
}
// Sets the upgradeAutomatically property value. The site's automatic upgrade setting. True for automatic upgrades, false for manual control
// Parameters:
//  - value : Value to set for the upgradeAutomatically property.
func (m *MicrosoftTunnelSite) SetUpgradeAutomatically(value *bool)() {
    m.upgradeAutomatically = value
}
// Sets the upgradeAvailable property value. True if an upgrade is available
// Parameters:
//  - value : Value to set for the upgradeAvailable property.
func (m *MicrosoftTunnelSite) SetUpgradeAvailable(value *bool)() {
    m.upgradeAvailable = value
}
// Sets the upgradeWindowEndTime property value. The site's upgrade window end time of day
// Parameters:
//  - value : Value to set for the upgradeWindowEndTime property.
func (m *MicrosoftTunnelSite) SetUpgradeWindowEndTime(value *string)() {
    m.upgradeWindowEndTime = value
}
// Sets the upgradeWindowStartTime property value. The site's upgrade window start time of day
// Parameters:
//  - value : Value to set for the upgradeWindowStartTime property.
func (m *MicrosoftTunnelSite) SetUpgradeWindowStartTime(value *string)() {
    m.upgradeWindowStartTime = value
}
// Sets the upgradeWindowUtcOffsetInMinutes property value. The site's timezone represented as a minute offset from UTC
// Parameters:
//  - value : Value to set for the upgradeWindowUtcOffsetInMinutes property.
func (m *MicrosoftTunnelSite) SetUpgradeWindowUtcOffsetInMinutes(value *int32)() {
    m.upgradeWindowUtcOffsetInMinutes = value
}
