package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// MicrosoftTunnelSite entity that represents a Microsoft Tunnel site
type MicrosoftTunnelSite struct {
    Entity
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // The MicrosoftTunnelSite's description
    description *string
    // The MicrosoftTunnelSite's display name
    displayName *string
    // The MicrosoftTunnelSite's Internal Network Access Probe URL
    internalNetworkProbeUrl *string
    // The MicrosoftTunnelConfiguration that has been applied to this MicrosoftTunnelSite
    microsoftTunnelConfiguration MicrosoftTunnelConfigurationable
    // A list of MicrosoftTunnelServers that are registered to this MicrosoftTunnelSite
    microsoftTunnelServers []MicrosoftTunnelServerable
    // The MicrosoftTunnelSite's public domain name or IP address
    publicAddress *string
    // List of Scope Tags for this Entity instance.
    roleScopeTagIds []string
    // The site's automatic upgrade setting. True for automatic upgrades, false for manual control
    upgradeAutomatically *bool
    // True if an upgrade is available
    upgradeAvailable *bool
    // The site's upgrade window end time of day
    upgradeWindowEndTime *i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.TimeOnly
    // The site's upgrade window start time of day
    upgradeWindowStartTime *i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.TimeOnly
    // The site's timezone represented as a minute offset from UTC
    upgradeWindowUtcOffsetInMinutes *int32
}
// NewMicrosoftTunnelSite instantiates a new microsoftTunnelSite and sets the default values.
func NewMicrosoftTunnelSite()(*MicrosoftTunnelSite) {
    m := &MicrosoftTunnelSite{
        Entity: *NewEntity(),
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateMicrosoftTunnelSiteFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateMicrosoftTunnelSiteFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewMicrosoftTunnelSite(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *MicrosoftTunnelSite) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
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
// GetFieldDeserializers the deserialization information for the current model
func (m *MicrosoftTunnelSite) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
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
    res["internalNetworkProbeUrl"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetInternalNetworkProbeUrl(val)
        }
        return nil
    }
    res["microsoftTunnelConfiguration"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateMicrosoftTunnelConfigurationFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMicrosoftTunnelConfiguration(val.(MicrosoftTunnelConfigurationable))
        }
        return nil
    }
    res["microsoftTunnelServers"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateMicrosoftTunnelServerFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]MicrosoftTunnelServerable, len(val))
            for i, v := range val {
                res[i] = v.(MicrosoftTunnelServerable)
            }
            m.SetMicrosoftTunnelServers(res)
        }
        return nil
    }
    res["publicAddress"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPublicAddress(val)
        }
        return nil
    }
    res["roleScopeTagIds"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
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
    res["upgradeAutomatically"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUpgradeAutomatically(val)
        }
        return nil
    }
    res["upgradeAvailable"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUpgradeAvailable(val)
        }
        return nil
    }
    res["upgradeWindowEndTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeOnlyValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUpgradeWindowEndTime(val)
        }
        return nil
    }
    res["upgradeWindowStartTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeOnlyValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUpgradeWindowStartTime(val)
        }
        return nil
    }
    res["upgradeWindowUtcOffsetInMinutes"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
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
// GetInternalNetworkProbeUrl gets the internalNetworkProbeUrl property value. The MicrosoftTunnelSite's Internal Network Access Probe URL
func (m *MicrosoftTunnelSite) GetInternalNetworkProbeUrl()(*string) {
    if m == nil {
        return nil
    } else {
        return m.internalNetworkProbeUrl
    }
}
// GetMicrosoftTunnelConfiguration gets the microsoftTunnelConfiguration property value. The MicrosoftTunnelConfiguration that has been applied to this MicrosoftTunnelSite
func (m *MicrosoftTunnelSite) GetMicrosoftTunnelConfiguration()(MicrosoftTunnelConfigurationable) {
    if m == nil {
        return nil
    } else {
        return m.microsoftTunnelConfiguration
    }
}
// GetMicrosoftTunnelServers gets the microsoftTunnelServers property value. A list of MicrosoftTunnelServers that are registered to this MicrosoftTunnelSite
func (m *MicrosoftTunnelSite) GetMicrosoftTunnelServers()([]MicrosoftTunnelServerable) {
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
func (m *MicrosoftTunnelSite) GetUpgradeWindowEndTime()(*i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.TimeOnly) {
    if m == nil {
        return nil
    } else {
        return m.upgradeWindowEndTime
    }
}
// GetUpgradeWindowStartTime gets the upgradeWindowStartTime property value. The site's upgrade window start time of day
func (m *MicrosoftTunnelSite) GetUpgradeWindowStartTime()(*i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.TimeOnly) {
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
// Serialize serializes information the current object
func (m *MicrosoftTunnelSite) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
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
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetMicrosoftTunnelServers()))
        for i, v := range m.GetMicrosoftTunnelServers() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
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
    {
        err = writer.WriteAdditionalData(m.GetAdditionalData())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAdditionalData sets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *MicrosoftTunnelSite) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
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
func (m *MicrosoftTunnelSite) SetMicrosoftTunnelConfiguration(value MicrosoftTunnelConfigurationable)() {
    if m != nil {
        m.microsoftTunnelConfiguration = value
    }
}
// SetMicrosoftTunnelServers sets the microsoftTunnelServers property value. A list of MicrosoftTunnelServers that are registered to this MicrosoftTunnelSite
func (m *MicrosoftTunnelSite) SetMicrosoftTunnelServers(value []MicrosoftTunnelServerable)() {
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
func (m *MicrosoftTunnelSite) SetUpgradeWindowEndTime(value *i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.TimeOnly)() {
    if m != nil {
        m.upgradeWindowEndTime = value
    }
}
// SetUpgradeWindowStartTime sets the upgradeWindowStartTime property value. The site's upgrade window start time of day
func (m *MicrosoftTunnelSite) SetUpgradeWindowStartTime(value *i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.TimeOnly)() {
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
