package ediscovery

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph"
)

// CaseSettings provides operations to manage the compliance singleton.
type CaseSettings struct {
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.Entity
    // The OCR (Optical Character Recognition) settings for the case.
    ocr OcrSettingsable;
    // The redundancy (near duplicate and email threading) detection settings for the case.
    redundancyDetection RedundancyDetectionSettingsable;
    // The Topic Modeling (Themes) settings for the case.
    topicModeling TopicModelingSettingsable;
}
// NewCaseSettings instantiates a new caseSettings and sets the default values.
func NewCaseSettings()(*CaseSettings) {
    m := &CaseSettings{
        Entity: *i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.NewEntity(),
    }
    return m
}
// CreateCaseSettingsFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateCaseSettingsFromDiscriminatorValue(parseNode i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, error) {
    return NewCaseSettings(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *CaseSettings) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["ocr"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(CreateOcrSettingsFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOcr(val.(OcrSettingsable))
        }
        return nil
    }
    res["redundancyDetection"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(CreateRedundancyDetectionSettingsFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRedundancyDetection(val.(RedundancyDetectionSettingsable))
        }
        return nil
    }
    res["topicModeling"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(CreateTopicModelingSettingsFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTopicModeling(val.(TopicModelingSettingsable))
        }
        return nil
    }
    return res
}
// GetOcr gets the ocr property value. The OCR (Optical Character Recognition) settings for the case.
func (m *CaseSettings) GetOcr()(OcrSettingsable) {
    if m == nil {
        return nil
    } else {
        return m.ocr
    }
}
// GetRedundancyDetection gets the redundancyDetection property value. The redundancy (near duplicate and email threading) detection settings for the case.
func (m *CaseSettings) GetRedundancyDetection()(RedundancyDetectionSettingsable) {
    if m == nil {
        return nil
    } else {
        return m.redundancyDetection
    }
}
// GetTopicModeling gets the topicModeling property value. The Topic Modeling (Themes) settings for the case.
func (m *CaseSettings) GetTopicModeling()(TopicModelingSettingsable) {
    if m == nil {
        return nil
    } else {
        return m.topicModeling
    }
}
func (m *CaseSettings) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
func (m *CaseSettings) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteObjectValue("ocr", m.GetOcr())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("redundancyDetection", m.GetRedundancyDetection())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("topicModeling", m.GetTopicModeling())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetOcr sets the ocr property value. The OCR (Optical Character Recognition) settings for the case.
func (m *CaseSettings) SetOcr(value OcrSettingsable)() {
    if m != nil {
        m.ocr = value
    }
}
// SetRedundancyDetection sets the redundancyDetection property value. The redundancy (near duplicate and email threading) detection settings for the case.
func (m *CaseSettings) SetRedundancyDetection(value RedundancyDetectionSettingsable)() {
    if m != nil {
        m.redundancyDetection = value
    }
}
// SetTopicModeling sets the topicModeling property value. The Topic Modeling (Themes) settings for the case.
func (m *CaseSettings) SetTopicModeling(value TopicModelingSettingsable)() {
    if m != nil {
        m.topicModeling = value
    }
}
