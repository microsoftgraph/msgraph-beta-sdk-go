package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    if9792b9f78e4933d4010103471e3a29a1979828da75ed72eb040984ee8d10c0f "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph/callrecords"
)

type FailureInfo struct {
    additionalData map[string]interface{};
    reason *string;
    stage *if9792b9f78e4933d4010103471e3a29a1979828da75ed72eb040984ee8d10c0f.FailureStage;
}
func NewFailureInfo()(*FailureInfo) {
    m := &FailureInfo{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
func (m *FailureInfo) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
func (m *FailureInfo) GetReason()(*string) {
    if m == nil {
        return nil
    } else {
        return m.reason
    }
}
func (m *FailureInfo) GetStage()(*if9792b9f78e4933d4010103471e3a29a1979828da75ed72eb040984ee8d10c0f.FailureStage) {
    if m == nil {
        return nil
    } else {
        return m.stage
    }
}
func (m *FailureInfo) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["reason"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetReason(val)
        return nil
    }
    res["stage"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(if9792b9f78e4933d4010103471e3a29a1979828da75ed72eb040984ee8d10c0f.ParseFailureStage)
        if err != nil {
            return err
        }
        cast := val.(if9792b9f78e4933d4010103471e3a29a1979828da75ed72eb040984ee8d10c0f.FailureStage)
        m.SetStage(&cast)
        return nil
    }
    return res
}
func (m *FailureInfo) IsNil()(bool) {
    return m == nil
}
func (m *FailureInfo) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("reason", m.GetReason())
        if err != nil {
            return err
        }
    }
    if m.GetStage() != nil {
        cast := m.GetStage().String()
        err := writer.WriteStringValue("stage", &cast)
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
func (m *FailureInfo) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
func (m *FailureInfo) SetReason(value *string)() {
    m.reason = value
}
func (m *FailureInfo) SetStage(value *if9792b9f78e4933d4010103471e3a29a1979828da75ed72eb040984ee8d10c0f.FailureStage)() {
    m.stage = value
}
