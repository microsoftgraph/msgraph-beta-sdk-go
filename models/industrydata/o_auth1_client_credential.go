package industrydata

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type OAuth1ClientCredential struct {
    OAuthClientCredential
}
// NewOAuth1ClientCredential instantiates a new OAuth1ClientCredential and sets the default values.
func NewOAuth1ClientCredential()(*OAuth1ClientCredential) {
    m := &OAuth1ClientCredential{
        OAuthClientCredential: *NewOAuthClientCredential(),
    }
    odataTypeValue := "#microsoft.graph.industryData.oAuth1ClientCredential"
    m.SetOdataType(&odataTypeValue)
    return m
}
// CreateOAuth1ClientCredentialFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateOAuth1ClientCredentialFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewOAuth1ClientCredential(), nil
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *OAuth1ClientCredential) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.OAuthClientCredential.GetFieldDeserializers()
    return res
}
// Serialize serializes information the current object
func (m *OAuth1ClientCredential) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.OAuthClientCredential.Serialize(writer)
    if err != nil {
        return err
    }
    return nil
}
type OAuth1ClientCredentialable interface {
    OAuthClientCredentialable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
}
