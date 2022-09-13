package directorysettingtemplates

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i03094d9d1021b32e492d609c092d84ed4cfb558c32f0c8ade3ad08ceb1eead66 "github.com/microsoftgraph/msgraph-beta-sdk-go/directorysettingtemplates/validateproperties"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
    i8b3159fe7053462b6a5b7687099a8e653ffaaa3e49ac1c52748ad001ac05d4d9 "github.com/microsoftgraph/msgraph-beta-sdk-go/directorysettingtemplates/getbyids"
    if2eec762f4a60cfa480122020666674cae3dd0c0dd67417b7da42d6d7d9efc9f "github.com/microsoftgraph/msgraph-beta-sdk-go/directorysettingtemplates/getuserownedobjects"
)

// DirectorySettingTemplatesRequestBuilder provides operations to manage the collection of directorySettingTemplate entities.
type DirectorySettingTemplatesRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// DirectorySettingTemplatesRequestBuilderGetQueryParameters directory setting templates represents a set of templates of directory settings, from which directory settings may be created and used within a tenant.  This operation retrieves the list of available **directorySettingTemplates** objects.
type DirectorySettingTemplatesRequestBuilderGetQueryParameters struct {
    // Order items by property values
    Orderby []string `uriparametername:"%24orderby"`
    // Search items by search phrases
    Search *string `uriparametername:"%24search"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
    // Show only the first n items
    Top *int32 `uriparametername:"%24top"`
}
// DirectorySettingTemplatesRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type DirectorySettingTemplatesRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *DirectorySettingTemplatesRequestBuilderGetQueryParameters
}
// DirectorySettingTemplatesRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type DirectorySettingTemplatesRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewDirectorySettingTemplatesRequestBuilderInternal instantiates a new DirectorySettingTemplatesRequestBuilder and sets the default values.
func NewDirectorySettingTemplatesRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*DirectorySettingTemplatesRequestBuilder) {
    m := &DirectorySettingTemplatesRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/directorySettingTemplates{?%24top*,%24search*,%24orderby,%24select}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewDirectorySettingTemplatesRequestBuilder instantiates a new DirectorySettingTemplatesRequestBuilder and sets the default values.
func NewDirectorySettingTemplatesRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*DirectorySettingTemplatesRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewDirectorySettingTemplatesRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateGetRequestInformation directory setting templates represents a set of templates of directory settings, from which directory settings may be created and used within a tenant.  This operation retrieves the list of available **directorySettingTemplates** objects.
func (m *DirectorySettingTemplatesRequestBuilder) CreateGetRequestInformation()(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    return m.CreateGetRequestInformationWithRequestConfiguration(nil);
}
// CreateGetRequestInformationWithRequestConfiguration directory setting templates represents a set of templates of directory settings, from which directory settings may be created and used within a tenant.  This operation retrieves the list of available **directorySettingTemplates** objects.
func (m *DirectorySettingTemplatesRequestBuilder) CreateGetRequestInformationWithRequestConfiguration(requestConfiguration *DirectorySettingTemplatesRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// CreatePostRequestInformation add new entity to directorySettingTemplates
func (m *DirectorySettingTemplatesRequestBuilder) CreatePostRequestInformation(body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DirectorySettingTemplateable)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    return m.CreatePostRequestInformationWithRequestConfiguration(body, nil);
}
// CreatePostRequestInformationWithRequestConfiguration add new entity to directorySettingTemplates
func (m *DirectorySettingTemplatesRequestBuilder) CreatePostRequestInformationWithRequestConfiguration(body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DirectorySettingTemplateable, requestConfiguration *DirectorySettingTemplatesRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST
    requestInfo.Headers["Accept"] = "application/json"
    requestInfo.SetContentFromParsable(m.requestAdapter, "application/json", body)
    if requestConfiguration != nil {
        requestInfo.AddRequestHeaders(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// Get directory setting templates represents a set of templates of directory settings, from which directory settings may be created and used within a tenant.  This operation retrieves the list of available **directorySettingTemplates** objects.
func (m *DirectorySettingTemplatesRequestBuilder) Get(ctx context.Context, requestConfiguration *DirectorySettingTemplatesRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DirectorySettingTemplateCollectionResponseable, error) {
    requestInfo, err := m.CreateGetRequestInformationWithRequestConfiguration(requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateDirectorySettingTemplateCollectionResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DirectorySettingTemplateCollectionResponseable), nil
}
// GetByIds the getByIds property
func (m *DirectorySettingTemplatesRequestBuilder) GetByIds()(*i8b3159fe7053462b6a5b7687099a8e653ffaaa3e49ac1c52748ad001ac05d4d9.GetByIdsRequestBuilder) {
    return i8b3159fe7053462b6a5b7687099a8e653ffaaa3e49ac1c52748ad001ac05d4d9.NewGetByIdsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// GetUserOwnedObjects the getUserOwnedObjects property
func (m *DirectorySettingTemplatesRequestBuilder) GetUserOwnedObjects()(*if2eec762f4a60cfa480122020666674cae3dd0c0dd67417b7da42d6d7d9efc9f.GetUserOwnedObjectsRequestBuilder) {
    return if2eec762f4a60cfa480122020666674cae3dd0c0dd67417b7da42d6d7d9efc9f.NewGetUserOwnedObjectsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Post add new entity to directorySettingTemplates
func (m *DirectorySettingTemplatesRequestBuilder) Post(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DirectorySettingTemplateable, requestConfiguration *DirectorySettingTemplatesRequestBuilderPostRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DirectorySettingTemplateable, error) {
    requestInfo, err := m.CreatePostRequestInformationWithRequestConfiguration(body, requestConfiguration);
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
// ValidateProperties the validateProperties property
func (m *DirectorySettingTemplatesRequestBuilder) ValidateProperties()(*i03094d9d1021b32e492d609c092d84ed4cfb558c32f0c8ade3ad08ceb1eead66.ValidatePropertiesRequestBuilder) {
    return i03094d9d1021b32e492d609c092d84ed4cfb558c32f0c8ade3ad08ceb1eead66.NewValidatePropertiesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
