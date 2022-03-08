package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// TopicModelingSettings provides operations to manage the compliance singleton.
type TopicModelingSettings struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // To learn more, see Adjust maximum number of themes dynamically.
    dynamicallyAdjustTopicCount *bool;
    // To learn more, see Include numbers in themes.
    ignoreNumbers *bool;
    // Indicates whether themes is enabled for the case.
    isEnabled *bool;
    // To learn more, see Maximum number of themes.
    topicCount *int32;
}
// NewTopicModelingSettings instantiates a new topicModelingSettings and sets the default values.
func NewTopicModelingSettings()(*TopicModelingSettings) {
    m := &TopicModelingSettings{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateTopicModelingSettingsFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateTopicModelingSettingsFromDiscriminatorValue(parseNode i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, error) {
    return NewTopicModelingSettings(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *TopicModelingSettings) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetDynamicallyAdjustTopicCount gets the dynamicallyAdjustTopicCount property value. To learn more, see Adjust maximum number of themes dynamically.
func (m *TopicModelingSettings) GetDynamicallyAdjustTopicCount()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.dynamicallyAdjustTopicCount
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *TopicModelingSettings) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["dynamicallyAdjustTopicCount"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDynamicallyAdjustTopicCount(val)
        }
        return nil
    }
    res["ignoreNumbers"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIgnoreNumbers(val)
        }
        return nil
    }
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
    res["topicCount"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTopicCount(val)
        }
        return nil
    }
    return res
}
// GetIgnoreNumbers gets the ignoreNumbers property value. To learn more, see Include numbers in themes.
func (m *TopicModelingSettings) GetIgnoreNumbers()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.ignoreNumbers
    }
}
// GetIsEnabled gets the isEnabled property value. Indicates whether themes is enabled for the case.
func (m *TopicModelingSettings) GetIsEnabled()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.isEnabled
    }
}
// GetTopicCount gets the topicCount property value. To learn more, see Maximum number of themes.
func (m *TopicModelingSettings) GetTopicCount()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.topicCount
    }
}
func (m *TopicModelingSettings) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
func (m *TopicModelingSettings) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    {
        err := writer.WriteBoolValue("dynamicallyAdjustTopicCount", m.GetDynamicallyAdjustTopicCount())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("ignoreNumbers", m.GetIgnoreNumbers())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("isEnabled", m.GetIsEnabled())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("topicCount", m.GetTopicCount())
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
func (m *TopicModelingSettings) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
// SetDynamicallyAdjustTopicCount sets the dynamicallyAdjustTopicCount property value. To learn more, see Adjust maximum number of themes dynamically.
func (m *TopicModelingSettings) SetDynamicallyAdjustTopicCount(value *bool)() {
    if m != nil {
        m.dynamicallyAdjustTopicCount = value
    }
}
// SetIgnoreNumbers sets the ignoreNumbers property value. To learn more, see Include numbers in themes.
func (m *TopicModelingSettings) SetIgnoreNumbers(value *bool)() {
    if m != nil {
        m.ignoreNumbers = value
    }
}
// SetIsEnabled sets the isEnabled property value. Indicates whether themes is enabled for the case.
func (m *TopicModelingSettings) SetIsEnabled(value *bool)() {
    if m != nil {
        m.isEnabled = value
    }
}
// SetTopicCount sets the topicCount property value. To learn more, see Maximum number of themes.
func (m *TopicModelingSettings) SetTopicCount(value *int32)() {
    if m != nil {
        m.topicCount = value
    }
}
