package models

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// AndroidDeviceOwnerImportedPFXCertificateProfile 
type AndroidDeviceOwnerImportedPFXCertificateProfile struct {
    AndroidDeviceOwnerCertificateProfileBase
    // Certificate access type. Possible values are: userApproval, specificApps, unknownFutureValue.
    certificateAccessType *AndroidDeviceOwnerCertificateAccessType
    // PFX Import Options.
    intendedPurpose *IntendedPurpose
    // Certificate state for devices. This collection can contain a maximum of 2147483647 elements.
    managedDeviceCertificateStates []ManagedDeviceCertificateStateable
    // Certificate access information. This collection can contain a maximum of 50 elements.
    silentCertificateAccessDetails []AndroidDeviceOwnerSilentCertificateAccessable
}
// NewAndroidDeviceOwnerImportedPFXCertificateProfile instantiates a new AndroidDeviceOwnerImportedPFXCertificateProfile and sets the default values.
func NewAndroidDeviceOwnerImportedPFXCertificateProfile()(*AndroidDeviceOwnerImportedPFXCertificateProfile) {
    m := &AndroidDeviceOwnerImportedPFXCertificateProfile{
        AndroidDeviceOwnerCertificateProfileBase: *NewAndroidDeviceOwnerCertificateProfileBase(),
    }
    odataTypeValue := "#microsoft.graph.androidDeviceOwnerImportedPFXCertificateProfile";
    m.SetOdataType(&odataTypeValue);
    return m
}
// CreateAndroidDeviceOwnerImportedPFXCertificateProfileFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateAndroidDeviceOwnerImportedPFXCertificateProfileFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewAndroidDeviceOwnerImportedPFXCertificateProfile(), nil
}
// GetCertificateAccessType gets the certificateAccessType property value. Certificate access type. Possible values are: userApproval, specificApps, unknownFutureValue.
func (m *AndroidDeviceOwnerImportedPFXCertificateProfile) GetCertificateAccessType()(*AndroidDeviceOwnerCertificateAccessType) {
    return m.certificateAccessType
}
// GetFieldDeserializers the deserialization information for the current model
func (m *AndroidDeviceOwnerImportedPFXCertificateProfile) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.AndroidDeviceOwnerCertificateProfileBase.GetFieldDeserializers()
    res["certificateAccessType"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetEnumValue(ParseAndroidDeviceOwnerCertificateAccessType , m.SetCertificateAccessType)
    res["intendedPurpose"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetEnumValue(ParseIntendedPurpose , m.SetIntendedPurpose)
    res["managedDeviceCertificateStates"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfObjectValues(CreateManagedDeviceCertificateStateFromDiscriminatorValue , m.SetManagedDeviceCertificateStates)
    res["silentCertificateAccessDetails"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfObjectValues(CreateAndroidDeviceOwnerSilentCertificateAccessFromDiscriminatorValue , m.SetSilentCertificateAccessDetails)
    return res
}
// GetIntendedPurpose gets the intendedPurpose property value. PFX Import Options.
func (m *AndroidDeviceOwnerImportedPFXCertificateProfile) GetIntendedPurpose()(*IntendedPurpose) {
    return m.intendedPurpose
}
// GetManagedDeviceCertificateStates gets the managedDeviceCertificateStates property value. Certificate state for devices. This collection can contain a maximum of 2147483647 elements.
func (m *AndroidDeviceOwnerImportedPFXCertificateProfile) GetManagedDeviceCertificateStates()([]ManagedDeviceCertificateStateable) {
    return m.managedDeviceCertificateStates
}
// GetSilentCertificateAccessDetails gets the silentCertificateAccessDetails property value. Certificate access information. This collection can contain a maximum of 50 elements.
func (m *AndroidDeviceOwnerImportedPFXCertificateProfile) GetSilentCertificateAccessDetails()([]AndroidDeviceOwnerSilentCertificateAccessable) {
    return m.silentCertificateAccessDetails
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
        cast := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.CollectionCast[i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable](m.GetManagedDeviceCertificateStates())
        err = writer.WriteCollectionOfObjectValues("managedDeviceCertificateStates", cast)
        if err != nil {
            return err
        }
    }
    if m.GetSilentCertificateAccessDetails() != nil {
        cast := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.CollectionCast[i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable](m.GetSilentCertificateAccessDetails())
        err = writer.WriteCollectionOfObjectValues("silentCertificateAccessDetails", cast)
        if err != nil {
            return err
        }
    }
    return nil
}
// SetCertificateAccessType sets the certificateAccessType property value. Certificate access type. Possible values are: userApproval, specificApps, unknownFutureValue.
func (m *AndroidDeviceOwnerImportedPFXCertificateProfile) SetCertificateAccessType(value *AndroidDeviceOwnerCertificateAccessType)() {
    m.certificateAccessType = value
}
// SetIntendedPurpose sets the intendedPurpose property value. PFX Import Options.
func (m *AndroidDeviceOwnerImportedPFXCertificateProfile) SetIntendedPurpose(value *IntendedPurpose)() {
    m.intendedPurpose = value
}
// SetManagedDeviceCertificateStates sets the managedDeviceCertificateStates property value. Certificate state for devices. This collection can contain a maximum of 2147483647 elements.
func (m *AndroidDeviceOwnerImportedPFXCertificateProfile) SetManagedDeviceCertificateStates(value []ManagedDeviceCertificateStateable)() {
    m.managedDeviceCertificateStates = value
}
// SetSilentCertificateAccessDetails sets the silentCertificateAccessDetails property value. Certificate access information. This collection can contain a maximum of 50 elements.
func (m *AndroidDeviceOwnerImportedPFXCertificateProfile) SetSilentCertificateAccessDetails(value []AndroidDeviceOwnerSilentCertificateAccessable)() {
    m.silentCertificateAccessDetails = value
}
