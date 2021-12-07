package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// RedirectUriSettings 
type RedirectUriSettings struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // Identifies the specific URI within the redirectURIs collection in SAML SSO flows. Defaults to null. The index is unique across all the redirectUris for the application.
    index *int32;
    // Specifies the URI that tokens are sent to.
    uri *string;
}
// NewRedirectUriSettings instantiates a new redirectUriSettings and sets the default values.
func NewRedirectUriSettings()(*RedirectUriSettings) {
    m := &RedirectUriSettings{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *RedirectUriSettings) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetIndex gets the index property value. Identifies the specific URI within the redirectURIs collection in SAML SSO flows. Defaults to null. The index is unique across all the redirectUris for the application.
func (m *RedirectUriSettings) GetIndex()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.index
    }
}
// GetUri gets the uri property value. Specifies the URI that tokens are sent to.
func (m *RedirectUriSettings) GetUri()(*string) {
    if m == nil {
        return nil
    } else {
        return m.uri
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *RedirectUriSettings) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["index"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIndex(val)
        }
        return nil
    }
    res["uri"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUri(val)
        }
        return nil
    }
    return res
}
func (m *RedirectUriSettings) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
func (m *RedirectUriSettings) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    {
        err := writer.WriteInt32Value("index", m.GetIndex())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("uri", m.GetUri())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteAdditionalData(m.GetAdditionalData())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAdditionalData sets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *RedirectUriSettings) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
// SetIndex sets the index property value. Identifies the specific URI within the redirectURIs collection in SAML SSO flows. Defaults to null. The index is unique across all the redirectUris for the application.
func (m *RedirectUriSettings) SetIndex(value *int32)() {
    if m != nil {
        m.index = value
    }
}
// SetUri sets the uri property value. Specifies the URI that tokens are sent to.
func (m *RedirectUriSettings) SetUri(value *string)() {
    if m != nil {
        m.uri = value
    }
}
