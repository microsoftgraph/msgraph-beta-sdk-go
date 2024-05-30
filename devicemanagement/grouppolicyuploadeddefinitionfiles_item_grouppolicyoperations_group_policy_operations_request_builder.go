package devicemanagement

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// GrouppolicyuploadeddefinitionfilesItemGrouppolicyoperationsGroupPolicyOperationsRequestBuilder provides operations to manage the groupPolicyOperations property of the microsoft.graph.groupPolicyUploadedDefinitionFile entity.
type GrouppolicyuploadeddefinitionfilesItemGrouppolicyoperationsGroupPolicyOperationsRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// GrouppolicyuploadeddefinitionfilesItemGrouppolicyoperationsGroupPolicyOperationsRequestBuilderGetQueryParameters the list of operations on the uploaded ADMX file.
type GrouppolicyuploadeddefinitionfilesItemGrouppolicyoperationsGroupPolicyOperationsRequestBuilderGetQueryParameters struct {
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
// GrouppolicyuploadeddefinitionfilesItemGrouppolicyoperationsGroupPolicyOperationsRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type GrouppolicyuploadeddefinitionfilesItemGrouppolicyoperationsGroupPolicyOperationsRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *GrouppolicyuploadeddefinitionfilesItemGrouppolicyoperationsGroupPolicyOperationsRequestBuilderGetQueryParameters
}
// GrouppolicyuploadeddefinitionfilesItemGrouppolicyoperationsGroupPolicyOperationsRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type GrouppolicyuploadeddefinitionfilesItemGrouppolicyoperationsGroupPolicyOperationsRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ByGroupPolicyOperationId provides operations to manage the groupPolicyOperations property of the microsoft.graph.groupPolicyUploadedDefinitionFile entity.
// returns a *GrouppolicyuploadeddefinitionfilesItemGrouppolicyoperationsGroupPolicyOperationItemRequestBuilder when successful
func (m *GrouppolicyuploadeddefinitionfilesItemGrouppolicyoperationsGroupPolicyOperationsRequestBuilder) ByGroupPolicyOperationId(groupPolicyOperationId string)(*GrouppolicyuploadeddefinitionfilesItemGrouppolicyoperationsGroupPolicyOperationItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    if groupPolicyOperationId != "" {
        urlTplParams["groupPolicyOperation%2Did"] = groupPolicyOperationId
    }
    return NewGrouppolicyuploadeddefinitionfilesItemGrouppolicyoperationsGroupPolicyOperationItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// NewGrouppolicyuploadeddefinitionfilesItemGrouppolicyoperationsGroupPolicyOperationsRequestBuilderInternal instantiates a new GrouppolicyuploadeddefinitionfilesItemGrouppolicyoperationsGroupPolicyOperationsRequestBuilder and sets the default values.
func NewGrouppolicyuploadeddefinitionfilesItemGrouppolicyoperationsGroupPolicyOperationsRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*GrouppolicyuploadeddefinitionfilesItemGrouppolicyoperationsGroupPolicyOperationsRequestBuilder) {
    m := &GrouppolicyuploadeddefinitionfilesItemGrouppolicyoperationsGroupPolicyOperationsRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/deviceManagement/groupPolicyUploadedDefinitionFiles/{groupPolicyUploadedDefinitionFile%2Did}/groupPolicyOperations{?%24count,%24expand,%24filter,%24orderby,%24search,%24select,%24skip,%24top}", pathParameters),
    }
    return m
}
// NewGrouppolicyuploadeddefinitionfilesItemGrouppolicyoperationsGroupPolicyOperationsRequestBuilder instantiates a new GrouppolicyuploadeddefinitionfilesItemGrouppolicyoperationsGroupPolicyOperationsRequestBuilder and sets the default values.
func NewGrouppolicyuploadeddefinitionfilesItemGrouppolicyoperationsGroupPolicyOperationsRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*GrouppolicyuploadeddefinitionfilesItemGrouppolicyoperationsGroupPolicyOperationsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewGrouppolicyuploadeddefinitionfilesItemGrouppolicyoperationsGroupPolicyOperationsRequestBuilderInternal(urlParams, requestAdapter)
}
// Count provides operations to count the resources in the collection.
// returns a *GrouppolicyuploadeddefinitionfilesItemGrouppolicyoperationsCountRequestBuilder when successful
func (m *GrouppolicyuploadeddefinitionfilesItemGrouppolicyoperationsGroupPolicyOperationsRequestBuilder) Count()(*GrouppolicyuploadeddefinitionfilesItemGrouppolicyoperationsCountRequestBuilder) {
    return NewGrouppolicyuploadeddefinitionfilesItemGrouppolicyoperationsCountRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Get the list of operations on the uploaded ADMX file.
// returns a GroupPolicyOperationCollectionResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *GrouppolicyuploadeddefinitionfilesItemGrouppolicyoperationsGroupPolicyOperationsRequestBuilder) Get(ctx context.Context, requestConfiguration *GrouppolicyuploadeddefinitionfilesItemGrouppolicyoperationsGroupPolicyOperationsRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.GroupPolicyOperationCollectionResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateGroupPolicyOperationCollectionResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.GroupPolicyOperationCollectionResponseable), nil
}
// Post create new navigation property to groupPolicyOperations for deviceManagement
// returns a GroupPolicyOperationable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *GrouppolicyuploadeddefinitionfilesItemGrouppolicyoperationsGroupPolicyOperationsRequestBuilder) Post(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.GroupPolicyOperationable, requestConfiguration *GrouppolicyuploadeddefinitionfilesItemGrouppolicyoperationsGroupPolicyOperationsRequestBuilderPostRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.GroupPolicyOperationable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateGroupPolicyOperationFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.GroupPolicyOperationable), nil
}
// ToGetRequestInformation the list of operations on the uploaded ADMX file.
// returns a *RequestInformation when successful
func (m *GrouppolicyuploadeddefinitionfilesItemGrouppolicyoperationsGroupPolicyOperationsRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *GrouppolicyuploadeddefinitionfilesItemGrouppolicyoperationsGroupPolicyOperationsRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPostRequestInformation create new navigation property to groupPolicyOperations for deviceManagement
// returns a *RequestInformation when successful
func (m *GrouppolicyuploadeddefinitionfilesItemGrouppolicyoperationsGroupPolicyOperationsRequestBuilder) ToPostRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.GroupPolicyOperationable, requestConfiguration *GrouppolicyuploadeddefinitionfilesItemGrouppolicyoperationsGroupPolicyOperationsRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *GrouppolicyuploadeddefinitionfilesItemGrouppolicyoperationsGroupPolicyOperationsRequestBuilder when successful
func (m *GrouppolicyuploadeddefinitionfilesItemGrouppolicyoperationsGroupPolicyOperationsRequestBuilder) WithUrl(rawUrl string)(*GrouppolicyuploadeddefinitionfilesItemGrouppolicyoperationsGroupPolicyOperationsRequestBuilder) {
    return NewGrouppolicyuploadeddefinitionfilesItemGrouppolicyoperationsGroupPolicyOperationsRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
