package devicemanagement

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// GrouppolicyconfigurationsItemDefinitionvaluesItemPresentationvaluesCountRequestBuilder provides operations to count the resources in the collection.
type GrouppolicyconfigurationsItemDefinitionvaluesItemPresentationvaluesCountRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// GrouppolicyconfigurationsItemDefinitionvaluesItemPresentationvaluesCountRequestBuilderGetQueryParameters get the number of the resource
type GrouppolicyconfigurationsItemDefinitionvaluesItemPresentationvaluesCountRequestBuilderGetQueryParameters struct {
    // Filter items by property values
    Filter *string `uriparametername:"%24filter"`
    // Search items by search phrases
    Search *string `uriparametername:"%24search"`
}
// GrouppolicyconfigurationsItemDefinitionvaluesItemPresentationvaluesCountRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type GrouppolicyconfigurationsItemDefinitionvaluesItemPresentationvaluesCountRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *GrouppolicyconfigurationsItemDefinitionvaluesItemPresentationvaluesCountRequestBuilderGetQueryParameters
}
// NewGrouppolicyconfigurationsItemDefinitionvaluesItemPresentationvaluesCountRequestBuilderInternal instantiates a new GrouppolicyconfigurationsItemDefinitionvaluesItemPresentationvaluesCountRequestBuilder and sets the default values.
func NewGrouppolicyconfigurationsItemDefinitionvaluesItemPresentationvaluesCountRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*GrouppolicyconfigurationsItemDefinitionvaluesItemPresentationvaluesCountRequestBuilder) {
    m := &GrouppolicyconfigurationsItemDefinitionvaluesItemPresentationvaluesCountRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/deviceManagement/groupPolicyConfigurations/{groupPolicyConfiguration%2Did}/definitionValues/{groupPolicyDefinitionValue%2Did}/presentationValues/$count{?%24filter,%24search}", pathParameters),
    }
    return m
}
// NewGrouppolicyconfigurationsItemDefinitionvaluesItemPresentationvaluesCountRequestBuilder instantiates a new GrouppolicyconfigurationsItemDefinitionvaluesItemPresentationvaluesCountRequestBuilder and sets the default values.
func NewGrouppolicyconfigurationsItemDefinitionvaluesItemPresentationvaluesCountRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*GrouppolicyconfigurationsItemDefinitionvaluesItemPresentationvaluesCountRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewGrouppolicyconfigurationsItemDefinitionvaluesItemPresentationvaluesCountRequestBuilderInternal(urlParams, requestAdapter)
}
// Get get the number of the resource
// returns a *int32 when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *GrouppolicyconfigurationsItemDefinitionvaluesItemPresentationvaluesCountRequestBuilder) Get(ctx context.Context, requestConfiguration *GrouppolicyconfigurationsItemDefinitionvaluesItemPresentationvaluesCountRequestBuilderGetRequestConfiguration)(*int32, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.SendPrimitive(ctx, requestInfo, "int32", errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(*int32), nil
}
// ToGetRequestInformation get the number of the resource
// returns a *RequestInformation when successful
func (m *GrouppolicyconfigurationsItemDefinitionvaluesItemPresentationvaluesCountRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *GrouppolicyconfigurationsItemDefinitionvaluesItemPresentationvaluesCountRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        if requestConfiguration.QueryParameters != nil {
            requestInfo.AddQueryParameters(*(requestConfiguration.QueryParameters))
        }
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "text/plain;q=0.9")
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *GrouppolicyconfigurationsItemDefinitionvaluesItemPresentationvaluesCountRequestBuilder when successful
func (m *GrouppolicyconfigurationsItemDefinitionvaluesItemPresentationvaluesCountRequestBuilder) WithUrl(rawUrl string)(*GrouppolicyconfigurationsItemDefinitionvaluesItemPresentationvaluesCountRequestBuilder) {
    return NewGrouppolicyconfigurationsItemDefinitionvaluesItemPresentationvaluesCountRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
