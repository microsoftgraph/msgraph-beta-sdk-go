package graph

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// 
type MobileApp struct {
    Entity
    // The list of group assignments for this mobile app.
    assignments []MobileAppAssignment;
    // The list of categories for this app.
    categories []MobileAppCategory;
    // The date and time the app was created.
    createdDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time;
    // The total number of dependencies the child app has.
    dependentAppCount *int32;
    // The description of the app.
    description *string;
    // The developer of the app.
    developer *string;
    // The list of installation states for this mobile app.
    deviceStatuses []MobileAppInstallStatus;
    // The admin provided or imported title of the app.
    displayName *string;
    // The more information Url.
    informationUrl *string;
    // Mobile App Install Summary.
    installSummary *MobileAppInstallSummary;
    // The value indicating whether the app is assigned to at least one group.
    isAssigned *bool;
    // The value indicating whether the app is marked as featured by the admin.
    isFeatured *bool;
    // The large icon, to be displayed in the app details and used for upload of the icon.
    largeIcon *MimeContent;
    // The date and time the app was last modified.
    lastModifiedDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time;
    // Notes for the app.
    notes *string;
    // The owner of the app.
    owner *string;
    // The privacy statement Url.
    privacyInformationUrl *string;
    // The publisher of the app.
    publisher *string;
    // The publishing state for the app. The app cannot be assigned unless the app is published. Possible values are: notPublished, processing, published.
    publishingState *MobileAppPublishingState;
    // List of relationships for this mobile app.
    relationships []MobileAppRelationship;
    // List of scope tag ids for this mobile app.
    roleScopeTagIds []string;
    // The total number of apps this app is directly or indirectly superseded by.
    supersededAppCount *int32;
    // The total number of apps this app directly or indirectly supersedes.
    supersedingAppCount *int32;
    // The upload state.
    uploadState *int32;
    // The list of installation states for this mobile app.
    userStatuses []UserAppInstallStatus;
}
// Instantiates a new mobileApp and sets the default values.
func NewMobileApp()(*MobileApp) {
    m := &MobileApp{
        Entity: *NewEntity(),
    }
    return m
}
// Gets the assignments property value. The list of group assignments for this mobile app.
func (m *MobileApp) GetAssignments()([]MobileAppAssignment) {
    if m == nil {
        return nil
    } else {
        return m.assignments
    }
}
// Gets the categories property value. The list of categories for this app.
func (m *MobileApp) GetCategories()([]MobileAppCategory) {
    if m == nil {
        return nil
    } else {
        return m.categories
    }
}
// Gets the createdDateTime property value. The date and time the app was created.
func (m *MobileApp) GetCreatedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.createdDateTime
    }
}
// Gets the dependentAppCount property value. The total number of dependencies the child app has.
func (m *MobileApp) GetDependentAppCount()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.dependentAppCount
    }
}
// Gets the description property value. The description of the app.
func (m *MobileApp) GetDescription()(*string) {
    if m == nil {
        return nil
    } else {
        return m.description
    }
}
// Gets the developer property value. The developer of the app.
func (m *MobileApp) GetDeveloper()(*string) {
    if m == nil {
        return nil
    } else {
        return m.developer
    }
}
// Gets the deviceStatuses property value. The list of installation states for this mobile app.
func (m *MobileApp) GetDeviceStatuses()([]MobileAppInstallStatus) {
    if m == nil {
        return nil
    } else {
        return m.deviceStatuses
    }
}
// Gets the displayName property value. The admin provided or imported title of the app.
func (m *MobileApp) GetDisplayName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.displayName
    }
}
// Gets the informationUrl property value. The more information Url.
func (m *MobileApp) GetInformationUrl()(*string) {
    if m == nil {
        return nil
    } else {
        return m.informationUrl
    }
}
// Gets the installSummary property value. Mobile App Install Summary.
func (m *MobileApp) GetInstallSummary()(*MobileAppInstallSummary) {
    if m == nil {
        return nil
    } else {
        return m.installSummary
    }
}
// Gets the isAssigned property value. The value indicating whether the app is assigned to at least one group.
func (m *MobileApp) GetIsAssigned()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.isAssigned
    }
}
// Gets the isFeatured property value. The value indicating whether the app is marked as featured by the admin.
func (m *MobileApp) GetIsFeatured()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.isFeatured
    }
}
// Gets the largeIcon property value. The large icon, to be displayed in the app details and used for upload of the icon.
func (m *MobileApp) GetLargeIcon()(*MimeContent) {
    if m == nil {
        return nil
    } else {
        return m.largeIcon
    }
}
// Gets the lastModifiedDateTime property value. The date and time the app was last modified.
func (m *MobileApp) GetLastModifiedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.lastModifiedDateTime
    }
}
// Gets the notes property value. Notes for the app.
func (m *MobileApp) GetNotes()(*string) {
    if m == nil {
        return nil
    } else {
        return m.notes
    }
}
// Gets the owner property value. The owner of the app.
func (m *MobileApp) GetOwner()(*string) {
    if m == nil {
        return nil
    } else {
        return m.owner
    }
}
// Gets the privacyInformationUrl property value. The privacy statement Url.
func (m *MobileApp) GetPrivacyInformationUrl()(*string) {
    if m == nil {
        return nil
    } else {
        return m.privacyInformationUrl
    }
}
// Gets the publisher property value. The publisher of the app.
func (m *MobileApp) GetPublisher()(*string) {
    if m == nil {
        return nil
    } else {
        return m.publisher
    }
}
// Gets the publishingState property value. The publishing state for the app. The app cannot be assigned unless the app is published. Possible values are: notPublished, processing, published.
func (m *MobileApp) GetPublishingState()(*MobileAppPublishingState) {
    if m == nil {
        return nil
    } else {
        return m.publishingState
    }
}
// Gets the relationships property value. List of relationships for this mobile app.
func (m *MobileApp) GetRelationships()([]MobileAppRelationship) {
    if m == nil {
        return nil
    } else {
        return m.relationships
    }
}
// Gets the roleScopeTagIds property value. List of scope tag ids for this mobile app.
func (m *MobileApp) GetRoleScopeTagIds()([]string) {
    if m == nil {
        return nil
    } else {
        return m.roleScopeTagIds
    }
}
// Gets the supersededAppCount property value. The total number of apps this app is directly or indirectly superseded by.
func (m *MobileApp) GetSupersededAppCount()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.supersededAppCount
    }
}
// Gets the supersedingAppCount property value. The total number of apps this app directly or indirectly supersedes.
func (m *MobileApp) GetSupersedingAppCount()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.supersedingAppCount
    }
}
// Gets the uploadState property value. The upload state.
func (m *MobileApp) GetUploadState()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.uploadState
    }
}
// Gets the userStatuses property value. The list of installation states for this mobile app.
func (m *MobileApp) GetUserStatuses()([]UserAppInstallStatus) {
    if m == nil {
        return nil
    } else {
        return m.userStatuses
    }
}
// The deserialization information for the current model
func (m *MobileApp) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["assignments"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewMobileAppAssignment() })
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]MobileAppAssignment, len(val))
            for i, v := range val {
                res[i] = *(v.(*MobileAppAssignment))
            }
            m.SetAssignments(res)
        }
        return nil
    }
    res["categories"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewMobileAppCategory() })
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]MobileAppCategory, len(val))
            for i, v := range val {
                res[i] = *(v.(*MobileAppCategory))
            }
            m.SetCategories(res)
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
    res["dependentAppCount"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDependentAppCount(val)
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
    res["developer"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDeveloper(val)
        }
        return nil
    }
    res["deviceStatuses"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewMobileAppInstallStatus() })
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]MobileAppInstallStatus, len(val))
            for i, v := range val {
                res[i] = *(v.(*MobileAppInstallStatus))
            }
            m.SetDeviceStatuses(res)
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
    res["informationUrl"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetInformationUrl(val)
        }
        return nil
    }
    res["installSummary"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewMobileAppInstallSummary() })
        if err != nil {
            return err
        }
        if val != nil {
            m.SetInstallSummary(val.(*MobileAppInstallSummary))
        }
        return nil
    }
    res["isAssigned"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsAssigned(val)
        }
        return nil
    }
    res["isFeatured"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsFeatured(val)
        }
        return nil
    }
    res["largeIcon"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewMimeContent() })
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLargeIcon(val.(*MimeContent))
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
    res["notes"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetNotes(val)
        }
        return nil
    }
    res["owner"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOwner(val)
        }
        return nil
    }
    res["privacyInformationUrl"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPrivacyInformationUrl(val)
        }
        return nil
    }
    res["publisher"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPublisher(val)
        }
        return nil
    }
    res["publishingState"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParseMobileAppPublishingState)
        if err != nil {
            return err
        }
        if val != nil {
            cast := val.(MobileAppPublishingState)
            m.SetPublishingState(&cast)
        }
        return nil
    }
    res["relationships"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewMobileAppRelationship() })
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]MobileAppRelationship, len(val))
            for i, v := range val {
                res[i] = *(v.(*MobileAppRelationship))
            }
            m.SetRelationships(res)
        }
        return nil
    }
    res["roleScopeTagIds"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]string, len(val))
            for i, v := range val {
                res[i] = *(v.(*string))
            }
            m.SetRoleScopeTagIds(res)
        }
        return nil
    }
    res["supersededAppCount"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSupersededAppCount(val)
        }
        return nil
    }
    res["supersedingAppCount"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSupersedingAppCount(val)
        }
        return nil
    }
    res["uploadState"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUploadState(val)
        }
        return nil
    }
    res["userStatuses"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewUserAppInstallStatus() })
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]UserAppInstallStatus, len(val))
            for i, v := range val {
                res[i] = *(v.(*UserAppInstallStatus))
            }
            m.SetUserStatuses(res)
        }
        return nil
    }
    return res
}
func (m *MobileApp) IsNil()(bool) {
    return m == nil
}
// Serializes information the current object
// Parameters:
//  - writer : Serialization writer to use to serialize this model
func (m *MobileApp) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetAssignments()))
        for i, v := range m.GetAssignments() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("assignments", cast)
        if err != nil {
            return err
        }
    }
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetCategories()))
        for i, v := range m.GetCategories() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("categories", cast)
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
        err = writer.WriteInt32Value("dependentAppCount", m.GetDependentAppCount())
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
        err = writer.WriteStringValue("developer", m.GetDeveloper())
        if err != nil {
            return err
        }
    }
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetDeviceStatuses()))
        for i, v := range m.GetDeviceStatuses() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("deviceStatuses", cast)
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
        err = writer.WriteStringValue("informationUrl", m.GetInformationUrl())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("installSummary", m.GetInstallSummary())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("isAssigned", m.GetIsAssigned())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("isFeatured", m.GetIsFeatured())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("largeIcon", m.GetLargeIcon())
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
        err = writer.WriteStringValue("notes", m.GetNotes())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("owner", m.GetOwner())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("privacyInformationUrl", m.GetPrivacyInformationUrl())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("publisher", m.GetPublisher())
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
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetRelationships()))
        for i, v := range m.GetRelationships() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("relationships", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteCollectionOfStringValues("roleScopeTagIds", m.GetRoleScopeTagIds())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("supersededAppCount", m.GetSupersededAppCount())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("supersedingAppCount", m.GetSupersedingAppCount())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("uploadState", m.GetUploadState())
        if err != nil {
            return err
        }
    }
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetUserStatuses()))
        for i, v := range m.GetUserStatuses() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("userStatuses", cast)
        if err != nil {
            return err
        }
    }
    return nil
}
// Sets the assignments property value. The list of group assignments for this mobile app.
// Parameters:
//  - value : Value to set for the assignments property.
func (m *MobileApp) SetAssignments(value []MobileAppAssignment)() {
    m.assignments = value
}
// Sets the categories property value. The list of categories for this app.
// Parameters:
//  - value : Value to set for the categories property.
func (m *MobileApp) SetCategories(value []MobileAppCategory)() {
    m.categories = value
}
// Sets the createdDateTime property value. The date and time the app was created.
// Parameters:
//  - value : Value to set for the createdDateTime property.
func (m *MobileApp) SetCreatedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.createdDateTime = value
}
// Sets the dependentAppCount property value. The total number of dependencies the child app has.
// Parameters:
//  - value : Value to set for the dependentAppCount property.
func (m *MobileApp) SetDependentAppCount(value *int32)() {
    m.dependentAppCount = value
}
// Sets the description property value. The description of the app.
// Parameters:
//  - value : Value to set for the description property.
func (m *MobileApp) SetDescription(value *string)() {
    m.description = value
}
// Sets the developer property value. The developer of the app.
// Parameters:
//  - value : Value to set for the developer property.
func (m *MobileApp) SetDeveloper(value *string)() {
    m.developer = value
}
// Sets the deviceStatuses property value. The list of installation states for this mobile app.
// Parameters:
//  - value : Value to set for the deviceStatuses property.
func (m *MobileApp) SetDeviceStatuses(value []MobileAppInstallStatus)() {
    m.deviceStatuses = value
}
// Sets the displayName property value. The admin provided or imported title of the app.
// Parameters:
//  - value : Value to set for the displayName property.
func (m *MobileApp) SetDisplayName(value *string)() {
    m.displayName = value
}
// Sets the informationUrl property value. The more information Url.
// Parameters:
//  - value : Value to set for the informationUrl property.
func (m *MobileApp) SetInformationUrl(value *string)() {
    m.informationUrl = value
}
// Sets the installSummary property value. Mobile App Install Summary.
// Parameters:
//  - value : Value to set for the installSummary property.
func (m *MobileApp) SetInstallSummary(value *MobileAppInstallSummary)() {
    m.installSummary = value
}
// Sets the isAssigned property value. The value indicating whether the app is assigned to at least one group.
// Parameters:
//  - value : Value to set for the isAssigned property.
func (m *MobileApp) SetIsAssigned(value *bool)() {
    m.isAssigned = value
}
// Sets the isFeatured property value. The value indicating whether the app is marked as featured by the admin.
// Parameters:
//  - value : Value to set for the isFeatured property.
func (m *MobileApp) SetIsFeatured(value *bool)() {
    m.isFeatured = value
}
// Sets the largeIcon property value. The large icon, to be displayed in the app details and used for upload of the icon.
// Parameters:
//  - value : Value to set for the largeIcon property.
func (m *MobileApp) SetLargeIcon(value *MimeContent)() {
    m.largeIcon = value
}
// Sets the lastModifiedDateTime property value. The date and time the app was last modified.
// Parameters:
//  - value : Value to set for the lastModifiedDateTime property.
func (m *MobileApp) SetLastModifiedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.lastModifiedDateTime = value
}
// Sets the notes property value. Notes for the app.
// Parameters:
//  - value : Value to set for the notes property.
func (m *MobileApp) SetNotes(value *string)() {
    m.notes = value
}
// Sets the owner property value. The owner of the app.
// Parameters:
//  - value : Value to set for the owner property.
func (m *MobileApp) SetOwner(value *string)() {
    m.owner = value
}
// Sets the privacyInformationUrl property value. The privacy statement Url.
// Parameters:
//  - value : Value to set for the privacyInformationUrl property.
func (m *MobileApp) SetPrivacyInformationUrl(value *string)() {
    m.privacyInformationUrl = value
}
// Sets the publisher property value. The publisher of the app.
// Parameters:
//  - value : Value to set for the publisher property.
func (m *MobileApp) SetPublisher(value *string)() {
    m.publisher = value
}
// Sets the publishingState property value. The publishing state for the app. The app cannot be assigned unless the app is published. Possible values are: notPublished, processing, published.
// Parameters:
//  - value : Value to set for the publishingState property.
func (m *MobileApp) SetPublishingState(value *MobileAppPublishingState)() {
    m.publishingState = value
}
// Sets the relationships property value. List of relationships for this mobile app.
// Parameters:
//  - value : Value to set for the relationships property.
func (m *MobileApp) SetRelationships(value []MobileAppRelationship)() {
    m.relationships = value
}
// Sets the roleScopeTagIds property value. List of scope tag ids for this mobile app.
// Parameters:
//  - value : Value to set for the roleScopeTagIds property.
func (m *MobileApp) SetRoleScopeTagIds(value []string)() {
    m.roleScopeTagIds = value
}
// Sets the supersededAppCount property value. The total number of apps this app is directly or indirectly superseded by.
// Parameters:
//  - value : Value to set for the supersededAppCount property.
func (m *MobileApp) SetSupersededAppCount(value *int32)() {
    m.supersededAppCount = value
}
// Sets the supersedingAppCount property value. The total number of apps this app directly or indirectly supersedes.
// Parameters:
//  - value : Value to set for the supersedingAppCount property.
func (m *MobileApp) SetSupersedingAppCount(value *int32)() {
    m.supersedingAppCount = value
}
// Sets the uploadState property value. The upload state.
// Parameters:
//  - value : Value to set for the uploadState property.
func (m *MobileApp) SetUploadState(value *int32)() {
    m.uploadState = value
}
// Sets the userStatuses property value. The list of installation states for this mobile app.
// Parameters:
//  - value : Value to set for the userStatuses property.
func (m *MobileApp) SetUserStatuses(value []UserAppInstallStatus)() {
    m.userStatuses = value
}
