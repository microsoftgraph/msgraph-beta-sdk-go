package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// OnTokenIssuanceStartCustomExtension 
type OnTokenIssuanceStartCustomExtension struct {
    CustomAuthenticationExtension
    // The claimsForTokenConfiguration property
    claimsForTokenConfiguration []OnTokenIssuanceStartReturnClaimable
}
// NewOnTokenIssuanceStartCustomExtension instantiates a new OnTokenIssuanceStartCustomExtension and sets the default values.
func NewOnTokenIssuanceStartCustomExtension()(*OnTokenIssuanceStartCustomExtension) {
    m := &OnTokenIssuanceStartCustomExtension{
        CustomAuthenticationExtension: *NewCustomAuthenticationExtension(),
    }
    odataTypeValue := "#microsoft.graph.onTokenIssuanceStartCustomExtension";
    m.SetOdataType(&odataTypeValue);
    return m
}
// CreateOnTokenIssuanceStartCustomExtensionFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateOnTokenIssuanceStartCustomExtensionFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewOnTokenIssuanceStartCustomExtension(), nil
}
// GetClaimsForTokenConfiguration gets the claimsForTokenConfiguration property value. The claimsForTokenConfiguration property
func (m *OnTokenIssuanceStartCustomExtension) GetClaimsForTokenConfiguration()([]OnTokenIssuanceStartReturnClaimable) {
    return m.claimsForTokenConfiguration
}
// GetFieldDeserializers the deserialization information for the current model
func (m *OnTokenIssuanceStartCustomExtension) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.CustomAuthenticationExtension.GetFieldDeserializers()
    res["claimsForTokenConfiguration"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfObjectValues(CreateOnTokenIssuanceStartReturnClaimFromDiscriminatorValue , m.SetClaimsForTokenConfiguration)
    return res
}
// Serialize serializes information the current object
func (m *OnTokenIssuanceStartCustomExtension) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.CustomAuthenticationExtension.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetClaimsForTokenConfiguration() != nil {
        cast := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.CollectionCast[i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable](m.GetClaimsForTokenConfiguration())
        err = writer.WriteCollectionOfObjectValues("claimsForTokenConfiguration", cast)
        if err != nil {
            return err
        }
    }
    return nil
}
// SetClaimsForTokenConfiguration sets the claimsForTokenConfiguration property value. The claimsForTokenConfiguration property
func (m *OnTokenIssuanceStartCustomExtension) SetClaimsForTokenConfiguration(value []OnTokenIssuanceStartReturnClaimable)() {
    m.claimsForTokenConfiguration = value
}
