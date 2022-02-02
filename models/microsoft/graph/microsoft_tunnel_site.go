package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// MicrosoftTunnelSite 
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
    upgradeWindowEndTime *i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.TimeOnly;
    // The site's upgrade window start time of day
    upgradeWindowStartTime *i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.TimeOnly;
    // The site's timezone represented as a minute offset from UTC
    upgradeWindowUtcOffsetInMinutes *int32;
}
// NewMicrosoftTunnelSite instantiates a new microsoftTunnelSite and sets the default values.
func NewMicrosoftTunnelSite()(*MicrosoftTunnelSite) {
    m := &MicrosoftTunnelSite{
        Entity: *NewEntity(),
    }
    return m
}
// GetDescription gets the description property value. The MicrosoftTunnelSite's description
func (m *MicrosoftTunnelSite) GetDescription()(*string) {
    if m == nil {
        return nil
    } else {
        return m.description
    }
}
// GetDisplayName gets the displayName property value. The MicrosoftTunnelSite's display name
func (m *MicrosoftTunnelSite) GetDisplayName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.displayName
    }
}
// GetInternalNetworkProbeUrl gets the internalNetworkProbeUrl property value. The MicrosoftTunnelSite's Internal Network Access Probe URL
func (m *MicrosoftTunnelSite) GetInternalNetworkProbeUrl()(*string) {
    if m == nil {
        return nil
    } else {
        return m.internalNetworkProbeUrl
    }
}
// GetMicrosoftTunnelConfiguration gets the microsoftTunnelConfiguration property value. The MicrosoftTunnelConfiguration that has been applied to this MicrosoftTunnelSite
func (m *MicrosoftTunnelSite) GetMicrosoftTunnelConfiguration()(*MicrosoftTunnelConfiguration) {
    if m == nil {
        return nil
    } else {
        return m.microsoftTunnelConfiguration
    }
}
// GetMicrosoftTunnelServers gets the microsoftTunnelServers property value. A list of MicrosoftTunnelServers that are registered to this MicrosoftTunnelSite
func (m *MicrosoftTunnelSite) GetMicrosoftTunnelServers()([]MicrosoftTunnelServer) {
    if m == nil {
        return nil
    } else {
        return m.microsoftTunnelServers
    }
}
// GetPublicAddress gets the publicAddress property value. The MicrosoftTunnelSite's public domain name or IP address
func (m *MicrosoftTunnelSite) GetPublicAddress()(*string) {
    if m == nil {
        return nil
    } else {
        return m.publicAddress
    }
}
// GetRoleScopeTagIds gets the roleScopeTagIds property value. List of Scope Tags for this Entity instance.
func (m *MicrosoftTunnelSite) GetRoleScopeTagIds()([]string) {
    if m == nil {
        return nil
    } else {
        return m.roleScopeTagIds
    }
}
// GetUpgradeAutomatically gets the upgradeAutomatically property value. The site's automatic upgrade setting. True for automatic upgrades, false for manual control
func (m *MicrosoftTunnelSite) GetUpgradeAutomatically()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.upgradeAutomatically
    }
}
// GetUpgradeAvailable gets the upgradeAvailable property value. True if an upgrade is available
func (m *MicrosoftTunnelSite) GetUpgradeAvailable()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.upgradeAvailable
    }
}
// GetUpgradeWindowEndTime gets the upgradeWindowEndTime property value. The site's upgrade window end time of day
func (m *MicrosoftTunnelSite) GetUpgradeWindowEndTime()(*i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.TimeOnly) {
    if m == nil {
        return nil
    } else {
        return m.upgradeWindowEndTime
    }
}
// GetUpgradeWindowStartTime gets the upgradeWindowStartTime property value. The site's upgrade window start time of day
func (m *MicrosoftTunnelSite) GetUpgradeWindowStartTime()(*i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.TimeOnly) {
    if m == nil {
        return nil
    } else {
        return m.upgradeWindowStartTime
    }
}
// GetUpgradeWindowUtcOffsetInMinutes gets the upgradeWindowUtcOffsetInMinutes property value. The site's timezone represented as a minute offset from UTC
func (m *MicrosoftTunnelSite) GetUpgradeWindowUtcOffsetInMinutes()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.upgradeWindowUtcOffsetInMinutes
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *MicrosoftTunnelSite) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
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
    res["internalNetworkProbeUrl"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetInternalNetworkProbeUrl(val)
        }
        return nil
    }
    res["microsoftTunnelConfiguration"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewMicrosoftTunnelConfiguration() })
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMicrosoftTunnelConfiguration(val.(*MicrosoftTunnelConfiguration))
        }
        return nil
    }
    res["microsoftTunnelServers"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewMicrosoftTunnelServer() })
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]MicrosoftTunnelServer, len(val))
            for i, v := range val {
                res[i] = *(v.(*MicrosoftTunnelServer))
            }
            m.SetMicrosoftTunnelServers(res)
        }
        return nil
    }
    res["publicAddress"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPublicAddress(val)
        }
        return nil
    }
    res["roleScopeTagIds"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]string, len(val))
            for i, v := range val {
                res[i] = *(v.(*string))
            }
            m.SetRoleScopeTagIds(res)
        }
        return nil
    }
    res["upgradeAutomatically"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUpgradeAutomatically(val)
        }
        return nil
    }
    res["upgradeAvailable"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUpgradeAvailable(val)
        }
        return nil
    }
    res["upgradeWindowEndTime"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetTimeOnlyValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUpgradeWindowEndTime(val)
        }
        return nil
    }
    res["upgradeWindowStartTime"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetTimeOnlyValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUpgradeWindowStartTime(val)
        }
        return nil
    }
    res["upgradeWindowUtcOffsetInMinutes"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUpgradeWindowUtcOffsetInMinutes(val)
        }
        return nil
    }
    return res
}
func (m *MicrosoftTunnelSite) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
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
    if m.GetMicrosoftTunnelServers() != nil {
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
    if m.GetRoleScopeTagIds() != nil {
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
        err = writer.WriteTimeOnlyValue("upgradeWindowEndTime", m.GetUpgradeWindowEndTime())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteTimeOnlyValue("upgradeWindowStartTime", m.GetUpgradeWindowStartTime())
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
// SetDescription sets the description property value. The MicrosoftTunnelSite's description
func (m *MicrosoftTunnelSite) SetDescription(value *string)() {
    if m != nil {
        m.description = value
    }
}
// SetDisplayName sets the displayName property value. The MicrosoftTunnelSite's display name
func (m *MicrosoftTunnelSite) SetDisplayName(value *string)() {
    if m != nil {
        m.displayName = value
    }
}
// SetInternalNetworkProbeUrl sets the internalNetworkProbeUrl property value. The MicrosoftTunnelSite's Internal Network Access Probe URL
func (m *MicrosoftTunnelSite) SetInternalNetworkProbeUrl(value *string)() {
    if m != nil {
        m.internalNetworkProbeUrl = value
    }
}
// SetMicrosoftTunnelConfiguration sets the microsoftTunnelConfiguration property value. The MicrosoftTunnelConfiguration that has been applied to this MicrosoftTunnelSite
func (m *MicrosoftTunnelSite) SetMicrosoftTunnelConfiguration(value *MicrosoftTunnelConfiguration)() {
    if m != nil {
        m.microsoftTunnelConfiguration = value
    }
}
// SetMicrosoftTunnelServers sets the microsoftTunnelServers property value. A list of MicrosoftTunnelServers that are registered to this MicrosoftTunnelSite
func (m *MicrosoftTunnelSite) SetMicrosoftTunnelServers(value []MicrosoftTunnelServer)() {
    if m != nil {
        m.microsoftTunnelServers = value
    }
}
// SetPublicAddress sets the publicAddress property value. The MicrosoftTunnelSite's public domain name or IP address
func (m *MicrosoftTunnelSite) SetPublicAddress(value *string)() {
    if m != nil {
        m.publicAddress = value
    }
}
// SetRoleScopeTagIds sets the roleScopeTagIds property value. List of Scope Tags for this Entity instance.
func (m *MicrosoftTunnelSite) SetRoleScopeTagIds(value []string)() {
    if m != nil {
        m.roleScopeTagIds = value
    }
}
// SetUpgradeAutomatically sets the upgradeAutomatically property value. The site's automatic upgrade setting. True for automatic upgrades, false for manual control
func (m *MicrosoftTunnelSite) SetUpgradeAutomatically(value *bool)() {
    if m != nil {
        m.upgradeAutomatically = value
    }
}
// SetUpgradeAvailable sets the upgradeAvailable property value. True if an upgrade is available
func (m *MicrosoftTunnelSite) SetUpgradeAvailable(value *bool)() {
    if m != nil {
        m.upgradeAvailable = value
    }
}
// SetUpgradeWindowEndTime sets the upgradeWindowEndTime property value. The site's upgrade window end time of day
func (m *MicrosoftTunnelSite) SetUpgradeWindowEndTime(value *i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.TimeOnly)() {
    if m != nil {
        m.upgradeWindowEndTime = value
    }
}
// SetUpgradeWindowStartTime sets the upgradeWindowStartTime property value. The site's upgrade window start time of day
func (m *MicrosoftTunnelSite) SetUpgradeWindowStartTime(value *i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.TimeOnly)() {
    if m != nil {
        m.upgradeWindowStartTime = value
    }
}
// SetUpgradeWindowUtcOffsetInMinutes sets the upgradeWindowUtcOffsetInMinutes property value. The site's timezone represented as a minute offset from UTC
func (m *MicrosoftTunnelSite) SetUpgradeWindowUtcOffsetInMinutes(value *int32)() {
    if m != nil {
        m.upgradeWindowUtcOffsetInMinutes = value
    }
}
