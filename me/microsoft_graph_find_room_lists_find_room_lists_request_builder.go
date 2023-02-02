package me

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// MicrosoftGraphFindRoomListsFindRoomListsRequestBuilder provides operations to call the findRoomLists method.
type MicrosoftGraphFindRoomListsFindRoomListsRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// MicrosoftGraphFindRoomListsFindRoomListsRequestBuilderGetQueryParameters get the room lists defined in a tenant, as represented by their emailAddress objects. Tenants can organize meeting rooms into room lists. In this API, each meeting room and room list is represented by an emailAddress instance.You can get all the room lists in the tenant, get all the rooms in the tenant, or get all the rooms in a specific room list.
type MicrosoftGraphFindRoomListsFindRoomListsRequestBuilderGetQueryParameters struct {
    // Include count of items
    Count *bool `uriparametername:"%24count"`
    // Filter items by property values
    Filter *string `uriparametername:"%24filter"`
    // Search items by search phrases
    Search *string `uriparametername:"%24search"`
    // Skip the first n items
    Skip *int32 `uriparametername:"%24skip"`
    // Show only the first n items
    Top *int32 `uriparametername:"%24top"`
}
// MicrosoftGraphFindRoomListsFindRoomListsRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type MicrosoftGraphFindRoomListsFindRoomListsRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *MicrosoftGraphFindRoomListsFindRoomListsRequestBuilderGetQueryParameters
}
// NewMicrosoftGraphFindRoomListsFindRoomListsRequestBuilderInternal instantiates a new FindRoomListsRequestBuilder and sets the default values.
func NewMicrosoftGraphFindRoomListsFindRoomListsRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*MicrosoftGraphFindRoomListsFindRoomListsRequestBuilder) {
    m := &MicrosoftGraphFindRoomListsFindRoomListsRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/me/microsoft.graph.findRoomLists(){?%24top,%24skip,%24search,%24filter,%24count}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams
    m.requestAdapter = requestAdapter
    return m
}
// NewMicrosoftGraphFindRoomListsFindRoomListsRequestBuilder instantiates a new FindRoomListsRequestBuilder and sets the default values.
func NewMicrosoftGraphFindRoomListsFindRoomListsRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*MicrosoftGraphFindRoomListsFindRoomListsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewMicrosoftGraphFindRoomListsFindRoomListsRequestBuilderInternal(urlParams, requestAdapter)
}
// Get get the room lists defined in a tenant, as represented by their emailAddress objects. Tenants can organize meeting rooms into room lists. In this API, each meeting room and room list is represented by an emailAddress instance.You can get all the room lists in the tenant, get all the rooms in the tenant, or get all the rooms in a specific room list.
// [Find more info here]
// 
// [Find more info here]: https://docs.microsoft.com/graph/api/user-findroomlists?view=graph-rest-1.0
func (m *MicrosoftGraphFindRoomListsFindRoomListsRequestBuilder) Get(ctx context.Context, requestConfiguration *MicrosoftGraphFindRoomListsFindRoomListsRequestBuilderGetRequestConfiguration)(MicrosoftGraphFindRoomListsFindRoomListsResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.Send(ctx, requestInfo, CreateMicrosoftGraphFindRoomListsFindRoomListsResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(MicrosoftGraphFindRoomListsFindRoomListsResponseable), nil
}
// ToGetRequestInformation get the room lists defined in a tenant, as represented by their emailAddress objects. Tenants can organize meeting rooms into room lists. In this API, each meeting room and room list is represented by an emailAddress instance.You can get all the room lists in the tenant, get all the rooms in the tenant, or get all the rooms in a specific room list.
func (m *MicrosoftGraphFindRoomListsFindRoomListsRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *MicrosoftGraphFindRoomListsFindRoomListsRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET
    requestInfo.Headers.Add("Accept", "application/json")
    if requestConfiguration != nil {
        if requestConfiguration.QueryParameters != nil {
            requestInfo.AddQueryParameters(*(requestConfiguration.QueryParameters))
        }
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
