package joinedgroups

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
    i14e2b33f6eebbf36864b149113f785d4d4fb69c0b082832eecbb0ade5d528ba0 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/joinedgroups/evaluatedynamicmembership"
    i9ea560dd3a1ef2bd400a3c9b5a56daafabdf774542c588e5bcbc538ff7be9937 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/joinedgroups/validateproperties"
    ia4b17264996a812540aa98efa53ac88ff9f3de6dc41b3a974d28cd88490f84e8 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/joinedgroups/getuserownedobjects"
    ib4b15e2432646bfa285596911d9fb56acb5ace28adfc5f56de348210723dbc40 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/joinedgroups/delta"
    ic01f1a71e9d81daccfa59f2811480ab0b9b91b235a54a1bf38301799ab722207 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/joinedgroups/getbyids"
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
// JoinedGroupsRequestBuilderGetQueryParameters get joinedGroups from me
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
// NewJoinedGroupsRequestBuilderInternal instantiates a new JoinedGroupsRequestBuilder and sets the default values.
func NewJoinedGroupsRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*JoinedGroupsRequestBuilder) {
    m := &JoinedGroupsRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/me/joinedGroups{?%24top,%24skip,%24filter,%24count,%24orderby,%24select}";
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
// CreateGetRequestInformation get joinedGroups from me
func (m *JoinedGroupsRequestBuilder) CreateGetRequestInformation()(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    return m.CreateGetRequestInformationWithRequestConfiguration(nil);
}
// CreateGetRequestInformationWithRequestConfiguration get joinedGroups from me
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
// Delta provides operations to call the delta method.
func (m *JoinedGroupsRequestBuilder) Delta()(*ib4b15e2432646bfa285596911d9fb56acb5ace28adfc5f56de348210723dbc40.DeltaRequestBuilder) {
    return ib4b15e2432646bfa285596911d9fb56acb5ace28adfc5f56de348210723dbc40.NewDeltaRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// EvaluateDynamicMembership the evaluateDynamicMembership property
func (m *JoinedGroupsRequestBuilder) EvaluateDynamicMembership()(*i14e2b33f6eebbf36864b149113f785d4d4fb69c0b082832eecbb0ade5d528ba0.EvaluateDynamicMembershipRequestBuilder) {
    return i14e2b33f6eebbf36864b149113f785d4d4fb69c0b082832eecbb0ade5d528ba0.NewEvaluateDynamicMembershipRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Get get joinedGroups from me
func (m *JoinedGroupsRequestBuilder) Get(ctx context.Context, requestConfiguration *JoinedGroupsRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.GroupCollectionResponseable, error) {
    requestInfo, err := m.CreateGetRequestInformationWithRequestConfiguration(requestConfiguration);
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
// GetByIds the getByIds property
func (m *JoinedGroupsRequestBuilder) GetByIds()(*ic01f1a71e9d81daccfa59f2811480ab0b9b91b235a54a1bf38301799ab722207.GetByIdsRequestBuilder) {
    return ic01f1a71e9d81daccfa59f2811480ab0b9b91b235a54a1bf38301799ab722207.NewGetByIdsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// GetUserOwnedObjects the getUserOwnedObjects property
func (m *JoinedGroupsRequestBuilder) GetUserOwnedObjects()(*ia4b17264996a812540aa98efa53ac88ff9f3de6dc41b3a974d28cd88490f84e8.GetUserOwnedObjectsRequestBuilder) {
    return ia4b17264996a812540aa98efa53ac88ff9f3de6dc41b3a974d28cd88490f84e8.NewGetUserOwnedObjectsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ValidateProperties the validateProperties property
func (m *JoinedGroupsRequestBuilder) ValidateProperties()(*i9ea560dd3a1ef2bd400a3c9b5a56daafabdf774542c588e5bcbc538ff7be9937.ValidatePropertiesRequestBuilder) {
    return i9ea560dd3a1ef2bd400a3c9b5a56daafabdf774542c588e5bcbc538ff7be9937.NewValidatePropertiesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
