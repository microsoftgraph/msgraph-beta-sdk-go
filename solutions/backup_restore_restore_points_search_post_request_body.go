package solutions

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e "github.com/microsoft/kiota-abstractions-go/store"
)

type BackupRestoreRestorePointsSearchPostRequestBody struct {
    // Stores model information.
    backingStore ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore
}
// NewBackupRestoreRestorePointsSearchPostRequestBody instantiates a new BackupRestoreRestorePointsSearchPostRequestBody and sets the default values.
func NewBackupRestoreRestorePointsSearchPostRequestBody()(*BackupRestoreRestorePointsSearchPostRequestBody) {
    m := &BackupRestoreRestorePointsSearchPostRequestBody{
    }
    m.backingStore = ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStoreFactoryInstance();
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateBackupRestoreRestorePointsSearchPostRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateBackupRestoreRestorePointsSearchPostRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewBackupRestoreRestorePointsSearchPostRequestBody(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *BackupRestoreRestorePointsSearchPostRequestBody) GetAdditionalData()(map[string]any) {
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
// GetArtifactQuery gets the artifactQuery property value. The artifactQuery property
// returns a ArtifactQueryable when successful
func (m *BackupRestoreRestorePointsSearchPostRequestBody) GetArtifactQuery()(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ArtifactQueryable) {
    val, err := m.GetBackingStore().Get("artifactQuery")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ArtifactQueryable)
    }
    return nil
}
// GetBackingStore gets the BackingStore property value. Stores model information.
// returns a BackingStore when successful
func (m *BackupRestoreRestorePointsSearchPostRequestBody) GetBackingStore()(ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore) {
    return m.backingStore
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *BackupRestoreRestorePointsSearchPostRequestBody) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["artifactQuery"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateArtifactQueryFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetArtifactQuery(val.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ArtifactQueryable))
        }
        return nil
    }
    res["protectionTimePeriod"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateTimePeriodFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetProtectionTimePeriod(val.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.TimePeriodable))
        }
        return nil
    }
    res["protectionUnitIds"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]string, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = *(v.(*string))
                }
            }
            m.SetProtectionUnitIds(res)
        }
        return nil
    }
    res["restorePointPreference"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ParseRestorePointPreference)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRestorePointPreference(val.(*ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.RestorePointPreference))
        }
        return nil
    }
    res["tags"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ParseRestorePointTags)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTags(val.(*ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.RestorePointTags))
        }
        return nil
    }
    return res
}
// GetProtectionTimePeriod gets the protectionTimePeriod property value. The protectionTimePeriod property
// returns a TimePeriodable when successful
func (m *BackupRestoreRestorePointsSearchPostRequestBody) GetProtectionTimePeriod()(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.TimePeriodable) {
    val, err := m.GetBackingStore().Get("protectionTimePeriod")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.TimePeriodable)
    }
    return nil
}
// GetProtectionUnitIds gets the protectionUnitIds property value. The protectionUnitIds property
// returns a []string when successful
func (m *BackupRestoreRestorePointsSearchPostRequestBody) GetProtectionUnitIds()([]string) {
    val, err := m.GetBackingStore().Get("protectionUnitIds")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]string)
    }
    return nil
}
// GetRestorePointPreference gets the restorePointPreference property value. The restorePointPreference property
// returns a *RestorePointPreference when successful
func (m *BackupRestoreRestorePointsSearchPostRequestBody) GetRestorePointPreference()(*ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.RestorePointPreference) {
    val, err := m.GetBackingStore().Get("restorePointPreference")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.RestorePointPreference)
    }
    return nil
}
// GetTags gets the tags property value. The tags property
// returns a *RestorePointTags when successful
func (m *BackupRestoreRestorePointsSearchPostRequestBody) GetTags()(*ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.RestorePointTags) {
    val, err := m.GetBackingStore().Get("tags")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.RestorePointTags)
    }
    return nil
}
// Serialize serializes information the current object
func (m *BackupRestoreRestorePointsSearchPostRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteObjectValue("artifactQuery", m.GetArtifactQuery())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("protectionTimePeriod", m.GetProtectionTimePeriod())
        if err != nil {
            return err
        }
    }
    if m.GetProtectionUnitIds() != nil {
        err := writer.WriteCollectionOfStringValues("protectionUnitIds", m.GetProtectionUnitIds())
        if err != nil {
            return err
        }
    }
    if m.GetRestorePointPreference() != nil {
        cast := (*m.GetRestorePointPreference()).String()
        err := writer.WriteStringValue("restorePointPreference", &cast)
        if err != nil {
            return err
        }
    }
    if m.GetTags() != nil {
        cast := (*m.GetTags()).String()
        err := writer.WriteStringValue("tags", &cast)
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
func (m *BackupRestoreRestorePointsSearchPostRequestBody) SetAdditionalData(value map[string]any)() {
    err := m.GetBackingStore().Set("additionalData", value)
    if err != nil {
        panic(err)
    }
}
// SetArtifactQuery sets the artifactQuery property value. The artifactQuery property
func (m *BackupRestoreRestorePointsSearchPostRequestBody) SetArtifactQuery(value ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ArtifactQueryable)() {
    err := m.GetBackingStore().Set("artifactQuery", value)
    if err != nil {
        panic(err)
    }
}
// SetBackingStore sets the BackingStore property value. Stores model information.
func (m *BackupRestoreRestorePointsSearchPostRequestBody) SetBackingStore(value ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)() {
    m.backingStore = value
}
// SetProtectionTimePeriod sets the protectionTimePeriod property value. The protectionTimePeriod property
func (m *BackupRestoreRestorePointsSearchPostRequestBody) SetProtectionTimePeriod(value ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.TimePeriodable)() {
    err := m.GetBackingStore().Set("protectionTimePeriod", value)
    if err != nil {
        panic(err)
    }
}
// SetProtectionUnitIds sets the protectionUnitIds property value. The protectionUnitIds property
func (m *BackupRestoreRestorePointsSearchPostRequestBody) SetProtectionUnitIds(value []string)() {
    err := m.GetBackingStore().Set("protectionUnitIds", value)
    if err != nil {
        panic(err)
    }
}
// SetRestorePointPreference sets the restorePointPreference property value. The restorePointPreference property
func (m *BackupRestoreRestorePointsSearchPostRequestBody) SetRestorePointPreference(value *ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.RestorePointPreference)() {
    err := m.GetBackingStore().Set("restorePointPreference", value)
    if err != nil {
        panic(err)
    }
}
// SetTags sets the tags property value. The tags property
func (m *BackupRestoreRestorePointsSearchPostRequestBody) SetTags(value *ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.RestorePointTags)() {
    err := m.GetBackingStore().Set("tags", value)
    if err != nil {
        panic(err)
    }
}
type BackupRestoreRestorePointsSearchPostRequestBodyable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackedModel
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetArtifactQuery()(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ArtifactQueryable)
    GetBackingStore()(ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)
    GetProtectionTimePeriod()(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.TimePeriodable)
    GetProtectionUnitIds()([]string)
    GetRestorePointPreference()(*ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.RestorePointPreference)
    GetTags()(*ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.RestorePointTags)
    SetArtifactQuery(value ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ArtifactQueryable)()
    SetBackingStore(value ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)()
    SetProtectionTimePeriod(value ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.TimePeriodable)()
    SetProtectionUnitIds(value []string)()
    SetRestorePointPreference(value *ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.RestorePointPreference)()
    SetTags(value *ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.RestorePointTags)()
}
