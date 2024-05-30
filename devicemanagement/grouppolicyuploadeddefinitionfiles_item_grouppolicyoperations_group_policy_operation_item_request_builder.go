package devicemanagement

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// GrouppolicyuploadeddefinitionfilesItemGrouppolicyoperationsGroupPolicyOperationItemRequestBuilder provides operations to manage the groupPolicyOperations property of the microsoft.graph.groupPolicyUploadedDefinitionFile entity.
type GrouppolicyuploadeddefinitionfilesItemGrouppolicyoperationsGroupPolicyOperationItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// GrouppolicyuploadeddefinitionfilesItemGrouppolicyoperationsGroupPolicyOperationItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type GrouppolicyuploadeddefinitionfilesItemGrouppolicyoperationsGroupPolicyOperationItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// GrouppolicyuploadeddefinitionfilesItemGrouppolicyoperationsGroupPolicyOperationItemRequestBuilderGetQueryParameters the list of operations on the uploaded ADMX file.
type GrouppolicyuploadeddefinitionfilesItemGrouppolicyoperationsGroupPolicyOperationItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// GrouppolicyuploadeddefinitionfilesItemGrouppolicyoperationsGroupPolicyOperationItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type GrouppolicyuploadeddefinitionfilesItemGrouppolicyoperationsGroupPolicyOperationItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *GrouppolicyuploadeddefinitionfilesItemGrouppolicyoperationsGroupPolicyOperationItemRequestBuilderGetQueryParameters
}
// GrouppolicyuploadeddefinitionfilesItemGrouppolicyoperationsGroupPolicyOperationItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type GrouppolicyuploadeddefinitionfilesItemGrouppolicyoperationsGroupPolicyOperationItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewGrouppolicyuploadeddefinitionfilesItemGrouppolicyoperationsGroupPolicyOperationItemRequestBuilderInternal instantiates a new GrouppolicyuploadeddefinitionfilesItemGrouppolicyoperationsGroupPolicyOperationItemRequestBuilder and sets the default values.
func NewGrouppolicyuploadeddefinitionfilesItemGrouppolicyoperationsGroupPolicyOperationItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*GrouppolicyuploadeddefinitionfilesItemGrouppolicyoperationsGroupPolicyOperationItemRequestBuilder) {
    m := &GrouppolicyuploadeddefinitionfilesItemGrouppolicyoperationsGroupPolicyOperationItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/deviceManagement/groupPolicyUploadedDefinitionFiles/{groupPolicyUploadedDefinitionFile%2Did}/groupPolicyOperations/{groupPolicyOperation%2Did}{?%24expand,%24select}", pathParameters),
    }
    return m
}
// NewGrouppolicyuploadeddefinitionfilesItemGrouppolicyoperationsGroupPolicyOperationItemRequestBuilder instantiates a new GrouppolicyuploadeddefinitionfilesItemGrouppolicyoperationsGroupPolicyOperationItemRequestBuilder and sets the default values.
func NewGrouppolicyuploadeddefinitionfilesItemGrouppolicyoperationsGroupPolicyOperationItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*GrouppolicyuploadeddefinitionfilesItemGrouppolicyoperationsGroupPolicyOperationItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewGrouppolicyuploadeddefinitionfilesItemGrouppolicyoperationsGroupPolicyOperationItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Delete delete navigation property groupPolicyOperations for deviceManagement
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *GrouppolicyuploadeddefinitionfilesItemGrouppolicyoperationsGroupPolicyOperationItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *GrouppolicyuploadeddefinitionfilesItemGrouppolicyoperationsGroupPolicyOperationItemRequestBuilderDeleteRequestConfiguration)(error) {
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
func (m *GrouppolicyuploadeddefinitionfilesItemGrouppolicyoperationsGroupPolicyOperationItemRequestBuilder) Get(ctx context.Context, requestConfiguration *GrouppolicyuploadeddefinitionfilesItemGrouppolicyoperationsGroupPolicyOperationItemRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.GroupPolicyOperationable, error) {
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
func (m *GrouppolicyuploadeddefinitionfilesItemGrouppolicyoperationsGroupPolicyOperationItemRequestBuilder) Patch(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.GroupPolicyOperationable, requestConfiguration *GrouppolicyuploadeddefinitionfilesItemGrouppolicyoperationsGroupPolicyOperationItemRequestBuilderPatchRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.GroupPolicyOperationable, error) {
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
func (m *GrouppolicyuploadeddefinitionfilesItemGrouppolicyoperationsGroupPolicyOperationItemRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *GrouppolicyuploadeddefinitionfilesItemGrouppolicyoperationsGroupPolicyOperationItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
func (m *GrouppolicyuploadeddefinitionfilesItemGrouppolicyoperationsGroupPolicyOperationItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *GrouppolicyuploadeddefinitionfilesItemGrouppolicyoperationsGroupPolicyOperationItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
func (m *GrouppolicyuploadeddefinitionfilesItemGrouppolicyoperationsGroupPolicyOperationItemRequestBuilder) ToPatchRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.GroupPolicyOperationable, requestConfiguration *GrouppolicyuploadeddefinitionfilesItemGrouppolicyoperationsGroupPolicyOperationItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *GrouppolicyuploadeddefinitionfilesItemGrouppolicyoperationsGroupPolicyOperationItemRequestBuilder when successful
func (m *GrouppolicyuploadeddefinitionfilesItemGrouppolicyoperationsGroupPolicyOperationItemRequestBuilder) WithUrl(rawUrl string)(*GrouppolicyuploadeddefinitionfilesItemGrouppolicyoperationsGroupPolicyOperationItemRequestBuilder) {
    return NewGrouppolicyuploadeddefinitionfilesItemGrouppolicyoperationsGroupPolicyOperationItemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
