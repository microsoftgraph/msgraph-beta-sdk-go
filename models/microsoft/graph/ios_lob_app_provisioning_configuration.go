package graph

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// IosLobAppProvisioningConfiguration 
type IosLobAppProvisioningConfiguration struct {
    Entity
    // The associated group assignments for IosLobAppProvisioningConfiguration.
    assignments []IosLobAppProvisioningConfigurationAssignment;
    // DateTime the object was created.
    createdDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time;
    // Admin provided description of the Device Configuration.
    description *string;
    // The list of device installation states for this mobile app configuration.
    deviceStatuses []ManagedDeviceMobileAppConfigurationDeviceStatus;
    // Admin provided name of the device configuration.
    displayName *string;
    // Optional profile expiration date and time.
    expirationDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time;
    // The associated group assignments.
    groupAssignments []MobileAppProvisioningConfigGroupAssignment;
    // DateTime the object was last modified.
    lastModifiedDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time;
    // Payload. (UTF8 encoded byte array)
    payload []byte;
    // Payload file name (.mobileprovision
    payloadFileName *string;
    // List of Scope Tags for this iOS LOB app provisioning configuration entity.
    roleScopeTagIds []string;
    // The list of user installation states for this mobile app configuration.
    userStatuses []ManagedDeviceMobileAppConfigurationUserStatus;
    // Version of the device configuration.
    version *int32;
}
// NewIosLobAppProvisioningConfiguration instantiates a new iosLobAppProvisioningConfiguration and sets the default values.
func NewIosLobAppProvisioningConfiguration()(*IosLobAppProvisioningConfiguration) {
    m := &IosLobAppProvisioningConfiguration{
        Entity: *NewEntity(),
    }
    return m
}
// GetAssignments gets the assignments property value. The associated group assignments for IosLobAppProvisioningConfiguration.
func (m *IosLobAppProvisioningConfiguration) GetAssignments()([]IosLobAppProvisioningConfigurationAssignment) {
    if m == nil {
        return nil
    } else {
        return m.assignments
    }
}
// GetCreatedDateTime gets the createdDateTime property value. DateTime the object was created.
func (m *IosLobAppProvisioningConfiguration) GetCreatedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.createdDateTime
    }
}
// GetDescription gets the description property value. Admin provided description of the Device Configuration.
func (m *IosLobAppProvisioningConfiguration) GetDescription()(*string) {
    if m == nil {
        return nil
    } else {
        return m.description
    }
}
// GetDeviceStatuses gets the deviceStatuses property value. The list of device installation states for this mobile app configuration.
func (m *IosLobAppProvisioningConfiguration) GetDeviceStatuses()([]ManagedDeviceMobileAppConfigurationDeviceStatus) {
    if m == nil {
        return nil
    } else {
        return m.deviceStatuses
    }
}
// GetDisplayName gets the displayName property value. Admin provided name of the device configuration.
func (m *IosLobAppProvisioningConfiguration) GetDisplayName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.displayName
    }
}
// GetExpirationDateTime gets the expirationDateTime property value. Optional profile expiration date and time.
func (m *IosLobAppProvisioningConfiguration) GetExpirationDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.expirationDateTime
    }
}
// GetGroupAssignments gets the groupAssignments property value. The associated group assignments.
func (m *IosLobAppProvisioningConfiguration) GetGroupAssignments()([]MobileAppProvisioningConfigGroupAssignment) {
    if m == nil {
        return nil
    } else {
        return m.groupAssignments
    }
}
// GetLastModifiedDateTime gets the lastModifiedDateTime property value. DateTime the object was last modified.
func (m *IosLobAppProvisioningConfiguration) GetLastModifiedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.lastModifiedDateTime
    }
}
// GetPayload gets the payload property value. Payload. (UTF8 encoded byte array)
func (m *IosLobAppProvisioningConfiguration) GetPayload()([]byte) {
    if m == nil {
        return nil
    } else {
        return m.payload
    }
}
// GetPayloadFileName gets the payloadFileName property value. Payload file name (.mobileprovision
func (m *IosLobAppProvisioningConfiguration) GetPayloadFileName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.payloadFileName
    }
}
// GetRoleScopeTagIds gets the roleScopeTagIds property value. List of Scope Tags for this iOS LOB app provisioning configuration entity.
func (m *IosLobAppProvisioningConfiguration) GetRoleScopeTagIds()([]string) {
    if m == nil {
        return nil
    } else {
        return m.roleScopeTagIds
    }
}
// GetUserStatuses gets the userStatuses property value. The list of user installation states for this mobile app configuration.
func (m *IosLobAppProvisioningConfiguration) GetUserStatuses()([]ManagedDeviceMobileAppConfigurationUserStatus) {
    if m == nil {
        return nil
    } else {
        return m.userStatuses
    }
}
// GetVersion gets the version property value. Version of the device configuration.
func (m *IosLobAppProvisioningConfiguration) GetVersion()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.version
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *IosLobAppProvisioningConfiguration) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["assignments"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewIosLobAppProvisioningConfigurationAssignment() })
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]IosLobAppProvisioningConfigurationAssignment, len(val))
            for i, v := range val {
                res[i] = *(v.(*IosLobAppProvisioningConfigurationAssignment))
            }
            m.SetAssignments(res)
        }
        return nil
    }
    res["createdDateTime"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCreatedDateTime(val)
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
    res["deviceStatuses"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewManagedDeviceMobileAppConfigurationDeviceStatus() })
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]ManagedDeviceMobileAppConfigurationDeviceStatus, len(val))
            for i, v := range val {
                res[i] = *(v.(*ManagedDeviceMobileAppConfigurationDeviceStatus))
            }
            m.SetDeviceStatuses(res)
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
    res["expirationDateTime"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetExpirationDateTime(val)
        }
        return nil
    }
    res["groupAssignments"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewMobileAppProvisioningConfigGroupAssignment() })
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]MobileAppProvisioningConfigGroupAssignment, len(val))
            for i, v := range val {
                res[i] = *(v.(*MobileAppProvisioningConfigGroupAssignment))
            }
            m.SetGroupAssignments(res)
        }
        return nil
    }
    res["lastModifiedDateTime"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLastModifiedDateTime(val)
        }
        return nil
    }
    res["payload"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetByteArrayValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPayload(val)
        }
        return nil
    }
    res["payloadFileName"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPayloadFileName(val)
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
    res["userStatuses"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewManagedDeviceMobileAppConfigurationUserStatus() })
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]ManagedDeviceMobileAppConfigurationUserStatus, len(val))
            for i, v := range val {
                res[i] = *(v.(*ManagedDeviceMobileAppConfigurationUserStatus))
            }
            m.SetUserStatuses(res)
        }
        return nil
    }
    res["version"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetVersion(val)
        }
        return nil
    }
    return res
}
func (m *IosLobAppProvisioningConfiguration) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
func (m *IosLobAppProvisioningConfiguration) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetAssignments() != nil {
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
        err = writer.WriteTimeValue("createdDateTime", m.GetCreatedDateTime())
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
    if m.GetDeviceStatuses() != nil {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetDeviceStatuses()))
        for i, v := range m.GetDeviceStatuses() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("deviceStatuses", cast)
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
        err = writer.WriteTimeValue("expirationDateTime", m.GetExpirationDateTime())
        if err != nil {
            return err
        }
    }
    if m.GetGroupAssignments() != nil {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetGroupAssignments()))
        for i, v := range m.GetGroupAssignments() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("groupAssignments", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteTimeValue("lastModifiedDateTime", m.GetLastModifiedDateTime())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteByteArrayValue("payload", m.GetPayload())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("payloadFileName", m.GetPayloadFileName())
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
    if m.GetUserStatuses() != nil {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetUserStatuses()))
        for i, v := range m.GetUserStatuses() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("userStatuses", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("version", m.GetVersion())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAssignments sets the assignments property value. The associated group assignments for IosLobAppProvisioningConfiguration.
func (m *IosLobAppProvisioningConfiguration) SetAssignments(value []IosLobAppProvisioningConfigurationAssignment)() {
    if m != nil {
        m.assignments = value
    }
}
// SetCreatedDateTime sets the createdDateTime property value. DateTime the object was created.
func (m *IosLobAppProvisioningConfiguration) SetCreatedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    if m != nil {
        m.createdDateTime = value
    }
}
// SetDescription sets the description property value. Admin provided description of the Device Configuration.
func (m *IosLobAppProvisioningConfiguration) SetDescription(value *string)() {
    if m != nil {
        m.description = value
    }
}
// SetDeviceStatuses sets the deviceStatuses property value. The list of device installation states for this mobile app configuration.
func (m *IosLobAppProvisioningConfiguration) SetDeviceStatuses(value []ManagedDeviceMobileAppConfigurationDeviceStatus)() {
    if m != nil {
        m.deviceStatuses = value
    }
}
// SetDisplayName sets the displayName property value. Admin provided name of the device configuration.
func (m *IosLobAppProvisioningConfiguration) SetDisplayName(value *string)() {
    if m != nil {
        m.displayName = value
    }
}
// SetExpirationDateTime sets the expirationDateTime property value. Optional profile expiration date and time.
func (m *IosLobAppProvisioningConfiguration) SetExpirationDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    if m != nil {
        m.expirationDateTime = value
    }
}
// SetGroupAssignments sets the groupAssignments property value. The associated group assignments.
func (m *IosLobAppProvisioningConfiguration) SetGroupAssignments(value []MobileAppProvisioningConfigGroupAssignment)() {
    if m != nil {
        m.groupAssignments = value
    }
}
// SetLastModifiedDateTime sets the lastModifiedDateTime property value. DateTime the object was last modified.
func (m *IosLobAppProvisioningConfiguration) SetLastModifiedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    if m != nil {
        m.lastModifiedDateTime = value
    }
}
// SetPayload sets the payload property value. Payload. (UTF8 encoded byte array)
func (m *IosLobAppProvisioningConfiguration) SetPayload(value []byte)() {
    if m != nil {
        m.payload = value
    }
}
// SetPayloadFileName sets the payloadFileName property value. Payload file name (.mobileprovision
func (m *IosLobAppProvisioningConfiguration) SetPayloadFileName(value *string)() {
    if m != nil {
        m.payloadFileName = value
    }
}
// SetRoleScopeTagIds sets the roleScopeTagIds property value. List of Scope Tags for this iOS LOB app provisioning configuration entity.
func (m *IosLobAppProvisioningConfiguration) SetRoleScopeTagIds(value []string)() {
    if m != nil {
        m.roleScopeTagIds = value
    }
}
// SetUserStatuses sets the userStatuses property value. The list of user installation states for this mobile app configuration.
func (m *IosLobAppProvisioningConfiguration) SetUserStatuses(value []ManagedDeviceMobileAppConfigurationUserStatus)() {
    if m != nil {
        m.userStatuses = value
    }
}
// SetVersion sets the version property value. Version of the device configuration.
func (m *IosLobAppProvisioningConfiguration) SetVersion(value *int32)() {
    if m != nil {
        m.version = value
    }
}
