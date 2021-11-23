package getoffice365activationcounts

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph"
)

// GetOffice365ActivationCounts 
type GetOffice365ActivationCounts struct {
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.Entity
    // The activation count on an Android device.
    android *int64;
    // The activation count on iOS.
    ios *int64;
    // The activation count on Mac OS.
    mac *int64;
    // The product type, such as 'Microsoft 365 ProPlus' or 'Project Client'.
    productType *string;
    // The latest date of the content.
    reportRefreshDate *string;
    // The activation count on Windows. This number includes every activation on any Windows computer.
    windows *int64;
    // The activation count on Windows 10 mobile.
    windows10Mobile *int64;
}
// NewGetOffice365ActivationCounts instantiates a new getOffice365ActivationCounts and sets the default values.
func NewGetOffice365ActivationCounts()(*GetOffice365ActivationCounts) {
    m := &GetOffice365ActivationCounts{
        Entity: *i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.NewEntity(),
    }
    return m
}
// GetAndroid gets the android property value. The activation count on an Android device.
func (m *GetOffice365ActivationCounts) GetAndroid()(*int64) {
    if m == nil {
        return nil
    } else {
        return m.android
    }
}
// GetIos gets the ios property value. The activation count on iOS.
func (m *GetOffice365ActivationCounts) GetIos()(*int64) {
    if m == nil {
        return nil
    } else {
        return m.ios
    }
}
// GetMac gets the mac property value. The activation count on Mac OS.
func (m *GetOffice365ActivationCounts) GetMac()(*int64) {
    if m == nil {
        return nil
    } else {
        return m.mac
    }
}
// GetProductType gets the productType property value. The product type, such as 'Microsoft 365 ProPlus' or 'Project Client'.
func (m *GetOffice365ActivationCounts) GetProductType()(*string) {
    if m == nil {
        return nil
    } else {
        return m.productType
    }
}
// GetReportRefreshDate gets the reportRefreshDate property value. The latest date of the content.
func (m *GetOffice365ActivationCounts) GetReportRefreshDate()(*string) {
    if m == nil {
        return nil
    } else {
        return m.reportRefreshDate
    }
}
// GetWindows gets the windows property value. The activation count on Windows. This number includes every activation on any Windows computer.
func (m *GetOffice365ActivationCounts) GetWindows()(*int64) {
    if m == nil {
        return nil
    } else {
        return m.windows
    }
}
// GetWindows10Mobile gets the windows10Mobile property value. The activation count on Windows 10 mobile.
func (m *GetOffice365ActivationCounts) GetWindows10Mobile()(*int64) {
    if m == nil {
        return nil
    } else {
        return m.windows10Mobile
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *GetOffice365ActivationCounts) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["android"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAndroid(val)
        }
        return nil
    }
    res["ios"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIos(val)
        }
        return nil
    }
    res["mac"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMac(val)
        }
        return nil
    }
    res["productType"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetProductType(val)
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
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetWindows(val)
        }
        return nil
    }
    res["windows10Mobile"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetWindows10Mobile(val)
        }
        return nil
    }
    return res
}
func (m *GetOffice365ActivationCounts) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
func (m *GetOffice365ActivationCounts) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteInt64Value("android", m.GetAndroid())
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
        err = writer.WriteInt64Value("mac", m.GetMac())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("productType", m.GetProductType())
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
        err = writer.WriteInt64Value("windows", m.GetWindows())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt64Value("windows10Mobile", m.GetWindows10Mobile())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAndroid sets the android property value. The activation count on an Android device.
func (m *GetOffice365ActivationCounts) SetAndroid(value *int64)() {
    m.android = value
}
// SetIos sets the ios property value. The activation count on iOS.
func (m *GetOffice365ActivationCounts) SetIos(value *int64)() {
    m.ios = value
}
// SetMac sets the mac property value. The activation count on Mac OS.
func (m *GetOffice365ActivationCounts) SetMac(value *int64)() {
    m.mac = value
}
// SetProductType sets the productType property value. The product type, such as 'Microsoft 365 ProPlus' or 'Project Client'.
func (m *GetOffice365ActivationCounts) SetProductType(value *string)() {
    m.productType = value
}
// SetReportRefreshDate sets the reportRefreshDate property value. The latest date of the content.
func (m *GetOffice365ActivationCounts) SetReportRefreshDate(value *string)() {
    m.reportRefreshDate = value
}
// SetWindows sets the windows property value. The activation count on Windows. This number includes every activation on any Windows computer.
func (m *GetOffice365ActivationCounts) SetWindows(value *int64)() {
    m.windows = value
}
// SetWindows10Mobile sets the windows10Mobile property value. The activation count on Windows 10 mobile.
func (m *GetOffice365ActivationCounts) SetWindows10Mobile(value *int64)() {
    m.windows10Mobile = value
}
