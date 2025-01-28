package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e "github.com/microsoft/kiota-abstractions-go/store"
)

// WindowsQualityUpdateCatalogProductRevision the operating system product revisions that are released as part of this quality update.
type WindowsQualityUpdateCatalogProductRevision struct {
    // Stores model information.
    backingStore ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore
}
// NewWindowsQualityUpdateCatalogProductRevision instantiates a new WindowsQualityUpdateCatalogProductRevision and sets the default values.
func NewWindowsQualityUpdateCatalogProductRevision()(*WindowsQualityUpdateCatalogProductRevision) {
    m := &WindowsQualityUpdateCatalogProductRevision{
    }
    m.backingStore = ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStoreFactoryInstance();
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateWindowsQualityUpdateCatalogProductRevisionFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateWindowsQualityUpdateCatalogProductRevisionFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewWindowsQualityUpdateCatalogProductRevision(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *WindowsQualityUpdateCatalogProductRevision) GetAdditionalData()(map[string]any) {
    val , err :=  m.backingStore.Get("additionalData")
    if err != nil {
        panic(err)
    }
    if val == nil {
        var value = make(map[string]any);
        m.SetAdditionalData(value);
    }
    return val.(map[string]any)
}
// GetBackingStore gets the BackingStore property value. Stores model information.
// returns a BackingStore when successful
func (m *WindowsQualityUpdateCatalogProductRevision) GetBackingStore()(ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore) {
    return m.backingStore
}
// GetDisplayName gets the displayName property value. The display name of the windows quality update catalog product revision. For example, 'Windows 11, version 22H2, build 22621.4112'. Read-only
// returns a *string when successful
func (m *WindowsQualityUpdateCatalogProductRevision) GetDisplayName()(*string) {
    val, err := m.GetBackingStore().Get("displayName")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *WindowsQualityUpdateCatalogProductRevision) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["displayName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDisplayName(val)
        }
        return nil
    }
    res["knowledgeBaseArticle"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateWindowsQualityUpdateProductKnowledgeBaseArticleFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetKnowledgeBaseArticle(val.(WindowsQualityUpdateProductKnowledgeBaseArticleable))
        }
        return nil
    }
    res["@odata.type"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOdataType(val)
        }
        return nil
    }
    res["osBuild"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateWindowsQualityUpdateProductBuildVersionDetailFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOsBuild(val.(WindowsQualityUpdateProductBuildVersionDetailable))
        }
        return nil
    }
    res["productName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetProductName(val)
        }
        return nil
    }
    res["releaseDateTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetReleaseDateTime(val)
        }
        return nil
    }
    res["versionName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetVersionName(val)
        }
        return nil
    }
    return res
}
// GetKnowledgeBaseArticle gets the knowledgeBaseArticle property value. Represents a knowledge base (KB) article.
// returns a WindowsQualityUpdateProductKnowledgeBaseArticleable when successful
func (m *WindowsQualityUpdateCatalogProductRevision) GetKnowledgeBaseArticle()(WindowsQualityUpdateProductKnowledgeBaseArticleable) {
    val, err := m.GetBackingStore().Get("knowledgeBaseArticle")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(WindowsQualityUpdateProductKnowledgeBaseArticleable)
    }
    return nil
}
// GetOdataType gets the @odata.type property value. The OdataType property
// returns a *string when successful
func (m *WindowsQualityUpdateCatalogProductRevision) GetOdataType()(*string) {
    val, err := m.GetBackingStore().Get("odataType")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetOsBuild gets the osBuild property value. Represents the build version details of a product revision that is associated with a quality update.
// returns a WindowsQualityUpdateProductBuildVersionDetailable when successful
func (m *WindowsQualityUpdateCatalogProductRevision) GetOsBuild()(WindowsQualityUpdateProductBuildVersionDetailable) {
    val, err := m.GetBackingStore().Get("osBuild")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(WindowsQualityUpdateProductBuildVersionDetailable)
    }
    return nil
}
// GetProductName gets the productName property value. The product name of the windows quality update catalog product revision. For example, 'Windows 11'. Read-only
// returns a *string when successful
func (m *WindowsQualityUpdateCatalogProductRevision) GetProductName()(*string) {
    val, err := m.GetBackingStore().Get("productName")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetReleaseDateTime gets the releaseDateTime property value. The date and time when the windows quality update catalog product revision was released. The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. Read-only
// returns a *Time when successful
func (m *WindowsQualityUpdateCatalogProductRevision) GetReleaseDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    val, err := m.GetBackingStore().Get("releaseDateTime")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    }
    return nil
}
// GetVersionName gets the versionName property value. The version name of the windows quality update catalog product revision. For example, '22H2'. Read-only
// returns a *string when successful
func (m *WindowsQualityUpdateCatalogProductRevision) GetVersionName()(*string) {
    val, err := m.GetBackingStore().Get("versionName")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// Serialize serializes information the current object
func (m *WindowsQualityUpdateCatalogProductRevision) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("displayName", m.GetDisplayName())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("knowledgeBaseArticle", m.GetKnowledgeBaseArticle())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("@odata.type", m.GetOdataType())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("osBuild", m.GetOsBuild())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("productName", m.GetProductName())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteTimeValue("releaseDateTime", m.GetReleaseDateTime())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("versionName", m.GetVersionName())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteAdditionalData(m.GetAdditionalData())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAdditionalData sets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *WindowsQualityUpdateCatalogProductRevision) SetAdditionalData(value map[string]any)() {
    err := m.GetBackingStore().Set("additionalData", value)
    if err != nil {
        panic(err)
    }
}
// SetBackingStore sets the BackingStore property value. Stores model information.
func (m *WindowsQualityUpdateCatalogProductRevision) SetBackingStore(value ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)() {
    m.backingStore = value
}
// SetDisplayName sets the displayName property value. The display name of the windows quality update catalog product revision. For example, 'Windows 11, version 22H2, build 22621.4112'. Read-only
func (m *WindowsQualityUpdateCatalogProductRevision) SetDisplayName(value *string)() {
    err := m.GetBackingStore().Set("displayName", value)
    if err != nil {
        panic(err)
    }
}
// SetKnowledgeBaseArticle sets the knowledgeBaseArticle property value. Represents a knowledge base (KB) article.
func (m *WindowsQualityUpdateCatalogProductRevision) SetKnowledgeBaseArticle(value WindowsQualityUpdateProductKnowledgeBaseArticleable)() {
    err := m.GetBackingStore().Set("knowledgeBaseArticle", value)
    if err != nil {
        panic(err)
    }
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *WindowsQualityUpdateCatalogProductRevision) SetOdataType(value *string)() {
    err := m.GetBackingStore().Set("odataType", value)
    if err != nil {
        panic(err)
    }
}
// SetOsBuild sets the osBuild property value. Represents the build version details of a product revision that is associated with a quality update.
func (m *WindowsQualityUpdateCatalogProductRevision) SetOsBuild(value WindowsQualityUpdateProductBuildVersionDetailable)() {
    err := m.GetBackingStore().Set("osBuild", value)
    if err != nil {
        panic(err)
    }
}
// SetProductName sets the productName property value. The product name of the windows quality update catalog product revision. For example, 'Windows 11'. Read-only
func (m *WindowsQualityUpdateCatalogProductRevision) SetProductName(value *string)() {
    err := m.GetBackingStore().Set("productName", value)
    if err != nil {
        panic(err)
    }
}
// SetReleaseDateTime sets the releaseDateTime property value. The date and time when the windows quality update catalog product revision was released. The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. Read-only
func (m *WindowsQualityUpdateCatalogProductRevision) SetReleaseDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    err := m.GetBackingStore().Set("releaseDateTime", value)
    if err != nil {
        panic(err)
    }
}
// SetVersionName sets the versionName property value. The version name of the windows quality update catalog product revision. For example, '22H2'. Read-only
func (m *WindowsQualityUpdateCatalogProductRevision) SetVersionName(value *string)() {
    err := m.GetBackingStore().Set("versionName", value)
    if err != nil {
        panic(err)
    }
}
type WindowsQualityUpdateCatalogProductRevisionable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackedModel
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetBackingStore()(ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)
    GetDisplayName()(*string)
    GetKnowledgeBaseArticle()(WindowsQualityUpdateProductKnowledgeBaseArticleable)
    GetOdataType()(*string)
    GetOsBuild()(WindowsQualityUpdateProductBuildVersionDetailable)
    GetProductName()(*string)
    GetReleaseDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetVersionName()(*string)
    SetBackingStore(value ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)()
    SetDisplayName(value *string)()
    SetKnowledgeBaseArticle(value WindowsQualityUpdateProductKnowledgeBaseArticleable)()
    SetOdataType(value *string)()
    SetOsBuild(value WindowsQualityUpdateProductBuildVersionDetailable)()
    SetProductName(value *string)()
    SetReleaseDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetVersionName(value *string)()
}
