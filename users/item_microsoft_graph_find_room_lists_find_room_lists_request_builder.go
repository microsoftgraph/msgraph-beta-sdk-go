package users

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// ItemMicrosoftGraphFindRoomListsFindRoomListsRequestBuilder provides operations to call the findRoomLists method.
type ItemMicrosoftGraphFindRoomListsFindRoomListsRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// ItemMicrosoftGraphFindRoomListsFindRoomListsRequestBuilderGetQueryParameters get the room lists defined in a tenant, as represented by their emailAddress objects. Tenants can organize meeting rooms into room lists. In this API, each meeting room and room list is represented by an emailAddress instance.You can get all the room lists in the tenant, get all the rooms in the tenant, or get all the rooms in a specific room list.
type ItemMicrosoftGraphFindRoomListsFindRoomListsRequestBuilderGetQueryParameters struct {
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
// ItemMicrosoftGraphFindRoomListsFindRoomListsRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemMicrosoftGraphFindRoomListsFindRoomListsRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ItemMicrosoftGraphFindRoomListsFindRoomListsRequestBuilderGetQueryParameters
}
// NewItemMicrosoftGraphFindRoomListsFindRoomListsRequestBuilderInternal instantiates a new FindRoomListsRequestBuilder and sets the default values.
func NewItemMicrosoftGraphFindRoomListsFindRoomListsRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemMicrosoftGraphFindRoomListsFindRoomListsRequestBuilder) {
    m := &ItemMicrosoftGraphFindRoomListsFindRoomListsRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/users/{user%2Did}/microsoft.graph.findRoomLists(){?%24top,%24skip,%24search,%24filter,%24count}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams
    m.requestAdapter = requestAdapter
    return m
}
// NewItemMicrosoftGraphFindRoomListsFindRoomListsRequestBuilder instantiates a new FindRoomListsRequestBuilder and sets the default values.
func NewItemMicrosoftGraphFindRoomListsFindRoomListsRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemMicrosoftGraphFindRoomListsFindRoomListsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemMicrosoftGraphFindRoomListsFindRoomListsRequestBuilderInternal(urlParams, requestAdapter)
}
// Get get the room lists defined in a tenant, as represented by their emailAddress objects. Tenants can organize meeting rooms into room lists. In this API, each meeting room and room list is represented by an emailAddress instance.You can get all the room lists in the tenant, get all the rooms in the tenant, or get all the rooms in a specific room list.
// [Find more info here]
// 
// [Find more info here]: https://docs.microsoft.com/graph/api/user-findroomlists?view=graph-rest-1.0
func (m *ItemMicrosoftGraphFindRoomListsFindRoomListsRequestBuilder) Get(ctx context.Context, requestConfiguration *ItemMicrosoftGraphFindRoomListsFindRoomListsRequestBuilderGetRequestConfiguration)(ItemMicrosoftGraphFindRoomListsFindRoomListsResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.Send(ctx, requestInfo, CreateItemMicrosoftGraphFindRoomListsFindRoomListsResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ItemMicrosoftGraphFindRoomListsFindRoomListsResponseable), nil
}
// ToGetRequestInformation get the room lists defined in a tenant, as represented by their emailAddress objects. Tenants can organize meeting rooms into room lists. In this API, each meeting room and room list is represented by an emailAddress instance.You can get all the room lists in the tenant, get all the rooms in the tenant, or get all the rooms in a specific room list.
func (m *ItemMicrosoftGraphFindRoomListsFindRoomListsRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ItemMicrosoftGraphFindRoomListsFindRoomListsRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
