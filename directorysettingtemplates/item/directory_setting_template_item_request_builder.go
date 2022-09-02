package item

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
    i09e0fb5ad1ad570bae5ac63f3e01cd3eeb4d487f97efe0766919eec487abbd14 "github.com/microsoftgraph/msgraph-beta-sdk-go/directorysettingtemplates/item/getmembergroups"
    i350af3bac14d3575d18b98f3718632a63e909339a7e1d311aeed01b007c6326e "github.com/microsoftgraph/msgraph-beta-sdk-go/directorysettingtemplates/item/getmemberobjects"
    i800ea0955c4120bd8ef6b78925bd3a2e6b2ce9f4c22580b1ebfc361c3da0db2e "github.com/microsoftgraph/msgraph-beta-sdk-go/directorysettingtemplates/item/checkmembergroups"
    i9aa94a0c065ec17979e9f863f95b6a7634f5da9003d00118acd3920a42faad24 "github.com/microsoftgraph/msgraph-beta-sdk-go/directorysettingtemplates/item/checkmemberobjects"
    i9c52bd45abe999d984316ab14274dc18a68a5c506277c7ca4405f3607edc74ee "github.com/microsoftgraph/msgraph-beta-sdk-go/directorysettingtemplates/item/restore"
)

// DirectorySettingTemplateItemRequestBuilder provides operations to manage the collection of directorySettingTemplate entities.
type DirectorySettingTemplateItemRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// DirectorySettingTemplateItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type DirectorySettingTemplateItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// DirectorySettingTemplateItemRequestBuilderGetQueryParameters a directory setting template represents a template of settings from which settings may be created within a tenant. This operation allows retrieval of the properties of the **directorySettingTemplate** object, including the available settings and their defaults.
type DirectorySettingTemplateItemRequestBuilderGetQueryParameters struct {
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// DirectorySettingTemplateItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type DirectorySettingTemplateItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *DirectorySettingTemplateItemRequestBuilderGetQueryParameters
}
// DirectorySettingTemplateItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type DirectorySettingTemplateItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// CheckMemberGroups the checkMemberGroups property
func (m *DirectorySettingTemplateItemRequestBuilder) CheckMemberGroups()(*i800ea0955c4120bd8ef6b78925bd3a2e6b2ce9f4c22580b1ebfc361c3da0db2e.CheckMemberGroupsRequestBuilder) {
    return i800ea0955c4120bd8ef6b78925bd3a2e6b2ce9f4c22580b1ebfc361c3da0db2e.NewCheckMemberGroupsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// CheckMemberObjects the checkMemberObjects property
func (m *DirectorySettingTemplateItemRequestBuilder) CheckMemberObjects()(*i9aa94a0c065ec17979e9f863f95b6a7634f5da9003d00118acd3920a42faad24.CheckMemberObjectsRequestBuilder) {
    return i9aa94a0c065ec17979e9f863f95b6a7634f5da9003d00118acd3920a42faad24.NewCheckMemberObjectsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// NewDirectorySettingTemplateItemRequestBuilderInternal instantiates a new DirectorySettingTemplateItemRequestBuilder and sets the default values.
func NewDirectorySettingTemplateItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*DirectorySettingTemplateItemRequestBuilder) {
    m := &DirectorySettingTemplateItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/directorySettingTemplates/{directorySettingTemplate%2Did}{?%24select}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewDirectorySettingTemplateItemRequestBuilder instantiates a new DirectorySettingTemplateItemRequestBuilder and sets the default values.
func NewDirectorySettingTemplateItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*DirectorySettingTemplateItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewDirectorySettingTemplateItemRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateDeleteRequestInformation delete entity from directorySettingTemplates
func (m *DirectorySettingTemplateItemRequestBuilder) CreateDeleteRequestInformation()(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    return m.CreateDeleteRequestInformationWithRequestConfiguration(nil);
}
// CreateDeleteRequestInformationWithRequestConfiguration delete entity from directorySettingTemplates
func (m *DirectorySettingTemplateItemRequestBuilder) CreateDeleteRequestInformationWithRequestConfiguration(requestConfiguration *DirectorySettingTemplateItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
func (m *DirectorySettingTemplateItemRequestBuilder) CreateGetRequestInformation()(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    return m.CreateGetRequestInformationWithRequestConfiguration(nil);
}
// CreateGetRequestInformationWithRequestConfiguration a directory setting template represents a template of settings from which settings may be created within a tenant. This operation allows retrieval of the properties of the **directorySettingTemplate** object, including the available settings and their defaults.
func (m *DirectorySettingTemplateItemRequestBuilder) CreateGetRequestInformationWithRequestConfiguration(requestConfiguration *DirectorySettingTemplateItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
func (m *DirectorySettingTemplateItemRequestBuilder) CreatePatchRequestInformation(body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DirectorySettingTemplateable)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    return m.CreatePatchRequestInformationWithRequestConfiguration(body, nil);
}
// CreatePatchRequestInformationWithRequestConfiguration update entity in directorySettingTemplates
func (m *DirectorySettingTemplateItemRequestBuilder) CreatePatchRequestInformationWithRequestConfiguration(body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DirectorySettingTemplateable, requestConfiguration *DirectorySettingTemplateItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.PATCH
    requestInfo.SetContentFromParsable(m.requestAdapter, "application/json", body)
    if requestConfiguration != nil {
        requestInfo.AddRequestHeaders(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// Delete delete entity from directorySettingTemplates
func (m *DirectorySettingTemplateItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *DirectorySettingTemplateItemRequestBuilderDeleteRequestConfiguration)(error) {
    requestInfo, err := m.CreateDeleteRequestInformationWithRequestConfiguration(requestConfiguration);
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
func (m *DirectorySettingTemplateItemRequestBuilder) Get(ctx context.Context, requestConfiguration *DirectorySettingTemplateItemRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DirectorySettingTemplateable, error) {
    requestInfo, err := m.CreateGetRequestInformationWithRequestConfiguration(requestConfiguration);
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
// GetMemberGroups the getMemberGroups property
func (m *DirectorySettingTemplateItemRequestBuilder) GetMemberGroups()(*i09e0fb5ad1ad570bae5ac63f3e01cd3eeb4d487f97efe0766919eec487abbd14.GetMemberGroupsRequestBuilder) {
    return i09e0fb5ad1ad570bae5ac63f3e01cd3eeb4d487f97efe0766919eec487abbd14.NewGetMemberGroupsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// GetMemberObjects the getMemberObjects property
func (m *DirectorySettingTemplateItemRequestBuilder) GetMemberObjects()(*i350af3bac14d3575d18b98f3718632a63e909339a7e1d311aeed01b007c6326e.GetMemberObjectsRequestBuilder) {
    return i350af3bac14d3575d18b98f3718632a63e909339a7e1d311aeed01b007c6326e.NewGetMemberObjectsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Patch update entity in directorySettingTemplates
func (m *DirectorySettingTemplateItemRequestBuilder) Patch(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DirectorySettingTemplateable, requestConfiguration *DirectorySettingTemplateItemRequestBuilderPatchRequestConfiguration)(error) {
    requestInfo, err := m.CreatePatchRequestInformationWithRequestConfiguration(body, requestConfiguration);
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
// Restore the restore property
func (m *DirectorySettingTemplateItemRequestBuilder) Restore()(*i9c52bd45abe999d984316ab14274dc18a68a5c506277c7ca4405f3607edc74ee.RestoreRequestBuilder) {
    return i9c52bd45abe999d984316ab14274dc18a68a5c506277c7ca4405f3607edc74ee.NewRestoreRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
