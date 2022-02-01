package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// TeamworkSoftwareUpdateHealth 
type TeamworkSoftwareUpdateHealth struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // The software update available for the admin agent.
    adminAgentSoftwareUpdateStatus *TeamworkSoftwareUpdateStatus;
    // The software update available for the company portal.
    companyPortalSoftwareUpdateStatus *TeamworkSoftwareUpdateStatus;
    // The software update available for the firmware.
    firmwareSoftwareUpdateStatus *TeamworkSoftwareUpdateStatus;
    // The software update available for the operating system.
    operatingSystemSoftwareUpdateStatus *TeamworkSoftwareUpdateStatus;
    // The software update available for the partner agent.
    partnerAgentSoftwareUpdateStatus *TeamworkSoftwareUpdateStatus;
    // The software update available for the Teams client.
    teamsClientSoftwareUpdateStatus *TeamworkSoftwareUpdateStatus;
}
// NewTeamworkSoftwareUpdateHealth instantiates a new teamworkSoftwareUpdateHealth and sets the default values.
func NewTeamworkSoftwareUpdateHealth()(*TeamworkSoftwareUpdateHealth) {
    m := &TeamworkSoftwareUpdateHealth{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *TeamworkSoftwareUpdateHealth) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetAdminAgentSoftwareUpdateStatus gets the adminAgentSoftwareUpdateStatus property value. The software update available for the admin agent.
func (m *TeamworkSoftwareUpdateHealth) GetAdminAgentSoftwareUpdateStatus()(*TeamworkSoftwareUpdateStatus) {
    if m == nil {
        return nil
    } else {
        return m.adminAgentSoftwareUpdateStatus
    }
}
// GetCompanyPortalSoftwareUpdateStatus gets the companyPortalSoftwareUpdateStatus property value. The software update available for the company portal.
func (m *TeamworkSoftwareUpdateHealth) GetCompanyPortalSoftwareUpdateStatus()(*TeamworkSoftwareUpdateStatus) {
    if m == nil {
        return nil
    } else {
        return m.companyPortalSoftwareUpdateStatus
    }
}
// GetFirmwareSoftwareUpdateStatus gets the firmwareSoftwareUpdateStatus property value. The software update available for the firmware.
func (m *TeamworkSoftwareUpdateHealth) GetFirmwareSoftwareUpdateStatus()(*TeamworkSoftwareUpdateStatus) {
    if m == nil {
        return nil
    } else {
        return m.firmwareSoftwareUpdateStatus
    }
}
// GetOperatingSystemSoftwareUpdateStatus gets the operatingSystemSoftwareUpdateStatus property value. The software update available for the operating system.
func (m *TeamworkSoftwareUpdateHealth) GetOperatingSystemSoftwareUpdateStatus()(*TeamworkSoftwareUpdateStatus) {
    if m == nil {
        return nil
    } else {
        return m.operatingSystemSoftwareUpdateStatus
    }
}
// GetPartnerAgentSoftwareUpdateStatus gets the partnerAgentSoftwareUpdateStatus property value. The software update available for the partner agent.
func (m *TeamworkSoftwareUpdateHealth) GetPartnerAgentSoftwareUpdateStatus()(*TeamworkSoftwareUpdateStatus) {
    if m == nil {
        return nil
    } else {
        return m.partnerAgentSoftwareUpdateStatus
    }
}
// GetTeamsClientSoftwareUpdateStatus gets the teamsClientSoftwareUpdateStatus property value. The software update available for the Teams client.
func (m *TeamworkSoftwareUpdateHealth) GetTeamsClientSoftwareUpdateStatus()(*TeamworkSoftwareUpdateStatus) {
    if m == nil {
        return nil
    } else {
        return m.teamsClientSoftwareUpdateStatus
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *TeamworkSoftwareUpdateHealth) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["adminAgentSoftwareUpdateStatus"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewTeamworkSoftwareUpdateStatus() })
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAdminAgentSoftwareUpdateStatus(val.(*TeamworkSoftwareUpdateStatus))
        }
        return nil
    }
    res["companyPortalSoftwareUpdateStatus"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewTeamworkSoftwareUpdateStatus() })
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCompanyPortalSoftwareUpdateStatus(val.(*TeamworkSoftwareUpdateStatus))
        }
        return nil
    }
    res["firmwareSoftwareUpdateStatus"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewTeamworkSoftwareUpdateStatus() })
        if err != nil {
            return err
        }
        if val != nil {
            m.SetFirmwareSoftwareUpdateStatus(val.(*TeamworkSoftwareUpdateStatus))
        }
        return nil
    }
    res["operatingSystemSoftwareUpdateStatus"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewTeamworkSoftwareUpdateStatus() })
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOperatingSystemSoftwareUpdateStatus(val.(*TeamworkSoftwareUpdateStatus))
        }
        return nil
    }
    res["partnerAgentSoftwareUpdateStatus"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewTeamworkSoftwareUpdateStatus() })
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPartnerAgentSoftwareUpdateStatus(val.(*TeamworkSoftwareUpdateStatus))
        }
        return nil
    }
    res["teamsClientSoftwareUpdateStatus"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewTeamworkSoftwareUpdateStatus() })
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTeamsClientSoftwareUpdateStatus(val.(*TeamworkSoftwareUpdateStatus))
        }
        return nil
    }
    return res
}
func (m *TeamworkSoftwareUpdateHealth) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
func (m *TeamworkSoftwareUpdateHealth) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    {
        err := writer.WriteObjectValue("adminAgentSoftwareUpdateStatus", m.GetAdminAgentSoftwareUpdateStatus())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("companyPortalSoftwareUpdateStatus", m.GetCompanyPortalSoftwareUpdateStatus())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("firmwareSoftwareUpdateStatus", m.GetFirmwareSoftwareUpdateStatus())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("operatingSystemSoftwareUpdateStatus", m.GetOperatingSystemSoftwareUpdateStatus())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("partnerAgentSoftwareUpdateStatus", m.GetPartnerAgentSoftwareUpdateStatus())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("teamsClientSoftwareUpdateStatus", m.GetTeamsClientSoftwareUpdateStatus())
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
func (m *TeamworkSoftwareUpdateHealth) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
// SetAdminAgentSoftwareUpdateStatus sets the adminAgentSoftwareUpdateStatus property value. The software update available for the admin agent.
func (m *TeamworkSoftwareUpdateHealth) SetAdminAgentSoftwareUpdateStatus(value *TeamworkSoftwareUpdateStatus)() {
    if m != nil {
        m.adminAgentSoftwareUpdateStatus = value
    }
}
// SetCompanyPortalSoftwareUpdateStatus sets the companyPortalSoftwareUpdateStatus property value. The software update available for the company portal.
func (m *TeamworkSoftwareUpdateHealth) SetCompanyPortalSoftwareUpdateStatus(value *TeamworkSoftwareUpdateStatus)() {
    if m != nil {
        m.companyPortalSoftwareUpdateStatus = value
    }
}
// SetFirmwareSoftwareUpdateStatus sets the firmwareSoftwareUpdateStatus property value. The software update available for the firmware.
func (m *TeamworkSoftwareUpdateHealth) SetFirmwareSoftwareUpdateStatus(value *TeamworkSoftwareUpdateStatus)() {
    if m != nil {
        m.firmwareSoftwareUpdateStatus = value
    }
}
// SetOperatingSystemSoftwareUpdateStatus sets the operatingSystemSoftwareUpdateStatus property value. The software update available for the operating system.
func (m *TeamworkSoftwareUpdateHealth) SetOperatingSystemSoftwareUpdateStatus(value *TeamworkSoftwareUpdateStatus)() {
    if m != nil {
        m.operatingSystemSoftwareUpdateStatus = value
    }
}
// SetPartnerAgentSoftwareUpdateStatus sets the partnerAgentSoftwareUpdateStatus property value. The software update available for the partner agent.
func (m *TeamworkSoftwareUpdateHealth) SetPartnerAgentSoftwareUpdateStatus(value *TeamworkSoftwareUpdateStatus)() {
    if m != nil {
        m.partnerAgentSoftwareUpdateStatus = value
    }
}
// SetTeamsClientSoftwareUpdateStatus sets the teamsClientSoftwareUpdateStatus property value. The software update available for the Teams client.
func (m *TeamworkSoftwareUpdateHealth) SetTeamsClientSoftwareUpdateStatus(value *TeamworkSoftwareUpdateStatus)() {
    if m != nil {
        m.teamsClientSoftwareUpdateStatus = value
    }
}
