package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// Profile 
type Profile struct {
    Entity
    // 
    account []UserAccountInformation;
    // Represents details of addresses associated with the user.
    addresses []ItemAddress;
    // Represents the details of meaningful dates associated with a person.
    anniversaries []PersonAnnualEvent;
    // Represents the details of awards or honors associated with a person.
    awards []PersonAward;
    // Represents the details of certifications associated with a person.
    certifications []PersonCertification;
    // Represents data that a user has supplied related to undergraduate, graduate, postgraduate or other educational activities.
    educationalActivities []EducationalActivity;
    // Represents detailed information about email addresses associated with the user.
    emails []ItemEmail;
    // Provides detailed information about interests the user has associated with themselves in various services.
    interests []PersonInterest;
    // Represents detailed information about languages that a user has added to their profile.
    languages []LanguageProficiency;
    // Represents the names a user has added to their profile.
    names []PersonName;
    // Represents notes that a user has added to their profile.
    notes []PersonAnnotation;
    // Represents patents that a user has added to their profile.
    patents []ItemPatent;
    // Represents detailed information about phone numbers associated with a user in various services.
    phones []ItemPhone;
    // Represents detailed information about work positions associated with a user's profile.
    positions []WorkPosition;
    // Represents detailed information about projects associated with a user.
    projects []ProjectParticipation;
    // Represents details of any publications a user has added to their profile.
    publications []ItemPublication;
    // Represents detailed information about skills associated with a user in various services.
    skills []SkillProficiency;
    // Represents web accounts the user has indicated they use or has added to their user profile.
    webAccounts []WebAccount;
    // Represents detailed information about websites associated with a user in various services.
    websites []PersonWebsite;
}
// NewProfile instantiates a new profile and sets the default values.
func NewProfile()(*Profile) {
    m := &Profile{
        Entity: *NewEntity(),
    }
    return m
}
// GetAccount gets the account property value. 
func (m *Profile) GetAccount()([]UserAccountInformation) {
    if m == nil {
        return nil
    } else {
        return m.account
    }
}
// GetAddresses gets the addresses property value. Represents details of addresses associated with the user.
func (m *Profile) GetAddresses()([]ItemAddress) {
    if m == nil {
        return nil
    } else {
        return m.addresses
    }
}
// GetAnniversaries gets the anniversaries property value. Represents the details of meaningful dates associated with a person.
func (m *Profile) GetAnniversaries()([]PersonAnnualEvent) {
    if m == nil {
        return nil
    } else {
        return m.anniversaries
    }
}
// GetAwards gets the awards property value. Represents the details of awards or honors associated with a person.
func (m *Profile) GetAwards()([]PersonAward) {
    if m == nil {
        return nil
    } else {
        return m.awards
    }
}
// GetCertifications gets the certifications property value. Represents the details of certifications associated with a person.
func (m *Profile) GetCertifications()([]PersonCertification) {
    if m == nil {
        return nil
    } else {
        return m.certifications
    }
}
// GetEducationalActivities gets the educationalActivities property value. Represents data that a user has supplied related to undergraduate, graduate, postgraduate or other educational activities.
func (m *Profile) GetEducationalActivities()([]EducationalActivity) {
    if m == nil {
        return nil
    } else {
        return m.educationalActivities
    }
}
// GetEmails gets the emails property value. Represents detailed information about email addresses associated with the user.
func (m *Profile) GetEmails()([]ItemEmail) {
    if m == nil {
        return nil
    } else {
        return m.emails
    }
}
// GetInterests gets the interests property value. Provides detailed information about interests the user has associated with themselves in various services.
func (m *Profile) GetInterests()([]PersonInterest) {
    if m == nil {
        return nil
    } else {
        return m.interests
    }
}
// GetLanguages gets the languages property value. Represents detailed information about languages that a user has added to their profile.
func (m *Profile) GetLanguages()([]LanguageProficiency) {
    if m == nil {
        return nil
    } else {
        return m.languages
    }
}
// GetNames gets the names property value. Represents the names a user has added to their profile.
func (m *Profile) GetNames()([]PersonName) {
    if m == nil {
        return nil
    } else {
        return m.names
    }
}
// GetNotes gets the notes property value. Represents notes that a user has added to their profile.
func (m *Profile) GetNotes()([]PersonAnnotation) {
    if m == nil {
        return nil
    } else {
        return m.notes
    }
}
// GetPatents gets the patents property value. Represents patents that a user has added to their profile.
func (m *Profile) GetPatents()([]ItemPatent) {
    if m == nil {
        return nil
    } else {
        return m.patents
    }
}
// GetPhones gets the phones property value. Represents detailed information about phone numbers associated with a user in various services.
func (m *Profile) GetPhones()([]ItemPhone) {
    if m == nil {
        return nil
    } else {
        return m.phones
    }
}
// GetPositions gets the positions property value. Represents detailed information about work positions associated with a user's profile.
func (m *Profile) GetPositions()([]WorkPosition) {
    if m == nil {
        return nil
    } else {
        return m.positions
    }
}
// GetProjects gets the projects property value. Represents detailed information about projects associated with a user.
func (m *Profile) GetProjects()([]ProjectParticipation) {
    if m == nil {
        return nil
    } else {
        return m.projects
    }
}
// GetPublications gets the publications property value. Represents details of any publications a user has added to their profile.
func (m *Profile) GetPublications()([]ItemPublication) {
    if m == nil {
        return nil
    } else {
        return m.publications
    }
}
// GetSkills gets the skills property value. Represents detailed information about skills associated with a user in various services.
func (m *Profile) GetSkills()([]SkillProficiency) {
    if m == nil {
        return nil
    } else {
        return m.skills
    }
}
// GetWebAccounts gets the webAccounts property value. Represents web accounts the user has indicated they use or has added to their user profile.
func (m *Profile) GetWebAccounts()([]WebAccount) {
    if m == nil {
        return nil
    } else {
        return m.webAccounts
    }
}
// GetWebsites gets the websites property value. Represents detailed information about websites associated with a user in various services.
func (m *Profile) GetWebsites()([]PersonWebsite) {
    if m == nil {
        return nil
    } else {
        return m.websites
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *Profile) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["account"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewUserAccountInformation() })
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]UserAccountInformation, len(val))
            for i, v := range val {
                res[i] = *(v.(*UserAccountInformation))
            }
            m.SetAccount(res)
        }
        return nil
    }
    res["addresses"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewItemAddress() })
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]ItemAddress, len(val))
            for i, v := range val {
                res[i] = *(v.(*ItemAddress))
            }
            m.SetAddresses(res)
        }
        return nil
    }
    res["anniversaries"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewPersonAnnualEvent() })
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]PersonAnnualEvent, len(val))
            for i, v := range val {
                res[i] = *(v.(*PersonAnnualEvent))
            }
            m.SetAnniversaries(res)
        }
        return nil
    }
    res["awards"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewPersonAward() })
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]PersonAward, len(val))
            for i, v := range val {
                res[i] = *(v.(*PersonAward))
            }
            m.SetAwards(res)
        }
        return nil
    }
    res["certifications"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewPersonCertification() })
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]PersonCertification, len(val))
            for i, v := range val {
                res[i] = *(v.(*PersonCertification))
            }
            m.SetCertifications(res)
        }
        return nil
    }
    res["educationalActivities"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewEducationalActivity() })
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]EducationalActivity, len(val))
            for i, v := range val {
                res[i] = *(v.(*EducationalActivity))
            }
            m.SetEducationalActivities(res)
        }
        return nil
    }
    res["emails"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewItemEmail() })
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]ItemEmail, len(val))
            for i, v := range val {
                res[i] = *(v.(*ItemEmail))
            }
            m.SetEmails(res)
        }
        return nil
    }
    res["interests"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewPersonInterest() })
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]PersonInterest, len(val))
            for i, v := range val {
                res[i] = *(v.(*PersonInterest))
            }
            m.SetInterests(res)
        }
        return nil
    }
    res["languages"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewLanguageProficiency() })
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]LanguageProficiency, len(val))
            for i, v := range val {
                res[i] = *(v.(*LanguageProficiency))
            }
            m.SetLanguages(res)
        }
        return nil
    }
    res["names"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewPersonName() })
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]PersonName, len(val))
            for i, v := range val {
                res[i] = *(v.(*PersonName))
            }
            m.SetNames(res)
        }
        return nil
    }
    res["notes"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewPersonAnnotation() })
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]PersonAnnotation, len(val))
            for i, v := range val {
                res[i] = *(v.(*PersonAnnotation))
            }
            m.SetNotes(res)
        }
        return nil
    }
    res["patents"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewItemPatent() })
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]ItemPatent, len(val))
            for i, v := range val {
                res[i] = *(v.(*ItemPatent))
            }
            m.SetPatents(res)
        }
        return nil
    }
    res["phones"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewItemPhone() })
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]ItemPhone, len(val))
            for i, v := range val {
                res[i] = *(v.(*ItemPhone))
            }
            m.SetPhones(res)
        }
        return nil
    }
    res["positions"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewWorkPosition() })
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]WorkPosition, len(val))
            for i, v := range val {
                res[i] = *(v.(*WorkPosition))
            }
            m.SetPositions(res)
        }
        return nil
    }
    res["projects"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewProjectParticipation() })
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]ProjectParticipation, len(val))
            for i, v := range val {
                res[i] = *(v.(*ProjectParticipation))
            }
            m.SetProjects(res)
        }
        return nil
    }
    res["publications"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewItemPublication() })
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]ItemPublication, len(val))
            for i, v := range val {
                res[i] = *(v.(*ItemPublication))
            }
            m.SetPublications(res)
        }
        return nil
    }
    res["skills"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewSkillProficiency() })
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]SkillProficiency, len(val))
            for i, v := range val {
                res[i] = *(v.(*SkillProficiency))
            }
            m.SetSkills(res)
        }
        return nil
    }
    res["webAccounts"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewWebAccount() })
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]WebAccount, len(val))
            for i, v := range val {
                res[i] = *(v.(*WebAccount))
            }
            m.SetWebAccounts(res)
        }
        return nil
    }
    res["websites"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewPersonWebsite() })
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]PersonWebsite, len(val))
            for i, v := range val {
                res[i] = *(v.(*PersonWebsite))
            }
            m.SetWebsites(res)
        }
        return nil
    }
    return res
}
func (m *Profile) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
func (m *Profile) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetAccount() != nil {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetAccount()))
        for i, v := range m.GetAccount() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("account", cast)
        if err != nil {
            return err
        }
    }
    if m.GetAddresses() != nil {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetAddresses()))
        for i, v := range m.GetAddresses() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("addresses", cast)
        if err != nil {
            return err
        }
    }
    if m.GetAnniversaries() != nil {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetAnniversaries()))
        for i, v := range m.GetAnniversaries() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("anniversaries", cast)
        if err != nil {
            return err
        }
    }
    if m.GetAwards() != nil {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetAwards()))
        for i, v := range m.GetAwards() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("awards", cast)
        if err != nil {
            return err
        }
    }
    if m.GetCertifications() != nil {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetCertifications()))
        for i, v := range m.GetCertifications() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("certifications", cast)
        if err != nil {
            return err
        }
    }
    if m.GetEducationalActivities() != nil {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetEducationalActivities()))
        for i, v := range m.GetEducationalActivities() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("educationalActivities", cast)
        if err != nil {
            return err
        }
    }
    if m.GetEmails() != nil {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetEmails()))
        for i, v := range m.GetEmails() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("emails", cast)
        if err != nil {
            return err
        }
    }
    if m.GetInterests() != nil {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetInterests()))
        for i, v := range m.GetInterests() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("interests", cast)
        if err != nil {
            return err
        }
    }
    if m.GetLanguages() != nil {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetLanguages()))
        for i, v := range m.GetLanguages() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("languages", cast)
        if err != nil {
            return err
        }
    }
    if m.GetNames() != nil {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetNames()))
        for i, v := range m.GetNames() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("names", cast)
        if err != nil {
            return err
        }
    }
    if m.GetNotes() != nil {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetNotes()))
        for i, v := range m.GetNotes() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("notes", cast)
        if err != nil {
            return err
        }
    }
    if m.GetPatents() != nil {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetPatents()))
        for i, v := range m.GetPatents() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("patents", cast)
        if err != nil {
            return err
        }
    }
    if m.GetPhones() != nil {
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
    if m.GetPositions() != nil {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetPositions()))
        for i, v := range m.GetPositions() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("positions", cast)
        if err != nil {
            return err
        }
    }
    if m.GetProjects() != nil {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetProjects()))
        for i, v := range m.GetProjects() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("projects", cast)
        if err != nil {
            return err
        }
    }
    if m.GetPublications() != nil {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetPublications()))
        for i, v := range m.GetPublications() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("publications", cast)
        if err != nil {
            return err
        }
    }
    if m.GetSkills() != nil {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetSkills()))
        for i, v := range m.GetSkills() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("skills", cast)
        if err != nil {
            return err
        }
    }
    if m.GetWebAccounts() != nil {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetWebAccounts()))
        for i, v := range m.GetWebAccounts() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("webAccounts", cast)
        if err != nil {
            return err
        }
    }
    if m.GetWebsites() != nil {
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
    return nil
}
// SetAccount sets the account property value. 
func (m *Profile) SetAccount(value []UserAccountInformation)() {
    if m != nil {
        m.account = value
    }
}
// SetAddresses sets the addresses property value. Represents details of addresses associated with the user.
func (m *Profile) SetAddresses(value []ItemAddress)() {
    if m != nil {
        m.addresses = value
    }
}
// SetAnniversaries sets the anniversaries property value. Represents the details of meaningful dates associated with a person.
func (m *Profile) SetAnniversaries(value []PersonAnnualEvent)() {
    if m != nil {
        m.anniversaries = value
    }
}
// SetAwards sets the awards property value. Represents the details of awards or honors associated with a person.
func (m *Profile) SetAwards(value []PersonAward)() {
    if m != nil {
        m.awards = value
    }
}
// SetCertifications sets the certifications property value. Represents the details of certifications associated with a person.
func (m *Profile) SetCertifications(value []PersonCertification)() {
    if m != nil {
        m.certifications = value
    }
}
// SetEducationalActivities sets the educationalActivities property value. Represents data that a user has supplied related to undergraduate, graduate, postgraduate or other educational activities.
func (m *Profile) SetEducationalActivities(value []EducationalActivity)() {
    if m != nil {
        m.educationalActivities = value
    }
}
// SetEmails sets the emails property value. Represents detailed information about email addresses associated with the user.
func (m *Profile) SetEmails(value []ItemEmail)() {
    if m != nil {
        m.emails = value
    }
}
// SetInterests sets the interests property value. Provides detailed information about interests the user has associated with themselves in various services.
func (m *Profile) SetInterests(value []PersonInterest)() {
    if m != nil {
        m.interests = value
    }
}
// SetLanguages sets the languages property value. Represents detailed information about languages that a user has added to their profile.
func (m *Profile) SetLanguages(value []LanguageProficiency)() {
    if m != nil {
        m.languages = value
    }
}
// SetNames sets the names property value. Represents the names a user has added to their profile.
func (m *Profile) SetNames(value []PersonName)() {
    if m != nil {
        m.names = value
    }
}
// SetNotes sets the notes property value. Represents notes that a user has added to their profile.
func (m *Profile) SetNotes(value []PersonAnnotation)() {
    if m != nil {
        m.notes = value
    }
}
// SetPatents sets the patents property value. Represents patents that a user has added to their profile.
func (m *Profile) SetPatents(value []ItemPatent)() {
    if m != nil {
        m.patents = value
    }
}
// SetPhones sets the phones property value. Represents detailed information about phone numbers associated with a user in various services.
func (m *Profile) SetPhones(value []ItemPhone)() {
    if m != nil {
        m.phones = value
    }
}
// SetPositions sets the positions property value. Represents detailed information about work positions associated with a user's profile.
func (m *Profile) SetPositions(value []WorkPosition)() {
    if m != nil {
        m.positions = value
    }
}
// SetProjects sets the projects property value. Represents detailed information about projects associated with a user.
func (m *Profile) SetProjects(value []ProjectParticipation)() {
    if m != nil {
        m.projects = value
    }
}
// SetPublications sets the publications property value. Represents details of any publications a user has added to their profile.
func (m *Profile) SetPublications(value []ItemPublication)() {
    if m != nil {
        m.publications = value
    }
}
// SetSkills sets the skills property value. Represents detailed information about skills associated with a user in various services.
func (m *Profile) SetSkills(value []SkillProficiency)() {
    if m != nil {
        m.skills = value
    }
}
// SetWebAccounts sets the webAccounts property value. Represents web accounts the user has indicated they use or has added to their user profile.
func (m *Profile) SetWebAccounts(value []WebAccount)() {
    if m != nil {
        m.webAccounts = value
    }
}
// SetWebsites sets the websites property value. Represents detailed information about websites associated with a user in various services.
func (m *Profile) SetWebsites(value []PersonWebsite)() {
    if m != nil {
        m.websites = value
    }
}
