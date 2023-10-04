package devicemanagement

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// DepOnboardingSettingsItemGetEncryptionPublicKeyResponse 
// Deprecated: This class is obsolete. Use getEncryptionPublicKeyGetResponse instead.
type DepOnboardingSettingsItemGetEncryptionPublicKeyResponse struct {
    DepOnboardingSettingsItemGetEncryptionPublicKeyGetResponse
}
// NewDepOnboardingSettingsItemGetEncryptionPublicKeyResponse instantiates a new DepOnboardingSettingsItemGetEncryptionPublicKeyResponse and sets the default values.
func NewDepOnboardingSettingsItemGetEncryptionPublicKeyResponse()(*DepOnboardingSettingsItemGetEncryptionPublicKeyResponse) {
    m := &DepOnboardingSettingsItemGetEncryptionPublicKeyResponse{
        DepOnboardingSettingsItemGetEncryptionPublicKeyGetResponse: *NewDepOnboardingSettingsItemGetEncryptionPublicKeyGetResponse(),
    }
    return m
}
// CreateDepOnboardingSettingsItemGetEncryptionPublicKeyResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateDepOnboardingSettingsItemGetEncryptionPublicKeyResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewDepOnboardingSettingsItemGetEncryptionPublicKeyResponse(), nil
}
// DepOnboardingSettingsItemGetEncryptionPublicKeyResponseable 
// Deprecated: This class is obsolete. Use getEncryptionPublicKeyGetResponse instead.
type DepOnboardingSettingsItemGetEncryptionPublicKeyResponseable interface {
    DepOnboardingSettingsItemGetEncryptionPublicKeyGetResponseable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
}
