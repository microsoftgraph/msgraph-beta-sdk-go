package devicemanagement

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// GroupPolicyUploadedDefinitionFilesGroupPolicyUploadedDefinitionFileItemRequestBuilder provides operations to manage the groupPolicyUploadedDefinitionFiles property of the microsoft.graph.deviceManagement entity.
type GroupPolicyUploadedDefinitionFilesGroupPolicyUploadedDefinitionFileItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// GroupPolicyUploadedDefinitionFilesGroupPolicyUploadedDefinitionFileItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type GroupPolicyUploadedDefinitionFilesGroupPolicyUploadedDefinitionFileItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// GroupPolicyUploadedDefinitionFilesGroupPolicyUploadedDefinitionFileItemRequestBuilderGetQueryParameters the available group policy uploaded definition files for this account.
type GroupPolicyUploadedDefinitionFilesGroupPolicyUploadedDefinitionFileItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// GroupPolicyUploadedDefinitionFilesGroupPolicyUploadedDefinitionFileItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type GroupPolicyUploadedDefinitionFilesGroupPolicyUploadedDefinitionFileItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *GroupPolicyUploadedDefinitionFilesGroupPolicyUploadedDefinitionFileItemRequestBuilderGetQueryParameters
}
// GroupPolicyUploadedDefinitionFilesGroupPolicyUploadedDefinitionFileItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type GroupPolicyUploadedDefinitionFilesGroupPolicyUploadedDefinitionFileItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// AddLanguageFiles provides operations to call the addLanguageFiles method.
func (m *GroupPolicyUploadedDefinitionFilesGroupPolicyUploadedDefinitionFileItemRequestBuilder) AddLanguageFiles()(*GroupPolicyUploadedDefinitionFilesItemAddLanguageFilesRequestBuilder) {
    return NewGroupPolicyUploadedDefinitionFilesItemAddLanguageFilesRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// NewGroupPolicyUploadedDefinitionFilesGroupPolicyUploadedDefinitionFileItemRequestBuilderInternal instantiates a new GroupPolicyUploadedDefinitionFileItemRequestBuilder and sets the default values.
func NewGroupPolicyUploadedDefinitionFilesGroupPolicyUploadedDefinitionFileItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*GroupPolicyUploadedDefinitionFilesGroupPolicyUploadedDefinitionFileItemRequestBuilder) {
    m := &GroupPolicyUploadedDefinitionFilesGroupPolicyUploadedDefinitionFileItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/deviceManagement/groupPolicyUploadedDefinitionFiles/{groupPolicyUploadedDefinitionFile%2Did}{?%24select,%24expand}", pathParameters),
    }
    return m
}
// NewGroupPolicyUploadedDefinitionFilesGroupPolicyUploadedDefinitionFileItemRequestBuilder instantiates a new GroupPolicyUploadedDefinitionFileItemRequestBuilder and sets the default values.
func NewGroupPolicyUploadedDefinitionFilesGroupPolicyUploadedDefinitionFileItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*GroupPolicyUploadedDefinitionFilesGroupPolicyUploadedDefinitionFileItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewGroupPolicyUploadedDefinitionFilesGroupPolicyUploadedDefinitionFileItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Delete delete navigation property groupPolicyUploadedDefinitionFiles for deviceManagement
func (m *GroupPolicyUploadedDefinitionFilesGroupPolicyUploadedDefinitionFileItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *GroupPolicyUploadedDefinitionFilesGroupPolicyUploadedDefinitionFileItemRequestBuilderDeleteRequestConfiguration)([]byte, error) {
    requestInfo, err := m.ToDeleteRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.SendPrimitive(ctx, requestInfo, "[]byte", errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.([]byte), nil
}
// Get the available group policy uploaded definition files for this account.
func (m *GroupPolicyUploadedDefinitionFilesGroupPolicyUploadedDefinitionFileItemRequestBuilder) Get(ctx context.Context, requestConfiguration *GroupPolicyUploadedDefinitionFilesGroupPolicyUploadedDefinitionFileItemRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.GroupPolicyUploadedDefinitionFileable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
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
// GroupPolicyOperations provides operations to manage the groupPolicyOperations property of the microsoft.graph.groupPolicyUploadedDefinitionFile entity.
func (m *GroupPolicyUploadedDefinitionFilesGroupPolicyUploadedDefinitionFileItemRequestBuilder) GroupPolicyOperations()(*GroupPolicyUploadedDefinitionFilesItemGroupPolicyOperationsRequestBuilder) {
    return NewGroupPolicyUploadedDefinitionFilesItemGroupPolicyOperationsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// GroupPolicyOperationsById provides operations to manage the groupPolicyOperations property of the microsoft.graph.groupPolicyUploadedDefinitionFile entity.
func (m *GroupPolicyUploadedDefinitionFilesGroupPolicyUploadedDefinitionFileItemRequestBuilder) GroupPolicyOperationsById(id string)(*GroupPolicyUploadedDefinitionFilesItemGroupPolicyOperationsGroupPolicyOperationItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["groupPolicyOperation%2Did"] = id
    }
    return NewGroupPolicyUploadedDefinitionFilesItemGroupPolicyOperationsGroupPolicyOperationItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// Patch update the navigation property groupPolicyUploadedDefinitionFiles in deviceManagement
func (m *GroupPolicyUploadedDefinitionFilesGroupPolicyUploadedDefinitionFileItemRequestBuilder) Patch(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.GroupPolicyUploadedDefinitionFileable, requestConfiguration *GroupPolicyUploadedDefinitionFilesGroupPolicyUploadedDefinitionFileItemRequestBuilderPatchRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.GroupPolicyUploadedDefinitionFileable, error) {
    requestInfo, err := m.ToPatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
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
// Remove provides operations to call the remove method.
func (m *GroupPolicyUploadedDefinitionFilesGroupPolicyUploadedDefinitionFileItemRequestBuilder) Remove()(*GroupPolicyUploadedDefinitionFilesItemRemoveRequestBuilder) {
    return NewGroupPolicyUploadedDefinitionFilesItemRemoveRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// RemoveLanguageFiles provides operations to call the removeLanguageFiles method.
func (m *GroupPolicyUploadedDefinitionFilesGroupPolicyUploadedDefinitionFileItemRequestBuilder) RemoveLanguageFiles()(*GroupPolicyUploadedDefinitionFilesItemRemoveLanguageFilesRequestBuilder) {
    return NewGroupPolicyUploadedDefinitionFilesItemRemoveLanguageFilesRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ToDeleteRequestInformation delete navigation property groupPolicyUploadedDefinitionFiles for deviceManagement
func (m *GroupPolicyUploadedDefinitionFilesGroupPolicyUploadedDefinitionFileItemRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *GroupPolicyUploadedDefinitionFilesGroupPolicyUploadedDefinitionFileItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.BaseRequestBuilder.UrlTemplate
    requestInfo.PathParameters = m.BaseRequestBuilder.PathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// ToGetRequestInformation the available group policy uploaded definition files for this account.
func (m *GroupPolicyUploadedDefinitionFilesGroupPolicyUploadedDefinitionFileItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *GroupPolicyUploadedDefinitionFilesGroupPolicyUploadedDefinitionFileItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.BaseRequestBuilder.UrlTemplate
    requestInfo.PathParameters = m.BaseRequestBuilder.PathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET
    requestInfo.Headers.Add("Accept", "application/json")
    if requestConfiguration != nil {
        if requestConfiguration.QueryParameters != nil {
            requestInfo.AddQueryParameters(*(requestConfiguration.QueryParameters))
        }
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// ToPatchRequestInformation update the navigation property groupPolicyUploadedDefinitionFiles in deviceManagement
func (m *GroupPolicyUploadedDefinitionFilesGroupPolicyUploadedDefinitionFileItemRequestBuilder) ToPatchRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.GroupPolicyUploadedDefinitionFileable, requestConfiguration *GroupPolicyUploadedDefinitionFilesGroupPolicyUploadedDefinitionFileItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.BaseRequestBuilder.UrlTemplate
    requestInfo.PathParameters = m.BaseRequestBuilder.PathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.PATCH
    requestInfo.Headers.Add("Accept", "application/json")
    err := requestInfo.SetContentFromParsable(ctx, m.BaseRequestBuilder.RequestAdapter, "application/json", body)
    if err != nil {
        return nil, err
    }
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// UpdateLanguageFiles provides operations to call the updateLanguageFiles method.
func (m *GroupPolicyUploadedDefinitionFilesGroupPolicyUploadedDefinitionFileItemRequestBuilder) UpdateLanguageFiles()(*GroupPolicyUploadedDefinitionFilesItemUpdateLanguageFilesRequestBuilder) {
    return NewGroupPolicyUploadedDefinitionFilesItemUpdateLanguageFilesRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// UploadNewVersion provides operations to call the uploadNewVersion method.
func (m *GroupPolicyUploadedDefinitionFilesGroupPolicyUploadedDefinitionFileItemRequestBuilder) UploadNewVersion()(*GroupPolicyUploadedDefinitionFilesItemUploadNewVersionRequestBuilder) {
    return NewGroupPolicyUploadedDefinitionFilesItemUploadNewVersionRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
