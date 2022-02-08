package graph

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// WindowsDriverUpdateInventory 
type WindowsDriverUpdateInventory struct {
    Entity
    // The number of devices for which this driver is applicable.
    applicableDeviceCount *int32;
    // The approval status for this driver. Possible values are: needsReview, declined, approved, suspended.
    approvalStatus *DriverApprovalStatus;
    // The category for this driver. Possible values are: recommended, previouslyApproved, other.
    category *DriverCategory;
    // The date time when a driver should be deployed if approvalStatus is approved.
    deployDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time;
    // The class of the driver.
    driverClass *string;
    // The manufacturer of the driver.
    manufacturer *string;
    // The name of the driver.
    name *string;
    // The release date time of the driver.
    releaseDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time;
    // The version of the driver.
    version *string;
}
// NewWindowsDriverUpdateInventory instantiates a new windowsDriverUpdateInventory and sets the default values.
func NewWindowsDriverUpdateInventory()(*WindowsDriverUpdateInventory) {
    m := &WindowsDriverUpdateInventory{
        Entity: *NewEntity(),
    }
    return m
}
// GetApplicableDeviceCount gets the applicableDeviceCount property value. The number of devices for which this driver is applicable.
func (m *WindowsDriverUpdateInventory) GetApplicableDeviceCount()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.applicableDeviceCount
    }
}
// GetApprovalStatus gets the approvalStatus property value. The approval status for this driver. Possible values are: needsReview, declined, approved, suspended.
func (m *WindowsDriverUpdateInventory) GetApprovalStatus()(*DriverApprovalStatus) {
    if m == nil {
        return nil
    } else {
        return m.approvalStatus
    }
}
// GetCategory gets the category property value. The category for this driver. Possible values are: recommended, previouslyApproved, other.
func (m *WindowsDriverUpdateInventory) GetCategory()(*DriverCategory) {
    if m == nil {
        return nil
    } else {
        return m.category
    }
}
// GetDeployDateTime gets the deployDateTime property value. The date time when a driver should be deployed if approvalStatus is approved.
func (m *WindowsDriverUpdateInventory) GetDeployDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.deployDateTime
    }
}
// GetDriverClass gets the driverClass property value. The class of the driver.
func (m *WindowsDriverUpdateInventory) GetDriverClass()(*string) {
    if m == nil {
        return nil
    } else {
        return m.driverClass
    }
}
// GetManufacturer gets the manufacturer property value. The manufacturer of the driver.
func (m *WindowsDriverUpdateInventory) GetManufacturer()(*string) {
    if m == nil {
        return nil
    } else {
        return m.manufacturer
    }
}
// GetName gets the name property value. The name of the driver.
func (m *WindowsDriverUpdateInventory) GetName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.name
    }
}
// GetReleaseDateTime gets the releaseDateTime property value. The release date time of the driver.
func (m *WindowsDriverUpdateInventory) GetReleaseDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.releaseDateTime
    }
}
// GetVersion gets the version property value. The version of the driver.
func (m *WindowsDriverUpdateInventory) GetVersion()(*string) {
    if m == nil {
        return nil
    } else {
        return m.version
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *WindowsDriverUpdateInventory) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["applicableDeviceCount"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetApplicableDeviceCount(val)
        }
        return nil
    }
    res["approvalStatus"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParseDriverApprovalStatus)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetApprovalStatus(val.(*DriverApprovalStatus))
        }
        return nil
    }
    res["category"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParseDriverCategory)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCategory(val.(*DriverCategory))
        }
        return nil
    }
    res["deployDateTime"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDeployDateTime(val)
        }
        return nil
    }
    res["driverClass"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDriverClass(val)
        }
        return nil
    }
    res["manufacturer"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetManufacturer(val)
        }
        return nil
    }
    res["name"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetName(val)
        }
        return nil
    }
    res["releaseDateTime"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetReleaseDateTime(val)
        }
        return nil
    }
    res["version"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
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
func (m *WindowsDriverUpdateInventory) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
func (m *WindowsDriverUpdateInventory) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteInt32Value("applicableDeviceCount", m.GetApplicableDeviceCount())
        if err != nil {
            return err
        }
    }
    if m.GetApprovalStatus() != nil {
        cast := (*m.GetApprovalStatus()).String()
        err = writer.WriteStringValue("approvalStatus", &cast)
        if err != nil {
            return err
        }
    }
    if m.GetCategory() != nil {
        cast := (*m.GetCategory()).String()
        err = writer.WriteStringValue("category", &cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteTimeValue("deployDateTime", m.GetDeployDateTime())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("driverClass", m.GetDriverClass())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("manufacturer", m.GetManufacturer())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("name", m.GetName())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteTimeValue("releaseDateTime", m.GetReleaseDateTime())
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
// SetApplicableDeviceCount sets the applicableDeviceCount property value. The number of devices for which this driver is applicable.
func (m *WindowsDriverUpdateInventory) SetApplicableDeviceCount(value *int32)() {
    if m != nil {
        m.applicableDeviceCount = value
    }
}
// SetApprovalStatus sets the approvalStatus property value. The approval status for this driver. Possible values are: needsReview, declined, approved, suspended.
func (m *WindowsDriverUpdateInventory) SetApprovalStatus(value *DriverApprovalStatus)() {
    if m != nil {
        m.approvalStatus = value
    }
}
// SetCategory sets the category property value. The category for this driver. Possible values are: recommended, previouslyApproved, other.
func (m *WindowsDriverUpdateInventory) SetCategory(value *DriverCategory)() {
    if m != nil {
        m.category = value
    }
}
// SetDeployDateTime sets the deployDateTime property value. The date time when a driver should be deployed if approvalStatus is approved.
func (m *WindowsDriverUpdateInventory) SetDeployDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    if m != nil {
        m.deployDateTime = value
    }
}
// SetDriverClass sets the driverClass property value. The class of the driver.
func (m *WindowsDriverUpdateInventory) SetDriverClass(value *string)() {
    if m != nil {
        m.driverClass = value
    }
}
// SetManufacturer sets the manufacturer property value. The manufacturer of the driver.
func (m *WindowsDriverUpdateInventory) SetManufacturer(value *string)() {
    if m != nil {
        m.manufacturer = value
    }
}
// SetName sets the name property value. The name of the driver.
func (m *WindowsDriverUpdateInventory) SetName(value *string)() {
    if m != nil {
        m.name = value
    }
}
// SetReleaseDateTime sets the releaseDateTime property value. The release date time of the driver.
func (m *WindowsDriverUpdateInventory) SetReleaseDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    if m != nil {
        m.releaseDateTime = value
    }
}
// SetVersion sets the version property value. The version of the driver.
func (m *WindowsDriverUpdateInventory) SetVersion(value *string)() {
    if m != nil {
        m.version = value
    }
}
