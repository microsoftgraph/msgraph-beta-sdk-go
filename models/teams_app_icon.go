package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// TeamsAppIcon provides operations to manage the collection of accessReview entities.
type TeamsAppIcon struct {
    Entity
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // The contents of the app icon if the icon is hosted within the Teams infrastructure.
    hostedContent TeamworkHostedContentable
    // The web URL that can be used for downloading the image.
    webUrl *string
}
// NewTeamsAppIcon instantiates a new teamsAppIcon and sets the default values.
func NewTeamsAppIcon()(*TeamsAppIcon) {
    m := &TeamsAppIcon{
        Entity: *NewEntity(),
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateTeamsAppIconFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateTeamsAppIconFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewTeamsAppIcon(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *TeamsAppIcon) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *TeamsAppIcon) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["hostedContent"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateTeamworkHostedContentFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetHostedContent(val.(TeamworkHostedContentable))
        }
        return nil
    }
    res["webUrl"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
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
// Serialize serializes information the current object
func (m *TeamsAppIcon) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
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
    {
        err = writer.WriteAdditionalData(m.GetAdditionalData())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAdditionalData sets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *TeamsAppIcon) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
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
