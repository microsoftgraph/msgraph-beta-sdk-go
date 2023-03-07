package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// WindowsPhone81AppX 
type WindowsPhone81AppX struct {
    MobileLobApp
}
// NewWindowsPhone81AppX instantiates a new WindowsPhone81AppX and sets the default values.
func NewWindowsPhone81AppX()(*WindowsPhone81AppX) {
    m := &WindowsPhone81AppX{
        MobileLobApp: *NewMobileLobApp(),
    }
    odataTypeValue := "#microsoft.graph.windowsPhone81AppX"
    m.SetOdataType(&odataTypeValue)
    return m
}
// CreateWindowsPhone81AppXFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateWindowsPhone81AppXFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    if parseNode != nil {
        mappingValueNode, err := parseNode.GetChildNode("@odata.type")
        if err != nil {
            return nil, err
        }
        if mappingValueNode != nil {
            mappingValue, err := mappingValueNode.GetStringValue()
            if err != nil {
                return nil, err
            }
            if mappingValue != nil {
                switch *mappingValue {
                    case "#microsoft.graph.windowsPhone81AppXBundle":
                        return NewWindowsPhone81AppXBundle(), nil
                }
            }
        }
    }
    return NewWindowsPhone81AppX(), nil
}
// GetApplicableArchitectures gets the applicableArchitectures property value. Contains properties for Windows architecture.
func (m *WindowsPhone81AppX) GetApplicableArchitectures()(*WindowsArchitecture) {
    val, err := m.GetBackingStore().Get("applicableArchitectures")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*WindowsArchitecture)
    }
    return nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *WindowsPhone81AppX) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.MobileLobApp.GetFieldDeserializers()
    res["applicableArchitectures"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseWindowsArchitecture)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetApplicableArchitectures(val.(*WindowsArchitecture))
        }
        return nil
    }
    res["identityName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIdentityName(val)
        }
        return nil
    }
    res["identityPublisherHash"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIdentityPublisherHash(val)
        }
        return nil
    }
    res["identityResourceIdentifier"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIdentityResourceIdentifier(val)
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
    res["minimumSupportedOperatingSystem"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateWindowsMinimumOperatingSystemFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMinimumSupportedOperatingSystem(val.(WindowsMinimumOperatingSystemable))
        }
        return nil
    }
    res["phoneProductIdentifier"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPhoneProductIdentifier(val)
        }
        return nil
    }
    res["phonePublisherId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPhonePublisherId(val)
        }
        return nil
    }
    return res
}
// GetIdentityName gets the identityName property value. The Identity Name.
func (m *WindowsPhone81AppX) GetIdentityName()(*string) {
    val, err := m.GetBackingStore().Get("identityName")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetIdentityPublisherHash gets the identityPublisherHash property value. The Identity Publisher Hash.
func (m *WindowsPhone81AppX) GetIdentityPublisherHash()(*string) {
    val, err := m.GetBackingStore().Get("identityPublisherHash")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetIdentityResourceIdentifier gets the identityResourceIdentifier property value. The Identity Resource Identifier.
func (m *WindowsPhone81AppX) GetIdentityResourceIdentifier()(*string) {
    val, err := m.GetBackingStore().Get("identityResourceIdentifier")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetIdentityVersion gets the identityVersion property value. The identity version.
func (m *WindowsPhone81AppX) GetIdentityVersion()(*string) {
    val, err := m.GetBackingStore().Get("identityVersion")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetMinimumSupportedOperatingSystem gets the minimumSupportedOperatingSystem property value. The minimum operating system required for a Windows mobile app.
func (m *WindowsPhone81AppX) GetMinimumSupportedOperatingSystem()(WindowsMinimumOperatingSystemable) {
    val, err := m.GetBackingStore().Get("minimumSupportedOperatingSystem")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(WindowsMinimumOperatingSystemable)
    }
    return nil
}
// GetPhoneProductIdentifier gets the phoneProductIdentifier property value. The Phone Product Identifier.
func (m *WindowsPhone81AppX) GetPhoneProductIdentifier()(*string) {
    val, err := m.GetBackingStore().Get("phoneProductIdentifier")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetPhonePublisherId gets the phonePublisherId property value. The Phone Publisher Id.
func (m *WindowsPhone81AppX) GetPhonePublisherId()(*string) {
    val, err := m.GetBackingStore().Get("phonePublisherId")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// Serialize serializes information the current object
func (m *WindowsPhone81AppX) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.MobileLobApp.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetApplicableArchitectures() != nil {
        cast := (*m.GetApplicableArchitectures()).String()
        err = writer.WriteStringValue("applicableArchitectures", &cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("identityName", m.GetIdentityName())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("identityPublisherHash", m.GetIdentityPublisherHash())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("identityResourceIdentifier", m.GetIdentityResourceIdentifier())
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
        err = writer.WriteObjectValue("minimumSupportedOperatingSystem", m.GetMinimumSupportedOperatingSystem())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("phoneProductIdentifier", m.GetPhoneProductIdentifier())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("phonePublisherId", m.GetPhonePublisherId())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetApplicableArchitectures sets the applicableArchitectures property value. Contains properties for Windows architecture.
func (m *WindowsPhone81AppX) SetApplicableArchitectures(value *WindowsArchitecture)() {
    err := m.GetBackingStore().Set("applicableArchitectures", value)
    if err != nil {
        panic(err)
    }
}
// SetIdentityName sets the identityName property value. The Identity Name.
func (m *WindowsPhone81AppX) SetIdentityName(value *string)() {
    err := m.GetBackingStore().Set("identityName", value)
    if err != nil {
        panic(err)
    }
}
// SetIdentityPublisherHash sets the identityPublisherHash property value. The Identity Publisher Hash.
func (m *WindowsPhone81AppX) SetIdentityPublisherHash(value *string)() {
    err := m.GetBackingStore().Set("identityPublisherHash", value)
    if err != nil {
        panic(err)
    }
}
// SetIdentityResourceIdentifier sets the identityResourceIdentifier property value. The Identity Resource Identifier.
func (m *WindowsPhone81AppX) SetIdentityResourceIdentifier(value *string)() {
    err := m.GetBackingStore().Set("identityResourceIdentifier", value)
    if err != nil {
        panic(err)
    }
}
// SetIdentityVersion sets the identityVersion property value. The identity version.
func (m *WindowsPhone81AppX) SetIdentityVersion(value *string)() {
    err := m.GetBackingStore().Set("identityVersion", value)
    if err != nil {
        panic(err)
    }
}
// SetMinimumSupportedOperatingSystem sets the minimumSupportedOperatingSystem property value. The minimum operating system required for a Windows mobile app.
func (m *WindowsPhone81AppX) SetMinimumSupportedOperatingSystem(value WindowsMinimumOperatingSystemable)() {
    err := m.GetBackingStore().Set("minimumSupportedOperatingSystem", value)
    if err != nil {
        panic(err)
    }
}
// SetPhoneProductIdentifier sets the phoneProductIdentifier property value. The Phone Product Identifier.
func (m *WindowsPhone81AppX) SetPhoneProductIdentifier(value *string)() {
    err := m.GetBackingStore().Set("phoneProductIdentifier", value)
    if err != nil {
        panic(err)
    }
}
// SetPhonePublisherId sets the phonePublisherId property value. The Phone Publisher Id.
func (m *WindowsPhone81AppX) SetPhonePublisherId(value *string)() {
    err := m.GetBackingStore().Set("phonePublisherId", value)
    if err != nil {
        panic(err)
    }
}
// WindowsPhone81AppXable 
type WindowsPhone81AppXable interface {
    MobileLobAppable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetApplicableArchitectures()(*WindowsArchitecture)
    GetIdentityName()(*string)
    GetIdentityPublisherHash()(*string)
    GetIdentityResourceIdentifier()(*string)
    GetIdentityVersion()(*string)
    GetMinimumSupportedOperatingSystem()(WindowsMinimumOperatingSystemable)
    GetPhoneProductIdentifier()(*string)
    GetPhonePublisherId()(*string)
    SetApplicableArchitectures(value *WindowsArchitecture)()
    SetIdentityName(value *string)()
    SetIdentityPublisherHash(value *string)()
    SetIdentityResourceIdentifier(value *string)()
    SetIdentityVersion(value *string)()
    SetMinimumSupportedOperatingSystem(value WindowsMinimumOperatingSystemable)()
    SetPhoneProductIdentifier(value *string)()
    SetPhonePublisherId(value *string)()
}
