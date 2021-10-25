package getteamsdeviceusagedistributionusercountswithperiod

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph"
)

type GetTeamsDeviceUsageDistributionUserCountsWithPeriod struct {
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.Entity
    androidPhone *int64;
    chromeOS *int64;
    ios *int64;
    linux *int64;
    mac *int64;
    reportPeriod *string;
    reportRefreshDate *string;
    web *int64;
    windows *int64;
    windowsPhone *int64;
}
func NewGetTeamsDeviceUsageDistributionUserCountsWithPeriod()(*GetTeamsDeviceUsageDistributionUserCountsWithPeriod) {
    m := &GetTeamsDeviceUsageDistributionUserCountsWithPeriod{
        Entity: *i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.NewEntity(),
    }
    return m
}
func (m *GetTeamsDeviceUsageDistributionUserCountsWithPeriod) GetAndroidPhone()(*int64) {
    if m == nil {
        return nil
    } else {
        return m.androidPhone
    }
}
func (m *GetTeamsDeviceUsageDistributionUserCountsWithPeriod) GetChromeOS()(*int64) {
    if m == nil {
        return nil
    } else {
        return m.chromeOS
    }
}
func (m *GetTeamsDeviceUsageDistributionUserCountsWithPeriod) GetIos()(*int64) {
    if m == nil {
        return nil
    } else {
        return m.ios
    }
}
func (m *GetTeamsDeviceUsageDistributionUserCountsWithPeriod) GetLinux()(*int64) {
    if m == nil {
        return nil
    } else {
        return m.linux
    }
}
func (m *GetTeamsDeviceUsageDistributionUserCountsWithPeriod) GetMac()(*int64) {
    if m == nil {
        return nil
    } else {
        return m.mac
    }
}
func (m *GetTeamsDeviceUsageDistributionUserCountsWithPeriod) GetReportPeriod()(*string) {
    if m == nil {
        return nil
    } else {
        return m.reportPeriod
    }
}
func (m *GetTeamsDeviceUsageDistributionUserCountsWithPeriod) GetReportRefreshDate()(*string) {
    if m == nil {
        return nil
    } else {
        return m.reportRefreshDate
    }
}
func (m *GetTeamsDeviceUsageDistributionUserCountsWithPeriod) GetWeb()(*int64) {
    if m == nil {
        return nil
    } else {
        return m.web
    }
}
func (m *GetTeamsDeviceUsageDistributionUserCountsWithPeriod) GetWindows()(*int64) {
    if m == nil {
        return nil
    } else {
        return m.windows
    }
}
func (m *GetTeamsDeviceUsageDistributionUserCountsWithPeriod) GetWindowsPhone()(*int64) {
    if m == nil {
        return nil
    } else {
        return m.windowsPhone
    }
}
func (m *GetTeamsDeviceUsageDistributionUserCountsWithPeriod) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["androidPhone"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        m.SetAndroidPhone(val)
        return nil
    }
    res["chromeOS"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        m.SetChromeOS(val)
        return nil
    }
    res["ios"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        m.SetIos(val)
        return nil
    }
    res["linux"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        m.SetLinux(val)
        return nil
    }
    res["mac"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        m.SetMac(val)
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
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        m.SetWeb(val)
        return nil
    }
    res["windows"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        m.SetWindows(val)
        return nil
    }
    res["windowsPhone"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        m.SetWindowsPhone(val)
        return nil
    }
    return res
}
func (m *GetTeamsDeviceUsageDistributionUserCountsWithPeriod) IsNil()(bool) {
    return m == nil
}
func (m *GetTeamsDeviceUsageDistributionUserCountsWithPeriod) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteInt64Value("androidPhone", m.GetAndroidPhone())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt64Value("chromeOS", m.GetChromeOS())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt64Value("ios", m.GetIos())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt64Value("linux", m.GetLinux())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt64Value("mac", m.GetMac())
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
        err = writer.WriteInt64Value("web", m.GetWeb())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt64Value("windows", m.GetWindows())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt64Value("windowsPhone", m.GetWindowsPhone())
        if err != nil {
            return err
        }
    }
    return nil
}
func (m *GetTeamsDeviceUsageDistributionUserCountsWithPeriod) SetAndroidPhone(value *int64)() {
    m.androidPhone = value
}
func (m *GetTeamsDeviceUsageDistributionUserCountsWithPeriod) SetChromeOS(value *int64)() {
    m.chromeOS = value
}
func (m *GetTeamsDeviceUsageDistributionUserCountsWithPeriod) SetIos(value *int64)() {
    m.ios = value
}
func (m *GetTeamsDeviceUsageDistributionUserCountsWithPeriod) SetLinux(value *int64)() {
    m.linux = value
}
func (m *GetTeamsDeviceUsageDistributionUserCountsWithPeriod) SetMac(value *int64)() {
    m.mac = value
}
func (m *GetTeamsDeviceUsageDistributionUserCountsWithPeriod) SetReportPeriod(value *string)() {
    m.reportPeriod = value
}
func (m *GetTeamsDeviceUsageDistributionUserCountsWithPeriod) SetReportRefreshDate(value *string)() {
    m.reportRefreshDate = value
}
func (m *GetTeamsDeviceUsageDistributionUserCountsWithPeriod) SetWeb(value *int64)() {
    m.web = value
}
func (m *GetTeamsDeviceUsageDistributionUserCountsWithPeriod) SetWindows(value *int64)() {
    m.windows = value
}
func (m *GetTeamsDeviceUsageDistributionUserCountsWithPeriod) SetWindowsPhone(value *int64)() {
    m.windowsPhone = value
}
