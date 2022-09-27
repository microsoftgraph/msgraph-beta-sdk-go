package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// Contact 
type Contact struct {
    OutlookItem
    // The name of the contact's assistant.
    assistantName *string
    // The contact's birthday. The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z
    birthday *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // The names of the contact's children.
    children []string
    // The name of the contact's company.
    companyName *string
    // The contact's department.
    department *string
    // The contact's display name. You can specify the display name in a create or update operation. Note that later updates to other properties may cause an automatically generated value to overwrite the displayName value you have specified. To preserve a pre-existing value, always include it as displayName in an update operation.
    displayName *string
    // The contact's email addresses.
    emailAddresses []TypedEmailAddressable
    // The collection of open extensions defined for the contact. Nullable.
    extensions []Extensionable
    // The name the contact is filed under.
    fileAs *string
    // The flag value that indicates the status, start date, due date, or completion date for the contact.
    flag FollowupFlagable
    // The contact's gender.
    gender *string
    // The contact's generation.
    generation *string
    // The contact's given name.
    givenName *string
    // The imAddresses property
    imAddresses []string
    // The initials property
    initials *string
    // The isFavorite property
    isFavorite *bool
    // The jobTitle property
    jobTitle *string
    // The manager property
    manager *string
    // The middleName property
    middleName *string
    // The collection of multi-value extended properties defined for the contact. Read-only. Nullable.
    multiValueExtendedProperties []MultiValueLegacyExtendedPropertyable
    // The nickName property
    nickName *string
    // The officeLocation property
    officeLocation *string
    // The parentFolderId property
    parentFolderId *string
    // The personalNotes property
    personalNotes *string
    // The phones property
    phones []Phoneable
    // Optional contact picture. You can get or set a photo for a contact.
    photo ProfilePhotoable
    // The postalAddresses property
    postalAddresses []PhysicalAddressable
    // The profession property
    profession *string
    // The collection of single-value extended properties defined for the contact. Read-only. Nullable.
    singleValueExtendedProperties []SingleValueLegacyExtendedPropertyable
    // The spouseName property
    spouseName *string
    // The surname property
    surname *string
    // The title property
    title *string
    // The websites property
    websites []Websiteable
    // The weddingAnniversary property
    weddingAnniversary *i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.DateOnly
    // The yomiCompanyName property
    yomiCompanyName *string
    // The yomiGivenName property
    yomiGivenName *string
    // The yomiSurname property
    yomiSurname *string
}
// NewContact instantiates a new Contact and sets the default values.
func NewContact()(*Contact) {
    m := &Contact{
        OutlookItem: *NewOutlookItem(),
    }
    odataTypeValue := "#microsoft.graph.contact";
    m.SetOdataType(&odataTypeValue);
    return m
}
// CreateContactFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateContactFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewContact(), nil
}
// GetAssistantName gets the assistantName property value. The name of the contact's assistant.
func (m *Contact) GetAssistantName()(*string) {
    return m.assistantName
}
// GetBirthday gets the birthday property value. The contact's birthday. The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z
func (m *Contact) GetBirthday()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.birthday
}
// GetChildren gets the children property value. The names of the contact's children.
func (m *Contact) GetChildren()([]string) {
    return m.children
}
// GetCompanyName gets the companyName property value. The name of the contact's company.
func (m *Contact) GetCompanyName()(*string) {
    return m.companyName
}
// GetDepartment gets the department property value. The contact's department.
func (m *Contact) GetDepartment()(*string) {
    return m.department
}
// GetDisplayName gets the displayName property value. The contact's display name. You can specify the display name in a create or update operation. Note that later updates to other properties may cause an automatically generated value to overwrite the displayName value you have specified. To preserve a pre-existing value, always include it as displayName in an update operation.
func (m *Contact) GetDisplayName()(*string) {
    return m.displayName
}
// GetEmailAddresses gets the emailAddresses property value. The contact's email addresses.
func (m *Contact) GetEmailAddresses()([]TypedEmailAddressable) {
    return m.emailAddresses
}
// GetExtensions gets the extensions property value. The collection of open extensions defined for the contact. Nullable.
func (m *Contact) GetExtensions()([]Extensionable) {
    return m.extensions
}
// GetFieldDeserializers the deserialization information for the current model
func (m *Contact) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.OutlookItem.GetFieldDeserializers()
    res["assistantName"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetAssistantName)
    res["birthday"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetTimeValue(m.SetBirthday)
    res["children"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfPrimitiveValues("string" , m.SetChildren)
    res["companyName"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetCompanyName)
    res["department"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetDepartment)
    res["displayName"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetDisplayName)
    res["emailAddresses"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfObjectValues(CreateTypedEmailAddressFromDiscriminatorValue , m.SetEmailAddresses)
    res["extensions"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfObjectValues(CreateExtensionFromDiscriminatorValue , m.SetExtensions)
    res["fileAs"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetFileAs)
    res["flag"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetObjectValue(CreateFollowupFlagFromDiscriminatorValue , m.SetFlag)
    res["gender"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetGender)
    res["generation"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetGeneration)
    res["givenName"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetGivenName)
    res["imAddresses"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfPrimitiveValues("string" , m.SetImAddresses)
    res["initials"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetInitials)
    res["isFavorite"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetBoolValue(m.SetIsFavorite)
    res["jobTitle"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetJobTitle)
    res["manager"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetManager)
    res["middleName"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetMiddleName)
    res["multiValueExtendedProperties"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfObjectValues(CreateMultiValueLegacyExtendedPropertyFromDiscriminatorValue , m.SetMultiValueExtendedProperties)
    res["nickName"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetNickName)
    res["officeLocation"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetOfficeLocation)
    res["parentFolderId"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetParentFolderId)
    res["personalNotes"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetPersonalNotes)
    res["phones"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfObjectValues(CreatePhoneFromDiscriminatorValue , m.SetPhones)
    res["photo"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetObjectValue(CreateProfilePhotoFromDiscriminatorValue , m.SetPhoto)
    res["postalAddresses"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfObjectValues(CreatePhysicalAddressFromDiscriminatorValue , m.SetPostalAddresses)
    res["profession"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetProfession)
    res["singleValueExtendedProperties"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfObjectValues(CreateSingleValueLegacyExtendedPropertyFromDiscriminatorValue , m.SetSingleValueExtendedProperties)
    res["spouseName"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetSpouseName)
    res["surname"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetSurname)
    res["title"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetTitle)
    res["websites"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfObjectValues(CreateWebsiteFromDiscriminatorValue , m.SetWebsites)
    res["weddingAnniversary"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetDateOnlyValue(m.SetWeddingAnniversary)
    res["yomiCompanyName"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetYomiCompanyName)
    res["yomiGivenName"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetYomiGivenName)
    res["yomiSurname"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetYomiSurname)
    return res
}
// GetFileAs gets the fileAs property value. The name the contact is filed under.
func (m *Contact) GetFileAs()(*string) {
    return m.fileAs
}
// GetFlag gets the flag property value. The flag value that indicates the status, start date, due date, or completion date for the contact.
func (m *Contact) GetFlag()(FollowupFlagable) {
    return m.flag
}
// GetGender gets the gender property value. The contact's gender.
func (m *Contact) GetGender()(*string) {
    return m.gender
}
// GetGeneration gets the generation property value. The contact's generation.
func (m *Contact) GetGeneration()(*string) {
    return m.generation
}
// GetGivenName gets the givenName property value. The contact's given name.
func (m *Contact) GetGivenName()(*string) {
    return m.givenName
}
// GetImAddresses gets the imAddresses property value. The imAddresses property
func (m *Contact) GetImAddresses()([]string) {
    return m.imAddresses
}
// GetInitials gets the initials property value. The initials property
func (m *Contact) GetInitials()(*string) {
    return m.initials
}
// GetIsFavorite gets the isFavorite property value. The isFavorite property
func (m *Contact) GetIsFavorite()(*bool) {
    return m.isFavorite
}
// GetJobTitle gets the jobTitle property value. The jobTitle property
func (m *Contact) GetJobTitle()(*string) {
    return m.jobTitle
}
// GetManager gets the manager property value. The manager property
func (m *Contact) GetManager()(*string) {
    return m.manager
}
// GetMiddleName gets the middleName property value. The middleName property
func (m *Contact) GetMiddleName()(*string) {
    return m.middleName
}
// GetMultiValueExtendedProperties gets the multiValueExtendedProperties property value. The collection of multi-value extended properties defined for the contact. Read-only. Nullable.
func (m *Contact) GetMultiValueExtendedProperties()([]MultiValueLegacyExtendedPropertyable) {
    return m.multiValueExtendedProperties
}
// GetNickName gets the nickName property value. The nickName property
func (m *Contact) GetNickName()(*string) {
    return m.nickName
}
// GetOfficeLocation gets the officeLocation property value. The officeLocation property
func (m *Contact) GetOfficeLocation()(*string) {
    return m.officeLocation
}
// GetParentFolderId gets the parentFolderId property value. The parentFolderId property
func (m *Contact) GetParentFolderId()(*string) {
    return m.parentFolderId
}
// GetPersonalNotes gets the personalNotes property value. The personalNotes property
func (m *Contact) GetPersonalNotes()(*string) {
    return m.personalNotes
}
// GetPhones gets the phones property value. The phones property
func (m *Contact) GetPhones()([]Phoneable) {
    return m.phones
}
// GetPhoto gets the photo property value. Optional contact picture. You can get or set a photo for a contact.
func (m *Contact) GetPhoto()(ProfilePhotoable) {
    return m.photo
}
// GetPostalAddresses gets the postalAddresses property value. The postalAddresses property
func (m *Contact) GetPostalAddresses()([]PhysicalAddressable) {
    return m.postalAddresses
}
// GetProfession gets the profession property value. The profession property
func (m *Contact) GetProfession()(*string) {
    return m.profession
}
// GetSingleValueExtendedProperties gets the singleValueExtendedProperties property value. The collection of single-value extended properties defined for the contact. Read-only. Nullable.
func (m *Contact) GetSingleValueExtendedProperties()([]SingleValueLegacyExtendedPropertyable) {
    return m.singleValueExtendedProperties
}
// GetSpouseName gets the spouseName property value. The spouseName property
func (m *Contact) GetSpouseName()(*string) {
    return m.spouseName
}
// GetSurname gets the surname property value. The surname property
func (m *Contact) GetSurname()(*string) {
    return m.surname
}
// GetTitle gets the title property value. The title property
func (m *Contact) GetTitle()(*string) {
    return m.title
}
// GetWebsites gets the websites property value. The websites property
func (m *Contact) GetWebsites()([]Websiteable) {
    return m.websites
}
// GetWeddingAnniversary gets the weddingAnniversary property value. The weddingAnniversary property
func (m *Contact) GetWeddingAnniversary()(*i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.DateOnly) {
    return m.weddingAnniversary
}
// GetYomiCompanyName gets the yomiCompanyName property value. The yomiCompanyName property
func (m *Contact) GetYomiCompanyName()(*string) {
    return m.yomiCompanyName
}
// GetYomiGivenName gets the yomiGivenName property value. The yomiGivenName property
func (m *Contact) GetYomiGivenName()(*string) {
    return m.yomiGivenName
}
// GetYomiSurname gets the yomiSurname property value. The yomiSurname property
func (m *Contact) GetYomiSurname()(*string) {
    return m.yomiSurname
}
// Serialize serializes information the current object
func (m *Contact) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.OutlookItem.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteStringValue("assistantName", m.GetAssistantName())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteTimeValue("birthday", m.GetBirthday())
        if err != nil {
            return err
        }
    }
    if m.GetChildren() != nil {
        err = writer.WriteCollectionOfStringValues("children", m.GetChildren())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("companyName", m.GetCompanyName())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("department", m.GetDepartment())
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
    if m.GetEmailAddresses() != nil {
        cast := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.CollectionCast[i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable](m.GetEmailAddresses())
        err = writer.WriteCollectionOfObjectValues("emailAddresses", cast)
        if err != nil {
            return err
        }
    }
    if m.GetExtensions() != nil {
        cast := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.CollectionCast[i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable](m.GetExtensions())
        err = writer.WriteCollectionOfObjectValues("extensions", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("fileAs", m.GetFileAs())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("flag", m.GetFlag())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("gender", m.GetGender())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("generation", m.GetGeneration())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("givenName", m.GetGivenName())
        if err != nil {
            return err
        }
    }
    if m.GetImAddresses() != nil {
        err = writer.WriteCollectionOfStringValues("imAddresses", m.GetImAddresses())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("initials", m.GetInitials())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("isFavorite", m.GetIsFavorite())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("jobTitle", m.GetJobTitle())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("manager", m.GetManager())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("middleName", m.GetMiddleName())
        if err != nil {
            return err
        }
    }
    if m.GetMultiValueExtendedProperties() != nil {
        cast := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.CollectionCast[i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable](m.GetMultiValueExtendedProperties())
        err = writer.WriteCollectionOfObjectValues("multiValueExtendedProperties", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("nickName", m.GetNickName())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("officeLocation", m.GetOfficeLocation())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("parentFolderId", m.GetParentFolderId())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("personalNotes", m.GetPersonalNotes())
        if err != nil {
            return err
        }
    }
    if m.GetPhones() != nil {
        cast := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.CollectionCast[i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable](m.GetPhones())
        err = writer.WriteCollectionOfObjectValues("phones", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("photo", m.GetPhoto())
        if err != nil {
            return err
        }
    }
    if m.GetPostalAddresses() != nil {
        cast := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.CollectionCast[i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable](m.GetPostalAddresses())
        err = writer.WriteCollectionOfObjectValues("postalAddresses", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("profession", m.GetProfession())
        if err != nil {
            return err
        }
    }
    if m.GetSingleValueExtendedProperties() != nil {
        cast := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.CollectionCast[i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable](m.GetSingleValueExtendedProperties())
        err = writer.WriteCollectionOfObjectValues("singleValueExtendedProperties", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("spouseName", m.GetSpouseName())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("surname", m.GetSurname())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("title", m.GetTitle())
        if err != nil {
            return err
        }
    }
    if m.GetWebsites() != nil {
        cast := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.CollectionCast[i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable](m.GetWebsites())
        err = writer.WriteCollectionOfObjectValues("websites", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteDateOnlyValue("weddingAnniversary", m.GetWeddingAnniversary())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("yomiCompanyName", m.GetYomiCompanyName())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("yomiGivenName", m.GetYomiGivenName())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("yomiSurname", m.GetYomiSurname())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAssistantName sets the assistantName property value. The name of the contact's assistant.
func (m *Contact) SetAssistantName(value *string)() {
    m.assistantName = value
}
// SetBirthday sets the birthday property value. The contact's birthday. The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z
func (m *Contact) SetBirthday(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.birthday = value
}
// SetChildren sets the children property value. The names of the contact's children.
func (m *Contact) SetChildren(value []string)() {
    m.children = value
}
// SetCompanyName sets the companyName property value. The name of the contact's company.
func (m *Contact) SetCompanyName(value *string)() {
    m.companyName = value
}
// SetDepartment sets the department property value. The contact's department.
func (m *Contact) SetDepartment(value *string)() {
    m.department = value
}
// SetDisplayName sets the displayName property value. The contact's display name. You can specify the display name in a create or update operation. Note that later updates to other properties may cause an automatically generated value to overwrite the displayName value you have specified. To preserve a pre-existing value, always include it as displayName in an update operation.
func (m *Contact) SetDisplayName(value *string)() {
    m.displayName = value
}
// SetEmailAddresses sets the emailAddresses property value. The contact's email addresses.
func (m *Contact) SetEmailAddresses(value []TypedEmailAddressable)() {
    m.emailAddresses = value
}
// SetExtensions sets the extensions property value. The collection of open extensions defined for the contact. Nullable.
func (m *Contact) SetExtensions(value []Extensionable)() {
    m.extensions = value
}
// SetFileAs sets the fileAs property value. The name the contact is filed under.
func (m *Contact) SetFileAs(value *string)() {
    m.fileAs = value
}
// SetFlag sets the flag property value. The flag value that indicates the status, start date, due date, or completion date for the contact.
func (m *Contact) SetFlag(value FollowupFlagable)() {
    m.flag = value
}
// SetGender sets the gender property value. The contact's gender.
func (m *Contact) SetGender(value *string)() {
    m.gender = value
}
// SetGeneration sets the generation property value. The contact's generation.
func (m *Contact) SetGeneration(value *string)() {
    m.generation = value
}
// SetGivenName sets the givenName property value. The contact's given name.
func (m *Contact) SetGivenName(value *string)() {
    m.givenName = value
}
// SetImAddresses sets the imAddresses property value. The imAddresses property
func (m *Contact) SetImAddresses(value []string)() {
    m.imAddresses = value
}
// SetInitials sets the initials property value. The initials property
func (m *Contact) SetInitials(value *string)() {
    m.initials = value
}
// SetIsFavorite sets the isFavorite property value. The isFavorite property
func (m *Contact) SetIsFavorite(value *bool)() {
    m.isFavorite = value
}
// SetJobTitle sets the jobTitle property value. The jobTitle property
func (m *Contact) SetJobTitle(value *string)() {
    m.jobTitle = value
}
// SetManager sets the manager property value. The manager property
func (m *Contact) SetManager(value *string)() {
    m.manager = value
}
// SetMiddleName sets the middleName property value. The middleName property
func (m *Contact) SetMiddleName(value *string)() {
    m.middleName = value
}
// SetMultiValueExtendedProperties sets the multiValueExtendedProperties property value. The collection of multi-value extended properties defined for the contact. Read-only. Nullable.
func (m *Contact) SetMultiValueExtendedProperties(value []MultiValueLegacyExtendedPropertyable)() {
    m.multiValueExtendedProperties = value
}
// SetNickName sets the nickName property value. The nickName property
func (m *Contact) SetNickName(value *string)() {
    m.nickName = value
}
// SetOfficeLocation sets the officeLocation property value. The officeLocation property
func (m *Contact) SetOfficeLocation(value *string)() {
    m.officeLocation = value
}
// SetParentFolderId sets the parentFolderId property value. The parentFolderId property
func (m *Contact) SetParentFolderId(value *string)() {
    m.parentFolderId = value
}
// SetPersonalNotes sets the personalNotes property value. The personalNotes property
func (m *Contact) SetPersonalNotes(value *string)() {
    m.personalNotes = value
}
// SetPhones sets the phones property value. The phones property
func (m *Contact) SetPhones(value []Phoneable)() {
    m.phones = value
}
// SetPhoto sets the photo property value. Optional contact picture. You can get or set a photo for a contact.
func (m *Contact) SetPhoto(value ProfilePhotoable)() {
    m.photo = value
}
// SetPostalAddresses sets the postalAddresses property value. The postalAddresses property
func (m *Contact) SetPostalAddresses(value []PhysicalAddressable)() {
    m.postalAddresses = value
}
// SetProfession sets the profession property value. The profession property
func (m *Contact) SetProfession(value *string)() {
    m.profession = value
}
// SetSingleValueExtendedProperties sets the singleValueExtendedProperties property value. The collection of single-value extended properties defined for the contact. Read-only. Nullable.
func (m *Contact) SetSingleValueExtendedProperties(value []SingleValueLegacyExtendedPropertyable)() {
    m.singleValueExtendedProperties = value
}
// SetSpouseName sets the spouseName property value. The spouseName property
func (m *Contact) SetSpouseName(value *string)() {
    m.spouseName = value
}
// SetSurname sets the surname property value. The surname property
func (m *Contact) SetSurname(value *string)() {
    m.surname = value
}
// SetTitle sets the title property value. The title property
func (m *Contact) SetTitle(value *string)() {
    m.title = value
}
// SetWebsites sets the websites property value. The websites property
func (m *Contact) SetWebsites(value []Websiteable)() {
    m.websites = value
}
// SetWeddingAnniversary sets the weddingAnniversary property value. The weddingAnniversary property
func (m *Contact) SetWeddingAnniversary(value *i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.DateOnly)() {
    m.weddingAnniversary = value
}
// SetYomiCompanyName sets the yomiCompanyName property value. The yomiCompanyName property
func (m *Contact) SetYomiCompanyName(value *string)() {
    m.yomiCompanyName = value
}
// SetYomiGivenName sets the yomiGivenName property value. The yomiGivenName property
func (m *Contact) SetYomiGivenName(value *string)() {
    m.yomiGivenName = value
}
// SetYomiSurname sets the yomiSurname property value. The yomiSurname property
func (m *Contact) SetYomiSurname(value *string)() {
    m.yomiSurname = value
}
