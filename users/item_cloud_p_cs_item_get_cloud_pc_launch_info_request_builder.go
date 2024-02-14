package users

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// ItemCloudPCsItemGetCloudPcLaunchInfoRequestBuilder provides operations to call the getCloudPcLaunchInfo method.
type ItemCloudPCsItemGetCloudPcLaunchInfoRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemCloudPCsItemGetCloudPcLaunchInfoRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemCloudPCsItemGetCloudPcLaunchInfoRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewItemCloudPCsItemGetCloudPcLaunchInfoRequestBuilderInternal instantiates a new ItemCloudPCsItemGetCloudPcLaunchInfoRequestBuilder and sets the default values.
func NewItemCloudPCsItemGetCloudPcLaunchInfoRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemCloudPCsItemGetCloudPcLaunchInfoRequestBuilder) {
    m := &ItemCloudPCsItemGetCloudPcLaunchInfoRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/users/{user%2Did}/cloudPCs/{cloudPC%2Did}/getCloudPcLaunchInfo()", pathParameters),
    }
    return m
}
// NewItemCloudPCsItemGetCloudPcLaunchInfoRequestBuilder instantiates a new ItemCloudPCsItemGetCloudPcLaunchInfoRequestBuilder and sets the default values.
func NewItemCloudPCsItemGetCloudPcLaunchInfoRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemCloudPCsItemGetCloudPcLaunchInfoRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemCloudPCsItemGetCloudPcLaunchInfoRequestBuilderInternal(urlParams, requestAdapter)
}
// Get invoke function getCloudPcLaunchInfo
// returns a CloudPcLaunchInfoable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ItemCloudPCsItemGetCloudPcLaunchInfoRequestBuilder) Get(ctx context.Context, requestConfiguration *ItemCloudPCsItemGetCloudPcLaunchInfoRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CloudPcLaunchInfoable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateCloudPcLaunchInfoFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CloudPcLaunchInfoable), nil
}
// ToGetRequestInformation invoke function getCloudPcLaunchInfo
// returns a *RequestInformation when successful
func (m *ItemCloudPCsItemGetCloudPcLaunchInfoRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ItemCloudPCsItemGetCloudPcLaunchInfoRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *ItemCloudPCsItemGetCloudPcLaunchInfoRequestBuilder when successful
func (m *ItemCloudPCsItemGetCloudPcLaunchInfoRequestBuilder) WithUrl(rawUrl string)(*ItemCloudPCsItemGetCloudPcLaunchInfoRequestBuilder) {
    return NewItemCloudPCsItemGetCloudPcLaunchInfoRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
