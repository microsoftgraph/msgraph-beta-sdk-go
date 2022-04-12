package importeddeviceidentities

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
    i2eb32d0eb94562e9b6f79567b48fe799dded573649031029fb86f8965772271e "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/importeddeviceidentities/searchexistingidentities"
    i480d1fa84691984b46caafa0c9d056d9c38c0452d47b9d4829c099dde2a6f199 "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/importeddeviceidentities/count"
    i76339fd2d40042286c7a9b4333ff7246712608a8ec9a407e5fc00e4a39473d6a "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/importeddeviceidentities/importdeviceidentitylist"
)

// ImportedDeviceIdentitiesRequestBuilder provides operations to manage the importedDeviceIdentities property of the microsoft.graph.deviceManagement entity.
type ImportedDeviceIdentitiesRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// ImportedDeviceIdentitiesRequestBuilderGetOptions options for Get
type ImportedDeviceIdentitiesRequestBuilderGetOptions struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ImportedDeviceIdentitiesRequestBuilderGetQueryParameters
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler
}
// ImportedDeviceIdentitiesRequestBuilderGetQueryParameters the imported device identities.
type ImportedDeviceIdentitiesRequestBuilderGetQueryParameters struct {
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
// ImportedDeviceIdentitiesRequestBuilderPostOptions options for Post
type ImportedDeviceIdentitiesRequestBuilderPostOptions struct {
    // 
    Body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ImportedDeviceIdentityable
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler
}
// NewImportedDeviceIdentitiesRequestBuilderInternal instantiates a new ImportedDeviceIdentitiesRequestBuilder and sets the default values.
func NewImportedDeviceIdentitiesRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ImportedDeviceIdentitiesRequestBuilder) {
    m := &ImportedDeviceIdentitiesRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/deviceManagement/importedDeviceIdentities{?%24top,%24skip,%24search,%24filter,%24count,%24orderby,%24select,%24expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewImportedDeviceIdentitiesRequestBuilder instantiates a new ImportedDeviceIdentitiesRequestBuilder and sets the default values.
func NewImportedDeviceIdentitiesRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ImportedDeviceIdentitiesRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewImportedDeviceIdentitiesRequestBuilderInternal(urlParams, requestAdapter)
}
// Count the count property
func (m *ImportedDeviceIdentitiesRequestBuilder) Count()(*i480d1fa84691984b46caafa0c9d056d9c38c0452d47b9d4829c099dde2a6f199.CountRequestBuilder) {
    return i480d1fa84691984b46caafa0c9d056d9c38c0452d47b9d4829c099dde2a6f199.NewCountRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// CreateGetRequestInformation the imported device identities.
func (m *ImportedDeviceIdentitiesRequestBuilder) CreateGetRequestInformation(options *ImportedDeviceIdentitiesRequestBuilderGetOptions)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET
    if options != nil && options.QueryParameters != nil {
        requestInfo.AddQueryParameters(*(options.QueryParameters))
    }
    if options != nil && options.Headers != nil {
        requestInfo.Headers = options.Headers
    }
    if options != nil && len(options.Options) != 0 {
        err := requestInfo.AddRequestOptions(options.Options...)
        if err != nil {
            return nil, err
        }
    }
    return requestInfo, nil
}
// CreatePostRequestInformation create new navigation property to importedDeviceIdentities for deviceManagement
func (m *ImportedDeviceIdentitiesRequestBuilder) CreatePostRequestInformation(options *ImportedDeviceIdentitiesRequestBuilderPostOptions)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST
    requestInfo.SetContentFromParsable(m.requestAdapter, "application/json", options.Body)
    if options != nil && options.Headers != nil {
        requestInfo.Headers = options.Headers
    }
    if options != nil && len(options.Options) != 0 {
        err := requestInfo.AddRequestOptions(options.Options...)
        if err != nil {
            return nil, err
        }
    }
    return requestInfo, nil
}
// Get the imported device identities.
func (m *ImportedDeviceIdentitiesRequestBuilder) Get(options *ImportedDeviceIdentitiesRequestBuilderGetOptions)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ImportedDeviceIdentityCollectionResponseable, error) {
    requestInfo, err := m.CreateGetRequestInformation(options);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateImportedDeviceIdentityCollectionResponseFromDiscriminatorValue, nil, errorMapping)
    if err != nil {
        return nil, err
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ImportedDeviceIdentityCollectionResponseable), nil
}
// ImportDeviceIdentityList the importDeviceIdentityList property
func (m *ImportedDeviceIdentitiesRequestBuilder) ImportDeviceIdentityList()(*i76339fd2d40042286c7a9b4333ff7246712608a8ec9a407e5fc00e4a39473d6a.ImportDeviceIdentityListRequestBuilder) {
    return i76339fd2d40042286c7a9b4333ff7246712608a8ec9a407e5fc00e4a39473d6a.NewImportDeviceIdentityListRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Post create new navigation property to importedDeviceIdentities for deviceManagement
func (m *ImportedDeviceIdentitiesRequestBuilder) Post(options *ImportedDeviceIdentitiesRequestBuilderPostOptions)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ImportedDeviceIdentityable, error) {
    requestInfo, err := m.CreatePostRequestInformation(options);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateImportedDeviceIdentityFromDiscriminatorValue, nil, errorMapping)
    if err != nil {
        return nil, err
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ImportedDeviceIdentityable), nil
}
// SearchExistingIdentities the searchExistingIdentities property
func (m *ImportedDeviceIdentitiesRequestBuilder) SearchExistingIdentities()(*i2eb32d0eb94562e9b6f79567b48fe799dded573649031029fb86f8965772271e.SearchExistingIdentitiesRequestBuilder) {
    return i2eb32d0eb94562e9b6f79567b48fe799dded573649031029fb86f8965772271e.NewSearchExistingIdentitiesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
