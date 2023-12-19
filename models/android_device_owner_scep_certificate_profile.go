package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// AndroidDeviceOwnerScepCertificateProfile android Device Owner SCEP certificate profile
type AndroidDeviceOwnerScepCertificateProfile struct {
    AndroidDeviceOwnerCertificateProfileBase
}
// NewAndroidDeviceOwnerScepCertificateProfile instantiates a new androidDeviceOwnerScepCertificateProfile and sets the default values.
func NewAndroidDeviceOwnerScepCertificateProfile()(*AndroidDeviceOwnerScepCertificateProfile) {
    m := &AndroidDeviceOwnerScepCertificateProfile{
        AndroidDeviceOwnerCertificateProfileBase: *NewAndroidDeviceOwnerCertificateProfileBase(),
    }
    odataTypeValue := "#microsoft.graph.androidDeviceOwnerScepCertificateProfile"
    m.SetOdataType(&odataTypeValue)
    return m
}
// CreateAndroidDeviceOwnerScepCertificateProfileFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateAndroidDeviceOwnerScepCertificateProfileFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewAndroidDeviceOwnerScepCertificateProfile(), nil
}
// GetCertificateAccessType gets the certificateAccessType property value. Certificate access type. Possible values are: userApproval, specificApps, unknownFutureValue.
func (m *AndroidDeviceOwnerScepCertificateProfile) GetCertificateAccessType()(*AndroidDeviceOwnerScepCertificateProfile_certificateAccessType) {
    val, err := m.GetBackingStore().Get("certificateAccessType")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*AndroidDeviceOwnerScepCertificateProfile_certificateAccessType)
    }
    return nil
}
// GetCertificateStore gets the certificateStore property value. Target store certificate. Possible values are: user, machine.
func (m *AndroidDeviceOwnerScepCertificateProfile) GetCertificateStore()(*AndroidDeviceOwnerScepCertificateProfile_certificateStore) {
    val, err := m.GetBackingStore().Get("certificateStore")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*AndroidDeviceOwnerScepCertificateProfile_certificateStore)
    }
    return nil
}
// GetCustomSubjectAlternativeNames gets the customSubjectAlternativeNames property value. Custom Subject Alternative Name Settings. This collection can contain a maximum of 500 elements.
func (m *AndroidDeviceOwnerScepCertificateProfile) GetCustomSubjectAlternativeNames()([]CustomSubjectAlternativeNameable) {
    val, err := m.GetBackingStore().Get("customSubjectAlternativeNames")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]CustomSubjectAlternativeNameable)
    }
    return nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *AndroidDeviceOwnerScepCertificateProfile) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.AndroidDeviceOwnerCertificateProfileBase.GetFieldDeserializers()
    res["certificateAccessType"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseAndroidDeviceOwnerScepCertificateProfile_certificateAccessType)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCertificateAccessType(val.(*AndroidDeviceOwnerScepCertificateProfile_certificateAccessType))
        }
        return nil
    }
    res["certificateStore"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseAndroidDeviceOwnerScepCertificateProfile_certificateStore)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCertificateStore(val.(*AndroidDeviceOwnerScepCertificateProfile_certificateStore))
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
                if v != nil {
                    res[i] = v.(CustomSubjectAlternativeNameable)
                }
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
                if v != nil {
                    res[i] = v.(ManagedDeviceCertificateStateable)
                }
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
                if v != nil {
                    res[i] = *(v.(*string))
                }
            }
            m.SetScepServerUrls(res)
        }
        return nil
    }
    res["silentCertificateAccessDetails"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateAndroidDeviceOwnerSilentCertificateAccessFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]AndroidDeviceOwnerSilentCertificateAccessable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(AndroidDeviceOwnerSilentCertificateAccessable)
                }
            }
            m.SetSilentCertificateAccessDetails(res)
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
func (m *AndroidDeviceOwnerScepCertificateProfile) GetHashAlgorithm()(*HashAlgorithms) {
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
func (m *AndroidDeviceOwnerScepCertificateProfile) GetKeySize()(*KeySize) {
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
func (m *AndroidDeviceOwnerScepCertificateProfile) GetKeyUsage()(*KeyUsages) {
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
func (m *AndroidDeviceOwnerScepCertificateProfile) GetManagedDeviceCertificateStates()([]ManagedDeviceCertificateStateable) {
    val, err := m.GetBackingStore().Get("managedDeviceCertificateStates")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]ManagedDeviceCertificateStateable)
    }
    return nil
}
// GetScepServerUrls gets the scepServerUrls property value. SCEP Server Url(s)
func (m *AndroidDeviceOwnerScepCertificateProfile) GetScepServerUrls()([]string) {
    val, err := m.GetBackingStore().Get("scepServerUrls")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]string)
    }
    return nil
}
// GetSilentCertificateAccessDetails gets the silentCertificateAccessDetails property value. Certificate access information. This collection can contain a maximum of 50 elements.
func (m *AndroidDeviceOwnerScepCertificateProfile) GetSilentCertificateAccessDetails()([]AndroidDeviceOwnerSilentCertificateAccessable) {
    val, err := m.GetBackingStore().Get("silentCertificateAccessDetails")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]AndroidDeviceOwnerSilentCertificateAccessable)
    }
    return nil
}
// GetSubjectAlternativeNameFormatString gets the subjectAlternativeNameFormatString property value. Custom String that defines the AAD Attribute.
func (m *AndroidDeviceOwnerScepCertificateProfile) GetSubjectAlternativeNameFormatString()(*string) {
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
func (m *AndroidDeviceOwnerScepCertificateProfile) GetSubjectNameFormatString()(*string) {
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
func (m *AndroidDeviceOwnerScepCertificateProfile) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.AndroidDeviceOwnerCertificateProfileBase.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetCertificateAccessType() != nil {
        cast := (*m.GetCertificateAccessType()).String()
        err = writer.WriteStringValue("certificateAccessType", &cast)
        if err != nil {
            return err
        }
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
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
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
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
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
    if m.GetSilentCertificateAccessDetails() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetSilentCertificateAccessDetails()))
        for i, v := range m.GetSilentCertificateAccessDetails() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err = writer.WriteCollectionOfObjectValues("silentCertificateAccessDetails", cast)
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
// SetCertificateAccessType sets the certificateAccessType property value. Certificate access type. Possible values are: userApproval, specificApps, unknownFutureValue.
func (m *AndroidDeviceOwnerScepCertificateProfile) SetCertificateAccessType(value *AndroidDeviceOwnerScepCertificateProfile_certificateAccessType)() {
    err := m.GetBackingStore().Set("certificateAccessType", value)
    if err != nil {
        panic(err)
    }
}
// SetCertificateStore sets the certificateStore property value. Target store certificate. Possible values are: user, machine.
func (m *AndroidDeviceOwnerScepCertificateProfile) SetCertificateStore(value *AndroidDeviceOwnerScepCertificateProfile_certificateStore)() {
    err := m.GetBackingStore().Set("certificateStore", value)
    if err != nil {
        panic(err)
    }
}
// SetCustomSubjectAlternativeNames sets the customSubjectAlternativeNames property value. Custom Subject Alternative Name Settings. This collection can contain a maximum of 500 elements.
func (m *AndroidDeviceOwnerScepCertificateProfile) SetCustomSubjectAlternativeNames(value []CustomSubjectAlternativeNameable)() {
    err := m.GetBackingStore().Set("customSubjectAlternativeNames", value)
    if err != nil {
        panic(err)
    }
}
// SetHashAlgorithm sets the hashAlgorithm property value. Hash Algorithm Options.
func (m *AndroidDeviceOwnerScepCertificateProfile) SetHashAlgorithm(value *HashAlgorithms)() {
    err := m.GetBackingStore().Set("hashAlgorithm", value)
    if err != nil {
        panic(err)
    }
}
// SetKeySize sets the keySize property value. Key Size Options.
func (m *AndroidDeviceOwnerScepCertificateProfile) SetKeySize(value *KeySize)() {
    err := m.GetBackingStore().Set("keySize", value)
    if err != nil {
        panic(err)
    }
}
// SetKeyUsage sets the keyUsage property value. Key Usage Options.
func (m *AndroidDeviceOwnerScepCertificateProfile) SetKeyUsage(value *KeyUsages)() {
    err := m.GetBackingStore().Set("keyUsage", value)
    if err != nil {
        panic(err)
    }
}
// SetManagedDeviceCertificateStates sets the managedDeviceCertificateStates property value. Certificate state for devices. This collection can contain a maximum of 2147483647 elements.
func (m *AndroidDeviceOwnerScepCertificateProfile) SetManagedDeviceCertificateStates(value []ManagedDeviceCertificateStateable)() {
    err := m.GetBackingStore().Set("managedDeviceCertificateStates", value)
    if err != nil {
        panic(err)
    }
}
// SetScepServerUrls sets the scepServerUrls property value. SCEP Server Url(s)
func (m *AndroidDeviceOwnerScepCertificateProfile) SetScepServerUrls(value []string)() {
    err := m.GetBackingStore().Set("scepServerUrls", value)
    if err != nil {
        panic(err)
    }
}
// SetSilentCertificateAccessDetails sets the silentCertificateAccessDetails property value. Certificate access information. This collection can contain a maximum of 50 elements.
func (m *AndroidDeviceOwnerScepCertificateProfile) SetSilentCertificateAccessDetails(value []AndroidDeviceOwnerSilentCertificateAccessable)() {
    err := m.GetBackingStore().Set("silentCertificateAccessDetails", value)
    if err != nil {
        panic(err)
    }
}
// SetSubjectAlternativeNameFormatString sets the subjectAlternativeNameFormatString property value. Custom String that defines the AAD Attribute.
func (m *AndroidDeviceOwnerScepCertificateProfile) SetSubjectAlternativeNameFormatString(value *string)() {
    err := m.GetBackingStore().Set("subjectAlternativeNameFormatString", value)
    if err != nil {
        panic(err)
    }
}
// SetSubjectNameFormatString sets the subjectNameFormatString property value. Custom format to use with SubjectNameFormat = Custom. Example: CN={{EmailAddress}},E={{EmailAddress}},OU=Enterprise Users,O=Contoso Corporation,L=Redmond,ST=WA,C=US
func (m *AndroidDeviceOwnerScepCertificateProfile) SetSubjectNameFormatString(value *string)() {
    err := m.GetBackingStore().Set("subjectNameFormatString", value)
    if err != nil {
        panic(err)
    }
}
// AndroidDeviceOwnerScepCertificateProfileable 
type AndroidDeviceOwnerScepCertificateProfileable interface {
    AndroidDeviceOwnerCertificateProfileBaseable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetCertificateAccessType()(*AndroidDeviceOwnerScepCertificateProfile_certificateAccessType)
    GetCertificateStore()(*AndroidDeviceOwnerScepCertificateProfile_certificateStore)
    GetCustomSubjectAlternativeNames()([]CustomSubjectAlternativeNameable)
    GetHashAlgorithm()(*HashAlgorithms)
    GetKeySize()(*KeySize)
    GetKeyUsage()(*KeyUsages)
    GetManagedDeviceCertificateStates()([]ManagedDeviceCertificateStateable)
    GetScepServerUrls()([]string)
    GetSilentCertificateAccessDetails()([]AndroidDeviceOwnerSilentCertificateAccessable)
    GetSubjectAlternativeNameFormatString()(*string)
    GetSubjectNameFormatString()(*string)
    SetCertificateAccessType(value *AndroidDeviceOwnerScepCertificateProfile_certificateAccessType)()
    SetCertificateStore(value *AndroidDeviceOwnerScepCertificateProfile_certificateStore)()
    SetCustomSubjectAlternativeNames(value []CustomSubjectAlternativeNameable)()
    SetHashAlgorithm(value *HashAlgorithms)()
    SetKeySize(value *KeySize)()
    SetKeyUsage(value *KeyUsages)()
    SetManagedDeviceCertificateStates(value []ManagedDeviceCertificateStateable)()
    SetScepServerUrls(value []string)()
    SetSilentCertificateAccessDetails(value []AndroidDeviceOwnerSilentCertificateAccessable)()
    SetSubjectAlternativeNameFormatString(value *string)()
    SetSubjectNameFormatString(value *string)()
}
