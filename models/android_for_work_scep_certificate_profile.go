package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// AndroidForWorkScepCertificateProfile 
type AndroidForWorkScepCertificateProfile struct {
    AndroidForWorkCertificateProfileBase
    // Target store certificate. Possible values are: user, machine.
    certificateStore *CertificateStore
    // Custom Subject Alternative Name Settings. This collection can contain a maximum of 500 elements.
    customSubjectAlternativeNames []CustomSubjectAlternativeNameable
    // SCEP Hash Algorithm. Possible values are: sha1, sha2.
    hashAlgorithm *HashAlgorithms
    // SCEP Key Size. Possible values are: size1024, size2048, size4096.
    keySize *KeySize
    // SCEP Key Usage. Possible values are: keyEncipherment, digitalSignature.
    keyUsage *KeyUsages
    // Certificate state for devices. This collection can contain a maximum of 2147483647 elements.
    managedDeviceCertificateStates []ManagedDeviceCertificateStateable
    // SCEP Server Url(s)
    scepServerUrls []string
    // Custom String that defines the AAD Attribute.
    subjectAlternativeNameFormatString *string
    // Custom format to use with SubjectNameFormat = Custom. Example: CN={{EmailAddress}},E={{EmailAddress}},OU=Enterprise Users,O=Contoso Corporation,L=Redmond,ST=WA,C=US
    subjectNameFormatString *string
}
// NewAndroidForWorkScepCertificateProfile instantiates a new AndroidForWorkScepCertificateProfile and sets the default values.
func NewAndroidForWorkScepCertificateProfile()(*AndroidForWorkScepCertificateProfile) {
    m := &AndroidForWorkScepCertificateProfile{
        AndroidForWorkCertificateProfileBase: *NewAndroidForWorkCertificateProfileBase(),
    }
    return m
}
// CreateAndroidForWorkScepCertificateProfileFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateAndroidForWorkScepCertificateProfileFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewAndroidForWorkScepCertificateProfile(), nil
}
// GetCertificateStore gets the certificateStore property value. Target store certificate. Possible values are: user, machine.
func (m *AndroidForWorkScepCertificateProfile) GetCertificateStore()(*CertificateStore) {
    if m == nil {
        return nil
    } else {
        return m.certificateStore
    }
}
// GetCustomSubjectAlternativeNames gets the customSubjectAlternativeNames property value. Custom Subject Alternative Name Settings. This collection can contain a maximum of 500 elements.
func (m *AndroidForWorkScepCertificateProfile) GetCustomSubjectAlternativeNames()([]CustomSubjectAlternativeNameable) {
    if m == nil {
        return nil
    } else {
        return m.customSubjectAlternativeNames
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *AndroidForWorkScepCertificateProfile) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.AndroidForWorkCertificateProfileBase.GetFieldDeserializers()
    res["certificateStore"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseCertificateStore)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCertificateStore(val.(*CertificateStore))
        }
        return nil
    }
    res["customSubjectAlternativeNames"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateCustomSubjectAlternativeNameFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]CustomSubjectAlternativeNameable, len(val))
            for i, v := range val {
                res[i] = v.(CustomSubjectAlternativeNameable)
            }
            m.SetCustomSubjectAlternativeNames(res)
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
                res[i] = v.(ManagedDeviceCertificateStateable)
            }
            m.SetManagedDeviceCertificateStates(res)
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
// GetHashAlgorithm gets the hashAlgorithm property value. SCEP Hash Algorithm. Possible values are: sha1, sha2.
func (m *AndroidForWorkScepCertificateProfile) GetHashAlgorithm()(*HashAlgorithms) {
    if m == nil {
        return nil
    } else {
        return m.hashAlgorithm
    }
}
// GetKeySize gets the keySize property value. SCEP Key Size. Possible values are: size1024, size2048, size4096.
func (m *AndroidForWorkScepCertificateProfile) GetKeySize()(*KeySize) {
    if m == nil {
        return nil
    } else {
        return m.keySize
    }
}
// GetKeyUsage gets the keyUsage property value. SCEP Key Usage. Possible values are: keyEncipherment, digitalSignature.
func (m *AndroidForWorkScepCertificateProfile) GetKeyUsage()(*KeyUsages) {
    if m == nil {
        return nil
    } else {
        return m.keyUsage
    }
}
// GetManagedDeviceCertificateStates gets the managedDeviceCertificateStates property value. Certificate state for devices. This collection can contain a maximum of 2147483647 elements.
func (m *AndroidForWorkScepCertificateProfile) GetManagedDeviceCertificateStates()([]ManagedDeviceCertificateStateable) {
    if m == nil {
        return nil
    } else {
        return m.managedDeviceCertificateStates
    }
}
// GetScepServerUrls gets the scepServerUrls property value. SCEP Server Url(s)
func (m *AndroidForWorkScepCertificateProfile) GetScepServerUrls()([]string) {
    if m == nil {
        return nil
    } else {
        return m.scepServerUrls
    }
}
// GetSubjectAlternativeNameFormatString gets the subjectAlternativeNameFormatString property value. Custom String that defines the AAD Attribute.
func (m *AndroidForWorkScepCertificateProfile) GetSubjectAlternativeNameFormatString()(*string) {
    if m == nil {
        return nil
    } else {
        return m.subjectAlternativeNameFormatString
    }
}
// GetSubjectNameFormatString gets the subjectNameFormatString property value. Custom format to use with SubjectNameFormat = Custom. Example: CN={{EmailAddress}},E={{EmailAddress}},OU=Enterprise Users,O=Contoso Corporation,L=Redmond,ST=WA,C=US
func (m *AndroidForWorkScepCertificateProfile) GetSubjectNameFormatString()(*string) {
    if m == nil {
        return nil
    } else {
        return m.subjectNameFormatString
    }
}
// Serialize serializes information the current object
func (m *AndroidForWorkScepCertificateProfile) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.AndroidForWorkCertificateProfileBase.Serialize(writer)
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
    if m.GetCustomSubjectAlternativeNames() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetCustomSubjectAlternativeNames()))
        for i, v := range m.GetCustomSubjectAlternativeNames() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("customSubjectAlternativeNames", cast)
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
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("managedDeviceCertificateStates", cast)
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
func (m *AndroidForWorkScepCertificateProfile) SetCertificateStore(value *CertificateStore)() {
    if m != nil {
        m.certificateStore = value
    }
}
// SetCustomSubjectAlternativeNames sets the customSubjectAlternativeNames property value. Custom Subject Alternative Name Settings. This collection can contain a maximum of 500 elements.
func (m *AndroidForWorkScepCertificateProfile) SetCustomSubjectAlternativeNames(value []CustomSubjectAlternativeNameable)() {
    if m != nil {
        m.customSubjectAlternativeNames = value
    }
}
// SetHashAlgorithm sets the hashAlgorithm property value. SCEP Hash Algorithm. Possible values are: sha1, sha2.
func (m *AndroidForWorkScepCertificateProfile) SetHashAlgorithm(value *HashAlgorithms)() {
    if m != nil {
        m.hashAlgorithm = value
    }
}
// SetKeySize sets the keySize property value. SCEP Key Size. Possible values are: size1024, size2048, size4096.
func (m *AndroidForWorkScepCertificateProfile) SetKeySize(value *KeySize)() {
    if m != nil {
        m.keySize = value
    }
}
// SetKeyUsage sets the keyUsage property value. SCEP Key Usage. Possible values are: keyEncipherment, digitalSignature.
func (m *AndroidForWorkScepCertificateProfile) SetKeyUsage(value *KeyUsages)() {
    if m != nil {
        m.keyUsage = value
    }
}
// SetManagedDeviceCertificateStates sets the managedDeviceCertificateStates property value. Certificate state for devices. This collection can contain a maximum of 2147483647 elements.
func (m *AndroidForWorkScepCertificateProfile) SetManagedDeviceCertificateStates(value []ManagedDeviceCertificateStateable)() {
    if m != nil {
        m.managedDeviceCertificateStates = value
    }
}
// SetScepServerUrls sets the scepServerUrls property value. SCEP Server Url(s)
func (m *AndroidForWorkScepCertificateProfile) SetScepServerUrls(value []string)() {
    if m != nil {
        m.scepServerUrls = value
    }
}
// SetSubjectAlternativeNameFormatString sets the subjectAlternativeNameFormatString property value. Custom String that defines the AAD Attribute.
func (m *AndroidForWorkScepCertificateProfile) SetSubjectAlternativeNameFormatString(value *string)() {
    if m != nil {
        m.subjectAlternativeNameFormatString = value
    }
}
// SetSubjectNameFormatString sets the subjectNameFormatString property value. Custom format to use with SubjectNameFormat = Custom. Example: CN={{EmailAddress}},E={{EmailAddress}},OU=Enterprise Users,O=Contoso Corporation,L=Redmond,ST=WA,C=US
func (m *AndroidForWorkScepCertificateProfile) SetSubjectNameFormatString(value *string)() {
    if m != nil {
        m.subjectNameFormatString = value
    }
}
