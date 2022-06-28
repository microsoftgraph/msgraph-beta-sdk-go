package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// WindowsAppX 
type WindowsAppX struct {
    MobileLobApp
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // The Windows architecture(s) for which this app can run on. Possible values are: none, x86, x64, arm, neutral, arm64.
    applicableArchitectures *WindowsArchitecture
    // The Identity Name.
    identityName *string
    // The Identity Publisher Hash.
    identityPublisherHash *string
    // The Identity Resource Identifier.
    identityResourceIdentifier *string
    // The identity version.
    identityVersion *string
    // Whether or not the app is a bundle.
    isBundle *bool
    // The minimum operating system required for a Windows mobile app.
    minimumSupportedOperatingSystem WindowsMinimumOperatingSystemable
}
// NewWindowsAppX instantiates a new WindowsAppX and sets the default values.
func NewWindowsAppX()(*WindowsAppX) {
    m := &WindowsAppX{
        MobileLobApp: *NewMobileLobApp(),
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateWindowsAppXFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateWindowsAppXFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewWindowsAppX(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *WindowsAppX) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetApplicableArchitectures gets the applicableArchitectures property value. The Windows architecture(s) for which this app can run on. Possible values are: none, x86, x64, arm, neutral, arm64.
func (m *WindowsAppX) GetApplicableArchitectures()(*WindowsArchitecture) {
    if m == nil {
        return nil
    } else {
        return m.applicableArchitectures
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *WindowsAppX) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
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
    res["isBundle"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsBundle(val)
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
    return res
}
// GetIdentityName gets the identityName property value. The Identity Name.
func (m *WindowsAppX) GetIdentityName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.identityName
    }
}
// GetIdentityPublisherHash gets the identityPublisherHash property value. The Identity Publisher Hash.
func (m *WindowsAppX) GetIdentityPublisherHash()(*string) {
    if m == nil {
        return nil
    } else {
        return m.identityPublisherHash
    }
}
// GetIdentityResourceIdentifier gets the identityResourceIdentifier property value. The Identity Resource Identifier.
func (m *WindowsAppX) GetIdentityResourceIdentifier()(*string) {
    if m == nil {
        return nil
    } else {
        return m.identityResourceIdentifier
    }
}
// GetIdentityVersion gets the identityVersion property value. The identity version.
func (m *WindowsAppX) GetIdentityVersion()(*string) {
    if m == nil {
        return nil
    } else {
        return m.identityVersion
    }
}
// GetIsBundle gets the isBundle property value. Whether or not the app is a bundle.
func (m *WindowsAppX) GetIsBundle()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.isBundle
    }
}
// GetMinimumSupportedOperatingSystem gets the minimumSupportedOperatingSystem property value. The minimum operating system required for a Windows mobile app.
func (m *WindowsAppX) GetMinimumSupportedOperatingSystem()(WindowsMinimumOperatingSystemable) {
    if m == nil {
        return nil
    } else {
        return m.minimumSupportedOperatingSystem
    }
}
// Serialize serializes information the current object
func (m *WindowsAppX) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
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
        err = writer.WriteBoolValue("isBundle", m.GetIsBundle())
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
        err = writer.WriteAdditionalData(m.GetAdditionalData())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAdditionalData sets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *WindowsAppX) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
// SetApplicableArchitectures sets the applicableArchitectures property value. The Windows architecture(s) for which this app can run on. Possible values are: none, x86, x64, arm, neutral, arm64.
func (m *WindowsAppX) SetApplicableArchitectures(value *WindowsArchitecture)() {
    if m != nil {
        m.applicableArchitectures = value
    }
}
// SetIdentityName sets the identityName property value. The Identity Name.
func (m *WindowsAppX) SetIdentityName(value *string)() {
    if m != nil {
        m.identityName = value
    }
}
// SetIdentityPublisherHash sets the identityPublisherHash property value. The Identity Publisher Hash.
func (m *WindowsAppX) SetIdentityPublisherHash(value *string)() {
    if m != nil {
        m.identityPublisherHash = value
    }
}
// SetIdentityResourceIdentifier sets the identityResourceIdentifier property value. The Identity Resource Identifier.
func (m *WindowsAppX) SetIdentityResourceIdentifier(value *string)() {
    if m != nil {
        m.identityResourceIdentifier = value
    }
}
// SetIdentityVersion sets the identityVersion property value. The identity version.
func (m *WindowsAppX) SetIdentityVersion(value *string)() {
    if m != nil {
        m.identityVersion = value
    }
}
// SetIsBundle sets the isBundle property value. Whether or not the app is a bundle.
func (m *WindowsAppX) SetIsBundle(value *bool)() {
    if m != nil {
        m.isBundle = value
    }
}
// SetMinimumSupportedOperatingSystem sets the minimumSupportedOperatingSystem property value. The minimum operating system required for a Windows mobile app.
func (m *WindowsAppX) SetMinimumSupportedOperatingSystem(value WindowsMinimumOperatingSystemable)() {
    if m != nil {
        m.minimumSupportedOperatingSystem = value
    }
}
