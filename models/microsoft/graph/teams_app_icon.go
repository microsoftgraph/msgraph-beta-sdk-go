package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

type TeamsAppIcon struct {
    Entity
    hostedContent *TeamworkHostedContent;
    webUrl *string;
}
func NewTeamsAppIcon()(*TeamsAppIcon) {
    m := &TeamsAppIcon{
        Entity: *NewEntity(),
    }
    return m
}
func (m *TeamsAppIcon) GetHostedContent()(*TeamworkHostedContent) {
    if m == nil {
        return nil
    } else {
        return m.hostedContent
    }
}
func (m *TeamsAppIcon) GetWebUrl()(*string) {
    if m == nil {
        return nil
    } else {
        return m.webUrl
    }
}
func (m *TeamsAppIcon) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["hostedContent"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewTeamworkHostedContent() })
        if err != nil {
            return err
        }
        m.SetHostedContent(val.(*TeamworkHostedContent))
        return nil
    }
    res["webUrl"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetWebUrl(val)
        return nil
    }
    return res
}
func (m *TeamsAppIcon) IsNil()(bool) {
    return m == nil
}
func (m *TeamsAppIcon) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteObjectValue("hostedContent", m.GetHostedContent())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("webUrl", m.GetWebUrl())
        if err != nil {
            return err
        }
    }
    return nil
}
func (m *TeamsAppIcon) SetHostedContent(value *TeamworkHostedContent)() {
    m.hostedContent = value
}
func (m *TeamsAppIcon) SetWebUrl(value *string)() {
    m.webUrl = value
}
