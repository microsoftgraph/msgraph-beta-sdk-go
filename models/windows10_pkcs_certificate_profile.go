package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// Windows10PkcsCertificateProfile 
type Windows10PkcsCertificateProfile struct {
    Windows10CertificateProfileBase
}
// NewWindows10PkcsCertificateProfile instantiates a new Windows10PkcsCertificateProfile and sets the default values.
func NewWindows10PkcsCertificateProfile()(*Windows10PkcsCertificateProfile) {
    m := &Windows10PkcsCertificateProfile{
        Windows10CertificateProfileBase: *NewWindows10CertificateProfileBase(),
    }
    odataTypeValue := "#microsoft.graph.windows10PkcsCertificateProfile"
    m.SetOdataType(&odataTypeValue)
    return m
}
// CreateWindows10PkcsCertificateProfileFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateWindows10PkcsCertificateProfileFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewWindows10PkcsCertificateProfile(), nil
}
// GetCertificateStore gets the certificateStore property value. Target store certificate. Possible values are: user, machine.
func (m *Windows10PkcsCertificateProfile) GetCertificateStore()(*CertificateStore) {
    val, err := m.GetBackingStore().Get("certificateStore")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*CertificateStore)
    }
    return nil
}
// GetCertificateTemplateName gets the certificateTemplateName property value. PKCS Certificate Template Name
func (m *Windows10PkcsCertificateProfile) GetCertificateTemplateName()(*string) {
    val, err := m.GetBackingStore().Get("certificateTemplateName")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetCertificationAuthority gets the certificationAuthority property value. PKCS Certification Authority
func (m *Windows10PkcsCertificateProfile) GetCertificationAuthority()(*string) {
    val, err := m.GetBackingStore().Get("certificationAuthority")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetCertificationAuthorityName gets the certificationAuthorityName property value. PKCS Certification Authority Name
func (m *Windows10PkcsCertificateProfile) GetCertificationAuthorityName()(*string) {
    val, err := m.GetBackingStore().Get("certificationAuthorityName")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetCustomSubjectAlternativeNames gets the customSubjectAlternativeNames property value. Custom Subject Alternative Name Settings. This collection can contain a maximum of 500 elements.
func (m *Windows10PkcsCertificateProfile) GetCustomSubjectAlternativeNames()([]CustomSubjectAlternativeNameable) {
    val, err := m.GetBackingStore().Get("customSubjectAlternativeNames")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]CustomSubjectAlternativeNameable)
    }
    return nil
}
// GetExtendedKeyUsages gets the extendedKeyUsages property value. Extended Key Usage (EKU) settings. This collection can contain a maximum of 500 elements.
func (m *Windows10PkcsCertificateProfile) GetExtendedKeyUsages()([]ExtendedKeyUsageable) {
    val, err := m.GetBackingStore().Get("extendedKeyUsages")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]ExtendedKeyUsageable)
    }
    return nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *Windows10PkcsCertificateProfile) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Windows10CertificateProfileBase.GetFieldDeserializers()
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
    res["certificateTemplateName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCertificateTemplateName(val)
        }
        return nil
    }
    res["certificationAuthority"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCertificationAuthority(val)
        }
        return nil
    }
    res["certificationAuthorityName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCertificationAuthorityName(val)
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
    res["extendedKeyUsages"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateExtendedKeyUsageFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]ExtendedKeyUsageable, len(val))
            for i, v := range val {
                res[i] = v.(ExtendedKeyUsageable)
            }
            m.SetExtendedKeyUsages(res)
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
// GetManagedDeviceCertificateStates gets the managedDeviceCertificateStates property value. Certificate state for devices. This collection can contain a maximum of 2147483647 elements.
func (m *Windows10PkcsCertificateProfile) GetManagedDeviceCertificateStates()([]ManagedDeviceCertificateStateable) {
    val, err := m.GetBackingStore().Get("managedDeviceCertificateStates")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]ManagedDeviceCertificateStateable)
    }
    return nil
}
// GetSubjectAlternativeNameFormatString gets the subjectAlternativeNameFormatString property value. Custom String that defines the AAD Attribute.
func (m *Windows10PkcsCertificateProfile) GetSubjectAlternativeNameFormatString()(*string) {
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
func (m *Windows10PkcsCertificateProfile) GetSubjectNameFormatString()(*string) {
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
func (m *Windows10PkcsCertificateProfile) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Windows10CertificateProfileBase.Serialize(writer)
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
    {
        err = writer.WriteStringValue("certificateTemplateName", m.GetCertificateTemplateName())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("certificationAuthority", m.GetCertificationAuthority())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("certificationAuthorityName", m.GetCertificationAuthorityName())
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
    if m.GetExtendedKeyUsages() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetExtendedKeyUsages()))
        for i, v := range m.GetExtendedKeyUsages() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("extendedKeyUsages", cast)
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
func (m *Windows10PkcsCertificateProfile) SetCertificateStore(value *CertificateStore)() {
    err := m.GetBackingStore().Set("certificateStore", value)
    if err != nil {
        panic(err)
    }
}
// SetCertificateTemplateName sets the certificateTemplateName property value. PKCS Certificate Template Name
func (m *Windows10PkcsCertificateProfile) SetCertificateTemplateName(value *string)() {
    err := m.GetBackingStore().Set("certificateTemplateName", value)
    if err != nil {
        panic(err)
    }
}
// SetCertificationAuthority sets the certificationAuthority property value. PKCS Certification Authority
func (m *Windows10PkcsCertificateProfile) SetCertificationAuthority(value *string)() {
    err := m.GetBackingStore().Set("certificationAuthority", value)
    if err != nil {
        panic(err)
    }
}
// SetCertificationAuthorityName sets the certificationAuthorityName property value. PKCS Certification Authority Name
func (m *Windows10PkcsCertificateProfile) SetCertificationAuthorityName(value *string)() {
    err := m.GetBackingStore().Set("certificationAuthorityName", value)
    if err != nil {
        panic(err)
    }
}
// SetCustomSubjectAlternativeNames sets the customSubjectAlternativeNames property value. Custom Subject Alternative Name Settings. This collection can contain a maximum of 500 elements.
func (m *Windows10PkcsCertificateProfile) SetCustomSubjectAlternativeNames(value []CustomSubjectAlternativeNameable)() {
    err := m.GetBackingStore().Set("customSubjectAlternativeNames", value)
    if err != nil {
        panic(err)
    }
}
// SetExtendedKeyUsages sets the extendedKeyUsages property value. Extended Key Usage (EKU) settings. This collection can contain a maximum of 500 elements.
func (m *Windows10PkcsCertificateProfile) SetExtendedKeyUsages(value []ExtendedKeyUsageable)() {
    err := m.GetBackingStore().Set("extendedKeyUsages", value)
    if err != nil {
        panic(err)
    }
}
// SetManagedDeviceCertificateStates sets the managedDeviceCertificateStates property value. Certificate state for devices. This collection can contain a maximum of 2147483647 elements.
func (m *Windows10PkcsCertificateProfile) SetManagedDeviceCertificateStates(value []ManagedDeviceCertificateStateable)() {
    err := m.GetBackingStore().Set("managedDeviceCertificateStates", value)
    if err != nil {
        panic(err)
    }
}
// SetSubjectAlternativeNameFormatString sets the subjectAlternativeNameFormatString property value. Custom String that defines the AAD Attribute.
func (m *Windows10PkcsCertificateProfile) SetSubjectAlternativeNameFormatString(value *string)() {
    err := m.GetBackingStore().Set("subjectAlternativeNameFormatString", value)
    if err != nil {
        panic(err)
    }
}
// SetSubjectNameFormatString sets the subjectNameFormatString property value. Custom format to use with SubjectNameFormat = Custom. Example: CN={{EmailAddress}},E={{EmailAddress}},OU=Enterprise Users,O=Contoso Corporation,L=Redmond,ST=WA,C=US
func (m *Windows10PkcsCertificateProfile) SetSubjectNameFormatString(value *string)() {
    err := m.GetBackingStore().Set("subjectNameFormatString", value)
    if err != nil {
        panic(err)
    }
}
// Windows10PkcsCertificateProfileable 
type Windows10PkcsCertificateProfileable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    Windows10CertificateProfileBaseable
    GetCertificateStore()(*CertificateStore)
    GetCertificateTemplateName()(*string)
    GetCertificationAuthority()(*string)
    GetCertificationAuthorityName()(*string)
    GetCustomSubjectAlternativeNames()([]CustomSubjectAlternativeNameable)
    GetExtendedKeyUsages()([]ExtendedKeyUsageable)
    GetManagedDeviceCertificateStates()([]ManagedDeviceCertificateStateable)
    GetSubjectAlternativeNameFormatString()(*string)
    GetSubjectNameFormatString()(*string)
    SetCertificateStore(value *CertificateStore)()
    SetCertificateTemplateName(value *string)()
    SetCertificationAuthority(value *string)()
    SetCertificationAuthorityName(value *string)()
    SetCustomSubjectAlternativeNames(value []CustomSubjectAlternativeNameable)()
    SetExtendedKeyUsages(value []ExtendedKeyUsageable)()
    SetManagedDeviceCertificateStates(value []ManagedDeviceCertificateStateable)()
    SetSubjectAlternativeNameFormatString(value *string)()
    SetSubjectNameFormatString(value *string)()
}
