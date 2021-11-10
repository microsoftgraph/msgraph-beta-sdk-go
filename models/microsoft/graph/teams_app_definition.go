package graph

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// 
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
// Instantiates a new teamsAppDefinition and sets the default values.
func NewTeamsAppDefinition()(*TeamsAppDefinition) {
    m := &TeamsAppDefinition{
        Entity: *NewEntity(),
    }
    return m
}
// Gets the allowedInstallationScopes property value. A collection of scopes where the Teams app can be installed. Possible values are:team — Indicates that the Teams app can be installed within a team and is authorized to access that team's data. groupChat  — Indicates that the Teams app can be installed within a group chat and is authorized to access that group chat's data.  personal — Indicates that the Teams app can be installed in the personal scope of a user and is authorized to access that user's data.
func (m *TeamsAppDefinition) GetAllowedInstallationScopes()(*TeamsAppInstallationScopes) {
    if m == nil {
        return nil
    } else {
        return m.allowedInstallationScopes
    }
}
// Gets the azureADAppId property value. The WebApplicationInfo.Id from the Teams app manifest.
func (m *TeamsAppDefinition) GetAzureADAppId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.azureADAppId
    }
}
// Gets the bot property value. The details of the bot specified in the Teams app manifest.
func (m *TeamsAppDefinition) GetBot()(*TeamworkBot) {
    if m == nil {
        return nil
    } else {
        return m.bot
    }
}
// Gets the colorIcon property value. The color version of the Teams app's icon.
func (m *TeamsAppDefinition) GetColorIcon()(*TeamsAppIcon) {
    if m == nil {
        return nil
    } else {
        return m.colorIcon
    }
}
// Gets the createdBy property value. 
func (m *TeamsAppDefinition) GetCreatedBy()(*IdentitySet) {
    if m == nil {
        return nil
    } else {
        return m.createdBy
    }
}
// Gets the description property value. Verbose description of the application.
func (m *TeamsAppDefinition) GetDescription()(*string) {
    if m == nil {
        return nil
    } else {
        return m.description
    }
}
// Gets the displayName property value. The name of the app provided by the app developer.
func (m *TeamsAppDefinition) GetDisplayName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.displayName
    }
}
// Gets the lastModifiedDateTime property value. 
func (m *TeamsAppDefinition) GetLastModifiedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.lastModifiedDateTime
    }
}
// Gets the outlineIcon property value. The outline version of the Teams app's icon.
func (m *TeamsAppDefinition) GetOutlineIcon()(*TeamsAppIcon) {
    if m == nil {
        return nil
    } else {
        return m.outlineIcon
    }
}
// Gets the publishingState property value. The published status of a specific version of a Teams app. Possible values are:submitted — The specific version of the Teams app has been submitted and is under review. published  — The request to publish the specific version of the Teams app has been approved by the admin and the app is published.  rejected — The request to publish the specific version of the Teams app was rejected by the admin.
func (m *TeamsAppDefinition) GetPublishingState()(*TeamsAppPublishingState) {
    if m == nil {
        return nil
    } else {
        return m.publishingState
    }
}
// Gets the shortdescription property value. 
func (m *TeamsAppDefinition) GetShortdescription()(*string) {
    if m == nil {
        return nil
    } else {
        return m.shortdescription
    }
}
// Gets the teamsAppId property value. The ID from the Teams app manifest.
func (m *TeamsAppDefinition) GetTeamsAppId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.teamsAppId
    }
}
// Gets the version property value. The version number of the application.
func (m *TeamsAppDefinition) GetVersion()(*string) {
    if m == nil {
        return nil
    } else {
        return m.version
    }
}
// The deserialization information for the current model
func (m *TeamsAppDefinition) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["allowedInstallationScopes"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParseTeamsAppInstallationScopes)
        if err != nil {
            return err
        }
        if val != nil {
            cast := val.(TeamsAppInstallationScopes)
            m.SetAllowedInstallationScopes(&cast)
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
            cast := val.(TeamsAppPublishingState)
            m.SetPublishingState(&cast)
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
// Serializes information the current object
// Parameters:
//  - writer : Serialization writer to use to serialize this model
func (m *TeamsAppDefinition) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetAllowedInstallationScopes() != nil {
        cast := m.GetAllowedInstallationScopes().String()
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
        cast := m.GetPublishingState().String()
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
// Sets the allowedInstallationScopes property value. A collection of scopes where the Teams app can be installed. Possible values are:team — Indicates that the Teams app can be installed within a team and is authorized to access that team's data. groupChat  — Indicates that the Teams app can be installed within a group chat and is authorized to access that group chat's data.  personal — Indicates that the Teams app can be installed in the personal scope of a user and is authorized to access that user's data.
// Parameters:
//  - value : Value to set for the allowedInstallationScopes property.
func (m *TeamsAppDefinition) SetAllowedInstallationScopes(value *TeamsAppInstallationScopes)() {
    m.allowedInstallationScopes = value
}
// Sets the azureADAppId property value. The WebApplicationInfo.Id from the Teams app manifest.
// Parameters:
//  - value : Value to set for the azureADAppId property.
func (m *TeamsAppDefinition) SetAzureADAppId(value *string)() {
    m.azureADAppId = value
}
// Sets the bot property value. The details of the bot specified in the Teams app manifest.
// Parameters:
//  - value : Value to set for the bot property.
func (m *TeamsAppDefinition) SetBot(value *TeamworkBot)() {
    m.bot = value
}
// Sets the colorIcon property value. The color version of the Teams app's icon.
// Parameters:
//  - value : Value to set for the colorIcon property.
func (m *TeamsAppDefinition) SetColorIcon(value *TeamsAppIcon)() {
    m.colorIcon = value
}
// Sets the createdBy property value. 
// Parameters:
//  - value : Value to set for the createdBy property.
func (m *TeamsAppDefinition) SetCreatedBy(value *IdentitySet)() {
    m.createdBy = value
}
// Sets the description property value. Verbose description of the application.
// Parameters:
//  - value : Value to set for the description property.
func (m *TeamsAppDefinition) SetDescription(value *string)() {
    m.description = value
}
// Sets the displayName property value. The name of the app provided by the app developer.
// Parameters:
//  - value : Value to set for the displayName property.
func (m *TeamsAppDefinition) SetDisplayName(value *string)() {
    m.displayName = value
}
// Sets the lastModifiedDateTime property value. 
// Parameters:
//  - value : Value to set for the lastModifiedDateTime property.
func (m *TeamsAppDefinition) SetLastModifiedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.lastModifiedDateTime = value
}
// Sets the outlineIcon property value. The outline version of the Teams app's icon.
// Parameters:
//  - value : Value to set for the outlineIcon property.
func (m *TeamsAppDefinition) SetOutlineIcon(value *TeamsAppIcon)() {
    m.outlineIcon = value
}
// Sets the publishingState property value. The published status of a specific version of a Teams app. Possible values are:submitted — The specific version of the Teams app has been submitted and is under review. published  — The request to publish the specific version of the Teams app has been approved by the admin and the app is published.  rejected — The request to publish the specific version of the Teams app was rejected by the admin.
// Parameters:
//  - value : Value to set for the publishingState property.
func (m *TeamsAppDefinition) SetPublishingState(value *TeamsAppPublishingState)() {
    m.publishingState = value
}
// Sets the shortdescription property value. 
// Parameters:
//  - value : Value to set for the shortdescription property.
func (m *TeamsAppDefinition) SetShortdescription(value *string)() {
    m.shortdescription = value
}
// Sets the teamsAppId property value. The ID from the Teams app manifest.
// Parameters:
//  - value : Value to set for the teamsAppId property.
func (m *TeamsAppDefinition) SetTeamsAppId(value *string)() {
    m.teamsAppId = value
}
// Sets the version property value. The version number of the application.
// Parameters:
//  - value : Value to set for the version property.
func (m *TeamsAppDefinition) SetVersion(value *string)() {
    m.version = value
}
