package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// AndroidDeviceOwnerImportedPFXCertificateProfile 
type AndroidDeviceOwnerImportedPFXCertificateProfile struct {
    AndroidDeviceOwnerCertificateProfileBase
}
// NewAndroidDeviceOwnerImportedPFXCertificateProfile instantiates a new AndroidDeviceOwnerImportedPFXCertificateProfile and sets the default values.
func NewAndroidDeviceOwnerImportedPFXCertificateProfile()(*AndroidDeviceOwnerImportedPFXCertificateProfile) {
    m := &AndroidDeviceOwnerImportedPFXCertificateProfile{
        AndroidDeviceOwnerCertificateProfileBase: *NewAndroidDeviceOwnerCertificateProfileBase(),
    }
    odataTypeValue := "#microsoft.graph.androidDeviceOwnerImportedPFXCertificateProfile"
    m.SetOdataType(&odataTypeValue)
    return m
}
// CreateAndroidDeviceOwnerImportedPFXCertificateProfileFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateAndroidDeviceOwnerImportedPFXCertificateProfileFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewAndroidDeviceOwnerImportedPFXCertificateProfile(), nil
}
// GetCertificateAccessType gets the certificateAccessType property value. Certificate access type. Possible values are: userApproval, specificApps, unknownFutureValue.
func (m *AndroidDeviceOwnerImportedPFXCertificateProfile) GetCertificateAccessType()(*AndroidDeviceOwnerCertificateAccessType) {
    val, err := m.GetBackingStore().Get("certificateAccessType")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*AndroidDeviceOwnerCertificateAccessType)
    }
    return nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *AndroidDeviceOwnerImportedPFXCertificateProfile) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.AndroidDeviceOwnerCertificateProfileBase.GetFieldDeserializers()
    res["certificateAccessType"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseAndroidDeviceOwnerCertificateAccessType)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCertificateAccessType(val.(*AndroidDeviceOwnerCertificateAccessType))
        }
        return nil
    }
    res["intendedPurpose"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseIntendedPurpose)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIntendedPurpose(val.(*IntendedPurpose))
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
    res["silentCertificateAccessDetails"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateAndroidDeviceOwnerSilentCertificateAccessFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]AndroidDeviceOwnerSilentCertificateAccessable, len(val))
            for i, v := range val {
                res[i] = v.(AndroidDeviceOwnerSilentCertificateAccessable)
            }
            m.SetSilentCertificateAccessDetails(res)
        }
        return nil
    }
    return res
}
// GetIntendedPurpose gets the intendedPurpose property value. PFX Import Options.
func (m *AndroidDeviceOwnerImportedPFXCertificateProfile) GetIntendedPurpose()(*IntendedPurpose) {
    val, err := m.GetBackingStore().Get("intendedPurpose")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*IntendedPurpose)
    }
    return nil
}
// GetManagedDeviceCertificateStates gets the managedDeviceCertificateStates property value. Certificate state for devices. This collection can contain a maximum of 2147483647 elements.
func (m *AndroidDeviceOwnerImportedPFXCertificateProfile) GetManagedDeviceCertificateStates()([]ManagedDeviceCertificateStateable) {
    val, err := m.GetBackingStore().Get("managedDeviceCertificateStates")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]ManagedDeviceCertificateStateable)
    }
    return nil
}
// GetSilentCertificateAccessDetails gets the silentCertificateAccessDetails property value. Certificate access information. This collection can contain a maximum of 50 elements.
func (m *AndroidDeviceOwnerImportedPFXCertificateProfile) GetSilentCertificateAccessDetails()([]AndroidDeviceOwnerSilentCertificateAccessable) {
    val, err := m.GetBackingStore().Get("silentCertificateAccessDetails")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]AndroidDeviceOwnerSilentCertificateAccessable)
    }
    return nil
}
// Serialize serializes information the current object
func (m *AndroidDeviceOwnerImportedPFXCertificateProfile) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
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
    if m.GetIntendedPurpose() != nil {
        cast := (*m.GetIntendedPurpose()).String()
        err = writer.WriteStringValue("intendedPurpose", &cast)
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
    if m.GetSilentCertificateAccessDetails() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetSilentCertificateAccessDetails()))
        for i, v := range m.GetSilentCertificateAccessDetails() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("silentCertificateAccessDetails", cast)
        if err != nil {
            return err
        }
    }
    return nil
}
// SetCertificateAccessType sets the certificateAccessType property value. Certificate access type. Possible values are: userApproval, specificApps, unknownFutureValue.
func (m *AndroidDeviceOwnerImportedPFXCertificateProfile) SetCertificateAccessType(value *AndroidDeviceOwnerCertificateAccessType)() {
    err := m.GetBackingStore().Set("certificateAccessType", value)
    if err != nil {
        panic(err)
    }
}
// SetIntendedPurpose sets the intendedPurpose property value. PFX Import Options.
func (m *AndroidDeviceOwnerImportedPFXCertificateProfile) SetIntendedPurpose(value *IntendedPurpose)() {
    err := m.GetBackingStore().Set("intendedPurpose", value)
    if err != nil {
        panic(err)
    }
}
// SetManagedDeviceCertificateStates sets the managedDeviceCertificateStates property value. Certificate state for devices. This collection can contain a maximum of 2147483647 elements.
func (m *AndroidDeviceOwnerImportedPFXCertificateProfile) SetManagedDeviceCertificateStates(value []ManagedDeviceCertificateStateable)() {
    err := m.GetBackingStore().Set("managedDeviceCertificateStates", value)
    if err != nil {
        panic(err)
    }
}
// SetSilentCertificateAccessDetails sets the silentCertificateAccessDetails property value. Certificate access information. This collection can contain a maximum of 50 elements.
func (m *AndroidDeviceOwnerImportedPFXCertificateProfile) SetSilentCertificateAccessDetails(value []AndroidDeviceOwnerSilentCertificateAccessable)() {
    err := m.GetBackingStore().Set("silentCertificateAccessDetails", value)
    if err != nil {
        panic(err)
    }
}
// AndroidDeviceOwnerImportedPFXCertificateProfileable 
type AndroidDeviceOwnerImportedPFXCertificateProfileable interface {
    AndroidDeviceOwnerCertificateProfileBaseable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetCertificateAccessType()(*AndroidDeviceOwnerCertificateAccessType)
    GetIntendedPurpose()(*IntendedPurpose)
    GetManagedDeviceCertificateStates()([]ManagedDeviceCertificateStateable)
    GetSilentCertificateAccessDetails()([]AndroidDeviceOwnerSilentCertificateAccessable)
    SetCertificateAccessType(value *AndroidDeviceOwnerCertificateAccessType)()
    SetIntendedPurpose(value *IntendedPurpose)()
    SetManagedDeviceCertificateStates(value []ManagedDeviceCertificateStateable)()
    SetSilentCertificateAccessDetails(value []AndroidDeviceOwnerSilentCertificateAccessable)()
}
