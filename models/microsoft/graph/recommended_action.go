package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// 
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
// Instantiates a new recommendedAction and sets the default values.
func NewRecommendedAction()(*RecommendedAction) {
    m := &RecommendedAction{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// Gets the actionWebUrl property value. Web URL to the recommended action.
func (m *RecommendedAction) GetActionWebUrl()(*string) {
    if m == nil {
        return nil
    } else {
        return m.actionWebUrl
    }
}
// Gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *RecommendedAction) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// Gets the potentialScoreImpact property value. Potential improvement in security score of the tenant from the recommended action.
func (m *RecommendedAction) GetPotentialScoreImpact()(*float64) {
    if m == nil {
        return nil
    } else {
        return m.potentialScoreImpact
    }
}
// Gets the title property value. Title of the recommended action.
func (m *RecommendedAction) GetTitle()(*string) {
    if m == nil {
        return nil
    } else {
        return m.title
    }
}
// The deserialization information for the current model
func (m *RecommendedAction) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["actionWebUrl"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetActionWebUrl(val)
        return nil
    }
    res["potentialScoreImpact"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetFloat64Value()
        if err != nil {
            return err
        }
        m.SetPotentialScoreImpact(val)
        return nil
    }
    res["title"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetTitle(val)
        return nil
    }
    return res
}
func (m *RecommendedAction) IsNil()(bool) {
    return m == nil
}
// Serializes information the current object
// Parameters:
//  - writer : Serialization writer to use to serialize this model
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
// Sets the actionWebUrl property value. Web URL to the recommended action.
// Parameters:
//  - value : Value to set for the actionWebUrl property.
func (m *RecommendedAction) SetActionWebUrl(value *string)() {
    m.actionWebUrl = value
}
// Sets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// Parameters:
//  - value : Value to set for the AdditionalData property.
func (m *RecommendedAction) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// Sets the potentialScoreImpact property value. Potential improvement in security score of the tenant from the recommended action.
// Parameters:
//  - value : Value to set for the potentialScoreImpact property.
func (m *RecommendedAction) SetPotentialScoreImpact(value *float64)() {
    m.potentialScoreImpact = value
}
// Sets the title property value. Title of the recommended action.
// Parameters:
//  - value : Value to set for the title property.
func (m *RecommendedAction) SetTitle(value *string)() {
    m.title = value
}
