package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

type MoveAction struct {
    additionalData map[string]interface{};
    from *string;
    to *string;
}
func NewMoveAction()(*MoveAction) {
    m := &MoveAction{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
func (m *MoveAction) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
func (m *MoveAction) GetFrom()(*string) {
    if m == nil {
        return nil
    } else {
        return m.from
    }
}
func (m *MoveAction) GetTo()(*string) {
    if m == nil {
        return nil
    } else {
        return m.to
    }
}
func (m *MoveAction) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["from"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetFrom(val)
        return nil
    }
    res["to"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetTo(val)
        return nil
    }
    return res
}
func (m *MoveAction) IsNil()(bool) {
    return m == nil
}
func (m *MoveAction) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("from", m.GetFrom())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("to", m.GetTo())
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
func (m *MoveAction) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
func (m *MoveAction) SetFrom(value *string)() {
    m.from = value
}
func (m *MoveAction) SetTo(value *string)() {
    m.to = value
}
