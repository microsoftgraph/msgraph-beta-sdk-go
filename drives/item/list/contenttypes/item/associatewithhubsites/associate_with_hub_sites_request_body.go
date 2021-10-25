package associatewithhubsites

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

type AssociateWithHubSitesRequestBody struct {
    additionalData map[string]interface{};
    hubSiteUrls []string;
    propagateToExistingLists *bool;
}
func NewAssociateWithHubSitesRequestBody()(*AssociateWithHubSitesRequestBody) {
    m := &AssociateWithHubSitesRequestBody{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
func (m *AssociateWithHubSitesRequestBody) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
func (m *AssociateWithHubSitesRequestBody) GetHubSiteUrls()([]string) {
    if m == nil {
        return nil
    } else {
        return m.hubSiteUrls
    }
}
func (m *AssociateWithHubSitesRequestBody) GetPropagateToExistingLists()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.propagateToExistingLists
    }
}
func (m *AssociateWithHubSitesRequestBody) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["hubSiteUrls"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        res := make([]string, len(val))
        for i, v := range val {
            res[i] = v.(string)
        }
        m.SetHubSiteUrls(res)
        return nil
    }
    res["propagateToExistingLists"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        m.SetPropagateToExistingLists(val)
        return nil
    }
    return res
}
func (m *AssociateWithHubSitesRequestBody) IsNil()(bool) {
    return m == nil
}
func (m *AssociateWithHubSitesRequestBody) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    {
        err := writer.WriteCollectionOfStringValues("hubSiteUrls", m.GetHubSiteUrls())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("propagateToExistingLists", m.GetPropagateToExistingLists())
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
func (m *AssociateWithHubSitesRequestBody) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
func (m *AssociateWithHubSitesRequestBody) SetHubSiteUrls(value []string)() {
    m.hubSiteUrls = value
}
func (m *AssociateWithHubSitesRequestBody) SetPropagateToExistingLists(value *bool)() {
    m.propagateToExistingLists = value
}
