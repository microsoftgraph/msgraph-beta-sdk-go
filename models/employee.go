package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// Employee 
type Employee struct {
    Entity
}
// NewEmployee instantiates a new employee and sets the default values.
func NewEmployee()(*Employee) {
    m := &Employee{
        Entity: *NewEntity(),
    }
    return m
}
// CreateEmployeeFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateEmployeeFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewEmployee(), nil
}
// GetAddress gets the address property value. The address property
func (m *Employee) GetAddress()(PostalAddressTypeable) {
    val, err := m.GetBackingStore().Get("address")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(PostalAddressTypeable)
    }
    return nil
}
// GetBirthDate gets the birthDate property value. The birthDate property
func (m *Employee) GetBirthDate()(*i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.DateOnly) {
    val, err := m.GetBackingStore().Get("birthDate")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.DateOnly)
    }
    return nil
}
// GetDisplayName gets the displayName property value. The displayName property
func (m *Employee) GetDisplayName()(*string) {
    val, err := m.GetBackingStore().Get("displayName")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetEmail gets the email property value. The email property
func (m *Employee) GetEmail()(*string) {
    val, err := m.GetBackingStore().Get("email")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetEmploymentDate gets the employmentDate property value. The employmentDate property
func (m *Employee) GetEmploymentDate()(*i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.DateOnly) {
    val, err := m.GetBackingStore().Get("employmentDate")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.DateOnly)
    }
    return nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *Employee) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["address"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreatePostalAddressTypeFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAddress(val.(PostalAddressTypeable))
        }
        return nil
    }
    res["birthDate"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetDateOnlyValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetBirthDate(val)
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
    res["email"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetEmail(val)
        }
        return nil
    }
    res["employmentDate"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetDateOnlyValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetEmploymentDate(val)
        }
        return nil
    }
    res["givenName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetGivenName(val)
        }
        return nil
    }
    res["jobTitle"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetJobTitle(val)
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
    res["middleName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMiddleName(val)
        }
        return nil
    }
    res["mobilePhone"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMobilePhone(val)
        }
        return nil
    }
    res["number"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetNumber(val)
        }
        return nil
    }
    res["personalEmail"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPersonalEmail(val)
        }
        return nil
    }
    res["phoneNumber"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPhoneNumber(val)
        }
        return nil
    }
    res["picture"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreatePictureFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]Pictureable, len(val))
            for i, v := range val {
                res[i] = v.(Pictureable)
            }
            m.SetPicture(res)
        }
        return nil
    }
    res["statisticsGroupCode"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetStatisticsGroupCode(val)
        }
        return nil
    }
    res["status"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetStatus(val)
        }
        return nil
    }
    res["surname"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSurname(val)
        }
        return nil
    }
    res["terminationDate"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetDateOnlyValue()
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
// GetGivenName gets the givenName property value. The givenName property
func (m *Employee) GetGivenName()(*string) {
    val, err := m.GetBackingStore().Get("givenName")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetJobTitle gets the jobTitle property value. The jobTitle property
func (m *Employee) GetJobTitle()(*string) {
    val, err := m.GetBackingStore().Get("jobTitle")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetLastModifiedDateTime gets the lastModifiedDateTime property value. The lastModifiedDateTime property
func (m *Employee) GetLastModifiedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    val, err := m.GetBackingStore().Get("lastModifiedDateTime")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    }
    return nil
}
// GetMiddleName gets the middleName property value. The middleName property
func (m *Employee) GetMiddleName()(*string) {
    val, err := m.GetBackingStore().Get("middleName")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetMobilePhone gets the mobilePhone property value. The mobilePhone property
func (m *Employee) GetMobilePhone()(*string) {
    val, err := m.GetBackingStore().Get("mobilePhone")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetNumber gets the number property value. The number property
func (m *Employee) GetNumber()(*string) {
    val, err := m.GetBackingStore().Get("number")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetPersonalEmail gets the personalEmail property value. The personalEmail property
func (m *Employee) GetPersonalEmail()(*string) {
    val, err := m.GetBackingStore().Get("personalEmail")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetPhoneNumber gets the phoneNumber property value. The phoneNumber property
func (m *Employee) GetPhoneNumber()(*string) {
    val, err := m.GetBackingStore().Get("phoneNumber")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetPicture gets the picture property value. The picture property
func (m *Employee) GetPicture()([]Pictureable) {
    val, err := m.GetBackingStore().Get("picture")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]Pictureable)
    }
    return nil
}
// GetStatisticsGroupCode gets the statisticsGroupCode property value. The statisticsGroupCode property
func (m *Employee) GetStatisticsGroupCode()(*string) {
    val, err := m.GetBackingStore().Get("statisticsGroupCode")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetStatus gets the status property value. The status property
func (m *Employee) GetStatus()(*string) {
    val, err := m.GetBackingStore().Get("status")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetSurname gets the surname property value. The surname property
func (m *Employee) GetSurname()(*string) {
    val, err := m.GetBackingStore().Get("surname")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetTerminationDate gets the terminationDate property value. The terminationDate property
func (m *Employee) GetTerminationDate()(*i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.DateOnly) {
    val, err := m.GetBackingStore().Get("terminationDate")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.DateOnly)
    }
    return nil
}
// Serialize serializes information the current object
func (m *Employee) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
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
        err = writer.WriteDateOnlyValue("birthDate", m.GetBirthDate())
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
        err = writer.WriteDateOnlyValue("employmentDate", m.GetEmploymentDate())
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
    if m.GetPicture() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetPicture()))
        for i, v := range m.GetPicture() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
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
        err = writer.WriteDateOnlyValue("terminationDate", m.GetTerminationDate())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAddress sets the address property value. The address property
func (m *Employee) SetAddress(value PostalAddressTypeable)() {
    err := m.GetBackingStore().Set("address", value)
    if err != nil {
        panic(err)
    }
}
// SetBirthDate sets the birthDate property value. The birthDate property
func (m *Employee) SetBirthDate(value *i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.DateOnly)() {
    err := m.GetBackingStore().Set("birthDate", value)
    if err != nil {
        panic(err)
    }
}
// SetDisplayName sets the displayName property value. The displayName property
func (m *Employee) SetDisplayName(value *string)() {
    err := m.GetBackingStore().Set("displayName", value)
    if err != nil {
        panic(err)
    }
}
// SetEmail sets the email property value. The email property
func (m *Employee) SetEmail(value *string)() {
    err := m.GetBackingStore().Set("email", value)
    if err != nil {
        panic(err)
    }
}
// SetEmploymentDate sets the employmentDate property value. The employmentDate property
func (m *Employee) SetEmploymentDate(value *i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.DateOnly)() {
    err := m.GetBackingStore().Set("employmentDate", value)
    if err != nil {
        panic(err)
    }
}
// SetGivenName sets the givenName property value. The givenName property
func (m *Employee) SetGivenName(value *string)() {
    err := m.GetBackingStore().Set("givenName", value)
    if err != nil {
        panic(err)
    }
}
// SetJobTitle sets the jobTitle property value. The jobTitle property
func (m *Employee) SetJobTitle(value *string)() {
    err := m.GetBackingStore().Set("jobTitle", value)
    if err != nil {
        panic(err)
    }
}
// SetLastModifiedDateTime sets the lastModifiedDateTime property value. The lastModifiedDateTime property
func (m *Employee) SetLastModifiedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    err := m.GetBackingStore().Set("lastModifiedDateTime", value)
    if err != nil {
        panic(err)
    }
}
// SetMiddleName sets the middleName property value. The middleName property
func (m *Employee) SetMiddleName(value *string)() {
    err := m.GetBackingStore().Set("middleName", value)
    if err != nil {
        panic(err)
    }
}
// SetMobilePhone sets the mobilePhone property value. The mobilePhone property
func (m *Employee) SetMobilePhone(value *string)() {
    err := m.GetBackingStore().Set("mobilePhone", value)
    if err != nil {
        panic(err)
    }
}
// SetNumber sets the number property value. The number property
func (m *Employee) SetNumber(value *string)() {
    err := m.GetBackingStore().Set("number", value)
    if err != nil {
        panic(err)
    }
}
// SetPersonalEmail sets the personalEmail property value. The personalEmail property
func (m *Employee) SetPersonalEmail(value *string)() {
    err := m.GetBackingStore().Set("personalEmail", value)
    if err != nil {
        panic(err)
    }
}
// SetPhoneNumber sets the phoneNumber property value. The phoneNumber property
func (m *Employee) SetPhoneNumber(value *string)() {
    err := m.GetBackingStore().Set("phoneNumber", value)
    if err != nil {
        panic(err)
    }
}
// SetPicture sets the picture property value. The picture property
func (m *Employee) SetPicture(value []Pictureable)() {
    err := m.GetBackingStore().Set("picture", value)
    if err != nil {
        panic(err)
    }
}
// SetStatisticsGroupCode sets the statisticsGroupCode property value. The statisticsGroupCode property
func (m *Employee) SetStatisticsGroupCode(value *string)() {
    err := m.GetBackingStore().Set("statisticsGroupCode", value)
    if err != nil {
        panic(err)
    }
}
// SetStatus sets the status property value. The status property
func (m *Employee) SetStatus(value *string)() {
    err := m.GetBackingStore().Set("status", value)
    if err != nil {
        panic(err)
    }
}
// SetSurname sets the surname property value. The surname property
func (m *Employee) SetSurname(value *string)() {
    err := m.GetBackingStore().Set("surname", value)
    if err != nil {
        panic(err)
    }
}
// SetTerminationDate sets the terminationDate property value. The terminationDate property
func (m *Employee) SetTerminationDate(value *i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.DateOnly)() {
    err := m.GetBackingStore().Set("terminationDate", value)
    if err != nil {
        panic(err)
    }
}
// Employeeable 
type Employeeable interface {
    Entityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAddress()(PostalAddressTypeable)
    GetBirthDate()(*i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.DateOnly)
    GetDisplayName()(*string)
    GetEmail()(*string)
    GetEmploymentDate()(*i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.DateOnly)
    GetGivenName()(*string)
    GetJobTitle()(*string)
    GetLastModifiedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetMiddleName()(*string)
    GetMobilePhone()(*string)
    GetNumber()(*string)
    GetPersonalEmail()(*string)
    GetPhoneNumber()(*string)
    GetPicture()([]Pictureable)
    GetStatisticsGroupCode()(*string)
    GetStatus()(*string)
    GetSurname()(*string)
    GetTerminationDate()(*i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.DateOnly)
    SetAddress(value PostalAddressTypeable)()
    SetBirthDate(value *i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.DateOnly)()
    SetDisplayName(value *string)()
    SetEmail(value *string)()
    SetEmploymentDate(value *i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.DateOnly)()
    SetGivenName(value *string)()
    SetJobTitle(value *string)()
    SetLastModifiedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetMiddleName(value *string)()
    SetMobilePhone(value *string)()
    SetNumber(value *string)()
    SetPersonalEmail(value *string)()
    SetPhoneNumber(value *string)()
    SetPicture(value []Pictureable)()
    SetStatisticsGroupCode(value *string)()
    SetStatus(value *string)()
    SetSurname(value *string)()
    SetTerminationDate(value *i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.DateOnly)()
}
