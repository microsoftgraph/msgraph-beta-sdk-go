package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// TeamworkLoginStatus 
type TeamworkLoginStatus struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // Information about the Exchange connection.
    exchangeConnection *TeamworkConnection;
    // Information about the Skype for Business connection.
    skypeConnection *TeamworkConnection;
    // Information about the Teams connection.
    teamsConnection *TeamworkConnection;
}
// NewTeamworkLoginStatus instantiates a new teamworkLoginStatus and sets the default values.
func NewTeamworkLoginStatus()(*TeamworkLoginStatus) {
    m := &TeamworkLoginStatus{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *TeamworkLoginStatus) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetExchangeConnection gets the exchangeConnection property value. Information about the Exchange connection.
func (m *TeamworkLoginStatus) GetExchangeConnection()(*TeamworkConnection) {
    if m == nil {
        return nil
    } else {
        return m.exchangeConnection
    }
}
// GetSkypeConnection gets the skypeConnection property value. Information about the Skype for Business connection.
func (m *TeamworkLoginStatus) GetSkypeConnection()(*TeamworkConnection) {
    if m == nil {
        return nil
    } else {
        return m.skypeConnection
    }
}
// GetTeamsConnection gets the teamsConnection property value. Information about the Teams connection.
func (m *TeamworkLoginStatus) GetTeamsConnection()(*TeamworkConnection) {
    if m == nil {
        return nil
    } else {
        return m.teamsConnection
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *TeamworkLoginStatus) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["exchangeConnection"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewTeamworkConnection() })
        if err != nil {
            return err
        }
        if val != nil {
            m.SetExchangeConnection(val.(*TeamworkConnection))
        }
        return nil
    }
    res["skypeConnection"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewTeamworkConnection() })
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSkypeConnection(val.(*TeamworkConnection))
        }
        return nil
    }
    res["teamsConnection"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewTeamworkConnection() })
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTeamsConnection(val.(*TeamworkConnection))
        }
        return nil
    }
    return res
}
func (m *TeamworkLoginStatus) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
func (m *TeamworkLoginStatus) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    {
        err := writer.WriteObjectValue("exchangeConnection", m.GetExchangeConnection())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("skypeConnection", m.GetSkypeConnection())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("teamsConnection", m.GetTeamsConnection())
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
func (m *TeamworkLoginStatus) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
// SetExchangeConnection sets the exchangeConnection property value. Information about the Exchange connection.
func (m *TeamworkLoginStatus) SetExchangeConnection(value *TeamworkConnection)() {
    if m != nil {
        m.exchangeConnection = value
    }
}
// SetSkypeConnection sets the skypeConnection property value. Information about the Skype for Business connection.
func (m *TeamworkLoginStatus) SetSkypeConnection(value *TeamworkConnection)() {
    if m != nil {
        m.skypeConnection = value
    }
}
// SetTeamsConnection sets the teamsConnection property value. Information about the Teams connection.
func (m *TeamworkLoginStatus) SetTeamsConnection(value *TeamworkConnection)() {
    if m != nil {
        m.teamsConnection = value
    }
}
