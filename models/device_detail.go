package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// DeviceDetail 
type DeviceDetail struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // Indicates the browser information of the used for signing in.
    browser *string
    // The browserId property
    browserId *string
    // Refers to the UniqueID of the device used for signing in.
    deviceId *string
    // Refers to the name of the device used for signing in.
    displayName *string
    // Indicates whether the device is compliant.
    isCompliant *bool
    // Indicates whether the device is managed.
    isManaged *bool
    // Indicates the operating system name and version used for signing in.
    operatingSystem *string
    // Provides information about whether the signed-in device is Workplace Joined, AzureAD Joined, Domain Joined.
    trustType *string
}
// NewDeviceDetail instantiates a new deviceDetail and sets the default values.
func NewDeviceDetail()(*DeviceDetail) {
    m := &DeviceDetail{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateDeviceDetailFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateDeviceDetailFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewDeviceDetail(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *DeviceDetail) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetBrowser gets the browser property value. Indicates the browser information of the used for signing in.
func (m *DeviceDetail) GetBrowser()(*string) {
    if m == nil {
        return nil
    } else {
        return m.browser
    }
}
// GetBrowserId gets the browserId property value. The browserId property
func (m *DeviceDetail) GetBrowserId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.browserId
    }
}
// GetDeviceId gets the deviceId property value. Refers to the UniqueID of the device used for signing in.
func (m *DeviceDetail) GetDeviceId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.deviceId
    }
}
// GetDisplayName gets the displayName property value. Refers to the name of the device used for signing in.
func (m *DeviceDetail) GetDisplayName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.displayName
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *DeviceDetail) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["browser"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetBrowser(val)
        }
        return nil
    }
    res["browserId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetBrowserId(val)
        }
        return nil
    }
    res["deviceId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDeviceId(val)
        }
        return nil
    }
    res["displayName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDisplayName(val)
        }
        return nil
    }
    res["isCompliant"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsCompliant(val)
        }
        return nil
    }
    res["isManaged"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsManaged(val)
        }
        return nil
    }
    res["operatingSystem"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOperatingSystem(val)
        }
        return nil
    }
    res["trustType"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTrustType(val)
        }
        return nil
    }
    return res
}
// GetIsCompliant gets the isCompliant property value. Indicates whether the device is compliant.
func (m *DeviceDetail) GetIsCompliant()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.isCompliant
    }
}
// GetIsManaged gets the isManaged property value. Indicates whether the device is managed.
func (m *DeviceDetail) GetIsManaged()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.isManaged
    }
}
// GetOperatingSystem gets the operatingSystem property value. Indicates the operating system name and version used for signing in.
func (m *DeviceDetail) GetOperatingSystem()(*string) {
    if m == nil {
        return nil
    } else {
        return m.operatingSystem
    }
}
// GetTrustType gets the trustType property value. Provides information about whether the signed-in device is Workplace Joined, AzureAD Joined, Domain Joined.
func (m *DeviceDetail) GetTrustType()(*string) {
    if m == nil {
        return nil
    } else {
        return m.trustType
    }
}
// Serialize serializes information the current object
func (m *DeviceDetail) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("browser", m.GetBrowser())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("browserId", m.GetBrowserId())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("deviceId", m.GetDeviceId())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("displayName", m.GetDisplayName())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("isCompliant", m.GetIsCompliant())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("isManaged", m.GetIsManaged())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("operatingSystem", m.GetOperatingSystem())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("trustType", m.GetTrustType())
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
// SetAdditionalData sets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *DeviceDetail) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
// SetBrowser sets the browser property value. Indicates the browser information of the used for signing in.
func (m *DeviceDetail) SetBrowser(value *string)() {
    if m != nil {
        m.browser = value
    }
}
// SetBrowserId sets the browserId property value. The browserId property
func (m *DeviceDetail) SetBrowserId(value *string)() {
    if m != nil {
        m.browserId = value
    }
}
// SetDeviceId sets the deviceId property value. Refers to the UniqueID of the device used for signing in.
func (m *DeviceDetail) SetDeviceId(value *string)() {
    if m != nil {
        m.deviceId = value
    }
}
// SetDisplayName sets the displayName property value. Refers to the name of the device used for signing in.
func (m *DeviceDetail) SetDisplayName(value *string)() {
    if m != nil {
        m.displayName = value
    }
}
// SetIsCompliant sets the isCompliant property value. Indicates whether the device is compliant.
func (m *DeviceDetail) SetIsCompliant(value *bool)() {
    if m != nil {
        m.isCompliant = value
    }
}
// SetIsManaged sets the isManaged property value. Indicates whether the device is managed.
func (m *DeviceDetail) SetIsManaged(value *bool)() {
    if m != nil {
        m.isManaged = value
    }
}
// SetOperatingSystem sets the operatingSystem property value. Indicates the operating system name and version used for signing in.
func (m *DeviceDetail) SetOperatingSystem(value *string)() {
    if m != nil {
        m.operatingSystem = value
    }
}
// SetTrustType sets the trustType property value. Provides information about whether the signed-in device is Workplace Joined, AzureAD Joined, Domain Joined.
func (m *DeviceDetail) SetTrustType(value *string)() {
    if m != nil {
        m.trustType = value
    }
}
