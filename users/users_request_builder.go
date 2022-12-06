package users

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i05710ad707ca5df305bac808ada934e670542cf0f0910503bdc6f9dee7a4e50c "github.com/microsoftgraph/msgraph-beta-sdk-go/users/validateproperties"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
    i2c83b78d88f4ae5d43180bd3ceb5b2b18f835891b4d285c93889e5b4e7dadfa7 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/validatepassword"
    i4955c8ead1dfb331d7b6c5cb84c83e17fe654822a84f8a3c4f9681b987a8f982 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/count"
    i5c97c863f904b1cf0ebde61c61539e9095b7a10833dbc1282cd7123559514275 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/getmanagedappblockedusers"
    i7c59b4c4e2615cc579520999d700128005226f408f7b8111977ed0c3a6a45d70 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/getuserownedobjects"
    i94b36155ccc2661432b4d3a1f7cfdb2c9c0142dd79eb78527be593b4136386be "github.com/microsoftgraph/msgraph-beta-sdk-go/users/delta"
    ia70b48171890bd6eba4caf9c2081ddd6213e4cd72a7d11bfda6ecbcc636d32ef "github.com/microsoftgraph/msgraph-beta-sdk-go/users/getbyids"
)

// UsersRequestBuilder provides operations to manage the collection of user entities.
type UsersRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// UsersRequestBuilderGetQueryParameters retrieve a list of user objects. This operation returns by default only a subset of the more commonly used properties for each user. These _default_ properties are noted in the Properties section. To get properties that are _not_ returned by default, do a GET operation for the user and specify the properties in a `$select` OData query option.
type UsersRequestBuilderGetQueryParameters struct {
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
// UsersRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type UsersRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *UsersRequestBuilderGetQueryParameters
}
// UsersRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type UsersRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewUsersRequestBuilderInternal instantiates a new UsersRequestBuilder and sets the default values.
func NewUsersRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*UsersRequestBuilder) {
    m := &UsersRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/users{?%24top,%24skip,%24search,%24filter,%24count,%24orderby,%24select,%24expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewUsersRequestBuilder instantiates a new UsersRequestBuilder and sets the default values.
func NewUsersRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*UsersRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewUsersRequestBuilderInternal(urlParams, requestAdapter)
}
// Count provides operations to count the resources in the collection.
func (m *UsersRequestBuilder) Count()(*i4955c8ead1dfb331d7b6c5cb84c83e17fe654822a84f8a3c4f9681b987a8f982.CountRequestBuilder) {
    return i4955c8ead1dfb331d7b6c5cb84c83e17fe654822a84f8a3c4f9681b987a8f982.NewCountRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// CreateGetRequestInformation retrieve a list of user objects. This operation returns by default only a subset of the more commonly used properties for each user. These _default_ properties are noted in the Properties section. To get properties that are _not_ returned by default, do a GET operation for the user and specify the properties in a `$select` OData query option.
func (m *UsersRequestBuilder) CreateGetRequestInformation(ctx context.Context, requestConfiguration *UsersRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// CreatePostRequestInformation create a new user.The request body contains the user to create. At a minimum, you must specify the required properties for the user. You can optionally specify any other writable properties. This operation returns by default only a subset of the properties for each user. These default properties are noted in the Properties section. To get properties that are not returned by default, do a GET operation and specify the properties in a `$select` OData query option.
func (m *UsersRequestBuilder) CreatePostRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Userable, requestConfiguration *UsersRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST
    requestInfo.Headers["Accept"] = "application/json"
    requestInfo.SetContentFromParsable(ctx, m.requestAdapter, "application/json", body)
    if requestConfiguration != nil {
        requestInfo.AddRequestHeaders(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// Delta provides operations to call the delta method.
func (m *UsersRequestBuilder) Delta()(*i94b36155ccc2661432b4d3a1f7cfdb2c9c0142dd79eb78527be593b4136386be.DeltaRequestBuilder) {
    return i94b36155ccc2661432b4d3a1f7cfdb2c9c0142dd79eb78527be593b4136386be.NewDeltaRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Get retrieve a list of user objects. This operation returns by default only a subset of the more commonly used properties for each user. These _default_ properties are noted in the Properties section. To get properties that are _not_ returned by default, do a GET operation for the user and specify the properties in a `$select` OData query option.
func (m *UsersRequestBuilder) Get(ctx context.Context, requestConfiguration *UsersRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.UserCollectionResponseable, error) {
    requestInfo, err := m.CreateGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateUserCollectionResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.UserCollectionResponseable), nil
}
// GetByIds provides operations to call the getByIds method.
func (m *UsersRequestBuilder) GetByIds()(*ia70b48171890bd6eba4caf9c2081ddd6213e4cd72a7d11bfda6ecbcc636d32ef.GetByIdsRequestBuilder) {
    return ia70b48171890bd6eba4caf9c2081ddd6213e4cd72a7d11bfda6ecbcc636d32ef.NewGetByIdsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// GetManagedAppBlockedUsers provides operations to call the getManagedAppBlockedUsers method.
func (m *UsersRequestBuilder) GetManagedAppBlockedUsers()(*i5c97c863f904b1cf0ebde61c61539e9095b7a10833dbc1282cd7123559514275.GetManagedAppBlockedUsersRequestBuilder) {
    return i5c97c863f904b1cf0ebde61c61539e9095b7a10833dbc1282cd7123559514275.NewGetManagedAppBlockedUsersRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// GetUserOwnedObjects provides operations to call the getUserOwnedObjects method.
func (m *UsersRequestBuilder) GetUserOwnedObjects()(*i7c59b4c4e2615cc579520999d700128005226f408f7b8111977ed0c3a6a45d70.GetUserOwnedObjectsRequestBuilder) {
    return i7c59b4c4e2615cc579520999d700128005226f408f7b8111977ed0c3a6a45d70.NewGetUserOwnedObjectsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Post create a new user.The request body contains the user to create. At a minimum, you must specify the required properties for the user. You can optionally specify any other writable properties. This operation returns by default only a subset of the properties for each user. These default properties are noted in the Properties section. To get properties that are not returned by default, do a GET operation and specify the properties in a `$select` OData query option.
func (m *UsersRequestBuilder) Post(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Userable, requestConfiguration *UsersRequestBuilderPostRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Userable, error) {
    requestInfo, err := m.CreatePostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateUserFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Userable), nil
}
// ValidatePassword provides operations to call the validatePassword method.
func (m *UsersRequestBuilder) ValidatePassword()(*i2c83b78d88f4ae5d43180bd3ceb5b2b18f835891b4d285c93889e5b4e7dadfa7.ValidatePasswordRequestBuilder) {
    return i2c83b78d88f4ae5d43180bd3ceb5b2b18f835891b4d285c93889e5b4e7dadfa7.NewValidatePasswordRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ValidateProperties provides operations to call the validateProperties method.
func (m *UsersRequestBuilder) ValidateProperties()(*i05710ad707ca5df305bac808ada934e670542cf0f0910503bdc6f9dee7a4e50c.ValidatePropertiesRequestBuilder) {
    return i05710ad707ca5df305bac808ada934e670542cf0f0910503bdc6f9dee7a4e50c.NewValidatePropertiesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
