package devicemanagement

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// Deprecated: This class is obsolete. Use IntentsItemGetcustomizedsettingsGetCustomizedSettingsGetResponseable instead.
type IntentsItemGetcustomizedsettingsGetCustomizedSettingsResponse struct {
    IntentsItemGetcustomizedsettingsGetCustomizedSettingsGetResponse
}
// NewIntentsItemGetcustomizedsettingsGetCustomizedSettingsResponse instantiates a new IntentsItemGetcustomizedsettingsGetCustomizedSettingsResponse and sets the default values.
func NewIntentsItemGetcustomizedsettingsGetCustomizedSettingsResponse()(*IntentsItemGetcustomizedsettingsGetCustomizedSettingsResponse) {
    m := &IntentsItemGetcustomizedsettingsGetCustomizedSettingsResponse{
        IntentsItemGetcustomizedsettingsGetCustomizedSettingsGetResponse: *NewIntentsItemGetcustomizedsettingsGetCustomizedSettingsGetResponse(),
    }
    return m
}
// CreateIntentsItemGetcustomizedsettingsGetCustomizedSettingsResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateIntentsItemGetcustomizedsettingsGetCustomizedSettingsResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewIntentsItemGetcustomizedsettingsGetCustomizedSettingsResponse(), nil
}
// Deprecated: This class is obsolete. Use IntentsItemGetcustomizedsettingsGetCustomizedSettingsGetResponseable instead.
type IntentsItemGetcustomizedsettingsGetCustomizedSettingsResponseable interface {
    IntentsItemGetcustomizedsettingsGetCustomizedSettingsGetResponseable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
}
