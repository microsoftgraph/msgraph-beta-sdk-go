package getskypeforbusinessdeviceusageuserdetailwithperiod

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph"
)

// 
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
// Instantiates a new getSkypeForBusinessDeviceUsageUserDetailWithPeriod and sets the default values.
func NewGetSkypeForBusinessDeviceUsageUserDetailWithPeriod()(*GetSkypeForBusinessDeviceUsageUserDetailWithPeriod) {
    m := &GetSkypeForBusinessDeviceUsageUserDetailWithPeriod{
        Entity: *i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.NewEntity(),
    }
    return m
}
// Gets the lastActivityDate property value. 
func (m *GetSkypeForBusinessDeviceUsageUserDetailWithPeriod) GetLastActivityDate()(*string) {
    if m == nil {
        return nil
    } else {
        return m.lastActivityDate
    }
}
// Gets the reportPeriod property value. 
func (m *GetSkypeForBusinessDeviceUsageUserDetailWithPeriod) GetReportPeriod()(*string) {
    if m == nil {
        return nil
    } else {
        return m.reportPeriod
    }
}
// Gets the reportRefreshDate property value. 
func (m *GetSkypeForBusinessDeviceUsageUserDetailWithPeriod) GetReportRefreshDate()(*string) {
    if m == nil {
        return nil
    } else {
        return m.reportRefreshDate
    }
}
// Gets the usedAndroidPhone property value. 
func (m *GetSkypeForBusinessDeviceUsageUserDetailWithPeriod) GetUsedAndroidPhone()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.usedAndroidPhone
    }
}
// Gets the usediPad property value. 
func (m *GetSkypeForBusinessDeviceUsageUserDetailWithPeriod) GetUsediPad()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.usediPad
    }
}
// Gets the usediPhone property value. 
func (m *GetSkypeForBusinessDeviceUsageUserDetailWithPeriod) GetUsediPhone()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.usediPhone
    }
}
// Gets the usedWindows property value. 
func (m *GetSkypeForBusinessDeviceUsageUserDetailWithPeriod) GetUsedWindows()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.usedWindows
    }
}
// Gets the usedWindowsPhone property value. 
func (m *GetSkypeForBusinessDeviceUsageUserDetailWithPeriod) GetUsedWindowsPhone()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.usedWindowsPhone
    }
}
// Gets the userPrincipalName property value. 
func (m *GetSkypeForBusinessDeviceUsageUserDetailWithPeriod) GetUserPrincipalName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.userPrincipalName
    }
}
// The deserialization information for the current model
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
// Serializes information the current object
// Parameters:
//  - writer : Serialization writer to use to serialize this model
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
// Sets the lastActivityDate property value. 
// Parameters:
//  - value : Value to set for the lastActivityDate property.
func (m *GetSkypeForBusinessDeviceUsageUserDetailWithPeriod) SetLastActivityDate(value *string)() {
    m.lastActivityDate = value
}
// Sets the reportPeriod property value. 
// Parameters:
//  - value : Value to set for the reportPeriod property.
func (m *GetSkypeForBusinessDeviceUsageUserDetailWithPeriod) SetReportPeriod(value *string)() {
    m.reportPeriod = value
}
// Sets the reportRefreshDate property value. 
// Parameters:
//  - value : Value to set for the reportRefreshDate property.
func (m *GetSkypeForBusinessDeviceUsageUserDetailWithPeriod) SetReportRefreshDate(value *string)() {
    m.reportRefreshDate = value
}
// Sets the usedAndroidPhone property value. 
// Parameters:
//  - value : Value to set for the usedAndroidPhone property.
func (m *GetSkypeForBusinessDeviceUsageUserDetailWithPeriod) SetUsedAndroidPhone(value *bool)() {
    m.usedAndroidPhone = value
}
// Sets the usediPad property value. 
// Parameters:
//  - value : Value to set for the usediPad property.
func (m *GetSkypeForBusinessDeviceUsageUserDetailWithPeriod) SetUsediPad(value *bool)() {
    m.usediPad = value
}
// Sets the usediPhone property value. 
// Parameters:
//  - value : Value to set for the usediPhone property.
func (m *GetSkypeForBusinessDeviceUsageUserDetailWithPeriod) SetUsediPhone(value *bool)() {
    m.usediPhone = value
}
// Sets the usedWindows property value. 
// Parameters:
//  - value : Value to set for the usedWindows property.
func (m *GetSkypeForBusinessDeviceUsageUserDetailWithPeriod) SetUsedWindows(value *bool)() {
    m.usedWindows = value
}
// Sets the usedWindowsPhone property value. 
// Parameters:
//  - value : Value to set for the usedWindowsPhone property.
func (m *GetSkypeForBusinessDeviceUsageUserDetailWithPeriod) SetUsedWindowsPhone(value *bool)() {
    m.usedWindowsPhone = value
}
// Sets the userPrincipalName property value. 
// Parameters:
//  - value : Value to set for the userPrincipalName property.
func (m *GetSkypeForBusinessDeviceUsageUserDetailWithPeriod) SetUserPrincipalName(value *string)() {
    m.userPrincipalName = value
}
