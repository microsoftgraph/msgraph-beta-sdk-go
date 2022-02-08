package graph

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// TeamworkDevice 
type TeamworkDevice struct {
    Entity
    // The activity properties that change based on the device usage.
    activity *TeamworkDeviceActivity;
    // The activity state of the device. The possible values are: unknown, busy, idle, unavailable, unknownFutureValue.
    activityState *TeamworkDeviceActivityState;
    // The company asset tag assigned by the admin on the device.
    companyAssetTag *string;
    // The configuration properties of the device.
    configuration *TeamworkDeviceConfiguration;
    // Identity of the user who enrolled the device to the tenant.
    createdBy *IdentitySet;
    // The UTC date and time when the device was enrolled to the tenant.
    createdDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time;
    // The signed-in user on the device.
    currentUser *TeamworkUserIdentity;
    // The type of device. The possible values are: unknown, ipPhone, teamsRoom, surfaceHub, collaborationBar, teamsDisplay, touchConsole, lowCostPhone, teamsPanel, sip, unknownFutureValue.
    deviceType *TeamworkDeviceType;
    // 
    hardwareDetail *TeamworkHardwareDetail;
    // The health properties of the device.
    health *TeamworkDeviceHealth;
    // The health status of the device. The possible values are: unknown, offline, critical, nonUrgent, healthy, unknownFutureValue.
    healthStatus *TeamworkDeviceHealthStatus;
    // Identity of the user who last modified the device details.
    lastModifiedBy *IdentitySet;
    // The UTC date and time when the device detail was last modified.
    lastModifiedDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time;
    // The notes added by the admin to the device.
    notes *string;
    // The async operations on the device.
    operations []TeamworkDeviceOperation;
}
// NewTeamworkDevice instantiates a new teamworkDevice and sets the default values.
func NewTeamworkDevice()(*TeamworkDevice) {
    m := &TeamworkDevice{
        Entity: *NewEntity(),
    }
    return m
}
// GetActivity gets the activity property value. The activity properties that change based on the device usage.
func (m *TeamworkDevice) GetActivity()(*TeamworkDeviceActivity) {
    if m == nil {
        return nil
    } else {
        return m.activity
    }
}
// GetActivityState gets the activityState property value. The activity state of the device. The possible values are: unknown, busy, idle, unavailable, unknownFutureValue.
func (m *TeamworkDevice) GetActivityState()(*TeamworkDeviceActivityState) {
    if m == nil {
        return nil
    } else {
        return m.activityState
    }
}
// GetCompanyAssetTag gets the companyAssetTag property value. The company asset tag assigned by the admin on the device.
func (m *TeamworkDevice) GetCompanyAssetTag()(*string) {
    if m == nil {
        return nil
    } else {
        return m.companyAssetTag
    }
}
// GetConfiguration gets the configuration property value. The configuration properties of the device.
func (m *TeamworkDevice) GetConfiguration()(*TeamworkDeviceConfiguration) {
    if m == nil {
        return nil
    } else {
        return m.configuration
    }
}
// GetCreatedBy gets the createdBy property value. Identity of the user who enrolled the device to the tenant.
func (m *TeamworkDevice) GetCreatedBy()(*IdentitySet) {
    if m == nil {
        return nil
    } else {
        return m.createdBy
    }
}
// GetCreatedDateTime gets the createdDateTime property value. The UTC date and time when the device was enrolled to the tenant.
func (m *TeamworkDevice) GetCreatedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.createdDateTime
    }
}
// GetCurrentUser gets the currentUser property value. The signed-in user on the device.
func (m *TeamworkDevice) GetCurrentUser()(*TeamworkUserIdentity) {
    if m == nil {
        return nil
    } else {
        return m.currentUser
    }
}
// GetDeviceType gets the deviceType property value. The type of device. The possible values are: unknown, ipPhone, teamsRoom, surfaceHub, collaborationBar, teamsDisplay, touchConsole, lowCostPhone, teamsPanel, sip, unknownFutureValue.
func (m *TeamworkDevice) GetDeviceType()(*TeamworkDeviceType) {
    if m == nil {
        return nil
    } else {
        return m.deviceType
    }
}
// GetHardwareDetail gets the hardwareDetail property value. 
func (m *TeamworkDevice) GetHardwareDetail()(*TeamworkHardwareDetail) {
    if m == nil {
        return nil
    } else {
        return m.hardwareDetail
    }
}
// GetHealth gets the health property value. The health properties of the device.
func (m *TeamworkDevice) GetHealth()(*TeamworkDeviceHealth) {
    if m == nil {
        return nil
    } else {
        return m.health
    }
}
// GetHealthStatus gets the healthStatus property value. The health status of the device. The possible values are: unknown, offline, critical, nonUrgent, healthy, unknownFutureValue.
func (m *TeamworkDevice) GetHealthStatus()(*TeamworkDeviceHealthStatus) {
    if m == nil {
        return nil
    } else {
        return m.healthStatus
    }
}
// GetLastModifiedBy gets the lastModifiedBy property value. Identity of the user who last modified the device details.
func (m *TeamworkDevice) GetLastModifiedBy()(*IdentitySet) {
    if m == nil {
        return nil
    } else {
        return m.lastModifiedBy
    }
}
// GetLastModifiedDateTime gets the lastModifiedDateTime property value. The UTC date and time when the device detail was last modified.
func (m *TeamworkDevice) GetLastModifiedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.lastModifiedDateTime
    }
}
// GetNotes gets the notes property value. The notes added by the admin to the device.
func (m *TeamworkDevice) GetNotes()(*string) {
    if m == nil {
        return nil
    } else {
        return m.notes
    }
}
// GetOperations gets the operations property value. The async operations on the device.
func (m *TeamworkDevice) GetOperations()([]TeamworkDeviceOperation) {
    if m == nil {
        return nil
    } else {
        return m.operations
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *TeamworkDevice) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["activity"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewTeamworkDeviceActivity() })
        if err != nil {
            return err
        }
        if val != nil {
            m.SetActivity(val.(*TeamworkDeviceActivity))
        }
        return nil
    }
    res["activityState"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParseTeamworkDeviceActivityState)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetActivityState(val.(*TeamworkDeviceActivityState))
        }
        return nil
    }
    res["companyAssetTag"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCompanyAssetTag(val)
        }
        return nil
    }
    res["configuration"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewTeamworkDeviceConfiguration() })
        if err != nil {
            return err
        }
        if val != nil {
            m.SetConfiguration(val.(*TeamworkDeviceConfiguration))
        }
        return nil
    }
    res["createdBy"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewIdentitySet() })
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCreatedBy(val.(*IdentitySet))
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
    res["currentUser"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewTeamworkUserIdentity() })
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCurrentUser(val.(*TeamworkUserIdentity))
        }
        return nil
    }
    res["deviceType"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParseTeamworkDeviceType)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDeviceType(val.(*TeamworkDeviceType))
        }
        return nil
    }
    res["hardwareDetail"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewTeamworkHardwareDetail() })
        if err != nil {
            return err
        }
        if val != nil {
            m.SetHardwareDetail(val.(*TeamworkHardwareDetail))
        }
        return nil
    }
    res["health"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewTeamworkDeviceHealth() })
        if err != nil {
            return err
        }
        if val != nil {
            m.SetHealth(val.(*TeamworkDeviceHealth))
        }
        return nil
    }
    res["healthStatus"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParseTeamworkDeviceHealthStatus)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetHealthStatus(val.(*TeamworkDeviceHealthStatus))
        }
        return nil
    }
    res["lastModifiedBy"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewIdentitySet() })
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLastModifiedBy(val.(*IdentitySet))
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
    res["notes"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetNotes(val)
        }
        return nil
    }
    res["operations"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewTeamworkDeviceOperation() })
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]TeamworkDeviceOperation, len(val))
            for i, v := range val {
                res[i] = *(v.(*TeamworkDeviceOperation))
            }
            m.SetOperations(res)
        }
        return nil
    }
    return res
}
func (m *TeamworkDevice) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
func (m *TeamworkDevice) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteObjectValue("activity", m.GetActivity())
        if err != nil {
            return err
        }
    }
    if m.GetActivityState() != nil {
        cast := (*m.GetActivityState()).String()
        err = writer.WriteStringValue("activityState", &cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("companyAssetTag", m.GetCompanyAssetTag())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("configuration", m.GetConfiguration())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("createdBy", m.GetCreatedBy())
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
        err = writer.WriteObjectValue("currentUser", m.GetCurrentUser())
        if err != nil {
            return err
        }
    }
    if m.GetDeviceType() != nil {
        cast := (*m.GetDeviceType()).String()
        err = writer.WriteStringValue("deviceType", &cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("hardwareDetail", m.GetHardwareDetail())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("health", m.GetHealth())
        if err != nil {
            return err
        }
    }
    if m.GetHealthStatus() != nil {
        cast := (*m.GetHealthStatus()).String()
        err = writer.WriteStringValue("healthStatus", &cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("lastModifiedBy", m.GetLastModifiedBy())
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
        err = writer.WriteStringValue("notes", m.GetNotes())
        if err != nil {
            return err
        }
    }
    if m.GetOperations() != nil {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetOperations()))
        for i, v := range m.GetOperations() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("operations", cast)
        if err != nil {
            return err
        }
    }
    return nil
}
// SetActivity sets the activity property value. The activity properties that change based on the device usage.
func (m *TeamworkDevice) SetActivity(value *TeamworkDeviceActivity)() {
    if m != nil {
        m.activity = value
    }
}
// SetActivityState sets the activityState property value. The activity state of the device. The possible values are: unknown, busy, idle, unavailable, unknownFutureValue.
func (m *TeamworkDevice) SetActivityState(value *TeamworkDeviceActivityState)() {
    if m != nil {
        m.activityState = value
    }
}
// SetCompanyAssetTag sets the companyAssetTag property value. The company asset tag assigned by the admin on the device.
func (m *TeamworkDevice) SetCompanyAssetTag(value *string)() {
    if m != nil {
        m.companyAssetTag = value
    }
}
// SetConfiguration sets the configuration property value. The configuration properties of the device.
func (m *TeamworkDevice) SetConfiguration(value *TeamworkDeviceConfiguration)() {
    if m != nil {
        m.configuration = value
    }
}
// SetCreatedBy sets the createdBy property value. Identity of the user who enrolled the device to the tenant.
func (m *TeamworkDevice) SetCreatedBy(value *IdentitySet)() {
    if m != nil {
        m.createdBy = value
    }
}
// SetCreatedDateTime sets the createdDateTime property value. The UTC date and time when the device was enrolled to the tenant.
func (m *TeamworkDevice) SetCreatedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    if m != nil {
        m.createdDateTime = value
    }
}
// SetCurrentUser sets the currentUser property value. The signed-in user on the device.
func (m *TeamworkDevice) SetCurrentUser(value *TeamworkUserIdentity)() {
    if m != nil {
        m.currentUser = value
    }
}
// SetDeviceType sets the deviceType property value. The type of device. The possible values are: unknown, ipPhone, teamsRoom, surfaceHub, collaborationBar, teamsDisplay, touchConsole, lowCostPhone, teamsPanel, sip, unknownFutureValue.
func (m *TeamworkDevice) SetDeviceType(value *TeamworkDeviceType)() {
    if m != nil {
        m.deviceType = value
    }
}
// SetHardwareDetail sets the hardwareDetail property value. 
func (m *TeamworkDevice) SetHardwareDetail(value *TeamworkHardwareDetail)() {
    if m != nil {
        m.hardwareDetail = value
    }
}
// SetHealth sets the health property value. The health properties of the device.
func (m *TeamworkDevice) SetHealth(value *TeamworkDeviceHealth)() {
    if m != nil {
        m.health = value
    }
}
// SetHealthStatus sets the healthStatus property value. The health status of the device. The possible values are: unknown, offline, critical, nonUrgent, healthy, unknownFutureValue.
func (m *TeamworkDevice) SetHealthStatus(value *TeamworkDeviceHealthStatus)() {
    if m != nil {
        m.healthStatus = value
    }
}
// SetLastModifiedBy sets the lastModifiedBy property value. Identity of the user who last modified the device details.
func (m *TeamworkDevice) SetLastModifiedBy(value *IdentitySet)() {
    if m != nil {
        m.lastModifiedBy = value
    }
}
// SetLastModifiedDateTime sets the lastModifiedDateTime property value. The UTC date and time when the device detail was last modified.
func (m *TeamworkDevice) SetLastModifiedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    if m != nil {
        m.lastModifiedDateTime = value
    }
}
// SetNotes sets the notes property value. The notes added by the admin to the device.
func (m *TeamworkDevice) SetNotes(value *string)() {
    if m != nil {
        m.notes = value
    }
}
// SetOperations sets the operations property value. The async operations on the device.
func (m *TeamworkDevice) SetOperations(value []TeamworkDeviceOperation)() {
    if m != nil {
        m.operations = value
    }
}
