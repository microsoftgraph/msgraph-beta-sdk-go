package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// TeamworkDeviceSoftwareVersions 
type TeamworkDeviceSoftwareVersions struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // The software version for the admin agent running on the device.
    adminAgentSoftwareVersion *string;
    // The software version for the firmware running on the device.
    firmwareSoftwareVersion *string;
    // The software version for the operating system on the device.
    operatingSystemSoftwareVersion *string;
    // The software version for the partner agent running on the device.
    partnerAgentSoftwareVersion *string;
    // The software version for the Teams client running on the device.
    teamsClientSoftwareVersion *string;
}
// NewTeamworkDeviceSoftwareVersions instantiates a new teamworkDeviceSoftwareVersions and sets the default values.
func NewTeamworkDeviceSoftwareVersions()(*TeamworkDeviceSoftwareVersions) {
    m := &TeamworkDeviceSoftwareVersions{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *TeamworkDeviceSoftwareVersions) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetAdminAgentSoftwareVersion gets the adminAgentSoftwareVersion property value. The software version for the admin agent running on the device.
func (m *TeamworkDeviceSoftwareVersions) GetAdminAgentSoftwareVersion()(*string) {
    if m == nil {
        return nil
    } else {
        return m.adminAgentSoftwareVersion
    }
}
// GetFirmwareSoftwareVersion gets the firmwareSoftwareVersion property value. The software version for the firmware running on the device.
func (m *TeamworkDeviceSoftwareVersions) GetFirmwareSoftwareVersion()(*string) {
    if m == nil {
        return nil
    } else {
        return m.firmwareSoftwareVersion
    }
}
// GetOperatingSystemSoftwareVersion gets the operatingSystemSoftwareVersion property value. The software version for the operating system on the device.
func (m *TeamworkDeviceSoftwareVersions) GetOperatingSystemSoftwareVersion()(*string) {
    if m == nil {
        return nil
    } else {
        return m.operatingSystemSoftwareVersion
    }
}
// GetPartnerAgentSoftwareVersion gets the partnerAgentSoftwareVersion property value. The software version for the partner agent running on the device.
func (m *TeamworkDeviceSoftwareVersions) GetPartnerAgentSoftwareVersion()(*string) {
    if m == nil {
        return nil
    } else {
        return m.partnerAgentSoftwareVersion
    }
}
// GetTeamsClientSoftwareVersion gets the teamsClientSoftwareVersion property value. The software version for the Teams client running on the device.
func (m *TeamworkDeviceSoftwareVersions) GetTeamsClientSoftwareVersion()(*string) {
    if m == nil {
        return nil
    } else {
        return m.teamsClientSoftwareVersion
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *TeamworkDeviceSoftwareVersions) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["adminAgentSoftwareVersion"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAdminAgentSoftwareVersion(val)
        }
        return nil
    }
    res["firmwareSoftwareVersion"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetFirmwareSoftwareVersion(val)
        }
        return nil
    }
    res["operatingSystemSoftwareVersion"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOperatingSystemSoftwareVersion(val)
        }
        return nil
    }
    res["partnerAgentSoftwareVersion"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPartnerAgentSoftwareVersion(val)
        }
        return nil
    }
    res["teamsClientSoftwareVersion"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTeamsClientSoftwareVersion(val)
        }
        return nil
    }
    return res
}
func (m *TeamworkDeviceSoftwareVersions) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
func (m *TeamworkDeviceSoftwareVersions) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("adminAgentSoftwareVersion", m.GetAdminAgentSoftwareVersion())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("firmwareSoftwareVersion", m.GetFirmwareSoftwareVersion())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("operatingSystemSoftwareVersion", m.GetOperatingSystemSoftwareVersion())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("partnerAgentSoftwareVersion", m.GetPartnerAgentSoftwareVersion())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("teamsClientSoftwareVersion", m.GetTeamsClientSoftwareVersion())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteAdditionalData(m.GetAdditionalData())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAdditionalData sets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *TeamworkDeviceSoftwareVersions) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
// SetAdminAgentSoftwareVersion sets the adminAgentSoftwareVersion property value. The software version for the admin agent running on the device.
func (m *TeamworkDeviceSoftwareVersions) SetAdminAgentSoftwareVersion(value *string)() {
    if m != nil {
        m.adminAgentSoftwareVersion = value
    }
}
// SetFirmwareSoftwareVersion sets the firmwareSoftwareVersion property value. The software version for the firmware running on the device.
func (m *TeamworkDeviceSoftwareVersions) SetFirmwareSoftwareVersion(value *string)() {
    if m != nil {
        m.firmwareSoftwareVersion = value
    }
}
// SetOperatingSystemSoftwareVersion sets the operatingSystemSoftwareVersion property value. The software version for the operating system on the device.
func (m *TeamworkDeviceSoftwareVersions) SetOperatingSystemSoftwareVersion(value *string)() {
    if m != nil {
        m.operatingSystemSoftwareVersion = value
    }
}
// SetPartnerAgentSoftwareVersion sets the partnerAgentSoftwareVersion property value. The software version for the partner agent running on the device.
func (m *TeamworkDeviceSoftwareVersions) SetPartnerAgentSoftwareVersion(value *string)() {
    if m != nil {
        m.partnerAgentSoftwareVersion = value
    }
}
// SetTeamsClientSoftwareVersion sets the teamsClientSoftwareVersion property value. The software version for the Teams client running on the device.
func (m *TeamworkDeviceSoftwareVersions) SetTeamsClientSoftwareVersion(value *string)() {
    if m != nil {
        m.teamsClientSoftwareVersion = value
    }
}
