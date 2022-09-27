package models

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// IosImportedPFXCertificateProfile 
type IosImportedPFXCertificateProfile struct {
    IosCertificateProfile
    // PFX Import Options.
    intendedPurpose *IntendedPurpose
    // Certificate state for devices. This collection can contain a maximum of 2147483647 elements.
    managedDeviceCertificateStates []ManagedDeviceCertificateStateable
}
// NewIosImportedPFXCertificateProfile instantiates a new IosImportedPFXCertificateProfile and sets the default values.
func NewIosImportedPFXCertificateProfile()(*IosImportedPFXCertificateProfile) {
    m := &IosImportedPFXCertificateProfile{
        IosCertificateProfile: *NewIosCertificateProfile(),
    }
    odataTypeValue := "#microsoft.graph.iosImportedPFXCertificateProfile";
    m.SetOdataType(&odataTypeValue);
    return m
}
// CreateIosImportedPFXCertificateProfileFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateIosImportedPFXCertificateProfileFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewIosImportedPFXCertificateProfile(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *IosImportedPFXCertificateProfile) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.IosCertificateProfile.GetFieldDeserializers()
    res["intendedPurpose"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetEnumValue(ParseIntendedPurpose , m.SetIntendedPurpose)
    res["managedDeviceCertificateStates"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfObjectValues(CreateManagedDeviceCertificateStateFromDiscriminatorValue , m.SetManagedDeviceCertificateStates)
    return res
}
// GetIntendedPurpose gets the intendedPurpose property value. PFX Import Options.
func (m *IosImportedPFXCertificateProfile) GetIntendedPurpose()(*IntendedPurpose) {
    return m.intendedPurpose
}
// GetManagedDeviceCertificateStates gets the managedDeviceCertificateStates property value. Certificate state for devices. This collection can contain a maximum of 2147483647 elements.
func (m *IosImportedPFXCertificateProfile) GetManagedDeviceCertificateStates()([]ManagedDeviceCertificateStateable) {
    return m.managedDeviceCertificateStates
}
// Serialize serializes information the current object
func (m *IosImportedPFXCertificateProfile) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.IosCertificateProfile.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetIntendedPurpose() != nil {
        cast := (*m.GetIntendedPurpose()).String()
        err = writer.WriteStringValue("intendedPurpose", &cast)
        if err != nil {
            return err
        }
    }
    if m.GetManagedDeviceCertificateStates() != nil {
        cast := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.CollectionCast[i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable](m.GetManagedDeviceCertificateStates())
        err = writer.WriteCollectionOfObjectValues("managedDeviceCertificateStates", cast)
        if err != nil {
            return err
        }
    }
    return nil
}
// SetIntendedPurpose sets the intendedPurpose property value. PFX Import Options.
func (m *IosImportedPFXCertificateProfile) SetIntendedPurpose(value *IntendedPurpose)() {
    m.intendedPurpose = value
}
// SetManagedDeviceCertificateStates sets the managedDeviceCertificateStates property value. Certificate state for devices. This collection can contain a maximum of 2147483647 elements.
func (m *IosImportedPFXCertificateProfile) SetManagedDeviceCertificateStates(value []ManagedDeviceCertificateStateable)() {
    m.managedDeviceCertificateStates = value
}
