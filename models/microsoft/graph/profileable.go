package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// Profileable 
type Profileable interface {
    Entityable
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable
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
