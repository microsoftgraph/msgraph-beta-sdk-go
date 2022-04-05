package directoryroles

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
    i28acc41a9ddaa5012f518de9ebb17e81fc5d01a5bb443c4043b0cabb5ae0b5ac "github.com/microsoftgraph/msgraph-beta-sdk-go/directoryroles/count"
    i686e52110d509abde81e2ec02a4b1cacb26be255a28efb139b5fe07005585dcb "github.com/microsoftgraph/msgraph-beta-sdk-go/directoryroles/getuserownedobjects"
    i6c75901f09b564ec245b7b3a868272e273083db1b9879e0b05d7329117fa6e34 "github.com/microsoftgraph/msgraph-beta-sdk-go/directoryroles/getbyids"
    i89f3d18e33abe74e12cff2eb112be1d34453a000675f0e7670b07807419b57ff "github.com/microsoftgraph/msgraph-beta-sdk-go/directoryroles/delta"
    ifd6f8f935853943530de1cc46edc30d54c90acd20d9483c7525660fa0c70d33f "github.com/microsoftgraph/msgraph-beta-sdk-go/directoryroles/validateproperties"
)

// DirectoryRolesRequestBuilder provides operations to manage the collection of directoryRole entities.
type DirectoryRolesRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// DirectoryRolesRequestBuilderGetOptions options for Get
type DirectoryRolesRequestBuilderGetOptions struct {
    // Request headers
    Headers map[string]string;
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption;
    // Request query parameters
    QueryParameters *DirectoryRolesRequestBuilderGetQueryParameters;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler;
}
// DirectoryRolesRequestBuilderGetQueryParameters get entities from directoryRoles
type DirectoryRolesRequestBuilderGetQueryParameters struct {
    // Include count of items
    Count *bool;
    // Expand related entities
    Expand []string;
    // Filter items by property values
    Filter *string;
    // Order items by property values
    Orderby []string;
    // Search items by search phrases
    Search *string;
    // Select properties to be returned
    Select []string;
    // Skip the first n items
    Skip *int32;
}
// DirectoryRolesRequestBuilderPostOptions options for Post
type DirectoryRolesRequestBuilderPostOptions struct {
    // 
    Body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DirectoryRoleable;
    // Request headers
    Headers map[string]string;
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler;
}
// NewDirectoryRolesRequestBuilderInternal instantiates a new DirectoryRolesRequestBuilder and sets the default values.
func NewDirectoryRolesRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*DirectoryRolesRequestBuilder) {
    m := &DirectoryRolesRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/directoryRoles{?skip,search,filter,count,orderby,select,expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewDirectoryRolesRequestBuilder instantiates a new DirectoryRolesRequestBuilder and sets the default values.
func NewDirectoryRolesRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*DirectoryRolesRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewDirectoryRolesRequestBuilderInternal(urlParams, requestAdapter)
}
// Count the count property
func (m *DirectoryRolesRequestBuilder) Count()(*i28acc41a9ddaa5012f518de9ebb17e81fc5d01a5bb443c4043b0cabb5ae0b5ac.CountRequestBuilder) {
    return i28acc41a9ddaa5012f518de9ebb17e81fc5d01a5bb443c4043b0cabb5ae0b5ac.NewCountRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// CreateGetRequestInformation get entities from directoryRoles
func (m *DirectoryRolesRequestBuilder) CreateGetRequestInformation(options *DirectoryRolesRequestBuilderGetOptions)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// CreatePostRequestInformation add new entity to directoryRoles
func (m *DirectoryRolesRequestBuilder) CreatePostRequestInformation(options *DirectoryRolesRequestBuilderPostOptions)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// Delta provides operations to call the delta method.
func (m *DirectoryRolesRequestBuilder) Delta()(*i89f3d18e33abe74e12cff2eb112be1d34453a000675f0e7670b07807419b57ff.DeltaRequestBuilder) {
    return i89f3d18e33abe74e12cff2eb112be1d34453a000675f0e7670b07807419b57ff.NewDeltaRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Get get entities from directoryRoles
func (m *DirectoryRolesRequestBuilder) Get(options *DirectoryRolesRequestBuilderGetOptions)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DirectoryRoleCollectionResponseable, error) {
    requestInfo, err := m.CreateGetRequestInformation(options);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateDirectoryRoleCollectionResponseFromDiscriminatorValue, nil, errorMapping)
    if err != nil {
        return nil, err
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DirectoryRoleCollectionResponseable), nil
}
// GetByIds the getByIds property
func (m *DirectoryRolesRequestBuilder) GetByIds()(*i6c75901f09b564ec245b7b3a868272e273083db1b9879e0b05d7329117fa6e34.GetByIdsRequestBuilder) {
    return i6c75901f09b564ec245b7b3a868272e273083db1b9879e0b05d7329117fa6e34.NewGetByIdsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// GetUserOwnedObjects the getUserOwnedObjects property
func (m *DirectoryRolesRequestBuilder) GetUserOwnedObjects()(*i686e52110d509abde81e2ec02a4b1cacb26be255a28efb139b5fe07005585dcb.GetUserOwnedObjectsRequestBuilder) {
    return i686e52110d509abde81e2ec02a4b1cacb26be255a28efb139b5fe07005585dcb.NewGetUserOwnedObjectsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Post add new entity to directoryRoles
func (m *DirectoryRolesRequestBuilder) Post(options *DirectoryRolesRequestBuilderPostOptions)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DirectoryRoleable, error) {
    requestInfo, err := m.CreatePostRequestInformation(options);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateDirectoryRoleFromDiscriminatorValue, nil, errorMapping)
    if err != nil {
        return nil, err
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DirectoryRoleable), nil
}
// ValidateProperties the validateProperties property
func (m *DirectoryRolesRequestBuilder) ValidateProperties()(*ifd6f8f935853943530de1cc46edc30d54c90acd20d9483c7525660fa0c70d33f.ValidatePropertiesRequestBuilder) {
    return ifd6f8f935853943530de1cc46edc30d54c90acd20d9483c7525660fa0c70d33f.NewValidatePropertiesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
