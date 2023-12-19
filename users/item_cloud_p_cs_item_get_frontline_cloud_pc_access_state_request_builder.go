package users

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
    i88f3a6af88aa2aea2fd7d6de613104f05bcd862d575941b44d37158780bcdd7b "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/cloudpcs/item/getfrontlinecloudpcaccessstate"
)

// ItemCloudPCsItemGetFrontlineCloudPcAccessStateRequestBuilder provides operations to call the getFrontlineCloudPcAccessState method.
type ItemCloudPCsItemGetFrontlineCloudPcAccessStateRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemCloudPCsItemGetFrontlineCloudPcAccessStateRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemCloudPCsItemGetFrontlineCloudPcAccessStateRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewItemCloudPCsItemGetFrontlineCloudPcAccessStateRequestBuilderInternal instantiates a new GetFrontlineCloudPcAccessStateRequestBuilder and sets the default values.
func NewItemCloudPCsItemGetFrontlineCloudPcAccessStateRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemCloudPCsItemGetFrontlineCloudPcAccessStateRequestBuilder) {
    m := &ItemCloudPCsItemGetFrontlineCloudPcAccessStateRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/users/{user%2Did}/cloudPCs/{cloudPC%2Did}/getFrontlineCloudPcAccessState()", pathParameters),
    }
    return m
}
// NewItemCloudPCsItemGetFrontlineCloudPcAccessStateRequestBuilder instantiates a new GetFrontlineCloudPcAccessStateRequestBuilder and sets the default values.
func NewItemCloudPCsItemGetFrontlineCloudPcAccessStateRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemCloudPCsItemGetFrontlineCloudPcAccessStateRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemCloudPCsItemGetFrontlineCloudPcAccessStateRequestBuilderInternal(urlParams, requestAdapter)
}
// Get invoke function getFrontlineCloudPcAccessState
func (m *ItemCloudPCsItemGetFrontlineCloudPcAccessStateRequestBuilder) Get(ctx context.Context, requestConfiguration *ItemCloudPCsItemGetFrontlineCloudPcAccessStateRequestBuilderGetRequestConfiguration)(*i88f3a6af88aa2aea2fd7d6de613104f05bcd862d575941b44d37158780bcdd7b.GetFrontlineCloudPcAccessStateGetResponse, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.SendEnum(ctx, requestInfo, i88f3a6af88aa2aea2fd7d6de613104f05bcd862d575941b44d37158780bcdd7b.ParseGetFrontlineCloudPcAccessStateGetResponse, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(*i88f3a6af88aa2aea2fd7d6de613104f05bcd862d575941b44d37158780bcdd7b.GetFrontlineCloudPcAccessStateGetResponse), nil
}
// ToGetRequestInformation invoke function getFrontlineCloudPcAccessState
func (m *ItemCloudPCsItemGetFrontlineCloudPcAccessStateRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ItemCloudPCsItemGetFrontlineCloudPcAccessStateRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
func (m *ItemCloudPCsItemGetFrontlineCloudPcAccessStateRequestBuilder) WithUrl(rawUrl string)(*ItemCloudPCsItemGetFrontlineCloudPcAccessStateRequestBuilder) {
    return NewItemCloudPCsItemGetFrontlineCloudPcAccessStateRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
