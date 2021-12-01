package getskypeforbusinessdeviceusageuserdetailwithdate

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph"
)

// GetSkypeForBusinessDeviceUsageUserDetailWithDate 
type GetSkypeForBusinessDeviceUsageUserDetailWithDate struct {
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
// NewGetSkypeForBusinessDeviceUsageUserDetailWithDate instantiates a new getSkypeForBusinessDeviceUsageUserDetailWithDate and sets the default values.
func NewGetSkypeForBusinessDeviceUsageUserDetailWithDate()(*GetSkypeForBusinessDeviceUsageUserDetailWithDate) {
    m := &GetSkypeForBusinessDeviceUsageUserDetailWithDate{
        Entity: *i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.NewEntity(),
    }
    return m
}
// GetLastActivityDate gets the lastActivityDate property value. 
func (m *GetSkypeForBusinessDeviceUsageUserDetailWithDate) GetLastActivityDate()(*string) {
    if m == nil {
        return nil
    } else {
        return m.lastActivityDate
    }
}
// GetReportPeriod gets the reportPeriod property value. 
func (m *GetSkypeForBusinessDeviceUsageUserDetailWithDate) GetReportPeriod()(*string) {
    if m == nil {
        return nil
    } else {
        return m.reportPeriod
    }
}
// GetReportRefreshDate gets the reportRefreshDate property value. 
func (m *GetSkypeForBusinessDeviceUsageUserDetailWithDate) GetReportRefreshDate()(*string) {
    if m == nil {
        return nil
    } else {
        return m.reportRefreshDate
    }
}
// GetUsedAndroidPhone gets the usedAndroidPhone property value. 
func (m *GetSkypeForBusinessDeviceUsageUserDetailWithDate) GetUsedAndroidPhone()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.usedAndroidPhone
    }
}
// GetUsediPad gets the usediPad property value. 
func (m *GetSkypeForBusinessDeviceUsageUserDetailWithDate) GetUsediPad()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.usediPad
    }
}
// GetUsediPhone gets the usediPhone property value. 
func (m *GetSkypeForBusinessDeviceUsageUserDetailWithDate) GetUsediPhone()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.usediPhone
    }
}
// GetUsedWindows gets the usedWindows property value. 
func (m *GetSkypeForBusinessDeviceUsageUserDetailWithDate) GetUsedWindows()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.usedWindows
    }
}
// GetUsedWindowsPhone gets the usedWindowsPhone property value. 
func (m *GetSkypeForBusinessDeviceUsageUserDetailWithDate) GetUsedWindowsPhone()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.usedWindowsPhone
    }
}
// GetUserPrincipalName gets the userPrincipalName property value. 
func (m *GetSkypeForBusinessDeviceUsageUserDetailWithDate) GetUserPrincipalName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.userPrincipalName
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *GetSkypeForBusinessDeviceUsageUserDetailWithDate) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
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
func (m *GetSkypeForBusinessDeviceUsageUserDetailWithDate) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
func (m *GetSkypeForBusinessDeviceUsageUserDetailWithDate) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
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
func (m *GetSkypeForBusinessDeviceUsageUserDetailWithDate) SetLastActivityDate(value *string)() {
    if m != nil {
        m.lastActivityDate = value
    }
}
// SetReportPeriod sets the reportPeriod property value. 
func (m *GetSkypeForBusinessDeviceUsageUserDetailWithDate) SetReportPeriod(value *string)() {
    if m != nil {
        m.reportPeriod = value
    }
}
// SetReportRefreshDate sets the reportRefreshDate property value. 
func (m *GetSkypeForBusinessDeviceUsageUserDetailWithDate) SetReportRefreshDate(value *string)() {
    if m != nil {
        m.reportRefreshDate = value
    }
}
// SetUsedAndroidPhone sets the usedAndroidPhone property value. 
func (m *GetSkypeForBusinessDeviceUsageUserDetailWithDate) SetUsedAndroidPhone(value *bool)() {
    if m != nil {
        m.usedAndroidPhone = value
    }
}
// SetUsediPad sets the usediPad property value. 
func (m *GetSkypeForBusinessDeviceUsageUserDetailWithDate) SetUsediPad(value *bool)() {
    if m != nil {
        m.usediPad = value
    }
}
// SetUsediPhone sets the usediPhone property value. 
func (m *GetSkypeForBusinessDeviceUsageUserDetailWithDate) SetUsediPhone(value *bool)() {
    if m != nil {
        m.usediPhone = value
    }
}
// SetUsedWindows sets the usedWindows property value. 
func (m *GetSkypeForBusinessDeviceUsageUserDetailWithDate) SetUsedWindows(value *bool)() {
    if m != nil {
        m.usedWindows = value
    }
}
// SetUsedWindowsPhone sets the usedWindowsPhone property value. 
func (m *GetSkypeForBusinessDeviceUsageUserDetailWithDate) SetUsedWindowsPhone(value *bool)() {
    if m != nil {
        m.usedWindowsPhone = value
    }
}
// SetUserPrincipalName sets the userPrincipalName property value. 
func (m *GetSkypeForBusinessDeviceUsageUserDetailWithDate) SetUserPrincipalName(value *string)() {
    if m != nil {
        m.userPrincipalName = value
    }
}
