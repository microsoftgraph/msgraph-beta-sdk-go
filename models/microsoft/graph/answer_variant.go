package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// AnswerVariant provides operations to manage the searchEntity singleton.
type AnswerVariant struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // Answer variation description shown on search results page.
    description *string;
    // Answer variation name displayed in search results.
    displayName *string;
    // 
    languageTag *string;
    // 
    platform *DevicePlatformType;
    // Answer variation URL link. When users click this answer variation in search results, they will go to this URL.
    webUrl *string;
}
// NewAnswerVariant instantiates a new answerVariant and sets the default values.
func NewAnswerVariant()(*AnswerVariant) {
    m := &AnswerVariant{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateAnswerVariantFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateAnswerVariantFromDiscriminatorValue(parseNode i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, error) {
    return NewAnswerVariant(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *AnswerVariant) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetDescription gets the description property value. Answer variation description shown on search results page.
func (m *AnswerVariant) GetDescription()(*string) {
    if m == nil {
        return nil
    } else {
        return m.description
    }
}
// GetDisplayName gets the displayName property value. Answer variation name displayed in search results.
func (m *AnswerVariant) GetDisplayName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.displayName
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *AnswerVariant) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["description"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDescription(val)
        }
        return nil
    }
    res["displayName"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDisplayName(val)
        }
        return nil
    }
    res["languageTag"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLanguageTag(val)
        }
        return nil
    }
    res["platform"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParseDevicePlatformType)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPlatform(val.(*DevicePlatformType))
        }
        return nil
    }
    res["webUrl"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetWebUrl(val)
        }
        return nil
    }
    return res
}
// GetLanguageTag gets the languageTag property value. 
func (m *AnswerVariant) GetLanguageTag()(*string) {
    if m == nil {
        return nil
    } else {
        return m.languageTag
    }
}
// GetPlatform gets the platform property value. 
func (m *AnswerVariant) GetPlatform()(*DevicePlatformType) {
    if m == nil {
        return nil
    } else {
        return m.platform
    }
}
// GetWebUrl gets the webUrl property value. Answer variation URL link. When users click this answer variation in search results, they will go to this URL.
func (m *AnswerVariant) GetWebUrl()(*string) {
    if m == nil {
        return nil
    } else {
        return m.webUrl
    }
}
func (m *AnswerVariant) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
func (m *AnswerVariant) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("description", m.GetDescription())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("displayName", m.GetDisplayName())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("languageTag", m.GetLanguageTag())
        if err != nil {
            return err
        }
    }
    if m.GetPlatform() != nil {
        cast := (*m.GetPlatform()).String()
        err := writer.WriteStringValue("platform", &cast)
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("webUrl", m.GetWebUrl())
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
func (m *AnswerVariant) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
// SetDescription sets the description property value. Answer variation description shown on search results page.
func (m *AnswerVariant) SetDescription(value *string)() {
    if m != nil {
        m.description = value
    }
}
// SetDisplayName sets the displayName property value. Answer variation name displayed in search results.
func (m *AnswerVariant) SetDisplayName(value *string)() {
    if m != nil {
        m.displayName = value
    }
}
// SetLanguageTag sets the languageTag property value. 
func (m *AnswerVariant) SetLanguageTag(value *string)() {
    if m != nil {
        m.languageTag = value
    }
}
// SetPlatform sets the platform property value. 
func (m *AnswerVariant) SetPlatform(value *DevicePlatformType)() {
    if m != nil {
        m.platform = value
    }
}
// SetWebUrl sets the webUrl property value. Answer variation URL link. When users click this answer variation in search results, they will go to this URL.
func (m *AnswerVariant) SetWebUrl(value *string)() {
    if m != nil {
        m.webUrl = value
    }
}
