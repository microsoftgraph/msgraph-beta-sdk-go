package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// Windows81SCEPCertificateProfile windows 8.1+ SCEP certificate profile
type Windows81SCEPCertificateProfile struct {
    Windows81CertificateProfileBase
}
// NewWindows81SCEPCertificateProfile instantiates a new windows81SCEPCertificateProfile and sets the default values.
func NewWindows81SCEPCertificateProfile()(*Windows81SCEPCertificateProfile) {
    m := &Windows81SCEPCertificateProfile{
        Windows81CertificateProfileBase: *NewWindows81CertificateProfileBase(),
    }
    odataTypeValue := "#microsoft.graph.windows81SCEPCertificateProfile"
    m.SetOdataType(&odataTypeValue)
    return m
}
// CreateWindows81SCEPCertificateProfileFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateWindows81SCEPCertificateProfileFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewWindows81SCEPCertificateProfile(), nil
}
// GetCertificateStore gets the certificateStore property value. Target store certificate. Possible values are: user, machine.
func (m *Windows81SCEPCertificateProfile) GetCertificateStore()(*Windows81SCEPCertificateProfile_certificateStore) {
    val, err := m.GetBackingStore().Get("certificateStore")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*Windows81SCEPCertificateProfile_certificateStore)
    }
    return nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *Windows81SCEPCertificateProfile) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Windows81CertificateProfileBase.GetFieldDeserializers()
    res["certificateStore"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseWindows81SCEPCertificateProfile_certificateStore)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCertificateStore(val.(*Windows81SCEPCertificateProfile_certificateStore))
        }
        return nil
    }
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
                if v != nil {
                    res[i] = v.(ManagedDeviceCertificateStateable)
                }
            }
            m.SetManagedDeviceCertificateStates(res)
        }
        return nil
    }
    res["rootCertificate"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateWindows81TrustedRootCertificateFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRootCertificate(val.(Windows81TrustedRootCertificateable))
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
                if v != nil {
                    res[i] = *(v.(*string))
                }
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
func (m *Windows81SCEPCertificateProfile) GetHashAlgorithm()(*HashAlgorithms) {
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
func (m *Windows81SCEPCertificateProfile) GetKeySize()(*KeySize) {
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
func (m *Windows81SCEPCertificateProfile) GetKeyUsage()(*KeyUsages) {
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
func (m *Windows81SCEPCertificateProfile) GetManagedDeviceCertificateStates()([]ManagedDeviceCertificateStateable) {
    val, err := m.GetBackingStore().Get("managedDeviceCertificateStates")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]ManagedDeviceCertificateStateable)
    }
    return nil
}
// GetRootCertificate gets the rootCertificate property value. Trusted Root Certificate
func (m *Windows81SCEPCertificateProfile) GetRootCertificate()(Windows81TrustedRootCertificateable) {
    val, err := m.GetBackingStore().Get("rootCertificate")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(Windows81TrustedRootCertificateable)
    }
    return nil
}
// GetScepServerUrls gets the scepServerUrls property value. SCEP Server Url(s).
func (m *Windows81SCEPCertificateProfile) GetScepServerUrls()([]string) {
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
func (m *Windows81SCEPCertificateProfile) GetSubjectAlternativeNameFormatString()(*string) {
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
func (m *Windows81SCEPCertificateProfile) GetSubjectNameFormatString()(*string) {
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
func (m *Windows81SCEPCertificateProfile) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Windows81CertificateProfileBase.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetCertificateStore() != nil {
        cast := (*m.GetCertificateStore()).String()
        err = writer.WriteStringValue("certificateStore", &cast)
        if err != nil {
            return err
        }
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
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
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
// SetCertificateStore sets the certificateStore property value. Target store certificate. Possible values are: user, machine.
func (m *Windows81SCEPCertificateProfile) SetCertificateStore(value *Windows81SCEPCertificateProfile_certificateStore)() {
    err := m.GetBackingStore().Set("certificateStore", value)
    if err != nil {
        panic(err)
    }
}
// SetHashAlgorithm sets the hashAlgorithm property value. Hash Algorithm Options.
func (m *Windows81SCEPCertificateProfile) SetHashAlgorithm(value *HashAlgorithms)() {
    err := m.GetBackingStore().Set("hashAlgorithm", value)
    if err != nil {
        panic(err)
    }
}
// SetKeySize sets the keySize property value. Key Size Options.
func (m *Windows81SCEPCertificateProfile) SetKeySize(value *KeySize)() {
    err := m.GetBackingStore().Set("keySize", value)
    if err != nil {
        panic(err)
    }
}
// SetKeyUsage sets the keyUsage property value. Key Usage Options.
func (m *Windows81SCEPCertificateProfile) SetKeyUsage(value *KeyUsages)() {
    err := m.GetBackingStore().Set("keyUsage", value)
    if err != nil {
        panic(err)
    }
}
// SetManagedDeviceCertificateStates sets the managedDeviceCertificateStates property value. Certificate state for devices. This collection can contain a maximum of 2147483647 elements.
func (m *Windows81SCEPCertificateProfile) SetManagedDeviceCertificateStates(value []ManagedDeviceCertificateStateable)() {
    err := m.GetBackingStore().Set("managedDeviceCertificateStates", value)
    if err != nil {
        panic(err)
    }
}
// SetRootCertificate sets the rootCertificate property value. Trusted Root Certificate
func (m *Windows81SCEPCertificateProfile) SetRootCertificate(value Windows81TrustedRootCertificateable)() {
    err := m.GetBackingStore().Set("rootCertificate", value)
    if err != nil {
        panic(err)
    }
}
// SetScepServerUrls sets the scepServerUrls property value. SCEP Server Url(s).
func (m *Windows81SCEPCertificateProfile) SetScepServerUrls(value []string)() {
    err := m.GetBackingStore().Set("scepServerUrls", value)
    if err != nil {
        panic(err)
    }
}
// SetSubjectAlternativeNameFormatString sets the subjectAlternativeNameFormatString property value. Custom String that defines the AAD Attribute.
func (m *Windows81SCEPCertificateProfile) SetSubjectAlternativeNameFormatString(value *string)() {
    err := m.GetBackingStore().Set("subjectAlternativeNameFormatString", value)
    if err != nil {
        panic(err)
    }
}
// SetSubjectNameFormatString sets the subjectNameFormatString property value. Custom format to use with SubjectNameFormat = Custom. Example: CN={{EmailAddress}},E={{EmailAddress}},OU=Enterprise Users,O=Contoso Corporation,L=Redmond,ST=WA,C=US
func (m *Windows81SCEPCertificateProfile) SetSubjectNameFormatString(value *string)() {
    err := m.GetBackingStore().Set("subjectNameFormatString", value)
    if err != nil {
        panic(err)
    }
}
// Windows81SCEPCertificateProfileable 
type Windows81SCEPCertificateProfileable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    Windows81CertificateProfileBaseable
    GetCertificateStore()(*Windows81SCEPCertificateProfile_certificateStore)
    GetHashAlgorithm()(*HashAlgorithms)
    GetKeySize()(*KeySize)
    GetKeyUsage()(*KeyUsages)
    GetManagedDeviceCertificateStates()([]ManagedDeviceCertificateStateable)
    GetRootCertificate()(Windows81TrustedRootCertificateable)
    GetScepServerUrls()([]string)
    GetSubjectAlternativeNameFormatString()(*string)
    GetSubjectNameFormatString()(*string)
    SetCertificateStore(value *Windows81SCEPCertificateProfile_certificateStore)()
    SetHashAlgorithm(value *HashAlgorithms)()
    SetKeySize(value *KeySize)()
    SetKeyUsage(value *KeyUsages)()
    SetManagedDeviceCertificateStates(value []ManagedDeviceCertificateStateable)()
    SetRootCertificate(value Windows81TrustedRootCertificateable)()
    SetScepServerUrls(value []string)()
    SetSubjectAlternativeNameFormatString(value *string)()
    SetSubjectNameFormatString(value *string)()
}
