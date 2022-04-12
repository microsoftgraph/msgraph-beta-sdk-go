package eventmessagerequest

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i460d2d85f1298f703ab8e16cc109596a88c98f98d7caf830d187e2be0c58851f "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/mailfolders/item/childfolders/item/messages/item/eventmessagerequest/decline"
    i4fbcdf2144e0a6ffbd2526bbb85a41ca066e69666b134c995373d1822c184eb6 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/mailfolders/item/childfolders/item/messages/item/eventmessagerequest/tentativelyaccept"
    ib3fb4a21c906926cd084c40c6aec815393074c717161d3518075fcfaaedbcb5d "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/mailfolders/item/childfolders/item/messages/item/eventmessagerequest/accept"
)

// EventMessageRequestRequestBuilder builds and executes requests for operations under \users\{user-id}\mailFolders\{mailFolder-id}\childFolders\{mailFolder-id1}\messages\{message-id}\microsoft.graph.eventMessageRequest
type EventMessageRequestRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// Accept the accept property
func (m *EventMessageRequestRequestBuilder) Accept()(*ib3fb4a21c906926cd084c40c6aec815393074c717161d3518075fcfaaedbcb5d.AcceptRequestBuilder) {
    return ib3fb4a21c906926cd084c40c6aec815393074c717161d3518075fcfaaedbcb5d.NewAcceptRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// NewEventMessageRequestRequestBuilderInternal instantiates a new EventMessageRequestRequestBuilder and sets the default values.
func NewEventMessageRequestRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*EventMessageRequestRequestBuilder) {
    m := &EventMessageRequestRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/users/{user%2Did}/mailFolders/{mailFolder%2Did}/childFolders/{mailFolder%2Did1}/messages/{message%2Did}/microsoft.graph.eventMessageRequest";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewEventMessageRequestRequestBuilder instantiates a new EventMessageRequestRequestBuilder and sets the default values.
func NewEventMessageRequestRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*EventMessageRequestRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewEventMessageRequestRequestBuilderInternal(urlParams, requestAdapter)
}
// Decline the decline property
func (m *EventMessageRequestRequestBuilder) Decline()(*i460d2d85f1298f703ab8e16cc109596a88c98f98d7caf830d187e2be0c58851f.DeclineRequestBuilder) {
    return i460d2d85f1298f703ab8e16cc109596a88c98f98d7caf830d187e2be0c58851f.NewDeclineRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// TentativelyAccept the tentativelyAccept property
func (m *EventMessageRequestRequestBuilder) TentativelyAccept()(*i4fbcdf2144e0a6ffbd2526bbb85a41ca066e69666b134c995373d1822c184eb6.TentativelyAcceptRequestBuilder) {
    return i4fbcdf2144e0a6ffbd2526bbb85a41ca066e69666b134c995373d1822c184eb6.NewTentativelyAcceptRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
