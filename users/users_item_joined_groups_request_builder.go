package users

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// UsersItemJoinedGroupsRequestBuilder provides operations to manage the joinedGroups property of the microsoft.graph.user entity.
type UsersItemJoinedGroupsRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// UsersItemJoinedGroupsRequestBuilderGetQueryParameters get joinedGroups from users
type UsersItemJoinedGroupsRequestBuilderGetQueryParameters struct {
    // Include count of items
    Count *bool `uriparametername:"%24count"`
    // Filter items by property values
    Filter *string `uriparametername:"%24filter"`
    // Order items by property values
    Orderby []string `uriparametername:"%24orderby"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
    // Skip the first n items
    Skip *int32 `uriparametername:"%24skip"`
    // Show only the first n items
    Top *int32 `uriparametername:"%24top"`
}
// UsersItemJoinedGroupsRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type UsersItemJoinedGroupsRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *UsersItemJoinedGroupsRequestBuilderGetQueryParameters
}
// NewUsersItemJoinedGroupsRequestBuilderInternal instantiates a new JoinedGroupsRequestBuilder and sets the default values.
func NewUsersItemJoinedGroupsRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*UsersItemJoinedGroupsRequestBuilder) {
    m := &UsersItemJoinedGroupsRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/users/{user%2Did}/joinedGroups{?%24top,%24skip,%24filter,%24count,%24orderby,%24select}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewUsersItemJoinedGroupsRequestBuilder instantiates a new JoinedGroupsRequestBuilder and sets the default values.
func NewUsersItemJoinedGroupsRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*UsersItemJoinedGroupsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewUsersItemJoinedGroupsRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateGetRequestInformation get joinedGroups from users
func (m *UsersItemJoinedGroupsRequestBuilder) CreateGetRequestInformation(ctx context.Context, requestConfiguration *UsersItemJoinedGroupsRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// Delta provides operations to call the delta method.
func (m *UsersItemJoinedGroupsRequestBuilder) Delta()(*UsersItemJoinedGroupsDeltaRequestBuilder) {
    return NewUsersItemJoinedGroupsDeltaRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// EvaluateDynamicMembership provides operations to call the evaluateDynamicMembership method.
func (m *UsersItemJoinedGroupsRequestBuilder) EvaluateDynamicMembership()(*UsersItemJoinedGroupsEvaluateDynamicMembershipRequestBuilder) {
    return NewUsersItemJoinedGroupsEvaluateDynamicMembershipRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Get get joinedGroups from users
func (m *UsersItemJoinedGroupsRequestBuilder) Get(ctx context.Context, requestConfiguration *UsersItemJoinedGroupsRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.GroupCollectionResponseable, error) {
    requestInfo, err := m.CreateGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateGroupCollectionResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.GroupCollectionResponseable), nil
}
// GetByIds provides operations to call the getByIds method.
func (m *UsersItemJoinedGroupsRequestBuilder) GetByIds()(*UsersItemJoinedGroupsGetByIdsRequestBuilder) {
    return NewUsersItemJoinedGroupsGetByIdsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// GetUserOwnedObjects provides operations to call the getUserOwnedObjects method.
func (m *UsersItemJoinedGroupsRequestBuilder) GetUserOwnedObjects()(*UsersItemJoinedGroupsGetUserOwnedObjectsRequestBuilder) {
    return NewUsersItemJoinedGroupsGetUserOwnedObjectsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ValidateProperties provides operations to call the validateProperties method.
func (m *UsersItemJoinedGroupsRequestBuilder) ValidateProperties()(*UsersItemJoinedGroupsValidatePropertiesRequestBuilder) {
    return NewUsersItemJoinedGroupsValidatePropertiesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
