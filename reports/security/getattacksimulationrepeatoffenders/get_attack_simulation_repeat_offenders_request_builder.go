package getattacksimulationrepeatoffenders

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// GetAttackSimulationRepeatOffendersRequestBuilder provides operations to call the getAttackSimulationRepeatOffenders method.
type GetAttackSimulationRepeatOffendersRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// GetAttackSimulationRepeatOffendersRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type GetAttackSimulationRepeatOffendersRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewGetAttackSimulationRepeatOffendersRequestBuilderInternal instantiates a new GetAttackSimulationRepeatOffendersRequestBuilder and sets the default values.
func NewGetAttackSimulationRepeatOffendersRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*GetAttackSimulationRepeatOffendersRequestBuilder) {
    m := &GetAttackSimulationRepeatOffendersRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/reports/security/microsoft.graph.getAttackSimulationRepeatOffenders()";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewGetAttackSimulationRepeatOffendersRequestBuilder instantiates a new GetAttackSimulationRepeatOffendersRequestBuilder and sets the default values.
func NewGetAttackSimulationRepeatOffendersRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*GetAttackSimulationRepeatOffendersRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewGetAttackSimulationRepeatOffendersRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateGetRequestInformation invoke function getAttackSimulationRepeatOffenders
func (m *GetAttackSimulationRepeatOffendersRequestBuilder) CreateGetRequestInformation()(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    return m.CreateGetRequestInformationWithRequestConfiguration(nil);
}
// CreateGetRequestInformationWithRequestConfiguration invoke function getAttackSimulationRepeatOffenders
func (m *GetAttackSimulationRepeatOffendersRequestBuilder) CreateGetRequestInformationWithRequestConfiguration(requestConfiguration *GetAttackSimulationRepeatOffendersRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET
    requestInfo.Headers["Accept"] = "application/json"
    if requestConfiguration != nil {
        requestInfo.AddRequestHeaders(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// Get invoke function getAttackSimulationRepeatOffenders
func (m *GetAttackSimulationRepeatOffendersRequestBuilder) Get(ctx context.Context, requestConfiguration *GetAttackSimulationRepeatOffendersRequestBuilderGetRequestConfiguration)(GetAttackSimulationRepeatOffendersResponseable, error) {
    requestInfo, err := m.CreateGetRequestInformationWithRequestConfiguration(requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(ctx, requestInfo, CreateGetAttackSimulationRepeatOffendersResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(GetAttackSimulationRepeatOffendersResponseable), nil
}