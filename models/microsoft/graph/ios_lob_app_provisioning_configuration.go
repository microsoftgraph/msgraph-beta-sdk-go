package graph

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// 
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
// Instantiates a new iosLobAppProvisioningConfiguration and sets the default values.
func NewIosLobAppProvisioningConfiguration()(*IosLobAppProvisioningConfiguration) {
    m := &IosLobAppProvisioningConfiguration{
        Entity: *NewEntity(),
    }
    return m
}
// Gets the assignments property value. The associated group assignments for IosLobAppProvisioningConfiguration.
func (m *IosLobAppProvisioningConfiguration) GetAssignments()([]IosLobAppProvisioningConfigurationAssignment) {
    if m == nil {
        return nil
    } else {
        return m.assignments
    }
}
// Gets the createdDateTime property value. DateTime the object was created.
func (m *IosLobAppProvisioningConfiguration) GetCreatedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.createdDateTime
    }
}
// Gets the description property value. Admin provided description of the Device Configuration.
func (m *IosLobAppProvisioningConfiguration) GetDescription()(*string) {
    if m == nil {
        return nil
    } else {
        return m.description
    }
}
// Gets the deviceStatuses property value. The list of device installation states for this mobile app configuration.
func (m *IosLobAppProvisioningConfiguration) GetDeviceStatuses()([]ManagedDeviceMobileAppConfigurationDeviceStatus) {
    if m == nil {
        return nil
    } else {
        return m.deviceStatuses
    }
}
// Gets the displayName property value. Admin provided name of the device configuration.
func (m *IosLobAppProvisioningConfiguration) GetDisplayName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.displayName
    }
}
// Gets the expirationDateTime property value. Optional profile expiration date and time.
func (m *IosLobAppProvisioningConfiguration) GetExpirationDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.expirationDateTime
    }
}
// Gets the groupAssignments property value. The associated group assignments.
func (m *IosLobAppProvisioningConfiguration) GetGroupAssignments()([]MobileAppProvisioningConfigGroupAssignment) {
    if m == nil {
        return nil
    } else {
        return m.groupAssignments
    }
}
// Gets the lastModifiedDateTime property value. DateTime the object was last modified.
func (m *IosLobAppProvisioningConfiguration) GetLastModifiedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.lastModifiedDateTime
    }
}
// Gets the payload property value. Payload. (UTF8 encoded byte array)
func (m *IosLobAppProvisioningConfiguration) GetPayload()([]byte) {
    if m == nil {
        return nil
    } else {
        return m.payload
    }
}
// Gets the payloadFileName property value. Payload file name (.mobileprovision
func (m *IosLobAppProvisioningConfiguration) GetPayloadFileName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.payloadFileName
    }
}
// Gets the roleScopeTagIds property value. List of Scope Tags for this iOS LOB app provisioning configuration entity.
func (m *IosLobAppProvisioningConfiguration) GetRoleScopeTagIds()([]string) {
    if m == nil {
        return nil
    } else {
        return m.roleScopeTagIds
    }
}
// Gets the userStatuses property value. The list of user installation states for this mobile app configuration.
func (m *IosLobAppProvisioningConfiguration) GetUserStatuses()([]ManagedDeviceMobileAppConfigurationUserStatus) {
    if m == nil {
        return nil
    } else {
        return m.userStatuses
    }
}
// Gets the version property value. Version of the device configuration.
func (m *IosLobAppProvisioningConfiguration) GetVersion()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.version
    }
}
// The deserialization information for the current model
func (m *IosLobAppProvisioningConfiguration) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["assignments"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewIosLobAppProvisioningConfigurationAssignment() })
        if err != nil {
            return err
        }
        res := make([]IosLobAppProvisioningConfigurationAssignment, len(val))
        for i, v := range val {
            res[i] = *(v.(*IosLobAppProvisioningConfigurationAssignment))
        }
        m.SetAssignments(res)
        return nil
    }
    res["createdDateTime"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        m.SetCreatedDateTime(val)
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
    res["deviceStatuses"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewManagedDeviceMobileAppConfigurationDeviceStatus() })
        if err != nil {
            return err
        }
        res := make([]ManagedDeviceMobileAppConfigurationDeviceStatus, len(val))
        for i, v := range val {
            res[i] = *(v.(*ManagedDeviceMobileAppConfigurationDeviceStatus))
        }
        m.SetDeviceStatuses(res)
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
    res["expirationDateTime"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        m.SetExpirationDateTime(val)
        return nil
    }
    res["groupAssignments"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewMobileAppProvisioningConfigGroupAssignment() })
        if err != nil {
            return err
        }
        res := make([]MobileAppProvisioningConfigGroupAssignment, len(val))
        for i, v := range val {
            res[i] = *(v.(*MobileAppProvisioningConfigGroupAssignment))
        }
        m.SetGroupAssignments(res)
        return nil
    }
    res["lastModifiedDateTime"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        m.SetLastModifiedDateTime(val)
        return nil
    }
    res["payload"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetByteArrayValue()
        if err != nil {
            return err
        }
        m.SetPayload(val)
        return nil
    }
    res["payloadFileName"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetPayloadFileName(val)
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
    res["userStatuses"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewManagedDeviceMobileAppConfigurationUserStatus() })
        if err != nil {
            return err
        }
        res := make([]ManagedDeviceMobileAppConfigurationUserStatus, len(val))
        for i, v := range val {
            res[i] = *(v.(*ManagedDeviceMobileAppConfigurationUserStatus))
        }
        m.SetUserStatuses(res)
        return nil
    }
    res["version"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        m.SetVersion(val)
        return nil
    }
    return res
}
func (m *IosLobAppProvisioningConfiguration) IsNil()(bool) {
    return m == nil
}
// Serializes information the current object
// Parameters:
//  - writer : Serialization writer to use to serialize this model
func (m *IosLobAppProvisioningConfiguration) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
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
    {
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
    {
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
    {
        err = writer.WriteCollectionOfStringValues("roleScopeTagIds", m.GetRoleScopeTagIds())
        if err != nil {
            return err
        }
    }
    {
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
// Sets the assignments property value. The associated group assignments for IosLobAppProvisioningConfiguration.
// Parameters:
//  - value : Value to set for the assignments property.
func (m *IosLobAppProvisioningConfiguration) SetAssignments(value []IosLobAppProvisioningConfigurationAssignment)() {
    m.assignments = value
}
// Sets the createdDateTime property value. DateTime the object was created.
// Parameters:
//  - value : Value to set for the createdDateTime property.
func (m *IosLobAppProvisioningConfiguration) SetCreatedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.createdDateTime = value
}
// Sets the description property value. Admin provided description of the Device Configuration.
// Parameters:
//  - value : Value to set for the description property.
func (m *IosLobAppProvisioningConfiguration) SetDescription(value *string)() {
    m.description = value
}
// Sets the deviceStatuses property value. The list of device installation states for this mobile app configuration.
// Parameters:
//  - value : Value to set for the deviceStatuses property.
func (m *IosLobAppProvisioningConfiguration) SetDeviceStatuses(value []ManagedDeviceMobileAppConfigurationDeviceStatus)() {
    m.deviceStatuses = value
}
// Sets the displayName property value. Admin provided name of the device configuration.
// Parameters:
//  - value : Value to set for the displayName property.
func (m *IosLobAppProvisioningConfiguration) SetDisplayName(value *string)() {
    m.displayName = value
}
// Sets the expirationDateTime property value. Optional profile expiration date and time.
// Parameters:
//  - value : Value to set for the expirationDateTime property.
func (m *IosLobAppProvisioningConfiguration) SetExpirationDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.expirationDateTime = value
}
// Sets the groupAssignments property value. The associated group assignments.
// Parameters:
//  - value : Value to set for the groupAssignments property.
func (m *IosLobAppProvisioningConfiguration) SetGroupAssignments(value []MobileAppProvisioningConfigGroupAssignment)() {
    m.groupAssignments = value
}
// Sets the lastModifiedDateTime property value. DateTime the object was last modified.
// Parameters:
//  - value : Value to set for the lastModifiedDateTime property.
func (m *IosLobAppProvisioningConfiguration) SetLastModifiedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.lastModifiedDateTime = value
}
// Sets the payload property value. Payload. (UTF8 encoded byte array)
// Parameters:
//  - value : Value to set for the payload property.
func (m *IosLobAppProvisioningConfiguration) SetPayload(value []byte)() {
    m.payload = value
}
// Sets the payloadFileName property value. Payload file name (.mobileprovision
// Parameters:
//  - value : Value to set for the payloadFileName property.
func (m *IosLobAppProvisioningConfiguration) SetPayloadFileName(value *string)() {
    m.payloadFileName = value
}
// Sets the roleScopeTagIds property value. List of Scope Tags for this iOS LOB app provisioning configuration entity.
// Parameters:
//  - value : Value to set for the roleScopeTagIds property.
func (m *IosLobAppProvisioningConfiguration) SetRoleScopeTagIds(value []string)() {
    m.roleScopeTagIds = value
}
// Sets the userStatuses property value. The list of user installation states for this mobile app configuration.
// Parameters:
//  - value : Value to set for the userStatuses property.
func (m *IosLobAppProvisioningConfiguration) SetUserStatuses(value []ManagedDeviceMobileAppConfigurationUserStatus)() {
    m.userStatuses = value
}
// Sets the version property value. Version of the device configuration.
// Parameters:
//  - value : Value to set for the version property.
func (m *IosLobAppProvisioningConfiguration) SetVersion(value *int32)() {
    m.version = value
}
