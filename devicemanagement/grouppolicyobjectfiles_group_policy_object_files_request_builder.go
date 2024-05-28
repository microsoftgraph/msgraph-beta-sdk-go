package devicemanagement

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// GrouppolicyobjectfilesGroupPolicyObjectFilesRequestBuilder provides operations to manage the groupPolicyObjectFiles property of the microsoft.graph.deviceManagement entity.
type GrouppolicyobjectfilesGroupPolicyObjectFilesRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// GrouppolicyobjectfilesGroupPolicyObjectFilesRequestBuilderGetQueryParameters a list of Group Policy Object files uploaded.
type GrouppolicyobjectfilesGroupPolicyObjectFilesRequestBuilderGetQueryParameters struct {
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
// GrouppolicyobjectfilesGroupPolicyObjectFilesRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type GrouppolicyobjectfilesGroupPolicyObjectFilesRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *GrouppolicyobjectfilesGroupPolicyObjectFilesRequestBuilderGetQueryParameters
}
// GrouppolicyobjectfilesGroupPolicyObjectFilesRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type GrouppolicyobjectfilesGroupPolicyObjectFilesRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ByGroupPolicyObjectFileId provides operations to manage the groupPolicyObjectFiles property of the microsoft.graph.deviceManagement entity.
// returns a *GrouppolicyobjectfilesGroupPolicyObjectFileItemRequestBuilder when successful
func (m *GrouppolicyobjectfilesGroupPolicyObjectFilesRequestBuilder) ByGroupPolicyObjectFileId(groupPolicyObjectFileId string)(*GrouppolicyobjectfilesGroupPolicyObjectFileItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    if groupPolicyObjectFileId != "" {
        urlTplParams["groupPolicyObjectFile%2Did"] = groupPolicyObjectFileId
    }
    return NewGrouppolicyobjectfilesGroupPolicyObjectFileItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// NewGrouppolicyobjectfilesGroupPolicyObjectFilesRequestBuilderInternal instantiates a new GrouppolicyobjectfilesGroupPolicyObjectFilesRequestBuilder and sets the default values.
func NewGrouppolicyobjectfilesGroupPolicyObjectFilesRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*GrouppolicyobjectfilesGroupPolicyObjectFilesRequestBuilder) {
    m := &GrouppolicyobjectfilesGroupPolicyObjectFilesRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/deviceManagement/groupPolicyObjectFiles{?%24count,%24expand,%24filter,%24orderby,%24search,%24select,%24skip,%24top}", pathParameters),
    }
    return m
}
// NewGrouppolicyobjectfilesGroupPolicyObjectFilesRequestBuilder instantiates a new GrouppolicyobjectfilesGroupPolicyObjectFilesRequestBuilder and sets the default values.
func NewGrouppolicyobjectfilesGroupPolicyObjectFilesRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*GrouppolicyobjectfilesGroupPolicyObjectFilesRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewGrouppolicyobjectfilesGroupPolicyObjectFilesRequestBuilderInternal(urlParams, requestAdapter)
}
// Count provides operations to count the resources in the collection.
// returns a *GrouppolicyobjectfilesCountRequestBuilder when successful
func (m *GrouppolicyobjectfilesGroupPolicyObjectFilesRequestBuilder) Count()(*GrouppolicyobjectfilesCountRequestBuilder) {
    return NewGrouppolicyobjectfilesCountRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Get a list of Group Policy Object files uploaded.
// returns a GroupPolicyObjectFileCollectionResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *GrouppolicyobjectfilesGroupPolicyObjectFilesRequestBuilder) Get(ctx context.Context, requestConfiguration *GrouppolicyobjectfilesGroupPolicyObjectFilesRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.GroupPolicyObjectFileCollectionResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateGroupPolicyObjectFileCollectionResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.GroupPolicyObjectFileCollectionResponseable), nil
}
// Post create new navigation property to groupPolicyObjectFiles for deviceManagement
// returns a GroupPolicyObjectFileable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *GrouppolicyobjectfilesGroupPolicyObjectFilesRequestBuilder) Post(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.GroupPolicyObjectFileable, requestConfiguration *GrouppolicyobjectfilesGroupPolicyObjectFilesRequestBuilderPostRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.GroupPolicyObjectFileable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateGroupPolicyObjectFileFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.GroupPolicyObjectFileable), nil
}
// ToGetRequestInformation a list of Group Policy Object files uploaded.
// returns a *RequestInformation when successful
func (m *GrouppolicyobjectfilesGroupPolicyObjectFilesRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *GrouppolicyobjectfilesGroupPolicyObjectFilesRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPostRequestInformation create new navigation property to groupPolicyObjectFiles for deviceManagement
// returns a *RequestInformation when successful
func (m *GrouppolicyobjectfilesGroupPolicyObjectFilesRequestBuilder) ToPostRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.GroupPolicyObjectFileable, requestConfiguration *GrouppolicyobjectfilesGroupPolicyObjectFilesRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    err := requestInfo.SetContentFromParsable(ctx, m.BaseRequestBuilder.RequestAdapter, "application/json", body)
    if err != nil {
        return nil, err
    }
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *GrouppolicyobjectfilesGroupPolicyObjectFilesRequestBuilder when successful
func (m *GrouppolicyobjectfilesGroupPolicyObjectFilesRequestBuilder) WithUrl(rawUrl string)(*GrouppolicyobjectfilesGroupPolicyObjectFilesRequestBuilder) {
    return NewGrouppolicyobjectfilesGroupPolicyObjectFilesRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
