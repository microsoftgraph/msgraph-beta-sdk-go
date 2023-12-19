package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ItemFacet 
type ItemFacet struct {
    Entity
}
// NewItemFacet instantiates a new itemFacet and sets the default values.
func NewItemFacet()(*ItemFacet) {
    m := &ItemFacet{
        Entity: *NewEntity(),
    }
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
func (m *ItemFacet) GetAllowedAudiences()(*ItemFacet_allowedAudiences) {
    val, err := m.GetBackingStore().Get("allowedAudiences")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*ItemFacet_allowedAudiences)
    }
    return nil
}
// GetCreatedBy gets the createdBy property value. The createdBy property
func (m *ItemFacet) GetCreatedBy()(IdentitySetable) {
    val, err := m.GetBackingStore().Get("createdBy")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(IdentitySetable)
    }
    return nil
}
// GetCreatedDateTime gets the createdDateTime property value. Provides the dateTimeOffset for when the entity was created.
func (m *ItemFacet) GetCreatedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    val, err := m.GetBackingStore().Get("createdDateTime")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    }
    return nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ItemFacet) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["allowedAudiences"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseItemFacet_allowedAudiences)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAllowedAudiences(val.(*ItemFacet_allowedAudiences))
        }
        return nil
    }
    res["createdBy"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateIdentitySetFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCreatedBy(val.(IdentitySetable))
        }
        return nil
    }
    res["createdDateTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCreatedDateTime(val)
        }
        return nil
    }
    res["inference"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateInferenceDataFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetInference(val.(InferenceDataable))
        }
        return nil
    }
    res["isSearchable"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsSearchable(val)
        }
        return nil
    }
    res["lastModifiedBy"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateIdentitySetFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLastModifiedBy(val.(IdentitySetable))
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
    res["source"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreatePersonDataSourcesFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSource(val.(PersonDataSourcesable))
        }
        return nil
    }
    return res
}
// GetInference gets the inference property value. Contains inference detail if the entity is inferred by the creating or modifying application.
func (m *ItemFacet) GetInference()(InferenceDataable) {
    val, err := m.GetBackingStore().Get("inference")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(InferenceDataable)
    }
    return nil
}
// GetIsSearchable gets the isSearchable property value. The isSearchable property
func (m *ItemFacet) GetIsSearchable()(*bool) {
    val, err := m.GetBackingStore().Get("isSearchable")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// GetLastModifiedBy gets the lastModifiedBy property value. The lastModifiedBy property
func (m *ItemFacet) GetLastModifiedBy()(IdentitySetable) {
    val, err := m.GetBackingStore().Get("lastModifiedBy")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(IdentitySetable)
    }
    return nil
}
// GetLastModifiedDateTime gets the lastModifiedDateTime property value. Provides the dateTimeOffset for when the entity was created.
func (m *ItemFacet) GetLastModifiedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    val, err := m.GetBackingStore().Get("lastModifiedDateTime")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    }
    return nil
}
// GetSource gets the source property value. Where the values within an entity originated if synced from another service.
func (m *ItemFacet) GetSource()(PersonDataSourcesable) {
    val, err := m.GetBackingStore().Get("source")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(PersonDataSourcesable)
    }
    return nil
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
func (m *ItemFacet) SetAllowedAudiences(value *ItemFacet_allowedAudiences)() {
    err := m.GetBackingStore().Set("allowedAudiences", value)
    if err != nil {
        panic(err)
    }
}
// SetCreatedBy sets the createdBy property value. The createdBy property
func (m *ItemFacet) SetCreatedBy(value IdentitySetable)() {
    err := m.GetBackingStore().Set("createdBy", value)
    if err != nil {
        panic(err)
    }
}
// SetCreatedDateTime sets the createdDateTime property value. Provides the dateTimeOffset for when the entity was created.
func (m *ItemFacet) SetCreatedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    err := m.GetBackingStore().Set("createdDateTime", value)
    if err != nil {
        panic(err)
    }
}
// SetInference sets the inference property value. Contains inference detail if the entity is inferred by the creating or modifying application.
func (m *ItemFacet) SetInference(value InferenceDataable)() {
    err := m.GetBackingStore().Set("inference", value)
    if err != nil {
        panic(err)
    }
}
// SetIsSearchable sets the isSearchable property value. The isSearchable property
func (m *ItemFacet) SetIsSearchable(value *bool)() {
    err := m.GetBackingStore().Set("isSearchable", value)
    if err != nil {
        panic(err)
    }
}
// SetLastModifiedBy sets the lastModifiedBy property value. The lastModifiedBy property
func (m *ItemFacet) SetLastModifiedBy(value IdentitySetable)() {
    err := m.GetBackingStore().Set("lastModifiedBy", value)
    if err != nil {
        panic(err)
    }
}
// SetLastModifiedDateTime sets the lastModifiedDateTime property value. Provides the dateTimeOffset for when the entity was created.
func (m *ItemFacet) SetLastModifiedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    err := m.GetBackingStore().Set("lastModifiedDateTime", value)
    if err != nil {
        panic(err)
    }
}
// SetSource sets the source property value. Where the values within an entity originated if synced from another service.
func (m *ItemFacet) SetSource(value PersonDataSourcesable)() {
    err := m.GetBackingStore().Set("source", value)
    if err != nil {
        panic(err)
    }
}
// ItemFacetable 
type ItemFacetable interface {
    Entityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAllowedAudiences()(*ItemFacet_allowedAudiences)
    GetCreatedBy()(IdentitySetable)
    GetCreatedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetInference()(InferenceDataable)
    GetIsSearchable()(*bool)
    GetLastModifiedBy()(IdentitySetable)
    GetLastModifiedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetSource()(PersonDataSourcesable)
    SetAllowedAudiences(value *ItemFacet_allowedAudiences)()
    SetCreatedBy(value IdentitySetable)()
    SetCreatedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetInference(value InferenceDataable)()
    SetIsSearchable(value *bool)()
    SetLastModifiedBy(value IdentitySetable)()
    SetLastModifiedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetSource(value PersonDataSourcesable)()
}
