package graph

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// 
type MacOSSoftwareUpdateCategorySummary struct {
    Entity
    // The device ID.
    deviceId *string;
    // The name of the report
    displayName *string;
    // Number of failed updates on the device
    failedUpdateCount *int32;
    // Last date time the report for this device was updated.
    lastUpdatedDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time;
    // Number of successful updates on the device
    successfulUpdateCount *int32;
    // Number of total updates on the device
    totalUpdateCount *int32;
    // Software update type. Possible values are: critical, configurationDataFile, firmware, other.
    updateCategory *MacOSSoftwareUpdateCategory;
    // Summary of the update states.
    updateStateSummaries []MacOSSoftwareUpdateStateSummary;
    // The user ID.
    userId *string;
}
// Instantiates a new macOSSoftwareUpdateCategorySummary and sets the default values.
func NewMacOSSoftwareUpdateCategorySummary()(*MacOSSoftwareUpdateCategorySummary) {
    m := &MacOSSoftwareUpdateCategorySummary{
        Entity: *NewEntity(),
    }
    return m
}
// Gets the deviceId property value. The device ID.
func (m *MacOSSoftwareUpdateCategorySummary) GetDeviceId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.deviceId
    }
}
// Gets the displayName property value. The name of the report
func (m *MacOSSoftwareUpdateCategorySummary) GetDisplayName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.displayName
    }
}
// Gets the failedUpdateCount property value. Number of failed updates on the device
func (m *MacOSSoftwareUpdateCategorySummary) GetFailedUpdateCount()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.failedUpdateCount
    }
}
// Gets the lastUpdatedDateTime property value. Last date time the report for this device was updated.
func (m *MacOSSoftwareUpdateCategorySummary) GetLastUpdatedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.lastUpdatedDateTime
    }
}
// Gets the successfulUpdateCount property value. Number of successful updates on the device
func (m *MacOSSoftwareUpdateCategorySummary) GetSuccessfulUpdateCount()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.successfulUpdateCount
    }
}
// Gets the totalUpdateCount property value. Number of total updates on the device
func (m *MacOSSoftwareUpdateCategorySummary) GetTotalUpdateCount()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.totalUpdateCount
    }
}
// Gets the updateCategory property value. Software update type. Possible values are: critical, configurationDataFile, firmware, other.
func (m *MacOSSoftwareUpdateCategorySummary) GetUpdateCategory()(*MacOSSoftwareUpdateCategory) {
    if m == nil {
        return nil
    } else {
        return m.updateCategory
    }
}
// Gets the updateStateSummaries property value. Summary of the update states.
func (m *MacOSSoftwareUpdateCategorySummary) GetUpdateStateSummaries()([]MacOSSoftwareUpdateStateSummary) {
    if m == nil {
        return nil
    } else {
        return m.updateStateSummaries
    }
}
// Gets the userId property value. The user ID.
func (m *MacOSSoftwareUpdateCategorySummary) GetUserId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.userId
    }
}
// The deserialization information for the current model
func (m *MacOSSoftwareUpdateCategorySummary) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["deviceId"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetDeviceId(val)
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
    res["failedUpdateCount"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        m.SetFailedUpdateCount(val)
        return nil
    }
    res["lastUpdatedDateTime"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        m.SetLastUpdatedDateTime(val)
        return nil
    }
    res["successfulUpdateCount"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        m.SetSuccessfulUpdateCount(val)
        return nil
    }
    res["totalUpdateCount"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        m.SetTotalUpdateCount(val)
        return nil
    }
    res["updateCategory"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParseMacOSSoftwareUpdateCategory)
        if err != nil {
            return err
        }
        cast := val.(MacOSSoftwareUpdateCategory)
        m.SetUpdateCategory(&cast)
        return nil
    }
    res["updateStateSummaries"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewMacOSSoftwareUpdateStateSummary() })
        if err != nil {
            return err
        }
        res := make([]MacOSSoftwareUpdateStateSummary, len(val))
        for i, v := range val {
            res[i] = *(v.(*MacOSSoftwareUpdateStateSummary))
        }
        m.SetUpdateStateSummaries(res)
        return nil
    }
    res["userId"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetUserId(val)
        return nil
    }
    return res
}
func (m *MacOSSoftwareUpdateCategorySummary) IsNil()(bool) {
    return m == nil
}
// Serializes information the current object
// Parameters:
//  - writer : Serialization writer to use to serialize this model
func (m *MacOSSoftwareUpdateCategorySummary) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteStringValue("deviceId", m.GetDeviceId())
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
        err = writer.WriteInt32Value("failedUpdateCount", m.GetFailedUpdateCount())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteTimeValue("lastUpdatedDateTime", m.GetLastUpdatedDateTime())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("successfulUpdateCount", m.GetSuccessfulUpdateCount())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("totalUpdateCount", m.GetTotalUpdateCount())
        if err != nil {
            return err
        }
    }
    if m.GetUpdateCategory() != nil {
        cast := m.GetUpdateCategory().String()
        err = writer.WriteStringValue("updateCategory", &cast)
        if err != nil {
            return err
        }
    }
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetUpdateStateSummaries()))
        for i, v := range m.GetUpdateStateSummaries() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("updateStateSummaries", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("userId", m.GetUserId())
        if err != nil {
            return err
        }
    }
    return nil
}
// Sets the deviceId property value. The device ID.
// Parameters:
//  - value : Value to set for the deviceId property.
func (m *MacOSSoftwareUpdateCategorySummary) SetDeviceId(value *string)() {
    m.deviceId = value
}
// Sets the displayName property value. The name of the report
// Parameters:
//  - value : Value to set for the displayName property.
func (m *MacOSSoftwareUpdateCategorySummary) SetDisplayName(value *string)() {
    m.displayName = value
}
// Sets the failedUpdateCount property value. Number of failed updates on the device
// Parameters:
//  - value : Value to set for the failedUpdateCount property.
func (m *MacOSSoftwareUpdateCategorySummary) SetFailedUpdateCount(value *int32)() {
    m.failedUpdateCount = value
}
// Sets the lastUpdatedDateTime property value. Last date time the report for this device was updated.
// Parameters:
//  - value : Value to set for the lastUpdatedDateTime property.
func (m *MacOSSoftwareUpdateCategorySummary) SetLastUpdatedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.lastUpdatedDateTime = value
}
// Sets the successfulUpdateCount property value. Number of successful updates on the device
// Parameters:
//  - value : Value to set for the successfulUpdateCount property.
func (m *MacOSSoftwareUpdateCategorySummary) SetSuccessfulUpdateCount(value *int32)() {
    m.successfulUpdateCount = value
}
// Sets the totalUpdateCount property value. Number of total updates on the device
// Parameters:
//  - value : Value to set for the totalUpdateCount property.
func (m *MacOSSoftwareUpdateCategorySummary) SetTotalUpdateCount(value *int32)() {
    m.totalUpdateCount = value
}
// Sets the updateCategory property value. Software update type. Possible values are: critical, configurationDataFile, firmware, other.
// Parameters:
//  - value : Value to set for the updateCategory property.
func (m *MacOSSoftwareUpdateCategorySummary) SetUpdateCategory(value *MacOSSoftwareUpdateCategory)() {
    m.updateCategory = value
}
// Sets the updateStateSummaries property value. Summary of the update states.
// Parameters:
//  - value : Value to set for the updateStateSummaries property.
func (m *MacOSSoftwareUpdateCategorySummary) SetUpdateStateSummaries(value []MacOSSoftwareUpdateStateSummary)() {
    m.updateStateSummaries = value
}
// Sets the userId property value. The user ID.
// Parameters:
//  - value : Value to set for the userId property.
func (m *MacOSSoftwareUpdateCategorySummary) SetUserId(value *string)() {
    m.userId = value
}
