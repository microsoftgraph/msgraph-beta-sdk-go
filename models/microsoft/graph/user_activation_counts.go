package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

type UserActivationCounts struct {
    activatedOnSharedComputer *bool;
    additionalData map[string]interface{};
    android *int64;
    ios *int64;
    lastActivatedDate *string;
    mac *int64;
    productType *string;
    windows *int64;
    windows10Mobile *int64;
}
func NewUserActivationCounts()(*UserActivationCounts) {
    m := &UserActivationCounts{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
func (m *UserActivationCounts) GetActivatedOnSharedComputer()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.activatedOnSharedComputer
    }
}
func (m *UserActivationCounts) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
func (m *UserActivationCounts) GetAndroid()(*int64) {
    if m == nil {
        return nil
    } else {
        return m.android
    }
}
func (m *UserActivationCounts) GetIos()(*int64) {
    if m == nil {
        return nil
    } else {
        return m.ios
    }
}
func (m *UserActivationCounts) GetLastActivatedDate()(*string) {
    if m == nil {
        return nil
    } else {
        return m.lastActivatedDate
    }
}
func (m *UserActivationCounts) GetMac()(*int64) {
    if m == nil {
        return nil
    } else {
        return m.mac
    }
}
func (m *UserActivationCounts) GetProductType()(*string) {
    if m == nil {
        return nil
    } else {
        return m.productType
    }
}
func (m *UserActivationCounts) GetWindows()(*int64) {
    if m == nil {
        return nil
    } else {
        return m.windows
    }
}
func (m *UserActivationCounts) GetWindows10Mobile()(*int64) {
    if m == nil {
        return nil
    } else {
        return m.windows10Mobile
    }
}
func (m *UserActivationCounts) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["activatedOnSharedComputer"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        m.SetActivatedOnSharedComputer(val)
        return nil
    }
    res["android"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        m.SetAndroid(val)
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
    res["lastActivatedDate"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetLastActivatedDate(val)
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
    res["productType"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetProductType(val)
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
    res["windows10Mobile"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        m.SetWindows10Mobile(val)
        return nil
    }
    return res
}
func (m *UserActivationCounts) IsNil()(bool) {
    return m == nil
}
func (m *UserActivationCounts) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    {
        err := writer.WriteBoolValue("activatedOnSharedComputer", m.GetActivatedOnSharedComputer())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt64Value("android", m.GetAndroid())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt64Value("ios", m.GetIos())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("lastActivatedDate", m.GetLastActivatedDate())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt64Value("mac", m.GetMac())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("productType", m.GetProductType())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt64Value("windows", m.GetWindows())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt64Value("windows10Mobile", m.GetWindows10Mobile())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteAdditionalData(m.GetAdditionalData())
        if err != nil {
            return err
        }
    }
    return nil
}
func (m *UserActivationCounts) SetActivatedOnSharedComputer(value *bool)() {
    m.activatedOnSharedComputer = value
}
func (m *UserActivationCounts) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
func (m *UserActivationCounts) SetAndroid(value *int64)() {
    m.android = value
}
func (m *UserActivationCounts) SetIos(value *int64)() {
    m.ios = value
}
func (m *UserActivationCounts) SetLastActivatedDate(value *string)() {
    m.lastActivatedDate = value
}
func (m *UserActivationCounts) SetMac(value *int64)() {
    m.mac = value
}
func (m *UserActivationCounts) SetProductType(value *string)() {
    m.productType = value
}
func (m *UserActivationCounts) SetWindows(value *int64)() {
    m.windows = value
}
func (m *UserActivationCounts) SetWindows10Mobile(value *int64)() {
    m.windows10Mobile = value
}
