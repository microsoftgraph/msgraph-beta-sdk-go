package ediscovery

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
)

// ReviewSetQuery provides operations to manage the collection of accessReview entities.
type ReviewSetQuery struct {
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Entity
    // The user who created the query.
    createdBy ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.IdentitySetable
    // The time and date when the query was created. The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z
    createdDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // The name of the query.
    displayName *string
    // The user who last modified the query.
    lastModifiedBy ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.IdentitySetable
    // The date and time the query was last modified. The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z
    lastModifiedDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // The query string in KQL (Keyword Query Language) query. For details, see Document metadata fields in Advanced eDiscovery.  This field maps directly to the keywords condition.  You can refine searches by using fields listed in the searchable field name paired with values; for example, subject:'Quarterly Financials' AND Date>=06/01/2016 AND Date<=07/01/2016.
    query *string
}
// NewReviewSetQuery instantiates a new reviewSetQuery and sets the default values.
func NewReviewSetQuery()(*ReviewSetQuery) {
    m := &ReviewSetQuery{
        Entity: *ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.NewEntity(),
    }
    odataTypeValue := "#microsoft.graph.ediscovery.reviewSetQuery";
    m.SetOdataType(&odataTypeValue);
    return m
}
// CreateReviewSetQueryFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateReviewSetQueryFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewReviewSetQuery(), nil
}
// GetCreatedBy gets the createdBy property value. The user who created the query.
func (m *ReviewSetQuery) GetCreatedBy()(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.IdentitySetable) {
    return m.createdBy
}
// GetCreatedDateTime gets the createdDateTime property value. The time and date when the query was created. The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z
func (m *ReviewSetQuery) GetCreatedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.createdDateTime
}
// GetDisplayName gets the displayName property value. The name of the query.
func (m *ReviewSetQuery) GetDisplayName()(*string) {
    return m.displayName
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ReviewSetQuery) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["createdBy"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetObjectValue(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateIdentitySetFromDiscriminatorValue , m.SetCreatedBy)
    res["createdDateTime"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetTimeValue(m.SetCreatedDateTime)
    res["displayName"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetDisplayName)
    res["lastModifiedBy"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetObjectValue(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateIdentitySetFromDiscriminatorValue , m.SetLastModifiedBy)
    res["lastModifiedDateTime"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetTimeValue(m.SetLastModifiedDateTime)
    res["query"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetQuery)
    return res
}
// GetLastModifiedBy gets the lastModifiedBy property value. The user who last modified the query.
func (m *ReviewSetQuery) GetLastModifiedBy()(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.IdentitySetable) {
    return m.lastModifiedBy
}
// GetLastModifiedDateTime gets the lastModifiedDateTime property value. The date and time the query was last modified. The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z
func (m *ReviewSetQuery) GetLastModifiedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.lastModifiedDateTime
}
// GetQuery gets the query property value. The query string in KQL (Keyword Query Language) query. For details, see Document metadata fields in Advanced eDiscovery.  This field maps directly to the keywords condition.  You can refine searches by using fields listed in the searchable field name paired with values; for example, subject:'Quarterly Financials' AND Date>=06/01/2016 AND Date<=07/01/2016.
func (m *ReviewSetQuery) GetQuery()(*string) {
    return m.query
}
// Serialize serializes information the current object
func (m *ReviewSetQuery) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
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
        err = writer.WriteStringValue("displayName", m.GetDisplayName())
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
        err = writer.WriteStringValue("query", m.GetQuery())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetCreatedBy sets the createdBy property value. The user who created the query.
func (m *ReviewSetQuery) SetCreatedBy(value ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.IdentitySetable)() {
    m.createdBy = value
}
// SetCreatedDateTime sets the createdDateTime property value. The time and date when the query was created. The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z
func (m *ReviewSetQuery) SetCreatedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.createdDateTime = value
}
// SetDisplayName sets the displayName property value. The name of the query.
func (m *ReviewSetQuery) SetDisplayName(value *string)() {
    m.displayName = value
}
// SetLastModifiedBy sets the lastModifiedBy property value. The user who last modified the query.
func (m *ReviewSetQuery) SetLastModifiedBy(value ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.IdentitySetable)() {
    m.lastModifiedBy = value
}
// SetLastModifiedDateTime sets the lastModifiedDateTime property value. The date and time the query was last modified. The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z
func (m *ReviewSetQuery) SetLastModifiedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.lastModifiedDateTime = value
}
// SetQuery sets the query property value. The query string in KQL (Keyword Query Language) query. For details, see Document metadata fields in Advanced eDiscovery.  This field maps directly to the keywords condition.  You can refine searches by using fields listed in the searchable field name paired with values; for example, subject:'Quarterly Financials' AND Date>=06/01/2016 AND Date<=07/01/2016.
func (m *ReviewSetQuery) SetQuery(value *string)() {
    m.query = value
}
