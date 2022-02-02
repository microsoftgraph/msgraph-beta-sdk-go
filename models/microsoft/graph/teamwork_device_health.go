package graph

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// TeamworkDeviceHealth 
type TeamworkDeviceHealth struct {
    Entity
    // 
    connection *TeamworkConnection;
    // Identity of the user who created the device health document.
    createdBy *IdentitySet;
    // The UTC date and time when the device health document was created.
    createdDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time;
    // Health details about the device hardware.
    hardwareHealth *TeamworkHardwareHealth;
    // Identity of the user who last modified the device health details.
    lastModifiedBy *IdentitySet;
    // The UTC date and time when the device health detail was last modified.
    lastModifiedDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time;
    // The login status of Microsoft Teams, Skype for Business, and Exchange.
    loginStatus *TeamworkLoginStatus;
    // Health details about all peripherals (for example, speaker and microphone) attached to a device.
    peripheralsHealth *TeamworkPeripheralsHealth;
    // Software updates available for the device.
    softwareUpdateHealth *TeamworkSoftwareUpdateHealth;
}
// NewTeamworkDeviceHealth instantiates a new teamworkDeviceHealth and sets the default values.
func NewTeamworkDeviceHealth()(*TeamworkDeviceHealth) {
    m := &TeamworkDeviceHealth{
        Entity: *NewEntity(),
    }
    return m
}
// GetConnection gets the connection property value. 
func (m *TeamworkDeviceHealth) GetConnection()(*TeamworkConnection) {
    if m == nil {
        return nil
    } else {
        return m.connection
    }
}
// GetCreatedBy gets the createdBy property value. Identity of the user who created the device health document.
func (m *TeamworkDeviceHealth) GetCreatedBy()(*IdentitySet) {
    if m == nil {
        return nil
    } else {
        return m.createdBy
    }
}
// GetCreatedDateTime gets the createdDateTime property value. The UTC date and time when the device health document was created.
func (m *TeamworkDeviceHealth) GetCreatedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.createdDateTime
    }
}
// GetHardwareHealth gets the hardwareHealth property value. Health details about the device hardware.
func (m *TeamworkDeviceHealth) GetHardwareHealth()(*TeamworkHardwareHealth) {
    if m == nil {
        return nil
    } else {
        return m.hardwareHealth
    }
}
// GetLastModifiedBy gets the lastModifiedBy property value. Identity of the user who last modified the device health details.
func (m *TeamworkDeviceHealth) GetLastModifiedBy()(*IdentitySet) {
    if m == nil {
        return nil
    } else {
        return m.lastModifiedBy
    }
}
// GetLastModifiedDateTime gets the lastModifiedDateTime property value. The UTC date and time when the device health detail was last modified.
func (m *TeamworkDeviceHealth) GetLastModifiedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.lastModifiedDateTime
    }
}
// GetLoginStatus gets the loginStatus property value. The login status of Microsoft Teams, Skype for Business, and Exchange.
func (m *TeamworkDeviceHealth) GetLoginStatus()(*TeamworkLoginStatus) {
    if m == nil {
        return nil
    } else {
        return m.loginStatus
    }
}
// GetPeripheralsHealth gets the peripheralsHealth property value. Health details about all peripherals (for example, speaker and microphone) attached to a device.
func (m *TeamworkDeviceHealth) GetPeripheralsHealth()(*TeamworkPeripheralsHealth) {
    if m == nil {
        return nil
    } else {
        return m.peripheralsHealth
    }
}
// GetSoftwareUpdateHealth gets the softwareUpdateHealth property value. Software updates available for the device.
func (m *TeamworkDeviceHealth) GetSoftwareUpdateHealth()(*TeamworkSoftwareUpdateHealth) {
    if m == nil {
        return nil
    } else {
        return m.softwareUpdateHealth
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *TeamworkDeviceHealth) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["connection"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewTeamworkConnection() })
        if err != nil {
            return err
        }
        if val != nil {
            m.SetConnection(val.(*TeamworkConnection))
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
    res["hardwareHealth"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewTeamworkHardwareHealth() })
        if err != nil {
            return err
        }
        if val != nil {
            m.SetHardwareHealth(val.(*TeamworkHardwareHealth))
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
    res["loginStatus"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewTeamworkLoginStatus() })
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLoginStatus(val.(*TeamworkLoginStatus))
        }
        return nil
    }
    res["peripheralsHealth"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewTeamworkPeripheralsHealth() })
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPeripheralsHealth(val.(*TeamworkPeripheralsHealth))
        }
        return nil
    }
    res["softwareUpdateHealth"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewTeamworkSoftwareUpdateHealth() })
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSoftwareUpdateHealth(val.(*TeamworkSoftwareUpdateHealth))
        }
        return nil
    }
    return res
}
func (m *TeamworkDeviceHealth) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
func (m *TeamworkDeviceHealth) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteObjectValue("connection", m.GetConnection())
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
        err = writer.WriteObjectValue("hardwareHealth", m.GetHardwareHealth())
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
        err = writer.WriteObjectValue("loginStatus", m.GetLoginStatus())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("peripheralsHealth", m.GetPeripheralsHealth())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("softwareUpdateHealth", m.GetSoftwareUpdateHealth())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetConnection sets the connection property value. 
func (m *TeamworkDeviceHealth) SetConnection(value *TeamworkConnection)() {
    if m != nil {
        m.connection = value
    }
}
// SetCreatedBy sets the createdBy property value. Identity of the user who created the device health document.
func (m *TeamworkDeviceHealth) SetCreatedBy(value *IdentitySet)() {
    if m != nil {
        m.createdBy = value
    }
}
// SetCreatedDateTime sets the createdDateTime property value. The UTC date and time when the device health document was created.
func (m *TeamworkDeviceHealth) SetCreatedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    if m != nil {
        m.createdDateTime = value
    }
}
// SetHardwareHealth sets the hardwareHealth property value. Health details about the device hardware.
func (m *TeamworkDeviceHealth) SetHardwareHealth(value *TeamworkHardwareHealth)() {
    if m != nil {
        m.hardwareHealth = value
    }
}
// SetLastModifiedBy sets the lastModifiedBy property value. Identity of the user who last modified the device health details.
func (m *TeamworkDeviceHealth) SetLastModifiedBy(value *IdentitySet)() {
    if m != nil {
        m.lastModifiedBy = value
    }
}
// SetLastModifiedDateTime sets the lastModifiedDateTime property value. The UTC date and time when the device health detail was last modified.
func (m *TeamworkDeviceHealth) SetLastModifiedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    if m != nil {
        m.lastModifiedDateTime = value
    }
}
// SetLoginStatus sets the loginStatus property value. The login status of Microsoft Teams, Skype for Business, and Exchange.
func (m *TeamworkDeviceHealth) SetLoginStatus(value *TeamworkLoginStatus)() {
    if m != nil {
        m.loginStatus = value
    }
}
// SetPeripheralsHealth sets the peripheralsHealth property value. Health details about all peripherals (for example, speaker and microphone) attached to a device.
func (m *TeamworkDeviceHealth) SetPeripheralsHealth(value *TeamworkPeripheralsHealth)() {
    if m != nil {
        m.peripheralsHealth = value
    }
}
// SetSoftwareUpdateHealth sets the softwareUpdateHealth property value. Software updates available for the device.
func (m *TeamworkDeviceHealth) SetSoftwareUpdateHealth(value *TeamworkSoftwareUpdateHealth)() {
    if m != nil {
        m.softwareUpdateHealth = value
    }
}
