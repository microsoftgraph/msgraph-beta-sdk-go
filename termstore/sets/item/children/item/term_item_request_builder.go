package item

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i11bbafbd56c37d7251e6346f4524f5f88f83f0c6ce462b8a4d27e4a6d969d4c3 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph/termstore"
    i428a28d14ab585560ab266716b214a45f45f18468b52fdb0f932c81a7f9706e4 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph/odataerrors"
    i259c4974ad05369136c91a47b0bb33bb51caf6c65426de4ff720a1fd6e261ab3 "github.com/microsoftgraph/msgraph-beta-sdk-go/termstore/sets/item/children/item/relations"
    i75e7c2ca91afeb714ccd80a1f59aba614c728edc1b8d97bf9ea98821a5d24a0b "github.com/microsoftgraph/msgraph-beta-sdk-go/termstore/sets/item/children/item/children"
    ie19ca3393f56e45ebb87ac0b256d0a128e831cee1559ed8f5611b26e6df1de3c "github.com/microsoftgraph/msgraph-beta-sdk-go/termstore/sets/item/children/item/set"
    i60d4636b3dc73e96047b220d2a344ac838391e14ed10d97c599064719c02861f "github.com/microsoftgraph/msgraph-beta-sdk-go/termstore/sets/item/children/item/children/item"
    i8c41d0c998190697de4b3e9f8db43880fe77b5f1eacab2484f48b7c66d6d6836 "github.com/microsoftgraph/msgraph-beta-sdk-go/termstore/sets/item/children/item/relations/item"
)

// TermItemRequestBuilder provides operations to manage the children property of the microsoft.graph.termStore.set entity.
type TermItemRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// TermItemRequestBuilderDeleteOptions options for Delete
type TermItemRequestBuilderDeleteOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// TermItemRequestBuilderGetOptions options for Get
type TermItemRequestBuilderGetOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Request query parameters
    Q *TermItemRequestBuilderGetQueryParameters;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// TermItemRequestBuilderGetQueryParameters children terms of set in term [store].
type TermItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string;
    // Select properties to be returned
    Select []string;
}
// TermItemRequestBuilderPatchOptions options for Patch
type TermItemRequestBuilderPatchOptions struct {
    // 
    Body i11bbafbd56c37d7251e6346f4524f5f88f83f0c6ce462b8a4d27e4a6d969d4c3.Termable;
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
func (m *TermItemRequestBuilder) Children()(*i75e7c2ca91afeb714ccd80a1f59aba614c728edc1b8d97bf9ea98821a5d24a0b.ChildrenRequestBuilder) {
    return i75e7c2ca91afeb714ccd80a1f59aba614c728edc1b8d97bf9ea98821a5d24a0b.NewChildrenRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ChildrenById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.termStore.sets.item.children.item.children.item collection
func (m *TermItemRequestBuilder) ChildrenById(id string)(*i60d4636b3dc73e96047b220d2a344ac838391e14ed10d97c599064719c02861f.TermItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["term_id1"] = id
    }
    return i60d4636b3dc73e96047b220d2a344ac838391e14ed10d97c599064719c02861f.NewTermItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// NewTermItemRequestBuilderInternal instantiates a new TermItemRequestBuilder and sets the default values.
func NewTermItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*TermItemRequestBuilder) {
    m := &TermItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/termStore/sets/{set_id}/children/{term_id}{?select,expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewTermItemRequestBuilder instantiates a new TermItemRequestBuilder and sets the default values.
func NewTermItemRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*TermItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewTermItemRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateDeleteRequestInformation delete navigation property children for termStore
func (m *TermItemRequestBuilder) CreateDeleteRequestInformation(options *TermItemRequestBuilderDeleteOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
    requestInfo := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.DELETE
    if options != nil && options.H != nil {
        requestInfo.Headers = options.H
    }
    if options != nil && len(options.O) != 0 {
        err := requestInfo.AddRequestOptions(options.O...)
        if err != nil {
            return nil, err
        }
    }
    return requestInfo, nil
}
// CreateGetRequestInformation children terms of set in term [store].
func (m *TermItemRequestBuilder) CreateGetRequestInformation(options *TermItemRequestBuilderGetOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
    requestInfo := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.GET
    if options != nil && options.Q != nil {
        requestInfo.AddQueryParameters(*(options.Q))
    }
    if options != nil && options.H != nil {
        requestInfo.Headers = options.H
    }
    if options != nil && len(options.O) != 0 {
        err := requestInfo.AddRequestOptions(options.O...)
        if err != nil {
            return nil, err
        }
    }
    return requestInfo, nil
}
// CreatePatchRequestInformation update the navigation property children in termStore
func (m *TermItemRequestBuilder) CreatePatchRequestInformation(options *TermItemRequestBuilderPatchOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
    requestInfo := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.PATCH
    requestInfo.SetContentFromParsable(m.requestAdapter, "application/json", options.Body)
    if options != nil && options.H != nil {
        requestInfo.Headers = options.H
    }
    if options != nil && len(options.O) != 0 {
        err := requestInfo.AddRequestOptions(options.O...)
        if err != nil {
            return nil, err
        }
    }
    return requestInfo, nil
}
// Delete delete navigation property children for termStore
func (m *TermItemRequestBuilder) Delete(options *TermItemRequestBuilderDeleteOptions)(error) {
    requestInfo, err := m.CreateDeleteRequestInformation(options);
    if err != nil {
        return err
    }
    errorMapping := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ErrorMappings {
        "4XX": i428a28d14ab585560ab266716b214a45f45f18468b52fdb0f932c81a7f9706e4.CreateODataErrorFromDiscriminatorValue,
        "5XX": i428a28d14ab585560ab266716b214a45f45f18468b52fdb0f932c81a7f9706e4.CreateODataErrorFromDiscriminatorValue,
    }
    err = m.requestAdapter.SendNoContentAsync(requestInfo, nil, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
// Get children terms of set in term [store].
func (m *TermItemRequestBuilder) Get(options *TermItemRequestBuilderGetOptions)(i11bbafbd56c37d7251e6346f4524f5f88f83f0c6ce462b8a4d27e4a6d969d4c3.Termable, error) {
    requestInfo, err := m.CreateGetRequestInformation(options);
    if err != nil {
        return nil, err
    }
    errorMapping := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ErrorMappings {
        "4XX": i428a28d14ab585560ab266716b214a45f45f18468b52fdb0f932c81a7f9706e4.CreateODataErrorFromDiscriminatorValue,
        "5XX": i428a28d14ab585560ab266716b214a45f45f18468b52fdb0f932c81a7f9706e4.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(requestInfo, i11bbafbd56c37d7251e6346f4524f5f88f83f0c6ce462b8a4d27e4a6d969d4c3.CreateTermFromDiscriminatorValue, nil, errorMapping)
    if err != nil {
        return nil, err
    }
    return res.(i11bbafbd56c37d7251e6346f4524f5f88f83f0c6ce462b8a4d27e4a6d969d4c3.Termable), nil
}
// Patch update the navigation property children in termStore
func (m *TermItemRequestBuilder) Patch(options *TermItemRequestBuilderPatchOptions)(error) {
    requestInfo, err := m.CreatePatchRequestInformation(options);
    if err != nil {
        return err
    }
    errorMapping := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ErrorMappings {
        "4XX": i428a28d14ab585560ab266716b214a45f45f18468b52fdb0f932c81a7f9706e4.CreateODataErrorFromDiscriminatorValue,
        "5XX": i428a28d14ab585560ab266716b214a45f45f18468b52fdb0f932c81a7f9706e4.CreateODataErrorFromDiscriminatorValue,
    }
    err = m.requestAdapter.SendNoContentAsync(requestInfo, nil, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
func (m *TermItemRequestBuilder) Relations()(*i259c4974ad05369136c91a47b0bb33bb51caf6c65426de4ff720a1fd6e261ab3.RelationsRequestBuilder) {
    return i259c4974ad05369136c91a47b0bb33bb51caf6c65426de4ff720a1fd6e261ab3.NewRelationsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// RelationsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.termStore.sets.item.children.item.relations.item collection
func (m *TermItemRequestBuilder) RelationsById(id string)(*i8c41d0c998190697de4b3e9f8db43880fe77b5f1eacab2484f48b7c66d6d6836.RelationItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["relation_id"] = id
    }
    return i8c41d0c998190697de4b3e9f8db43880fe77b5f1eacab2484f48b7c66d6d6836.NewRelationItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *TermItemRequestBuilder) Set()(*ie19ca3393f56e45ebb87ac0b256d0a128e831cee1559ed8f5611b26e6df1de3c.SetRequestBuilder) {
    return ie19ca3393f56e45ebb87ac0b256d0a128e831cee1559ed8f5611b26e6df1de3c.NewSetRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
