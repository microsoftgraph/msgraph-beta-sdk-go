package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// TeamsAppIcon provides operations to manage the appCatalogs singleton.
type TeamsAppIcon struct {
    Entity
    // The contents of the app icon if the icon is hosted within the Teams infrastructure.
    hostedContent TeamworkHostedContentable;
    // The web URL that can be used for downloading the image.
    webUrl *string;
}
// NewTeamsAppIcon instantiates a new teamsAppIcon and sets the default values.
func NewTeamsAppIcon()(*TeamsAppIcon) {
    m := &TeamsAppIcon{
        Entity: *NewEntity(),
    }
    return m
}
// CreateTeamsAppIconFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateTeamsAppIconFromDiscriminatorValue(parseNode i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, error) {
    return NewTeamsAppIcon(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *TeamsAppIcon) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["hostedContent"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(CreateTeamworkHostedContentFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetHostedContent(val.(TeamworkHostedContentable))
        }
        return nil
    }
    res["webUrl"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetWebUrl(val)
        }
        return nil
    }
    return res
}
// GetHostedContent gets the hostedContent property value. The contents of the app icon if the icon is hosted within the Teams infrastructure.
func (m *TeamsAppIcon) GetHostedContent()(TeamworkHostedContentable) {
    if m == nil {
        return nil
    } else {
        return m.hostedContent
    }
}
// GetWebUrl gets the webUrl property value. The web URL that can be used for downloading the image.
func (m *TeamsAppIcon) GetWebUrl()(*string) {
    if m == nil {
        return nil
    } else {
        return m.webUrl
    }
}
func (m *TeamsAppIcon) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
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
// SetHostedContent sets the hostedContent property value. The contents of the app icon if the icon is hosted within the Teams infrastructure.
func (m *TeamsAppIcon) SetHostedContent(value TeamworkHostedContentable)() {
    if m != nil {
        m.hostedContent = value
    }
}
// SetWebUrl sets the webUrl property value. The web URL that can be used for downloading the image.
func (m *TeamsAppIcon) SetWebUrl(value *string)() {
    if m != nil {
        m.webUrl = value
    }
}
