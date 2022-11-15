package search

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
)

// SearchAnswer provides operations to manage the collection of accessReviewDecision entities.
type SearchAnswer struct {
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Entity
    // Search answer description shown on search results page.
    description *string
    // Search answer name displayed in search results.
    displayName *string
    // Details of the user that created or last modified the search answer. Read-only.
    lastModifiedBy IdentitySetable
    // Timestamp of when the search answer is created or edited. Read-only.
    lastModifiedDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // Search answer URL link. When users click this search answer in search results, they will go to this URL.
    webUrl *string
}
// NewSearchAnswer instantiates a new searchAnswer and sets the default values.
func NewSearchAnswer()(*SearchAnswer) {
    m := &SearchAnswer{
        Entity: *ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.NewEntity(),
    }
    odataTypeValue := "#microsoft.graph.search.searchAnswer";
    m.SetOdataType(&odataTypeValue);
    return m
}
// CreateSearchAnswerFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateSearchAnswerFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
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
                    case "#microsoft.graph.search.acronym":
                        return NewAcronym(), nil
                    case "#microsoft.graph.search.bookmark":
                        return NewBookmark(), nil
                    case "#microsoft.graph.search.qna":
                        return NewQna(), nil
                }
            }
        }
    }
    return NewSearchAnswer(), nil
}
// GetDescription gets the description property value. Search answer description shown on search results page.
func (m *SearchAnswer) GetDescription()(*string) {
    return m.description
}
// GetDisplayName gets the displayName property value. Search answer name displayed in search results.
func (m *SearchAnswer) GetDisplayName()(*string) {
    return m.displayName
}
// GetFieldDeserializers the deserialization information for the current model
func (m *SearchAnswer) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["description"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetDescription)
    res["displayName"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetDisplayName)
    res["lastModifiedBy"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetObjectValue(CreateIdentitySetFromDiscriminatorValue , m.SetLastModifiedBy)
    res["lastModifiedDateTime"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetTimeValue(m.SetLastModifiedDateTime)
    res["webUrl"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetWebUrl)
    return res
}
// GetLastModifiedBy gets the lastModifiedBy property value. Details of the user that created or last modified the search answer. Read-only.
func (m *SearchAnswer) GetLastModifiedBy()(IdentitySetable) {
    return m.lastModifiedBy
}
// GetLastModifiedDateTime gets the lastModifiedDateTime property value. Timestamp of when the search answer is created or edited. Read-only.
func (m *SearchAnswer) GetLastModifiedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.lastModifiedDateTime
}
// GetWebUrl gets the webUrl property value. Search answer URL link. When users click this search answer in search results, they will go to this URL.
func (m *SearchAnswer) GetWebUrl()(*string) {
    return m.webUrl
}
// Serialize serializes information the current object
func (m *SearchAnswer) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteStringValue("description", m.GetDescription())
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
        err = writer.WriteStringValue("webUrl", m.GetWebUrl())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetDescription sets the description property value. Search answer description shown on search results page.
func (m *SearchAnswer) SetDescription(value *string)() {
    m.description = value
}
// SetDisplayName sets the displayName property value. Search answer name displayed in search results.
func (m *SearchAnswer) SetDisplayName(value *string)() {
    m.displayName = value
}
// SetLastModifiedBy sets the lastModifiedBy property value. Details of the user that created or last modified the search answer. Read-only.
func (m *SearchAnswer) SetLastModifiedBy(value IdentitySetable)() {
    m.lastModifiedBy = value
}
// SetLastModifiedDateTime sets the lastModifiedDateTime property value. Timestamp of when the search answer is created or edited. Read-only.
func (m *SearchAnswer) SetLastModifiedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.lastModifiedDateTime = value
}
// SetWebUrl sets the webUrl property value. Search answer URL link. When users click this search answer in search results, they will go to this URL.
func (m *SearchAnswer) SetWebUrl(value *string)() {
    m.webUrl = value
}
