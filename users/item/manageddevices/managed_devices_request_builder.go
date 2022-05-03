package manageddevices

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
    i221b9cde99e476699da097a8d903ef5a79f3532dc9f4f0e5ee20a66cefd5214f "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/manageddevices/bulkreprovisioncloudpc"
    i354da0a7c77a3e4dbbf4c2908e40eb9ba102f2173af0e90bc4d10956f0e44dee "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/manageddevices/movedevicestoou"
    i8d8a434efe54b7a3fa263217b0c04ff177a4bc1c2230b98a73b80cd1d561c96d "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/manageddevices/executeaction"
    i992082cec9bab654e1839fe55796615f8620ef8a4078776f6e65f83f16c347b1 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/manageddevices/bulkrestorecloudpc"
    ia17acf45d59e7e028076f52b9be902e17efbcf8ed8876140e71381907e497649 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/manageddevices/bulksetcloudpcreviewstatus"
    ic3bd1e20a7b0e931aaa4561b6ef97f202035d17ae0819323a2b55935709cdfa7 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/manageddevices/count"
)

// ManagedDevicesRequestBuilder provides operations to manage the managedDevices property of the microsoft.graph.user entity.
type ManagedDevicesRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// ManagedDevicesRequestBuilderGetQueryParameters the managed devices associated with the user.
type ManagedDevicesRequestBuilderGetQueryParameters struct {
    // Include count of items
    Count *bool `uriparametername:"%24count"`
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Filter items by property values
    Filter *string `uriparametername:"%24filter"`
    // Order items by property values
    Orderby []string `uriparametername:"%24orderby"`
    // Search items by search phrases
    Search *string `uriparametername:"%24search"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
    // Skip the first n items
    Skip *int32 `uriparametername:"%24skip"`
    // Show only the first n items
    Top *int32 `uriparametername:"%24top"`
}
// ManagedDevicesRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ManagedDevicesRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ManagedDevicesRequestBuilderGetQueryParameters
}
// ManagedDevicesRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ManagedDevicesRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// BulkReprovisionCloudPc the bulkReprovisionCloudPc property
func (m *ManagedDevicesRequestBuilder) BulkReprovisionCloudPc()(*i221b9cde99e476699da097a8d903ef5a79f3532dc9f4f0e5ee20a66cefd5214f.BulkReprovisionCloudPcRequestBuilder) {
    return i221b9cde99e476699da097a8d903ef5a79f3532dc9f4f0e5ee20a66cefd5214f.NewBulkReprovisionCloudPcRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// BulkRestoreCloudPc the bulkRestoreCloudPc property
func (m *ManagedDevicesRequestBuilder) BulkRestoreCloudPc()(*i992082cec9bab654e1839fe55796615f8620ef8a4078776f6e65f83f16c347b1.BulkRestoreCloudPcRequestBuilder) {
    return i992082cec9bab654e1839fe55796615f8620ef8a4078776f6e65f83f16c347b1.NewBulkRestoreCloudPcRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// BulkSetCloudPcReviewStatus the bulkSetCloudPcReviewStatus property
func (m *ManagedDevicesRequestBuilder) BulkSetCloudPcReviewStatus()(*ia17acf45d59e7e028076f52b9be902e17efbcf8ed8876140e71381907e497649.BulkSetCloudPcReviewStatusRequestBuilder) {
    return ia17acf45d59e7e028076f52b9be902e17efbcf8ed8876140e71381907e497649.NewBulkSetCloudPcReviewStatusRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// NewManagedDevicesRequestBuilderInternal instantiates a new ManagedDevicesRequestBuilder and sets the default values.
func NewManagedDevicesRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ManagedDevicesRequestBuilder) {
    m := &ManagedDevicesRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/users/{user%2Did}/managedDevices{?%24top,%24skip,%24search,%24filter,%24count,%24orderby,%24select,%24expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewManagedDevicesRequestBuilder instantiates a new ManagedDevicesRequestBuilder and sets the default values.
func NewManagedDevicesRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ManagedDevicesRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewManagedDevicesRequestBuilderInternal(urlParams, requestAdapter)
}
// Count the count property
func (m *ManagedDevicesRequestBuilder) Count()(*ic3bd1e20a7b0e931aaa4561b6ef97f202035d17ae0819323a2b55935709cdfa7.CountRequestBuilder) {
    return ic3bd1e20a7b0e931aaa4561b6ef97f202035d17ae0819323a2b55935709cdfa7.NewCountRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// CreateGetRequestInformationWithRequestConfiguration the managed devices associated with the user.
func (m *ManagedDevicesRequestBuilder) CreateGetRequestInformationWithRequestConfiguration()(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    return m.CreateGetRequestInformationWithRequestConfiguration(nil);
}
// CreateGetRequestInformationWithRequestConfiguration the managed devices associated with the user.
func (m *ManagedDevicesRequestBuilder) CreateGetRequestInformationWithRequestConfiguration(requestConfiguration *ManagedDevicesRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET
    if requestConfiguration != nil {
        if requestConfiguration.QueryParameters != nil {
            requestInfo.AddQueryParameters(*(requestConfiguration.QueryParameters))
        }
        requestInfo.AddRequestHeaders(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// CreatePostRequestInformationWithRequestConfiguration create new navigation property to managedDevices for users
func (m *ManagedDevicesRequestBuilder) CreatePostRequestInformationWithRequestConfiguration(body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ManagedDeviceable)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    return m.CreatePostRequestInformationWithRequestConfiguration(body, nil);
}
// CreatePostRequestInformationWithRequestConfiguration create new navigation property to managedDevices for users
func (m *ManagedDevicesRequestBuilder) CreatePostRequestInformationWithRequestConfiguration(body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ManagedDeviceable, requestConfiguration *ManagedDevicesRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST
    requestInfo.SetContentFromParsable(m.requestAdapter, "application/json", body)
    if requestConfiguration != nil {
        requestInfo.AddRequestHeaders(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// ExecuteAction the executeAction property
func (m *ManagedDevicesRequestBuilder) ExecuteAction()(*i8d8a434efe54b7a3fa263217b0c04ff177a4bc1c2230b98a73b80cd1d561c96d.ExecuteActionRequestBuilder) {
    return i8d8a434efe54b7a3fa263217b0c04ff177a4bc1c2230b98a73b80cd1d561c96d.NewExecuteActionRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// GetWithResponseHandler the managed devices associated with the user.
func (m *ManagedDevicesRequestBuilder) GetWithResponseHandler(requestConfiguration *ManagedDevicesRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ManagedDeviceCollectionResponseable, error) {
    return m.GetWithResponseHandler(requestConfiguration, nil);
}
// GetWithResponseHandler the managed devices associated with the user.
func (m *ManagedDevicesRequestBuilder) GetWithResponseHandler(requestConfiguration *ManagedDevicesRequestBuilderGetRequestConfiguration, responseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ManagedDeviceCollectionResponseable, error) {
    requestInfo, err := m.CreateGetRequestInformationWithRequestConfiguration(requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateManagedDeviceCollectionResponseFromDiscriminatorValue, responseHandler, errorMapping)
    if err != nil {
        return nil, err
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ManagedDeviceCollectionResponseable), nil
}
// MoveDevicesToOU the moveDevicesToOU property
func (m *ManagedDevicesRequestBuilder) MoveDevicesToOU()(*i354da0a7c77a3e4dbbf4c2908e40eb9ba102f2173af0e90bc4d10956f0e44dee.MoveDevicesToOURequestBuilder) {
    return i354da0a7c77a3e4dbbf4c2908e40eb9ba102f2173af0e90bc4d10956f0e44dee.NewMoveDevicesToOURequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// PostWithResponseHandler create new navigation property to managedDevices for users
func (m *ManagedDevicesRequestBuilder) PostWithResponseHandler(body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ManagedDeviceable, requestConfiguration *ManagedDevicesRequestBuilderPostRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ManagedDeviceable, error) {
    return m.PostWithResponseHandler(body, requestConfiguration, nil);
}
// PostWithResponseHandler create new navigation property to managedDevices for users
func (m *ManagedDevicesRequestBuilder) PostWithResponseHandler(body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ManagedDeviceable, requestConfiguration *ManagedDevicesRequestBuilderPostRequestConfiguration, responseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ManagedDeviceable, error) {
    requestInfo, err := m.CreatePostRequestInformationWithRequestConfiguration(body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateManagedDeviceFromDiscriminatorValue, responseHandler, errorMapping)
    if err != nil {
        return nil, err
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ManagedDeviceable), nil
}
