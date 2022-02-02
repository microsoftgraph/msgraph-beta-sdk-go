package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// ReportRoot 
type ReportRoot struct {
    Entity
    // Represents a detailed summary of an application sign-in.
    applicationSignInDetailedSummary []ApplicationSignInDetailedSummary;
    // Container for navigation properties for Azure AD authentication methods resources.
    authenticationMethods *AuthenticationMethodsRoot;
    // Details of the usage of self-service password reset and multi-factor authentication (MFA) for all registered users.
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
    // Represents the self-service password reset (SSPR) usage for a given tenant.
    userCredentialUsageDetails []UserCredentialUsageDetails;
}
// NewReportRoot instantiates a new reportRoot and sets the default values.
func NewReportRoot()(*ReportRoot) {
    m := &ReportRoot{
        Entity: *NewEntity(),
    }
    return m
}
// GetApplicationSignInDetailedSummary gets the applicationSignInDetailedSummary property value. Represents a detailed summary of an application sign-in.
func (m *ReportRoot) GetApplicationSignInDetailedSummary()([]ApplicationSignInDetailedSummary) {
    if m == nil {
        return nil
    } else {
        return m.applicationSignInDetailedSummary
    }
}
// GetAuthenticationMethods gets the authenticationMethods property value. Container for navigation properties for Azure AD authentication methods resources.
func (m *ReportRoot) GetAuthenticationMethods()(*AuthenticationMethodsRoot) {
    if m == nil {
        return nil
    } else {
        return m.authenticationMethods
    }
}
// GetCredentialUserRegistrationDetails gets the credentialUserRegistrationDetails property value. Details of the usage of self-service password reset and multi-factor authentication (MFA) for all registered users.
func (m *ReportRoot) GetCredentialUserRegistrationDetails()([]CredentialUserRegistrationDetails) {
    if m == nil {
        return nil
    } else {
        return m.credentialUserRegistrationDetails
    }
}
// GetDailyPrintUsageByPrinter gets the dailyPrintUsageByPrinter property value. 
func (m *ReportRoot) GetDailyPrintUsageByPrinter()([]PrintUsageByPrinter) {
    if m == nil {
        return nil
    } else {
        return m.dailyPrintUsageByPrinter
    }
}
// GetDailyPrintUsageByUser gets the dailyPrintUsageByUser property value. 
func (m *ReportRoot) GetDailyPrintUsageByUser()([]PrintUsageByUser) {
    if m == nil {
        return nil
    } else {
        return m.dailyPrintUsageByUser
    }
}
// GetDailyPrintUsageSummariesByPrinter gets the dailyPrintUsageSummariesByPrinter property value. 
func (m *ReportRoot) GetDailyPrintUsageSummariesByPrinter()([]PrintUsageByPrinter) {
    if m == nil {
        return nil
    } else {
        return m.dailyPrintUsageSummariesByPrinter
    }
}
// GetDailyPrintUsageSummariesByUser gets the dailyPrintUsageSummariesByUser property value. 
func (m *ReportRoot) GetDailyPrintUsageSummariesByUser()([]PrintUsageByUser) {
    if m == nil {
        return nil
    } else {
        return m.dailyPrintUsageSummariesByUser
    }
}
// GetMonthlyPrintUsageByPrinter gets the monthlyPrintUsageByPrinter property value. 
func (m *ReportRoot) GetMonthlyPrintUsageByPrinter()([]PrintUsageByPrinter) {
    if m == nil {
        return nil
    } else {
        return m.monthlyPrintUsageByPrinter
    }
}
// GetMonthlyPrintUsageByUser gets the monthlyPrintUsageByUser property value. 
func (m *ReportRoot) GetMonthlyPrintUsageByUser()([]PrintUsageByUser) {
    if m == nil {
        return nil
    } else {
        return m.monthlyPrintUsageByUser
    }
}
// GetMonthlyPrintUsageSummariesByPrinter gets the monthlyPrintUsageSummariesByPrinter property value. 
func (m *ReportRoot) GetMonthlyPrintUsageSummariesByPrinter()([]PrintUsageByPrinter) {
    if m == nil {
        return nil
    } else {
        return m.monthlyPrintUsageSummariesByPrinter
    }
}
// GetMonthlyPrintUsageSummariesByUser gets the monthlyPrintUsageSummariesByUser property value. 
func (m *ReportRoot) GetMonthlyPrintUsageSummariesByUser()([]PrintUsageByUser) {
    if m == nil {
        return nil
    } else {
        return m.monthlyPrintUsageSummariesByUser
    }
}
// GetUserCredentialUsageDetails gets the userCredentialUsageDetails property value. Represents the self-service password reset (SSPR) usage for a given tenant.
func (m *ReportRoot) GetUserCredentialUsageDetails()([]UserCredentialUsageDetails) {
    if m == nil {
        return nil
    } else {
        return m.userCredentialUsageDetails
    }
}
// GetFieldDeserializers the deserialization information for the current model
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
// Serialize serializes information the current object
func (m *ReportRoot) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetApplicationSignInDetailedSummary() != nil {
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
    if m.GetCredentialUserRegistrationDetails() != nil {
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
    if m.GetDailyPrintUsageByPrinter() != nil {
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
    if m.GetDailyPrintUsageByUser() != nil {
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
    if m.GetDailyPrintUsageSummariesByPrinter() != nil {
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
    if m.GetDailyPrintUsageSummariesByUser() != nil {
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
    if m.GetMonthlyPrintUsageByPrinter() != nil {
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
    if m.GetMonthlyPrintUsageByUser() != nil {
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
    if m.GetMonthlyPrintUsageSummariesByPrinter() != nil {
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
    if m.GetMonthlyPrintUsageSummariesByUser() != nil {
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
    if m.GetUserCredentialUsageDetails() != nil {
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
// SetApplicationSignInDetailedSummary sets the applicationSignInDetailedSummary property value. Represents a detailed summary of an application sign-in.
func (m *ReportRoot) SetApplicationSignInDetailedSummary(value []ApplicationSignInDetailedSummary)() {
    if m != nil {
        m.applicationSignInDetailedSummary = value
    }
}
// SetAuthenticationMethods sets the authenticationMethods property value. Container for navigation properties for Azure AD authentication methods resources.
func (m *ReportRoot) SetAuthenticationMethods(value *AuthenticationMethodsRoot)() {
    if m != nil {
        m.authenticationMethods = value
    }
}
// SetCredentialUserRegistrationDetails sets the credentialUserRegistrationDetails property value. Details of the usage of self-service password reset and multi-factor authentication (MFA) for all registered users.
func (m *ReportRoot) SetCredentialUserRegistrationDetails(value []CredentialUserRegistrationDetails)() {
    if m != nil {
        m.credentialUserRegistrationDetails = value
    }
}
// SetDailyPrintUsageByPrinter sets the dailyPrintUsageByPrinter property value. 
func (m *ReportRoot) SetDailyPrintUsageByPrinter(value []PrintUsageByPrinter)() {
    if m != nil {
        m.dailyPrintUsageByPrinter = value
    }
}
// SetDailyPrintUsageByUser sets the dailyPrintUsageByUser property value. 
func (m *ReportRoot) SetDailyPrintUsageByUser(value []PrintUsageByUser)() {
    if m != nil {
        m.dailyPrintUsageByUser = value
    }
}
// SetDailyPrintUsageSummariesByPrinter sets the dailyPrintUsageSummariesByPrinter property value. 
func (m *ReportRoot) SetDailyPrintUsageSummariesByPrinter(value []PrintUsageByPrinter)() {
    if m != nil {
        m.dailyPrintUsageSummariesByPrinter = value
    }
}
// SetDailyPrintUsageSummariesByUser sets the dailyPrintUsageSummariesByUser property value. 
func (m *ReportRoot) SetDailyPrintUsageSummariesByUser(value []PrintUsageByUser)() {
    if m != nil {
        m.dailyPrintUsageSummariesByUser = value
    }
}
// SetMonthlyPrintUsageByPrinter sets the monthlyPrintUsageByPrinter property value. 
func (m *ReportRoot) SetMonthlyPrintUsageByPrinter(value []PrintUsageByPrinter)() {
    if m != nil {
        m.monthlyPrintUsageByPrinter = value
    }
}
// SetMonthlyPrintUsageByUser sets the monthlyPrintUsageByUser property value. 
func (m *ReportRoot) SetMonthlyPrintUsageByUser(value []PrintUsageByUser)() {
    if m != nil {
        m.monthlyPrintUsageByUser = value
    }
}
// SetMonthlyPrintUsageSummariesByPrinter sets the monthlyPrintUsageSummariesByPrinter property value. 
func (m *ReportRoot) SetMonthlyPrintUsageSummariesByPrinter(value []PrintUsageByPrinter)() {
    if m != nil {
        m.monthlyPrintUsageSummariesByPrinter = value
    }
}
// SetMonthlyPrintUsageSummariesByUser sets the monthlyPrintUsageSummariesByUser property value. 
func (m *ReportRoot) SetMonthlyPrintUsageSummariesByUser(value []PrintUsageByUser)() {
    if m != nil {
        m.monthlyPrintUsageSummariesByUser = value
    }
}
// SetUserCredentialUsageDetails sets the userCredentialUsageDetails property value. Represents the self-service password reset (SSPR) usage for a given tenant.
func (m *ReportRoot) SetUserCredentialUsageDetails(value []UserCredentialUsageDetails)() {
    if m != nil {
        m.userCredentialUsageDetails = value
    }
}
