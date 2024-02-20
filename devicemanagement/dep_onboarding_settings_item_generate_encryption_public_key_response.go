package devicemanagement

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// Deprecated: This class is obsolete. Use DepOnboardingSettingsItemGenerateEncryptionPublicKeyPostResponseable instead.
type DepOnboardingSettingsItemGenerateEncryptionPublicKeyResponse struct {
    DepOnboardingSettingsItemGenerateEncryptionPublicKeyPostResponse
}
// NewDepOnboardingSettingsItemGenerateEncryptionPublicKeyResponse instantiates a new DepOnboardingSettingsItemGenerateEncryptionPublicKeyResponse and sets the default values.
func NewDepOnboardingSettingsItemGenerateEncryptionPublicKeyResponse()(*DepOnboardingSettingsItemGenerateEncryptionPublicKeyResponse) {
    m := &DepOnboardingSettingsItemGenerateEncryptionPublicKeyResponse{
        DepOnboardingSettingsItemGenerateEncryptionPublicKeyPostResponse: *NewDepOnboardingSettingsItemGenerateEncryptionPublicKeyPostResponse(),
    }
    return m
}
// CreateDepOnboardingSettingsItemGenerateEncryptionPublicKeyResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateDepOnboardingSettingsItemGenerateEncryptionPublicKeyResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewDepOnboardingSettingsItemGenerateEncryptionPublicKeyResponse(), nil
}
// Deprecated: This class is obsolete. Use DepOnboardingSettingsItemGenerateEncryptionPublicKeyPostResponseable instead.
type DepOnboardingSettingsItemGenerateEncryptionPublicKeyResponseable interface {
    DepOnboardingSettingsItemGenerateEncryptionPublicKeyPostResponseable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
}
