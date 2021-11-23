package getskypeforbusinessdeviceusageuserdetailwithperiod

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph"
)

// getSkypeForBusinessDeviceUsageUserDetailWithPeriod 
type GetSkypeForBusinessDeviceUsageUserDetailWithPeriod struct {
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.Entity
    // 
    lastActivityDate *string;
    // 
    reportPeriod *string;
    // 
    reportRefreshDate *string;
    // 
    usedAndroidPhone *bool;
    // 
    usediPad *bool;
    // 
    usediPhone *bool;
    // 
    usedWindows *bool;
    // 
    usedWindowsPhone *bool;
    // 
    userPrincipalName *string;
}
// NewGetSkypeForBusinessDeviceUsageUserDetailWithPeriod instantiates a new getSkypeForBusinessDeviceUsageUserDetailWithPeriod and sets the default values.
func NewGetSkypeForBusinessDeviceUsageUserDetailWithPeriod()(*GetSkypeForBusinessDeviceUsageUserDetailWithPeriod) {
    m := &GetSkypeForBusinessDeviceUsageUserDetailWithPeriod{
        Entity: *i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.NewEntity(),
    }
    return m
}
// GetLastActivityDate gets the lastActivityDate property value. 
func (m *GetSkypeForBusinessDeviceUsageUserDetailWithPeriod) GetLastActivityDate()(*string) {
    if m == nil {
        return nil
    } else {
        return m.lastActivityDate
    }
}
// GetReportPeriod gets the reportPeriod property value. 
func (m *GetSkypeForBusinessDeviceUsageUserDetailWithPeriod) GetReportPeriod()(*string) {
    if m == nil {
        return nil
    } else {
        return m.reportPeriod
    }
}
// GetReportRefreshDate gets the reportRefreshDate property value. 
func (m *GetSkypeForBusinessDeviceUsageUserDetailWithPeriod) GetReportRefreshDate()(*string) {
    if m == nil {
        return nil
    } else {
        return m.reportRefreshDate
    }
}
// GetUsedAndroidPhone gets the usedAndroidPhone property value. 
func (m *GetSkypeForBusinessDeviceUsageUserDetailWithPeriod) GetUsedAndroidPhone()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.usedAndroidPhone
    }
}
// GetUsediPad gets the usediPad property value. 
func (m *GetSkypeForBusinessDeviceUsageUserDetailWithPeriod) GetUsediPad()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.usediPad
    }
}
// GetUsediPhone gets the usediPhone property value. 
func (m *GetSkypeForBusinessDeviceUsageUserDetailWithPeriod) GetUsediPhone()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.usediPhone
    }
}
// GetUsedWindows gets the usedWindows property value. 
func (m *GetSkypeForBusinessDeviceUsageUserDetailWithPeriod) GetUsedWindows()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.usedWindows
    }
}
// GetUsedWindowsPhone gets the usedWindowsPhone property value. 
func (m *GetSkypeForBusinessDeviceUsageUserDetailWithPeriod) GetUsedWindowsPhone()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.usedWindowsPhone
    }
}
// GetUserPrincipalName gets the userPrincipalName property value. 
func (m *GetSkypeForBusinessDeviceUsageUserDetailWithPeriod) GetUserPrincipalName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.userPrincipalName
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *GetSkypeForBusinessDeviceUsageUserDetailWithPeriod) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["lastActivityDate"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLastActivityDate(val)
        }
        return nil
    }
    res["reportPeriod"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetReportPeriod(val)
        }
        return nil
    }
    res["reportRefreshDate"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetReportRefreshDate(val)
        }
        return nil
    }
    res["usedAndroidPhone"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUsedAndroidPhone(val)
        }
        return nil
    }
    res["usediPad"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUsediPad(val)
        }
        return nil
    }
    res["usediPhone"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUsediPhone(val)
        }
        return nil
    }
    res["usedWindows"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUsedWindows(val)
        }
        return nil
    }
    res["usedWindowsPhone"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUsedWindowsPhone(val)
        }
        return nil
    }
    res["userPrincipalName"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUserPrincipalName(val)
        }
        return nil
    }
    return res
}
func (m *GetSkypeForBusinessDeviceUsageUserDetailWithPeriod) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
func (m *GetSkypeForBusinessDeviceUsageUserDetailWithPeriod) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteStringValue("lastActivityDate", m.GetLastActivityDate())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("reportPeriod", m.GetReportPeriod())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("reportRefreshDate", m.GetReportRefreshDate())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("usedAndroidPhone", m.GetUsedAndroidPhone())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("usediPad", m.GetUsediPad())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("usediPhone", m.GetUsediPhone())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("usedWindows", m.GetUsedWindows())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("usedWindowsPhone", m.GetUsedWindowsPhone())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("userPrincipalName", m.GetUserPrincipalName())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetLastActivityDate sets the lastActivityDate property value. 
func (m *GetSkypeForBusinessDeviceUsageUserDetailWithPeriod) SetLastActivityDate(value *string)() {
    m.lastActivityDate = value
}
// SetReportPeriod sets the reportPeriod property value. 
func (m *GetSkypeForBusinessDeviceUsageUserDetailWithPeriod) SetReportPeriod(value *string)() {
    m.reportPeriod = value
}
// SetReportRefreshDate sets the reportRefreshDate property value. 
func (m *GetSkypeForBusinessDeviceUsageUserDetailWithPeriod) SetReportRefreshDate(value *string)() {
    m.reportRefreshDate = value
}
// SetUsedAndroidPhone sets the usedAndroidPhone property value. 
func (m *GetSkypeForBusinessDeviceUsageUserDetailWithPeriod) SetUsedAndroidPhone(value *bool)() {
    m.usedAndroidPhone = value
}
// SetUsediPad sets the usediPad property value. 
func (m *GetSkypeForBusinessDeviceUsageUserDetailWithPeriod) SetUsediPad(value *bool)() {
    m.usediPad = value
}
// SetUsediPhone sets the usediPhone property value. 
func (m *GetSkypeForBusinessDeviceUsageUserDetailWithPeriod) SetUsediPhone(value *bool)() {
    m.usediPhone = value
}
// SetUsedWindows sets the usedWindows property value. 
func (m *GetSkypeForBusinessDeviceUsageUserDetailWithPeriod) SetUsedWindows(value *bool)() {
    m.usedWindows = value
}
// SetUsedWindowsPhone sets the usedWindowsPhone property value. 
func (m *GetSkypeForBusinessDeviceUsageUserDetailWithPeriod) SetUsedWindowsPhone(value *bool)() {
    m.usedWindowsPhone = value
}
// SetUserPrincipalName sets the userPrincipalName property value. 
func (m *GetSkypeForBusinessDeviceUsageUserDetailWithPeriod) SetUserPrincipalName(value *string)() {
    m.userPrincipalName = value
}
