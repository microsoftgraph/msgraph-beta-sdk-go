package graph

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// ImportedDeviceIdentity 
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
// NewImportedDeviceIdentity instantiates a new importedDeviceIdentity and sets the default values.
func NewImportedDeviceIdentity()(*ImportedDeviceIdentity) {
    m := &ImportedDeviceIdentity{
        Entity: *NewEntity(),
    }
    return m
}
// GetCreatedDateTime gets the createdDateTime property value. Created Date Time of the device
func (m *ImportedDeviceIdentity) GetCreatedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.createdDateTime
    }
}
// GetDescription gets the description property value. The description of the device
func (m *ImportedDeviceIdentity) GetDescription()(*string) {
    if m == nil {
        return nil
    } else {
        return m.description
    }
}
// GetEnrollmentState gets the enrollmentState property value. The state of the device in Intune. Possible values are: unknown, enrolled, pendingReset, failed, notContacted, blocked.
func (m *ImportedDeviceIdentity) GetEnrollmentState()(*EnrollmentState) {
    if m == nil {
        return nil
    } else {
        return m.enrollmentState
    }
}
// GetImportedDeviceIdentifier gets the importedDeviceIdentifier property value. Imported Device Identifier
func (m *ImportedDeviceIdentity) GetImportedDeviceIdentifier()(*string) {
    if m == nil {
        return nil
    } else {
        return m.importedDeviceIdentifier
    }
}
// GetImportedDeviceIdentityType gets the importedDeviceIdentityType property value. Type of Imported Device Identity. Possible values are: unknown, imei, serialNumber.
func (m *ImportedDeviceIdentity) GetImportedDeviceIdentityType()(*ImportedDeviceIdentityType) {
    if m == nil {
        return nil
    } else {
        return m.importedDeviceIdentityType
    }
}
// GetLastContactedDateTime gets the lastContactedDateTime property value. Last Contacted Date Time of the device
func (m *ImportedDeviceIdentity) GetLastContactedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.lastContactedDateTime
    }
}
// GetLastModifiedDateTime gets the lastModifiedDateTime property value. Last Modified DateTime of the description
func (m *ImportedDeviceIdentity) GetLastModifiedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.lastModifiedDateTime
    }
}
// GetPlatform gets the platform property value. The platform of the Device. Possible values are: unknown, ios, android, windows, windowsMobile, macOS.
func (m *ImportedDeviceIdentity) GetPlatform()(*Platform) {
    if m == nil {
        return nil
    } else {
        return m.platform
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ImportedDeviceIdentity) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
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
    res["enrollmentState"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParseEnrollmentState)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetEnrollmentState(val.(*EnrollmentState))
        }
        return nil
    }
    res["importedDeviceIdentifier"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetImportedDeviceIdentifier(val)
        }
        return nil
    }
    res["importedDeviceIdentityType"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParseImportedDeviceIdentityType)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetImportedDeviceIdentityType(val.(*ImportedDeviceIdentityType))
        }
        return nil
    }
    res["lastContactedDateTime"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLastContactedDateTime(val)
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
    res["platform"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParsePlatform)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPlatform(val.(*Platform))
        }
        return nil
    }
    return res
}
func (m *ImportedDeviceIdentity) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
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
        cast := (*m.GetEnrollmentState()).String()
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
        cast := (*m.GetImportedDeviceIdentityType()).String()
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
        cast := (*m.GetPlatform()).String()
        err = writer.WriteStringValue("platform", &cast)
        if err != nil {
            return err
        }
    }
    return nil
}
// SetCreatedDateTime sets the createdDateTime property value. Created Date Time of the device
func (m *ImportedDeviceIdentity) SetCreatedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    if m != nil {
        m.createdDateTime = value
    }
}
// SetDescription sets the description property value. The description of the device
func (m *ImportedDeviceIdentity) SetDescription(value *string)() {
    if m != nil {
        m.description = value
    }
}
// SetEnrollmentState sets the enrollmentState property value. The state of the device in Intune. Possible values are: unknown, enrolled, pendingReset, failed, notContacted, blocked.
func (m *ImportedDeviceIdentity) SetEnrollmentState(value *EnrollmentState)() {
    if m != nil {
        m.enrollmentState = value
    }
}
// SetImportedDeviceIdentifier sets the importedDeviceIdentifier property value. Imported Device Identifier
func (m *ImportedDeviceIdentity) SetImportedDeviceIdentifier(value *string)() {
    if m != nil {
        m.importedDeviceIdentifier = value
    }
}
// SetImportedDeviceIdentityType sets the importedDeviceIdentityType property value. Type of Imported Device Identity. Possible values are: unknown, imei, serialNumber.
func (m *ImportedDeviceIdentity) SetImportedDeviceIdentityType(value *ImportedDeviceIdentityType)() {
    if m != nil {
        m.importedDeviceIdentityType = value
    }
}
// SetLastContactedDateTime sets the lastContactedDateTime property value. Last Contacted Date Time of the device
func (m *ImportedDeviceIdentity) SetLastContactedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    if m != nil {
        m.lastContactedDateTime = value
    }
}
// SetLastModifiedDateTime sets the lastModifiedDateTime property value. Last Modified DateTime of the description
func (m *ImportedDeviceIdentity) SetLastModifiedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    if m != nil {
        m.lastModifiedDateTime = value
    }
}
// SetPlatform sets the platform property value. The platform of the Device. Possible values are: unknown, ios, android, windows, windowsMobile, macOS.
func (m *ImportedDeviceIdentity) SetPlatform(value *Platform)() {
    if m != nil {
        m.platform = value
    }
}
