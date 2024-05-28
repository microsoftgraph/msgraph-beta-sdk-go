package groups

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e "github.com/microsoft/kiota-abstractions-go/store"
)

type ItemSitesItemListsItemContenttypesItemCopytodefaultcontentlocationCopyToDefaultContentLocationPostRequestBody struct {
    // Stores model information.
    backingStore ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore
}
// NewItemSitesItemListsItemContenttypesItemCopytodefaultcontentlocationCopyToDefaultContentLocationPostRequestBody instantiates a new ItemSitesItemListsItemContenttypesItemCopytodefaultcontentlocationCopyToDefaultContentLocationPostRequestBody and sets the default values.
func NewItemSitesItemListsItemContenttypesItemCopytodefaultcontentlocationCopyToDefaultContentLocationPostRequestBody()(*ItemSitesItemListsItemContenttypesItemCopytodefaultcontentlocationCopyToDefaultContentLocationPostRequestBody) {
    m := &ItemSitesItemListsItemContenttypesItemCopytodefaultcontentlocationCopyToDefaultContentLocationPostRequestBody{
    }
    m.backingStore = ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStoreFactoryInstance();
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateItemSitesItemListsItemContenttypesItemCopytodefaultcontentlocationCopyToDefaultContentLocationPostRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateItemSitesItemListsItemContenttypesItemCopytodefaultcontentlocationCopyToDefaultContentLocationPostRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewItemSitesItemListsItemContenttypesItemCopytodefaultcontentlocationCopyToDefaultContentLocationPostRequestBody(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *ItemSitesItemListsItemContenttypesItemCopytodefaultcontentlocationCopyToDefaultContentLocationPostRequestBody) GetAdditionalData()(map[string]any) {
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
func (m *ItemSitesItemListsItemContenttypesItemCopytodefaultcontentlocationCopyToDefaultContentLocationPostRequestBody) GetBackingStore()(ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore) {
    return m.backingStore
}
// GetDestinationFileName gets the destinationFileName property value. The destinationFileName property
// returns a *string when successful
func (m *ItemSitesItemListsItemContenttypesItemCopytodefaultcontentlocationCopyToDefaultContentLocationPostRequestBody) GetDestinationFileName()(*string) {
    val, err := m.GetBackingStore().Get("destinationFileName")
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
func (m *ItemSitesItemListsItemContenttypesItemCopytodefaultcontentlocationCopyToDefaultContentLocationPostRequestBody) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["destinationFileName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDestinationFileName(val)
        }
        return nil
    }
    res["sourceFile"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateItemReferenceFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSourceFile(val.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ItemReferenceable))
        }
        return nil
    }
    return res
}
// GetSourceFile gets the sourceFile property value. The sourceFile property
// returns a ItemReferenceable when successful
func (m *ItemSitesItemListsItemContenttypesItemCopytodefaultcontentlocationCopyToDefaultContentLocationPostRequestBody) GetSourceFile()(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ItemReferenceable) {
    val, err := m.GetBackingStore().Get("sourceFile")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ItemReferenceable)
    }
    return nil
}
// Serialize serializes information the current object
func (m *ItemSitesItemListsItemContenttypesItemCopytodefaultcontentlocationCopyToDefaultContentLocationPostRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("destinationFileName", m.GetDestinationFileName())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("sourceFile", m.GetSourceFile())
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
func (m *ItemSitesItemListsItemContenttypesItemCopytodefaultcontentlocationCopyToDefaultContentLocationPostRequestBody) SetAdditionalData(value map[string]any)() {
    err := m.GetBackingStore().Set("additionalData", value)
    if err != nil {
        panic(err)
    }
}
// SetBackingStore sets the BackingStore property value. Stores model information.
func (m *ItemSitesItemListsItemContenttypesItemCopytodefaultcontentlocationCopyToDefaultContentLocationPostRequestBody) SetBackingStore(value ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)() {
    m.backingStore = value
}
// SetDestinationFileName sets the destinationFileName property value. The destinationFileName property
func (m *ItemSitesItemListsItemContenttypesItemCopytodefaultcontentlocationCopyToDefaultContentLocationPostRequestBody) SetDestinationFileName(value *string)() {
    err := m.GetBackingStore().Set("destinationFileName", value)
    if err != nil {
        panic(err)
    }
}
// SetSourceFile sets the sourceFile property value. The sourceFile property
func (m *ItemSitesItemListsItemContenttypesItemCopytodefaultcontentlocationCopyToDefaultContentLocationPostRequestBody) SetSourceFile(value ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ItemReferenceable)() {
    err := m.GetBackingStore().Set("sourceFile", value)
    if err != nil {
        panic(err)
    }
}
type ItemSitesItemListsItemContenttypesItemCopytodefaultcontentlocationCopyToDefaultContentLocationPostRequestBodyable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackedModel
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetBackingStore()(ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)
    GetDestinationFileName()(*string)
    GetSourceFile()(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ItemReferenceable)
    SetBackingStore(value ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)()
    SetDestinationFileName(value *string)()
    SetSourceFile(value ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ItemReferenceable)()
}
