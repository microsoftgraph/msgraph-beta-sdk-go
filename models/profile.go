package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// Profile 
type Profile struct {
    Entity
}
// NewProfile instantiates a new profile and sets the default values.
func NewProfile()(*Profile) {
    m := &Profile{
        Entity: *NewEntity(),
    }
    return m
}
// CreateProfileFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateProfileFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewProfile(), nil
}
// GetAccount gets the account property value. The account property
func (m *Profile) GetAccount()([]UserAccountInformationable) {
    val, err := m.GetBackingStore().Get("account")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]UserAccountInformationable)
    }
    return nil
}
// GetAddresses gets the addresses property value. Represents details of addresses associated with the user.
func (m *Profile) GetAddresses()([]ItemAddressable) {
    val, err := m.GetBackingStore().Get("addresses")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]ItemAddressable)
    }
    return nil
}
// GetAnniversaries gets the anniversaries property value. Represents the details of meaningful dates associated with a person.
func (m *Profile) GetAnniversaries()([]PersonAnnualEventable) {
    val, err := m.GetBackingStore().Get("anniversaries")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]PersonAnnualEventable)
    }
    return nil
}
// GetAwards gets the awards property value. Represents the details of awards or honors associated with a person.
func (m *Profile) GetAwards()([]PersonAwardable) {
    val, err := m.GetBackingStore().Get("awards")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]PersonAwardable)
    }
    return nil
}
// GetCertifications gets the certifications property value. Represents the details of certifications associated with a person.
func (m *Profile) GetCertifications()([]PersonCertificationable) {
    val, err := m.GetBackingStore().Get("certifications")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]PersonCertificationable)
    }
    return nil
}
// GetEducationalActivities gets the educationalActivities property value. Represents data that a user has supplied related to undergraduate, graduate, postgraduate or other educational activities.
func (m *Profile) GetEducationalActivities()([]EducationalActivityable) {
    val, err := m.GetBackingStore().Get("educationalActivities")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]EducationalActivityable)
    }
    return nil
}
// GetEmails gets the emails property value. Represents detailed information about email addresses associated with the user.
func (m *Profile) GetEmails()([]ItemEmailable) {
    val, err := m.GetBackingStore().Get("emails")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]ItemEmailable)
    }
    return nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *Profile) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["account"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateUserAccountInformationFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]UserAccountInformationable, len(val))
            for i, v := range val {
                res[i] = v.(UserAccountInformationable)
            }
            m.SetAccount(res)
        }
        return nil
    }
    res["addresses"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateItemAddressFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]ItemAddressable, len(val))
            for i, v := range val {
                res[i] = v.(ItemAddressable)
            }
            m.SetAddresses(res)
        }
        return nil
    }
    res["anniversaries"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreatePersonAnnualEventFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]PersonAnnualEventable, len(val))
            for i, v := range val {
                res[i] = v.(PersonAnnualEventable)
            }
            m.SetAnniversaries(res)
        }
        return nil
    }
    res["awards"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreatePersonAwardFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]PersonAwardable, len(val))
            for i, v := range val {
                res[i] = v.(PersonAwardable)
            }
            m.SetAwards(res)
        }
        return nil
    }
    res["certifications"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreatePersonCertificationFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]PersonCertificationable, len(val))
            for i, v := range val {
                res[i] = v.(PersonCertificationable)
            }
            m.SetCertifications(res)
        }
        return nil
    }
    res["educationalActivities"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateEducationalActivityFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]EducationalActivityable, len(val))
            for i, v := range val {
                res[i] = v.(EducationalActivityable)
            }
            m.SetEducationalActivities(res)
        }
        return nil
    }
    res["emails"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateItemEmailFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]ItemEmailable, len(val))
            for i, v := range val {
                res[i] = v.(ItemEmailable)
            }
            m.SetEmails(res)
        }
        return nil
    }
    res["interests"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreatePersonInterestFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]PersonInterestable, len(val))
            for i, v := range val {
                res[i] = v.(PersonInterestable)
            }
            m.SetInterests(res)
        }
        return nil
    }
    res["languages"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateLanguageProficiencyFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]LanguageProficiencyable, len(val))
            for i, v := range val {
                res[i] = v.(LanguageProficiencyable)
            }
            m.SetLanguages(res)
        }
        return nil
    }
    res["names"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreatePersonNameFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]PersonNameable, len(val))
            for i, v := range val {
                res[i] = v.(PersonNameable)
            }
            m.SetNames(res)
        }
        return nil
    }
    res["notes"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreatePersonAnnotationFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]PersonAnnotationable, len(val))
            for i, v := range val {
                res[i] = v.(PersonAnnotationable)
            }
            m.SetNotes(res)
        }
        return nil
    }
    res["patents"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateItemPatentFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]ItemPatentable, len(val))
            for i, v := range val {
                res[i] = v.(ItemPatentable)
            }
            m.SetPatents(res)
        }
        return nil
    }
    res["phones"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateItemPhoneFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]ItemPhoneable, len(val))
            for i, v := range val {
                res[i] = v.(ItemPhoneable)
            }
            m.SetPhones(res)
        }
        return nil
    }
    res["positions"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateWorkPositionFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]WorkPositionable, len(val))
            for i, v := range val {
                res[i] = v.(WorkPositionable)
            }
            m.SetPositions(res)
        }
        return nil
    }
    res["projects"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateProjectParticipationFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]ProjectParticipationable, len(val))
            for i, v := range val {
                res[i] = v.(ProjectParticipationable)
            }
            m.SetProjects(res)
        }
        return nil
    }
    res["publications"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateItemPublicationFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]ItemPublicationable, len(val))
            for i, v := range val {
                res[i] = v.(ItemPublicationable)
            }
            m.SetPublications(res)
        }
        return nil
    }
    res["skills"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateSkillProficiencyFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]SkillProficiencyable, len(val))
            for i, v := range val {
                res[i] = v.(SkillProficiencyable)
            }
            m.SetSkills(res)
        }
        return nil
    }
    res["webAccounts"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateWebAccountFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]WebAccountable, len(val))
            for i, v := range val {
                res[i] = v.(WebAccountable)
            }
            m.SetWebAccounts(res)
        }
        return nil
    }
    res["websites"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreatePersonWebsiteFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]PersonWebsiteable, len(val))
            for i, v := range val {
                res[i] = v.(PersonWebsiteable)
            }
            m.SetWebsites(res)
        }
        return nil
    }
    return res
}
// GetInterests gets the interests property value. Provides detailed information about interests the user has associated with themselves in various services.
func (m *Profile) GetInterests()([]PersonInterestable) {
    val, err := m.GetBackingStore().Get("interests")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]PersonInterestable)
    }
    return nil
}
// GetLanguages gets the languages property value. Represents detailed information about languages that a user has added to their profile.
func (m *Profile) GetLanguages()([]LanguageProficiencyable) {
    val, err := m.GetBackingStore().Get("languages")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]LanguageProficiencyable)
    }
    return nil
}
// GetNames gets the names property value. Represents the names a user has added to their profile.
func (m *Profile) GetNames()([]PersonNameable) {
    val, err := m.GetBackingStore().Get("names")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]PersonNameable)
    }
    return nil
}
// GetNotes gets the notes property value. Represents notes that a user has added to their profile.
func (m *Profile) GetNotes()([]PersonAnnotationable) {
    val, err := m.GetBackingStore().Get("notes")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]PersonAnnotationable)
    }
    return nil
}
// GetPatents gets the patents property value. Represents patents that a user has added to their profile.
func (m *Profile) GetPatents()([]ItemPatentable) {
    val, err := m.GetBackingStore().Get("patents")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]ItemPatentable)
    }
    return nil
}
// GetPhones gets the phones property value. Represents detailed information about phone numbers associated with a user in various services.
func (m *Profile) GetPhones()([]ItemPhoneable) {
    val, err := m.GetBackingStore().Get("phones")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]ItemPhoneable)
    }
    return nil
}
// GetPositions gets the positions property value. Represents detailed information about work positions associated with a user's profile.
func (m *Profile) GetPositions()([]WorkPositionable) {
    val, err := m.GetBackingStore().Get("positions")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]WorkPositionable)
    }
    return nil
}
// GetProjects gets the projects property value. Represents detailed information about projects associated with a user.
func (m *Profile) GetProjects()([]ProjectParticipationable) {
    val, err := m.GetBackingStore().Get("projects")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]ProjectParticipationable)
    }
    return nil
}
// GetPublications gets the publications property value. Represents details of any publications a user has added to their profile.
func (m *Profile) GetPublications()([]ItemPublicationable) {
    val, err := m.GetBackingStore().Get("publications")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]ItemPublicationable)
    }
    return nil
}
// GetSkills gets the skills property value. Represents detailed information about skills associated with a user in various services.
func (m *Profile) GetSkills()([]SkillProficiencyable) {
    val, err := m.GetBackingStore().Get("skills")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]SkillProficiencyable)
    }
    return nil
}
// GetWebAccounts gets the webAccounts property value. Represents web accounts the user has indicated they use or has added to their user profile.
func (m *Profile) GetWebAccounts()([]WebAccountable) {
    val, err := m.GetBackingStore().Get("webAccounts")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]WebAccountable)
    }
    return nil
}
// GetWebsites gets the websites property value. Represents detailed information about websites associated with a user in various services.
func (m *Profile) GetWebsites()([]PersonWebsiteable) {
    val, err := m.GetBackingStore().Get("websites")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]PersonWebsiteable)
    }
    return nil
}
// Serialize serializes information the current object
func (m *Profile) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetAccount() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetAccount()))
        for i, v := range m.GetAccount() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("account", cast)
        if err != nil {
            return err
        }
    }
    if m.GetAddresses() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetAddresses()))
        for i, v := range m.GetAddresses() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("addresses", cast)
        if err != nil {
            return err
        }
    }
    if m.GetAnniversaries() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetAnniversaries()))
        for i, v := range m.GetAnniversaries() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("anniversaries", cast)
        if err != nil {
            return err
        }
    }
    if m.GetAwards() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetAwards()))
        for i, v := range m.GetAwards() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("awards", cast)
        if err != nil {
            return err
        }
    }
    if m.GetCertifications() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetCertifications()))
        for i, v := range m.GetCertifications() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("certifications", cast)
        if err != nil {
            return err
        }
    }
    if m.GetEducationalActivities() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetEducationalActivities()))
        for i, v := range m.GetEducationalActivities() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("educationalActivities", cast)
        if err != nil {
            return err
        }
    }
    if m.GetEmails() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetEmails()))
        for i, v := range m.GetEmails() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("emails", cast)
        if err != nil {
            return err
        }
    }
    if m.GetInterests() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetInterests()))
        for i, v := range m.GetInterests() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("interests", cast)
        if err != nil {
            return err
        }
    }
    if m.GetLanguages() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetLanguages()))
        for i, v := range m.GetLanguages() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("languages", cast)
        if err != nil {
            return err
        }
    }
    if m.GetNames() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetNames()))
        for i, v := range m.GetNames() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("names", cast)
        if err != nil {
            return err
        }
    }
    if m.GetNotes() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetNotes()))
        for i, v := range m.GetNotes() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("notes", cast)
        if err != nil {
            return err
        }
    }
    if m.GetPatents() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetPatents()))
        for i, v := range m.GetPatents() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("patents", cast)
        if err != nil {
            return err
        }
    }
    if m.GetPhones() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetPhones()))
        for i, v := range m.GetPhones() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("phones", cast)
        if err != nil {
            return err
        }
    }
    if m.GetPositions() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetPositions()))
        for i, v := range m.GetPositions() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("positions", cast)
        if err != nil {
            return err
        }
    }
    if m.GetProjects() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetProjects()))
        for i, v := range m.GetProjects() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("projects", cast)
        if err != nil {
            return err
        }
    }
    if m.GetPublications() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetPublications()))
        for i, v := range m.GetPublications() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("publications", cast)
        if err != nil {
            return err
        }
    }
    if m.GetSkills() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetSkills()))
        for i, v := range m.GetSkills() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("skills", cast)
        if err != nil {
            return err
        }
    }
    if m.GetWebAccounts() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetWebAccounts()))
        for i, v := range m.GetWebAccounts() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("webAccounts", cast)
        if err != nil {
            return err
        }
    }
    if m.GetWebsites() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetWebsites()))
        for i, v := range m.GetWebsites() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("websites", cast)
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAccount sets the account property value. The account property
func (m *Profile) SetAccount(value []UserAccountInformationable)() {
    err := m.GetBackingStore().Set("account", value)
    if err != nil {
        panic(err)
    }
}
// SetAddresses sets the addresses property value. Represents details of addresses associated with the user.
func (m *Profile) SetAddresses(value []ItemAddressable)() {
    err := m.GetBackingStore().Set("addresses", value)
    if err != nil {
        panic(err)
    }
}
// SetAnniversaries sets the anniversaries property value. Represents the details of meaningful dates associated with a person.
func (m *Profile) SetAnniversaries(value []PersonAnnualEventable)() {
    err := m.GetBackingStore().Set("anniversaries", value)
    if err != nil {
        panic(err)
    }
}
// SetAwards sets the awards property value. Represents the details of awards or honors associated with a person.
func (m *Profile) SetAwards(value []PersonAwardable)() {
    err := m.GetBackingStore().Set("awards", value)
    if err != nil {
        panic(err)
    }
}
// SetCertifications sets the certifications property value. Represents the details of certifications associated with a person.
func (m *Profile) SetCertifications(value []PersonCertificationable)() {
    err := m.GetBackingStore().Set("certifications", value)
    if err != nil {
        panic(err)
    }
}
// SetEducationalActivities sets the educationalActivities property value. Represents data that a user has supplied related to undergraduate, graduate, postgraduate or other educational activities.
func (m *Profile) SetEducationalActivities(value []EducationalActivityable)() {
    err := m.GetBackingStore().Set("educationalActivities", value)
    if err != nil {
        panic(err)
    }
}
// SetEmails sets the emails property value. Represents detailed information about email addresses associated with the user.
func (m *Profile) SetEmails(value []ItemEmailable)() {
    err := m.GetBackingStore().Set("emails", value)
    if err != nil {
        panic(err)
    }
}
// SetInterests sets the interests property value. Provides detailed information about interests the user has associated with themselves in various services.
func (m *Profile) SetInterests(value []PersonInterestable)() {
    err := m.GetBackingStore().Set("interests", value)
    if err != nil {
        panic(err)
    }
}
// SetLanguages sets the languages property value. Represents detailed information about languages that a user has added to their profile.
func (m *Profile) SetLanguages(value []LanguageProficiencyable)() {
    err := m.GetBackingStore().Set("languages", value)
    if err != nil {
        panic(err)
    }
}
// SetNames sets the names property value. Represents the names a user has added to their profile.
func (m *Profile) SetNames(value []PersonNameable)() {
    err := m.GetBackingStore().Set("names", value)
    if err != nil {
        panic(err)
    }
}
// SetNotes sets the notes property value. Represents notes that a user has added to their profile.
func (m *Profile) SetNotes(value []PersonAnnotationable)() {
    err := m.GetBackingStore().Set("notes", value)
    if err != nil {
        panic(err)
    }
}
// SetPatents sets the patents property value. Represents patents that a user has added to their profile.
func (m *Profile) SetPatents(value []ItemPatentable)() {
    err := m.GetBackingStore().Set("patents", value)
    if err != nil {
        panic(err)
    }
}
// SetPhones sets the phones property value. Represents detailed information about phone numbers associated with a user in various services.
func (m *Profile) SetPhones(value []ItemPhoneable)() {
    err := m.GetBackingStore().Set("phones", value)
    if err != nil {
        panic(err)
    }
}
// SetPositions sets the positions property value. Represents detailed information about work positions associated with a user's profile.
func (m *Profile) SetPositions(value []WorkPositionable)() {
    err := m.GetBackingStore().Set("positions", value)
    if err != nil {
        panic(err)
    }
}
// SetProjects sets the projects property value. Represents detailed information about projects associated with a user.
func (m *Profile) SetProjects(value []ProjectParticipationable)() {
    err := m.GetBackingStore().Set("projects", value)
    if err != nil {
        panic(err)
    }
}
// SetPublications sets the publications property value. Represents details of any publications a user has added to their profile.
func (m *Profile) SetPublications(value []ItemPublicationable)() {
    err := m.GetBackingStore().Set("publications", value)
    if err != nil {
        panic(err)
    }
}
// SetSkills sets the skills property value. Represents detailed information about skills associated with a user in various services.
func (m *Profile) SetSkills(value []SkillProficiencyable)() {
    err := m.GetBackingStore().Set("skills", value)
    if err != nil {
        panic(err)
    }
}
// SetWebAccounts sets the webAccounts property value. Represents web accounts the user has indicated they use or has added to their user profile.
func (m *Profile) SetWebAccounts(value []WebAccountable)() {
    err := m.GetBackingStore().Set("webAccounts", value)
    if err != nil {
        panic(err)
    }
}
// SetWebsites sets the websites property value. Represents detailed information about websites associated with a user in various services.
func (m *Profile) SetWebsites(value []PersonWebsiteable)() {
    err := m.GetBackingStore().Set("websites", value)
    if err != nil {
        panic(err)
    }
}
// Profileable 
type Profileable interface {
    Entityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAccount()([]UserAccountInformationable)
    GetAddresses()([]ItemAddressable)
    GetAnniversaries()([]PersonAnnualEventable)
    GetAwards()([]PersonAwardable)
    GetCertifications()([]PersonCertificationable)
    GetEducationalActivities()([]EducationalActivityable)
    GetEmails()([]ItemEmailable)
    GetInterests()([]PersonInterestable)
    GetLanguages()([]LanguageProficiencyable)
    GetNames()([]PersonNameable)
    GetNotes()([]PersonAnnotationable)
    GetPatents()([]ItemPatentable)
    GetPhones()([]ItemPhoneable)
    GetPositions()([]WorkPositionable)
    GetProjects()([]ProjectParticipationable)
    GetPublications()([]ItemPublicationable)
    GetSkills()([]SkillProficiencyable)
    GetWebAccounts()([]WebAccountable)
    GetWebsites()([]PersonWebsiteable)
    SetAccount(value []UserAccountInformationable)()
    SetAddresses(value []ItemAddressable)()
    SetAnniversaries(value []PersonAnnualEventable)()
    SetAwards(value []PersonAwardable)()
    SetCertifications(value []PersonCertificationable)()
    SetEducationalActivities(value []EducationalActivityable)()
    SetEmails(value []ItemEmailable)()
    SetInterests(value []PersonInterestable)()
    SetLanguages(value []LanguageProficiencyable)()
    SetNames(value []PersonNameable)()
    SetNotes(value []PersonAnnotationable)()
    SetPatents(value []ItemPatentable)()
    SetPhones(value []ItemPhoneable)()
    SetPositions(value []WorkPositionable)()
    SetProjects(value []ProjectParticipationable)()
    SetPublications(value []ItemPublicationable)()
    SetSkills(value []SkillProficiencyable)()
    SetWebAccounts(value []WebAccountable)()
    SetWebsites(value []PersonWebsiteable)()
}
