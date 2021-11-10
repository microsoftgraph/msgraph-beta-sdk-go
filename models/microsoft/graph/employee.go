package graph

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// 
type Employee struct {
    Entity
    // 
    address *PostalAddressType;
    // 
    birthDate *string;
    // 
    displayName *string;
    // 
    email *string;
    // 
    employmentDate *string;
    // 
    givenName *string;
    // 
    jobTitle *string;
    // 
    lastModifiedDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time;
    // 
    middleName *string;
    // 
    mobilePhone *string;
    // 
    number *string;
    // 
    personalEmail *string;
    // 
    phoneNumber *string;
    // 
    picture []Picture;
    // 
    statisticsGroupCode *string;
    // 
    status *string;
    // 
    surname *string;
    // 
    terminationDate *string;
}
// Instantiates a new employee and sets the default values.
func NewEmployee()(*Employee) {
    m := &Employee{
        Entity: *NewEntity(),
    }
    return m
}
// Gets the address property value. 
func (m *Employee) GetAddress()(*PostalAddressType) {
    if m == nil {
        return nil
    } else {
        return m.address
    }
}
// Gets the birthDate property value. 
func (m *Employee) GetBirthDate()(*string) {
    if m == nil {
        return nil
    } else {
        return m.birthDate
    }
}
// Gets the displayName property value. 
func (m *Employee) GetDisplayName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.displayName
    }
}
// Gets the email property value. 
func (m *Employee) GetEmail()(*string) {
    if m == nil {
        return nil
    } else {
        return m.email
    }
}
// Gets the employmentDate property value. 
func (m *Employee) GetEmploymentDate()(*string) {
    if m == nil {
        return nil
    } else {
        return m.employmentDate
    }
}
// Gets the givenName property value. 
func (m *Employee) GetGivenName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.givenName
    }
}
// Gets the jobTitle property value. 
func (m *Employee) GetJobTitle()(*string) {
    if m == nil {
        return nil
    } else {
        return m.jobTitle
    }
}
// Gets the lastModifiedDateTime property value. 
func (m *Employee) GetLastModifiedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.lastModifiedDateTime
    }
}
// Gets the middleName property value. 
func (m *Employee) GetMiddleName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.middleName
    }
}
// Gets the mobilePhone property value. 
func (m *Employee) GetMobilePhone()(*string) {
    if m == nil {
        return nil
    } else {
        return m.mobilePhone
    }
}
// Gets the number property value. 
func (m *Employee) GetNumber()(*string) {
    if m == nil {
        return nil
    } else {
        return m.number
    }
}
// Gets the personalEmail property value. 
func (m *Employee) GetPersonalEmail()(*string) {
    if m == nil {
        return nil
    } else {
        return m.personalEmail
    }
}
// Gets the phoneNumber property value. 
func (m *Employee) GetPhoneNumber()(*string) {
    if m == nil {
        return nil
    } else {
        return m.phoneNumber
    }
}
// Gets the picture property value. 
func (m *Employee) GetPicture()([]Picture) {
    if m == nil {
        return nil
    } else {
        return m.picture
    }
}
// Gets the statisticsGroupCode property value. 
func (m *Employee) GetStatisticsGroupCode()(*string) {
    if m == nil {
        return nil
    } else {
        return m.statisticsGroupCode
    }
}
// Gets the status property value. 
func (m *Employee) GetStatus()(*string) {
    if m == nil {
        return nil
    } else {
        return m.status
    }
}
// Gets the surname property value. 
func (m *Employee) GetSurname()(*string) {
    if m == nil {
        return nil
    } else {
        return m.surname
    }
}
// Gets the terminationDate property value. 
func (m *Employee) GetTerminationDate()(*string) {
    if m == nil {
        return nil
    } else {
        return m.terminationDate
    }
}
// The deserialization information for the current model
func (m *Employee) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["address"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewPostalAddressType() })
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAddress(val.(*PostalAddressType))
        }
        return nil
    }
    res["birthDate"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetBirthDate(val)
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
    res["email"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetEmail(val)
        }
        return nil
    }
    res["employmentDate"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetEmploymentDate(val)
        }
        return nil
    }
    res["givenName"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetGivenName(val)
        }
        return nil
    }
    res["jobTitle"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetJobTitle(val)
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
    res["middleName"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMiddleName(val)
        }
        return nil
    }
    res["mobilePhone"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMobilePhone(val)
        }
        return nil
    }
    res["number"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetNumber(val)
        }
        return nil
    }
    res["personalEmail"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPersonalEmail(val)
        }
        return nil
    }
    res["phoneNumber"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPhoneNumber(val)
        }
        return nil
    }
    res["picture"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewPicture() })
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]Picture, len(val))
            for i, v := range val {
                res[i] = *(v.(*Picture))
            }
            m.SetPicture(res)
        }
        return nil
    }
    res["statisticsGroupCode"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetStatisticsGroupCode(val)
        }
        return nil
    }
    res["status"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetStatus(val)
        }
        return nil
    }
    res["surname"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSurname(val)
        }
        return nil
    }
    res["terminationDate"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTerminationDate(val)
        }
        return nil
    }
    return res
}
func (m *Employee) IsNil()(bool) {
    return m == nil
}
// Serializes information the current object
// Parameters:
//  - writer : Serialization writer to use to serialize this model
func (m *Employee) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteObjectValue("address", m.GetAddress())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("birthDate", m.GetBirthDate())
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
        err = writer.WriteStringValue("email", m.GetEmail())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("employmentDate", m.GetEmploymentDate())
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
        err = writer.WriteStringValue("jobTitle", m.GetJobTitle())
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
        err = writer.WriteStringValue("middleName", m.GetMiddleName())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("mobilePhone", m.GetMobilePhone())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("number", m.GetNumber())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("personalEmail", m.GetPersonalEmail())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("phoneNumber", m.GetPhoneNumber())
        if err != nil {
            return err
        }
    }
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetPicture()))
        for i, v := range m.GetPicture() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("picture", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("statisticsGroupCode", m.GetStatisticsGroupCode())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("status", m.GetStatus())
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
        err = writer.WriteStringValue("terminationDate", m.GetTerminationDate())
        if err != nil {
            return err
        }
    }
    return nil
}
// Sets the address property value. 
// Parameters:
//  - value : Value to set for the address property.
func (m *Employee) SetAddress(value *PostalAddressType)() {
    m.address = value
}
// Sets the birthDate property value. 
// Parameters:
//  - value : Value to set for the birthDate property.
func (m *Employee) SetBirthDate(value *string)() {
    m.birthDate = value
}
// Sets the displayName property value. 
// Parameters:
//  - value : Value to set for the displayName property.
func (m *Employee) SetDisplayName(value *string)() {
    m.displayName = value
}
// Sets the email property value. 
// Parameters:
//  - value : Value to set for the email property.
func (m *Employee) SetEmail(value *string)() {
    m.email = value
}
// Sets the employmentDate property value. 
// Parameters:
//  - value : Value to set for the employmentDate property.
func (m *Employee) SetEmploymentDate(value *string)() {
    m.employmentDate = value
}
// Sets the givenName property value. 
// Parameters:
//  - value : Value to set for the givenName property.
func (m *Employee) SetGivenName(value *string)() {
    m.givenName = value
}
// Sets the jobTitle property value. 
// Parameters:
//  - value : Value to set for the jobTitle property.
func (m *Employee) SetJobTitle(value *string)() {
    m.jobTitle = value
}
// Sets the lastModifiedDateTime property value. 
// Parameters:
//  - value : Value to set for the lastModifiedDateTime property.
func (m *Employee) SetLastModifiedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.lastModifiedDateTime = value
}
// Sets the middleName property value. 
// Parameters:
//  - value : Value to set for the middleName property.
func (m *Employee) SetMiddleName(value *string)() {
    m.middleName = value
}
// Sets the mobilePhone property value. 
// Parameters:
//  - value : Value to set for the mobilePhone property.
func (m *Employee) SetMobilePhone(value *string)() {
    m.mobilePhone = value
}
// Sets the number property value. 
// Parameters:
//  - value : Value to set for the number property.
func (m *Employee) SetNumber(value *string)() {
    m.number = value
}
// Sets the personalEmail property value. 
// Parameters:
//  - value : Value to set for the personalEmail property.
func (m *Employee) SetPersonalEmail(value *string)() {
    m.personalEmail = value
}
// Sets the phoneNumber property value. 
// Parameters:
//  - value : Value to set for the phoneNumber property.
func (m *Employee) SetPhoneNumber(value *string)() {
    m.phoneNumber = value
}
// Sets the picture property value. 
// Parameters:
//  - value : Value to set for the picture property.
func (m *Employee) SetPicture(value []Picture)() {
    m.picture = value
}
// Sets the statisticsGroupCode property value. 
// Parameters:
//  - value : Value to set for the statisticsGroupCode property.
func (m *Employee) SetStatisticsGroupCode(value *string)() {
    m.statisticsGroupCode = value
}
// Sets the status property value. 
// Parameters:
//  - value : Value to set for the status property.
func (m *Employee) SetStatus(value *string)() {
    m.status = value
}
// Sets the surname property value. 
// Parameters:
//  - value : Value to set for the surname property.
func (m *Employee) SetSurname(value *string)() {
    m.surname = value
}
// Sets the terminationDate property value. 
// Parameters:
//  - value : Value to set for the terminationDate property.
func (m *Employee) SetTerminationDate(value *string)() {
    m.terminationDate = value
}
