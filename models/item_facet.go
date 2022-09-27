package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ItemFacet provides operations to manage the collection of activityStatistics entities.
type ItemFacet struct {
    Entity
    // The audiences that are able to see the values contained within the associated entity. Possible values are: me, family, contacts, groupMembers, organization, federatedOrganizations, everyone, unknownFutureValue.
    allowedAudiences *AllowedAudiences
    // The createdBy property
    createdBy IdentitySetable
    // Provides the dateTimeOffset for when the entity was created.
    createdDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // Contains inference detail if the entity is inferred by the creating or modifying application.
    inference InferenceDataable
    // The isSearchable property
    isSearchable *bool
    // The lastModifiedBy property
    lastModifiedBy IdentitySetable
    // Provides the dateTimeOffset for when the entity was created.
    lastModifiedDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // Where the values within an entity originated if synced from another service.
    source PersonDataSourcesable
}
// NewItemFacet instantiates a new itemFacet and sets the default values.
func NewItemFacet()(*ItemFacet) {
    m := &ItemFacet{
        Entity: *NewEntity(),
    }
    odataTypeValue := "#microsoft.graph.itemFacet";
    m.SetOdataType(&odataTypeValue);
    return m
}
// CreateItemFacetFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateItemFacetFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    if parseNode != nil {
        mappingValueNode, err := parseNode.GetChildNode("@odata.type")
        if err != nil {
            return nil, err
        }
        if mappingValueNode != nil {
            mappingValue, err := mappingValueNode.GetStringValue()
            if err != nil {
                return nil, err
            }
            if mappingValue != nil {
                switch *mappingValue {
                    case "#microsoft.graph.educationalActivity":
                        return NewEducationalActivity(), nil
                    case "#microsoft.graph.itemAddress":
                        return NewItemAddress(), nil
                    case "#microsoft.graph.itemEmail":
                        return NewItemEmail(), nil
                    case "#microsoft.graph.itemPatent":
                        return NewItemPatent(), nil
                    case "#microsoft.graph.itemPhone":
                        return NewItemPhone(), nil
                    case "#microsoft.graph.itemPublication":
                        return NewItemPublication(), nil
                    case "#microsoft.graph.languageProficiency":
                        return NewLanguageProficiency(), nil
                    case "#microsoft.graph.personAnnotation":
                        return NewPersonAnnotation(), nil
                    case "#microsoft.graph.personAnnualEvent":
                        return NewPersonAnnualEvent(), nil
                    case "#microsoft.graph.personAward":
                        return NewPersonAward(), nil
                    case "#microsoft.graph.personCertification":
                        return NewPersonCertification(), nil
                    case "#microsoft.graph.personInterest":
                        return NewPersonInterest(), nil
                    case "#microsoft.graph.personName":
                        return NewPersonName(), nil
                    case "#microsoft.graph.personResponsibility":
                        return NewPersonResponsibility(), nil
                    case "#microsoft.graph.personWebsite":
                        return NewPersonWebsite(), nil
                    case "#microsoft.graph.projectParticipation":
                        return NewProjectParticipation(), nil
                    case "#microsoft.graph.skillProficiency":
                        return NewSkillProficiency(), nil
                    case "#microsoft.graph.userAccountInformation":
                        return NewUserAccountInformation(), nil
                    case "#microsoft.graph.webAccount":
                        return NewWebAccount(), nil
                    case "#microsoft.graph.workPosition":
                        return NewWorkPosition(), nil
                }
            }
        }
    }
    return NewItemFacet(), nil
}
// GetAllowedAudiences gets the allowedAudiences property value. The audiences that are able to see the values contained within the associated entity. Possible values are: me, family, contacts, groupMembers, organization, federatedOrganizations, everyone, unknownFutureValue.
func (m *ItemFacet) GetAllowedAudiences()(*AllowedAudiences) {
    return m.allowedAudiences
}
// GetCreatedBy gets the createdBy property value. The createdBy property
func (m *ItemFacet) GetCreatedBy()(IdentitySetable) {
    return m.createdBy
}
// GetCreatedDateTime gets the createdDateTime property value. Provides the dateTimeOffset for when the entity was created.
func (m *ItemFacet) GetCreatedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.createdDateTime
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ItemFacet) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["allowedAudiences"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetEnumValue(ParseAllowedAudiences , m.SetAllowedAudiences)
    res["createdBy"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetObjectValue(CreateIdentitySetFromDiscriminatorValue , m.SetCreatedBy)
    res["createdDateTime"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetTimeValue(m.SetCreatedDateTime)
    res["inference"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetObjectValue(CreateInferenceDataFromDiscriminatorValue , m.SetInference)
    res["isSearchable"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetBoolValue(m.SetIsSearchable)
    res["lastModifiedBy"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetObjectValue(CreateIdentitySetFromDiscriminatorValue , m.SetLastModifiedBy)
    res["lastModifiedDateTime"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetTimeValue(m.SetLastModifiedDateTime)
    res["source"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetObjectValue(CreatePersonDataSourcesFromDiscriminatorValue , m.SetSource)
    return res
}
// GetInference gets the inference property value. Contains inference detail if the entity is inferred by the creating or modifying application.
func (m *ItemFacet) GetInference()(InferenceDataable) {
    return m.inference
}
// GetIsSearchable gets the isSearchable property value. The isSearchable property
func (m *ItemFacet) GetIsSearchable()(*bool) {
    return m.isSearchable
}
// GetLastModifiedBy gets the lastModifiedBy property value. The lastModifiedBy property
func (m *ItemFacet) GetLastModifiedBy()(IdentitySetable) {
    return m.lastModifiedBy
}
// GetLastModifiedDateTime gets the lastModifiedDateTime property value. Provides the dateTimeOffset for when the entity was created.
func (m *ItemFacet) GetLastModifiedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.lastModifiedDateTime
}
// GetSource gets the source property value. Where the values within an entity originated if synced from another service.
func (m *ItemFacet) GetSource()(PersonDataSourcesable) {
    return m.source
}
// Serialize serializes information the current object
func (m *ItemFacet) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetAllowedAudiences() != nil {
        cast := (*m.GetAllowedAudiences()).String()
        err = writer.WriteStringValue("allowedAudiences", &cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("createdBy", m.GetCreatedBy())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteTimeValue("createdDateTime", m.GetCreatedDateTime())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("inference", m.GetInference())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("isSearchable", m.GetIsSearchable())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("lastModifiedBy", m.GetLastModifiedBy())
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
        err = writer.WriteObjectValue("source", m.GetSource())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAllowedAudiences sets the allowedAudiences property value. The audiences that are able to see the values contained within the associated entity. Possible values are: me, family, contacts, groupMembers, organization, federatedOrganizations, everyone, unknownFutureValue.
func (m *ItemFacet) SetAllowedAudiences(value *AllowedAudiences)() {
    m.allowedAudiences = value
}
// SetCreatedBy sets the createdBy property value. The createdBy property
func (m *ItemFacet) SetCreatedBy(value IdentitySetable)() {
    m.createdBy = value
}
// SetCreatedDateTime sets the createdDateTime property value. Provides the dateTimeOffset for when the entity was created.
func (m *ItemFacet) SetCreatedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.createdDateTime = value
}
// SetInference sets the inference property value. Contains inference detail if the entity is inferred by the creating or modifying application.
func (m *ItemFacet) SetInference(value InferenceDataable)() {
    m.inference = value
}
// SetIsSearchable sets the isSearchable property value. The isSearchable property
func (m *ItemFacet) SetIsSearchable(value *bool)() {
    m.isSearchable = value
}
// SetLastModifiedBy sets the lastModifiedBy property value. The lastModifiedBy property
func (m *ItemFacet) SetLastModifiedBy(value IdentitySetable)() {
    m.lastModifiedBy = value
}
// SetLastModifiedDateTime sets the lastModifiedDateTime property value. Provides the dateTimeOffset for when the entity was created.
func (m *ItemFacet) SetLastModifiedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.lastModifiedDateTime = value
}
// SetSource sets the source property value. Where the values within an entity originated if synced from another service.
func (m *ItemFacet) SetSource(value PersonDataSourcesable)() {
    m.source = value
}
