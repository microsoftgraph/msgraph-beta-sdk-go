package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// RedundancyDetectionSettings 
type RedundancyDetectionSettings struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // Indicates whether email threading and near duplicate detection are enabled.
    isEnabled *bool;
    // See Minimum/maximum number of words to learn more.
    maxWords *int32;
    // See Minimum/maximum number of words to learn more.
    minWords *int32;
    // See Document and email similarity threshold to learn more.
    similarityThreshold *int32;
}
// NewRedundancyDetectionSettings instantiates a new redundancyDetectionSettings and sets the default values.
func NewRedundancyDetectionSettings()(*RedundancyDetectionSettings) {
    m := &RedundancyDetectionSettings{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *RedundancyDetectionSettings) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetIsEnabled gets the isEnabled property value. Indicates whether email threading and near duplicate detection are enabled.
func (m *RedundancyDetectionSettings) GetIsEnabled()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.isEnabled
    }
}
// GetMaxWords gets the maxWords property value. See Minimum/maximum number of words to learn more.
func (m *RedundancyDetectionSettings) GetMaxWords()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.maxWords
    }
}
// GetMinWords gets the minWords property value. See Minimum/maximum number of words to learn more.
func (m *RedundancyDetectionSettings) GetMinWords()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.minWords
    }
}
// GetSimilarityThreshold gets the similarityThreshold property value. See Document and email similarity threshold to learn more.
func (m *RedundancyDetectionSettings) GetSimilarityThreshold()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.similarityThreshold
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *RedundancyDetectionSettings) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["isEnabled"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsEnabled(val)
        }
        return nil
    }
    res["maxWords"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMaxWords(val)
        }
        return nil
    }
    res["minWords"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMinWords(val)
        }
        return nil
    }
    res["similarityThreshold"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSimilarityThreshold(val)
        }
        return nil
    }
    return res
}
func (m *RedundancyDetectionSettings) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
func (m *RedundancyDetectionSettings) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    {
        err := writer.WriteBoolValue("isEnabled", m.GetIsEnabled())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("maxWords", m.GetMaxWords())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("minWords", m.GetMinWords())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("similarityThreshold", m.GetSimilarityThreshold())
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
// SetAdditionalData sets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *RedundancyDetectionSettings) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
// SetIsEnabled sets the isEnabled property value. Indicates whether email threading and near duplicate detection are enabled.
func (m *RedundancyDetectionSettings) SetIsEnabled(value *bool)() {
    if m != nil {
        m.isEnabled = value
    }
}
// SetMaxWords sets the maxWords property value. See Minimum/maximum number of words to learn more.
func (m *RedundancyDetectionSettings) SetMaxWords(value *int32)() {
    if m != nil {
        m.maxWords = value
    }
}
// SetMinWords sets the minWords property value. See Minimum/maximum number of words to learn more.
func (m *RedundancyDetectionSettings) SetMinWords(value *int32)() {
    if m != nil {
        m.minWords = value
    }
}
// SetSimilarityThreshold sets the similarityThreshold property value. See Document and email similarity threshold to learn more.
func (m *RedundancyDetectionSettings) SetSimilarityThreshold(value *int32)() {
    if m != nil {
        m.similarityThreshold = value
    }
}
