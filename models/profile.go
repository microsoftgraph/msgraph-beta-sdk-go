package models

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// Profile 
type Profile struct {
    Entity
    // The account property
    account []UserAccountInformationable
    // Represents details of addresses associated with the user.
    addresses []ItemAddressable
    // Represents the details of meaningful dates associated with a person.
    anniversaries []PersonAnnualEventable
    // Represents the details of awards or honors associated with a person.
    awards []PersonAwardable
    // Represents the details of certifications associated with a person.
    certifications []PersonCertificationable
    // Represents data that a user has supplied related to undergraduate, graduate, postgraduate or other educational activities.
    educationalActivities []EducationalActivityable
    // Represents detailed information about email addresses associated with the user.
    emails []ItemEmailable
    // Provides detailed information about interests the user has associated with themselves in various services.
    interests []PersonInterestable
    // Represents detailed information about languages that a user has added to their profile.
    languages []LanguageProficiencyable
    // Represents the names a user has added to their profile.
    names []PersonNameable
    // Represents notes that a user has added to their profile.
    notes []PersonAnnotationable
    // Represents patents that a user has added to their profile.
    patents []ItemPatentable
    // Represents detailed information about phone numbers associated with a user in various services.
    phones []ItemPhoneable
    // Represents detailed information about work positions associated with a user's profile.
    positions []WorkPositionable
    // Represents detailed information about projects associated with a user.
    projects []ProjectParticipationable
    // Represents details of any publications a user has added to their profile.
    publications []ItemPublicationable
    // Represents detailed information about skills associated with a user in various services.
    skills []SkillProficiencyable
    // Represents web accounts the user has indicated they use or has added to their user profile.
    webAccounts []WebAccountable
    // Represents detailed information about websites associated with a user in various services.
    websites []PersonWebsiteable
}
// NewProfile instantiates a new Profile and sets the default values.
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
    return m.account
}
// GetAddresses gets the addresses property value. Represents details of addresses associated with the user.
func (m *Profile) GetAddresses()([]ItemAddressable) {
    return m.addresses
}
// GetAnniversaries gets the anniversaries property value. Represents the details of meaningful dates associated with a person.
func (m *Profile) GetAnniversaries()([]PersonAnnualEventable) {
    return m.anniversaries
}
// GetAwards gets the awards property value. Represents the details of awards or honors associated with a person.
func (m *Profile) GetAwards()([]PersonAwardable) {
    return m.awards
}
// GetCertifications gets the certifications property value. Represents the details of certifications associated with a person.
func (m *Profile) GetCertifications()([]PersonCertificationable) {
    return m.certifications
}
// GetEducationalActivities gets the educationalActivities property value. Represents data that a user has supplied related to undergraduate, graduate, postgraduate or other educational activities.
func (m *Profile) GetEducationalActivities()([]EducationalActivityable) {
    return m.educationalActivities
}
// GetEmails gets the emails property value. Represents detailed information about email addresses associated with the user.
func (m *Profile) GetEmails()([]ItemEmailable) {
    return m.emails
}
// GetFieldDeserializers the deserialization information for the current model
func (m *Profile) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["account"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfObjectValues(CreateUserAccountInformationFromDiscriminatorValue , m.SetAccount)
    res["addresses"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfObjectValues(CreateItemAddressFromDiscriminatorValue , m.SetAddresses)
    res["anniversaries"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfObjectValues(CreatePersonAnnualEventFromDiscriminatorValue , m.SetAnniversaries)
    res["awards"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfObjectValues(CreatePersonAwardFromDiscriminatorValue , m.SetAwards)
    res["certifications"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfObjectValues(CreatePersonCertificationFromDiscriminatorValue , m.SetCertifications)
    res["educationalActivities"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfObjectValues(CreateEducationalActivityFromDiscriminatorValue , m.SetEducationalActivities)
    res["emails"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfObjectValues(CreateItemEmailFromDiscriminatorValue , m.SetEmails)
    res["interests"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfObjectValues(CreatePersonInterestFromDiscriminatorValue , m.SetInterests)
    res["languages"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfObjectValues(CreateLanguageProficiencyFromDiscriminatorValue , m.SetLanguages)
    res["names"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfObjectValues(CreatePersonNameFromDiscriminatorValue , m.SetNames)
    res["notes"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfObjectValues(CreatePersonAnnotationFromDiscriminatorValue , m.SetNotes)
    res["patents"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfObjectValues(CreateItemPatentFromDiscriminatorValue , m.SetPatents)
    res["phones"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfObjectValues(CreateItemPhoneFromDiscriminatorValue , m.SetPhones)
    res["positions"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfObjectValues(CreateWorkPositionFromDiscriminatorValue , m.SetPositions)
    res["projects"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfObjectValues(CreateProjectParticipationFromDiscriminatorValue , m.SetProjects)
    res["publications"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfObjectValues(CreateItemPublicationFromDiscriminatorValue , m.SetPublications)
    res["skills"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfObjectValues(CreateSkillProficiencyFromDiscriminatorValue , m.SetSkills)
    res["webAccounts"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfObjectValues(CreateWebAccountFromDiscriminatorValue , m.SetWebAccounts)
    res["websites"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfObjectValues(CreatePersonWebsiteFromDiscriminatorValue , m.SetWebsites)
    return res
}
// GetInterests gets the interests property value. Provides detailed information about interests the user has associated with themselves in various services.
func (m *Profile) GetInterests()([]PersonInterestable) {
    return m.interests
}
// GetLanguages gets the languages property value. Represents detailed information about languages that a user has added to their profile.
func (m *Profile) GetLanguages()([]LanguageProficiencyable) {
    return m.languages
}
// GetNames gets the names property value. Represents the names a user has added to their profile.
func (m *Profile) GetNames()([]PersonNameable) {
    return m.names
}
// GetNotes gets the notes property value. Represents notes that a user has added to their profile.
func (m *Profile) GetNotes()([]PersonAnnotationable) {
    return m.notes
}
// GetPatents gets the patents property value. Represents patents that a user has added to their profile.
func (m *Profile) GetPatents()([]ItemPatentable) {
    return m.patents
}
// GetPhones gets the phones property value. Represents detailed information about phone numbers associated with a user in various services.
func (m *Profile) GetPhones()([]ItemPhoneable) {
    return m.phones
}
// GetPositions gets the positions property value. Represents detailed information about work positions associated with a user's profile.
func (m *Profile) GetPositions()([]WorkPositionable) {
    return m.positions
}
// GetProjects gets the projects property value. Represents detailed information about projects associated with a user.
func (m *Profile) GetProjects()([]ProjectParticipationable) {
    return m.projects
}
// GetPublications gets the publications property value. Represents details of any publications a user has added to their profile.
func (m *Profile) GetPublications()([]ItemPublicationable) {
    return m.publications
}
// GetSkills gets the skills property value. Represents detailed information about skills associated with a user in various services.
func (m *Profile) GetSkills()([]SkillProficiencyable) {
    return m.skills
}
// GetWebAccounts gets the webAccounts property value. Represents web accounts the user has indicated they use or has added to their user profile.
func (m *Profile) GetWebAccounts()([]WebAccountable) {
    return m.webAccounts
}
// GetWebsites gets the websites property value. Represents detailed information about websites associated with a user in various services.
func (m *Profile) GetWebsites()([]PersonWebsiteable) {
    return m.websites
}
// Serialize serializes information the current object
func (m *Profile) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetAccount() != nil {
        cast := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.CollectionCast[i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable](m.GetAccount())
        err = writer.WriteCollectionOfObjectValues("account", cast)
        if err != nil {
            return err
        }
    }
    if m.GetAddresses() != nil {
        cast := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.CollectionCast[i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable](m.GetAddresses())
        err = writer.WriteCollectionOfObjectValues("addresses", cast)
        if err != nil {
            return err
        }
    }
    if m.GetAnniversaries() != nil {
        cast := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.CollectionCast[i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable](m.GetAnniversaries())
        err = writer.WriteCollectionOfObjectValues("anniversaries", cast)
        if err != nil {
            return err
        }
    }
    if m.GetAwards() != nil {
        cast := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.CollectionCast[i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable](m.GetAwards())
        err = writer.WriteCollectionOfObjectValues("awards", cast)
        if err != nil {
            return err
        }
    }
    if m.GetCertifications() != nil {
        cast := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.CollectionCast[i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable](m.GetCertifications())
        err = writer.WriteCollectionOfObjectValues("certifications", cast)
        if err != nil {
            return err
        }
    }
    if m.GetEducationalActivities() != nil {
        cast := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.CollectionCast[i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable](m.GetEducationalActivities())
        err = writer.WriteCollectionOfObjectValues("educationalActivities", cast)
        if err != nil {
            return err
        }
    }
    if m.GetEmails() != nil {
        cast := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.CollectionCast[i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable](m.GetEmails())
        err = writer.WriteCollectionOfObjectValues("emails", cast)
        if err != nil {
            return err
        }
    }
    if m.GetInterests() != nil {
        cast := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.CollectionCast[i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable](m.GetInterests())
        err = writer.WriteCollectionOfObjectValues("interests", cast)
        if err != nil {
            return err
        }
    }
    if m.GetLanguages() != nil {
        cast := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.CollectionCast[i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable](m.GetLanguages())
        err = writer.WriteCollectionOfObjectValues("languages", cast)
        if err != nil {
            return err
        }
    }
    if m.GetNames() != nil {
        cast := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.CollectionCast[i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable](m.GetNames())
        err = writer.WriteCollectionOfObjectValues("names", cast)
        if err != nil {
            return err
        }
    }
    if m.GetNotes() != nil {
        cast := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.CollectionCast[i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable](m.GetNotes())
        err = writer.WriteCollectionOfObjectValues("notes", cast)
        if err != nil {
            return err
        }
    }
    if m.GetPatents() != nil {
        cast := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.CollectionCast[i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable](m.GetPatents())
        err = writer.WriteCollectionOfObjectValues("patents", cast)
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
    if m.GetPositions() != nil {
        cast := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.CollectionCast[i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable](m.GetPositions())
        err = writer.WriteCollectionOfObjectValues("positions", cast)
        if err != nil {
            return err
        }
    }
    if m.GetProjects() != nil {
        cast := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.CollectionCast[i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable](m.GetProjects())
        err = writer.WriteCollectionOfObjectValues("projects", cast)
        if err != nil {
            return err
        }
    }
    if m.GetPublications() != nil {
        cast := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.CollectionCast[i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable](m.GetPublications())
        err = writer.WriteCollectionOfObjectValues("publications", cast)
        if err != nil {
            return err
        }
    }
    if m.GetSkills() != nil {
        cast := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.CollectionCast[i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable](m.GetSkills())
        err = writer.WriteCollectionOfObjectValues("skills", cast)
        if err != nil {
            return err
        }
    }
    if m.GetWebAccounts() != nil {
        cast := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.CollectionCast[i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable](m.GetWebAccounts())
        err = writer.WriteCollectionOfObjectValues("webAccounts", cast)
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
    return nil
}
// SetAccount sets the account property value. The account property
func (m *Profile) SetAccount(value []UserAccountInformationable)() {
    m.account = value
}
// SetAddresses sets the addresses property value. Represents details of addresses associated with the user.
func (m *Profile) SetAddresses(value []ItemAddressable)() {
    m.addresses = value
}
// SetAnniversaries sets the anniversaries property value. Represents the details of meaningful dates associated with a person.
func (m *Profile) SetAnniversaries(value []PersonAnnualEventable)() {
    m.anniversaries = value
}
// SetAwards sets the awards property value. Represents the details of awards or honors associated with a person.
func (m *Profile) SetAwards(value []PersonAwardable)() {
    m.awards = value
}
// SetCertifications sets the certifications property value. Represents the details of certifications associated with a person.
func (m *Profile) SetCertifications(value []PersonCertificationable)() {
    m.certifications = value
}
// SetEducationalActivities sets the educationalActivities property value. Represents data that a user has supplied related to undergraduate, graduate, postgraduate or other educational activities.
func (m *Profile) SetEducationalActivities(value []EducationalActivityable)() {
    m.educationalActivities = value
}
// SetEmails sets the emails property value. Represents detailed information about email addresses associated with the user.
func (m *Profile) SetEmails(value []ItemEmailable)() {
    m.emails = value
}
// SetInterests sets the interests property value. Provides detailed information about interests the user has associated with themselves in various services.
func (m *Profile) SetInterests(value []PersonInterestable)() {
    m.interests = value
}
// SetLanguages sets the languages property value. Represents detailed information about languages that a user has added to their profile.
func (m *Profile) SetLanguages(value []LanguageProficiencyable)() {
    m.languages = value
}
// SetNames sets the names property value. Represents the names a user has added to their profile.
func (m *Profile) SetNames(value []PersonNameable)() {
    m.names = value
}
// SetNotes sets the notes property value. Represents notes that a user has added to their profile.
func (m *Profile) SetNotes(value []PersonAnnotationable)() {
    m.notes = value
}
// SetPatents sets the patents property value. Represents patents that a user has added to their profile.
func (m *Profile) SetPatents(value []ItemPatentable)() {
    m.patents = value
}
// SetPhones sets the phones property value. Represents detailed information about phone numbers associated with a user in various services.
func (m *Profile) SetPhones(value []ItemPhoneable)() {
    m.phones = value
}
// SetPositions sets the positions property value. Represents detailed information about work positions associated with a user's profile.
func (m *Profile) SetPositions(value []WorkPositionable)() {
    m.positions = value
}
// SetProjects sets the projects property value. Represents detailed information about projects associated with a user.
func (m *Profile) SetProjects(value []ProjectParticipationable)() {
    m.projects = value
}
// SetPublications sets the publications property value. Represents details of any publications a user has added to their profile.
func (m *Profile) SetPublications(value []ItemPublicationable)() {
    m.publications = value
}
// SetSkills sets the skills property value. Represents detailed information about skills associated with a user in various services.
func (m *Profile) SetSkills(value []SkillProficiencyable)() {
    m.skills = value
}
// SetWebAccounts sets the webAccounts property value. Represents web accounts the user has indicated they use or has added to their user profile.
func (m *Profile) SetWebAccounts(value []WebAccountable)() {
    m.webAccounts = value
}
// SetWebsites sets the websites property value. Represents detailed information about websites associated with a user in various services.
func (m *Profile) SetWebsites(value []PersonWebsiteable)() {
    m.websites = value
}
