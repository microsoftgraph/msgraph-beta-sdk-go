package users

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// UsersItemManagedDevicesItemGetCloudPcRemoteActionResultsRequestBuilder provides operations to call the getCloudPcRemoteActionResults method.
type UsersItemManagedDevicesItemGetCloudPcRemoteActionResultsRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// UsersItemManagedDevicesItemGetCloudPcRemoteActionResultsRequestBuilderGetQueryParameters invoke function getCloudPcRemoteActionResults
type UsersItemManagedDevicesItemGetCloudPcRemoteActionResultsRequestBuilderGetQueryParameters struct {
    // Include count of items
    Count *bool `uriparametername:"%24count"`
    // Filter items by property values
    Filter *string `uriparametername:"%24filter"`
    // Search items by search phrases
    Search *string `uriparametername:"%24search"`
    // Skip the first n items
    Skip *int32 `uriparametername:"%24skip"`
    // Show only the first n items
    Top *int32 `uriparametername:"%24top"`
}
// UsersItemManagedDevicesItemGetCloudPcRemoteActionResultsRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type UsersItemManagedDevicesItemGetCloudPcRemoteActionResultsRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *UsersItemManagedDevicesItemGetCloudPcRemoteActionResultsRequestBuilderGetQueryParameters
}
// NewUsersItemManagedDevicesItemGetCloudPcRemoteActionResultsRequestBuilderInternal instantiates a new GetCloudPcRemoteActionResultsRequestBuilder and sets the default values.
func NewUsersItemManagedDevicesItemGetCloudPcRemoteActionResultsRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*UsersItemManagedDevicesItemGetCloudPcRemoteActionResultsRequestBuilder) {
    m := &UsersItemManagedDevicesItemGetCloudPcRemoteActionResultsRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/users/{user%2Did}/managedDevices/{managedDevice%2Did}/microsoft.graph.getCloudPcRemoteActionResults(){?%24top,%24skip,%24search,%24filter,%24count}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewUsersItemManagedDevicesItemGetCloudPcRemoteActionResultsRequestBuilder instantiates a new GetCloudPcRemoteActionResultsRequestBuilder and sets the default values.
func NewUsersItemManagedDevicesItemGetCloudPcRemoteActionResultsRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*UsersItemManagedDevicesItemGetCloudPcRemoteActionResultsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewUsersItemManagedDevicesItemGetCloudPcRemoteActionResultsRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateGetRequestInformation invoke function getCloudPcRemoteActionResults
func (m *UsersItemManagedDevicesItemGetCloudPcRemoteActionResultsRequestBuilder) CreateGetRequestInformation(ctx context.Context, requestConfiguration *UsersItemManagedDevicesItemGetCloudPcRemoteActionResultsRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET
    requestInfo.Headers["Accept"] = "application/json"
    if requestConfiguration != nil {
        if requestConfiguration.QueryParameters != nil {
            requestInfo.AddQueryParameters(*(requestConfiguration.QueryParameters))
        }
        requestInfo.AddRequestHeaders(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// Get invoke function getCloudPcRemoteActionResults
func (m *UsersItemManagedDevicesItemGetCloudPcRemoteActionResultsRequestBuilder) Get(ctx context.Context, requestConfiguration *UsersItemManagedDevicesItemGetCloudPcRemoteActionResultsRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.UsersItemManagedDevicesItemGetCloudPcRemoteActionResultsResponseable, error) {
    requestInfo, err := m.CreateGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateUsersItemManagedDevicesItemGetCloudPcRemoteActionResultsResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.UsersItemManagedDevicesItemGetCloudPcRemoteActionResultsResponseable), nil
}
