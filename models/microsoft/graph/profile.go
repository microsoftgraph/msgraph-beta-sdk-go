package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

type Profile struct {
    Entity
    account []UserAccountInformation;
    addresses []ItemAddress;
    anniversaries []PersonAnnualEvent;
    awards []PersonAward;
    certifications []PersonCertification;
    educationalActivities []EducationalActivity;
    emails []ItemEmail;
    interests []PersonInterest;
    languages []LanguageProficiency;
    names []PersonName;
    notes []PersonAnnotation;
    patents []ItemPatent;
    phones []ItemPhone;
    positions []WorkPosition;
    projects []ProjectParticipation;
    publications []ItemPublication;
    skills []SkillProficiency;
    webAccounts []WebAccount;
    websites []PersonWebsite;
}
func NewProfile()(*Profile) {
    m := &Profile{
        Entity: *NewEntity(),
    }
    return m
}
func (m *Profile) GetAccount()([]UserAccountInformation) {
    if m == nil {
        return nil
    } else {
        return m.account
    }
}
func (m *Profile) GetAddresses()([]ItemAddress) {
    if m == nil {
        return nil
    } else {
        return m.addresses
    }
}
func (m *Profile) GetAnniversaries()([]PersonAnnualEvent) {
    if m == nil {
        return nil
    } else {
        return m.anniversaries
    }
}
func (m *Profile) GetAwards()([]PersonAward) {
    if m == nil {
        return nil
    } else {
        return m.awards
    }
}
func (m *Profile) GetCertifications()([]PersonCertification) {
    if m == nil {
        return nil
    } else {
        return m.certifications
    }
}
func (m *Profile) GetEducationalActivities()([]EducationalActivity) {
    if m == nil {
        return nil
    } else {
        return m.educationalActivities
    }
}
func (m *Profile) GetEmails()([]ItemEmail) {
    if m == nil {
        return nil
    } else {
        return m.emails
    }
}
func (m *Profile) GetInterests()([]PersonInterest) {
    if m == nil {
        return nil
    } else {
        return m.interests
    }
}
func (m *Profile) GetLanguages()([]LanguageProficiency) {
    if m == nil {
        return nil
    } else {
        return m.languages
    }
}
func (m *Profile) GetNames()([]PersonName) {
    if m == nil {
        return nil
    } else {
        return m.names
    }
}
func (m *Profile) GetNotes()([]PersonAnnotation) {
    if m == nil {
        return nil
    } else {
        return m.notes
    }
}
func (m *Profile) GetPatents()([]ItemPatent) {
    if m == nil {
        return nil
    } else {
        return m.patents
    }
}
func (m *Profile) GetPhones()([]ItemPhone) {
    if m == nil {
        return nil
    } else {
        return m.phones
    }
}
func (m *Profile) GetPositions()([]WorkPosition) {
    if m == nil {
        return nil
    } else {
        return m.positions
    }
}
func (m *Profile) GetProjects()([]ProjectParticipation) {
    if m == nil {
        return nil
    } else {
        return m.projects
    }
}
func (m *Profile) GetPublications()([]ItemPublication) {
    if m == nil {
        return nil
    } else {
        return m.publications
    }
}
func (m *Profile) GetSkills()([]SkillProficiency) {
    if m == nil {
        return nil
    } else {
        return m.skills
    }
}
func (m *Profile) GetWebAccounts()([]WebAccount) {
    if m == nil {
        return nil
    } else {
        return m.webAccounts
    }
}
func (m *Profile) GetWebsites()([]PersonWebsite) {
    if m == nil {
        return nil
    } else {
        return m.websites
    }
}
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
func (m *Profile) SetAccount(value []UserAccountInformation)() {
    m.account = value
}
func (m *Profile) SetAddresses(value []ItemAddress)() {
    m.addresses = value
}
func (m *Profile) SetAnniversaries(value []PersonAnnualEvent)() {
    m.anniversaries = value
}
func (m *Profile) SetAwards(value []PersonAward)() {
    m.awards = value
}
func (m *Profile) SetCertifications(value []PersonCertification)() {
    m.certifications = value
}
func (m *Profile) SetEducationalActivities(value []EducationalActivity)() {
    m.educationalActivities = value
}
func (m *Profile) SetEmails(value []ItemEmail)() {
    m.emails = value
}
func (m *Profile) SetInterests(value []PersonInterest)() {
    m.interests = value
}
func (m *Profile) SetLanguages(value []LanguageProficiency)() {
    m.languages = value
}
func (m *Profile) SetNames(value []PersonName)() {
    m.names = value
}
func (m *Profile) SetNotes(value []PersonAnnotation)() {
    m.notes = value
}
func (m *Profile) SetPatents(value []ItemPatent)() {
    m.patents = value
}
func (m *Profile) SetPhones(value []ItemPhone)() {
    m.phones = value
}
func (m *Profile) SetPositions(value []WorkPosition)() {
    m.positions = value
}
func (m *Profile) SetProjects(value []ProjectParticipation)() {
    m.projects = value
}
func (m *Profile) SetPublications(value []ItemPublication)() {
    m.publications = value
}
func (m *Profile) SetSkills(value []SkillProficiency)() {
    m.skills = value
}
func (m *Profile) SetWebAccounts(value []WebAccount)() {
    m.webAccounts = value
}
func (m *Profile) SetWebsites(value []PersonWebsite)() {
    m.websites = value
}
