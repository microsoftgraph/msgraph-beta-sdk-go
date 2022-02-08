package graph

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// DeviceManagementPartner 
type DeviceManagementPartner struct {
    Entity
    // Partner display name
    displayName *string;
    // User groups that specifies whether enrollment is through partner.
    groupsRequiringPartnerEnrollment []DeviceManagementPartnerAssignment;
    // Whether device management partner is configured or not
    isConfigured *bool;
    // Timestamp of last heartbeat after admin enabled option Connect to Device management Partner
    lastHeartbeatDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time;
    // Partner App type. Possible values are: unknown, singleTenantApp, multiTenantApp.
    partnerAppType *DeviceManagementPartnerAppType;
    // Partner state of this tenant. Possible values are: unknown, unavailable, enabled, terminated, rejected, unresponsive.
    partnerState *DeviceManagementPartnerTenantState;
    // Partner Single tenant App id
    singleTenantAppId *string;
    // DateTime in UTC when PartnerDevices will be marked as NonCompliant. This will become obselete soon.
    whenPartnerDevicesWillBeMarkedAsNonCompliant *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time;
    // DateTime in UTC when PartnerDevices will be marked as NonCompliant
    whenPartnerDevicesWillBeMarkedAsNonCompliantDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time;
    // DateTime in UTC when PartnerDevices will be removed. This will become obselete soon.
    whenPartnerDevicesWillBeRemoved *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time;
    // DateTime in UTC when PartnerDevices will be removed
    whenPartnerDevicesWillBeRemovedDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time;
}
// NewDeviceManagementPartner instantiates a new deviceManagementPartner and sets the default values.
func NewDeviceManagementPartner()(*DeviceManagementPartner) {
    m := &DeviceManagementPartner{
        Entity: *NewEntity(),
    }
    return m
}
// GetDisplayName gets the displayName property value. Partner display name
func (m *DeviceManagementPartner) GetDisplayName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.displayName
    }
}
// GetGroupsRequiringPartnerEnrollment gets the groupsRequiringPartnerEnrollment property value. User groups that specifies whether enrollment is through partner.
func (m *DeviceManagementPartner) GetGroupsRequiringPartnerEnrollment()([]DeviceManagementPartnerAssignment) {
    if m == nil {
        return nil
    } else {
        return m.groupsRequiringPartnerEnrollment
    }
}
// GetIsConfigured gets the isConfigured property value. Whether device management partner is configured or not
func (m *DeviceManagementPartner) GetIsConfigured()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.isConfigured
    }
}
// GetLastHeartbeatDateTime gets the lastHeartbeatDateTime property value. Timestamp of last heartbeat after admin enabled option Connect to Device management Partner
func (m *DeviceManagementPartner) GetLastHeartbeatDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.lastHeartbeatDateTime
    }
}
// GetPartnerAppType gets the partnerAppType property value. Partner App type. Possible values are: unknown, singleTenantApp, multiTenantApp.
func (m *DeviceManagementPartner) GetPartnerAppType()(*DeviceManagementPartnerAppType) {
    if m == nil {
        return nil
    } else {
        return m.partnerAppType
    }
}
// GetPartnerState gets the partnerState property value. Partner state of this tenant. Possible values are: unknown, unavailable, enabled, terminated, rejected, unresponsive.
func (m *DeviceManagementPartner) GetPartnerState()(*DeviceManagementPartnerTenantState) {
    if m == nil {
        return nil
    } else {
        return m.partnerState
    }
}
// GetSingleTenantAppId gets the singleTenantAppId property value. Partner Single tenant App id
func (m *DeviceManagementPartner) GetSingleTenantAppId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.singleTenantAppId
    }
}
// GetWhenPartnerDevicesWillBeMarkedAsNonCompliant gets the whenPartnerDevicesWillBeMarkedAsNonCompliant property value. DateTime in UTC when PartnerDevices will be marked as NonCompliant. This will become obselete soon.
func (m *DeviceManagementPartner) GetWhenPartnerDevicesWillBeMarkedAsNonCompliant()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.whenPartnerDevicesWillBeMarkedAsNonCompliant
    }
}
// GetWhenPartnerDevicesWillBeMarkedAsNonCompliantDateTime gets the whenPartnerDevicesWillBeMarkedAsNonCompliantDateTime property value. DateTime in UTC when PartnerDevices will be marked as NonCompliant
func (m *DeviceManagementPartner) GetWhenPartnerDevicesWillBeMarkedAsNonCompliantDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.whenPartnerDevicesWillBeMarkedAsNonCompliantDateTime
    }
}
// GetWhenPartnerDevicesWillBeRemoved gets the whenPartnerDevicesWillBeRemoved property value. DateTime in UTC when PartnerDevices will be removed. This will become obselete soon.
func (m *DeviceManagementPartner) GetWhenPartnerDevicesWillBeRemoved()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.whenPartnerDevicesWillBeRemoved
    }
}
// GetWhenPartnerDevicesWillBeRemovedDateTime gets the whenPartnerDevicesWillBeRemovedDateTime property value. DateTime in UTC when PartnerDevices will be removed
func (m *DeviceManagementPartner) GetWhenPartnerDevicesWillBeRemovedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.whenPartnerDevicesWillBeRemovedDateTime
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *DeviceManagementPartner) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
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
    res["groupsRequiringPartnerEnrollment"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewDeviceManagementPartnerAssignment() })
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]DeviceManagementPartnerAssignment, len(val))
            for i, v := range val {
                res[i] = *(v.(*DeviceManagementPartnerAssignment))
            }
            m.SetGroupsRequiringPartnerEnrollment(res)
        }
        return nil
    }
    res["isConfigured"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsConfigured(val)
        }
        return nil
    }
    res["lastHeartbeatDateTime"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLastHeartbeatDateTime(val)
        }
        return nil
    }
    res["partnerAppType"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParseDeviceManagementPartnerAppType)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPartnerAppType(val.(*DeviceManagementPartnerAppType))
        }
        return nil
    }
    res["partnerState"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParseDeviceManagementPartnerTenantState)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPartnerState(val.(*DeviceManagementPartnerTenantState))
        }
        return nil
    }
    res["singleTenantAppId"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSingleTenantAppId(val)
        }
        return nil
    }
    res["whenPartnerDevicesWillBeMarkedAsNonCompliant"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetWhenPartnerDevicesWillBeMarkedAsNonCompliant(val)
        }
        return nil
    }
    res["whenPartnerDevicesWillBeMarkedAsNonCompliantDateTime"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetWhenPartnerDevicesWillBeMarkedAsNonCompliantDateTime(val)
        }
        return nil
    }
    res["whenPartnerDevicesWillBeRemoved"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetWhenPartnerDevicesWillBeRemoved(val)
        }
        return nil
    }
    res["whenPartnerDevicesWillBeRemovedDateTime"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetWhenPartnerDevicesWillBeRemovedDateTime(val)
        }
        return nil
    }
    return res
}
func (m *DeviceManagementPartner) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
func (m *DeviceManagementPartner) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteStringValue("displayName", m.GetDisplayName())
        if err != nil {
            return err
        }
    }
    if m.GetGroupsRequiringPartnerEnrollment() != nil {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetGroupsRequiringPartnerEnrollment()))
        for i, v := range m.GetGroupsRequiringPartnerEnrollment() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("groupsRequiringPartnerEnrollment", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("isConfigured", m.GetIsConfigured())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteTimeValue("lastHeartbeatDateTime", m.GetLastHeartbeatDateTime())
        if err != nil {
            return err
        }
    }
    if m.GetPartnerAppType() != nil {
        cast := (*m.GetPartnerAppType()).String()
        err = writer.WriteStringValue("partnerAppType", &cast)
        if err != nil {
            return err
        }
    }
    if m.GetPartnerState() != nil {
        cast := (*m.GetPartnerState()).String()
        err = writer.WriteStringValue("partnerState", &cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("singleTenantAppId", m.GetSingleTenantAppId())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteTimeValue("whenPartnerDevicesWillBeMarkedAsNonCompliant", m.GetWhenPartnerDevicesWillBeMarkedAsNonCompliant())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteTimeValue("whenPartnerDevicesWillBeMarkedAsNonCompliantDateTime", m.GetWhenPartnerDevicesWillBeMarkedAsNonCompliantDateTime())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteTimeValue("whenPartnerDevicesWillBeRemoved", m.GetWhenPartnerDevicesWillBeRemoved())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteTimeValue("whenPartnerDevicesWillBeRemovedDateTime", m.GetWhenPartnerDevicesWillBeRemovedDateTime())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetDisplayName sets the displayName property value. Partner display name
func (m *DeviceManagementPartner) SetDisplayName(value *string)() {
    if m != nil {
        m.displayName = value
    }
}
// SetGroupsRequiringPartnerEnrollment sets the groupsRequiringPartnerEnrollment property value. User groups that specifies whether enrollment is through partner.
func (m *DeviceManagementPartner) SetGroupsRequiringPartnerEnrollment(value []DeviceManagementPartnerAssignment)() {
    if m != nil {
        m.groupsRequiringPartnerEnrollment = value
    }
}
// SetIsConfigured sets the isConfigured property value. Whether device management partner is configured or not
func (m *DeviceManagementPartner) SetIsConfigured(value *bool)() {
    if m != nil {
        m.isConfigured = value
    }
}
// SetLastHeartbeatDateTime sets the lastHeartbeatDateTime property value. Timestamp of last heartbeat after admin enabled option Connect to Device management Partner
func (m *DeviceManagementPartner) SetLastHeartbeatDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    if m != nil {
        m.lastHeartbeatDateTime = value
    }
}
// SetPartnerAppType sets the partnerAppType property value. Partner App type. Possible values are: unknown, singleTenantApp, multiTenantApp.
func (m *DeviceManagementPartner) SetPartnerAppType(value *DeviceManagementPartnerAppType)() {
    if m != nil {
        m.partnerAppType = value
    }
}
// SetPartnerState sets the partnerState property value. Partner state of this tenant. Possible values are: unknown, unavailable, enabled, terminated, rejected, unresponsive.
func (m *DeviceManagementPartner) SetPartnerState(value *DeviceManagementPartnerTenantState)() {
    if m != nil {
        m.partnerState = value
    }
}
// SetSingleTenantAppId sets the singleTenantAppId property value. Partner Single tenant App id
func (m *DeviceManagementPartner) SetSingleTenantAppId(value *string)() {
    if m != nil {
        m.singleTenantAppId = value
    }
}
// SetWhenPartnerDevicesWillBeMarkedAsNonCompliant sets the whenPartnerDevicesWillBeMarkedAsNonCompliant property value. DateTime in UTC when PartnerDevices will be marked as NonCompliant. This will become obselete soon.
func (m *DeviceManagementPartner) SetWhenPartnerDevicesWillBeMarkedAsNonCompliant(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    if m != nil {
        m.whenPartnerDevicesWillBeMarkedAsNonCompliant = value
    }
}
// SetWhenPartnerDevicesWillBeMarkedAsNonCompliantDateTime sets the whenPartnerDevicesWillBeMarkedAsNonCompliantDateTime property value. DateTime in UTC when PartnerDevices will be marked as NonCompliant
func (m *DeviceManagementPartner) SetWhenPartnerDevicesWillBeMarkedAsNonCompliantDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    if m != nil {
        m.whenPartnerDevicesWillBeMarkedAsNonCompliantDateTime = value
    }
}
// SetWhenPartnerDevicesWillBeRemoved sets the whenPartnerDevicesWillBeRemoved property value. DateTime in UTC when PartnerDevices will be removed. This will become obselete soon.
func (m *DeviceManagementPartner) SetWhenPartnerDevicesWillBeRemoved(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    if m != nil {
        m.whenPartnerDevicesWillBeRemoved = value
    }
}
// SetWhenPartnerDevicesWillBeRemovedDateTime sets the whenPartnerDevicesWillBeRemovedDateTime property value. DateTime in UTC when PartnerDevices will be removed
func (m *DeviceManagementPartner) SetWhenPartnerDevicesWillBeRemovedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    if m != nil {
        m.whenPartnerDevicesWillBeRemovedDateTime = value
    }
}
