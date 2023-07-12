package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// TeamTemplateDefinition 
type TeamTemplateDefinition struct {
    Entity
    // The OdataType property
    OdataType *string
}
// NewTeamTemplateDefinition instantiates a new teamTemplateDefinition and sets the default values.
func NewTeamTemplateDefinition()(*TeamTemplateDefinition) {
    m := &TeamTemplateDefinition{
        Entity: *NewEntity(),
    }
    return m
}
// CreateTeamTemplateDefinitionFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateTeamTemplateDefinitionFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewTeamTemplateDefinition(), nil
}
// GetAudience gets the audience property value. Describes the audience the team template is available to. The possible values are: organization, user, public, unknownFutureValue.
func (m *TeamTemplateDefinition) GetAudience()(*TeamTemplateAudience) {
    val, err := m.GetBackingStore().Get("audience")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*TeamTemplateAudience)
    }
    return nil
}
// GetCategories gets the categories property value. The assigned categories for the team template.
func (m *TeamTemplateDefinition) GetCategories()([]string) {
    val, err := m.GetBackingStore().Get("categories")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]string)
    }
    return nil
}
// GetDescription gets the description property value. A brief description of the team template as it will appear to the users in Microsoft Teams.
func (m *TeamTemplateDefinition) GetDescription()(*string) {
    val, err := m.GetBackingStore().Get("description")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetDisplayName gets the displayName property value. The user defined name of the team template.
func (m *TeamTemplateDefinition) GetDisplayName()(*string) {
    val, err := m.GetBackingStore().Get("displayName")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *TeamTemplateDefinition) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["audience"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseTeamTemplateAudience)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAudience(val.(*TeamTemplateAudience))
        }
        return nil
    }
    res["categories"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]string, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = *(v.(*string))
                }
            }
            m.SetCategories(res)
        }
        return nil
    }
    res["description"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDescription(val)
        }
        return nil
    }
    res["displayName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDisplayName(val)
        }
        return nil
    }
    res["iconUrl"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIconUrl(val)
        }
        return nil
    }
    res["languageTag"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLanguageTag(val)
        }
        return nil
    }
    res["lastModifiedBy"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateIdentitySetFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLastModifiedBy(val.(IdentitySetable))
        }
        return nil
    }
    res["lastModifiedDateTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLastModifiedDateTime(val)
        }
        return nil
    }
    res["parentTemplateId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetParentTemplateId(val)
        }
        return nil
    }
    res["publisherName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPublisherName(val)
        }
        return nil
    }
    res["shortDescription"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetShortDescription(val)
        }
        return nil
    }
    res["teamDefinition"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateTeamFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTeamDefinition(val.(Teamable))
        }
        return nil
    }
    return res
}
// GetIconUrl gets the iconUrl property value. The icon url for the team template.
func (m *TeamTemplateDefinition) GetIconUrl()(*string) {
    val, err := m.GetBackingStore().Get("iconUrl")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetLanguageTag gets the languageTag property value. Language the template is available in.
func (m *TeamTemplateDefinition) GetLanguageTag()(*string) {
    val, err := m.GetBackingStore().Get("languageTag")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetLastModifiedBy gets the lastModifiedBy property value. The identity of the user who last modified the team template.
func (m *TeamTemplateDefinition) GetLastModifiedBy()(IdentitySetable) {
    val, err := m.GetBackingStore().Get("lastModifiedBy")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(IdentitySetable)
    }
    return nil
}
// GetLastModifiedDateTime gets the lastModifiedDateTime property value. The date time of when the team template was last modified.
func (m *TeamTemplateDefinition) GetLastModifiedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    val, err := m.GetBackingStore().Get("lastModifiedDateTime")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    }
    return nil
}
// GetParentTemplateId gets the parentTemplateId property value. The templateId for the team template
func (m *TeamTemplateDefinition) GetParentTemplateId()(*string) {
    val, err := m.GetBackingStore().Get("parentTemplateId")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetPublisherName gets the publisherName property value. The organization which published the team template.
func (m *TeamTemplateDefinition) GetPublisherName()(*string) {
    val, err := m.GetBackingStore().Get("publisherName")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetShortDescription gets the shortDescription property value. A short-description of the team template as it will appear to the users in Microsoft Teams.
func (m *TeamTemplateDefinition) GetShortDescription()(*string) {
    val, err := m.GetBackingStore().Get("shortDescription")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetTeamDefinition gets the teamDefinition property value. Collection of channel objects. A channel represents a topic, and therefore a logical isolation of discussion, within a team.
func (m *TeamTemplateDefinition) GetTeamDefinition()(Teamable) {
    val, err := m.GetBackingStore().Get("teamDefinition")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(Teamable)
    }
    return nil
}
// Serialize serializes information the current object
func (m *TeamTemplateDefinition) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetAudience() != nil {
        cast := (*m.GetAudience()).String()
        err = writer.WriteStringValue("audience", &cast)
        if err != nil {
            return err
        }
    }
    if m.GetCategories() != nil {
        err = writer.WriteCollectionOfStringValues("categories", m.GetCategories())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("description", m.GetDescription())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("displayName", m.GetDisplayName())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("iconUrl", m.GetIconUrl())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("languageTag", m.GetLanguageTag())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("lastModifiedBy", m.GetLastModifiedBy())
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
        err = writer.WriteStringValue("parentTemplateId", m.GetParentTemplateId())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("publisherName", m.GetPublisherName())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("shortDescription", m.GetShortDescription())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("teamDefinition", m.GetTeamDefinition())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAudience sets the audience property value. Describes the audience the team template is available to. The possible values are: organization, user, public, unknownFutureValue.
func (m *TeamTemplateDefinition) SetAudience(value *TeamTemplateAudience)() {
    err := m.GetBackingStore().Set("audience", value)
    if err != nil {
        panic(err)
    }
}
// SetCategories sets the categories property value. The assigned categories for the team template.
func (m *TeamTemplateDefinition) SetCategories(value []string)() {
    err := m.GetBackingStore().Set("categories", value)
    if err != nil {
        panic(err)
    }
}
// SetDescription sets the description property value. A brief description of the team template as it will appear to the users in Microsoft Teams.
func (m *TeamTemplateDefinition) SetDescription(value *string)() {
    err := m.GetBackingStore().Set("description", value)
    if err != nil {
        panic(err)
    }
}
// SetDisplayName sets the displayName property value. The user defined name of the team template.
func (m *TeamTemplateDefinition) SetDisplayName(value *string)() {
    err := m.GetBackingStore().Set("displayName", value)
    if err != nil {
        panic(err)
    }
}
// SetIconUrl sets the iconUrl property value. The icon url for the team template.
func (m *TeamTemplateDefinition) SetIconUrl(value *string)() {
    err := m.GetBackingStore().Set("iconUrl", value)
    if err != nil {
        panic(err)
    }
}
// SetLanguageTag sets the languageTag property value. Language the template is available in.
func (m *TeamTemplateDefinition) SetLanguageTag(value *string)() {
    err := m.GetBackingStore().Set("languageTag", value)
    if err != nil {
        panic(err)
    }
}
// SetLastModifiedBy sets the lastModifiedBy property value. The identity of the user who last modified the team template.
func (m *TeamTemplateDefinition) SetLastModifiedBy(value IdentitySetable)() {
    err := m.GetBackingStore().Set("lastModifiedBy", value)
    if err != nil {
        panic(err)
    }
}
// SetLastModifiedDateTime sets the lastModifiedDateTime property value. The date time of when the team template was last modified.
func (m *TeamTemplateDefinition) SetLastModifiedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    err := m.GetBackingStore().Set("lastModifiedDateTime", value)
    if err != nil {
        panic(err)
    }
}
// SetParentTemplateId sets the parentTemplateId property value. The templateId for the team template
func (m *TeamTemplateDefinition) SetParentTemplateId(value *string)() {
    err := m.GetBackingStore().Set("parentTemplateId", value)
    if err != nil {
        panic(err)
    }
}
// SetPublisherName sets the publisherName property value. The organization which published the team template.
func (m *TeamTemplateDefinition) SetPublisherName(value *string)() {
    err := m.GetBackingStore().Set("publisherName", value)
    if err != nil {
        panic(err)
    }
}
// SetShortDescription sets the shortDescription property value. A short-description of the team template as it will appear to the users in Microsoft Teams.
func (m *TeamTemplateDefinition) SetShortDescription(value *string)() {
    err := m.GetBackingStore().Set("shortDescription", value)
    if err != nil {
        panic(err)
    }
}
// SetTeamDefinition sets the teamDefinition property value. Collection of channel objects. A channel represents a topic, and therefore a logical isolation of discussion, within a team.
func (m *TeamTemplateDefinition) SetTeamDefinition(value Teamable)() {
    err := m.GetBackingStore().Set("teamDefinition", value)
    if err != nil {
        panic(err)
    }
}
// TeamTemplateDefinitionable 
type TeamTemplateDefinitionable interface {
    Entityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAudience()(*TeamTemplateAudience)
    GetCategories()([]string)
    GetDescription()(*string)
    GetDisplayName()(*string)
    GetIconUrl()(*string)
    GetLanguageTag()(*string)
    GetLastModifiedBy()(IdentitySetable)
    GetLastModifiedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetParentTemplateId()(*string)
    GetPublisherName()(*string)
    GetShortDescription()(*string)
    GetTeamDefinition()(Teamable)
    SetAudience(value *TeamTemplateAudience)()
    SetCategories(value []string)()
    SetDescription(value *string)()
    SetDisplayName(value *string)()
    SetIconUrl(value *string)()
    SetLanguageTag(value *string)()
    SetLastModifiedBy(value IdentitySetable)()
    SetLastModifiedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetParentTemplateId(value *string)()
    SetPublisherName(value *string)()
    SetShortDescription(value *string)()
    SetTeamDefinition(value Teamable)()
}
