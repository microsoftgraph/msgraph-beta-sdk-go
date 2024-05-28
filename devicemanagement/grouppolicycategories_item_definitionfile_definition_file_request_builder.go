package devicemanagement

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// GrouppolicycategoriesItemDefinitionfileDefinitionFileRequestBuilder provides operations to manage the definitionFile property of the microsoft.graph.groupPolicyCategory entity.
type GrouppolicycategoriesItemDefinitionfileDefinitionFileRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// GrouppolicycategoriesItemDefinitionfileDefinitionFileRequestBuilderGetQueryParameters the id of the definition file the category came from
type GrouppolicycategoriesItemDefinitionfileDefinitionFileRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// GrouppolicycategoriesItemDefinitionfileDefinitionFileRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type GrouppolicycategoriesItemDefinitionfileDefinitionFileRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *GrouppolicycategoriesItemDefinitionfileDefinitionFileRequestBuilderGetQueryParameters
}
// NewGrouppolicycategoriesItemDefinitionfileDefinitionFileRequestBuilderInternal instantiates a new GrouppolicycategoriesItemDefinitionfileDefinitionFileRequestBuilder and sets the default values.
func NewGrouppolicycategoriesItemDefinitionfileDefinitionFileRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*GrouppolicycategoriesItemDefinitionfileDefinitionFileRequestBuilder) {
    m := &GrouppolicycategoriesItemDefinitionfileDefinitionFileRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/deviceManagement/groupPolicyCategories/{groupPolicyCategory%2Did}/definitionFile{?%24expand,%24select}", pathParameters),
    }
    return m
}
// NewGrouppolicycategoriesItemDefinitionfileDefinitionFileRequestBuilder instantiates a new GrouppolicycategoriesItemDefinitionfileDefinitionFileRequestBuilder and sets the default values.
func NewGrouppolicycategoriesItemDefinitionfileDefinitionFileRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*GrouppolicycategoriesItemDefinitionfileDefinitionFileRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewGrouppolicycategoriesItemDefinitionfileDefinitionFileRequestBuilderInternal(urlParams, requestAdapter)
}
// Get the id of the definition file the category came from
// returns a GroupPolicyDefinitionFileable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *GrouppolicycategoriesItemDefinitionfileDefinitionFileRequestBuilder) Get(ctx context.Context, requestConfiguration *GrouppolicycategoriesItemDefinitionfileDefinitionFileRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.GroupPolicyDefinitionFileable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateGroupPolicyDefinitionFileFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.GroupPolicyDefinitionFileable), nil
}
// ToGetRequestInformation the id of the definition file the category came from
// returns a *RequestInformation when successful
func (m *GrouppolicycategoriesItemDefinitionfileDefinitionFileRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *GrouppolicycategoriesItemDefinitionfileDefinitionFileRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        if requestConfiguration.QueryParameters != nil {
            requestInfo.AddQueryParameters(*(requestConfiguration.QueryParameters))
        }
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *GrouppolicycategoriesItemDefinitionfileDefinitionFileRequestBuilder when successful
func (m *GrouppolicycategoriesItemDefinitionfileDefinitionFileRequestBuilder) WithUrl(rawUrl string)(*GrouppolicycategoriesItemDefinitionfileDefinitionFileRequestBuilder) {
    return NewGrouppolicycategoriesItemDefinitionfileDefinitionFileRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
