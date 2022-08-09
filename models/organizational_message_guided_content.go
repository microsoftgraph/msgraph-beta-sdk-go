package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// OrganizationalMessageGuidedContent this will allow the admin to see the different templated organizational messages that can be created
type OrganizationalMessageGuidedContent struct {
    Entity
    // Example of the logo that will be displayed to customers and its size requirements
    logo OrganizationalMessageLogoGuideable
    // Contains the different types of text content that can be displayed to customers along with their localized values
    placementDetails []OrganizationalMessagePlacementDetailable
    // Indicates the scenario for the message
    scenario *OrganizationalMessageScenario
    // Indicates the area where content will be displayed to customers
    surface *OrganizationalMessageSurface
    // Indicates the theme for the guided content
    theme *OrganizationalMessageTheme
}
// NewOrganizationalMessageGuidedContent instantiates a new organizationalMessageGuidedContent and sets the default values.
func NewOrganizationalMessageGuidedContent()(*OrganizationalMessageGuidedContent) {
    m := &OrganizationalMessageGuidedContent{
        Entity: *NewEntity(),
    }
    odataTypeValue := "#microsoft.graph.organizationalMessageGuidedContent";
    m.SetOdataType(&odataTypeValue);
    return m
}
// CreateOrganizationalMessageGuidedContentFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateOrganizationalMessageGuidedContentFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewOrganizationalMessageGuidedContent(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *OrganizationalMessageGuidedContent) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["logo"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateOrganizationalMessageLogoGuideFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLogo(val.(OrganizationalMessageLogoGuideable))
        }
        return nil
    }
    res["placementDetails"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateOrganizationalMessagePlacementDetailFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]OrganizationalMessagePlacementDetailable, len(val))
            for i, v := range val {
                res[i] = v.(OrganizationalMessagePlacementDetailable)
            }
            m.SetPlacementDetails(res)
        }
        return nil
    }
    res["scenario"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseOrganizationalMessageScenario)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetScenario(val.(*OrganizationalMessageScenario))
        }
        return nil
    }
    res["surface"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseOrganizationalMessageSurface)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSurface(val.(*OrganizationalMessageSurface))
        }
        return nil
    }
    res["theme"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseOrganizationalMessageTheme)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTheme(val.(*OrganizationalMessageTheme))
        }
        return nil
    }
    return res
}
// GetLogo gets the logo property value. Example of the logo that will be displayed to customers and its size requirements
func (m *OrganizationalMessageGuidedContent) GetLogo()(OrganizationalMessageLogoGuideable) {
    if m == nil {
        return nil
    } else {
        return m.logo
    }
}
// GetPlacementDetails gets the placementDetails property value. Contains the different types of text content that can be displayed to customers along with their localized values
func (m *OrganizationalMessageGuidedContent) GetPlacementDetails()([]OrganizationalMessagePlacementDetailable) {
    if m == nil {
        return nil
    } else {
        return m.placementDetails
    }
}
// GetScenario gets the scenario property value. Indicates the scenario for the message
func (m *OrganizationalMessageGuidedContent) GetScenario()(*OrganizationalMessageScenario) {
    if m == nil {
        return nil
    } else {
        return m.scenario
    }
}
// GetSurface gets the surface property value. Indicates the area where content will be displayed to customers
func (m *OrganizationalMessageGuidedContent) GetSurface()(*OrganizationalMessageSurface) {
    if m == nil {
        return nil
    } else {
        return m.surface
    }
}
// GetTheme gets the theme property value. Indicates the theme for the guided content
func (m *OrganizationalMessageGuidedContent) GetTheme()(*OrganizationalMessageTheme) {
    if m == nil {
        return nil
    } else {
        return m.theme
    }
}
// Serialize serializes information the current object
func (m *OrganizationalMessageGuidedContent) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteObjectValue("logo", m.GetLogo())
        if err != nil {
            return err
        }
    }
    if m.GetPlacementDetails() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetPlacementDetails()))
        for i, v := range m.GetPlacementDetails() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("placementDetails", cast)
        if err != nil {
            return err
        }
    }
    if m.GetScenario() != nil {
        cast := (*m.GetScenario()).String()
        err = writer.WriteStringValue("scenario", &cast)
        if err != nil {
            return err
        }
    }
    if m.GetSurface() != nil {
        cast := (*m.GetSurface()).String()
        err = writer.WriteStringValue("surface", &cast)
        if err != nil {
            return err
        }
    }
    if m.GetTheme() != nil {
        cast := (*m.GetTheme()).String()
        err = writer.WriteStringValue("theme", &cast)
        if err != nil {
            return err
        }
    }
    return nil
}
// SetLogo sets the logo property value. Example of the logo that will be displayed to customers and its size requirements
func (m *OrganizationalMessageGuidedContent) SetLogo(value OrganizationalMessageLogoGuideable)() {
    if m != nil {
        m.logo = value
    }
}
// SetPlacementDetails sets the placementDetails property value. Contains the different types of text content that can be displayed to customers along with their localized values
func (m *OrganizationalMessageGuidedContent) SetPlacementDetails(value []OrganizationalMessagePlacementDetailable)() {
    if m != nil {
        m.placementDetails = value
    }
}
// SetScenario sets the scenario property value. Indicates the scenario for the message
func (m *OrganizationalMessageGuidedContent) SetScenario(value *OrganizationalMessageScenario)() {
    if m != nil {
        m.scenario = value
    }
}
// SetSurface sets the surface property value. Indicates the area where content will be displayed to customers
func (m *OrganizationalMessageGuidedContent) SetSurface(value *OrganizationalMessageSurface)() {
    if m != nil {
        m.surface = value
    }
}
// SetTheme sets the theme property value. Indicates the theme for the guided content
func (m *OrganizationalMessageGuidedContent) SetTheme(value *OrganizationalMessageTheme)() {
    if m != nil {
        m.theme = value
    }
}
