package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

type RecordingInfo struct {
    additionalData map[string]interface{};
    initiatedBy *ParticipantInfo;
    initiator *IdentitySet;
    recordingStatus *RecordingStatus;
}
func NewRecordingInfo()(*RecordingInfo) {
    m := &RecordingInfo{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
func (m *RecordingInfo) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
func (m *RecordingInfo) GetInitiatedBy()(*ParticipantInfo) {
    if m == nil {
        return nil
    } else {
        return m.initiatedBy
    }
}
func (m *RecordingInfo) GetInitiator()(*IdentitySet) {
    if m == nil {
        return nil
    } else {
        return m.initiator
    }
}
func (m *RecordingInfo) GetRecordingStatus()(*RecordingStatus) {
    if m == nil {
        return nil
    } else {
        return m.recordingStatus
    }
}
func (m *RecordingInfo) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["initiatedBy"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewParticipantInfo() })
        if err != nil {
            return err
        }
        m.SetInitiatedBy(val.(*ParticipantInfo))
        return nil
    }
    res["initiator"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewIdentitySet() })
        if err != nil {
            return err
        }
        m.SetInitiator(val.(*IdentitySet))
        return nil
    }
    res["recordingStatus"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParseRecordingStatus)
        if err != nil {
            return err
        }
        cast := val.(RecordingStatus)
        m.SetRecordingStatus(&cast)
        return nil
    }
    return res
}
func (m *RecordingInfo) IsNil()(bool) {
    return m == nil
}
func (m *RecordingInfo) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    {
        err := writer.WriteObjectValue("initiatedBy", m.GetInitiatedBy())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("initiator", m.GetInitiator())
        if err != nil {
            return err
        }
    }
    if m.GetRecordingStatus() != nil {
        cast := m.GetRecordingStatus().String()
        err := writer.WriteStringValue("recordingStatus", &cast)
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
func (m *RecordingInfo) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
func (m *RecordingInfo) SetInitiatedBy(value *ParticipantInfo)() {
    m.initiatedBy = value
}
func (m *RecordingInfo) SetInitiator(value *IdentitySet)() {
    m.initiator = value
}
func (m *RecordingInfo) SetRecordingStatus(value *RecordingStatus)() {
    m.recordingStatus = value
}
