package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// 
type AnswerKeyword struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // 
    keywords []string;
    // 
    matchSimilarKeywords *bool;
    // 
    reservedKeywords []string;
}
// Instantiates a new answerKeyword and sets the default values.
func NewAnswerKeyword()(*AnswerKeyword) {
    m := &AnswerKeyword{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// Gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *AnswerKeyword) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// Gets the keywords property value. 
func (m *AnswerKeyword) GetKeywords()([]string) {
    if m == nil {
        return nil
    } else {
        return m.keywords
    }
}
// Gets the matchSimilarKeywords property value. 
func (m *AnswerKeyword) GetMatchSimilarKeywords()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.matchSimilarKeywords
    }
}
// Gets the reservedKeywords property value. 
func (m *AnswerKeyword) GetReservedKeywords()([]string) {
    if m == nil {
        return nil
    } else {
        return m.reservedKeywords
    }
}
// The deserialization information for the current model
func (m *AnswerKeyword) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["keywords"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]string, len(val))
            for i, v := range val {
                res[i] = *(v.(*string))
            }
            m.SetKeywords(res)
        }
        return nil
    }
    res["matchSimilarKeywords"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMatchSimilarKeywords(val)
        }
        return nil
    }
    res["reservedKeywords"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]string, len(val))
            for i, v := range val {
                res[i] = *(v.(*string))
            }
            m.SetReservedKeywords(res)
        }
        return nil
    }
    return res
}
func (m *AnswerKeyword) IsNil()(bool) {
    return m == nil
}
// Serializes information the current object
// Parameters:
//  - writer : Serialization writer to use to serialize this model
func (m *AnswerKeyword) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    {
        err := writer.WriteCollectionOfStringValues("keywords", m.GetKeywords())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("matchSimilarKeywords", m.GetMatchSimilarKeywords())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteCollectionOfStringValues("reservedKeywords", m.GetReservedKeywords())
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
func (m *AnswerKeyword) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// Sets the keywords property value. 
// Parameters:
//  - value : Value to set for the keywords property.
func (m *AnswerKeyword) SetKeywords(value []string)() {
    m.keywords = value
}
// Sets the matchSimilarKeywords property value. 
// Parameters:
//  - value : Value to set for the matchSimilarKeywords property.
func (m *AnswerKeyword) SetMatchSimilarKeywords(value *bool)() {
    m.matchSimilarKeywords = value
}
// Sets the reservedKeywords property value. 
// Parameters:
//  - value : Value to set for the reservedKeywords property.
func (m *AnswerKeyword) SetReservedKeywords(value []string)() {
    m.reservedKeywords = value
}
