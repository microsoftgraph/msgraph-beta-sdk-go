package industrydata

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type ApiDataConnector struct {
    IndustryDataConnector
}
// NewApiDataConnector instantiates a new ApiDataConnector and sets the default values.
func NewApiDataConnector()(*ApiDataConnector) {
    m := &ApiDataConnector{
        IndustryDataConnector: *NewIndustryDataConnector(),
    }
    odataTypeValue := "#microsoft.graph.industryData.apiDataConnector"
    m.SetOdataType(&odataTypeValue)
    return m
}
// CreateApiDataConnectorFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateApiDataConnectorFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
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
                    case "#microsoft.graph.industryData.oneRosterApiDataConnector":
                        return NewOneRosterApiDataConnector(), nil
                }
            }
        }
    }
    return NewApiDataConnector(), nil
}
// GetApiFormat gets the apiFormat property value. The apiFormat property
// returns a *ApiFormat when successful
func (m *ApiDataConnector) GetApiFormat()(*ApiFormat) {
    val, err := m.GetBackingStore().Get("apiFormat")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*ApiFormat)
    }
    return nil
}
// GetBaseUrl gets the baseUrl property value. The base URL, including the scheme, host, and path for the API, with or without a trailing '/'. For example, 'https://example.com/ims/oneRoster/v1p1'
// returns a *string when successful
func (m *ApiDataConnector) GetBaseUrl()(*string) {
    val, err := m.GetBackingStore().Get("baseUrl")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetCredential gets the credential property value. The credential property
// returns a Credentialable when successful
func (m *ApiDataConnector) GetCredential()(Credentialable) {
    val, err := m.GetBackingStore().Get("credential")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(Credentialable)
    }
    return nil
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *ApiDataConnector) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.IndustryDataConnector.GetFieldDeserializers()
    res["apiFormat"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseApiFormat)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetApiFormat(val.(*ApiFormat))
        }
        return nil
    }
    res["baseUrl"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetBaseUrl(val)
        }
        return nil
    }
    res["credential"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateCredentialFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCredential(val.(Credentialable))
        }
        return nil
    }
    return res
}
// Serialize serializes information the current object
func (m *ApiDataConnector) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.IndustryDataConnector.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetApiFormat() != nil {
        cast := (*m.GetApiFormat()).String()
        err = writer.WriteStringValue("apiFormat", &cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("baseUrl", m.GetBaseUrl())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("credential", m.GetCredential())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetApiFormat sets the apiFormat property value. The apiFormat property
func (m *ApiDataConnector) SetApiFormat(value *ApiFormat)() {
    err := m.GetBackingStore().Set("apiFormat", value)
    if err != nil {
        panic(err)
    }
}
// SetBaseUrl sets the baseUrl property value. The base URL, including the scheme, host, and path for the API, with or without a trailing '/'. For example, 'https://example.com/ims/oneRoster/v1p1'
func (m *ApiDataConnector) SetBaseUrl(value *string)() {
    err := m.GetBackingStore().Set("baseUrl", value)
    if err != nil {
        panic(err)
    }
}
// SetCredential sets the credential property value. The credential property
func (m *ApiDataConnector) SetCredential(value Credentialable)() {
    err := m.GetBackingStore().Set("credential", value)
    if err != nil {
        panic(err)
    }
}
type ApiDataConnectorable interface {
    IndustryDataConnectorable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetApiFormat()(*ApiFormat)
    GetBaseUrl()(*string)
    GetCredential()(Credentialable)
    SetApiFormat(value *ApiFormat)()
    SetBaseUrl(value *string)()
    SetCredential(value Credentialable)()
}
