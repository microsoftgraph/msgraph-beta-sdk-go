package comanageddevices

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
    i0952f14bd3d84883aa9509ce20c436cb3a9becc286ff26416d1a9ab4f9a5188e "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/comanageddevices/executeaction"
    i21b46d2a36281265e82fedbdddef1c20e278c5a4bbf3cf85fa0d860eccee5b55 "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/comanageddevices/movedevicestoou"
    i7e96e1b990aeb4d47dec870ab144f92eef626fd02edc82312e6ebb7fb40dbd1b "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/comanageddevices/count"
    i7fe672ce2ded63745de2480f0a175b485195547bff2ac413663b22ef3172bb4e "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/comanageddevices/bulkreprovisioncloudpc"
    i97970601371fcbb5ba7ab85926c15aaf372bd11b13e0e69958d51b9047db72f5 "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/comanageddevices/bulkrestorecloudpc"
)

// ComanagedDevicesRequestBuilder provides operations to manage the comanagedDevices property of the microsoft.graph.deviceManagement entity.
type ComanagedDevicesRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// ComanagedDevicesRequestBuilderGetOptions options for Get
type ComanagedDevicesRequestBuilderGetOptions struct {
    // Request headers
    Headers map[string]string;
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption;
    // Request query parameters
    QueryParameters *ComanagedDevicesRequestBuilderGetQueryParameters;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler;
}
// ComanagedDevicesRequestBuilderGetQueryParameters the list of co-managed devices report
type ComanagedDevicesRequestBuilderGetQueryParameters struct {
    // Include count of items
    Count *bool;
    // Expand related entities
    Expand []string;
    // Filter items by property values
    Filter *string;
    // Order items by property values
    Orderby []string;
    // Search items by search phrases
    Search *string;
    // Select properties to be returned
    Select []string;
    // Skip the first n items
    Skip *int32;
    // Show only the first n items
    Top *int32;
}
// ComanagedDevicesRequestBuilderPostOptions options for Post
type ComanagedDevicesRequestBuilderPostOptions struct {
    // 
    Body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ManagedDeviceable;
    // Request headers
    Headers map[string]string;
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler;
}
// BulkReprovisionCloudPc the bulkReprovisionCloudPc property
func (m *ComanagedDevicesRequestBuilder) BulkReprovisionCloudPc()(*i7fe672ce2ded63745de2480f0a175b485195547bff2ac413663b22ef3172bb4e.BulkReprovisionCloudPcRequestBuilder) {
    return i7fe672ce2ded63745de2480f0a175b485195547bff2ac413663b22ef3172bb4e.NewBulkReprovisionCloudPcRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// BulkRestoreCloudPc the bulkRestoreCloudPc property
func (m *ComanagedDevicesRequestBuilder) BulkRestoreCloudPc()(*i97970601371fcbb5ba7ab85926c15aaf372bd11b13e0e69958d51b9047db72f5.BulkRestoreCloudPcRequestBuilder) {
    return i97970601371fcbb5ba7ab85926c15aaf372bd11b13e0e69958d51b9047db72f5.NewBulkRestoreCloudPcRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// NewComanagedDevicesRequestBuilderInternal instantiates a new ComanagedDevicesRequestBuilder and sets the default values.
func NewComanagedDevicesRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ComanagedDevicesRequestBuilder) {
    m := &ComanagedDevicesRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/deviceManagement/comanagedDevices{?top,skip,search,filter,count,orderby,select,expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewComanagedDevicesRequestBuilder instantiates a new ComanagedDevicesRequestBuilder and sets the default values.
func NewComanagedDevicesRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ComanagedDevicesRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewComanagedDevicesRequestBuilderInternal(urlParams, requestAdapter)
}
// Count the count property
func (m *ComanagedDevicesRequestBuilder) Count()(*i7e96e1b990aeb4d47dec870ab144f92eef626fd02edc82312e6ebb7fb40dbd1b.CountRequestBuilder) {
    return i7e96e1b990aeb4d47dec870ab144f92eef626fd02edc82312e6ebb7fb40dbd1b.NewCountRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// CreateGetRequestInformation the list of co-managed devices report
func (m *ComanagedDevicesRequestBuilder) CreateGetRequestInformation(options *ComanagedDevicesRequestBuilderGetOptions)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET
    if options != nil && options.QueryParameters != nil {
        requestInfo.AddQueryParameters(*(options.QueryParameters))
    }
    if options != nil && options.Headers != nil {
        requestInfo.Headers = options.Headers
    }
    if options != nil && len(options.Options) != 0 {
        err := requestInfo.AddRequestOptions(options.Options...)
        if err != nil {
            return nil, err
        }
    }
    return requestInfo, nil
}
// CreatePostRequestInformation create new navigation property to comanagedDevices for deviceManagement
func (m *ComanagedDevicesRequestBuilder) CreatePostRequestInformation(options *ComanagedDevicesRequestBuilderPostOptions)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST
    requestInfo.SetContentFromParsable(m.requestAdapter, "application/json", options.Body)
    if options != nil && options.Headers != nil {
        requestInfo.Headers = options.Headers
    }
    if options != nil && len(options.Options) != 0 {
        err := requestInfo.AddRequestOptions(options.Options...)
        if err != nil {
            return nil, err
        }
    }
    return requestInfo, nil
}
// ExecuteAction the executeAction property
func (m *ComanagedDevicesRequestBuilder) ExecuteAction()(*i0952f14bd3d84883aa9509ce20c436cb3a9becc286ff26416d1a9ab4f9a5188e.ExecuteActionRequestBuilder) {
    return i0952f14bd3d84883aa9509ce20c436cb3a9becc286ff26416d1a9ab4f9a5188e.NewExecuteActionRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Get the list of co-managed devices report
func (m *ComanagedDevicesRequestBuilder) Get(options *ComanagedDevicesRequestBuilderGetOptions)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ManagedDeviceCollectionResponseable, error) {
    requestInfo, err := m.CreateGetRequestInformation(options);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateManagedDeviceCollectionResponseFromDiscriminatorValue, nil, errorMapping)
    if err != nil {
        return nil, err
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ManagedDeviceCollectionResponseable), nil
}
// MoveDevicesToOU the moveDevicesToOU property
func (m *ComanagedDevicesRequestBuilder) MoveDevicesToOU()(*i21b46d2a36281265e82fedbdddef1c20e278c5a4bbf3cf85fa0d860eccee5b55.MoveDevicesToOURequestBuilder) {
    return i21b46d2a36281265e82fedbdddef1c20e278c5a4bbf3cf85fa0d860eccee5b55.NewMoveDevicesToOURequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Post create new navigation property to comanagedDevices for deviceManagement
func (m *ComanagedDevicesRequestBuilder) Post(options *ComanagedDevicesRequestBuilderPostOptions)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ManagedDeviceable, error) {
    requestInfo, err := m.CreatePostRequestInformation(options);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateManagedDeviceFromDiscriminatorValue, nil, errorMapping)
    if err != nil {
        return nil, err
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ManagedDeviceable), nil
}
