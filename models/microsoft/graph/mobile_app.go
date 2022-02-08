package graph

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// MobileApp 
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
// NewMobileApp instantiates a new mobileApp and sets the default values.
func NewMobileApp()(*MobileApp) {
    m := &MobileApp{
        Entity: *NewEntity(),
    }
    return m
}
// GetAssignments gets the assignments property value. The list of group assignments for this mobile app.
func (m *MobileApp) GetAssignments()([]MobileAppAssignment) {
    if m == nil {
        return nil
    } else {
        return m.assignments
    }
}
// GetCategories gets the categories property value. The list of categories for this app.
func (m *MobileApp) GetCategories()([]MobileAppCategory) {
    if m == nil {
        return nil
    } else {
        return m.categories
    }
}
// GetCreatedDateTime gets the createdDateTime property value. The date and time the app was created.
func (m *MobileApp) GetCreatedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.createdDateTime
    }
}
// GetDependentAppCount gets the dependentAppCount property value. The total number of dependencies the child app has.
func (m *MobileApp) GetDependentAppCount()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.dependentAppCount
    }
}
// GetDescription gets the description property value. The description of the app.
func (m *MobileApp) GetDescription()(*string) {
    if m == nil {
        return nil
    } else {
        return m.description
    }
}
// GetDeveloper gets the developer property value. The developer of the app.
func (m *MobileApp) GetDeveloper()(*string) {
    if m == nil {
        return nil
    } else {
        return m.developer
    }
}
// GetDeviceStatuses gets the deviceStatuses property value. The list of installation states for this mobile app.
func (m *MobileApp) GetDeviceStatuses()([]MobileAppInstallStatus) {
    if m == nil {
        return nil
    } else {
        return m.deviceStatuses
    }
}
// GetDisplayName gets the displayName property value. The admin provided or imported title of the app.
func (m *MobileApp) GetDisplayName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.displayName
    }
}
// GetInformationUrl gets the informationUrl property value. The more information Url.
func (m *MobileApp) GetInformationUrl()(*string) {
    if m == nil {
        return nil
    } else {
        return m.informationUrl
    }
}
// GetInstallSummary gets the installSummary property value. Mobile App Install Summary.
func (m *MobileApp) GetInstallSummary()(*MobileAppInstallSummary) {
    if m == nil {
        return nil
    } else {
        return m.installSummary
    }
}
// GetIsAssigned gets the isAssigned property value. The value indicating whether the app is assigned to at least one group.
func (m *MobileApp) GetIsAssigned()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.isAssigned
    }
}
// GetIsFeatured gets the isFeatured property value. The value indicating whether the app is marked as featured by the admin.
func (m *MobileApp) GetIsFeatured()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.isFeatured
    }
}
// GetLargeIcon gets the largeIcon property value. The large icon, to be displayed in the app details and used for upload of the icon.
func (m *MobileApp) GetLargeIcon()(*MimeContent) {
    if m == nil {
        return nil
    } else {
        return m.largeIcon
    }
}
// GetLastModifiedDateTime gets the lastModifiedDateTime property value. The date and time the app was last modified.
func (m *MobileApp) GetLastModifiedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.lastModifiedDateTime
    }
}
// GetNotes gets the notes property value. Notes for the app.
func (m *MobileApp) GetNotes()(*string) {
    if m == nil {
        return nil
    } else {
        return m.notes
    }
}
// GetOwner gets the owner property value. The owner of the app.
func (m *MobileApp) GetOwner()(*string) {
    if m == nil {
        return nil
    } else {
        return m.owner
    }
}
// GetPrivacyInformationUrl gets the privacyInformationUrl property value. The privacy statement Url.
func (m *MobileApp) GetPrivacyInformationUrl()(*string) {
    if m == nil {
        return nil
    } else {
        return m.privacyInformationUrl
    }
}
// GetPublisher gets the publisher property value. The publisher of the app.
func (m *MobileApp) GetPublisher()(*string) {
    if m == nil {
        return nil
    } else {
        return m.publisher
    }
}
// GetPublishingState gets the publishingState property value. The publishing state for the app. The app cannot be assigned unless the app is published. Possible values are: notPublished, processing, published.
func (m *MobileApp) GetPublishingState()(*MobileAppPublishingState) {
    if m == nil {
        return nil
    } else {
        return m.publishingState
    }
}
// GetRelationships gets the relationships property value. List of relationships for this mobile app.
func (m *MobileApp) GetRelationships()([]MobileAppRelationship) {
    if m == nil {
        return nil
    } else {
        return m.relationships
    }
}
// GetRoleScopeTagIds gets the roleScopeTagIds property value. List of scope tag ids for this mobile app.
func (m *MobileApp) GetRoleScopeTagIds()([]string) {
    if m == nil {
        return nil
    } else {
        return m.roleScopeTagIds
    }
}
// GetSupersededAppCount gets the supersededAppCount property value. The total number of apps this app is directly or indirectly superseded by.
func (m *MobileApp) GetSupersededAppCount()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.supersededAppCount
    }
}
// GetSupersedingAppCount gets the supersedingAppCount property value. The total number of apps this app directly or indirectly supersedes.
func (m *MobileApp) GetSupersedingAppCount()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.supersedingAppCount
    }
}
// GetUploadState gets the uploadState property value. The upload state.
func (m *MobileApp) GetUploadState()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.uploadState
    }
}
// GetUserStatuses gets the userStatuses property value. The list of installation states for this mobile app.
func (m *MobileApp) GetUserStatuses()([]UserAppInstallStatus) {
    if m == nil {
        return nil
    } else {
        return m.userStatuses
    }
}
// GetFieldDeserializers the deserialization information for the current model
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
            m.SetPublishingState(val.(*MobileAppPublishingState))
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
// Serialize serializes information the current object
func (m *MobileApp) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetAssignments() != nil {
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
    if m.GetCategories() != nil {
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
    if m.GetDeviceStatuses() != nil {
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
        cast := (*m.GetPublishingState()).String()
        err = writer.WriteStringValue("publishingState", &cast)
        if err != nil {
            return err
        }
    }
    if m.GetRelationships() != nil {
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
    if m.GetRoleScopeTagIds() != nil {
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
    if m.GetUserStatuses() != nil {
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
// SetAssignments sets the assignments property value. The list of group assignments for this mobile app.
func (m *MobileApp) SetAssignments(value []MobileAppAssignment)() {
    if m != nil {
        m.assignments = value
    }
}
// SetCategories sets the categories property value. The list of categories for this app.
func (m *MobileApp) SetCategories(value []MobileAppCategory)() {
    if m != nil {
        m.categories = value
    }
}
// SetCreatedDateTime sets the createdDateTime property value. The date and time the app was created.
func (m *MobileApp) SetCreatedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    if m != nil {
        m.createdDateTime = value
    }
}
// SetDependentAppCount sets the dependentAppCount property value. The total number of dependencies the child app has.
func (m *MobileApp) SetDependentAppCount(value *int32)() {
    if m != nil {
        m.dependentAppCount = value
    }
}
// SetDescription sets the description property value. The description of the app.
func (m *MobileApp) SetDescription(value *string)() {
    if m != nil {
        m.description = value
    }
}
// SetDeveloper sets the developer property value. The developer of the app.
func (m *MobileApp) SetDeveloper(value *string)() {
    if m != nil {
        m.developer = value
    }
}
// SetDeviceStatuses sets the deviceStatuses property value. The list of installation states for this mobile app.
func (m *MobileApp) SetDeviceStatuses(value []MobileAppInstallStatus)() {
    if m != nil {
        m.deviceStatuses = value
    }
}
// SetDisplayName sets the displayName property value. The admin provided or imported title of the app.
func (m *MobileApp) SetDisplayName(value *string)() {
    if m != nil {
        m.displayName = value
    }
}
// SetInformationUrl sets the informationUrl property value. The more information Url.
func (m *MobileApp) SetInformationUrl(value *string)() {
    if m != nil {
        m.informationUrl = value
    }
}
// SetInstallSummary sets the installSummary property value. Mobile App Install Summary.
func (m *MobileApp) SetInstallSummary(value *MobileAppInstallSummary)() {
    if m != nil {
        m.installSummary = value
    }
}
// SetIsAssigned sets the isAssigned property value. The value indicating whether the app is assigned to at least one group.
func (m *MobileApp) SetIsAssigned(value *bool)() {
    if m != nil {
        m.isAssigned = value
    }
}
// SetIsFeatured sets the isFeatured property value. The value indicating whether the app is marked as featured by the admin.
func (m *MobileApp) SetIsFeatured(value *bool)() {
    if m != nil {
        m.isFeatured = value
    }
}
// SetLargeIcon sets the largeIcon property value. The large icon, to be displayed in the app details and used for upload of the icon.
func (m *MobileApp) SetLargeIcon(value *MimeContent)() {
    if m != nil {
        m.largeIcon = value
    }
}
// SetLastModifiedDateTime sets the lastModifiedDateTime property value. The date and time the app was last modified.
func (m *MobileApp) SetLastModifiedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    if m != nil {
        m.lastModifiedDateTime = value
    }
}
// SetNotes sets the notes property value. Notes for the app.
func (m *MobileApp) SetNotes(value *string)() {
    if m != nil {
        m.notes = value
    }
}
// SetOwner sets the owner property value. The owner of the app.
func (m *MobileApp) SetOwner(value *string)() {
    if m != nil {
        m.owner = value
    }
}
// SetPrivacyInformationUrl sets the privacyInformationUrl property value. The privacy statement Url.
func (m *MobileApp) SetPrivacyInformationUrl(value *string)() {
    if m != nil {
        m.privacyInformationUrl = value
    }
}
// SetPublisher sets the publisher property value. The publisher of the app.
func (m *MobileApp) SetPublisher(value *string)() {
    if m != nil {
        m.publisher = value
    }
}
// SetPublishingState sets the publishingState property value. The publishing state for the app. The app cannot be assigned unless the app is published. Possible values are: notPublished, processing, published.
func (m *MobileApp) SetPublishingState(value *MobileAppPublishingState)() {
    if m != nil {
        m.publishingState = value
    }
}
// SetRelationships sets the relationships property value. List of relationships for this mobile app.
func (m *MobileApp) SetRelationships(value []MobileAppRelationship)() {
    if m != nil {
        m.relationships = value
    }
}
// SetRoleScopeTagIds sets the roleScopeTagIds property value. List of scope tag ids for this mobile app.
func (m *MobileApp) SetRoleScopeTagIds(value []string)() {
    if m != nil {
        m.roleScopeTagIds = value
    }
}
// SetSupersededAppCount sets the supersededAppCount property value. The total number of apps this app is directly or indirectly superseded by.
func (m *MobileApp) SetSupersededAppCount(value *int32)() {
    if m != nil {
        m.supersededAppCount = value
    }
}
// SetSupersedingAppCount sets the supersedingAppCount property value. The total number of apps this app directly or indirectly supersedes.
func (m *MobileApp) SetSupersedingAppCount(value *int32)() {
    if m != nil {
        m.supersedingAppCount = value
    }
}
// SetUploadState sets the uploadState property value. The upload state.
func (m *MobileApp) SetUploadState(value *int32)() {
    if m != nil {
        m.uploadState = value
    }
}
// SetUserStatuses sets the userStatuses property value. The list of installation states for this mobile app.
func (m *MobileApp) SetUserStatuses(value []UserAppInstallStatus)() {
    if m != nil {
        m.userStatuses = value
    }
}
