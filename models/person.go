package models

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// Person provides operations to manage the collection of accessReview entities.
type Person struct {
    Entity
    // The person's birthday.
    birthday *string
    // The name of the person's company.
    companyName *string
    // The person's department.
    department *string
    // The person's display name.
    displayName *string
    // The person's email addresses.
    emailAddresses []RankedEmailAddressable
    // The person's given name.
    givenName *string
    // true if the user has flagged this person as a favorite.
    isFavorite *bool
    // The type of mailbox that is represented by the person's email address.
    mailboxType *string
    // The location of the person's office.
    officeLocation *string
    // Free-form notes that the user has taken about this person.
    personNotes *string
    // The type of person, for example distribution list.
    personType *string
    // The person's phone numbers.
    phones []Phoneable
    // The person's addresses.
    postalAddresses []Locationable
    // The person's profession.
    profession *string
    // The sources the user data comes from, for example Directory or Outlook Contacts.
    sources []PersonDataSourceable
    // The person's surname.
    surname *string
    // The person's title.
    title *string
    // The user principal name (UPN) of the person. The UPN is an Internet-style login name for the person based on the Internet standard RFC 822. By convention, this should map to the person's email name. The general format is alias@domain.
    userPrincipalName *string
    // The person's websites.
    websites []Websiteable
    // The phonetic Japanese name of the person's company.
    yomiCompany *string
}
// NewPerson instantiates a new person and sets the default values.
func NewPerson()(*Person) {
    m := &Person{
        Entity: *NewEntity(),
    }
    odataTypeValue := "#microsoft.graph.person";
    m.SetOdataType(&odataTypeValue);
    return m
}
// CreatePersonFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreatePersonFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewPerson(), nil
}
// GetBirthday gets the birthday property value. The person's birthday.
func (m *Person) GetBirthday()(*string) {
    return m.birthday
}
// GetCompanyName gets the companyName property value. The name of the person's company.
func (m *Person) GetCompanyName()(*string) {
    return m.companyName
}
// GetDepartment gets the department property value. The person's department.
func (m *Person) GetDepartment()(*string) {
    return m.department
}
// GetDisplayName gets the displayName property value. The person's display name.
func (m *Person) GetDisplayName()(*string) {
    return m.displayName
}
// GetEmailAddresses gets the emailAddresses property value. The person's email addresses.
func (m *Person) GetEmailAddresses()([]RankedEmailAddressable) {
    return m.emailAddresses
}
// GetFieldDeserializers the deserialization information for the current model
func (m *Person) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["birthday"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetBirthday)
    res["companyName"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetCompanyName)
    res["department"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetDepartment)
    res["displayName"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetDisplayName)
    res["emailAddresses"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfObjectValues(CreateRankedEmailAddressFromDiscriminatorValue , m.SetEmailAddresses)
    res["givenName"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetGivenName)
    res["isFavorite"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetBoolValue(m.SetIsFavorite)
    res["mailboxType"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetMailboxType)
    res["officeLocation"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetOfficeLocation)
    res["personNotes"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetPersonNotes)
    res["personType"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetPersonType)
    res["phones"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfObjectValues(CreatePhoneFromDiscriminatorValue , m.SetPhones)
    res["postalAddresses"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfObjectValues(CreateLocationFromDiscriminatorValue , m.SetPostalAddresses)
    res["profession"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetProfession)
    res["sources"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfObjectValues(CreatePersonDataSourceFromDiscriminatorValue , m.SetSources)
    res["surname"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetSurname)
    res["title"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetTitle)
    res["userPrincipalName"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetUserPrincipalName)
    res["websites"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfObjectValues(CreateWebsiteFromDiscriminatorValue , m.SetWebsites)
    res["yomiCompany"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetYomiCompany)
    return res
}
// GetGivenName gets the givenName property value. The person's given name.
func (m *Person) GetGivenName()(*string) {
    return m.givenName
}
// GetIsFavorite gets the isFavorite property value. true if the user has flagged this person as a favorite.
func (m *Person) GetIsFavorite()(*bool) {
    return m.isFavorite
}
// GetMailboxType gets the mailboxType property value. The type of mailbox that is represented by the person's email address.
func (m *Person) GetMailboxType()(*string) {
    return m.mailboxType
}
// GetOfficeLocation gets the officeLocation property value. The location of the person's office.
func (m *Person) GetOfficeLocation()(*string) {
    return m.officeLocation
}
// GetPersonNotes gets the personNotes property value. Free-form notes that the user has taken about this person.
func (m *Person) GetPersonNotes()(*string) {
    return m.personNotes
}
// GetPersonType gets the personType property value. The type of person, for example distribution list.
func (m *Person) GetPersonType()(*string) {
    return m.personType
}
// GetPhones gets the phones property value. The person's phone numbers.
func (m *Person) GetPhones()([]Phoneable) {
    return m.phones
}
// GetPostalAddresses gets the postalAddresses property value. The person's addresses.
func (m *Person) GetPostalAddresses()([]Locationable) {
    return m.postalAddresses
}
// GetProfession gets the profession property value. The person's profession.
func (m *Person) GetProfession()(*string) {
    return m.profession
}
// GetSources gets the sources property value. The sources the user data comes from, for example Directory or Outlook Contacts.
func (m *Person) GetSources()([]PersonDataSourceable) {
    return m.sources
}
// GetSurname gets the surname property value. The person's surname.
func (m *Person) GetSurname()(*string) {
    return m.surname
}
// GetTitle gets the title property value. The person's title.
func (m *Person) GetTitle()(*string) {
    return m.title
}
// GetUserPrincipalName gets the userPrincipalName property value. The user principal name (UPN) of the person. The UPN is an Internet-style login name for the person based on the Internet standard RFC 822. By convention, this should map to the person's email name. The general format is alias@domain.
func (m *Person) GetUserPrincipalName()(*string) {
    return m.userPrincipalName
}
// GetWebsites gets the websites property value. The person's websites.
func (m *Person) GetWebsites()([]Websiteable) {
    return m.websites
}
// GetYomiCompany gets the yomiCompany property value. The phonetic Japanese name of the person's company.
func (m *Person) GetYomiCompany()(*string) {
    return m.yomiCompany
}
// Serialize serializes information the current object
func (m *Person) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteStringValue("birthday", m.GetBirthday())
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
    {
        err = writer.WriteStringValue("givenName", m.GetGivenName())
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
        err = writer.WriteStringValue("mailboxType", m.GetMailboxType())
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
        err = writer.WriteStringValue("personNotes", m.GetPersonNotes())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("personType", m.GetPersonType())
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
    if m.GetSources() != nil {
        cast := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.CollectionCast[i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable](m.GetSources())
        err = writer.WriteCollectionOfObjectValues("sources", cast)
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
    {
        err = writer.WriteStringValue("userPrincipalName", m.GetUserPrincipalName())
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
        err = writer.WriteStringValue("yomiCompany", m.GetYomiCompany())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetBirthday sets the birthday property value. The person's birthday.
func (m *Person) SetBirthday(value *string)() {
    m.birthday = value
}
// SetCompanyName sets the companyName property value. The name of the person's company.
func (m *Person) SetCompanyName(value *string)() {
    m.companyName = value
}
// SetDepartment sets the department property value. The person's department.
func (m *Person) SetDepartment(value *string)() {
    m.department = value
}
// SetDisplayName sets the displayName property value. The person's display name.
func (m *Person) SetDisplayName(value *string)() {
    m.displayName = value
}
// SetEmailAddresses sets the emailAddresses property value. The person's email addresses.
func (m *Person) SetEmailAddresses(value []RankedEmailAddressable)() {
    m.emailAddresses = value
}
// SetGivenName sets the givenName property value. The person's given name.
func (m *Person) SetGivenName(value *string)() {
    m.givenName = value
}
// SetIsFavorite sets the isFavorite property value. true if the user has flagged this person as a favorite.
func (m *Person) SetIsFavorite(value *bool)() {
    m.isFavorite = value
}
// SetMailboxType sets the mailboxType property value. The type of mailbox that is represented by the person's email address.
func (m *Person) SetMailboxType(value *string)() {
    m.mailboxType = value
}
// SetOfficeLocation sets the officeLocation property value. The location of the person's office.
func (m *Person) SetOfficeLocation(value *string)() {
    m.officeLocation = value
}
// SetPersonNotes sets the personNotes property value. Free-form notes that the user has taken about this person.
func (m *Person) SetPersonNotes(value *string)() {
    m.personNotes = value
}
// SetPersonType sets the personType property value. The type of person, for example distribution list.
func (m *Person) SetPersonType(value *string)() {
    m.personType = value
}
// SetPhones sets the phones property value. The person's phone numbers.
func (m *Person) SetPhones(value []Phoneable)() {
    m.phones = value
}
// SetPostalAddresses sets the postalAddresses property value. The person's addresses.
func (m *Person) SetPostalAddresses(value []Locationable)() {
    m.postalAddresses = value
}
// SetProfession sets the profession property value. The person's profession.
func (m *Person) SetProfession(value *string)() {
    m.profession = value
}
// SetSources sets the sources property value. The sources the user data comes from, for example Directory or Outlook Contacts.
func (m *Person) SetSources(value []PersonDataSourceable)() {
    m.sources = value
}
// SetSurname sets the surname property value. The person's surname.
func (m *Person) SetSurname(value *string)() {
    m.surname = value
}
// SetTitle sets the title property value. The person's title.
func (m *Person) SetTitle(value *string)() {
    m.title = value
}
// SetUserPrincipalName sets the userPrincipalName property value. The user principal name (UPN) of the person. The UPN is an Internet-style login name for the person based on the Internet standard RFC 822. By convention, this should map to the person's email name. The general format is alias@domain.
func (m *Person) SetUserPrincipalName(value *string)() {
    m.userPrincipalName = value
}
// SetWebsites sets the websites property value. The person's websites.
func (m *Person) SetWebsites(value []Websiteable)() {
    m.websites = value
}
// SetYomiCompany sets the yomiCompany property value. The phonetic Japanese name of the person's company.
func (m *Person) SetYomiCompany(value *string)() {
    m.yomiCompany = value
}
