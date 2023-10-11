package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e "github.com/microsoft/kiota-abstractions-go/store"
)

// WindowsPackageInformation contains properties for the package information for a Windows line of business app. Used as property within windowsPhone81AppXBundle object, which is also being deprecated. This complex type will be deprecated in February 2023.
type WindowsPackageInformation struct {
    // Stores model information.
    backingStore ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore
}
// NewWindowsPackageInformation instantiates a new windowsPackageInformation and sets the default values.
func NewWindowsPackageInformation()(*WindowsPackageInformation) {
    m := &WindowsPackageInformation{
    }
    m.backingStore = ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStoreFactoryInstance();
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateWindowsPackageInformationFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateWindowsPackageInformationFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewWindowsPackageInformation(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *WindowsPackageInformation) GetAdditionalData()(map[string]any) {
    val , err :=  m.backingStore.Get("additionalData")
    if err != nil {
        panic(err)
    }
    if val == nil {
        var value = make(map[string]any);
        m.SetAdditionalData(value);
    }
    return val.(map[string]any)
}
// GetApplicableArchitecture gets the applicableArchitecture property value. Contains properties for Windows architecture.
func (m *WindowsPackageInformation) GetApplicableArchitecture()(*WindowsArchitecture) {
    val, err := m.GetBackingStore().Get("applicableArchitecture")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*WindowsArchitecture)
    }
    return nil
}
// GetBackingStore gets the BackingStore property value. Stores model information.
func (m *WindowsPackageInformation) GetBackingStore()(ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore) {
    return m.backingStore
}
// GetDisplayName gets the displayName property value. The Display Name.
func (m *WindowsPackageInformation) GetDisplayName()(*string) {
    val, err := m.GetBackingStore().Get("displayName")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *WindowsPackageInformation) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["applicableArchitecture"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseWindowsArchitecture)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetApplicableArchitecture(val.(*WindowsArchitecture))
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
    res["identityPublisher"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIdentityPublisher(val)
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
    res["@odata.type"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOdataType(val)
        }
        return nil
    }
    return res
}
// GetIdentityName gets the identityName property value. The Identity Name.
func (m *WindowsPackageInformation) GetIdentityName()(*string) {
    val, err := m.GetBackingStore().Get("identityName")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetIdentityPublisher gets the identityPublisher property value. The Identity Publisher.
func (m *WindowsPackageInformation) GetIdentityPublisher()(*string) {
    val, err := m.GetBackingStore().Get("identityPublisher")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetIdentityResourceIdentifier gets the identityResourceIdentifier property value. The Identity Resource Identifier.
func (m *WindowsPackageInformation) GetIdentityResourceIdentifier()(*string) {
    val, err := m.GetBackingStore().Get("identityResourceIdentifier")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetIdentityVersion gets the identityVersion property value. The Identity Version.
func (m *WindowsPackageInformation) GetIdentityVersion()(*string) {
    val, err := m.GetBackingStore().Get("identityVersion")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetMinimumSupportedOperatingSystem gets the minimumSupportedOperatingSystem property value. The value for the minimum applicable operating system.
func (m *WindowsPackageInformation) GetMinimumSupportedOperatingSystem()(WindowsMinimumOperatingSystemable) {
    val, err := m.GetBackingStore().Get("minimumSupportedOperatingSystem")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(WindowsMinimumOperatingSystemable)
    }
    return nil
}
// GetOdataType gets the @odata.type property value. The OdataType property
func (m *WindowsPackageInformation) GetOdataType()(*string) {
    val, err := m.GetBackingStore().Get("odataType")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// Serialize serializes information the current object
func (m *WindowsPackageInformation) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    if m.GetApplicableArchitecture() != nil {
        cast := (*m.GetApplicableArchitecture()).String()
        err := writer.WriteStringValue("applicableArchitecture", &cast)
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
        err := writer.WriteStringValue("identityName", m.GetIdentityName())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("identityPublisher", m.GetIdentityPublisher())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("identityResourceIdentifier", m.GetIdentityResourceIdentifier())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("identityVersion", m.GetIdentityVersion())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("minimumSupportedOperatingSystem", m.GetMinimumSupportedOperatingSystem())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("@odata.type", m.GetOdataType())
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
// SetAdditionalData sets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *WindowsPackageInformation) SetAdditionalData(value map[string]any)() {
    err := m.GetBackingStore().Set("additionalData", value)
    if err != nil {
        panic(err)
    }
}
// SetApplicableArchitecture sets the applicableArchitecture property value. Contains properties for Windows architecture.
func (m *WindowsPackageInformation) SetApplicableArchitecture(value *WindowsArchitecture)() {
    err := m.GetBackingStore().Set("applicableArchitecture", value)
    if err != nil {
        panic(err)
    }
}
// SetBackingStore sets the BackingStore property value. Stores model information.
func (m *WindowsPackageInformation) SetBackingStore(value ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)() {
    m.backingStore = value
}
// SetDisplayName sets the displayName property value. The Display Name.
func (m *WindowsPackageInformation) SetDisplayName(value *string)() {
    err := m.GetBackingStore().Set("displayName", value)
    if err != nil {
        panic(err)
    }
}
// SetIdentityName sets the identityName property value. The Identity Name.
func (m *WindowsPackageInformation) SetIdentityName(value *string)() {
    err := m.GetBackingStore().Set("identityName", value)
    if err != nil {
        panic(err)
    }
}
// SetIdentityPublisher sets the identityPublisher property value. The Identity Publisher.
func (m *WindowsPackageInformation) SetIdentityPublisher(value *string)() {
    err := m.GetBackingStore().Set("identityPublisher", value)
    if err != nil {
        panic(err)
    }
}
// SetIdentityResourceIdentifier sets the identityResourceIdentifier property value. The Identity Resource Identifier.
func (m *WindowsPackageInformation) SetIdentityResourceIdentifier(value *string)() {
    err := m.GetBackingStore().Set("identityResourceIdentifier", value)
    if err != nil {
        panic(err)
    }
}
// SetIdentityVersion sets the identityVersion property value. The Identity Version.
func (m *WindowsPackageInformation) SetIdentityVersion(value *string)() {
    err := m.GetBackingStore().Set("identityVersion", value)
    if err != nil {
        panic(err)
    }
}
// SetMinimumSupportedOperatingSystem sets the minimumSupportedOperatingSystem property value. The value for the minimum applicable operating system.
func (m *WindowsPackageInformation) SetMinimumSupportedOperatingSystem(value WindowsMinimumOperatingSystemable)() {
    err := m.GetBackingStore().Set("minimumSupportedOperatingSystem", value)
    if err != nil {
        panic(err)
    }
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *WindowsPackageInformation) SetOdataType(value *string)() {
    err := m.GetBackingStore().Set("odataType", value)
    if err != nil {
        panic(err)
    }
}
// WindowsPackageInformationable 
type WindowsPackageInformationable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackedModel
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetApplicableArchitecture()(*WindowsArchitecture)
    GetBackingStore()(ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)
    GetDisplayName()(*string)
    GetIdentityName()(*string)
    GetIdentityPublisher()(*string)
    GetIdentityResourceIdentifier()(*string)
    GetIdentityVersion()(*string)
    GetMinimumSupportedOperatingSystem()(WindowsMinimumOperatingSystemable)
    GetOdataType()(*string)
    SetApplicableArchitecture(value *WindowsArchitecture)()
    SetBackingStore(value ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)()
    SetDisplayName(value *string)()
    SetIdentityName(value *string)()
    SetIdentityPublisher(value *string)()
    SetIdentityResourceIdentifier(value *string)()
    SetIdentityVersion(value *string)()
    SetMinimumSupportedOperatingSystem(value WindowsMinimumOperatingSystemable)()
    SetOdataType(value *string)()
}
