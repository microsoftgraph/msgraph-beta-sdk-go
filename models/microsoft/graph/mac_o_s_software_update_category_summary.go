package graph

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// MacOSSoftwareUpdateCategorySummary 
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
// NewMacOSSoftwareUpdateCategorySummary instantiates a new macOSSoftwareUpdateCategorySummary and sets the default values.
func NewMacOSSoftwareUpdateCategorySummary()(*MacOSSoftwareUpdateCategorySummary) {
    m := &MacOSSoftwareUpdateCategorySummary{
        Entity: *NewEntity(),
    }
    return m
}
// GetDeviceId gets the deviceId property value. The device ID.
func (m *MacOSSoftwareUpdateCategorySummary) GetDeviceId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.deviceId
    }
}
// GetDisplayName gets the displayName property value. The name of the report
func (m *MacOSSoftwareUpdateCategorySummary) GetDisplayName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.displayName
    }
}
// GetFailedUpdateCount gets the failedUpdateCount property value. Number of failed updates on the device
func (m *MacOSSoftwareUpdateCategorySummary) GetFailedUpdateCount()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.failedUpdateCount
    }
}
// GetLastUpdatedDateTime gets the lastUpdatedDateTime property value. Last date time the report for this device was updated.
func (m *MacOSSoftwareUpdateCategorySummary) GetLastUpdatedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.lastUpdatedDateTime
    }
}
// GetSuccessfulUpdateCount gets the successfulUpdateCount property value. Number of successful updates on the device
func (m *MacOSSoftwareUpdateCategorySummary) GetSuccessfulUpdateCount()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.successfulUpdateCount
    }
}
// GetTotalUpdateCount gets the totalUpdateCount property value. Number of total updates on the device
func (m *MacOSSoftwareUpdateCategorySummary) GetTotalUpdateCount()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.totalUpdateCount
    }
}
// GetUpdateCategory gets the updateCategory property value. Software update type. Possible values are: critical, configurationDataFile, firmware, other.
func (m *MacOSSoftwareUpdateCategorySummary) GetUpdateCategory()(*MacOSSoftwareUpdateCategory) {
    if m == nil {
        return nil
    } else {
        return m.updateCategory
    }
}
// GetUpdateStateSummaries gets the updateStateSummaries property value. Summary of the update states.
func (m *MacOSSoftwareUpdateCategorySummary) GetUpdateStateSummaries()([]MacOSSoftwareUpdateStateSummary) {
    if m == nil {
        return nil
    } else {
        return m.updateStateSummaries
    }
}
// GetUserId gets the userId property value. The user ID.
func (m *MacOSSoftwareUpdateCategorySummary) GetUserId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.userId
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *MacOSSoftwareUpdateCategorySummary) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["deviceId"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDeviceId(val)
        }
        return nil
    }
    res["displayName"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDisplayName(val)
        }
        return nil
    }
    res["failedUpdateCount"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetFailedUpdateCount(val)
        }
        return nil
    }
    res["lastUpdatedDateTime"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLastUpdatedDateTime(val)
        }
        return nil
    }
    res["successfulUpdateCount"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSuccessfulUpdateCount(val)
        }
        return nil
    }
    res["totalUpdateCount"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTotalUpdateCount(val)
        }
        return nil
    }
    res["updateCategory"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParseMacOSSoftwareUpdateCategory)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUpdateCategory(val.(*MacOSSoftwareUpdateCategory))
        }
        return nil
    }
    res["updateStateSummaries"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewMacOSSoftwareUpdateStateSummary() })
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]MacOSSoftwareUpdateStateSummary, len(val))
            for i, v := range val {
                res[i] = *(v.(*MacOSSoftwareUpdateStateSummary))
            }
            m.SetUpdateStateSummaries(res)
        }
        return nil
    }
    res["userId"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUserId(val)
        }
        return nil
    }
    return res
}
func (m *MacOSSoftwareUpdateCategorySummary) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
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
        cast := (*m.GetUpdateCategory()).String()
        err = writer.WriteStringValue("updateCategory", &cast)
        if err != nil {
            return err
        }
    }
    if m.GetUpdateStateSummaries() != nil {
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
// SetDeviceId sets the deviceId property value. The device ID.
func (m *MacOSSoftwareUpdateCategorySummary) SetDeviceId(value *string)() {
    if m != nil {
        m.deviceId = value
    }
}
// SetDisplayName sets the displayName property value. The name of the report
func (m *MacOSSoftwareUpdateCategorySummary) SetDisplayName(value *string)() {
    if m != nil {
        m.displayName = value
    }
}
// SetFailedUpdateCount sets the failedUpdateCount property value. Number of failed updates on the device
func (m *MacOSSoftwareUpdateCategorySummary) SetFailedUpdateCount(value *int32)() {
    if m != nil {
        m.failedUpdateCount = value
    }
}
// SetLastUpdatedDateTime sets the lastUpdatedDateTime property value. Last date time the report for this device was updated.
func (m *MacOSSoftwareUpdateCategorySummary) SetLastUpdatedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    if m != nil {
        m.lastUpdatedDateTime = value
    }
}
// SetSuccessfulUpdateCount sets the successfulUpdateCount property value. Number of successful updates on the device
func (m *MacOSSoftwareUpdateCategorySummary) SetSuccessfulUpdateCount(value *int32)() {
    if m != nil {
        m.successfulUpdateCount = value
    }
}
// SetTotalUpdateCount sets the totalUpdateCount property value. Number of total updates on the device
func (m *MacOSSoftwareUpdateCategorySummary) SetTotalUpdateCount(value *int32)() {
    if m != nil {
        m.totalUpdateCount = value
    }
}
// SetUpdateCategory sets the updateCategory property value. Software update type. Possible values are: critical, configurationDataFile, firmware, other.
func (m *MacOSSoftwareUpdateCategorySummary) SetUpdateCategory(value *MacOSSoftwareUpdateCategory)() {
    if m != nil {
        m.updateCategory = value
    }
}
// SetUpdateStateSummaries sets the updateStateSummaries property value. Summary of the update states.
func (m *MacOSSoftwareUpdateCategorySummary) SetUpdateStateSummaries(value []MacOSSoftwareUpdateStateSummary)() {
    if m != nil {
        m.updateStateSummaries = value
    }
}
// SetUserId sets the userId property value. The user ID.
func (m *MacOSSoftwareUpdateCategorySummary) SetUserId(value *string)() {
    if m != nil {
        m.userId = value
    }
}
