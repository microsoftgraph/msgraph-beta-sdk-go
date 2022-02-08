package graph

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// TeamsAppDefinition 
type TeamsAppDefinition struct {
    Entity
    // A collection of scopes where the Teams app can be installed. Possible values are:team — Indicates that the Teams app can be installed within a team and is authorized to access that team's data. groupChat  — Indicates that the Teams app can be installed within a group chat and is authorized to access that group chat's data.  personal — Indicates that the Teams app can be installed in the personal scope of a user and is authorized to access that user's data.
    allowedInstallationScopes *TeamsAppInstallationScopes;
    // The WebApplicationInfo.Id from the Teams app manifest.
    azureADAppId *string;
    // The details of the bot specified in the Teams app manifest.
    bot *TeamworkBot;
    // The color version of the Teams app's icon.
    colorIcon *TeamsAppIcon;
    // 
    createdBy *IdentitySet;
    // Verbose description of the application.
    description *string;
    // The name of the app provided by the app developer.
    displayName *string;
    // 
    lastModifiedDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time;
    // The outline version of the Teams app's icon.
    outlineIcon *TeamsAppIcon;
    // The published status of a specific version of a Teams app. Possible values are:submitted — The specific version of the Teams app has been submitted and is under review. published  — The request to publish the specific version of the Teams app has been approved by the admin and the app is published.  rejected — The request to publish the specific version of the Teams app was rejected by the admin.
    publishingState *TeamsAppPublishingState;
    // 
    shortdescription *string;
    // The ID from the Teams app manifest.
    teamsAppId *string;
    // The version number of the application.
    version *string;
}
// NewTeamsAppDefinition instantiates a new teamsAppDefinition and sets the default values.
func NewTeamsAppDefinition()(*TeamsAppDefinition) {
    m := &TeamsAppDefinition{
        Entity: *NewEntity(),
    }
    return m
}
// GetAllowedInstallationScopes gets the allowedInstallationScopes property value. A collection of scopes where the Teams app can be installed. Possible values are:team — Indicates that the Teams app can be installed within a team and is authorized to access that team's data. groupChat  — Indicates that the Teams app can be installed within a group chat and is authorized to access that group chat's data.  personal — Indicates that the Teams app can be installed in the personal scope of a user and is authorized to access that user's data.
func (m *TeamsAppDefinition) GetAllowedInstallationScopes()(*TeamsAppInstallationScopes) {
    if m == nil {
        return nil
    } else {
        return m.allowedInstallationScopes
    }
}
// GetAzureADAppId gets the azureADAppId property value. The WebApplicationInfo.Id from the Teams app manifest.
func (m *TeamsAppDefinition) GetAzureADAppId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.azureADAppId
    }
}
// GetBot gets the bot property value. The details of the bot specified in the Teams app manifest.
func (m *TeamsAppDefinition) GetBot()(*TeamworkBot) {
    if m == nil {
        return nil
    } else {
        return m.bot
    }
}
// GetColorIcon gets the colorIcon property value. The color version of the Teams app's icon.
func (m *TeamsAppDefinition) GetColorIcon()(*TeamsAppIcon) {
    if m == nil {
        return nil
    } else {
        return m.colorIcon
    }
}
// GetCreatedBy gets the createdBy property value. 
func (m *TeamsAppDefinition) GetCreatedBy()(*IdentitySet) {
    if m == nil {
        return nil
    } else {
        return m.createdBy
    }
}
// GetDescription gets the description property value. Verbose description of the application.
func (m *TeamsAppDefinition) GetDescription()(*string) {
    if m == nil {
        return nil
    } else {
        return m.description
    }
}
// GetDisplayName gets the displayName property value. The name of the app provided by the app developer.
func (m *TeamsAppDefinition) GetDisplayName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.displayName
    }
}
// GetLastModifiedDateTime gets the lastModifiedDateTime property value. 
func (m *TeamsAppDefinition) GetLastModifiedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.lastModifiedDateTime
    }
}
// GetOutlineIcon gets the outlineIcon property value. The outline version of the Teams app's icon.
func (m *TeamsAppDefinition) GetOutlineIcon()(*TeamsAppIcon) {
    if m == nil {
        return nil
    } else {
        return m.outlineIcon
    }
}
// GetPublishingState gets the publishingState property value. The published status of a specific version of a Teams app. Possible values are:submitted — The specific version of the Teams app has been submitted and is under review. published  — The request to publish the specific version of the Teams app has been approved by the admin and the app is published.  rejected — The request to publish the specific version of the Teams app was rejected by the admin.
func (m *TeamsAppDefinition) GetPublishingState()(*TeamsAppPublishingState) {
    if m == nil {
        return nil
    } else {
        return m.publishingState
    }
}
// GetShortdescription gets the shortdescription property value. 
func (m *TeamsAppDefinition) GetShortdescription()(*string) {
    if m == nil {
        return nil
    } else {
        return m.shortdescription
    }
}
// GetTeamsAppId gets the teamsAppId property value. The ID from the Teams app manifest.
func (m *TeamsAppDefinition) GetTeamsAppId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.teamsAppId
    }
}
// GetVersion gets the version property value. The version number of the application.
func (m *TeamsAppDefinition) GetVersion()(*string) {
    if m == nil {
        return nil
    } else {
        return m.version
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *TeamsAppDefinition) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["allowedInstallationScopes"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParseTeamsAppInstallationScopes)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAllowedInstallationScopes(val.(*TeamsAppInstallationScopes))
        }
        return nil
    }
    res["azureADAppId"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAzureADAppId(val)
        }
        return nil
    }
    res["bot"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewTeamworkBot() })
        if err != nil {
            return err
        }
        if val != nil {
            m.SetBot(val.(*TeamworkBot))
        }
        return nil
    }
    res["colorIcon"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewTeamsAppIcon() })
        if err != nil {
            return err
        }
        if val != nil {
            m.SetColorIcon(val.(*TeamsAppIcon))
        }
        return nil
    }
    res["createdBy"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewIdentitySet() })
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCreatedBy(val.(*IdentitySet))
        }
        return nil
    }
    res["description"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDescription(val)
        }
        return nil
    }
    res["displayName"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDisplayName(val)
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
    res["outlineIcon"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewTeamsAppIcon() })
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOutlineIcon(val.(*TeamsAppIcon))
        }
        return nil
    }
    res["publishingState"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParseTeamsAppPublishingState)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPublishingState(val.(*TeamsAppPublishingState))
        }
        return nil
    }
    res["shortdescription"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetShortdescription(val)
        }
        return nil
    }
    res["teamsAppId"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTeamsAppId(val)
        }
        return nil
    }
    res["version"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetVersion(val)
        }
        return nil
    }
    return res
}
func (m *TeamsAppDefinition) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
func (m *TeamsAppDefinition) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetAllowedInstallationScopes() != nil {
        cast := (*m.GetAllowedInstallationScopes()).String()
        err = writer.WriteStringValue("allowedInstallationScopes", &cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("azureADAppId", m.GetAzureADAppId())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("bot", m.GetBot())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("colorIcon", m.GetColorIcon())
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
        err = writer.WriteTimeValue("lastModifiedDateTime", m.GetLastModifiedDateTime())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("outlineIcon", m.GetOutlineIcon())
        if err != nil {
            return err
        }
    }
    if m.GetPublishingState() != nil {
        cast := (*m.GetPublishingState()).String()
        err = writer.WriteStringValue("publishingState", &cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("shortdescription", m.GetShortdescription())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("teamsAppId", m.GetTeamsAppId())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("version", m.GetVersion())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAllowedInstallationScopes sets the allowedInstallationScopes property value. A collection of scopes where the Teams app can be installed. Possible values are:team — Indicates that the Teams app can be installed within a team and is authorized to access that team's data. groupChat  — Indicates that the Teams app can be installed within a group chat and is authorized to access that group chat's data.  personal — Indicates that the Teams app can be installed in the personal scope of a user and is authorized to access that user's data.
func (m *TeamsAppDefinition) SetAllowedInstallationScopes(value *TeamsAppInstallationScopes)() {
    if m != nil {
        m.allowedInstallationScopes = value
    }
}
// SetAzureADAppId sets the azureADAppId property value. The WebApplicationInfo.Id from the Teams app manifest.
func (m *TeamsAppDefinition) SetAzureADAppId(value *string)() {
    if m != nil {
        m.azureADAppId = value
    }
}
// SetBot sets the bot property value. The details of the bot specified in the Teams app manifest.
func (m *TeamsAppDefinition) SetBot(value *TeamworkBot)() {
    if m != nil {
        m.bot = value
    }
}
// SetColorIcon sets the colorIcon property value. The color version of the Teams app's icon.
func (m *TeamsAppDefinition) SetColorIcon(value *TeamsAppIcon)() {
    if m != nil {
        m.colorIcon = value
    }
}
// SetCreatedBy sets the createdBy property value. 
func (m *TeamsAppDefinition) SetCreatedBy(value *IdentitySet)() {
    if m != nil {
        m.createdBy = value
    }
}
// SetDescription sets the description property value. Verbose description of the application.
func (m *TeamsAppDefinition) SetDescription(value *string)() {
    if m != nil {
        m.description = value
    }
}
// SetDisplayName sets the displayName property value. The name of the app provided by the app developer.
func (m *TeamsAppDefinition) SetDisplayName(value *string)() {
    if m != nil {
        m.displayName = value
    }
}
// SetLastModifiedDateTime sets the lastModifiedDateTime property value. 
func (m *TeamsAppDefinition) SetLastModifiedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    if m != nil {
        m.lastModifiedDateTime = value
    }
}
// SetOutlineIcon sets the outlineIcon property value. The outline version of the Teams app's icon.
func (m *TeamsAppDefinition) SetOutlineIcon(value *TeamsAppIcon)() {
    if m != nil {
        m.outlineIcon = value
    }
}
// SetPublishingState sets the publishingState property value. The published status of a specific version of a Teams app. Possible values are:submitted — The specific version of the Teams app has been submitted and is under review. published  — The request to publish the specific version of the Teams app has been approved by the admin and the app is published.  rejected — The request to publish the specific version of the Teams app was rejected by the admin.
func (m *TeamsAppDefinition) SetPublishingState(value *TeamsAppPublishingState)() {
    if m != nil {
        m.publishingState = value
    }
}
// SetShortdescription sets the shortdescription property value. 
func (m *TeamsAppDefinition) SetShortdescription(value *string)() {
    if m != nil {
        m.shortdescription = value
    }
}
// SetTeamsAppId sets the teamsAppId property value. The ID from the Teams app manifest.
func (m *TeamsAppDefinition) SetTeamsAppId(value *string)() {
    if m != nil {
        m.teamsAppId = value
    }
}
// SetVersion sets the version property value. The version number of the application.
func (m *TeamsAppDefinition) SetVersion(value *string)() {
    if m != nil {
        m.version = value
    }
}
