package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// 
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
// Instantiates a new profile and sets the default values.
func NewProfile()(*Profile) {
    m := &Profile{
        Entity: *NewEntity(),
    }
    return m
}
// Gets the account property value. 
func (m *Profile) GetAccount()([]UserAccountInformation) {
    if m == nil {
        return nil
    } else {
        return m.account
    }
}
// Gets the addresses property value. Represents details of addresses associated with the user.
func (m *Profile) GetAddresses()([]ItemAddress) {
    if m == nil {
        return nil
    } else {
        return m.addresses
    }
}
// Gets the anniversaries property value. Represents the details of meaningful dates associated with a person.
func (m *Profile) GetAnniversaries()([]PersonAnnualEvent) {
    if m == nil {
        return nil
    } else {
        return m.anniversaries
    }
}
// Gets the awards property value. Represents the details of awards or honors associated with a person.
func (m *Profile) GetAwards()([]PersonAward) {
    if m == nil {
        return nil
    } else {
        return m.awards
    }
}
// Gets the certifications property value. Represents the details of certifications associated with a person.
func (m *Profile) GetCertifications()([]PersonCertification) {
    if m == nil {
        return nil
    } else {
        return m.certifications
    }
}
// Gets the educationalActivities property value. Represents data that a user has supplied related to undergraduate, graduate, postgraduate or other educational activities.
func (m *Profile) GetEducationalActivities()([]EducationalActivity) {
    if m == nil {
        return nil
    } else {
        return m.educationalActivities
    }
}
// Gets the emails property value. Represents detailed information about email addresses associated with the user.
func (m *Profile) GetEmails()([]ItemEmail) {
    if m == nil {
        return nil
    } else {
        return m.emails
    }
}
// Gets the interests property value. Provides detailed information about interests the user has associated with themselves in various services.
func (m *Profile) GetInterests()([]PersonInterest) {
    if m == nil {
        return nil
    } else {
        return m.interests
    }
}
// Gets the languages property value. Represents detailed information about languages that a user has added to their profile.
func (m *Profile) GetLanguages()([]LanguageProficiency) {
    if m == nil {
        return nil
    } else {
        return m.languages
    }
}
// Gets the names property value. Represents the names a user has added to their profile.
func (m *Profile) GetNames()([]PersonName) {
    if m == nil {
        return nil
    } else {
        return m.names
    }
}
// Gets the notes property value. Represents notes that a user has added to their profile.
func (m *Profile) GetNotes()([]PersonAnnotation) {
    if m == nil {
        return nil
    } else {
        return m.notes
    }
}
// Gets the patents property value. Represents patents that a user has added to their profile.
func (m *Profile) GetPatents()([]ItemPatent) {
    if m == nil {
        return nil
    } else {
        return m.patents
    }
}
// Gets the phones property value. Represents detailed information about phone numbers associated with a user in various services.
func (m *Profile) GetPhones()([]ItemPhone) {
    if m == nil {
        return nil
    } else {
        return m.phones
    }
}
// Gets the positions property value. Represents detailed information about work positions associated with a user's profile.
func (m *Profile) GetPositions()([]WorkPosition) {
    if m == nil {
        return nil
    } else {
        return m.positions
    }
}
// Gets the projects property value. Represents detailed information about projects associated with a user.
func (m *Profile) GetProjects()([]ProjectParticipation) {
    if m == nil {
        return nil
    } else {
        return m.projects
    }
}
// Gets the publications property value. Represents details of any publications a user has added to their profile.
func (m *Profile) GetPublications()([]ItemPublication) {
    if m == nil {
        return nil
    } else {
        return m.publications
    }
}
// Gets the skills property value. Represents detailed information about skills associated with a user in various services.
func (m *Profile) GetSkills()([]SkillProficiency) {
    if m == nil {
        return nil
    } else {
        return m.skills
    }
}
// Gets the webAccounts property value. Represents web accounts the user has indicated they use or has added to their user profile.
func (m *Profile) GetWebAccounts()([]WebAccount) {
    if m == nil {
        return nil
    } else {
        return m.webAccounts
    }
}
// Gets the websites property value. Represents detailed information about websites associated with a user in various services.
func (m *Profile) GetWebsites()([]PersonWebsite) {
    if m == nil {
        return nil
    } else {
        return m.websites
    }
}
// The deserialization information for the current model
func (m *Profile) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["account"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewUserAccountInformation() })
        if err != nil {
            return err
        }
        res := make([]UserAccountInformation, len(val))
        for i, v := range val {
            res[i] = *(v.(*UserAccountInformation))
        }
        m.SetAccount(res)
        return nil
    }
    res["addresses"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewItemAddress() })
        if err != nil {
            return err
        }
        res := make([]ItemAddress, len(val))
        for i, v := range val {
            res[i] = *(v.(*ItemAddress))
        }
        m.SetAddresses(res)
        return nil
    }
    res["anniversaries"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewPersonAnnualEvent() })
        if err != nil {
            return err
        }
        res := make([]PersonAnnualEvent, len(val))
        for i, v := range val {
            res[i] = *(v.(*PersonAnnualEvent))
        }
        m.SetAnniversaries(res)
        return nil
    }
    res["awards"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewPersonAward() })
        if err != nil {
            return err
        }
        res := make([]PersonAward, len(val))
        for i, v := range val {
            res[i] = *(v.(*PersonAward))
        }
        m.SetAwards(res)
        return nil
    }
    res["certifications"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewPersonCertification() })
        if err != nil {
            return err
        }
        res := make([]PersonCertification, len(val))
        for i, v := range val {
            res[i] = *(v.(*PersonCertification))
        }
        m.SetCertifications(res)
        return nil
    }
    res["educationalActivities"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewEducationalActivity() })
        if err != nil {
            return err
        }
        res := make([]EducationalActivity, len(val))
        for i, v := range val {
            res[i] = *(v.(*EducationalActivity))
        }
        m.SetEducationalActivities(res)
        return nil
    }
    res["emails"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewItemEmail() })
        if err != nil {
            return err
        }
        res := make([]ItemEmail, len(val))
        for i, v := range val {
            res[i] = *(v.(*ItemEmail))
        }
        m.SetEmails(res)
        return nil
    }
    res["interests"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewPersonInterest() })
        if err != nil {
            return err
        }
        res := make([]PersonInterest, len(val))
        for i, v := range val {
            res[i] = *(v.(*PersonInterest))
        }
        m.SetInterests(res)
        return nil
    }
    res["languages"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewLanguageProficiency() })
        if err != nil {
            return err
        }
        res := make([]LanguageProficiency, len(val))
        for i, v := range val {
            res[i] = *(v.(*LanguageProficiency))
        }
        m.SetLanguages(res)
        return nil
    }
    res["names"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewPersonName() })
        if err != nil {
            return err
        }
        res := make([]PersonName, len(val))
        for i, v := range val {
            res[i] = *(v.(*PersonName))
        }
        m.SetNames(res)
        return nil
    }
    res["notes"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewPersonAnnotation() })
        if err != nil {
            return err
        }
        res := make([]PersonAnnotation, len(val))
        for i, v := range val {
            res[i] = *(v.(*PersonAnnotation))
        }
        m.SetNotes(res)
        return nil
    }
    res["patents"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewItemPatent() })
        if err != nil {
            return err
        }
        res := make([]ItemPatent, len(val))
        for i, v := range val {
            res[i] = *(v.(*ItemPatent))
        }
        m.SetPatents(res)
        return nil
    }
    res["phones"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewItemPhone() })
        if err != nil {
            return err
        }
        res := make([]ItemPhone, len(val))
        for i, v := range val {
            res[i] = *(v.(*ItemPhone))
        }
        m.SetPhones(res)
        return nil
    }
    res["positions"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewWorkPosition() })
        if err != nil {
            return err
        }
        res := make([]WorkPosition, len(val))
        for i, v := range val {
            res[i] = *(v.(*WorkPosition))
        }
        m.SetPositions(res)
        return nil
    }
    res["projects"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewProjectParticipation() })
        if err != nil {
            return err
        }
        res := make([]ProjectParticipation, len(val))
        for i, v := range val {
            res[i] = *(v.(*ProjectParticipation))
        }
        m.SetProjects(res)
        return nil
    }
    res["publications"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewItemPublication() })
        if err != nil {
            return err
        }
        res := make([]ItemPublication, len(val))
        for i, v := range val {
            res[i] = *(v.(*ItemPublication))
        }
        m.SetPublications(res)
        return nil
    }
    res["skills"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewSkillProficiency() })
        if err != nil {
            return err
        }
        res := make([]SkillProficiency, len(val))
        for i, v := range val {
            res[i] = *(v.(*SkillProficiency))
        }
        m.SetSkills(res)
        return nil
    }
    res["webAccounts"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewWebAccount() })
        if err != nil {
            return err
        }
        res := make([]WebAccount, len(val))
        for i, v := range val {
            res[i] = *(v.(*WebAccount))
        }
        m.SetWebAccounts(res)
        return nil
    }
    res["websites"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewPersonWebsite() })
        if err != nil {
            return err
        }
        res := make([]PersonWebsite, len(val))
        for i, v := range val {
            res[i] = *(v.(*PersonWebsite))
        }
        m.SetWebsites(res)
        return nil
    }
    return res
}
func (m *Profile) IsNil()(bool) {
    return m == nil
}
// Serializes information the current object
// Parameters:
//  - writer : Serialization writer to use to serialize this model
func (m *Profile) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
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
    {
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
    {
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
    {
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
    {
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
    {
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
    {
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
    {
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
    {
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
    {
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
    {
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
    {
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
    {
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
    {
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
    {
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
    {
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
    return nil
}
// Sets the account property value. 
// Parameters:
//  - value : Value to set for the account property.
func (m *Profile) SetAccount(value []UserAccountInformation)() {
    m.account = value
}
// Sets the addresses property value. Represents details of addresses associated with the user.
// Parameters:
//  - value : Value to set for the addresses property.
func (m *Profile) SetAddresses(value []ItemAddress)() {
    m.addresses = value
}
// Sets the anniversaries property value. Represents the details of meaningful dates associated with a person.
// Parameters:
//  - value : Value to set for the anniversaries property.
func (m *Profile) SetAnniversaries(value []PersonAnnualEvent)() {
    m.anniversaries = value
}
// Sets the awards property value. Represents the details of awards or honors associated with a person.
// Parameters:
//  - value : Value to set for the awards property.
func (m *Profile) SetAwards(value []PersonAward)() {
    m.awards = value
}
// Sets the certifications property value. Represents the details of certifications associated with a person.
// Parameters:
//  - value : Value to set for the certifications property.
func (m *Profile) SetCertifications(value []PersonCertification)() {
    m.certifications = value
}
// Sets the educationalActivities property value. Represents data that a user has supplied related to undergraduate, graduate, postgraduate or other educational activities.
// Parameters:
//  - value : Value to set for the educationalActivities property.
func (m *Profile) SetEducationalActivities(value []EducationalActivity)() {
    m.educationalActivities = value
}
// Sets the emails property value. Represents detailed information about email addresses associated with the user.
// Parameters:
//  - value : Value to set for the emails property.
func (m *Profile) SetEmails(value []ItemEmail)() {
    m.emails = value
}
// Sets the interests property value. Provides detailed information about interests the user has associated with themselves in various services.
// Parameters:
//  - value : Value to set for the interests property.
func (m *Profile) SetInterests(value []PersonInterest)() {
    m.interests = value
}
// Sets the languages property value. Represents detailed information about languages that a user has added to their profile.
// Parameters:
//  - value : Value to set for the languages property.
func (m *Profile) SetLanguages(value []LanguageProficiency)() {
    m.languages = value
}
// Sets the names property value. Represents the names a user has added to their profile.
// Parameters:
//  - value : Value to set for the names property.
func (m *Profile) SetNames(value []PersonName)() {
    m.names = value
}
// Sets the notes property value. Represents notes that a user has added to their profile.
// Parameters:
//  - value : Value to set for the notes property.
func (m *Profile) SetNotes(value []PersonAnnotation)() {
    m.notes = value
}
// Sets the patents property value. Represents patents that a user has added to their profile.
// Parameters:
//  - value : Value to set for the patents property.
func (m *Profile) SetPatents(value []ItemPatent)() {
    m.patents = value
}
// Sets the phones property value. Represents detailed information about phone numbers associated with a user in various services.
// Parameters:
//  - value : Value to set for the phones property.
func (m *Profile) SetPhones(value []ItemPhone)() {
    m.phones = value
}
// Sets the positions property value. Represents detailed information about work positions associated with a user's profile.
// Parameters:
//  - value : Value to set for the positions property.
func (m *Profile) SetPositions(value []WorkPosition)() {
    m.positions = value
}
// Sets the projects property value. Represents detailed information about projects associated with a user.
// Parameters:
//  - value : Value to set for the projects property.
func (m *Profile) SetProjects(value []ProjectParticipation)() {
    m.projects = value
}
// Sets the publications property value. Represents details of any publications a user has added to their profile.
// Parameters:
//  - value : Value to set for the publications property.
func (m *Profile) SetPublications(value []ItemPublication)() {
    m.publications = value
}
// Sets the skills property value. Represents detailed information about skills associated with a user in various services.
// Parameters:
//  - value : Value to set for the skills property.
func (m *Profile) SetSkills(value []SkillProficiency)() {
    m.skills = value
}
// Sets the webAccounts property value. Represents web accounts the user has indicated they use or has added to their user profile.
// Parameters:
//  - value : Value to set for the webAccounts property.
func (m *Profile) SetWebAccounts(value []WebAccount)() {
    m.webAccounts = value
}
// Sets the websites property value. Represents detailed information about websites associated with a user in various services.
// Parameters:
//  - value : Value to set for the websites property.
func (m *Profile) SetWebsites(value []PersonWebsite)() {
    m.websites = value
}
