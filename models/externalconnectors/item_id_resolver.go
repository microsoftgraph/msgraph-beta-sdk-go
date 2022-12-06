package externalconnectors

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ItemIdResolver 
type ItemIdResolver struct {
    UrlToItemResolverBase
    // Pattern that specifies how to form the ID of the external item that the URL represents. The named groups from the regular expression in urlPattern within the urlMatchInfo can be referenced by inserting the group name inside curly brackets.
    itemId *string
    // Configurations to match and resolve URL.
    urlMatchInfo UrlMatchInfoable
}
// NewItemIdResolver instantiates a new ItemIdResolver and sets the default values.
func NewItemIdResolver()(*ItemIdResolver) {
    m := &ItemIdResolver{
        UrlToItemResolverBase: *NewUrlToItemResolverBase(),
    }
    odataTypeValue := "#microsoft.graph.externalConnectors.itemIdResolver";
    m.SetOdataType(&odataTypeValue);
    return m
}
// CreateItemIdResolverFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateItemIdResolverFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewItemIdResolver(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ItemIdResolver) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.UrlToItemResolverBase.GetFieldDeserializers()
    res["itemId"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetItemId)
    res["urlMatchInfo"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetObjectValue(CreateUrlMatchInfoFromDiscriminatorValue , m.SetUrlMatchInfo)
    return res
}
// GetItemId gets the itemId property value. Pattern that specifies how to form the ID of the external item that the URL represents. The named groups from the regular expression in urlPattern within the urlMatchInfo can be referenced by inserting the group name inside curly brackets.
func (m *ItemIdResolver) GetItemId()(*string) {
    return m.itemId
}
// GetUrlMatchInfo gets the urlMatchInfo property value. Configurations to match and resolve URL.
func (m *ItemIdResolver) GetUrlMatchInfo()(UrlMatchInfoable) {
    return m.urlMatchInfo
}
// Serialize serializes information the current object
func (m *ItemIdResolver) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.UrlToItemResolverBase.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteStringValue("itemId", m.GetItemId())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("urlMatchInfo", m.GetUrlMatchInfo())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetItemId sets the itemId property value. Pattern that specifies how to form the ID of the external item that the URL represents. The named groups from the regular expression in urlPattern within the urlMatchInfo can be referenced by inserting the group name inside curly brackets.
func (m *ItemIdResolver) SetItemId(value *string)() {
    m.itemId = value
}
// SetUrlMatchInfo sets the urlMatchInfo property value. Configurations to match and resolve URL.
func (m *ItemIdResolver) SetUrlMatchInfo(value UrlMatchInfoable)() {
    m.urlMatchInfo = value
}
