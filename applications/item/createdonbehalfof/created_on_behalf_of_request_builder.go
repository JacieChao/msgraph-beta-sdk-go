package createdonbehalfof

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph"
    i43526d35583dd9080b4876a8e2b3266a57f601e7f67dd2aca6820e77a95d1611 "github.com/microsoftgraph/msgraph-beta-sdk-go/applications/item/createdonbehalfof/ref"
)

// createdOnBehalfOfRequestBuilder builds and executes requests for operations under \applications\{application-id}\createdOnBehalfOf
type CreatedOnBehalfOfRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// CreatedOnBehalfOfRequestBuilderGetOptions options for Get
type CreatedOnBehalfOfRequestBuilderGetOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Request query parameters
    Q *CreatedOnBehalfOfRequestBuilderGetQueryParameters;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// createdOnBehalfOfRequestBuilderGetQueryParameters read-only.
type CreatedOnBehalfOfRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string;
    // Select properties to be returned
    Select_escaped []string;
}
// NewCreatedOnBehalfOfRequestBuilderInternal instantiates a new CreatedOnBehalfOfRequestBuilder and sets the default values.
func NewCreatedOnBehalfOfRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*CreatedOnBehalfOfRequestBuilder) {
    m := &CreatedOnBehalfOfRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/applications/{application_id}/createdOnBehalfOf{?select,expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = pathParameters;
    m.requestAdapter = requestAdapter;
    return m
}
// NewCreatedOnBehalfOfRequestBuilder instantiates a new CreatedOnBehalfOfRequestBuilder and sets the default values.
func NewCreatedOnBehalfOfRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*CreatedOnBehalfOfRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewCreatedOnBehalfOfRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateGetRequestInformation read-only.
func (m *CreatedOnBehalfOfRequestBuilder) CreateGetRequestInformation(options *CreatedOnBehalfOfRequestBuilderGetOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// Get read-only.
func (m *CreatedOnBehalfOfRequestBuilder) Get(options *CreatedOnBehalfOfRequestBuilderGetOptions)(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.DirectoryObject, error) {
    requestInfo, err := m.CreateGetRequestInformation(options);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.SendAsync(*requestInfo, func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.NewDirectoryObject() }, nil)
    if err != nil {
        return nil, err
    }
    return res.(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.DirectoryObject), nil
}
func (m *CreatedOnBehalfOfRequestBuilder) Ref()(*i43526d35583dd9080b4876a8e2b3266a57f601e7f67dd2aca6820e77a95d1611.RefRequestBuilder) {
    return i43526d35583dd9080b4876a8e2b3266a57f601e7f67dd2aca6820e77a95d1611.NewRefRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
