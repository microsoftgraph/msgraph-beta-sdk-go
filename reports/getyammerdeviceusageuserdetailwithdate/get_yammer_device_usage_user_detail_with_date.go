package getyammerdeviceusageuserdetailwithdate

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph"
)

// 
type GetYammerDeviceUsageUserDetailWithDate struct {
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.Entity
    // 
    displayName *string;
    // 
    lastActivityDate *string;
    // 
    reportPeriod *string;
    // 
    reportRefreshDate *string;
    // 
    stateChangeDate *string;
    // 
    usedAndroidPhone *bool;
    // 
    usediPad *bool;
    // 
    usediPhone *bool;
    // 
    usedOthers *bool;
    // 
    usedWeb *bool;
    // 
    usedWindowsPhone *bool;
    // 
    userPrincipalName *string;
    // 
    userState *string;
}
// Instantiates a new getYammerDeviceUsageUserDetailWithDate and sets the default values.
func NewGetYammerDeviceUsageUserDetailWithDate()(*GetYammerDeviceUsageUserDetailWithDate) {
    m := &GetYammerDeviceUsageUserDetailWithDate{
        Entity: *i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.NewEntity(),
    }
    return m
}
// Gets the displayName property value. 
func (m *GetYammerDeviceUsageUserDetailWithDate) GetDisplayName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.displayName
    }
}
// Gets the lastActivityDate property value. 
func (m *GetYammerDeviceUsageUserDetailWithDate) GetLastActivityDate()(*string) {
    if m == nil {
        return nil
    } else {
        return m.lastActivityDate
    }
}
// Gets the reportPeriod property value. 
func (m *GetYammerDeviceUsageUserDetailWithDate) GetReportPeriod()(*string) {
    if m == nil {
        return nil
    } else {
        return m.reportPeriod
    }
}
// Gets the reportRefreshDate property value. 
func (m *GetYammerDeviceUsageUserDetailWithDate) GetReportRefreshDate()(*string) {
    if m == nil {
        return nil
    } else {
        return m.reportRefreshDate
    }
}
// Gets the stateChangeDate property value. 
func (m *GetYammerDeviceUsageUserDetailWithDate) GetStateChangeDate()(*string) {
    if m == nil {
        return nil
    } else {
        return m.stateChangeDate
    }
}
// Gets the usedAndroidPhone property value. 
func (m *GetYammerDeviceUsageUserDetailWithDate) GetUsedAndroidPhone()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.usedAndroidPhone
    }
}
// Gets the usediPad property value. 
func (m *GetYammerDeviceUsageUserDetailWithDate) GetUsediPad()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.usediPad
    }
}
// Gets the usediPhone property value. 
func (m *GetYammerDeviceUsageUserDetailWithDate) GetUsediPhone()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.usediPhone
    }
}
// Gets the usedOthers property value. 
func (m *GetYammerDeviceUsageUserDetailWithDate) GetUsedOthers()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.usedOthers
    }
}
// Gets the usedWeb property value. 
func (m *GetYammerDeviceUsageUserDetailWithDate) GetUsedWeb()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.usedWeb
    }
}
// Gets the usedWindowsPhone property value. 
func (m *GetYammerDeviceUsageUserDetailWithDate) GetUsedWindowsPhone()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.usedWindowsPhone
    }
}
// Gets the userPrincipalName property value. 
func (m *GetYammerDeviceUsageUserDetailWithDate) GetUserPrincipalName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.userPrincipalName
    }
}
// Gets the userState property value. 
func (m *GetYammerDeviceUsageUserDetailWithDate) GetUserState()(*string) {
    if m == nil {
        return nil
    } else {
        return m.userState
    }
}
// The deserialization information for the current model
func (m *GetYammerDeviceUsageUserDetailWithDate) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
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
    res["stateChangeDate"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetStateChangeDate(val)
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
    res["usedOthers"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUsedOthers(val)
        }
        return nil
    }
    res["usedWeb"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUsedWeb(val)
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
    res["userState"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUserState(val)
        }
        return nil
    }
    return res
}
func (m *GetYammerDeviceUsageUserDetailWithDate) IsNil()(bool) {
    return m == nil
}
// Serializes information the current object
// Parameters:
//  - writer : Serialization writer to use to serialize this model
func (m *GetYammerDeviceUsageUserDetailWithDate) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
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
        err = writer.WriteStringValue("stateChangeDate", m.GetStateChangeDate())
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
        err = writer.WriteBoolValue("usedOthers", m.GetUsedOthers())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("usedWeb", m.GetUsedWeb())
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
    {
        err = writer.WriteStringValue("userState", m.GetUserState())
        if err != nil {
            return err
        }
    }
    return nil
}
// Sets the displayName property value. 
// Parameters:
//  - value : Value to set for the displayName property.
func (m *GetYammerDeviceUsageUserDetailWithDate) SetDisplayName(value *string)() {
    m.displayName = value
}
// Sets the lastActivityDate property value. 
// Parameters:
//  - value : Value to set for the lastActivityDate property.
func (m *GetYammerDeviceUsageUserDetailWithDate) SetLastActivityDate(value *string)() {
    m.lastActivityDate = value
}
// Sets the reportPeriod property value. 
// Parameters:
//  - value : Value to set for the reportPeriod property.
func (m *GetYammerDeviceUsageUserDetailWithDate) SetReportPeriod(value *string)() {
    m.reportPeriod = value
}
// Sets the reportRefreshDate property value. 
// Parameters:
//  - value : Value to set for the reportRefreshDate property.
func (m *GetYammerDeviceUsageUserDetailWithDate) SetReportRefreshDate(value *string)() {
    m.reportRefreshDate = value
}
// Sets the stateChangeDate property value. 
// Parameters:
//  - value : Value to set for the stateChangeDate property.
func (m *GetYammerDeviceUsageUserDetailWithDate) SetStateChangeDate(value *string)() {
    m.stateChangeDate = value
}
// Sets the usedAndroidPhone property value. 
// Parameters:
//  - value : Value to set for the usedAndroidPhone property.
func (m *GetYammerDeviceUsageUserDetailWithDate) SetUsedAndroidPhone(value *bool)() {
    m.usedAndroidPhone = value
}
// Sets the usediPad property value. 
// Parameters:
//  - value : Value to set for the usediPad property.
func (m *GetYammerDeviceUsageUserDetailWithDate) SetUsediPad(value *bool)() {
    m.usediPad = value
}
// Sets the usediPhone property value. 
// Parameters:
//  - value : Value to set for the usediPhone property.
func (m *GetYammerDeviceUsageUserDetailWithDate) SetUsediPhone(value *bool)() {
    m.usediPhone = value
}
// Sets the usedOthers property value. 
// Parameters:
//  - value : Value to set for the usedOthers property.
func (m *GetYammerDeviceUsageUserDetailWithDate) SetUsedOthers(value *bool)() {
    m.usedOthers = value
}
// Sets the usedWeb property value. 
// Parameters:
//  - value : Value to set for the usedWeb property.
func (m *GetYammerDeviceUsageUserDetailWithDate) SetUsedWeb(value *bool)() {
    m.usedWeb = value
}
// Sets the usedWindowsPhone property value. 
// Parameters:
//  - value : Value to set for the usedWindowsPhone property.
func (m *GetYammerDeviceUsageUserDetailWithDate) SetUsedWindowsPhone(value *bool)() {
    m.usedWindowsPhone = value
}
// Sets the userPrincipalName property value. 
// Parameters:
//  - value : Value to set for the userPrincipalName property.
func (m *GetYammerDeviceUsageUserDetailWithDate) SetUserPrincipalName(value *string)() {
    m.userPrincipalName = value
}
// Sets the userState property value. 
// Parameters:
//  - value : Value to set for the userState property.
func (m *GetYammerDeviceUsageUserDetailWithDate) SetUserState(value *string)() {
    m.userState = value
}
