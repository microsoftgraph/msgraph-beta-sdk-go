package graph

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// 
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
// Instantiates a new mention and sets the default values.
func NewMention()(*Mention) {
    m := &Mention{
        Entity: *NewEntity(),
    }
    return m
}
// Gets the application property value. The name of the application where the mention is created. Optional. Not used and defaulted as null for message.
func (m *Mention) GetApplication()(*string) {
    if m == nil {
        return nil
    } else {
        return m.application
    }
}
// Gets the clientReference property value. A unique identifier that represents a parent of the resource instance. Optional. Not used and defaulted as null for message.
func (m *Mention) GetClientReference()(*string) {
    if m == nil {
        return nil
    } else {
        return m.clientReference
    }
}
// Gets the createdBy property value. The email information of the user who made the mention.
func (m *Mention) GetCreatedBy()(*EmailAddress) {
    if m == nil {
        return nil
    } else {
        return m.createdBy
    }
}
// Gets the createdDateTime property value. The date and time that the mention is created on the client.
func (m *Mention) GetCreatedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.createdDateTime
    }
}
// Gets the deepLink property value. A deep web link to the context of the mention in the resource instance. Optional. Not used and defaulted as null for message.
func (m *Mention) GetDeepLink()(*string) {
    if m == nil {
        return nil
    } else {
        return m.deepLink
    }
}
// Gets the mentioned property value. 
func (m *Mention) GetMentioned()(*EmailAddress) {
    if m == nil {
        return nil
    } else {
        return m.mentioned
    }
}
// Gets the mentionText property value. Optional. Not used and defaulted as null for message. To get the mentions in a message, see the bodyPreview property of the message instead.
func (m *Mention) GetMentionText()(*string) {
    if m == nil {
        return nil
    } else {
        return m.mentionText
    }
}
// Gets the serverCreatedDateTime property value. The date and time that the mention is created on the server. Optional. Not used and defaulted as null for message.
func (m *Mention) GetServerCreatedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.serverCreatedDateTime
    }
}
// The deserialization information for the current model
func (m *Mention) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["application"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetApplication(val)
        return nil
    }
    res["clientReference"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetClientReference(val)
        return nil
    }
    res["createdBy"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewEmailAddress() })
        if err != nil {
            return err
        }
        m.SetCreatedBy(val.(*EmailAddress))
        return nil
    }
    res["createdDateTime"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        m.SetCreatedDateTime(val)
        return nil
    }
    res["deepLink"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetDeepLink(val)
        return nil
    }
    res["mentioned"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewEmailAddress() })
        if err != nil {
            return err
        }
        m.SetMentioned(val.(*EmailAddress))
        return nil
    }
    res["mentionText"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetMentionText(val)
        return nil
    }
    res["serverCreatedDateTime"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        m.SetServerCreatedDateTime(val)
        return nil
    }
    return res
}
func (m *Mention) IsNil()(bool) {
    return m == nil
}
// Serializes information the current object
// Parameters:
//  - writer : Serialization writer to use to serialize this model
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
// Sets the application property value. The name of the application where the mention is created. Optional. Not used and defaulted as null for message.
// Parameters:
//  - value : Value to set for the application property.
func (m *Mention) SetApplication(value *string)() {
    m.application = value
}
// Sets the clientReference property value. A unique identifier that represents a parent of the resource instance. Optional. Not used and defaulted as null for message.
// Parameters:
//  - value : Value to set for the clientReference property.
func (m *Mention) SetClientReference(value *string)() {
    m.clientReference = value
}
// Sets the createdBy property value. The email information of the user who made the mention.
// Parameters:
//  - value : Value to set for the createdBy property.
func (m *Mention) SetCreatedBy(value *EmailAddress)() {
    m.createdBy = value
}
// Sets the createdDateTime property value. The date and time that the mention is created on the client.
// Parameters:
//  - value : Value to set for the createdDateTime property.
func (m *Mention) SetCreatedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.createdDateTime = value
}
// Sets the deepLink property value. A deep web link to the context of the mention in the resource instance. Optional. Not used and defaulted as null for message.
// Parameters:
//  - value : Value to set for the deepLink property.
func (m *Mention) SetDeepLink(value *string)() {
    m.deepLink = value
}
// Sets the mentioned property value. 
// Parameters:
//  - value : Value to set for the mentioned property.
func (m *Mention) SetMentioned(value *EmailAddress)() {
    m.mentioned = value
}
// Sets the mentionText property value. Optional. Not used and defaulted as null for message. To get the mentions in a message, see the bodyPreview property of the message instead.
// Parameters:
//  - value : Value to set for the mentionText property.
func (m *Mention) SetMentionText(value *string)() {
    m.mentionText = value
}
// Sets the serverCreatedDateTime property value. The date and time that the mention is created on the server. Optional. Not used and defaulted as null for message.
// Parameters:
//  - value : Value to set for the serverCreatedDateTime property.
func (m *Mention) SetServerCreatedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.serverCreatedDateTime = value
}
