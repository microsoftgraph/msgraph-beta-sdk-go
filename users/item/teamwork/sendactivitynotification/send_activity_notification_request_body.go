package sendactivitynotification

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph"
)

// 
type SendActivityNotificationRequestBody struct {
    // 
    activityType *string;
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // 
    chainId *int64;
    // 
    previewText *i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.ItemBody;
    // 
    templateParameters []i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.KeyValuePair;
    // 
    topic *i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.TeamworkActivityTopic;
}
// Instantiates a new sendActivityNotificationRequestBody and sets the default values.
func NewSendActivityNotificationRequestBody()(*SendActivityNotificationRequestBody) {
    m := &SendActivityNotificationRequestBody{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// Gets the activityType property value. 
func (m *SendActivityNotificationRequestBody) GetActivityType()(*string) {
    if m == nil {
        return nil
    } else {
        return m.activityType
    }
}
// Gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *SendActivityNotificationRequestBody) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// Gets the chainId property value. 
func (m *SendActivityNotificationRequestBody) GetChainId()(*int64) {
    if m == nil {
        return nil
    } else {
        return m.chainId
    }
}
// Gets the previewText property value. 
func (m *SendActivityNotificationRequestBody) GetPreviewText()(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.ItemBody) {
    if m == nil {
        return nil
    } else {
        return m.previewText
    }
}
// Gets the templateParameters property value. 
func (m *SendActivityNotificationRequestBody) GetTemplateParameters()([]i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.KeyValuePair) {
    if m == nil {
        return nil
    } else {
        return m.templateParameters
    }
}
// Gets the topic property value. 
func (m *SendActivityNotificationRequestBody) GetTopic()(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.TeamworkActivityTopic) {
    if m == nil {
        return nil
    } else {
        return m.topic
    }
}
// The deserialization information for the current model
func (m *SendActivityNotificationRequestBody) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["activityType"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetActivityType(val)
        return nil
    }
    res["chainId"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        m.SetChainId(val)
        return nil
    }
    res["previewText"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.NewItemBody() })
        if err != nil {
            return err
        }
        m.SetPreviewText(val.(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.ItemBody))
        return nil
    }
    res["templateParameters"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.NewKeyValuePair() })
        if err != nil {
            return err
        }
        res := make([]i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.KeyValuePair, len(val))
        for i, v := range val {
            res[i] = *(v.(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.KeyValuePair))
        }
        m.SetTemplateParameters(res)
        return nil
    }
    res["topic"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.NewTeamworkActivityTopic() })
        if err != nil {
            return err
        }
        m.SetTopic(val.(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.TeamworkActivityTopic))
        return nil
    }
    return res
}
func (m *SendActivityNotificationRequestBody) IsNil()(bool) {
    return m == nil
}
// Serializes information the current object
// Parameters:
//  - writer : Serialization writer to use to serialize this model
func (m *SendActivityNotificationRequestBody) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("activityType", m.GetActivityType())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt64Value("chainId", m.GetChainId())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("previewText", m.GetPreviewText())
        if err != nil {
            return err
        }
    }
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetTemplateParameters()))
        for i, v := range m.GetTemplateParameters() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err := writer.WriteCollectionOfObjectValues("templateParameters", cast)
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("topic", m.GetTopic())
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
// Sets the activityType property value. 
// Parameters:
//  - value : Value to set for the activityType property.
func (m *SendActivityNotificationRequestBody) SetActivityType(value *string)() {
    m.activityType = value
}
// Sets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// Parameters:
//  - value : Value to set for the AdditionalData property.
func (m *SendActivityNotificationRequestBody) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// Sets the chainId property value. 
// Parameters:
//  - value : Value to set for the chainId property.
func (m *SendActivityNotificationRequestBody) SetChainId(value *int64)() {
    m.chainId = value
}
// Sets the previewText property value. 
// Parameters:
//  - value : Value to set for the previewText property.
func (m *SendActivityNotificationRequestBody) SetPreviewText(value *i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.ItemBody)() {
    m.previewText = value
}
// Sets the templateParameters property value. 
// Parameters:
//  - value : Value to set for the templateParameters property.
func (m *SendActivityNotificationRequestBody) SetTemplateParameters(value []i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.KeyValuePair)() {
    m.templateParameters = value
}
// Sets the topic property value. 
// Parameters:
//  - value : Value to set for the topic property.
func (m *SendActivityNotificationRequestBody) SetTopic(value *i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.TeamworkActivityTopic)() {
    m.topic = value
}
