package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// WindowsPhone81SCEPCertificateProfile 
type WindowsPhone81SCEPCertificateProfile struct {
    WindowsPhone81CertificateProfileBase
}
// NewWindowsPhone81SCEPCertificateProfile instantiates a new WindowsPhone81SCEPCertificateProfile and sets the default values.
func NewWindowsPhone81SCEPCertificateProfile()(*WindowsPhone81SCEPCertificateProfile) {
    m := &WindowsPhone81SCEPCertificateProfile{
        WindowsPhone81CertificateProfileBase: *NewWindowsPhone81CertificateProfileBase(),
    }
    odataTypeValue := "#microsoft.graph.windowsPhone81SCEPCertificateProfile"
    m.SetOdataType(&odataTypeValue)
    return m
}
// CreateWindowsPhone81SCEPCertificateProfileFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateWindowsPhone81SCEPCertificateProfileFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewWindowsPhone81SCEPCertificateProfile(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *WindowsPhone81SCEPCertificateProfile) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.WindowsPhone81CertificateProfileBase.GetFieldDeserializers()
    res["hashAlgorithm"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseHashAlgorithms)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetHashAlgorithm(val.(*HashAlgorithms))
        }
        return nil
    }
    res["keySize"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseKeySize)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetKeySize(val.(*KeySize))
        }
        return nil
    }
    res["keyUsage"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseKeyUsages)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetKeyUsage(val.(*KeyUsages))
        }
        return nil
    }
    res["managedDeviceCertificateStates"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateManagedDeviceCertificateStateFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]ManagedDeviceCertificateStateable, len(val))
            for i, v := range val {
                res[i] = v.(ManagedDeviceCertificateStateable)
            }
            m.SetManagedDeviceCertificateStates(res)
        }
        return nil
    }
    res["rootCertificate"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateWindowsPhone81TrustedRootCertificateFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRootCertificate(val.(WindowsPhone81TrustedRootCertificateable))
        }
        return nil
    }
    res["scepServerUrls"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]string, len(val))
            for i, v := range val {
                res[i] = *(v.(*string))
            }
            m.SetScepServerUrls(res)
        }
        return nil
    }
    res["subjectAlternativeNameFormatString"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSubjectAlternativeNameFormatString(val)
        }
        return nil
    }
    res["subjectNameFormatString"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSubjectNameFormatString(val)
        }
        return nil
    }
    return res
}
// GetHashAlgorithm gets the hashAlgorithm property value. Hash Algorithm Options.
func (m *WindowsPhone81SCEPCertificateProfile) GetHashAlgorithm()(*HashAlgorithms) {
    val, err := m.GetBackingStore().Get("hashAlgorithm")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*HashAlgorithms)
    }
    return nil
}
// GetKeySize gets the keySize property value. Key Size Options.
func (m *WindowsPhone81SCEPCertificateProfile) GetKeySize()(*KeySize) {
    val, err := m.GetBackingStore().Get("keySize")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*KeySize)
    }
    return nil
}
// GetKeyUsage gets the keyUsage property value. Key Usage Options.
func (m *WindowsPhone81SCEPCertificateProfile) GetKeyUsage()(*KeyUsages) {
    val, err := m.GetBackingStore().Get("keyUsage")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*KeyUsages)
    }
    return nil
}
// GetManagedDeviceCertificateStates gets the managedDeviceCertificateStates property value. Certificate state for devices. This collection can contain a maximum of 2147483647 elements.
func (m *WindowsPhone81SCEPCertificateProfile) GetManagedDeviceCertificateStates()([]ManagedDeviceCertificateStateable) {
    val, err := m.GetBackingStore().Get("managedDeviceCertificateStates")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]ManagedDeviceCertificateStateable)
    }
    return nil
}
// GetRootCertificate gets the rootCertificate property value. Trusted Root Certificate.
func (m *WindowsPhone81SCEPCertificateProfile) GetRootCertificate()(WindowsPhone81TrustedRootCertificateable) {
    val, err := m.GetBackingStore().Get("rootCertificate")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(WindowsPhone81TrustedRootCertificateable)
    }
    return nil
}
// GetScepServerUrls gets the scepServerUrls property value. SCEP Server Url(s).
func (m *WindowsPhone81SCEPCertificateProfile) GetScepServerUrls()([]string) {
    val, err := m.GetBackingStore().Get("scepServerUrls")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]string)
    }
    return nil
}
// GetSubjectAlternativeNameFormatString gets the subjectAlternativeNameFormatString property value. Custom String that defines the AAD Attribute.
func (m *WindowsPhone81SCEPCertificateProfile) GetSubjectAlternativeNameFormatString()(*string) {
    val, err := m.GetBackingStore().Get("subjectAlternativeNameFormatString")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetSubjectNameFormatString gets the subjectNameFormatString property value. Custom format to use with SubjectNameFormat = Custom. Example: CN={{EmailAddress}},E={{EmailAddress}},OU=Enterprise Users,O=Contoso Corporation,L=Redmond,ST=WA,C=US
func (m *WindowsPhone81SCEPCertificateProfile) GetSubjectNameFormatString()(*string) {
    val, err := m.GetBackingStore().Get("subjectNameFormatString")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// Serialize serializes information the current object
func (m *WindowsPhone81SCEPCertificateProfile) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.WindowsPhone81CertificateProfileBase.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetHashAlgorithm() != nil {
        cast := (*m.GetHashAlgorithm()).String()
        err = writer.WriteStringValue("hashAlgorithm", &cast)
        if err != nil {
            return err
        }
    }
    if m.GetKeySize() != nil {
        cast := (*m.GetKeySize()).String()
        err = writer.WriteStringValue("keySize", &cast)
        if err != nil {
            return err
        }
    }
    if m.GetKeyUsage() != nil {
        cast := (*m.GetKeyUsage()).String()
        err = writer.WriteStringValue("keyUsage", &cast)
        if err != nil {
            return err
        }
    }
    if m.GetManagedDeviceCertificateStates() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetManagedDeviceCertificateStates()))
        for i, v := range m.GetManagedDeviceCertificateStates() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("managedDeviceCertificateStates", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("rootCertificate", m.GetRootCertificate())
        if err != nil {
            return err
        }
    }
    if m.GetScepServerUrls() != nil {
        err = writer.WriteCollectionOfStringValues("scepServerUrls", m.GetScepServerUrls())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("subjectAlternativeNameFormatString", m.GetSubjectAlternativeNameFormatString())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("subjectNameFormatString", m.GetSubjectNameFormatString())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetHashAlgorithm sets the hashAlgorithm property value. Hash Algorithm Options.
func (m *WindowsPhone81SCEPCertificateProfile) SetHashAlgorithm(value *HashAlgorithms)() {
    err := m.GetBackingStore().Set("hashAlgorithm", value)
    if err != nil {
        panic(err)
    }
}
// SetKeySize sets the keySize property value. Key Size Options.
func (m *WindowsPhone81SCEPCertificateProfile) SetKeySize(value *KeySize)() {
    err := m.GetBackingStore().Set("keySize", value)
    if err != nil {
        panic(err)
    }
}
// SetKeyUsage sets the keyUsage property value. Key Usage Options.
func (m *WindowsPhone81SCEPCertificateProfile) SetKeyUsage(value *KeyUsages)() {
    err := m.GetBackingStore().Set("keyUsage", value)
    if err != nil {
        panic(err)
    }
}
// SetManagedDeviceCertificateStates sets the managedDeviceCertificateStates property value. Certificate state for devices. This collection can contain a maximum of 2147483647 elements.
func (m *WindowsPhone81SCEPCertificateProfile) SetManagedDeviceCertificateStates(value []ManagedDeviceCertificateStateable)() {
    err := m.GetBackingStore().Set("managedDeviceCertificateStates", value)
    if err != nil {
        panic(err)
    }
}
// SetRootCertificate sets the rootCertificate property value. Trusted Root Certificate.
func (m *WindowsPhone81SCEPCertificateProfile) SetRootCertificate(value WindowsPhone81TrustedRootCertificateable)() {
    err := m.GetBackingStore().Set("rootCertificate", value)
    if err != nil {
        panic(err)
    }
}
// SetScepServerUrls sets the scepServerUrls property value. SCEP Server Url(s).
func (m *WindowsPhone81SCEPCertificateProfile) SetScepServerUrls(value []string)() {
    err := m.GetBackingStore().Set("scepServerUrls", value)
    if err != nil {
        panic(err)
    }
}
// SetSubjectAlternativeNameFormatString sets the subjectAlternativeNameFormatString property value. Custom String that defines the AAD Attribute.
func (m *WindowsPhone81SCEPCertificateProfile) SetSubjectAlternativeNameFormatString(value *string)() {
    err := m.GetBackingStore().Set("subjectAlternativeNameFormatString", value)
    if err != nil {
        panic(err)
    }
}
// SetSubjectNameFormatString sets the subjectNameFormatString property value. Custom format to use with SubjectNameFormat = Custom. Example: CN={{EmailAddress}},E={{EmailAddress}},OU=Enterprise Users,O=Contoso Corporation,L=Redmond,ST=WA,C=US
func (m *WindowsPhone81SCEPCertificateProfile) SetSubjectNameFormatString(value *string)() {
    err := m.GetBackingStore().Set("subjectNameFormatString", value)
    if err != nil {
        panic(err)
    }
}
// WindowsPhone81SCEPCertificateProfileable 
type WindowsPhone81SCEPCertificateProfileable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    WindowsPhone81CertificateProfileBaseable
    GetHashAlgorithm()(*HashAlgorithms)
    GetKeySize()(*KeySize)
    GetKeyUsage()(*KeyUsages)
    GetManagedDeviceCertificateStates()([]ManagedDeviceCertificateStateable)
    GetRootCertificate()(WindowsPhone81TrustedRootCertificateable)
    GetScepServerUrls()([]string)
    GetSubjectAlternativeNameFormatString()(*string)
    GetSubjectNameFormatString()(*string)
    SetHashAlgorithm(value *HashAlgorithms)()
    SetKeySize(value *KeySize)()
    SetKeyUsage(value *KeyUsages)()
    SetManagedDeviceCertificateStates(value []ManagedDeviceCertificateStateable)()
    SetRootCertificate(value WindowsPhone81TrustedRootCertificateable)()
    SetScepServerUrls(value []string)()
    SetSubjectAlternativeNameFormatString(value *string)()
    SetSubjectNameFormatString(value *string)()
}
