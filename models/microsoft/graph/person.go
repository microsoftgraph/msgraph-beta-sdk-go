package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// 
type Person struct {
    Entity
    // The person's birthday.
    birthday *string;
    // The name of the person's company.
    companyName *string;
    // The person's department.
    department *string;
    // The person's display name.
    displayName *string;
    // The person's email addresses.
    emailAddresses []RankedEmailAddress;
    // The person's given name.
    givenName *string;
    // true if the user has flagged this person as a favorite.
    isFavorite *bool;
    // The type of mailbox that is represented by the person's email address.
    mailboxType *string;
    // The location of the person's office.
    officeLocation *string;
    // Free-form notes that the user has taken about this person.
    personNotes *string;
    // The type of person.
    personType *string;
    // The person's phone numbers.
    phones []Phone;
    // The person's addresses.
    postalAddresses []Location;
    // The person's profession.
    profession *string;
    // The sources the user data comes from, for example Directory or Outlook Contacts.
    sources []PersonDataSource;
    // The person's surname.
    surname *string;
    // The person's title.
    title *string;
    // The user principal name (UPN) of the person. The UPN is an Internet-style login name for the person based on the Internet standard RFC 822. By convention, this should map to the person's email name. The general format is alias@domain.
    userPrincipalName *string;
    // The person's websites.
    websites []Website;
    // The phonetic Japanese name of the person's company.
    yomiCompany *string;
}
// Instantiates a new person and sets the default values.
func NewPerson()(*Person) {
    m := &Person{
        Entity: *NewEntity(),
    }
    return m
}
// Gets the birthday property value. The person's birthday.
func (m *Person) GetBirthday()(*string) {
    if m == nil {
        return nil
    } else {
        return m.birthday
    }
}
// Gets the companyName property value. The name of the person's company.
func (m *Person) GetCompanyName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.companyName
    }
}
// Gets the department property value. The person's department.
func (m *Person) GetDepartment()(*string) {
    if m == nil {
        return nil
    } else {
        return m.department
    }
}
// Gets the displayName property value. The person's display name.
func (m *Person) GetDisplayName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.displayName
    }
}
// Gets the emailAddresses property value. The person's email addresses.
func (m *Person) GetEmailAddresses()([]RankedEmailAddress) {
    if m == nil {
        return nil
    } else {
        return m.emailAddresses
    }
}
// Gets the givenName property value. The person's given name.
func (m *Person) GetGivenName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.givenName
    }
}
// Gets the isFavorite property value. true if the user has flagged this person as a favorite.
func (m *Person) GetIsFavorite()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.isFavorite
    }
}
// Gets the mailboxType property value. The type of mailbox that is represented by the person's email address.
func (m *Person) GetMailboxType()(*string) {
    if m == nil {
        return nil
    } else {
        return m.mailboxType
    }
}
// Gets the officeLocation property value. The location of the person's office.
func (m *Person) GetOfficeLocation()(*string) {
    if m == nil {
        return nil
    } else {
        return m.officeLocation
    }
}
// Gets the personNotes property value. Free-form notes that the user has taken about this person.
func (m *Person) GetPersonNotes()(*string) {
    if m == nil {
        return nil
    } else {
        return m.personNotes
    }
}
// Gets the personType property value. The type of person.
func (m *Person) GetPersonType()(*string) {
    if m == nil {
        return nil
    } else {
        return m.personType
    }
}
// Gets the phones property value. The person's phone numbers.
func (m *Person) GetPhones()([]Phone) {
    if m == nil {
        return nil
    } else {
        return m.phones
    }
}
// Gets the postalAddresses property value. The person's addresses.
func (m *Person) GetPostalAddresses()([]Location) {
    if m == nil {
        return nil
    } else {
        return m.postalAddresses
    }
}
// Gets the profession property value. The person's profession.
func (m *Person) GetProfession()(*string) {
    if m == nil {
        return nil
    } else {
        return m.profession
    }
}
// Gets the sources property value. The sources the user data comes from, for example Directory or Outlook Contacts.
func (m *Person) GetSources()([]PersonDataSource) {
    if m == nil {
        return nil
    } else {
        return m.sources
    }
}
// Gets the surname property value. The person's surname.
func (m *Person) GetSurname()(*string) {
    if m == nil {
        return nil
    } else {
        return m.surname
    }
}
// Gets the title property value. The person's title.
func (m *Person) GetTitle()(*string) {
    if m == nil {
        return nil
    } else {
        return m.title
    }
}
// Gets the userPrincipalName property value. The user principal name (UPN) of the person. The UPN is an Internet-style login name for the person based on the Internet standard RFC 822. By convention, this should map to the person's email name. The general format is alias@domain.
func (m *Person) GetUserPrincipalName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.userPrincipalName
    }
}
// Gets the websites property value. The person's websites.
func (m *Person) GetWebsites()([]Website) {
    if m == nil {
        return nil
    } else {
        return m.websites
    }
}
// Gets the yomiCompany property value. The phonetic Japanese name of the person's company.
func (m *Person) GetYomiCompany()(*string) {
    if m == nil {
        return nil
    } else {
        return m.yomiCompany
    }
}
// The deserialization information for the current model
func (m *Person) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["birthday"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetBirthday(val)
        return nil
    }
    res["companyName"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetCompanyName(val)
        return nil
    }
    res["department"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetDepartment(val)
        return nil
    }
    res["displayName"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetDisplayName(val)
        return nil
    }
    res["emailAddresses"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewRankedEmailAddress() })
        if err != nil {
            return err
        }
        res := make([]RankedEmailAddress, len(val))
        for i, v := range val {
            res[i] = *(v.(*RankedEmailAddress))
        }
        m.SetEmailAddresses(res)
        return nil
    }
    res["givenName"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetGivenName(val)
        return nil
    }
    res["isFavorite"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        m.SetIsFavorite(val)
        return nil
    }
    res["mailboxType"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetMailboxType(val)
        return nil
    }
    res["officeLocation"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetOfficeLocation(val)
        return nil
    }
    res["personNotes"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetPersonNotes(val)
        return nil
    }
    res["personType"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetPersonType(val)
        return nil
    }
    res["phones"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewPhone() })
        if err != nil {
            return err
        }
        res := make([]Phone, len(val))
        for i, v := range val {
            res[i] = *(v.(*Phone))
        }
        m.SetPhones(res)
        return nil
    }
    res["postalAddresses"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewLocation() })
        if err != nil {
            return err
        }
        res := make([]Location, len(val))
        for i, v := range val {
            res[i] = *(v.(*Location))
        }
        m.SetPostalAddresses(res)
        return nil
    }
    res["profession"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetProfession(val)
        return nil
    }
    res["sources"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewPersonDataSource() })
        if err != nil {
            return err
        }
        res := make([]PersonDataSource, len(val))
        for i, v := range val {
            res[i] = *(v.(*PersonDataSource))
        }
        m.SetSources(res)
        return nil
    }
    res["surname"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetSurname(val)
        return nil
    }
    res["title"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetTitle(val)
        return nil
    }
    res["userPrincipalName"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetUserPrincipalName(val)
        return nil
    }
    res["websites"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewWebsite() })
        if err != nil {
            return err
        }
        res := make([]Website, len(val))
        for i, v := range val {
            res[i] = *(v.(*Website))
        }
        m.SetWebsites(res)
        return nil
    }
    res["yomiCompany"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetYomiCompany(val)
        return nil
    }
    return res
}
func (m *Person) IsNil()(bool) {
    return m == nil
}
// Serializes information the current object
// Parameters:
//  - writer : Serialization writer to use to serialize this model
func (m *Person) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
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
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetEmailAddresses()))
        for i, v := range m.GetEmailAddresses() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
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
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetPhones()))
        for i, v := range m.GetPhones() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("phones", cast)
        if err != nil {
            return err
        }
    }
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetPostalAddresses()))
        for i, v := range m.GetPostalAddresses() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
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
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetSources()))
        for i, v := range m.GetSources() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
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
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetWebsites()))
        for i, v := range m.GetWebsites() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
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
// Sets the birthday property value. The person's birthday.
// Parameters:
//  - value : Value to set for the birthday property.
func (m *Person) SetBirthday(value *string)() {
    m.birthday = value
}
// Sets the companyName property value. The name of the person's company.
// Parameters:
//  - value : Value to set for the companyName property.
func (m *Person) SetCompanyName(value *string)() {
    m.companyName = value
}
// Sets the department property value. The person's department.
// Parameters:
//  - value : Value to set for the department property.
func (m *Person) SetDepartment(value *string)() {
    m.department = value
}
// Sets the displayName property value. The person's display name.
// Parameters:
//  - value : Value to set for the displayName property.
func (m *Person) SetDisplayName(value *string)() {
    m.displayName = value
}
// Sets the emailAddresses property value. The person's email addresses.
// Parameters:
//  - value : Value to set for the emailAddresses property.
func (m *Person) SetEmailAddresses(value []RankedEmailAddress)() {
    m.emailAddresses = value
}
// Sets the givenName property value. The person's given name.
// Parameters:
//  - value : Value to set for the givenName property.
func (m *Person) SetGivenName(value *string)() {
    m.givenName = value
}
// Sets the isFavorite property value. true if the user has flagged this person as a favorite.
// Parameters:
//  - value : Value to set for the isFavorite property.
func (m *Person) SetIsFavorite(value *bool)() {
    m.isFavorite = value
}
// Sets the mailboxType property value. The type of mailbox that is represented by the person's email address.
// Parameters:
//  - value : Value to set for the mailboxType property.
func (m *Person) SetMailboxType(value *string)() {
    m.mailboxType = value
}
// Sets the officeLocation property value. The location of the person's office.
// Parameters:
//  - value : Value to set for the officeLocation property.
func (m *Person) SetOfficeLocation(value *string)() {
    m.officeLocation = value
}
// Sets the personNotes property value. Free-form notes that the user has taken about this person.
// Parameters:
//  - value : Value to set for the personNotes property.
func (m *Person) SetPersonNotes(value *string)() {
    m.personNotes = value
}
// Sets the personType property value. The type of person.
// Parameters:
//  - value : Value to set for the personType property.
func (m *Person) SetPersonType(value *string)() {
    m.personType = value
}
// Sets the phones property value. The person's phone numbers.
// Parameters:
//  - value : Value to set for the phones property.
func (m *Person) SetPhones(value []Phone)() {
    m.phones = value
}
// Sets the postalAddresses property value. The person's addresses.
// Parameters:
//  - value : Value to set for the postalAddresses property.
func (m *Person) SetPostalAddresses(value []Location)() {
    m.postalAddresses = value
}
// Sets the profession property value. The person's profession.
// Parameters:
//  - value : Value to set for the profession property.
func (m *Person) SetProfession(value *string)() {
    m.profession = value
}
// Sets the sources property value. The sources the user data comes from, for example Directory or Outlook Contacts.
// Parameters:
//  - value : Value to set for the sources property.
func (m *Person) SetSources(value []PersonDataSource)() {
    m.sources = value
}
// Sets the surname property value. The person's surname.
// Parameters:
//  - value : Value to set for the surname property.
func (m *Person) SetSurname(value *string)() {
    m.surname = value
}
// Sets the title property value. The person's title.
// Parameters:
//  - value : Value to set for the title property.
func (m *Person) SetTitle(value *string)() {
    m.title = value
}
// Sets the userPrincipalName property value. The user principal name (UPN) of the person. The UPN is an Internet-style login name for the person based on the Internet standard RFC 822. By convention, this should map to the person's email name. The general format is alias@domain.
// Parameters:
//  - value : Value to set for the userPrincipalName property.
func (m *Person) SetUserPrincipalName(value *string)() {
    m.userPrincipalName = value
}
// Sets the websites property value. The person's websites.
// Parameters:
//  - value : Value to set for the websites property.
func (m *Person) SetWebsites(value []Website)() {
    m.websites = value
}
// Sets the yomiCompany property value. The phonetic Japanese name of the person's company.
// Parameters:
//  - value : Value to set for the yomiCompany property.
func (m *Person) SetYomiCompany(value *string)() {
    m.yomiCompany = value
}
