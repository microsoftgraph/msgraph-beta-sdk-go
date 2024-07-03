package industrydata

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type OAuthClientCredential struct {
    Credential
}
// NewOAuthClientCredential instantiates a new OAuthClientCredential and sets the default values.
func NewOAuthClientCredential()(*OAuthClientCredential) {
    m := &OAuthClientCredential{
        Credential: *NewCredential(),
    }
    odataTypeValue := "#microsoft.graph.industryData.oAuthClientCredential"
    m.SetOdataType(&odataTypeValue)
    return m
}
// CreateOAuthClientCredentialFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateOAuthClientCredentialFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
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
                    case "#microsoft.graph.industryData.oAuth1ClientCredential":
                        return NewOAuth1ClientCredential(), nil
                    case "#microsoft.graph.industryData.oAuth2ClientCredential":
                        return NewOAuth2ClientCredential(), nil
                }
            }
        }
    }
    return NewOAuthClientCredential(), nil
}
// GetClientId gets the clientId property value. The client identifier of the application that is authenticating.
// returns a *string when successful
func (m *OAuthClientCredential) GetClientId()(*string) {
    val, err := m.GetBackingStore().Get("clientId")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetClientSecret gets the clientSecret property value. The client secret that is used to authenticate (write-only).
// returns a *string when successful
func (m *OAuthClientCredential) GetClientSecret()(*string) {
    val, err := m.GetBackingStore().Get("clientSecret")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *OAuthClientCredential) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Credential.GetFieldDeserializers()
    res["clientId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetClientId(val)
        }
        return nil
    }
    res["clientSecret"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetClientSecret(val)
        }
        return nil
    }
    return res
}
// Serialize serializes information the current object
func (m *OAuthClientCredential) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Credential.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteStringValue("clientId", m.GetClientId())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("clientSecret", m.GetClientSecret())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetClientId sets the clientId property value. The client identifier of the application that is authenticating.
func (m *OAuthClientCredential) SetClientId(value *string)() {
    err := m.GetBackingStore().Set("clientId", value)
    if err != nil {
        panic(err)
    }
}
// SetClientSecret sets the clientSecret property value. The client secret that is used to authenticate (write-only).
func (m *OAuthClientCredential) SetClientSecret(value *string)() {
    err := m.GetBackingStore().Set("clientSecret", value)
    if err != nil {
        panic(err)
    }
}
type OAuthClientCredentialable interface {
    Credentialable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetClientId()(*string)
    GetClientSecret()(*string)
    SetClientId(value *string)()
    SetClientSecret(value *string)()
}
