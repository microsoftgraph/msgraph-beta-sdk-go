package graph

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// Mention 
type Mention struct {
    Entity
    // The name of the application where the mention is created. Optional. Not used and defaulted as null for message.
    application *string;
    // A unique identifier that represents a parent of the resource instance. Optional. Not used and defaulted as null for message.
    clientReference *string;
    // The email information of the user who made the mention.
    createdBy *EmailAddress;
    // The date and time that the mention is created on the client.
    createdDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time;
    // A deep web link to the context of the mention in the resource instance. Optional. Not used and defaulted as null for message.
    deepLink *string;
    // 
    mentioned *EmailAddress;
    // Optional. Not used and defaulted as null for message. To get the mentions in a message, see the bodyPreview property of the message instead.
    mentionText *string;
    // The date and time that the mention is created on the server. Optional. Not used and defaulted as null for message.
    serverCreatedDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time;
}
// NewMention instantiates a new mention and sets the default values.
func NewMention()(*Mention) {
    m := &Mention{
        Entity: *NewEntity(),
    }
    return m
}
// GetApplication gets the application property value. The name of the application where the mention is created. Optional. Not used and defaulted as null for message.
func (m *Mention) GetApplication()(*string) {
    if m == nil {
        return nil
    } else {
        return m.application
    }
}
// GetClientReference gets the clientReference property value. A unique identifier that represents a parent of the resource instance. Optional. Not used and defaulted as null for message.
func (m *Mention) GetClientReference()(*string) {
    if m == nil {
        return nil
    } else {
        return m.clientReference
    }
}
// GetCreatedBy gets the createdBy property value. The email information of the user who made the mention.
func (m *Mention) GetCreatedBy()(*EmailAddress) {
    if m == nil {
        return nil
    } else {
        return m.createdBy
    }
}
// GetCreatedDateTime gets the createdDateTime property value. The date and time that the mention is created on the client.
func (m *Mention) GetCreatedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.createdDateTime
    }
}
// GetDeepLink gets the deepLink property value. A deep web link to the context of the mention in the resource instance. Optional. Not used and defaulted as null for message.
func (m *Mention) GetDeepLink()(*string) {
    if m == nil {
        return nil
    } else {
        return m.deepLink
    }
}
// GetMentioned gets the mentioned property value. 
func (m *Mention) GetMentioned()(*EmailAddress) {
    if m == nil {
        return nil
    } else {
        return m.mentioned
    }
}
// GetMentionText gets the mentionText property value. Optional. Not used and defaulted as null for message. To get the mentions in a message, see the bodyPreview property of the message instead.
func (m *Mention) GetMentionText()(*string) {
    if m == nil {
        return nil
    } else {
        return m.mentionText
    }
}
// GetServerCreatedDateTime gets the serverCreatedDateTime property value. The date and time that the mention is created on the server. Optional. Not used and defaulted as null for message.
func (m *Mention) GetServerCreatedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.serverCreatedDateTime
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *Mention) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["application"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetApplication(val)
        }
        return nil
    }
    res["clientReference"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetClientReference(val)
        }
        return nil
    }
    res["createdBy"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewEmailAddress() })
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCreatedBy(val.(*EmailAddress))
        }
        return nil
    }
    res["createdDateTime"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCreatedDateTime(val)
        }
        return nil
    }
    res["deepLink"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDeepLink(val)
        }
        return nil
    }
    res["mentioned"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewEmailAddress() })
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMentioned(val.(*EmailAddress))
        }
        return nil
    }
    res["mentionText"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMentionText(val)
        }
        return nil
    }
    res["serverCreatedDateTime"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetServerCreatedDateTime(val)
        }
        return nil
    }
    return res
}
func (m *Mention) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
func (m *Mention) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteStringValue("application", m.GetApplication())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("clientReference", m.GetClientReference())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("createdBy", m.GetCreatedBy())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteTimeValue("createdDateTime", m.GetCreatedDateTime())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("deepLink", m.GetDeepLink())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("mentioned", m.GetMentioned())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("mentionText", m.GetMentionText())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteTimeValue("serverCreatedDateTime", m.GetServerCreatedDateTime())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetApplication sets the application property value. The name of the application where the mention is created. Optional. Not used and defaulted as null for message.
func (m *Mention) SetApplication(value *string)() {
    if m != nil {
        m.application = value
    }
}
// SetClientReference sets the clientReference property value. A unique identifier that represents a parent of the resource instance. Optional. Not used and defaulted as null for message.
func (m *Mention) SetClientReference(value *string)() {
    if m != nil {
        m.clientReference = value
    }
}
// SetCreatedBy sets the createdBy property value. The email information of the user who made the mention.
func (m *Mention) SetCreatedBy(value *EmailAddress)() {
    if m != nil {
        m.createdBy = value
    }
}
// SetCreatedDateTime sets the createdDateTime property value. The date and time that the mention is created on the client.
func (m *Mention) SetCreatedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    if m != nil {
        m.createdDateTime = value
    }
}
// SetDeepLink sets the deepLink property value. A deep web link to the context of the mention in the resource instance. Optional. Not used and defaulted as null for message.
func (m *Mention) SetDeepLink(value *string)() {
    if m != nil {
        m.deepLink = value
    }
}
// SetMentioned sets the mentioned property value. 
func (m *Mention) SetMentioned(value *EmailAddress)() {
    if m != nil {
        m.mentioned = value
    }
}
// SetMentionText sets the mentionText property value. Optional. Not used and defaulted as null for message. To get the mentions in a message, see the bodyPreview property of the message instead.
func (m *Mention) SetMentionText(value *string)() {
    if m != nil {
        m.mentionText = value
    }
}
// SetServerCreatedDateTime sets the serverCreatedDateTime property value. The date and time that the mention is created on the server. Optional. Not used and defaulted as null for message.
func (m *Mention) SetServerCreatedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    if m != nil {
        m.serverCreatedDateTime = value
    }
}
