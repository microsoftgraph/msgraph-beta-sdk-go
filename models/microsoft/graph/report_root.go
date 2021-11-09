package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// 
type ReportRoot struct {
    Entity
    // 
    applicationSignInDetailedSummary []ApplicationSignInDetailedSummary;
    // 
    authenticationMethods *AuthenticationMethodsRoot;
    // 
    credentialUserRegistrationDetails []CredentialUserRegistrationDetails;
    // 
    dailyPrintUsageByPrinter []PrintUsageByPrinter;
    // 
    dailyPrintUsageByUser []PrintUsageByUser;
    // 
    dailyPrintUsageSummariesByPrinter []PrintUsageByPrinter;
    // 
    dailyPrintUsageSummariesByUser []PrintUsageByUser;
    // 
    monthlyPrintUsageByPrinter []PrintUsageByPrinter;
    // 
    monthlyPrintUsageByUser []PrintUsageByUser;
    // 
    monthlyPrintUsageSummariesByPrinter []PrintUsageByPrinter;
    // 
    monthlyPrintUsageSummariesByUser []PrintUsageByUser;
    // 
    userCredentialUsageDetails []UserCredentialUsageDetails;
}
// Instantiates a new reportRoot and sets the default values.
func NewReportRoot()(*ReportRoot) {
    m := &ReportRoot{
        Entity: *NewEntity(),
    }
    return m
}
// Gets the applicationSignInDetailedSummary property value. 
func (m *ReportRoot) GetApplicationSignInDetailedSummary()([]ApplicationSignInDetailedSummary) {
    if m == nil {
        return nil
    } else {
        return m.applicationSignInDetailedSummary
    }
}
// Gets the authenticationMethods property value. 
func (m *ReportRoot) GetAuthenticationMethods()(*AuthenticationMethodsRoot) {
    if m == nil {
        return nil
    } else {
        return m.authenticationMethods
    }
}
// Gets the credentialUserRegistrationDetails property value. 
func (m *ReportRoot) GetCredentialUserRegistrationDetails()([]CredentialUserRegistrationDetails) {
    if m == nil {
        return nil
    } else {
        return m.credentialUserRegistrationDetails
    }
}
// Gets the dailyPrintUsageByPrinter property value. 
func (m *ReportRoot) GetDailyPrintUsageByPrinter()([]PrintUsageByPrinter) {
    if m == nil {
        return nil
    } else {
        return m.dailyPrintUsageByPrinter
    }
}
// Gets the dailyPrintUsageByUser property value. 
func (m *ReportRoot) GetDailyPrintUsageByUser()([]PrintUsageByUser) {
    if m == nil {
        return nil
    } else {
        return m.dailyPrintUsageByUser
    }
}
// Gets the dailyPrintUsageSummariesByPrinter property value. 
func (m *ReportRoot) GetDailyPrintUsageSummariesByPrinter()([]PrintUsageByPrinter) {
    if m == nil {
        return nil
    } else {
        return m.dailyPrintUsageSummariesByPrinter
    }
}
// Gets the dailyPrintUsageSummariesByUser property value. 
func (m *ReportRoot) GetDailyPrintUsageSummariesByUser()([]PrintUsageByUser) {
    if m == nil {
        return nil
    } else {
        return m.dailyPrintUsageSummariesByUser
    }
}
// Gets the monthlyPrintUsageByPrinter property value. 
func (m *ReportRoot) GetMonthlyPrintUsageByPrinter()([]PrintUsageByPrinter) {
    if m == nil {
        return nil
    } else {
        return m.monthlyPrintUsageByPrinter
    }
}
// Gets the monthlyPrintUsageByUser property value. 
func (m *ReportRoot) GetMonthlyPrintUsageByUser()([]PrintUsageByUser) {
    if m == nil {
        return nil
    } else {
        return m.monthlyPrintUsageByUser
    }
}
// Gets the monthlyPrintUsageSummariesByPrinter property value. 
func (m *ReportRoot) GetMonthlyPrintUsageSummariesByPrinter()([]PrintUsageByPrinter) {
    if m == nil {
        return nil
    } else {
        return m.monthlyPrintUsageSummariesByPrinter
    }
}
// Gets the monthlyPrintUsageSummariesByUser property value. 
func (m *ReportRoot) GetMonthlyPrintUsageSummariesByUser()([]PrintUsageByUser) {
    if m == nil {
        return nil
    } else {
        return m.monthlyPrintUsageSummariesByUser
    }
}
// Gets the userCredentialUsageDetails property value. 
func (m *ReportRoot) GetUserCredentialUsageDetails()([]UserCredentialUsageDetails) {
    if m == nil {
        return nil
    } else {
        return m.userCredentialUsageDetails
    }
}
// The deserialization information for the current model
func (m *ReportRoot) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["applicationSignInDetailedSummary"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewApplicationSignInDetailedSummary() })
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]ApplicationSignInDetailedSummary, len(val))
            for i, v := range val {
                res[i] = *(v.(*ApplicationSignInDetailedSummary))
            }
            m.SetApplicationSignInDetailedSummary(res)
        }
        return nil
    }
    res["authenticationMethods"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewAuthenticationMethodsRoot() })
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAuthenticationMethods(val.(*AuthenticationMethodsRoot))
        }
        return nil
    }
    res["credentialUserRegistrationDetails"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewCredentialUserRegistrationDetails() })
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]CredentialUserRegistrationDetails, len(val))
            for i, v := range val {
                res[i] = *(v.(*CredentialUserRegistrationDetails))
            }
            m.SetCredentialUserRegistrationDetails(res)
        }
        return nil
    }
    res["dailyPrintUsageByPrinter"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewPrintUsageByPrinter() })
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]PrintUsageByPrinter, len(val))
            for i, v := range val {
                res[i] = *(v.(*PrintUsageByPrinter))
            }
            m.SetDailyPrintUsageByPrinter(res)
        }
        return nil
    }
    res["dailyPrintUsageByUser"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewPrintUsageByUser() })
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]PrintUsageByUser, len(val))
            for i, v := range val {
                res[i] = *(v.(*PrintUsageByUser))
            }
            m.SetDailyPrintUsageByUser(res)
        }
        return nil
    }
    res["dailyPrintUsageSummariesByPrinter"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewPrintUsageByPrinter() })
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]PrintUsageByPrinter, len(val))
            for i, v := range val {
                res[i] = *(v.(*PrintUsageByPrinter))
            }
            m.SetDailyPrintUsageSummariesByPrinter(res)
        }
        return nil
    }
    res["dailyPrintUsageSummariesByUser"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewPrintUsageByUser() })
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]PrintUsageByUser, len(val))
            for i, v := range val {
                res[i] = *(v.(*PrintUsageByUser))
            }
            m.SetDailyPrintUsageSummariesByUser(res)
        }
        return nil
    }
    res["monthlyPrintUsageByPrinter"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewPrintUsageByPrinter() })
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]PrintUsageByPrinter, len(val))
            for i, v := range val {
                res[i] = *(v.(*PrintUsageByPrinter))
            }
            m.SetMonthlyPrintUsageByPrinter(res)
        }
        return nil
    }
    res["monthlyPrintUsageByUser"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewPrintUsageByUser() })
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]PrintUsageByUser, len(val))
            for i, v := range val {
                res[i] = *(v.(*PrintUsageByUser))
            }
            m.SetMonthlyPrintUsageByUser(res)
        }
        return nil
    }
    res["monthlyPrintUsageSummariesByPrinter"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewPrintUsageByPrinter() })
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]PrintUsageByPrinter, len(val))
            for i, v := range val {
                res[i] = *(v.(*PrintUsageByPrinter))
            }
            m.SetMonthlyPrintUsageSummariesByPrinter(res)
        }
        return nil
    }
    res["monthlyPrintUsageSummariesByUser"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewPrintUsageByUser() })
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]PrintUsageByUser, len(val))
            for i, v := range val {
                res[i] = *(v.(*PrintUsageByUser))
            }
            m.SetMonthlyPrintUsageSummariesByUser(res)
        }
        return nil
    }
    res["userCredentialUsageDetails"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewUserCredentialUsageDetails() })
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]UserCredentialUsageDetails, len(val))
            for i, v := range val {
                res[i] = *(v.(*UserCredentialUsageDetails))
            }
            m.SetUserCredentialUsageDetails(res)
        }
        return nil
    }
    return res
}
func (m *ReportRoot) IsNil()(bool) {
    return m == nil
}
// Serializes information the current object
// Parameters:
//  - writer : Serialization writer to use to serialize this model
func (m *ReportRoot) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetApplicationSignInDetailedSummary()))
        for i, v := range m.GetApplicationSignInDetailedSummary() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("applicationSignInDetailedSummary", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("authenticationMethods", m.GetAuthenticationMethods())
        if err != nil {
            return err
        }
    }
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetCredentialUserRegistrationDetails()))
        for i, v := range m.GetCredentialUserRegistrationDetails() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("credentialUserRegistrationDetails", cast)
        if err != nil {
            return err
        }
    }
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetDailyPrintUsageByPrinter()))
        for i, v := range m.GetDailyPrintUsageByPrinter() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("dailyPrintUsageByPrinter", cast)
        if err != nil {
            return err
        }
    }
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetDailyPrintUsageByUser()))
        for i, v := range m.GetDailyPrintUsageByUser() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("dailyPrintUsageByUser", cast)
        if err != nil {
            return err
        }
    }
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetDailyPrintUsageSummariesByPrinter()))
        for i, v := range m.GetDailyPrintUsageSummariesByPrinter() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("dailyPrintUsageSummariesByPrinter", cast)
        if err != nil {
            return err
        }
    }
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetDailyPrintUsageSummariesByUser()))
        for i, v := range m.GetDailyPrintUsageSummariesByUser() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("dailyPrintUsageSummariesByUser", cast)
        if err != nil {
            return err
        }
    }
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetMonthlyPrintUsageByPrinter()))
        for i, v := range m.GetMonthlyPrintUsageByPrinter() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("monthlyPrintUsageByPrinter", cast)
        if err != nil {
            return err
        }
    }
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetMonthlyPrintUsageByUser()))
        for i, v := range m.GetMonthlyPrintUsageByUser() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("monthlyPrintUsageByUser", cast)
        if err != nil {
            return err
        }
    }
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetMonthlyPrintUsageSummariesByPrinter()))
        for i, v := range m.GetMonthlyPrintUsageSummariesByPrinter() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("monthlyPrintUsageSummariesByPrinter", cast)
        if err != nil {
            return err
        }
    }
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetMonthlyPrintUsageSummariesByUser()))
        for i, v := range m.GetMonthlyPrintUsageSummariesByUser() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("monthlyPrintUsageSummariesByUser", cast)
        if err != nil {
            return err
        }
    }
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetUserCredentialUsageDetails()))
        for i, v := range m.GetUserCredentialUsageDetails() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("userCredentialUsageDetails", cast)
        if err != nil {
            return err
        }
    }
    return nil
}
// Sets the applicationSignInDetailedSummary property value. 
// Parameters:
//  - value : Value to set for the applicationSignInDetailedSummary property.
func (m *ReportRoot) SetApplicationSignInDetailedSummary(value []ApplicationSignInDetailedSummary)() {
    m.applicationSignInDetailedSummary = value
}
// Sets the authenticationMethods property value. 
// Parameters:
//  - value : Value to set for the authenticationMethods property.
func (m *ReportRoot) SetAuthenticationMethods(value *AuthenticationMethodsRoot)() {
    m.authenticationMethods = value
}
// Sets the credentialUserRegistrationDetails property value. 
// Parameters:
//  - value : Value to set for the credentialUserRegistrationDetails property.
func (m *ReportRoot) SetCredentialUserRegistrationDetails(value []CredentialUserRegistrationDetails)() {
    m.credentialUserRegistrationDetails = value
}
// Sets the dailyPrintUsageByPrinter property value. 
// Parameters:
//  - value : Value to set for the dailyPrintUsageByPrinter property.
func (m *ReportRoot) SetDailyPrintUsageByPrinter(value []PrintUsageByPrinter)() {
    m.dailyPrintUsageByPrinter = value
}
// Sets the dailyPrintUsageByUser property value. 
// Parameters:
//  - value : Value to set for the dailyPrintUsageByUser property.
func (m *ReportRoot) SetDailyPrintUsageByUser(value []PrintUsageByUser)() {
    m.dailyPrintUsageByUser = value
}
// Sets the dailyPrintUsageSummariesByPrinter property value. 
// Parameters:
//  - value : Value to set for the dailyPrintUsageSummariesByPrinter property.
func (m *ReportRoot) SetDailyPrintUsageSummariesByPrinter(value []PrintUsageByPrinter)() {
    m.dailyPrintUsageSummariesByPrinter = value
}
// Sets the dailyPrintUsageSummariesByUser property value. 
// Parameters:
//  - value : Value to set for the dailyPrintUsageSummariesByUser property.
func (m *ReportRoot) SetDailyPrintUsageSummariesByUser(value []PrintUsageByUser)() {
    m.dailyPrintUsageSummariesByUser = value
}
// Sets the monthlyPrintUsageByPrinter property value. 
// Parameters:
//  - value : Value to set for the monthlyPrintUsageByPrinter property.
func (m *ReportRoot) SetMonthlyPrintUsageByPrinter(value []PrintUsageByPrinter)() {
    m.monthlyPrintUsageByPrinter = value
}
// Sets the monthlyPrintUsageByUser property value. 
// Parameters:
//  - value : Value to set for the monthlyPrintUsageByUser property.
func (m *ReportRoot) SetMonthlyPrintUsageByUser(value []PrintUsageByUser)() {
    m.monthlyPrintUsageByUser = value
}
// Sets the monthlyPrintUsageSummariesByPrinter property value. 
// Parameters:
//  - value : Value to set for the monthlyPrintUsageSummariesByPrinter property.
func (m *ReportRoot) SetMonthlyPrintUsageSummariesByPrinter(value []PrintUsageByPrinter)() {
    m.monthlyPrintUsageSummariesByPrinter = value
}
// Sets the monthlyPrintUsageSummariesByUser property value. 
// Parameters:
//  - value : Value to set for the monthlyPrintUsageSummariesByUser property.
func (m *ReportRoot) SetMonthlyPrintUsageSummariesByUser(value []PrintUsageByUser)() {
    m.monthlyPrintUsageSummariesByUser = value
}
// Sets the userCredentialUsageDetails property value. 
// Parameters:
//  - value : Value to set for the userCredentialUsageDetails property.
func (m *ReportRoot) SetUserCredentialUsageDetails(value []UserCredentialUsageDetails)() {
    m.userCredentialUsageDetails = value
}
