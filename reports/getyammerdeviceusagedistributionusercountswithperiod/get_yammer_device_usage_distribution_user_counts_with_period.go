package getyammerdeviceusagedistributionusercountswithperiod

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph"
)

// GetYammerDeviceUsageDistributionUserCountsWithPeriod 
type GetYammerDeviceUsageDistributionUserCountsWithPeriod struct {
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.Entity
    // 
    androidPhone *int32;
    // 
    iPad *int32;
    // 
    iPhone *int32;
    // 
    other *int32;
    // 
    reportPeriod *string;
    // 
    reportRefreshDate *string;
    // 
    web *int32;
    // 
    windowsPhone *int32;
}
// NewGetYammerDeviceUsageDistributionUserCountsWithPeriod instantiates a new getYammerDeviceUsageDistributionUserCountsWithPeriod and sets the default values.
func NewGetYammerDeviceUsageDistributionUserCountsWithPeriod()(*GetYammerDeviceUsageDistributionUserCountsWithPeriod) {
    m := &GetYammerDeviceUsageDistributionUserCountsWithPeriod{
        Entity: *i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.NewEntity(),
    }
    return m
}
// GetAndroidPhone gets the androidPhone property value. 
func (m *GetYammerDeviceUsageDistributionUserCountsWithPeriod) GetAndroidPhone()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.androidPhone
    }
}
// GetIPad gets the iPad property value. 
func (m *GetYammerDeviceUsageDistributionUserCountsWithPeriod) GetIPad()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.iPad
    }
}
// GetIPhone gets the iPhone property value. 
func (m *GetYammerDeviceUsageDistributionUserCountsWithPeriod) GetIPhone()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.iPhone
    }
}
// GetOther gets the other property value. 
func (m *GetYammerDeviceUsageDistributionUserCountsWithPeriod) GetOther()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.other
    }
}
// GetReportPeriod gets the reportPeriod property value. 
func (m *GetYammerDeviceUsageDistributionUserCountsWithPeriod) GetReportPeriod()(*string) {
    if m == nil {
        return nil
    } else {
        return m.reportPeriod
    }
}
// GetReportRefreshDate gets the reportRefreshDate property value. 
func (m *GetYammerDeviceUsageDistributionUserCountsWithPeriod) GetReportRefreshDate()(*string) {
    if m == nil {
        return nil
    } else {
        return m.reportRefreshDate
    }
}
// GetWeb gets the web property value. 
func (m *GetYammerDeviceUsageDistributionUserCountsWithPeriod) GetWeb()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.web
    }
}
// GetWindowsPhone gets the windowsPhone property value. 
func (m *GetYammerDeviceUsageDistributionUserCountsWithPeriod) GetWindowsPhone()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.windowsPhone
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *GetYammerDeviceUsageDistributionUserCountsWithPeriod) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
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
    res["other"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOther(val)
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
    res["web"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetWeb(val)
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
func (m *GetYammerDeviceUsageDistributionUserCountsWithPeriod) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
func (m *GetYammerDeviceUsageDistributionUserCountsWithPeriod) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
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
        err = writer.WriteInt32Value("other", m.GetOther())
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
        err = writer.WriteInt32Value("web", m.GetWeb())
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
func (m *GetYammerDeviceUsageDistributionUserCountsWithPeriod) SetAndroidPhone(value *int32)() {
    if m != nil {
        m.androidPhone = value
    }
}
// SetIPad sets the iPad property value. 
func (m *GetYammerDeviceUsageDistributionUserCountsWithPeriod) SetIPad(value *int32)() {
    if m != nil {
        m.iPad = value
    }
}
// SetIPhone sets the iPhone property value. 
func (m *GetYammerDeviceUsageDistributionUserCountsWithPeriod) SetIPhone(value *int32)() {
    if m != nil {
        m.iPhone = value
    }
}
// SetOther sets the other property value. 
func (m *GetYammerDeviceUsageDistributionUserCountsWithPeriod) SetOther(value *int32)() {
    if m != nil {
        m.other = value
    }
}
// SetReportPeriod sets the reportPeriod property value. 
func (m *GetYammerDeviceUsageDistributionUserCountsWithPeriod) SetReportPeriod(value *string)() {
    if m != nil {
        m.reportPeriod = value
    }
}
// SetReportRefreshDate sets the reportRefreshDate property value. 
func (m *GetYammerDeviceUsageDistributionUserCountsWithPeriod) SetReportRefreshDate(value *string)() {
    if m != nil {
        m.reportRefreshDate = value
    }
}
// SetWeb sets the web property value. 
func (m *GetYammerDeviceUsageDistributionUserCountsWithPeriod) SetWeb(value *int32)() {
    if m != nil {
        m.web = value
    }
}
// SetWindowsPhone sets the windowsPhone property value. 
func (m *GetYammerDeviceUsageDistributionUserCountsWithPeriod) SetWindowsPhone(value *int32)() {
    if m != nil {
        m.windowsPhone = value
    }
}
