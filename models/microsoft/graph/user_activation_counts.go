package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// 
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
// Instantiates a new userActivationCounts and sets the default values.
func NewUserActivationCounts()(*UserActivationCounts) {
    m := &UserActivationCounts{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// Gets the activatedOnSharedComputer property value. True if the user used the product on a shared computer before.
func (m *UserActivationCounts) GetActivatedOnSharedComputer()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.activatedOnSharedComputer
    }
}
// Gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *UserActivationCounts) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// Gets the android property value. The activation count on an Android device.
func (m *UserActivationCounts) GetAndroid()(*int64) {
    if m == nil {
        return nil
    } else {
        return m.android
    }
}
// Gets the ios property value. The activation count on iOS.
func (m *UserActivationCounts) GetIos()(*int64) {
    if m == nil {
        return nil
    } else {
        return m.ios
    }
}
// Gets the lastActivatedDate property value. The date of the latest activation.
func (m *UserActivationCounts) GetLastActivatedDate()(*string) {
    if m == nil {
        return nil
    } else {
        return m.lastActivatedDate
    }
}
// Gets the mac property value. The activation count on Mac OS.
func (m *UserActivationCounts) GetMac()(*int64) {
    if m == nil {
        return nil
    } else {
        return m.mac
    }
}
// Gets the productType property value. The product type, such as 'Microsoft 365 ProPlus'or 'Project Client'.
func (m *UserActivationCounts) GetProductType()(*string) {
    if m == nil {
        return nil
    } else {
        return m.productType
    }
}
// Gets the windows property value. The activation count on Windows. This number includes every activation on any Windows computer.
func (m *UserActivationCounts) GetWindows()(*int64) {
    if m == nil {
        return nil
    } else {
        return m.windows
    }
}
// Gets the windows10Mobile property value. The activation count on Windows 10 mobile.
func (m *UserActivationCounts) GetWindows10Mobile()(*int64) {
    if m == nil {
        return nil
    } else {
        return m.windows10Mobile
    }
}
// The deserialization information for the current model
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
// Serializes information the current object
// Parameters:
//  - writer : Serialization writer to use to serialize this model
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
// Sets the activatedOnSharedComputer property value. True if the user used the product on a shared computer before.
// Parameters:
//  - value : Value to set for the activatedOnSharedComputer property.
func (m *UserActivationCounts) SetActivatedOnSharedComputer(value *bool)() {
    m.activatedOnSharedComputer = value
}
// Sets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// Parameters:
//  - value : Value to set for the AdditionalData property.
func (m *UserActivationCounts) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// Sets the android property value. The activation count on an Android device.
// Parameters:
//  - value : Value to set for the android property.
func (m *UserActivationCounts) SetAndroid(value *int64)() {
    m.android = value
}
// Sets the ios property value. The activation count on iOS.
// Parameters:
//  - value : Value to set for the ios property.
func (m *UserActivationCounts) SetIos(value *int64)() {
    m.ios = value
}
// Sets the lastActivatedDate property value. The date of the latest activation.
// Parameters:
//  - value : Value to set for the lastActivatedDate property.
func (m *UserActivationCounts) SetLastActivatedDate(value *string)() {
    m.lastActivatedDate = value
}
// Sets the mac property value. The activation count on Mac OS.
// Parameters:
//  - value : Value to set for the mac property.
func (m *UserActivationCounts) SetMac(value *int64)() {
    m.mac = value
}
// Sets the productType property value. The product type, such as 'Microsoft 365 ProPlus'or 'Project Client'.
// Parameters:
//  - value : Value to set for the productType property.
func (m *UserActivationCounts) SetProductType(value *string)() {
    m.productType = value
}
// Sets the windows property value. The activation count on Windows. This number includes every activation on any Windows computer.
// Parameters:
//  - value : Value to set for the windows property.
func (m *UserActivationCounts) SetWindows(value *int64)() {
    m.windows = value
}
// Sets the windows10Mobile property value. The activation count on Windows 10 mobile.
// Parameters:
//  - value : Value to set for the windows10Mobile property.
func (m *UserActivationCounts) SetWindows10Mobile(value *int64)() {
    m.windows10Mobile = value
}
