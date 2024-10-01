package healthmonitoring

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
)

type DirectoryObjectImpactSummary struct {
    ResourceImpactSummary
}
// NewDirectoryObjectImpactSummary instantiates a new DirectoryObjectImpactSummary and sets the default values.
func NewDirectoryObjectImpactSummary()(*DirectoryObjectImpactSummary) {
    m := &DirectoryObjectImpactSummary{
        ResourceImpactSummary: *NewResourceImpactSummary(),
    }
    odataTypeValue := "#microsoft.graph.healthMonitoring.directoryObjectImpactSummary"
    m.SetOdataType(&odataTypeValue)
    return m
}
// CreateDirectoryObjectImpactSummaryFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateDirectoryObjectImpactSummaryFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
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
                    case "#microsoft.graph.healthMonitoring.applicationImpactSummary":
                        return NewApplicationImpactSummary(), nil
                    case "#microsoft.graph.healthMonitoring.deviceImpactSummary":
                        return NewDeviceImpactSummary(), nil
                    case "#microsoft.graph.healthMonitoring.groupImpactSummary":
                        return NewGroupImpactSummary(), nil
                    case "#microsoft.graph.healthMonitoring.servicePrincipalImpactSummary":
                        return NewServicePrincipalImpactSummary(), nil
                    case "#microsoft.graph.healthMonitoring.userImpactSummary":
                        return NewUserImpactSummary(), nil
                }
            }
        }
    }
    return NewDirectoryObjectImpactSummary(), nil
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *DirectoryObjectImpactSummary) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.ResourceImpactSummary.GetFieldDeserializers()
    res["resourceSampling"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateDirectoryObjectFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DirectoryObjectable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DirectoryObjectable)
                }
            }
            m.SetResourceSampling(res)
        }
        return nil
    }
    return res
}
// GetResourceSampling gets the resourceSampling property value. The resourceSampling property
// returns a []DirectoryObjectable when successful
func (m *DirectoryObjectImpactSummary) GetResourceSampling()([]ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DirectoryObjectable) {
    val, err := m.GetBackingStore().Get("resourceSampling")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DirectoryObjectable)
    }
    return nil
}
// Serialize serializes information the current object
func (m *DirectoryObjectImpactSummary) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.ResourceImpactSummary.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetResourceSampling() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetResourceSampling()))
        for i, v := range m.GetResourceSampling() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err = writer.WriteCollectionOfObjectValues("resourceSampling", cast)
        if err != nil {
            return err
        }
    }
    return nil
}
// SetResourceSampling sets the resourceSampling property value. The resourceSampling property
func (m *DirectoryObjectImpactSummary) SetResourceSampling(value []ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DirectoryObjectable)() {
    err := m.GetBackingStore().Set("resourceSampling", value)
    if err != nil {
        panic(err)
    }
}
type DirectoryObjectImpactSummaryable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    ResourceImpactSummaryable
    GetResourceSampling()([]ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DirectoryObjectable)
    SetResourceSampling(value []ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DirectoryObjectable)()
}
