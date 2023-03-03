package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// Windows81CertificateProfileBase 
type Windows81CertificateProfileBase struct {
    WindowsCertificateProfileBase
}
// NewWindows81CertificateProfileBase instantiates a new Windows81CertificateProfileBase and sets the default values.
func NewWindows81CertificateProfileBase()(*Windows81CertificateProfileBase) {
    m := &Windows81CertificateProfileBase{
        WindowsCertificateProfileBase: *NewWindowsCertificateProfileBase(),
    }
    odataTypeValue := "#microsoft.graph.windows81CertificateProfileBase"
    m.SetOdataType(&odataTypeValue)
    return m
}
// CreateWindows81CertificateProfileBaseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateWindows81CertificateProfileBaseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
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
                    case "#microsoft.graph.windows81SCEPCertificateProfile":
                        return NewWindows81SCEPCertificateProfile(), nil
                }
            }
        }
    }
    return NewWindows81CertificateProfileBase(), nil
}
// GetCustomSubjectAlternativeNames gets the customSubjectAlternativeNames property value. Custom Subject Alternative Name Settings. This collection can contain a maximum of 500 elements.
func (m *Windows81CertificateProfileBase) GetCustomSubjectAlternativeNames()([]CustomSubjectAlternativeNameable) {
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
func (m *Windows81CertificateProfileBase) GetExtendedKeyUsages()([]ExtendedKeyUsageable) {
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
func (m *Windows81CertificateProfileBase) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.WindowsCertificateProfileBase.GetFieldDeserializers()
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
    return res
}
// Serialize serializes information the current object
func (m *Windows81CertificateProfileBase) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.WindowsCertificateProfileBase.Serialize(writer)
    if err != nil {
        return err
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
    return nil
}
// SetCustomSubjectAlternativeNames sets the customSubjectAlternativeNames property value. Custom Subject Alternative Name Settings. This collection can contain a maximum of 500 elements.
func (m *Windows81CertificateProfileBase) SetCustomSubjectAlternativeNames(value []CustomSubjectAlternativeNameable)() {
    err := m.GetBackingStore().Set("customSubjectAlternativeNames", value)
    if err != nil {
        panic(err)
    }
}
// SetExtendedKeyUsages sets the extendedKeyUsages property value. Extended Key Usage (EKU) settings. This collection can contain a maximum of 500 elements.
func (m *Windows81CertificateProfileBase) SetExtendedKeyUsages(value []ExtendedKeyUsageable)() {
    err := m.GetBackingStore().Set("extendedKeyUsages", value)
    if err != nil {
        panic(err)
    }
}
// Windows81CertificateProfileBaseable 
type Windows81CertificateProfileBaseable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    WindowsCertificateProfileBaseable
    GetCustomSubjectAlternativeNames()([]CustomSubjectAlternativeNameable)
    GetExtendedKeyUsages()([]ExtendedKeyUsageable)
    SetCustomSubjectAlternativeNames(value []CustomSubjectAlternativeNameable)()
    SetExtendedKeyUsages(value []ExtendedKeyUsageable)()
}
