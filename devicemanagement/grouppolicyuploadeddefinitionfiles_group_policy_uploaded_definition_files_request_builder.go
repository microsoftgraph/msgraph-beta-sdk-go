package devicemanagement

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// GrouppolicyuploadeddefinitionfilesGroupPolicyUploadedDefinitionFilesRequestBuilder provides operations to manage the groupPolicyUploadedDefinitionFiles property of the microsoft.graph.deviceManagement entity.
type GrouppolicyuploadeddefinitionfilesGroupPolicyUploadedDefinitionFilesRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// GrouppolicyuploadeddefinitionfilesGroupPolicyUploadedDefinitionFilesRequestBuilderGetQueryParameters the available group policy uploaded definition files for this account.
type GrouppolicyuploadeddefinitionfilesGroupPolicyUploadedDefinitionFilesRequestBuilderGetQueryParameters struct {
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
// GrouppolicyuploadeddefinitionfilesGroupPolicyUploadedDefinitionFilesRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type GrouppolicyuploadeddefinitionfilesGroupPolicyUploadedDefinitionFilesRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *GrouppolicyuploadeddefinitionfilesGroupPolicyUploadedDefinitionFilesRequestBuilderGetQueryParameters
}
// GrouppolicyuploadeddefinitionfilesGroupPolicyUploadedDefinitionFilesRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type GrouppolicyuploadeddefinitionfilesGroupPolicyUploadedDefinitionFilesRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ByGroupPolicyUploadedDefinitionFileId provides operations to manage the groupPolicyUploadedDefinitionFiles property of the microsoft.graph.deviceManagement entity.
// returns a *GrouppolicyuploadeddefinitionfilesGroupPolicyUploadedDefinitionFileItemRequestBuilder when successful
func (m *GrouppolicyuploadeddefinitionfilesGroupPolicyUploadedDefinitionFilesRequestBuilder) ByGroupPolicyUploadedDefinitionFileId(groupPolicyUploadedDefinitionFileId string)(*GrouppolicyuploadeddefinitionfilesGroupPolicyUploadedDefinitionFileItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    if groupPolicyUploadedDefinitionFileId != "" {
        urlTplParams["groupPolicyUploadedDefinitionFile%2Did"] = groupPolicyUploadedDefinitionFileId
    }
    return NewGrouppolicyuploadeddefinitionfilesGroupPolicyUploadedDefinitionFileItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// NewGrouppolicyuploadeddefinitionfilesGroupPolicyUploadedDefinitionFilesRequestBuilderInternal instantiates a new GrouppolicyuploadeddefinitionfilesGroupPolicyUploadedDefinitionFilesRequestBuilder and sets the default values.
func NewGrouppolicyuploadeddefinitionfilesGroupPolicyUploadedDefinitionFilesRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*GrouppolicyuploadeddefinitionfilesGroupPolicyUploadedDefinitionFilesRequestBuilder) {
    m := &GrouppolicyuploadeddefinitionfilesGroupPolicyUploadedDefinitionFilesRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/deviceManagement/groupPolicyUploadedDefinitionFiles{?%24count,%24expand,%24filter,%24orderby,%24search,%24select,%24skip,%24top}", pathParameters),
    }
    return m
}
// NewGrouppolicyuploadeddefinitionfilesGroupPolicyUploadedDefinitionFilesRequestBuilder instantiates a new GrouppolicyuploadeddefinitionfilesGroupPolicyUploadedDefinitionFilesRequestBuilder and sets the default values.
func NewGrouppolicyuploadeddefinitionfilesGroupPolicyUploadedDefinitionFilesRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*GrouppolicyuploadeddefinitionfilesGroupPolicyUploadedDefinitionFilesRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewGrouppolicyuploadeddefinitionfilesGroupPolicyUploadedDefinitionFilesRequestBuilderInternal(urlParams, requestAdapter)
}
// Count provides operations to count the resources in the collection.
// returns a *GrouppolicyuploadeddefinitionfilesCountRequestBuilder when successful
func (m *GrouppolicyuploadeddefinitionfilesGroupPolicyUploadedDefinitionFilesRequestBuilder) Count()(*GrouppolicyuploadeddefinitionfilesCountRequestBuilder) {
    return NewGrouppolicyuploadeddefinitionfilesCountRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Get the available group policy uploaded definition files for this account.
// returns a GroupPolicyUploadedDefinitionFileCollectionResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *GrouppolicyuploadeddefinitionfilesGroupPolicyUploadedDefinitionFilesRequestBuilder) Get(ctx context.Context, requestConfiguration *GrouppolicyuploadeddefinitionfilesGroupPolicyUploadedDefinitionFilesRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.GroupPolicyUploadedDefinitionFileCollectionResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateGroupPolicyUploadedDefinitionFileCollectionResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.GroupPolicyUploadedDefinitionFileCollectionResponseable), nil
}
// Post create new navigation property to groupPolicyUploadedDefinitionFiles for deviceManagement
// returns a GroupPolicyUploadedDefinitionFileable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *GrouppolicyuploadeddefinitionfilesGroupPolicyUploadedDefinitionFilesRequestBuilder) Post(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.GroupPolicyUploadedDefinitionFileable, requestConfiguration *GrouppolicyuploadeddefinitionfilesGroupPolicyUploadedDefinitionFilesRequestBuilderPostRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.GroupPolicyUploadedDefinitionFileable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateGroupPolicyUploadedDefinitionFileFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.GroupPolicyUploadedDefinitionFileable), nil
}
// ToGetRequestInformation the available group policy uploaded definition files for this account.
// returns a *RequestInformation when successful
func (m *GrouppolicyuploadeddefinitionfilesGroupPolicyUploadedDefinitionFilesRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *GrouppolicyuploadeddefinitionfilesGroupPolicyUploadedDefinitionFilesRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPostRequestInformation create new navigation property to groupPolicyUploadedDefinitionFiles for deviceManagement
// returns a *RequestInformation when successful
func (m *GrouppolicyuploadeddefinitionfilesGroupPolicyUploadedDefinitionFilesRequestBuilder) ToPostRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.GroupPolicyUploadedDefinitionFileable, requestConfiguration *GrouppolicyuploadeddefinitionfilesGroupPolicyUploadedDefinitionFilesRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *GrouppolicyuploadeddefinitionfilesGroupPolicyUploadedDefinitionFilesRequestBuilder when successful
func (m *GrouppolicyuploadeddefinitionfilesGroupPolicyUploadedDefinitionFilesRequestBuilder) WithUrl(rawUrl string)(*GrouppolicyuploadeddefinitionfilesGroupPolicyUploadedDefinitionFilesRequestBuilder) {
    return NewGrouppolicyuploadeddefinitionfilesGroupPolicyUploadedDefinitionFilesRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
