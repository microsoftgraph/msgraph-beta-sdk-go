package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// AnswerKeyword 
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
// NewAnswerKeyword instantiates a new answerKeyword and sets the default values.
func NewAnswerKeyword()(*AnswerKeyword) {
    m := &AnswerKeyword{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *AnswerKeyword) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetKeywords gets the keywords property value. 
func (m *AnswerKeyword) GetKeywords()([]string) {
    if m == nil {
        return nil
    } else {
        return m.keywords
    }
}
// GetMatchSimilarKeywords gets the matchSimilarKeywords property value. 
func (m *AnswerKeyword) GetMatchSimilarKeywords()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.matchSimilarKeywords
    }
}
// GetReservedKeywords gets the reservedKeywords property value. 
func (m *AnswerKeyword) GetReservedKeywords()([]string) {
    if m == nil {
        return nil
    } else {
        return m.reservedKeywords
    }
}
// GetFieldDeserializers the deserialization information for the current model
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
// Serialize serializes information the current object
func (m *AnswerKeyword) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    if m.GetKeywords() != nil {
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
    if m.GetReservedKeywords() != nil {
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
// SetAdditionalData sets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *AnswerKeyword) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
// SetKeywords sets the keywords property value. 
func (m *AnswerKeyword) SetKeywords(value []string)() {
    if m != nil {
        m.keywords = value
    }
}
// SetMatchSimilarKeywords sets the matchSimilarKeywords property value. 
func (m *AnswerKeyword) SetMatchSimilarKeywords(value *bool)() {
    if m != nil {
        m.matchSimilarKeywords = value
    }
}
// SetReservedKeywords sets the reservedKeywords property value. 
func (m *AnswerKeyword) SetReservedKeywords(value []string)() {
    if m != nil {
        m.reservedKeywords = value
    }
}
