package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// RecommendedAction 
type RecommendedAction struct {
    // Web URL to the recommended action.
    actionWebUrl *string;
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // Potential improvement in security score of the tenant from the recommended action.
    potentialScoreImpact *float64;
    // Title of the recommended action.
    title *string;
}
// NewRecommendedAction instantiates a new recommendedAction and sets the default values.
func NewRecommendedAction()(*RecommendedAction) {
    m := &RecommendedAction{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// GetActionWebUrl gets the actionWebUrl property value. Web URL to the recommended action.
func (m *RecommendedAction) GetActionWebUrl()(*string) {
    if m == nil {
        return nil
    } else {
        return m.actionWebUrl
    }
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *RecommendedAction) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetPotentialScoreImpact gets the potentialScoreImpact property value. Potential improvement in security score of the tenant from the recommended action.
func (m *RecommendedAction) GetPotentialScoreImpact()(*float64) {
    if m == nil {
        return nil
    } else {
        return m.potentialScoreImpact
    }
}
// GetTitle gets the title property value. Title of the recommended action.
func (m *RecommendedAction) GetTitle()(*string) {
    if m == nil {
        return nil
    } else {
        return m.title
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *RecommendedAction) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["actionWebUrl"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetActionWebUrl(val)
        }
        return nil
    }
    res["potentialScoreImpact"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetFloat64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPotentialScoreImpact(val)
        }
        return nil
    }
    res["title"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTitle(val)
        }
        return nil
    }
    return res
}
func (m *RecommendedAction) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
func (m *RecommendedAction) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("actionWebUrl", m.GetActionWebUrl())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteFloat64Value("potentialScoreImpact", m.GetPotentialScoreImpact())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("title", m.GetTitle())
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
// SetActionWebUrl sets the actionWebUrl property value. Web URL to the recommended action.
func (m *RecommendedAction) SetActionWebUrl(value *string)() {
    if m != nil {
        m.actionWebUrl = value
    }
}
// SetAdditionalData sets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *RecommendedAction) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
// SetPotentialScoreImpact sets the potentialScoreImpact property value. Potential improvement in security score of the tenant from the recommended action.
func (m *RecommendedAction) SetPotentialScoreImpact(value *float64)() {
    if m != nil {
        m.potentialScoreImpact = value
    }
}
// SetTitle sets the title property value. Title of the recommended action.
func (m *RecommendedAction) SetTitle(value *string)() {
    if m != nil {
        m.title = value
    }
}
