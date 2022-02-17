package item

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph"
    i6027adaace07f5ce2884ddb6c6f42f4f657337224011aed3f67e4f6d702b65d5 "github.com/microsoftgraph/msgraph-beta-sdk-go/chats/item/installedapps/item/upgrade"
    i7c9660d87d4841efba0bf04fafd25c0e1cfcebaac5121db728184507e522b5d8 "github.com/microsoftgraph/msgraph-beta-sdk-go/chats/item/installedapps/item/teamsappdefinition"
    if6be6a563953e577b7efcecb779024594ae836463f5524eb49aa6d5e1030a9c1 "github.com/microsoftgraph/msgraph-beta-sdk-go/chats/item/installedapps/item/teamsapp"
)

// TeamsAppInstallationRequestBuilder builds and executes requests for operations under \chats\{chat-id}\installedApps\{teamsAppInstallation-id}
type TeamsAppInstallationRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// TeamsAppInstallationRequestBuilderDeleteOptions options for Delete
type TeamsAppInstallationRequestBuilderDeleteOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// TeamsAppInstallationRequestBuilderGetOptions options for Get
type TeamsAppInstallationRequestBuilderGetOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Request query parameters
    Q *TeamsAppInstallationRequestBuilderGetQueryParameters;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// TeamsAppInstallationRequestBuilderGetQueryParameters a collection of all the apps in the chat. Nullable.
type TeamsAppInstallationRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string;
    // Select properties to be returned
    Select []string;
}
// TeamsAppInstallationRequestBuilderPatchOptions options for Patch
type TeamsAppInstallationRequestBuilderPatchOptions struct {
    // 
    Body *i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.TeamsAppInstallation;
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// NewTeamsAppInstallationRequestBuilderInternal instantiates a new TeamsAppInstallationRequestBuilder and sets the default values.
func NewTeamsAppInstallationRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*TeamsAppInstallationRequestBuilder) {
    m := &TeamsAppInstallationRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/chats/{chat_id}/installedApps/{teamsAppInstallation_id}{?select,expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = pathParameters;
    m.requestAdapter = requestAdapter;
    return m
}
// NewTeamsAppInstallationRequestBuilder instantiates a new TeamsAppInstallationRequestBuilder and sets the default values.
func NewTeamsAppInstallationRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*TeamsAppInstallationRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewTeamsAppInstallationRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateDeleteRequestInformation a collection of all the apps in the chat. Nullable.
func (m *TeamsAppInstallationRequestBuilder) CreateDeleteRequestInformation(options *TeamsAppInstallationRequestBuilderDeleteOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// CreateGetRequestInformation a collection of all the apps in the chat. Nullable.
func (m *TeamsAppInstallationRequestBuilder) CreateGetRequestInformation(options *TeamsAppInstallationRequestBuilderGetOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// CreatePatchRequestInformation a collection of all the apps in the chat. Nullable.
func (m *TeamsAppInstallationRequestBuilder) CreatePatchRequestInformation(options *TeamsAppInstallationRequestBuilderPatchOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// Delete a collection of all the apps in the chat. Nullable.
func (m *TeamsAppInstallationRequestBuilder) Delete(options *TeamsAppInstallationRequestBuilderDeleteOptions)(error) {
    requestInfo, err := m.CreateDeleteRequestInformation(options);
    if err != nil {
        return err
    }
    err = m.requestAdapter.SendNoContentAsync(*requestInfo, nil, nil)
    if err != nil {
        return err
    }
    return nil
}
// Get a collection of all the apps in the chat. Nullable.
func (m *TeamsAppInstallationRequestBuilder) Get(options *TeamsAppInstallationRequestBuilderGetOptions)(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.TeamsAppInstallation, error) {
    requestInfo, err := m.CreateGetRequestInformation(options);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.SendAsync(*requestInfo, func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.NewTeamsAppInstallation() }, nil, nil)
    if err != nil {
        return nil, err
    }
    return res.(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.TeamsAppInstallation), nil
}
// Patch a collection of all the apps in the chat. Nullable.
func (m *TeamsAppInstallationRequestBuilder) Patch(options *TeamsAppInstallationRequestBuilderPatchOptions)(error) {
    requestInfo, err := m.CreatePatchRequestInformation(options);
    if err != nil {
        return err
    }
    err = m.requestAdapter.SendNoContentAsync(*requestInfo, nil, nil)
    if err != nil {
        return err
    }
    return nil
}
func (m *TeamsAppInstallationRequestBuilder) TeamsApp()(*if6be6a563953e577b7efcecb779024594ae836463f5524eb49aa6d5e1030a9c1.TeamsAppRequestBuilder) {
    return if6be6a563953e577b7efcecb779024594ae836463f5524eb49aa6d5e1030a9c1.NewTeamsAppRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *TeamsAppInstallationRequestBuilder) TeamsAppDefinition()(*i7c9660d87d4841efba0bf04fafd25c0e1cfcebaac5121db728184507e522b5d8.TeamsAppDefinitionRequestBuilder) {
    return i7c9660d87d4841efba0bf04fafd25c0e1cfcebaac5121db728184507e522b5d8.NewTeamsAppDefinitionRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *TeamsAppInstallationRequestBuilder) Upgrade()(*i6027adaace07f5ce2884ddb6c6f42f4f657337224011aed3f67e4f6d702b65d5.UpgradeRequestBuilder) {
    return i6027adaace07f5ce2884ddb6c6f42f4f657337224011aed3f67e4f6d702b65d5.NewUpgradeRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
