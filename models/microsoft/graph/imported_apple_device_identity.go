package graph

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// 
type ImportedAppleDeviceIdentity struct {
    Entity
    // Created Date Time of the device
    createdDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time;
    // The description of the device
    description *string;
    // Apple device discovery source. Possible values are: unknown, adminImport, deviceEnrollmentProgram.
    discoverySource *DiscoverySource;
    // The state of the device in Intune. Possible values are: unknown, enrolled, pendingReset, failed, notContacted, blocked.
    enrollmentState *EnrollmentState;
    // Indicates if the device is deleted from Apple Business Manager
    isDeleted *bool;
    // Indicates if the Apple device is supervised. More information is at: https://support.apple.com/en-us/HT202837
    isSupervised *bool;
    // Last Contacted Date Time of the device
    lastContactedDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time;
    // The platform of the Device. Possible values are: unknown, ios, android, windows, windowsMobile, macOS.
    platform *Platform;
    // The time enrollment profile was assigned to the device
    requestedEnrollmentProfileAssignmentDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time;
    // Enrollment profile Id admin intends to apply to the device during next enrollment
    requestedEnrollmentProfileId *string;
    // Device serial number
    serialNumber *string;
}
// Instantiates a new importedAppleDeviceIdentity and sets the default values.
func NewImportedAppleDeviceIdentity()(*ImportedAppleDeviceIdentity) {
    m := &ImportedAppleDeviceIdentity{
        Entity: *NewEntity(),
    }
    return m
}
// Gets the createdDateTime property value. Created Date Time of the device
func (m *ImportedAppleDeviceIdentity) GetCreatedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.createdDateTime
    }
}
// Gets the description property value. The description of the device
func (m *ImportedAppleDeviceIdentity) GetDescription()(*string) {
    if m == nil {
        return nil
    } else {
        return m.description
    }
}
// Gets the discoverySource property value. Apple device discovery source. Possible values are: unknown, adminImport, deviceEnrollmentProgram.
func (m *ImportedAppleDeviceIdentity) GetDiscoverySource()(*DiscoverySource) {
    if m == nil {
        return nil
    } else {
        return m.discoverySource
    }
}
// Gets the enrollmentState property value. The state of the device in Intune. Possible values are: unknown, enrolled, pendingReset, failed, notContacted, blocked.
func (m *ImportedAppleDeviceIdentity) GetEnrollmentState()(*EnrollmentState) {
    if m == nil {
        return nil
    } else {
        return m.enrollmentState
    }
}
// Gets the isDeleted property value. Indicates if the device is deleted from Apple Business Manager
func (m *ImportedAppleDeviceIdentity) GetIsDeleted()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.isDeleted
    }
}
// Gets the isSupervised property value. Indicates if the Apple device is supervised. More information is at: https://support.apple.com/en-us/HT202837
func (m *ImportedAppleDeviceIdentity) GetIsSupervised()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.isSupervised
    }
}
// Gets the lastContactedDateTime property value. Last Contacted Date Time of the device
func (m *ImportedAppleDeviceIdentity) GetLastContactedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.lastContactedDateTime
    }
}
// Gets the platform property value. The platform of the Device. Possible values are: unknown, ios, android, windows, windowsMobile, macOS.
func (m *ImportedAppleDeviceIdentity) GetPlatform()(*Platform) {
    if m == nil {
        return nil
    } else {
        return m.platform
    }
}
// Gets the requestedEnrollmentProfileAssignmentDateTime property value. The time enrollment profile was assigned to the device
func (m *ImportedAppleDeviceIdentity) GetRequestedEnrollmentProfileAssignmentDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.requestedEnrollmentProfileAssignmentDateTime
    }
}
// Gets the requestedEnrollmentProfileId property value. Enrollment profile Id admin intends to apply to the device during next enrollment
func (m *ImportedAppleDeviceIdentity) GetRequestedEnrollmentProfileId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.requestedEnrollmentProfileId
    }
}
// Gets the serialNumber property value. Device serial number
func (m *ImportedAppleDeviceIdentity) GetSerialNumber()(*string) {
    if m == nil {
        return nil
    } else {
        return m.serialNumber
    }
}
// The deserialization information for the current model
func (m *ImportedAppleDeviceIdentity) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
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
    res["discoverySource"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParseDiscoverySource)
        if err != nil {
            return err
        }
        cast := val.(DiscoverySource)
        m.SetDiscoverySource(&cast)
        return nil
    }
    res["enrollmentState"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParseEnrollmentState)
        if err != nil {
            return err
        }
        cast := val.(EnrollmentState)
        m.SetEnrollmentState(&cast)
        return nil
    }
    res["isDeleted"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        m.SetIsDeleted(val)
        return nil
    }
    res["isSupervised"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        m.SetIsSupervised(val)
        return nil
    }
    res["lastContactedDateTime"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        m.SetLastContactedDateTime(val)
        return nil
    }
    res["platform"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParsePlatform)
        if err != nil {
            return err
        }
        cast := val.(Platform)
        m.SetPlatform(&cast)
        return nil
    }
    res["requestedEnrollmentProfileAssignmentDateTime"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        m.SetRequestedEnrollmentProfileAssignmentDateTime(val)
        return nil
    }
    res["requestedEnrollmentProfileId"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetRequestedEnrollmentProfileId(val)
        return nil
    }
    res["serialNumber"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetSerialNumber(val)
        return nil
    }
    return res
}
func (m *ImportedAppleDeviceIdentity) IsNil()(bool) {
    return m == nil
}
// Serializes information the current object
// Parameters:
//  - writer : Serialization writer to use to serialize this model
func (m *ImportedAppleDeviceIdentity) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
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
    if m.GetDiscoverySource() != nil {
        cast := m.GetDiscoverySource().String()
        err = writer.WriteStringValue("discoverySource", &cast)
        if err != nil {
            return err
        }
    }
    if m.GetEnrollmentState() != nil {
        cast := m.GetEnrollmentState().String()
        err = writer.WriteStringValue("enrollmentState", &cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("isDeleted", m.GetIsDeleted())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("isSupervised", m.GetIsSupervised())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteTimeValue("lastContactedDateTime", m.GetLastContactedDateTime())
        if err != nil {
            return err
        }
    }
    if m.GetPlatform() != nil {
        cast := m.GetPlatform().String()
        err = writer.WriteStringValue("platform", &cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteTimeValue("requestedEnrollmentProfileAssignmentDateTime", m.GetRequestedEnrollmentProfileAssignmentDateTime())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("requestedEnrollmentProfileId", m.GetRequestedEnrollmentProfileId())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("serialNumber", m.GetSerialNumber())
        if err != nil {
            return err
        }
    }
    return nil
}
// Sets the createdDateTime property value. Created Date Time of the device
// Parameters:
//  - value : Value to set for the createdDateTime property.
func (m *ImportedAppleDeviceIdentity) SetCreatedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.createdDateTime = value
}
// Sets the description property value. The description of the device
// Parameters:
//  - value : Value to set for the description property.
func (m *ImportedAppleDeviceIdentity) SetDescription(value *string)() {
    m.description = value
}
// Sets the discoverySource property value. Apple device discovery source. Possible values are: unknown, adminImport, deviceEnrollmentProgram.
// Parameters:
//  - value : Value to set for the discoverySource property.
func (m *ImportedAppleDeviceIdentity) SetDiscoverySource(value *DiscoverySource)() {
    m.discoverySource = value
}
// Sets the enrollmentState property value. The state of the device in Intune. Possible values are: unknown, enrolled, pendingReset, failed, notContacted, blocked.
// Parameters:
//  - value : Value to set for the enrollmentState property.
func (m *ImportedAppleDeviceIdentity) SetEnrollmentState(value *EnrollmentState)() {
    m.enrollmentState = value
}
// Sets the isDeleted property value. Indicates if the device is deleted from Apple Business Manager
// Parameters:
//  - value : Value to set for the isDeleted property.
func (m *ImportedAppleDeviceIdentity) SetIsDeleted(value *bool)() {
    m.isDeleted = value
}
// Sets the isSupervised property value. Indicates if the Apple device is supervised. More information is at: https://support.apple.com/en-us/HT202837
// Parameters:
//  - value : Value to set for the isSupervised property.
func (m *ImportedAppleDeviceIdentity) SetIsSupervised(value *bool)() {
    m.isSupervised = value
}
// Sets the lastContactedDateTime property value. Last Contacted Date Time of the device
// Parameters:
//  - value : Value to set for the lastContactedDateTime property.
func (m *ImportedAppleDeviceIdentity) SetLastContactedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.lastContactedDateTime = value
}
// Sets the platform property value. The platform of the Device. Possible values are: unknown, ios, android, windows, windowsMobile, macOS.
// Parameters:
//  - value : Value to set for the platform property.
func (m *ImportedAppleDeviceIdentity) SetPlatform(value *Platform)() {
    m.platform = value
}
// Sets the requestedEnrollmentProfileAssignmentDateTime property value. The time enrollment profile was assigned to the device
// Parameters:
//  - value : Value to set for the requestedEnrollmentProfileAssignmentDateTime property.
func (m *ImportedAppleDeviceIdentity) SetRequestedEnrollmentProfileAssignmentDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.requestedEnrollmentProfileAssignmentDateTime = value
}
// Sets the requestedEnrollmentProfileId property value. Enrollment profile Id admin intends to apply to the device during next enrollment
// Parameters:
//  - value : Value to set for the requestedEnrollmentProfileId property.
func (m *ImportedAppleDeviceIdentity) SetRequestedEnrollmentProfileId(value *string)() {
    m.requestedEnrollmentProfileId = value
}
// Sets the serialNumber property value. Device serial number
// Parameters:
//  - value : Value to set for the serialNumber property.
func (m *ImportedAppleDeviceIdentity) SetSerialNumber(value *string)() {
    m.serialNumber = value
}
