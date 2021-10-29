package getyammerdeviceusagedistributionusercountswithperiod

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph"
)

// 
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
// Instantiates a new getYammerDeviceUsageDistributionUserCountsWithPeriod and sets the default values.
func NewGetYammerDeviceUsageDistributionUserCountsWithPeriod()(*GetYammerDeviceUsageDistributionUserCountsWithPeriod) {
    m := &GetYammerDeviceUsageDistributionUserCountsWithPeriod{
        Entity: *i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.NewEntity(),
    }
    return m
}
// Gets the androidPhone property value. 
func (m *GetYammerDeviceUsageDistributionUserCountsWithPeriod) GetAndroidPhone()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.androidPhone
    }
}
// Gets the iPad property value. 
func (m *GetYammerDeviceUsageDistributionUserCountsWithPeriod) GetIPad()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.iPad
    }
}
// Gets the iPhone property value. 
func (m *GetYammerDeviceUsageDistributionUserCountsWithPeriod) GetIPhone()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.iPhone
    }
}
// Gets the other property value. 
func (m *GetYammerDeviceUsageDistributionUserCountsWithPeriod) GetOther()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.other
    }
}
// Gets the reportPeriod property value. 
func (m *GetYammerDeviceUsageDistributionUserCountsWithPeriod) GetReportPeriod()(*string) {
    if m == nil {
        return nil
    } else {
        return m.reportPeriod
    }
}
// Gets the reportRefreshDate property value. 
func (m *GetYammerDeviceUsageDistributionUserCountsWithPeriod) GetReportRefreshDate()(*string) {
    if m == nil {
        return nil
    } else {
        return m.reportRefreshDate
    }
}
// Gets the web property value. 
func (m *GetYammerDeviceUsageDistributionUserCountsWithPeriod) GetWeb()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.web
    }
}
// Gets the windowsPhone property value. 
func (m *GetYammerDeviceUsageDistributionUserCountsWithPeriod) GetWindowsPhone()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.windowsPhone
    }
}
// The deserialization information for the current model
func (m *GetYammerDeviceUsageDistributionUserCountsWithPeriod) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["androidPhone"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        m.SetAndroidPhone(val)
        return nil
    }
    res["iPad"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        m.SetIPad(val)
        return nil
    }
    res["iPhone"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        m.SetIPhone(val)
        return nil
    }
    res["other"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        m.SetOther(val)
        return nil
    }
    res["reportPeriod"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetReportPeriod(val)
        return nil
    }
    res["reportRefreshDate"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetReportRefreshDate(val)
        return nil
    }
    res["web"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        m.SetWeb(val)
        return nil
    }
    res["windowsPhone"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        m.SetWindowsPhone(val)
        return nil
    }
    return res
}
func (m *GetYammerDeviceUsageDistributionUserCountsWithPeriod) IsNil()(bool) {
    return m == nil
}
// Serializes information the current object
// Parameters:
//  - writer : Serialization writer to use to serialize this model
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
// Sets the androidPhone property value. 
// Parameters:
//  - value : Value to set for the androidPhone property.
func (m *GetYammerDeviceUsageDistributionUserCountsWithPeriod) SetAndroidPhone(value *int32)() {
    m.androidPhone = value
}
// Sets the iPad property value. 
// Parameters:
//  - value : Value to set for the iPad property.
func (m *GetYammerDeviceUsageDistributionUserCountsWithPeriod) SetIPad(value *int32)() {
    m.iPad = value
}
// Sets the iPhone property value. 
// Parameters:
//  - value : Value to set for the iPhone property.
func (m *GetYammerDeviceUsageDistributionUserCountsWithPeriod) SetIPhone(value *int32)() {
    m.iPhone = value
}
// Sets the other property value. 
// Parameters:
//  - value : Value to set for the other property.
func (m *GetYammerDeviceUsageDistributionUserCountsWithPeriod) SetOther(value *int32)() {
    m.other = value
}
// Sets the reportPeriod property value. 
// Parameters:
//  - value : Value to set for the reportPeriod property.
func (m *GetYammerDeviceUsageDistributionUserCountsWithPeriod) SetReportPeriod(value *string)() {
    m.reportPeriod = value
}
// Sets the reportRefreshDate property value. 
// Parameters:
//  - value : Value to set for the reportRefreshDate property.
func (m *GetYammerDeviceUsageDistributionUserCountsWithPeriod) SetReportRefreshDate(value *string)() {
    m.reportRefreshDate = value
}
// Sets the web property value. 
// Parameters:
//  - value : Value to set for the web property.
func (m *GetYammerDeviceUsageDistributionUserCountsWithPeriod) SetWeb(value *int32)() {
    m.web = value
}
// Sets the windowsPhone property value. 
// Parameters:
//  - value : Value to set for the windowsPhone property.
func (m *GetYammerDeviceUsageDistributionUserCountsWithPeriod) SetWindowsPhone(value *int32)() {
    m.windowsPhone = value
}
