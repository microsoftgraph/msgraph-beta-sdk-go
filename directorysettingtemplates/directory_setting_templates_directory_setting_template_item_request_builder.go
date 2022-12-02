package directorysettingtemplates

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// DirectorySettingTemplatesDirectorySettingTemplateItemRequestBuilder provides operations to manage the collection of directorySettingTemplate entities.
type DirectorySettingTemplatesDirectorySettingTemplateItemRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// DirectorySettingTemplatesDirectorySettingTemplateItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type DirectorySettingTemplatesDirectorySettingTemplateItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// DirectorySettingTemplatesDirectorySettingTemplateItemRequestBuilderGetQueryParameters a directory setting template represents a template of settings from which settings may be created within a tenant. This operation allows retrieval of the properties of the **directorySettingTemplate** object, including the available settings and their defaults.
type DirectorySettingTemplatesDirectorySettingTemplateItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// DirectorySettingTemplatesDirectorySettingTemplateItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type DirectorySettingTemplatesDirectorySettingTemplateItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *DirectorySettingTemplatesDirectorySettingTemplateItemRequestBuilderGetQueryParameters
}
// DirectorySettingTemplatesDirectorySettingTemplateItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type DirectorySettingTemplatesDirectorySettingTemplateItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// CheckMemberGroups provides operations to call the checkMemberGroups method.
func (m *DirectorySettingTemplatesDirectorySettingTemplateItemRequestBuilder) CheckMemberGroups()(*DirectorySettingTemplatesItemCheckMemberGroupsRequestBuilder) {
    return NewDirectorySettingTemplatesItemCheckMemberGroupsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// CheckMemberObjects provides operations to call the checkMemberObjects method.
func (m *DirectorySettingTemplatesDirectorySettingTemplateItemRequestBuilder) CheckMemberObjects()(*DirectorySettingTemplatesItemCheckMemberObjectsRequestBuilder) {
    return NewDirectorySettingTemplatesItemCheckMemberObjectsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// NewDirectorySettingTemplatesDirectorySettingTemplateItemRequestBuilderInternal instantiates a new DirectorySettingTemplateItemRequestBuilder and sets the default values.
func NewDirectorySettingTemplatesDirectorySettingTemplateItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*DirectorySettingTemplatesDirectorySettingTemplateItemRequestBuilder) {
    m := &DirectorySettingTemplatesDirectorySettingTemplateItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/directorySettingTemplates/{directorySettingTemplate%2Did}{?%24select,%24expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewDirectorySettingTemplatesDirectorySettingTemplateItemRequestBuilder instantiates a new DirectorySettingTemplateItemRequestBuilder and sets the default values.
func NewDirectorySettingTemplatesDirectorySettingTemplateItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*DirectorySettingTemplatesDirectorySettingTemplateItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewDirectorySettingTemplatesDirectorySettingTemplateItemRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateDeleteRequestInformation delete entity from directorySettingTemplates
func (m *DirectorySettingTemplatesDirectorySettingTemplateItemRequestBuilder) CreateDeleteRequestInformation(ctx context.Context, requestConfiguration *DirectorySettingTemplatesDirectorySettingTemplateItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE
    if requestConfiguration != nil {
        requestInfo.AddRequestHeaders(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// CreateGetRequestInformation a directory setting template represents a template of settings from which settings may be created within a tenant. This operation allows retrieval of the properties of the **directorySettingTemplate** object, including the available settings and their defaults.
func (m *DirectorySettingTemplatesDirectorySettingTemplateItemRequestBuilder) CreateGetRequestInformation(ctx context.Context, requestConfiguration *DirectorySettingTemplatesDirectorySettingTemplateItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET
    requestInfo.Headers["Accept"] = "application/json"
    if requestConfiguration != nil {
        if requestConfiguration.QueryParameters != nil {
            requestInfo.AddQueryParameters(*(requestConfiguration.QueryParameters))
        }
        requestInfo.AddRequestHeaders(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// CreatePatchRequestInformation update entity in directorySettingTemplates
func (m *DirectorySettingTemplatesDirectorySettingTemplateItemRequestBuilder) CreatePatchRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DirectorySettingTemplateable, requestConfiguration *DirectorySettingTemplatesDirectorySettingTemplateItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.PATCH
    requestInfo.Headers["Accept"] = "application/json"
    requestInfo.SetContentFromParsable(ctx, m.requestAdapter, "application/json", body)
    if requestConfiguration != nil {
        requestInfo.AddRequestHeaders(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// Delete delete entity from directorySettingTemplates
func (m *DirectorySettingTemplatesDirectorySettingTemplateItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *DirectorySettingTemplatesDirectorySettingTemplateItemRequestBuilderDeleteRequestConfiguration)(error) {
    requestInfo, err := m.CreateDeleteRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    err = m.requestAdapter.SendNoContentAsync(ctx, requestInfo, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
// Get a directory setting template represents a template of settings from which settings may be created within a tenant. This operation allows retrieval of the properties of the **directorySettingTemplate** object, including the available settings and their defaults.
func (m *DirectorySettingTemplatesDirectorySettingTemplateItemRequestBuilder) Get(ctx context.Context, requestConfiguration *DirectorySettingTemplatesDirectorySettingTemplateItemRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DirectorySettingTemplateable, error) {
    requestInfo, err := m.CreateGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateDirectorySettingTemplateFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DirectorySettingTemplateable), nil
}
// GetMemberGroups provides operations to call the getMemberGroups method.
func (m *DirectorySettingTemplatesDirectorySettingTemplateItemRequestBuilder) GetMemberGroups()(*DirectorySettingTemplatesItemGetMemberGroupsRequestBuilder) {
    return NewDirectorySettingTemplatesItemGetMemberGroupsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// GetMemberObjects provides operations to call the getMemberObjects method.
func (m *DirectorySettingTemplatesDirectorySettingTemplateItemRequestBuilder) GetMemberObjects()(*DirectorySettingTemplatesItemGetMemberObjectsRequestBuilder) {
    return NewDirectorySettingTemplatesItemGetMemberObjectsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Patch update entity in directorySettingTemplates
func (m *DirectorySettingTemplatesDirectorySettingTemplateItemRequestBuilder) Patch(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DirectorySettingTemplateable, requestConfiguration *DirectorySettingTemplatesDirectorySettingTemplateItemRequestBuilderPatchRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DirectorySettingTemplateable, error) {
    requestInfo, err := m.CreatePatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateDirectorySettingTemplateFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DirectorySettingTemplateable), nil
}
// Restore provides operations to call the restore method.
func (m *DirectorySettingTemplatesDirectorySettingTemplateItemRequestBuilder) Restore()(*DirectorySettingTemplatesItemRestoreRequestBuilder) {
    return NewDirectorySettingTemplatesItemRestoreRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
