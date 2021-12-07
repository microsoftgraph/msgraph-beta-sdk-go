package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// UserActivationCounts 
type UserActivationCounts struct {
    // True if the user used the product on a shared computer before.
    activatedOnSharedComputer *bool;
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // The activation count on an Android device.
    android *int64;
    // The activation count on iOS.
    ios *int64;
    // The date of the latest activation.
    lastActivatedDate *string;
    // The activation count on Mac OS.
    mac *int64;
    // The product type, such as 'Microsoft 365 ProPlus'or 'Project Client'.
    productType *string;
    // The activation count on Windows. This number includes every activation on any Windows computer.
    windows *int64;
    // The activation count on Windows 10 mobile.
    windows10Mobile *int64;
}
// NewUserActivationCounts instantiates a new userActivationCounts and sets the default values.
func NewUserActivationCounts()(*UserActivationCounts) {
    m := &UserActivationCounts{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// GetActivatedOnSharedComputer gets the activatedOnSharedComputer property value. True if the user used the product on a shared computer before.
func (m *UserActivationCounts) GetActivatedOnSharedComputer()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.activatedOnSharedComputer
    }
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *UserActivationCounts) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetAndroid gets the android property value. The activation count on an Android device.
func (m *UserActivationCounts) GetAndroid()(*int64) {
    if m == nil {
        return nil
    } else {
        return m.android
    }
}
// GetIos gets the ios property value. The activation count on iOS.
func (m *UserActivationCounts) GetIos()(*int64) {
    if m == nil {
        return nil
    } else {
        return m.ios
    }
}
// GetLastActivatedDate gets the lastActivatedDate property value. The date of the latest activation.
func (m *UserActivationCounts) GetLastActivatedDate()(*string) {
    if m == nil {
        return nil
    } else {
        return m.lastActivatedDate
    }
}
// GetMac gets the mac property value. The activation count on Mac OS.
func (m *UserActivationCounts) GetMac()(*int64) {
    if m == nil {
        return nil
    } else {
        return m.mac
    }
}
// GetProductType gets the productType property value. The product type, such as 'Microsoft 365 ProPlus'or 'Project Client'.
func (m *UserActivationCounts) GetProductType()(*string) {
    if m == nil {
        return nil
    } else {
        return m.productType
    }
}
// GetWindows gets the windows property value. The activation count on Windows. This number includes every activation on any Windows computer.
func (m *UserActivationCounts) GetWindows()(*int64) {
    if m == nil {
        return nil
    } else {
        return m.windows
    }
}
// GetWindows10Mobile gets the windows10Mobile property value. The activation count on Windows 10 mobile.
func (m *UserActivationCounts) GetWindows10Mobile()(*int64) {
    if m == nil {
        return nil
    } else {
        return m.windows10Mobile
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *UserActivationCounts) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["activatedOnSharedComputer"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetActivatedOnSharedComputer(val)
        }
        return nil
    }
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
    res["lastActivatedDate"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLastActivatedDate(val)
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
func (m *UserActivationCounts) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
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
// SetActivatedOnSharedComputer sets the activatedOnSharedComputer property value. True if the user used the product on a shared computer before.
func (m *UserActivationCounts) SetActivatedOnSharedComputer(value *bool)() {
    if m != nil {
        m.activatedOnSharedComputer = value
    }
}
// SetAdditionalData sets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *UserActivationCounts) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
// SetAndroid sets the android property value. The activation count on an Android device.
func (m *UserActivationCounts) SetAndroid(value *int64)() {
    if m != nil {
        m.android = value
    }
}
// SetIos sets the ios property value. The activation count on iOS.
func (m *UserActivationCounts) SetIos(value *int64)() {
    if m != nil {
        m.ios = value
    }
}
// SetLastActivatedDate sets the lastActivatedDate property value. The date of the latest activation.
func (m *UserActivationCounts) SetLastActivatedDate(value *string)() {
    if m != nil {
        m.lastActivatedDate = value
    }
}
// SetMac sets the mac property value. The activation count on Mac OS.
func (m *UserActivationCounts) SetMac(value *int64)() {
    if m != nil {
        m.mac = value
    }
}
// SetProductType sets the productType property value. The product type, such as 'Microsoft 365 ProPlus'or 'Project Client'.
func (m *UserActivationCounts) SetProductType(value *string)() {
    if m != nil {
        m.productType = value
    }
}
// SetWindows sets the windows property value. The activation count on Windows. This number includes every activation on any Windows computer.
func (m *UserActivationCounts) SetWindows(value *int64)() {
    if m != nil {
        m.windows = value
    }
}
// SetWindows10Mobile sets the windows10Mobile property value. The activation count on Windows 10 mobile.
func (m *UserActivationCounts) SetWindows10Mobile(value *int64)() {
    if m != nil {
        m.windows10Mobile = value
    }
}
