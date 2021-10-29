package graph

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// 
type ImportedDeviceIdentity struct {
    Entity
    // Created Date Time of the device
    createdDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time;
    // The description of the device
    description *string;
    // The state of the device in Intune. Possible values are: unknown, enrolled, pendingReset, failed, notContacted, blocked.
    enrollmentState *EnrollmentState;
    // Imported Device Identifier
    importedDeviceIdentifier *string;
    // Type of Imported Device Identity. Possible values are: unknown, imei, serialNumber.
    importedDeviceIdentityType *ImportedDeviceIdentityType;
    // Last Contacted Date Time of the device
    lastContactedDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time;
    // Last Modified DateTime of the description
    lastModifiedDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time;
    // The platform of the Device. Possible values are: unknown, ios, android, windows, windowsMobile, macOS.
    platform *Platform;
}
// Instantiates a new importedDeviceIdentity and sets the default values.
func NewImportedDeviceIdentity()(*ImportedDeviceIdentity) {
    m := &ImportedDeviceIdentity{
        Entity: *NewEntity(),
    }
    return m
}
// Gets the createdDateTime property value. Created Date Time of the device
func (m *ImportedDeviceIdentity) GetCreatedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.createdDateTime
    }
}
// Gets the description property value. The description of the device
func (m *ImportedDeviceIdentity) GetDescription()(*string) {
    if m == nil {
        return nil
    } else {
        return m.description
    }
}
// Gets the enrollmentState property value. The state of the device in Intune. Possible values are: unknown, enrolled, pendingReset, failed, notContacted, blocked.
func (m *ImportedDeviceIdentity) GetEnrollmentState()(*EnrollmentState) {
    if m == nil {
        return nil
    } else {
        return m.enrollmentState
    }
}
// Gets the importedDeviceIdentifier property value. Imported Device Identifier
func (m *ImportedDeviceIdentity) GetImportedDeviceIdentifier()(*string) {
    if m == nil {
        return nil
    } else {
        return m.importedDeviceIdentifier
    }
}
// Gets the importedDeviceIdentityType property value. Type of Imported Device Identity. Possible values are: unknown, imei, serialNumber.
func (m *ImportedDeviceIdentity) GetImportedDeviceIdentityType()(*ImportedDeviceIdentityType) {
    if m == nil {
        return nil
    } else {
        return m.importedDeviceIdentityType
    }
}
// Gets the lastContactedDateTime property value. Last Contacted Date Time of the device
func (m *ImportedDeviceIdentity) GetLastContactedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.lastContactedDateTime
    }
}
// Gets the lastModifiedDateTime property value. Last Modified DateTime of the description
func (m *ImportedDeviceIdentity) GetLastModifiedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.lastModifiedDateTime
    }
}
// Gets the platform property value. The platform of the Device. Possible values are: unknown, ios, android, windows, windowsMobile, macOS.
func (m *ImportedDeviceIdentity) GetPlatform()(*Platform) {
    if m == nil {
        return nil
    } else {
        return m.platform
    }
}
// The deserialization information for the current model
func (m *ImportedDeviceIdentity) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
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
    res["enrollmentState"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParseEnrollmentState)
        if err != nil {
            return err
        }
        cast := val.(EnrollmentState)
        m.SetEnrollmentState(&cast)
        return nil
    }
    res["importedDeviceIdentifier"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetImportedDeviceIdentifier(val)
        return nil
    }
    res["importedDeviceIdentityType"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParseImportedDeviceIdentityType)
        if err != nil {
            return err
        }
        cast := val.(ImportedDeviceIdentityType)
        m.SetImportedDeviceIdentityType(&cast)
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
    res["lastModifiedDateTime"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        m.SetLastModifiedDateTime(val)
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
    return res
}
func (m *ImportedDeviceIdentity) IsNil()(bool) {
    return m == nil
}
// Serializes information the current object
// Parameters:
//  - writer : Serialization writer to use to serialize this model
func (m *ImportedDeviceIdentity) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
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
    if m.GetEnrollmentState() != nil {
        cast := m.GetEnrollmentState().String()
        err = writer.WriteStringValue("enrollmentState", &cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("importedDeviceIdentifier", m.GetImportedDeviceIdentifier())
        if err != nil {
            return err
        }
    }
    if m.GetImportedDeviceIdentityType() != nil {
        cast := m.GetImportedDeviceIdentityType().String()
        err = writer.WriteStringValue("importedDeviceIdentityType", &cast)
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
    {
        err = writer.WriteTimeValue("lastModifiedDateTime", m.GetLastModifiedDateTime())
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
    return nil
}
// Sets the createdDateTime property value. Created Date Time of the device
// Parameters:
//  - value : Value to set for the createdDateTime property.
func (m *ImportedDeviceIdentity) SetCreatedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.createdDateTime = value
}
// Sets the description property value. The description of the device
// Parameters:
//  - value : Value to set for the description property.
func (m *ImportedDeviceIdentity) SetDescription(value *string)() {
    m.description = value
}
// Sets the enrollmentState property value. The state of the device in Intune. Possible values are: unknown, enrolled, pendingReset, failed, notContacted, blocked.
// Parameters:
//  - value : Value to set for the enrollmentState property.
func (m *ImportedDeviceIdentity) SetEnrollmentState(value *EnrollmentState)() {
    m.enrollmentState = value
}
// Sets the importedDeviceIdentifier property value. Imported Device Identifier
// Parameters:
//  - value : Value to set for the importedDeviceIdentifier property.
func (m *ImportedDeviceIdentity) SetImportedDeviceIdentifier(value *string)() {
    m.importedDeviceIdentifier = value
}
// Sets the importedDeviceIdentityType property value. Type of Imported Device Identity. Possible values are: unknown, imei, serialNumber.
// Parameters:
//  - value : Value to set for the importedDeviceIdentityType property.
func (m *ImportedDeviceIdentity) SetImportedDeviceIdentityType(value *ImportedDeviceIdentityType)() {
    m.importedDeviceIdentityType = value
}
// Sets the lastContactedDateTime property value. Last Contacted Date Time of the device
// Parameters:
//  - value : Value to set for the lastContactedDateTime property.
func (m *ImportedDeviceIdentity) SetLastContactedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.lastContactedDateTime = value
}
// Sets the lastModifiedDateTime property value. Last Modified DateTime of the description
// Parameters:
//  - value : Value to set for the lastModifiedDateTime property.
func (m *ImportedDeviceIdentity) SetLastModifiedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.lastModifiedDateTime = value
}
// Sets the platform property value. The platform of the Device. Possible values are: unknown, ios, android, windows, windowsMobile, macOS.
// Parameters:
//  - value : Value to set for the platform property.
func (m *ImportedDeviceIdentity) SetPlatform(value *Platform)() {
    m.platform = value
}
