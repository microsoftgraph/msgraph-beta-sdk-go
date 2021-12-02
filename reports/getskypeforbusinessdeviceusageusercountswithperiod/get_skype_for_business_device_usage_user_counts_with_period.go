package getskypeforbusinessdeviceusageusercountswithperiod

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph"
)

// GetSkypeForBusinessDeviceUsageUserCountsWithPeriod 
type GetSkypeForBusinessDeviceUsageUserCountsWithPeriod struct {
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.Entity
    // 
    androidPhone *int32;
    // 
    iPad *int32;
    // 
    iPhone *int32;
    // 
    reportDate *string;
    // 
    reportPeriod *string;
    // 
    reportRefreshDate *string;
    // 
    windows *int32;
    // 
    windowsPhone *int32;
}
// NewGetSkypeForBusinessDeviceUsageUserCountsWithPeriod instantiates a new getSkypeForBusinessDeviceUsageUserCountsWithPeriod and sets the default values.
func NewGetSkypeForBusinessDeviceUsageUserCountsWithPeriod()(*GetSkypeForBusinessDeviceUsageUserCountsWithPeriod) {
    m := &GetSkypeForBusinessDeviceUsageUserCountsWithPeriod{
        Entity: *i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.NewEntity(),
    }
    return m
}
// GetAndroidPhone gets the androidPhone property value. 
func (m *GetSkypeForBusinessDeviceUsageUserCountsWithPeriod) GetAndroidPhone()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.androidPhone
    }
}
// GetIPad gets the iPad property value. 
func (m *GetSkypeForBusinessDeviceUsageUserCountsWithPeriod) GetIPad()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.iPad
    }
}
// GetIPhone gets the iPhone property value. 
func (m *GetSkypeForBusinessDeviceUsageUserCountsWithPeriod) GetIPhone()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.iPhone
    }
}
// GetReportDate gets the reportDate property value. 
func (m *GetSkypeForBusinessDeviceUsageUserCountsWithPeriod) GetReportDate()(*string) {
    if m == nil {
        return nil
    } else {
        return m.reportDate
    }
}
// GetReportPeriod gets the reportPeriod property value. 
func (m *GetSkypeForBusinessDeviceUsageUserCountsWithPeriod) GetReportPeriod()(*string) {
    if m == nil {
        return nil
    } else {
        return m.reportPeriod
    }
}
// GetReportRefreshDate gets the reportRefreshDate property value. 
func (m *GetSkypeForBusinessDeviceUsageUserCountsWithPeriod) GetReportRefreshDate()(*string) {
    if m == nil {
        return nil
    } else {
        return m.reportRefreshDate
    }
}
// GetWindows gets the windows property value. 
func (m *GetSkypeForBusinessDeviceUsageUserCountsWithPeriod) GetWindows()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.windows
    }
}
// GetWindowsPhone gets the windowsPhone property value. 
func (m *GetSkypeForBusinessDeviceUsageUserCountsWithPeriod) GetWindowsPhone()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.windowsPhone
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *GetSkypeForBusinessDeviceUsageUserCountsWithPeriod) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["androidPhone"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAndroidPhone(val)
        }
        return nil
    }
    res["iPad"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIPad(val)
        }
        return nil
    }
    res["iPhone"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIPhone(val)
        }
        return nil
    }
    res["reportDate"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetReportDate(val)
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
    res["windows"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetWindows(val)
        }
        return nil
    }
    res["windowsPhone"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetWindowsPhone(val)
        }
        return nil
    }
    return res
}
func (m *GetSkypeForBusinessDeviceUsageUserCountsWithPeriod) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
func (m *GetSkypeForBusinessDeviceUsageUserCountsWithPeriod) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteInt32Value("androidPhone", m.GetAndroidPhone())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("iPad", m.GetIPad())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("iPhone", m.GetIPhone())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("reportDate", m.GetReportDate())
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
        err = writer.WriteInt32Value("windows", m.GetWindows())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("windowsPhone", m.GetWindowsPhone())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAndroidPhone sets the androidPhone property value. 
func (m *GetSkypeForBusinessDeviceUsageUserCountsWithPeriod) SetAndroidPhone(value *int32)() {
    if m != nil {
        m.androidPhone = value
    }
}
// SetIPad sets the iPad property value. 
func (m *GetSkypeForBusinessDeviceUsageUserCountsWithPeriod) SetIPad(value *int32)() {
    if m != nil {
        m.iPad = value
    }
}
// SetIPhone sets the iPhone property value. 
func (m *GetSkypeForBusinessDeviceUsageUserCountsWithPeriod) SetIPhone(value *int32)() {
    if m != nil {
        m.iPhone = value
    }
}
// SetReportDate sets the reportDate property value. 
func (m *GetSkypeForBusinessDeviceUsageUserCountsWithPeriod) SetReportDate(value *string)() {
    if m != nil {
        m.reportDate = value
    }
}
// SetReportPeriod sets the reportPeriod property value. 
func (m *GetSkypeForBusinessDeviceUsageUserCountsWithPeriod) SetReportPeriod(value *string)() {
    if m != nil {
        m.reportPeriod = value
    }
}
// SetReportRefreshDate sets the reportRefreshDate property value. 
func (m *GetSkypeForBusinessDeviceUsageUserCountsWithPeriod) SetReportRefreshDate(value *string)() {
    if m != nil {
        m.reportRefreshDate = value
    }
}
// SetWindows sets the windows property value. 
func (m *GetSkypeForBusinessDeviceUsageUserCountsWithPeriod) SetWindows(value *int32)() {
    if m != nil {
        m.windows = value
    }
}
// SetWindowsPhone sets the windowsPhone property value. 
func (m *GetSkypeForBusinessDeviceUsageUserCountsWithPeriod) SetWindowsPhone(value *int32)() {
    if m != nil {
        m.windowsPhone = value
    }
}
