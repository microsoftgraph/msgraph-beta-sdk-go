package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// MobileApp an abstract class containing the base properties for Intune mobile apps.
type MobileApp struct {
    Entity
    // The list of group assignments for this mobile app.
    assignments []MobileAppAssignmentable
    // The list of categories for this app.
    categories []MobileAppCategoryable
    // The date and time the app was created.
    createdDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // The total number of dependencies the child app has.
    dependentAppCount *int32
    // The description of the app.
    description *string
    // The developer of the app.
    developer *string
    // The list of installation states for this mobile app.
    deviceStatuses []MobileAppInstallStatusable
    // The admin provided or imported title of the app.
    displayName *string
    // The more information Url.
    informationUrl *string
    // Mobile App Install Summary.
    installSummary MobileAppInstallSummaryable
    // The value indicating whether the app is assigned to at least one group.
    isAssigned *bool
    // The value indicating whether the app is marked as featured by the admin.
    isFeatured *bool
    // The large icon, to be displayed in the app details and used for upload of the icon.
    largeIcon MimeContentable
    // The date and time the app was last modified.
    lastModifiedDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // Notes for the app.
    notes *string
    // The owner of the app.
    owner *string
    // The privacy statement Url.
    privacyInformationUrl *string
    // The publisher of the app.
    publisher *string
    // Indicates the publishing state of an app.
    publishingState *MobileAppPublishingState
    // List of relationships for this mobile app.
    relationships []MobileAppRelationshipable
    // List of scope tag ids for this mobile app.
    roleScopeTagIds []string
    // The total number of apps this app is directly or indirectly superseded by.
    supersededAppCount *int32
    // The total number of apps this app directly or indirectly supersedes.
    supersedingAppCount *int32
    // The upload state.
    uploadState *int32
    // The list of installation states for this mobile app.
    userStatuses []UserAppInstallStatusable
}
// NewMobileApp instantiates a new mobileApp and sets the default values.
func NewMobileApp()(*MobileApp) {
    m := &MobileApp{
        Entity: *NewEntity(),
    }
    odataTypeValue := "#microsoft.graph.mobileApp";
    m.SetOdataType(&odataTypeValue);
    return m
}
// CreateMobileAppFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateMobileAppFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    if parseNode != nil {
        mappingValueNode, err := parseNode.GetChildNode("@odata.type")
        if err != nil {
            return nil, err
        }
        if mappingValueNode != nil {
            mappingValue, err := mappingValueNode.GetStringValue()
            if err != nil {
                return nil, err
            }
            if mappingValue != nil {
                switch *mappingValue {
                    case "#microsoft.graph.androidForWorkApp":
                        return NewAndroidForWorkApp(), nil
                    case "#microsoft.graph.androidLobApp":
                        return NewAndroidLobApp(), nil
                    case "#microsoft.graph.androidManagedStoreApp":
                        return NewAndroidManagedStoreApp(), nil
                    case "#microsoft.graph.androidManagedStoreWebApp":
                        return NewAndroidManagedStoreWebApp(), nil
                    case "#microsoft.graph.androidStoreApp":
                        return NewAndroidStoreApp(), nil
                    case "#microsoft.graph.iosiPadOSWebClip":
                        return NewIosiPadOSWebClip(), nil
                    case "#microsoft.graph.iosLobApp":
                        return NewIosLobApp(), nil
                    case "#microsoft.graph.iosStoreApp":
                        return NewIosStoreApp(), nil
                    case "#microsoft.graph.iosVppApp":
                        return NewIosVppApp(), nil
                    case "#microsoft.graph.macOSDmgApp":
                        return NewMacOSDmgApp(), nil
                    case "#microsoft.graph.macOSLobApp":
                        return NewMacOSLobApp(), nil
                    case "#microsoft.graph.macOSMdatpApp":
                        return NewMacOSMdatpApp(), nil
                    case "#microsoft.graph.macOSMicrosoftEdgeApp":
                        return NewMacOSMicrosoftEdgeApp(), nil
                    case "#microsoft.graph.macOSOfficeSuiteApp":
                        return NewMacOSOfficeSuiteApp(), nil
                    case "#microsoft.graph.macOsVppApp":
                        return NewMacOsVppApp(), nil
                    case "#microsoft.graph.managedAndroidLobApp":
                        return NewManagedAndroidLobApp(), nil
                    case "#microsoft.graph.managedAndroidStoreApp":
                        return NewManagedAndroidStoreApp(), nil
                    case "#microsoft.graph.managedApp":
                        return NewManagedApp(), nil
                    case "#microsoft.graph.managedIOSLobApp":
                        return NewManagedIOSLobApp(), nil
                    case "#microsoft.graph.managedIOSStoreApp":
                        return NewManagedIOSStoreApp(), nil
                    case "#microsoft.graph.managedMobileLobApp":
                        return NewManagedMobileLobApp(), nil
                    case "#microsoft.graph.microsoftStoreForBusinessApp":
                        return NewMicrosoftStoreForBusinessApp(), nil
                    case "#microsoft.graph.mobileLobApp":
                        return NewMobileLobApp(), nil
                    case "#microsoft.graph.officeSuiteApp":
                        return NewOfficeSuiteApp(), nil
                    case "#microsoft.graph.webApp":
                        return NewWebApp(), nil
                    case "#microsoft.graph.win32LobApp":
                        return NewWin32LobApp(), nil
                    case "#microsoft.graph.windowsAppX":
                        return NewWindowsAppX(), nil
                    case "#microsoft.graph.windowsMicrosoftEdgeApp":
                        return NewWindowsMicrosoftEdgeApp(), nil
                    case "#microsoft.graph.windowsMobileMSI":
                        return NewWindowsMobileMSI(), nil
                    case "#microsoft.graph.windowsPhone81AppX":
                        return NewWindowsPhone81AppX(), nil
                    case "#microsoft.graph.windowsPhone81AppXBundle":
                        return NewWindowsPhone81AppXBundle(), nil
                    case "#microsoft.graph.windowsPhone81StoreApp":
                        return NewWindowsPhone81StoreApp(), nil
                    case "#microsoft.graph.windowsPhoneXAP":
                        return NewWindowsPhoneXAP(), nil
                    case "#microsoft.graph.windowsStoreApp":
                        return NewWindowsStoreApp(), nil
                    case "#microsoft.graph.windowsUniversalAppX":
                        return NewWindowsUniversalAppX(), nil
                    case "#microsoft.graph.windowsWebApp":
                        return NewWindowsWebApp(), nil
                    case "#microsoft.graph.winGetApp":
                        return NewWinGetApp(), nil
                }
            }
        }
    }
    return NewMobileApp(), nil
}
// GetAssignments gets the assignments property value. The list of group assignments for this mobile app.
func (m *MobileApp) GetAssignments()([]MobileAppAssignmentable) {
    return m.assignments
}
// GetCategories gets the categories property value. The list of categories for this app.
func (m *MobileApp) GetCategories()([]MobileAppCategoryable) {
    return m.categories
}
// GetCreatedDateTime gets the createdDateTime property value. The date and time the app was created.
func (m *MobileApp) GetCreatedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.createdDateTime
}
// GetDependentAppCount gets the dependentAppCount property value. The total number of dependencies the child app has.
func (m *MobileApp) GetDependentAppCount()(*int32) {
    return m.dependentAppCount
}
// GetDescription gets the description property value. The description of the app.
func (m *MobileApp) GetDescription()(*string) {
    return m.description
}
// GetDeveloper gets the developer property value. The developer of the app.
func (m *MobileApp) GetDeveloper()(*string) {
    return m.developer
}
// GetDeviceStatuses gets the deviceStatuses property value. The list of installation states for this mobile app.
func (m *MobileApp) GetDeviceStatuses()([]MobileAppInstallStatusable) {
    return m.deviceStatuses
}
// GetDisplayName gets the displayName property value. The admin provided or imported title of the app.
func (m *MobileApp) GetDisplayName()(*string) {
    return m.displayName
}
// GetFieldDeserializers the deserialization information for the current model
func (m *MobileApp) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["assignments"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfObjectValues(CreateMobileAppAssignmentFromDiscriminatorValue , m.SetAssignments)
    res["categories"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfObjectValues(CreateMobileAppCategoryFromDiscriminatorValue , m.SetCategories)
    res["createdDateTime"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetTimeValue(m.SetCreatedDateTime)
    res["dependentAppCount"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetInt32Value(m.SetDependentAppCount)
    res["description"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetDescription)
    res["developer"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetDeveloper)
    res["deviceStatuses"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfObjectValues(CreateMobileAppInstallStatusFromDiscriminatorValue , m.SetDeviceStatuses)
    res["displayName"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetDisplayName)
    res["informationUrl"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetInformationUrl)
    res["installSummary"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetObjectValue(CreateMobileAppInstallSummaryFromDiscriminatorValue , m.SetInstallSummary)
    res["isAssigned"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetBoolValue(m.SetIsAssigned)
    res["isFeatured"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetBoolValue(m.SetIsFeatured)
    res["largeIcon"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetObjectValue(CreateMimeContentFromDiscriminatorValue , m.SetLargeIcon)
    res["lastModifiedDateTime"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetTimeValue(m.SetLastModifiedDateTime)
    res["notes"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetNotes)
    res["owner"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetOwner)
    res["privacyInformationUrl"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetPrivacyInformationUrl)
    res["publisher"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetPublisher)
    res["publishingState"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetEnumValue(ParseMobileAppPublishingState , m.SetPublishingState)
    res["relationships"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfObjectValues(CreateMobileAppRelationshipFromDiscriminatorValue , m.SetRelationships)
    res["roleScopeTagIds"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfPrimitiveValues("string" , m.SetRoleScopeTagIds)
    res["supersededAppCount"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetInt32Value(m.SetSupersededAppCount)
    res["supersedingAppCount"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetInt32Value(m.SetSupersedingAppCount)
    res["uploadState"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetInt32Value(m.SetUploadState)
    res["userStatuses"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfObjectValues(CreateUserAppInstallStatusFromDiscriminatorValue , m.SetUserStatuses)
    return res
}
// GetInformationUrl gets the informationUrl property value. The more information Url.
func (m *MobileApp) GetInformationUrl()(*string) {
    return m.informationUrl
}
// GetInstallSummary gets the installSummary property value. Mobile App Install Summary.
func (m *MobileApp) GetInstallSummary()(MobileAppInstallSummaryable) {
    return m.installSummary
}
// GetIsAssigned gets the isAssigned property value. The value indicating whether the app is assigned to at least one group.
func (m *MobileApp) GetIsAssigned()(*bool) {
    return m.isAssigned
}
// GetIsFeatured gets the isFeatured property value. The value indicating whether the app is marked as featured by the admin.
func (m *MobileApp) GetIsFeatured()(*bool) {
    return m.isFeatured
}
// GetLargeIcon gets the largeIcon property value. The large icon, to be displayed in the app details and used for upload of the icon.
func (m *MobileApp) GetLargeIcon()(MimeContentable) {
    return m.largeIcon
}
// GetLastModifiedDateTime gets the lastModifiedDateTime property value. The date and time the app was last modified.
func (m *MobileApp) GetLastModifiedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.lastModifiedDateTime
}
// GetNotes gets the notes property value. Notes for the app.
func (m *MobileApp) GetNotes()(*string) {
    return m.notes
}
// GetOwner gets the owner property value. The owner of the app.
func (m *MobileApp) GetOwner()(*string) {
    return m.owner
}
// GetPrivacyInformationUrl gets the privacyInformationUrl property value. The privacy statement Url.
func (m *MobileApp) GetPrivacyInformationUrl()(*string) {
    return m.privacyInformationUrl
}
// GetPublisher gets the publisher property value. The publisher of the app.
func (m *MobileApp) GetPublisher()(*string) {
    return m.publisher
}
// GetPublishingState gets the publishingState property value. Indicates the publishing state of an app.
func (m *MobileApp) GetPublishingState()(*MobileAppPublishingState) {
    return m.publishingState
}
// GetRelationships gets the relationships property value. List of relationships for this mobile app.
func (m *MobileApp) GetRelationships()([]MobileAppRelationshipable) {
    return m.relationships
}
// GetRoleScopeTagIds gets the roleScopeTagIds property value. List of scope tag ids for this mobile app.
func (m *MobileApp) GetRoleScopeTagIds()([]string) {
    return m.roleScopeTagIds
}
// GetSupersededAppCount gets the supersededAppCount property value. The total number of apps this app is directly or indirectly superseded by.
func (m *MobileApp) GetSupersededAppCount()(*int32) {
    return m.supersededAppCount
}
// GetSupersedingAppCount gets the supersedingAppCount property value. The total number of apps this app directly or indirectly supersedes.
func (m *MobileApp) GetSupersedingAppCount()(*int32) {
    return m.supersedingAppCount
}
// GetUploadState gets the uploadState property value. The upload state.
func (m *MobileApp) GetUploadState()(*int32) {
    return m.uploadState
}
// GetUserStatuses gets the userStatuses property value. The list of installation states for this mobile app.
func (m *MobileApp) GetUserStatuses()([]UserAppInstallStatusable) {
    return m.userStatuses
}
// Serialize serializes information the current object
func (m *MobileApp) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetAssignments() != nil {
        cast := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.CollectionCast[i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable](m.GetAssignments())
        err = writer.WriteCollectionOfObjectValues("assignments", cast)
        if err != nil {
            return err
        }
    }
    if m.GetCategories() != nil {
        cast := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.CollectionCast[i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable](m.GetCategories())
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
        cast := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.CollectionCast[i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable](m.GetDeviceStatuses())
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
        cast := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.CollectionCast[i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable](m.GetRelationships())
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
        cast := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.CollectionCast[i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable](m.GetUserStatuses())
        err = writer.WriteCollectionOfObjectValues("userStatuses", cast)
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAssignments sets the assignments property value. The list of group assignments for this mobile app.
func (m *MobileApp) SetAssignments(value []MobileAppAssignmentable)() {
    m.assignments = value
}
// SetCategories sets the categories property value. The list of categories for this app.
func (m *MobileApp) SetCategories(value []MobileAppCategoryable)() {
    m.categories = value
}
// SetCreatedDateTime sets the createdDateTime property value. The date and time the app was created.
func (m *MobileApp) SetCreatedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.createdDateTime = value
}
// SetDependentAppCount sets the dependentAppCount property value. The total number of dependencies the child app has.
func (m *MobileApp) SetDependentAppCount(value *int32)() {
    m.dependentAppCount = value
}
// SetDescription sets the description property value. The description of the app.
func (m *MobileApp) SetDescription(value *string)() {
    m.description = value
}
// SetDeveloper sets the developer property value. The developer of the app.
func (m *MobileApp) SetDeveloper(value *string)() {
    m.developer = value
}
// SetDeviceStatuses sets the deviceStatuses property value. The list of installation states for this mobile app.
func (m *MobileApp) SetDeviceStatuses(value []MobileAppInstallStatusable)() {
    m.deviceStatuses = value
}
// SetDisplayName sets the displayName property value. The admin provided or imported title of the app.
func (m *MobileApp) SetDisplayName(value *string)() {
    m.displayName = value
}
// SetInformationUrl sets the informationUrl property value. The more information Url.
func (m *MobileApp) SetInformationUrl(value *string)() {
    m.informationUrl = value
}
// SetInstallSummary sets the installSummary property value. Mobile App Install Summary.
func (m *MobileApp) SetInstallSummary(value MobileAppInstallSummaryable)() {
    m.installSummary = value
}
// SetIsAssigned sets the isAssigned property value. The value indicating whether the app is assigned to at least one group.
func (m *MobileApp) SetIsAssigned(value *bool)() {
    m.isAssigned = value
}
// SetIsFeatured sets the isFeatured property value. The value indicating whether the app is marked as featured by the admin.
func (m *MobileApp) SetIsFeatured(value *bool)() {
    m.isFeatured = value
}
// SetLargeIcon sets the largeIcon property value. The large icon, to be displayed in the app details and used for upload of the icon.
func (m *MobileApp) SetLargeIcon(value MimeContentable)() {
    m.largeIcon = value
}
// SetLastModifiedDateTime sets the lastModifiedDateTime property value. The date and time the app was last modified.
func (m *MobileApp) SetLastModifiedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.lastModifiedDateTime = value
}
// SetNotes sets the notes property value. Notes for the app.
func (m *MobileApp) SetNotes(value *string)() {
    m.notes = value
}
// SetOwner sets the owner property value. The owner of the app.
func (m *MobileApp) SetOwner(value *string)() {
    m.owner = value
}
// SetPrivacyInformationUrl sets the privacyInformationUrl property value. The privacy statement Url.
func (m *MobileApp) SetPrivacyInformationUrl(value *string)() {
    m.privacyInformationUrl = value
}
// SetPublisher sets the publisher property value. The publisher of the app.
func (m *MobileApp) SetPublisher(value *string)() {
    m.publisher = value
}
// SetPublishingState sets the publishingState property value. Indicates the publishing state of an app.
func (m *MobileApp) SetPublishingState(value *MobileAppPublishingState)() {
    m.publishingState = value
}
// SetRelationships sets the relationships property value. List of relationships for this mobile app.
func (m *MobileApp) SetRelationships(value []MobileAppRelationshipable)() {
    m.relationships = value
}
// SetRoleScopeTagIds sets the roleScopeTagIds property value. List of scope tag ids for this mobile app.
func (m *MobileApp) SetRoleScopeTagIds(value []string)() {
    m.roleScopeTagIds = value
}
// SetSupersededAppCount sets the supersededAppCount property value. The total number of apps this app is directly or indirectly superseded by.
func (m *MobileApp) SetSupersededAppCount(value *int32)() {
    m.supersededAppCount = value
}
// SetSupersedingAppCount sets the supersedingAppCount property value. The total number of apps this app directly or indirectly supersedes.
func (m *MobileApp) SetSupersedingAppCount(value *int32)() {
    m.supersedingAppCount = value
}
// SetUploadState sets the uploadState property value. The upload state.
func (m *MobileApp) SetUploadState(value *int32)() {
    m.uploadState = value
}
// SetUserStatuses sets the userStatuses property value. The list of installation states for this mobile app.
func (m *MobileApp) SetUserStatuses(value []UserAppInstallStatusable)() {
    m.userStatuses = value
}
