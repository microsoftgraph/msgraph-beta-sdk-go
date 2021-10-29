package graph

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// 
type CloudPcDeviceImage struct {
    Entity
    // The image's display name.
    displayName *string;
    // The data and time that the image was last modified. The time is shown in ISO 8601 format and  Coordinated Universal Time (UTC) time. For example, midnight UTC on Jan 1, 2014 appears as '2014-01-01T00:00:00Z'.
    lastModifiedDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time;
    // The image's operating system. For example: Windows 10 Enterprise.
    operatingSystem *string;
    // The image's OS build version. For example: 1909.
    osBuildNumber *string;
    // The ID of the source image resource on Azure. Required format: '/subscriptions/{subscription-id}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/images/{imageName}'.
    sourceImageResourceId *string;
    // The status of the image on Cloud PC. Possible values are: pending, ready, failed.
    status *CloudPcDeviceImageStatus;
    // The details of the image's status, which indicates why the upload failed, if applicable. Possible values are: internalServerError, sourceImageNotFound, osVersionNotSupported, and sourceImageInvalid.
    statusDetails *CloudPcDeviceImageStatusDetails;
    // The image version. For example: 0.0.1, 1.5.13.
    version *string;
}
// Instantiates a new cloudPcDeviceImage and sets the default values.
func NewCloudPcDeviceImage()(*CloudPcDeviceImage) {
    m := &CloudPcDeviceImage{
        Entity: *NewEntity(),
    }
    return m
}
// Gets the displayName property value. The image's display name.
func (m *CloudPcDeviceImage) GetDisplayName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.displayName
    }
}
// Gets the lastModifiedDateTime property value. The data and time that the image was last modified. The time is shown in ISO 8601 format and  Coordinated Universal Time (UTC) time. For example, midnight UTC on Jan 1, 2014 appears as '2014-01-01T00:00:00Z'.
func (m *CloudPcDeviceImage) GetLastModifiedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.lastModifiedDateTime
    }
}
// Gets the operatingSystem property value. The image's operating system. For example: Windows 10 Enterprise.
func (m *CloudPcDeviceImage) GetOperatingSystem()(*string) {
    if m == nil {
        return nil
    } else {
        return m.operatingSystem
    }
}
// Gets the osBuildNumber property value. The image's OS build version. For example: 1909.
func (m *CloudPcDeviceImage) GetOsBuildNumber()(*string) {
    if m == nil {
        return nil
    } else {
        return m.osBuildNumber
    }
}
// Gets the sourceImageResourceId property value. The ID of the source image resource on Azure. Required format: '/subscriptions/{subscription-id}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/images/{imageName}'.
func (m *CloudPcDeviceImage) GetSourceImageResourceId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.sourceImageResourceId
    }
}
// Gets the status property value. The status of the image on Cloud PC. Possible values are: pending, ready, failed.
func (m *CloudPcDeviceImage) GetStatus()(*CloudPcDeviceImageStatus) {
    if m == nil {
        return nil
    } else {
        return m.status
    }
}
// Gets the statusDetails property value. The details of the image's status, which indicates why the upload failed, if applicable. Possible values are: internalServerError, sourceImageNotFound, osVersionNotSupported, and sourceImageInvalid.
func (m *CloudPcDeviceImage) GetStatusDetails()(*CloudPcDeviceImageStatusDetails) {
    if m == nil {
        return nil
    } else {
        return m.statusDetails
    }
}
// Gets the version property value. The image version. For example: 0.0.1, 1.5.13.
func (m *CloudPcDeviceImage) GetVersion()(*string) {
    if m == nil {
        return nil
    } else {
        return m.version
    }
}
// The deserialization information for the current model
func (m *CloudPcDeviceImage) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["displayName"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetDisplayName(val)
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
    res["operatingSystem"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetOperatingSystem(val)
        return nil
    }
    res["osBuildNumber"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetOsBuildNumber(val)
        return nil
    }
    res["sourceImageResourceId"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetSourceImageResourceId(val)
        return nil
    }
    res["status"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParseCloudPcDeviceImageStatus)
        if err != nil {
            return err
        }
        cast := val.(CloudPcDeviceImageStatus)
        m.SetStatus(&cast)
        return nil
    }
    res["statusDetails"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParseCloudPcDeviceImageStatusDetails)
        if err != nil {
            return err
        }
        cast := val.(CloudPcDeviceImageStatusDetails)
        m.SetStatusDetails(&cast)
        return nil
    }
    res["version"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetVersion(val)
        return nil
    }
    return res
}
func (m *CloudPcDeviceImage) IsNil()(bool) {
    return m == nil
}
// Serializes information the current object
// Parameters:
//  - writer : Serialization writer to use to serialize this model
func (m *CloudPcDeviceImage) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
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
    {
        err = writer.WriteTimeValue("lastModifiedDateTime", m.GetLastModifiedDateTime())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("operatingSystem", m.GetOperatingSystem())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("osBuildNumber", m.GetOsBuildNumber())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("sourceImageResourceId", m.GetSourceImageResourceId())
        if err != nil {
            return err
        }
    }
    if m.GetStatus() != nil {
        cast := m.GetStatus().String()
        err = writer.WriteStringValue("status", &cast)
        if err != nil {
            return err
        }
    }
    if m.GetStatusDetails() != nil {
        cast := m.GetStatusDetails().String()
        err = writer.WriteStringValue("statusDetails", &cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("version", m.GetVersion())
        if err != nil {
            return err
        }
    }
    return nil
}
// Sets the displayName property value. The image's display name.
// Parameters:
//  - value : Value to set for the displayName property.
func (m *CloudPcDeviceImage) SetDisplayName(value *string)() {
    m.displayName = value
}
// Sets the lastModifiedDateTime property value. The data and time that the image was last modified. The time is shown in ISO 8601 format and  Coordinated Universal Time (UTC) time. For example, midnight UTC on Jan 1, 2014 appears as '2014-01-01T00:00:00Z'.
// Parameters:
//  - value : Value to set for the lastModifiedDateTime property.
func (m *CloudPcDeviceImage) SetLastModifiedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.lastModifiedDateTime = value
}
// Sets the operatingSystem property value. The image's operating system. For example: Windows 10 Enterprise.
// Parameters:
//  - value : Value to set for the operatingSystem property.
func (m *CloudPcDeviceImage) SetOperatingSystem(value *string)() {
    m.operatingSystem = value
}
// Sets the osBuildNumber property value. The image's OS build version. For example: 1909.
// Parameters:
//  - value : Value to set for the osBuildNumber property.
func (m *CloudPcDeviceImage) SetOsBuildNumber(value *string)() {
    m.osBuildNumber = value
}
// Sets the sourceImageResourceId property value. The ID of the source image resource on Azure. Required format: '/subscriptions/{subscription-id}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/images/{imageName}'.
// Parameters:
//  - value : Value to set for the sourceImageResourceId property.
func (m *CloudPcDeviceImage) SetSourceImageResourceId(value *string)() {
    m.sourceImageResourceId = value
}
// Sets the status property value. The status of the image on Cloud PC. Possible values are: pending, ready, failed.
// Parameters:
//  - value : Value to set for the status property.
func (m *CloudPcDeviceImage) SetStatus(value *CloudPcDeviceImageStatus)() {
    m.status = value
}
// Sets the statusDetails property value. The details of the image's status, which indicates why the upload failed, if applicable. Possible values are: internalServerError, sourceImageNotFound, osVersionNotSupported, and sourceImageInvalid.
// Parameters:
//  - value : Value to set for the statusDetails property.
func (m *CloudPcDeviceImage) SetStatusDetails(value *CloudPcDeviceImageStatusDetails)() {
    m.statusDetails = value
}
// Sets the version property value. The image version. For example: 0.0.1, 1.5.13.
// Parameters:
//  - value : Value to set for the version property.
func (m *CloudPcDeviceImage) SetVersion(value *string)() {
    m.version = value
}
