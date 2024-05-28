package devicemanagement

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// GrouppolicyuploadeddefinitionfilesGroupPolicyUploadedDefinitionFileItemRequestBuilder provides operations to manage the groupPolicyUploadedDefinitionFiles property of the microsoft.graph.deviceManagement entity.
type GrouppolicyuploadeddefinitionfilesGroupPolicyUploadedDefinitionFileItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// GrouppolicyuploadeddefinitionfilesGroupPolicyUploadedDefinitionFileItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type GrouppolicyuploadeddefinitionfilesGroupPolicyUploadedDefinitionFileItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// GrouppolicyuploadeddefinitionfilesGroupPolicyUploadedDefinitionFileItemRequestBuilderGetQueryParameters the available group policy uploaded definition files for this account.
type GrouppolicyuploadeddefinitionfilesGroupPolicyUploadedDefinitionFileItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// GrouppolicyuploadeddefinitionfilesGroupPolicyUploadedDefinitionFileItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type GrouppolicyuploadeddefinitionfilesGroupPolicyUploadedDefinitionFileItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *GrouppolicyuploadeddefinitionfilesGroupPolicyUploadedDefinitionFileItemRequestBuilderGetQueryParameters
}
// GrouppolicyuploadeddefinitionfilesGroupPolicyUploadedDefinitionFileItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type GrouppolicyuploadeddefinitionfilesGroupPolicyUploadedDefinitionFileItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// AddLanguageFiles provides operations to call the addLanguageFiles method.
// returns a *GrouppolicyuploadeddefinitionfilesItemAddlanguagefilesAddLanguageFilesRequestBuilder when successful
func (m *GrouppolicyuploadeddefinitionfilesGroupPolicyUploadedDefinitionFileItemRequestBuilder) AddLanguageFiles()(*GrouppolicyuploadeddefinitionfilesItemAddlanguagefilesAddLanguageFilesRequestBuilder) {
    return NewGrouppolicyuploadeddefinitionfilesItemAddlanguagefilesAddLanguageFilesRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// NewGrouppolicyuploadeddefinitionfilesGroupPolicyUploadedDefinitionFileItemRequestBuilderInternal instantiates a new GrouppolicyuploadeddefinitionfilesGroupPolicyUploadedDefinitionFileItemRequestBuilder and sets the default values.
func NewGrouppolicyuploadeddefinitionfilesGroupPolicyUploadedDefinitionFileItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*GrouppolicyuploadeddefinitionfilesGroupPolicyUploadedDefinitionFileItemRequestBuilder) {
    m := &GrouppolicyuploadeddefinitionfilesGroupPolicyUploadedDefinitionFileItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/deviceManagement/groupPolicyUploadedDefinitionFiles/{groupPolicyUploadedDefinitionFile%2Did}{?%24expand,%24select}", pathParameters),
    }
    return m
}
// NewGrouppolicyuploadeddefinitionfilesGroupPolicyUploadedDefinitionFileItemRequestBuilder instantiates a new GrouppolicyuploadeddefinitionfilesGroupPolicyUploadedDefinitionFileItemRequestBuilder and sets the default values.
func NewGrouppolicyuploadeddefinitionfilesGroupPolicyUploadedDefinitionFileItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*GrouppolicyuploadeddefinitionfilesGroupPolicyUploadedDefinitionFileItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewGrouppolicyuploadeddefinitionfilesGroupPolicyUploadedDefinitionFileItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Definitions provides operations to manage the definitions property of the microsoft.graph.groupPolicyDefinitionFile entity.
// returns a *GrouppolicyuploadeddefinitionfilesItemDefinitionsRequestBuilder when successful
func (m *GrouppolicyuploadeddefinitionfilesGroupPolicyUploadedDefinitionFileItemRequestBuilder) Definitions()(*GrouppolicyuploadeddefinitionfilesItemDefinitionsRequestBuilder) {
    return NewGrouppolicyuploadeddefinitionfilesItemDefinitionsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Delete delete navigation property groupPolicyUploadedDefinitionFiles for deviceManagement
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *GrouppolicyuploadeddefinitionfilesGroupPolicyUploadedDefinitionFileItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *GrouppolicyuploadeddefinitionfilesGroupPolicyUploadedDefinitionFileItemRequestBuilderDeleteRequestConfiguration)(error) {
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
// Get the available group policy uploaded definition files for this account.
// returns a GroupPolicyUploadedDefinitionFileable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *GrouppolicyuploadeddefinitionfilesGroupPolicyUploadedDefinitionFileItemRequestBuilder) Get(ctx context.Context, requestConfiguration *GrouppolicyuploadeddefinitionfilesGroupPolicyUploadedDefinitionFileItemRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.GroupPolicyUploadedDefinitionFileable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
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
// GroupPolicyOperations provides operations to manage the groupPolicyOperations property of the microsoft.graph.groupPolicyUploadedDefinitionFile entity.
// returns a *GrouppolicyuploadeddefinitionfilesItemGrouppolicyoperationsGroupPolicyOperationsRequestBuilder when successful
func (m *GrouppolicyuploadeddefinitionfilesGroupPolicyUploadedDefinitionFileItemRequestBuilder) GroupPolicyOperations()(*GrouppolicyuploadeddefinitionfilesItemGrouppolicyoperationsGroupPolicyOperationsRequestBuilder) {
    return NewGrouppolicyuploadeddefinitionfilesItemGrouppolicyoperationsGroupPolicyOperationsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Patch update the navigation property groupPolicyUploadedDefinitionFiles in deviceManagement
// returns a GroupPolicyUploadedDefinitionFileable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *GrouppolicyuploadeddefinitionfilesGroupPolicyUploadedDefinitionFileItemRequestBuilder) Patch(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.GroupPolicyUploadedDefinitionFileable, requestConfiguration *GrouppolicyuploadeddefinitionfilesGroupPolicyUploadedDefinitionFileItemRequestBuilderPatchRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.GroupPolicyUploadedDefinitionFileable, error) {
    requestInfo, err := m.ToPatchRequestInformation(ctx, body, requestConfiguration);
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
// Remove provides operations to call the remove method.
// returns a *GrouppolicyuploadeddefinitionfilesItemRemoveRequestBuilder when successful
func (m *GrouppolicyuploadeddefinitionfilesGroupPolicyUploadedDefinitionFileItemRequestBuilder) Remove()(*GrouppolicyuploadeddefinitionfilesItemRemoveRequestBuilder) {
    return NewGrouppolicyuploadeddefinitionfilesItemRemoveRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// RemoveLanguageFiles provides operations to call the removeLanguageFiles method.
// returns a *GrouppolicyuploadeddefinitionfilesItemRemovelanguagefilesRemoveLanguageFilesRequestBuilder when successful
func (m *GrouppolicyuploadeddefinitionfilesGroupPolicyUploadedDefinitionFileItemRequestBuilder) RemoveLanguageFiles()(*GrouppolicyuploadeddefinitionfilesItemRemovelanguagefilesRemoveLanguageFilesRequestBuilder) {
    return NewGrouppolicyuploadeddefinitionfilesItemRemovelanguagefilesRemoveLanguageFilesRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ToDeleteRequestInformation delete navigation property groupPolicyUploadedDefinitionFiles for deviceManagement
// returns a *RequestInformation when successful
func (m *GrouppolicyuploadeddefinitionfilesGroupPolicyUploadedDefinitionFileItemRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *GrouppolicyuploadeddefinitionfilesGroupPolicyUploadedDefinitionFileItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// ToGetRequestInformation the available group policy uploaded definition files for this account.
// returns a *RequestInformation when successful
func (m *GrouppolicyuploadeddefinitionfilesGroupPolicyUploadedDefinitionFileItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *GrouppolicyuploadeddefinitionfilesGroupPolicyUploadedDefinitionFileItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPatchRequestInformation update the navigation property groupPolicyUploadedDefinitionFiles in deviceManagement
// returns a *RequestInformation when successful
func (m *GrouppolicyuploadeddefinitionfilesGroupPolicyUploadedDefinitionFileItemRequestBuilder) ToPatchRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.GroupPolicyUploadedDefinitionFileable, requestConfiguration *GrouppolicyuploadeddefinitionfilesGroupPolicyUploadedDefinitionFileItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// UpdateLanguageFiles provides operations to call the updateLanguageFiles method.
// returns a *GrouppolicyuploadeddefinitionfilesItemUpdatelanguagefilesUpdateLanguageFilesRequestBuilder when successful
func (m *GrouppolicyuploadeddefinitionfilesGroupPolicyUploadedDefinitionFileItemRequestBuilder) UpdateLanguageFiles()(*GrouppolicyuploadeddefinitionfilesItemUpdatelanguagefilesUpdateLanguageFilesRequestBuilder) {
    return NewGrouppolicyuploadeddefinitionfilesItemUpdatelanguagefilesUpdateLanguageFilesRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// UploadNewVersion provides operations to call the uploadNewVersion method.
// returns a *GrouppolicyuploadeddefinitionfilesItemUploadnewversionUploadNewVersionRequestBuilder when successful
func (m *GrouppolicyuploadeddefinitionfilesGroupPolicyUploadedDefinitionFileItemRequestBuilder) UploadNewVersion()(*GrouppolicyuploadeddefinitionfilesItemUploadnewversionUploadNewVersionRequestBuilder) {
    return NewGrouppolicyuploadeddefinitionfilesItemUploadnewversionUploadNewVersionRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *GrouppolicyuploadeddefinitionfilesGroupPolicyUploadedDefinitionFileItemRequestBuilder when successful
func (m *GrouppolicyuploadeddefinitionfilesGroupPolicyUploadedDefinitionFileItemRequestBuilder) WithUrl(rawUrl string)(*GrouppolicyuploadeddefinitionfilesGroupPolicyUploadedDefinitionFileItemRequestBuilder) {
    return NewGrouppolicyuploadeddefinitionfilesGroupPolicyUploadedDefinitionFileItemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
