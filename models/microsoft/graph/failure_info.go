package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    if9792b9f78e4933d4010103471e3a29a1979828da75ed72eb040984ee8d10c0f "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph/callrecords"
)

// FailureInfo 
type FailureInfo struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // Classification of why a call or portion of a call failed.
    reason *string;
    // The stage when the failure occurred. Possible values are: unknown, callSetup, midcall, unknownFutureValue.
    stage *if9792b9f78e4933d4010103471e3a29a1979828da75ed72eb040984ee8d10c0f.FailureStage;
}
// NewFailureInfo instantiates a new failureInfo and sets the default values.
func NewFailureInfo()(*FailureInfo) {
    m := &FailureInfo{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *FailureInfo) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetReason gets the reason property value. Classification of why a call or portion of a call failed.
func (m *FailureInfo) GetReason()(*string) {
    if m == nil {
        return nil
    } else {
        return m.reason
    }
}
// GetStage gets the stage property value. The stage when the failure occurred. Possible values are: unknown, callSetup, midcall, unknownFutureValue.
func (m *FailureInfo) GetStage()(*if9792b9f78e4933d4010103471e3a29a1979828da75ed72eb040984ee8d10c0f.FailureStage) {
    if m == nil {
        return nil
    } else {
        return m.stage
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *FailureInfo) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["reason"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetReason(val)
        }
        return nil
    }
    res["stage"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(if9792b9f78e4933d4010103471e3a29a1979828da75ed72eb040984ee8d10c0f.ParseFailureStage)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetStage(val.(*if9792b9f78e4933d4010103471e3a29a1979828da75ed72eb040984ee8d10c0f.FailureStage))
        }
        return nil
    }
    return res
}
func (m *FailureInfo) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
func (m *FailureInfo) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("reason", m.GetReason())
        if err != nil {
            return err
        }
    }
    if m.GetStage() != nil {
        cast := (*m.GetStage()).String()
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
// SetAdditionalData sets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *FailureInfo) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
// SetReason sets the reason property value. Classification of why a call or portion of a call failed.
func (m *FailureInfo) SetReason(value *string)() {
    if m != nil {
        m.reason = value
    }
}
// SetStage sets the stage property value. The stage when the failure occurred. Possible values are: unknown, callSetup, midcall, unknownFutureValue.
func (m *FailureInfo) SetStage(value *if9792b9f78e4933d4010103471e3a29a1979828da75ed72eb040984ee8d10c0f.FailureStage)() {
    if m != nil {
        m.stage = value
    }
}
