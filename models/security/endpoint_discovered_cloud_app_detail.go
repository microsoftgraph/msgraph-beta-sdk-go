package security

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type EndpointDiscoveredCloudAppDetail struct {
    DiscoveredCloudAppDetail
}
// NewEndpointDiscoveredCloudAppDetail instantiates a new EndpointDiscoveredCloudAppDetail and sets the default values.
func NewEndpointDiscoveredCloudAppDetail()(*EndpointDiscoveredCloudAppDetail) {
    m := &EndpointDiscoveredCloudAppDetail{
        DiscoveredCloudAppDetail: *NewDiscoveredCloudAppDetail(),
    }
    return m
}
// CreateEndpointDiscoveredCloudAppDetailFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateEndpointDiscoveredCloudAppDetailFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewEndpointDiscoveredCloudAppDetail(), nil
}
// GetDeviceCount gets the deviceCount property value. The number of devices that accessed the discovered app.
// returns a *int64 when successful
func (m *EndpointDiscoveredCloudAppDetail) GetDeviceCount()(*int64) {
    val, err := m.GetBackingStore().Get("deviceCount")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*int64)
    }
    return nil
}
// GetDevices gets the devices property value. Represents the devices that access the discovered apps.
// returns a []DiscoveredCloudAppDeviceable when successful
func (m *EndpointDiscoveredCloudAppDetail) GetDevices()([]DiscoveredCloudAppDeviceable) {
    val, err := m.GetBackingStore().Get("devices")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]DiscoveredCloudAppDeviceable)
    }
    return nil
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *EndpointDiscoveredCloudAppDetail) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.DiscoveredCloudAppDetail.GetFieldDeserializers()
    res["deviceCount"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDeviceCount(val)
        }
        return nil
    }
    res["devices"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateDiscoveredCloudAppDeviceFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]DiscoveredCloudAppDeviceable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(DiscoveredCloudAppDeviceable)
                }
            }
            m.SetDevices(res)
        }
        return nil
    }
    return res
}
// Serialize serializes information the current object
func (m *EndpointDiscoveredCloudAppDetail) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.DiscoveredCloudAppDetail.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteInt64Value("deviceCount", m.GetDeviceCount())
        if err != nil {
            return err
        }
    }
    if m.GetDevices() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetDevices()))
        for i, v := range m.GetDevices() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err = writer.WriteCollectionOfObjectValues("devices", cast)
        if err != nil {
            return err
        }
    }
    return nil
}
// SetDeviceCount sets the deviceCount property value. The number of devices that accessed the discovered app.
func (m *EndpointDiscoveredCloudAppDetail) SetDeviceCount(value *int64)() {
    err := m.GetBackingStore().Set("deviceCount", value)
    if err != nil {
        panic(err)
    }
}
// SetDevices sets the devices property value. Represents the devices that access the discovered apps.
func (m *EndpointDiscoveredCloudAppDetail) SetDevices(value []DiscoveredCloudAppDeviceable)() {
    err := m.GetBackingStore().Set("devices", value)
    if err != nil {
        panic(err)
    }
}
type EndpointDiscoveredCloudAppDetailable interface {
    DiscoveredCloudAppDetailable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetDeviceCount()(*int64)
    GetDevices()([]DiscoveredCloudAppDeviceable)
    SetDeviceCount(value *int64)()
    SetDevices(value []DiscoveredCloudAppDeviceable)()
}
