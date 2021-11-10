package graph

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// 
type Deployment struct {
    Entity
    // Specifies the audience to which content is deployed.
    audience *DeploymentAudience;
    // Specifies what content to deploy. Cannot be changed. Returned by default.
    content *DeployableContent;
    // The date and time the deployment was created. Returned by default. Read-only.
    createdDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time;
    // The date and time the deployment was last modified. Returned by default. Read-only.
    lastModifiedDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time;
    // Settings specified on the specific deployment governing how to deploy content. Returned by default.
    settings *DeploymentSettings;
    // Execution status of the deployment. Returned by default.
    state *DeploymentState;
}
// Instantiates a new deployment and sets the default values.
func NewDeployment()(*Deployment) {
    m := &Deployment{
        Entity: *NewEntity(),
    }
    return m
}
// Gets the audience property value. Specifies the audience to which content is deployed.
func (m *Deployment) GetAudience()(*DeploymentAudience) {
    if m == nil {
        return nil
    } else {
        return m.audience
    }
}
// Gets the content property value. Specifies what content to deploy. Cannot be changed. Returned by default.
func (m *Deployment) GetContent()(*DeployableContent) {
    if m == nil {
        return nil
    } else {
        return m.content
    }
}
// Gets the createdDateTime property value. The date and time the deployment was created. Returned by default. Read-only.
func (m *Deployment) GetCreatedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.createdDateTime
    }
}
// Gets the lastModifiedDateTime property value. The date and time the deployment was last modified. Returned by default. Read-only.
func (m *Deployment) GetLastModifiedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.lastModifiedDateTime
    }
}
// Gets the settings property value. Settings specified on the specific deployment governing how to deploy content. Returned by default.
func (m *Deployment) GetSettings()(*DeploymentSettings) {
    if m == nil {
        return nil
    } else {
        return m.settings
    }
}
// Gets the state property value. Execution status of the deployment. Returned by default.
func (m *Deployment) GetState()(*DeploymentState) {
    if m == nil {
        return nil
    } else {
        return m.state
    }
}
// The deserialization information for the current model
func (m *Deployment) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["audience"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewDeploymentAudience() })
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAudience(val.(*DeploymentAudience))
        }
        return nil
    }
    res["content"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewDeployableContent() })
        if err != nil {
            return err
        }
        if val != nil {
            m.SetContent(val.(*DeployableContent))
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
    res["lastModifiedDateTime"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLastModifiedDateTime(val)
        }
        return nil
    }
    res["settings"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewDeploymentSettings() })
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSettings(val.(*DeploymentSettings))
        }
        return nil
    }
    res["state"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewDeploymentState() })
        if err != nil {
            return err
        }
        if val != nil {
            m.SetState(val.(*DeploymentState))
        }
        return nil
    }
    return res
}
func (m *Deployment) IsNil()(bool) {
    return m == nil
}
// Serializes information the current object
// Parameters:
//  - writer : Serialization writer to use to serialize this model
func (m *Deployment) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteObjectValue("audience", m.GetAudience())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("content", m.GetContent())
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
        err = writer.WriteTimeValue("lastModifiedDateTime", m.GetLastModifiedDateTime())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("settings", m.GetSettings())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("state", m.GetState())
        if err != nil {
            return err
        }
    }
    return nil
}
// Sets the audience property value. Specifies the audience to which content is deployed.
// Parameters:
//  - value : Value to set for the audience property.
func (m *Deployment) SetAudience(value *DeploymentAudience)() {
    m.audience = value
}
// Sets the content property value. Specifies what content to deploy. Cannot be changed. Returned by default.
// Parameters:
//  - value : Value to set for the content property.
func (m *Deployment) SetContent(value *DeployableContent)() {
    m.content = value
}
// Sets the createdDateTime property value. The date and time the deployment was created. Returned by default. Read-only.
// Parameters:
//  - value : Value to set for the createdDateTime property.
func (m *Deployment) SetCreatedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.createdDateTime = value
}
// Sets the lastModifiedDateTime property value. The date and time the deployment was last modified. Returned by default. Read-only.
// Parameters:
//  - value : Value to set for the lastModifiedDateTime property.
func (m *Deployment) SetLastModifiedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.lastModifiedDateTime = value
}
// Sets the settings property value. Settings specified on the specific deployment governing how to deploy content. Returned by default.
// Parameters:
//  - value : Value to set for the settings property.
func (m *Deployment) SetSettings(value *DeploymentSettings)() {
    m.settings = value
}
// Sets the state property value. Execution status of the deployment. Returned by default.
// Parameters:
//  - value : Value to set for the state property.
func (m *Deployment) SetState(value *DeploymentState)() {
    m.state = value
}
