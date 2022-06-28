package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// WindowsMobileMSI 
type WindowsMobileMSI struct {
    MobileLobApp
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // The command line.
    commandLine *string
    // The identity version.
    identityVersion *string
    // A boolean to control whether the app's version will be used to detect the app after it is installed on a device. Set this to true for Windows Mobile MSI Line of Business (LoB) apps that use a self update feature.
    ignoreVersionDetection *bool
    // The product code.
    productCode *string
    // The product version of Windows Mobile MSI Line of Business (LoB) app.
    productVersion *string
    // Indicates whether to install a dual-mode MSI in the device context. If true, app will be installed for all users. If false, app will be installed per-user. If null, service will use the MSI package's default install context. In case of dual-mode MSI, this default will be per-user.  Cannot be set for non-dual-mode apps.  Cannot be changed after initial creation of the application.
    useDeviceContext *bool
}
// NewWindowsMobileMSI instantiates a new WindowsMobileMSI and sets the default values.
func NewWindowsMobileMSI()(*WindowsMobileMSI) {
    m := &WindowsMobileMSI{
        MobileLobApp: *NewMobileLobApp(),
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateWindowsMobileMSIFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateWindowsMobileMSIFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewWindowsMobileMSI(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *WindowsMobileMSI) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetCommandLine gets the commandLine property value. The command line.
func (m *WindowsMobileMSI) GetCommandLine()(*string) {
    if m == nil {
        return nil
    } else {
        return m.commandLine
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *WindowsMobileMSI) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.MobileLobApp.GetFieldDeserializers()
    res["commandLine"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCommandLine(val)
        }
        return nil
    }
    res["identityVersion"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIdentityVersion(val)
        }
        return nil
    }
    res["ignoreVersionDetection"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIgnoreVersionDetection(val)
        }
        return nil
    }
    res["productCode"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetProductCode(val)
        }
        return nil
    }
    res["productVersion"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetProductVersion(val)
        }
        return nil
    }
    res["useDeviceContext"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUseDeviceContext(val)
        }
        return nil
    }
    return res
}
// GetIdentityVersion gets the identityVersion property value. The identity version.
func (m *WindowsMobileMSI) GetIdentityVersion()(*string) {
    if m == nil {
        return nil
    } else {
        return m.identityVersion
    }
}
// GetIgnoreVersionDetection gets the ignoreVersionDetection property value. A boolean to control whether the app's version will be used to detect the app after it is installed on a device. Set this to true for Windows Mobile MSI Line of Business (LoB) apps that use a self update feature.
func (m *WindowsMobileMSI) GetIgnoreVersionDetection()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.ignoreVersionDetection
    }
}
// GetProductCode gets the productCode property value. The product code.
func (m *WindowsMobileMSI) GetProductCode()(*string) {
    if m == nil {
        return nil
    } else {
        return m.productCode
    }
}
// GetProductVersion gets the productVersion property value. The product version of Windows Mobile MSI Line of Business (LoB) app.
func (m *WindowsMobileMSI) GetProductVersion()(*string) {
    if m == nil {
        return nil
    } else {
        return m.productVersion
    }
}
// GetUseDeviceContext gets the useDeviceContext property value. Indicates whether to install a dual-mode MSI in the device context. If true, app will be installed for all users. If false, app will be installed per-user. If null, service will use the MSI package's default install context. In case of dual-mode MSI, this default will be per-user.  Cannot be set for non-dual-mode apps.  Cannot be changed after initial creation of the application.
func (m *WindowsMobileMSI) GetUseDeviceContext()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.useDeviceContext
    }
}
// Serialize serializes information the current object
func (m *WindowsMobileMSI) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.MobileLobApp.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteStringValue("commandLine", m.GetCommandLine())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("identityVersion", m.GetIdentityVersion())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("ignoreVersionDetection", m.GetIgnoreVersionDetection())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("productCode", m.GetProductCode())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("productVersion", m.GetProductVersion())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("useDeviceContext", m.GetUseDeviceContext())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteAdditionalData(m.GetAdditionalData())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAdditionalData sets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *WindowsMobileMSI) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
// SetCommandLine sets the commandLine property value. The command line.
func (m *WindowsMobileMSI) SetCommandLine(value *string)() {
    if m != nil {
        m.commandLine = value
    }
}
// SetIdentityVersion sets the identityVersion property value. The identity version.
func (m *WindowsMobileMSI) SetIdentityVersion(value *string)() {
    if m != nil {
        m.identityVersion = value
    }
}
// SetIgnoreVersionDetection sets the ignoreVersionDetection property value. A boolean to control whether the app's version will be used to detect the app after it is installed on a device. Set this to true for Windows Mobile MSI Line of Business (LoB) apps that use a self update feature.
func (m *WindowsMobileMSI) SetIgnoreVersionDetection(value *bool)() {
    if m != nil {
        m.ignoreVersionDetection = value
    }
}
// SetProductCode sets the productCode property value. The product code.
func (m *WindowsMobileMSI) SetProductCode(value *string)() {
    if m != nil {
        m.productCode = value
    }
}
// SetProductVersion sets the productVersion property value. The product version of Windows Mobile MSI Line of Business (LoB) app.
func (m *WindowsMobileMSI) SetProductVersion(value *string)() {
    if m != nil {
        m.productVersion = value
    }
}
// SetUseDeviceContext sets the useDeviceContext property value. Indicates whether to install a dual-mode MSI in the device context. If true, app will be installed for all users. If false, app will be installed per-user. If null, service will use the MSI package's default install context. In case of dual-mode MSI, this default will be per-user.  Cannot be set for non-dual-mode apps.  Cannot be changed after initial creation of the application.
func (m *WindowsMobileMSI) SetUseDeviceContext(value *bool)() {
    if m != nil {
        m.useDeviceContext = value
    }
}
