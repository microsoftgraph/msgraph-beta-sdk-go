package devicemanagement

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// GroupPolicyUploadedDefinitionFilesItemGroupPolicyOperationsGroupPolicyOperationItemRequestBuilder provides operations to manage the groupPolicyOperations property of the microsoft.graph.groupPolicyUploadedDefinitionFile entity.
type GroupPolicyUploadedDefinitionFilesItemGroupPolicyOperationsGroupPolicyOperationItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// GroupPolicyUploadedDefinitionFilesItemGroupPolicyOperationsGroupPolicyOperationItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type GroupPolicyUploadedDefinitionFilesItemGroupPolicyOperationsGroupPolicyOperationItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// GroupPolicyUploadedDefinitionFilesItemGroupPolicyOperationsGroupPolicyOperationItemRequestBuilderGetQueryParameters the list of operations on the uploaded ADMX file.
type GroupPolicyUploadedDefinitionFilesItemGroupPolicyOperationsGroupPolicyOperationItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// GroupPolicyUploadedDefinitionFilesItemGroupPolicyOperationsGroupPolicyOperationItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type GroupPolicyUploadedDefinitionFilesItemGroupPolicyOperationsGroupPolicyOperationItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *GroupPolicyUploadedDefinitionFilesItemGroupPolicyOperationsGroupPolicyOperationItemRequestBuilderGetQueryParameters
}
// GroupPolicyUploadedDefinitionFilesItemGroupPolicyOperationsGroupPolicyOperationItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type GroupPolicyUploadedDefinitionFilesItemGroupPolicyOperationsGroupPolicyOperationItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewGroupPolicyUploadedDefinitionFilesItemGroupPolicyOperationsGroupPolicyOperationItemRequestBuilderInternal instantiates a new GroupPolicyUploadedDefinitionFilesItemGroupPolicyOperationsGroupPolicyOperationItemRequestBuilder and sets the default values.
func NewGroupPolicyUploadedDefinitionFilesItemGroupPolicyOperationsGroupPolicyOperationItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*GroupPolicyUploadedDefinitionFilesItemGroupPolicyOperationsGroupPolicyOperationItemRequestBuilder) {
    m := &GroupPolicyUploadedDefinitionFilesItemGroupPolicyOperationsGroupPolicyOperationItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/deviceManagement/groupPolicyUploadedDefinitionFiles/{groupPolicyUploadedDefinitionFile%2Did}/groupPolicyOperations/{groupPolicyOperation%2Did}{?%24expand,%24select}", pathParameters),
    }
    return m
}
// NewGroupPolicyUploadedDefinitionFilesItemGroupPolicyOperationsGroupPolicyOperationItemRequestBuilder instantiates a new GroupPolicyUploadedDefinitionFilesItemGroupPolicyOperationsGroupPolicyOperationItemRequestBuilder and sets the default values.
func NewGroupPolicyUploadedDefinitionFilesItemGroupPolicyOperationsGroupPolicyOperationItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*GroupPolicyUploadedDefinitionFilesItemGroupPolicyOperationsGroupPolicyOperationItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewGroupPolicyUploadedDefinitionFilesItemGroupPolicyOperationsGroupPolicyOperationItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Delete delete navigation property groupPolicyOperations for deviceManagement
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *GroupPolicyUploadedDefinitionFilesItemGroupPolicyOperationsGroupPolicyOperationItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *GroupPolicyUploadedDefinitionFilesItemGroupPolicyOperationsGroupPolicyOperationItemRequestBuilderDeleteRequestConfiguration)(error) {
    requestInfo, err := m.ToDeleteRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    err = m.BaseRequestBuilder.RequestAdapter.SendNoContent(ctx, requestInfo, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
// Get the list of operations on the uploaded ADMX file.
// returns a GroupPolicyOperationable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *GroupPolicyUploadedDefinitionFilesItemGroupPolicyOperationsGroupPolicyOperationItemRequestBuilder) Get(ctx context.Context, requestConfiguration *GroupPolicyUploadedDefinitionFilesItemGroupPolicyOperationsGroupPolicyOperationItemRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.GroupPolicyOperationable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
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
// Patch update the navigation property groupPolicyOperations in deviceManagement
// returns a GroupPolicyOperationable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *GroupPolicyUploadedDefinitionFilesItemGroupPolicyOperationsGroupPolicyOperationItemRequestBuilder) Patch(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.GroupPolicyOperationable, requestConfiguration *GroupPolicyUploadedDefinitionFilesItemGroupPolicyOperationsGroupPolicyOperationItemRequestBuilderPatchRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.GroupPolicyOperationable, error) {
    requestInfo, err := m.ToPatchRequestInformation(ctx, body, requestConfiguration);
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
// ToDeleteRequestInformation delete navigation property groupPolicyOperations for deviceManagement
// returns a *RequestInformation when successful
func (m *GroupPolicyUploadedDefinitionFilesItemGroupPolicyOperationsGroupPolicyOperationItemRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *GroupPolicyUploadedDefinitionFilesItemGroupPolicyOperationsGroupPolicyOperationItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// ToGetRequestInformation the list of operations on the uploaded ADMX file.
// returns a *RequestInformation when successful
func (m *GroupPolicyUploadedDefinitionFilesItemGroupPolicyOperationsGroupPolicyOperationItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *GroupPolicyUploadedDefinitionFilesItemGroupPolicyOperationsGroupPolicyOperationItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPatchRequestInformation update the navigation property groupPolicyOperations in deviceManagement
// returns a *RequestInformation when successful
func (m *GroupPolicyUploadedDefinitionFilesItemGroupPolicyOperationsGroupPolicyOperationItemRequestBuilder) ToPatchRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.GroupPolicyOperationable, requestConfiguration *GroupPolicyUploadedDefinitionFilesItemGroupPolicyOperationsGroupPolicyOperationItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.PATCH, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
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
// returns a *GroupPolicyUploadedDefinitionFilesItemGroupPolicyOperationsGroupPolicyOperationItemRequestBuilder when successful
func (m *GroupPolicyUploadedDefinitionFilesItemGroupPolicyOperationsGroupPolicyOperationItemRequestBuilder) WithUrl(rawUrl string)(*GroupPolicyUploadedDefinitionFilesItemGroupPolicyOperationsGroupPolicyOperationItemRequestBuilder) {
    return NewGroupPolicyUploadedDefinitionFilesItemGroupPolicyOperationsGroupPolicyOperationItemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
