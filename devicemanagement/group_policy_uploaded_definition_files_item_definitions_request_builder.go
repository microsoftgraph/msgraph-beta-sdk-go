package devicemanagement

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// GroupPolicyUploadedDefinitionFilesItemDefinitionsRequestBuilder provides operations to manage the definitions property of the microsoft.graph.groupPolicyDefinitionFile entity.
type GroupPolicyUploadedDefinitionFilesItemDefinitionsRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// GroupPolicyUploadedDefinitionFilesItemDefinitionsRequestBuilderGetQueryParameters the group policy definitions associated with the file.
type GroupPolicyUploadedDefinitionFilesItemDefinitionsRequestBuilderGetQueryParameters struct {
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
// GroupPolicyUploadedDefinitionFilesItemDefinitionsRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type GroupPolicyUploadedDefinitionFilesItemDefinitionsRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *GroupPolicyUploadedDefinitionFilesItemDefinitionsRequestBuilderGetQueryParameters
}
// ByGroupPolicyDefinitionId provides operations to manage the definitions property of the microsoft.graph.groupPolicyDefinitionFile entity.
// returns a *GroupPolicyUploadedDefinitionFilesItemDefinitionsGroupPolicyDefinitionItemRequestBuilder when successful
func (m *GroupPolicyUploadedDefinitionFilesItemDefinitionsRequestBuilder) ByGroupPolicyDefinitionId(groupPolicyDefinitionId string)(*GroupPolicyUploadedDefinitionFilesItemDefinitionsGroupPolicyDefinitionItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    if groupPolicyDefinitionId != "" {
        urlTplParams["groupPolicyDefinition%2Did"] = groupPolicyDefinitionId
    }
    return NewGroupPolicyUploadedDefinitionFilesItemDefinitionsGroupPolicyDefinitionItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// NewGroupPolicyUploadedDefinitionFilesItemDefinitionsRequestBuilderInternal instantiates a new GroupPolicyUploadedDefinitionFilesItemDefinitionsRequestBuilder and sets the default values.
func NewGroupPolicyUploadedDefinitionFilesItemDefinitionsRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*GroupPolicyUploadedDefinitionFilesItemDefinitionsRequestBuilder) {
    m := &GroupPolicyUploadedDefinitionFilesItemDefinitionsRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/deviceManagement/groupPolicyUploadedDefinitionFiles/{groupPolicyUploadedDefinitionFile%2Did}/definitions{?%24count,%24expand,%24filter,%24orderby,%24search,%24select,%24skip,%24top}", pathParameters),
    }
    return m
}
// NewGroupPolicyUploadedDefinitionFilesItemDefinitionsRequestBuilder instantiates a new GroupPolicyUploadedDefinitionFilesItemDefinitionsRequestBuilder and sets the default values.
func NewGroupPolicyUploadedDefinitionFilesItemDefinitionsRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*GroupPolicyUploadedDefinitionFilesItemDefinitionsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewGroupPolicyUploadedDefinitionFilesItemDefinitionsRequestBuilderInternal(urlParams, requestAdapter)
}
// Count provides operations to count the resources in the collection.
// returns a *GroupPolicyUploadedDefinitionFilesItemDefinitionsCountRequestBuilder when successful
func (m *GroupPolicyUploadedDefinitionFilesItemDefinitionsRequestBuilder) Count()(*GroupPolicyUploadedDefinitionFilesItemDefinitionsCountRequestBuilder) {
    return NewGroupPolicyUploadedDefinitionFilesItemDefinitionsCountRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Get the group policy definitions associated with the file.
// returns a GroupPolicyDefinitionCollectionResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *GroupPolicyUploadedDefinitionFilesItemDefinitionsRequestBuilder) Get(ctx context.Context, requestConfiguration *GroupPolicyUploadedDefinitionFilesItemDefinitionsRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.GroupPolicyDefinitionCollectionResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateGroupPolicyDefinitionCollectionResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.GroupPolicyDefinitionCollectionResponseable), nil
}
// ToGetRequestInformation the group policy definitions associated with the file.
// returns a *RequestInformation when successful
func (m *GroupPolicyUploadedDefinitionFilesItemDefinitionsRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *GroupPolicyUploadedDefinitionFilesItemDefinitionsRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *GroupPolicyUploadedDefinitionFilesItemDefinitionsRequestBuilder when successful
func (m *GroupPolicyUploadedDefinitionFilesItemDefinitionsRequestBuilder) WithUrl(rawUrl string)(*GroupPolicyUploadedDefinitionFilesItemDefinitionsRequestBuilder) {
    return NewGroupPolicyUploadedDefinitionFilesItemDefinitionsRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
