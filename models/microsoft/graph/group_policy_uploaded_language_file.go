package graph

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// 
type GroupPolicyUploadedLanguageFile struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // The contents of the uploaded ADML file.
    content []byte;
    // The file name of the uploaded ADML file.
    fileName *string;
    // Key of the entity.
    id *string;
    // The language code of the uploaded ADML file.
    languageCode *string;
    // The date and time the entity was last modified.
    lastModifiedDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time;
}
// Instantiates a new groupPolicyUploadedLanguageFile and sets the default values.
func NewGroupPolicyUploadedLanguageFile()(*GroupPolicyUploadedLanguageFile) {
    m := &GroupPolicyUploadedLanguageFile{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// Gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *GroupPolicyUploadedLanguageFile) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// Gets the content property value. The contents of the uploaded ADML file.
func (m *GroupPolicyUploadedLanguageFile) GetContent()([]byte) {
    if m == nil {
        return nil
    } else {
        return m.content
    }
}
// Gets the fileName property value. The file name of the uploaded ADML file.
func (m *GroupPolicyUploadedLanguageFile) GetFileName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.fileName
    }
}
// Gets the id property value. Key of the entity.
func (m *GroupPolicyUploadedLanguageFile) GetId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.id
    }
}
// Gets the languageCode property value. The language code of the uploaded ADML file.
func (m *GroupPolicyUploadedLanguageFile) GetLanguageCode()(*string) {
    if m == nil {
        return nil
    } else {
        return m.languageCode
    }
}
// Gets the lastModifiedDateTime property value. The date and time the entity was last modified.
func (m *GroupPolicyUploadedLanguageFile) GetLastModifiedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.lastModifiedDateTime
    }
}
// The deserialization information for the current model
func (m *GroupPolicyUploadedLanguageFile) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["content"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetByteArrayValue()
        if err != nil {
            return err
        }
        m.SetContent(val)
        return nil
    }
    res["fileName"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetFileName(val)
        return nil
    }
    res["id"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetId(val)
        return nil
    }
    res["languageCode"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetLanguageCode(val)
        return nil
    }
    res["lastModifiedDateTime"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        m.SetLastModifiedDateTime(val)
        return nil
    }
    return res
}
func (m *GroupPolicyUploadedLanguageFile) IsNil()(bool) {
    return m == nil
}
// Serializes information the current object
// Parameters:
//  - writer : Serialization writer to use to serialize this model
func (m *GroupPolicyUploadedLanguageFile) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    {
        err := writer.WriteByteArrayValue("content", m.GetContent())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("fileName", m.GetFileName())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("id", m.GetId())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("languageCode", m.GetLanguageCode())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteTimeValue("lastModifiedDateTime", m.GetLastModifiedDateTime())
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
// Sets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// Parameters:
//  - value : Value to set for the AdditionalData property.
func (m *GroupPolicyUploadedLanguageFile) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// Sets the content property value. The contents of the uploaded ADML file.
// Parameters:
//  - value : Value to set for the content property.
func (m *GroupPolicyUploadedLanguageFile) SetContent(value []byte)() {
    m.content = value
}
// Sets the fileName property value. The file name of the uploaded ADML file.
// Parameters:
//  - value : Value to set for the fileName property.
func (m *GroupPolicyUploadedLanguageFile) SetFileName(value *string)() {
    m.fileName = value
}
// Sets the id property value. Key of the entity.
// Parameters:
//  - value : Value to set for the id property.
func (m *GroupPolicyUploadedLanguageFile) SetId(value *string)() {
    m.id = value
}
// Sets the languageCode property value. The language code of the uploaded ADML file.
// Parameters:
//  - value : Value to set for the languageCode property.
func (m *GroupPolicyUploadedLanguageFile) SetLanguageCode(value *string)() {
    m.languageCode = value
}
// Sets the lastModifiedDateTime property value. The date and time the entity was last modified.
// Parameters:
//  - value : Value to set for the lastModifiedDateTime property.
func (m *GroupPolicyUploadedLanguageFile) SetLastModifiedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.lastModifiedDateTime = value
}
