package getyammerdeviceusageuserdetailwithperiod

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph"
)

// GetYammerDeviceUsageUserDetailWithPeriod 
type GetYammerDeviceUsageUserDetailWithPeriod struct {
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
// NewGetYammerDeviceUsageUserDetailWithPeriod instantiates a new getYammerDeviceUsageUserDetailWithPeriod and sets the default values.
func NewGetYammerDeviceUsageUserDetailWithPeriod()(*GetYammerDeviceUsageUserDetailWithPeriod) {
    m := &GetYammerDeviceUsageUserDetailWithPeriod{
        Entity: *i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.NewEntity(),
    }
    return m
}
// GetDisplayName gets the displayName property value. 
func (m *GetYammerDeviceUsageUserDetailWithPeriod) GetDisplayName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.displayName
    }
}
// GetLastActivityDate gets the lastActivityDate property value. 
func (m *GetYammerDeviceUsageUserDetailWithPeriod) GetLastActivityDate()(*string) {
    if m == nil {
        return nil
    } else {
        return m.lastActivityDate
    }
}
// GetReportPeriod gets the reportPeriod property value. 
func (m *GetYammerDeviceUsageUserDetailWithPeriod) GetReportPeriod()(*string) {
    if m == nil {
        return nil
    } else {
        return m.reportPeriod
    }
}
// GetReportRefreshDate gets the reportRefreshDate property value. 
func (m *GetYammerDeviceUsageUserDetailWithPeriod) GetReportRefreshDate()(*string) {
    if m == nil {
        return nil
    } else {
        return m.reportRefreshDate
    }
}
// GetStateChangeDate gets the stateChangeDate property value. 
func (m *GetYammerDeviceUsageUserDetailWithPeriod) GetStateChangeDate()(*string) {
    if m == nil {
        return nil
    } else {
        return m.stateChangeDate
    }
}
// GetUsedAndroidPhone gets the usedAndroidPhone property value. 
func (m *GetYammerDeviceUsageUserDetailWithPeriod) GetUsedAndroidPhone()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.usedAndroidPhone
    }
}
// GetUsediPad gets the usediPad property value. 
func (m *GetYammerDeviceUsageUserDetailWithPeriod) GetUsediPad()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.usediPad
    }
}
// GetUsediPhone gets the usediPhone property value. 
func (m *GetYammerDeviceUsageUserDetailWithPeriod) GetUsediPhone()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.usediPhone
    }
}
// GetUsedOthers gets the usedOthers property value. 
func (m *GetYammerDeviceUsageUserDetailWithPeriod) GetUsedOthers()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.usedOthers
    }
}
// GetUsedWeb gets the usedWeb property value. 
func (m *GetYammerDeviceUsageUserDetailWithPeriod) GetUsedWeb()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.usedWeb
    }
}
// GetUsedWindowsPhone gets the usedWindowsPhone property value. 
func (m *GetYammerDeviceUsageUserDetailWithPeriod) GetUsedWindowsPhone()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.usedWindowsPhone
    }
}
// GetUserPrincipalName gets the userPrincipalName property value. 
func (m *GetYammerDeviceUsageUserDetailWithPeriod) GetUserPrincipalName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.userPrincipalName
    }
}
// GetUserState gets the userState property value. 
func (m *GetYammerDeviceUsageUserDetailWithPeriod) GetUserState()(*string) {
    if m == nil {
        return nil
    } else {
        return m.userState
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *GetYammerDeviceUsageUserDetailWithPeriod) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
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
func (m *GetYammerDeviceUsageUserDetailWithPeriod) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
func (m *GetYammerDeviceUsageUserDetailWithPeriod) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
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
// SetDisplayName sets the displayName property value. 
func (m *GetYammerDeviceUsageUserDetailWithPeriod) SetDisplayName(value *string)() {
    if m != nil {
        m.displayName = value
    }
}
// SetLastActivityDate sets the lastActivityDate property value. 
func (m *GetYammerDeviceUsageUserDetailWithPeriod) SetLastActivityDate(value *string)() {
    if m != nil {
        m.lastActivityDate = value
    }
}
// SetReportPeriod sets the reportPeriod property value. 
func (m *GetYammerDeviceUsageUserDetailWithPeriod) SetReportPeriod(value *string)() {
    if m != nil {
        m.reportPeriod = value
    }
}
// SetReportRefreshDate sets the reportRefreshDate property value. 
func (m *GetYammerDeviceUsageUserDetailWithPeriod) SetReportRefreshDate(value *string)() {
    if m != nil {
        m.reportRefreshDate = value
    }
}
// SetStateChangeDate sets the stateChangeDate property value. 
func (m *GetYammerDeviceUsageUserDetailWithPeriod) SetStateChangeDate(value *string)() {
    if m != nil {
        m.stateChangeDate = value
    }
}
// SetUsedAndroidPhone sets the usedAndroidPhone property value. 
func (m *GetYammerDeviceUsageUserDetailWithPeriod) SetUsedAndroidPhone(value *bool)() {
    if m != nil {
        m.usedAndroidPhone = value
    }
}
// SetUsediPad sets the usediPad property value. 
func (m *GetYammerDeviceUsageUserDetailWithPeriod) SetUsediPad(value *bool)() {
    if m != nil {
        m.usediPad = value
    }
}
// SetUsediPhone sets the usediPhone property value. 
func (m *GetYammerDeviceUsageUserDetailWithPeriod) SetUsediPhone(value *bool)() {
    if m != nil {
        m.usediPhone = value
    }
}
// SetUsedOthers sets the usedOthers property value. 
func (m *GetYammerDeviceUsageUserDetailWithPeriod) SetUsedOthers(value *bool)() {
    if m != nil {
        m.usedOthers = value
    }
}
// SetUsedWeb sets the usedWeb property value. 
func (m *GetYammerDeviceUsageUserDetailWithPeriod) SetUsedWeb(value *bool)() {
    if m != nil {
        m.usedWeb = value
    }
}
// SetUsedWindowsPhone sets the usedWindowsPhone property value. 
func (m *GetYammerDeviceUsageUserDetailWithPeriod) SetUsedWindowsPhone(value *bool)() {
    if m != nil {
        m.usedWindowsPhone = value
    }
}
// SetUserPrincipalName sets the userPrincipalName property value. 
func (m *GetYammerDeviceUsageUserDetailWithPeriod) SetUserPrincipalName(value *string)() {
    if m != nil {
        m.userPrincipalName = value
    }
}
// SetUserState sets the userState property value. 
func (m *GetYammerDeviceUsageUserDetailWithPeriod) SetUserState(value *string)() {
    if m != nil {
        m.userState = value
    }
}
