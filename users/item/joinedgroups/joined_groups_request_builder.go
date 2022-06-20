package joinedgroups

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
    i0674cf0527a848b5986f3a53c384773f88bd5e8c2d3afea1e622a5ae746b4dbd "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/joinedgroups/getuserownedobjects"
    i171b181e36fbf7b41352fc29fba26dbe0008f0bd7d3f6fbceb31919fcc9866ff "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/joinedgroups/validateproperties"
    i3f57c1ceec7f528c27ff19be63f0e9d48225fd907beb4bbedf279f35769ff893 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/joinedgroups/getbyids"
    i88857d6c6dbe9d9e7c50f2e09673dd0d012411f1b08616d9d2d304c958fcead1 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/joinedgroups/evaluatedynamicmembership"
    ide9e4ec7463b93fefb042f46f8edd992d184d05d68ef23ca96b5c17a45c5d6cb "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/joinedgroups/delta"
)

// JoinedGroupsRequestBuilder provides operations to manage the joinedGroups property of the microsoft.graph.user entity.
type JoinedGroupsRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// JoinedGroupsRequestBuilderGetQueryParameters get joinedGroups from users
type JoinedGroupsRequestBuilderGetQueryParameters struct {
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
// JoinedGroupsRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type JoinedGroupsRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *JoinedGroupsRequestBuilderGetQueryParameters
}
// JoinedGroupsRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type JoinedGroupsRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewJoinedGroupsRequestBuilderInternal instantiates a new JoinedGroupsRequestBuilder and sets the default values.
func NewJoinedGroupsRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*JoinedGroupsRequestBuilder) {
    m := &JoinedGroupsRequestBuilder{
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
// NewJoinedGroupsRequestBuilder instantiates a new JoinedGroupsRequestBuilder and sets the default values.
func NewJoinedGroupsRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*JoinedGroupsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewJoinedGroupsRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateGetRequestInformation get joinedGroups from users
func (m *JoinedGroupsRequestBuilder) CreateGetRequestInformation()(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    return m.CreateGetRequestInformationWithRequestConfiguration(nil);
}
// CreateGetRequestInformationWithRequestConfiguration get joinedGroups from users
func (m *JoinedGroupsRequestBuilder) CreateGetRequestInformationWithRequestConfiguration(requestConfiguration *JoinedGroupsRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// CreatePostRequestInformation create new navigation property to joinedGroups for users
func (m *JoinedGroupsRequestBuilder) CreatePostRequestInformation(body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Groupable)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    return m.CreatePostRequestInformationWithRequestConfiguration(body, nil);
}
// CreatePostRequestInformationWithRequestConfiguration create new navigation property to joinedGroups for users
func (m *JoinedGroupsRequestBuilder) CreatePostRequestInformationWithRequestConfiguration(body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Groupable, requestConfiguration *JoinedGroupsRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// Delta provides operations to call the delta method.
func (m *JoinedGroupsRequestBuilder) Delta()(*ide9e4ec7463b93fefb042f46f8edd992d184d05d68ef23ca96b5c17a45c5d6cb.DeltaRequestBuilder) {
    return ide9e4ec7463b93fefb042f46f8edd992d184d05d68ef23ca96b5c17a45c5d6cb.NewDeltaRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// EvaluateDynamicMembership the evaluateDynamicMembership property
func (m *JoinedGroupsRequestBuilder) EvaluateDynamicMembership()(*i88857d6c6dbe9d9e7c50f2e09673dd0d012411f1b08616d9d2d304c958fcead1.EvaluateDynamicMembershipRequestBuilder) {
    return i88857d6c6dbe9d9e7c50f2e09673dd0d012411f1b08616d9d2d304c958fcead1.NewEvaluateDynamicMembershipRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Get get joinedGroups from users
func (m *JoinedGroupsRequestBuilder) Get()(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.GroupCollectionResponseable, error) {
    return m.GetWithRequestConfigurationAndResponseHandler(nil, nil);
}
// GetByIds the getByIds property
func (m *JoinedGroupsRequestBuilder) GetByIds()(*i3f57c1ceec7f528c27ff19be63f0e9d48225fd907beb4bbedf279f35769ff893.GetByIdsRequestBuilder) {
    return i3f57c1ceec7f528c27ff19be63f0e9d48225fd907beb4bbedf279f35769ff893.NewGetByIdsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// GetUserOwnedObjects the getUserOwnedObjects property
func (m *JoinedGroupsRequestBuilder) GetUserOwnedObjects()(*i0674cf0527a848b5986f3a53c384773f88bd5e8c2d3afea1e622a5ae746b4dbd.GetUserOwnedObjectsRequestBuilder) {
    return i0674cf0527a848b5986f3a53c384773f88bd5e8c2d3afea1e622a5ae746b4dbd.NewGetUserOwnedObjectsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// GetWithRequestConfigurationAndResponseHandler get joinedGroups from users
func (m *JoinedGroupsRequestBuilder) GetWithRequestConfigurationAndResponseHandler(requestConfiguration *JoinedGroupsRequestBuilderGetRequestConfiguration, responseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.GroupCollectionResponseable, error) {
    requestInfo, err := m.CreateGetRequestInformationWithRequestConfiguration(requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateGroupCollectionResponseFromDiscriminatorValue, responseHandler, errorMapping)
    if err != nil {
        return nil, err
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.GroupCollectionResponseable), nil
}
// Post create new navigation property to joinedGroups for users
func (m *JoinedGroupsRequestBuilder) Post(body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Groupable)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Groupable, error) {
    return m.PostWithRequestConfigurationAndResponseHandler(body, nil, nil);
}
// PostWithRequestConfigurationAndResponseHandler create new navigation property to joinedGroups for users
func (m *JoinedGroupsRequestBuilder) PostWithRequestConfigurationAndResponseHandler(body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Groupable, requestConfiguration *JoinedGroupsRequestBuilderPostRequestConfiguration, responseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Groupable, error) {
    requestInfo, err := m.CreatePostRequestInformationWithRequestConfiguration(body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateGroupFromDiscriminatorValue, responseHandler, errorMapping)
    if err != nil {
        return nil, err
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Groupable), nil
}
// ValidateProperties the validateProperties property
func (m *JoinedGroupsRequestBuilder) ValidateProperties()(*i171b181e36fbf7b41352fc29fba26dbe0008f0bd7d3f6fbceb31919fcc9866ff.ValidatePropertiesRequestBuilder) {
    return i171b181e36fbf7b41352fc29fba26dbe0008f0bd7d3f6fbceb31919fcc9866ff.NewValidatePropertiesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
