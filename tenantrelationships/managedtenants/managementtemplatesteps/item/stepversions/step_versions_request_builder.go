package stepversions

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    iff673fb08c79c236f1dc447f618c13b3bc13d45c7a2b1408fd2383407bf599d6 "github.com/microsoftgraph/msgraph-beta-sdk-go/tenantrelationships/managedtenants/managementtemplatesteps/item/stepversions/ref"
)

// StepVersionsRequestBuilder builds and executes requests for operations under \tenantRelationships\managedTenants\managementTemplateSteps\{managementTemplateStep-id}\stepVersions
type StepVersionsRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// StepVersionsRequestBuilderGetOptions options for Get
type StepVersionsRequestBuilderGetOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Request query parameters
    Q *StepVersionsRequestBuilderGetQueryParameters;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// StepVersionsRequestBuilderGetQueryParameters get stepVersions from tenantRelationships
type StepVersionsRequestBuilderGetQueryParameters struct {
    // Include count of items
    Count *bool;
    // Expand related entities
    Expand []string;
    // Filter items by property values
    Filter *string;
    // Order items by property values
    Orderby []string;
    // Search items by search phrases
    Search *string;
    // Select properties to be returned
    Select []string;
    // Skip the first n items
    Skip *int32;
    // Show only the first n items
    Top *int32;
}
// NewStepVersionsRequestBuilderInternal instantiates a new StepVersionsRequestBuilder and sets the default values.
func NewStepVersionsRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*StepVersionsRequestBuilder) {
    m := &StepVersionsRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/tenantRelationships/managedTenants/managementTemplateSteps/{managementTemplateStep_id}/stepVersions{?top,skip,search,filter,count,orderby,select,expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = pathParameters;
    m.requestAdapter = requestAdapter;
    return m
}
// NewStepVersionsRequestBuilder instantiates a new StepVersionsRequestBuilder and sets the default values.
func NewStepVersionsRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*StepVersionsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewStepVersionsRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateGetRequestInformation get stepVersions from tenantRelationships
func (m *StepVersionsRequestBuilder) CreateGetRequestInformation(options *StepVersionsRequestBuilderGetOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// Get get stepVersions from tenantRelationships
func (m *StepVersionsRequestBuilder) Get(options *StepVersionsRequestBuilderGetOptions)(*StepVersionsResponse, error) {
    requestInfo, err := m.CreateGetRequestInformation(options);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.SendAsync(*requestInfo, func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewStepVersionsResponse() }, nil)
    if err != nil {
        return nil, err
    }
    return res.(*StepVersionsResponse), nil
}
func (m *StepVersionsRequestBuilder) Ref()(*iff673fb08c79c236f1dc447f618c13b3bc13d45c7a2b1408fd2383407bf599d6.RefRequestBuilder) {
    return iff673fb08c79c236f1dc447f618c13b3bc13d45c7a2b1408fd2383407bf599d6.NewRefRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}