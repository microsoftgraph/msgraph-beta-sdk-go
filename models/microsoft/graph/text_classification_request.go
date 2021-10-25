package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

type TextClassificationRequest struct {
    Entity
    fileExtension *string;
    matchTolerancesToInclude *MlClassificationMatchTolerance;
    scopesToRun *SensitiveTypeScope;
    sensitiveTypeIds []string;
    text *string;
}
func NewTextClassificationRequest()(*TextClassificationRequest) {
    m := &TextClassificationRequest{
        Entity: *NewEntity(),
    }
    return m
}
func (m *TextClassificationRequest) GetFileExtension()(*string) {
    if m == nil {
        return nil
    } else {
        return m.fileExtension
    }
}
func (m *TextClassificationRequest) GetMatchTolerancesToInclude()(*MlClassificationMatchTolerance) {
    if m == nil {
        return nil
    } else {
        return m.matchTolerancesToInclude
    }
}
func (m *TextClassificationRequest) GetScopesToRun()(*SensitiveTypeScope) {
    if m == nil {
        return nil
    } else {
        return m.scopesToRun
    }
}
func (m *TextClassificationRequest) GetSensitiveTypeIds()([]string) {
    if m == nil {
        return nil
    } else {
        return m.sensitiveTypeIds
    }
}
func (m *TextClassificationRequest) GetText()(*string) {
    if m == nil {
        return nil
    } else {
        return m.text
    }
}
func (m *TextClassificationRequest) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["fileExtension"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetFileExtension(val)
        return nil
    }
    res["matchTolerancesToInclude"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParseMlClassificationMatchTolerance)
        if err != nil {
            return err
        }
        cast := val.(MlClassificationMatchTolerance)
        m.SetMatchTolerancesToInclude(&cast)
        return nil
    }
    res["scopesToRun"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParseSensitiveTypeScope)
        if err != nil {
            return err
        }
        cast := val.(SensitiveTypeScope)
        m.SetScopesToRun(&cast)
        return nil
    }
    res["sensitiveTypeIds"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        res := make([]string, len(val))
        for i, v := range val {
            res[i] = v.(string)
        }
        m.SetSensitiveTypeIds(res)
        return nil
    }
    res["text"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetText(val)
        return nil
    }
    return res
}
func (m *TextClassificationRequest) IsNil()(bool) {
    return m == nil
}
func (m *TextClassificationRequest) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteStringValue("fileExtension", m.GetFileExtension())
        if err != nil {
            return err
        }
    }
    if m.GetMatchTolerancesToInclude() != nil {
        cast := m.GetMatchTolerancesToInclude().String()
        err = writer.WriteStringValue("matchTolerancesToInclude", &cast)
        if err != nil {
            return err
        }
    }
    if m.GetScopesToRun() != nil {
        cast := m.GetScopesToRun().String()
        err = writer.WriteStringValue("scopesToRun", &cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteCollectionOfStringValues("sensitiveTypeIds", m.GetSensitiveTypeIds())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("text", m.GetText())
        if err != nil {
            return err
        }
    }
    return nil
}
func (m *TextClassificationRequest) SetFileExtension(value *string)() {
    m.fileExtension = value
}
func (m *TextClassificationRequest) SetMatchTolerancesToInclude(value *MlClassificationMatchTolerance)() {
    m.matchTolerancesToInclude = value
}
func (m *TextClassificationRequest) SetScopesToRun(value *SensitiveTypeScope)() {
    m.scopesToRun = value
}
func (m *TextClassificationRequest) SetSensitiveTypeIds(value []string)() {
    m.sensitiveTypeIds = value
}
func (m *TextClassificationRequest) SetText(value *string)() {
    m.text = value
}
