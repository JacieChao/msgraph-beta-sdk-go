package reports

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// GetAzureADApplicationSignInSummaryWithPeriodRequestBuilder provides operations to call the getAzureADApplicationSignInSummary method.
type GetAzureADApplicationSignInSummaryWithPeriodRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// GetAzureADApplicationSignInSummaryWithPeriodRequestBuilderGetQueryParameters invoke function getAzureADApplicationSignInSummary
type GetAzureADApplicationSignInSummaryWithPeriodRequestBuilderGetQueryParameters struct {
    // Include count of items
    Count *bool `uriparametername:"%24count"`
    // Filter items by property values
    Filter *string `uriparametername:"%24filter"`
    // Order items by property values
    Orderby []string `uriparametername:"%24orderby"`
    // Search items by search phrases
    Search *string `uriparametername:"%24search"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
    // Skip the first n items
    Skip *int32 `uriparametername:"%24skip"`
    // Show only the first n items
    Top *int32 `uriparametername:"%24top"`
}
// GetAzureADApplicationSignInSummaryWithPeriodRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type GetAzureADApplicationSignInSummaryWithPeriodRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *GetAzureADApplicationSignInSummaryWithPeriodRequestBuilderGetQueryParameters
}
// NewGetAzureADApplicationSignInSummaryWithPeriodRequestBuilderInternal instantiates a new GetAzureADApplicationSignInSummaryWithPeriodRequestBuilder and sets the default values.
func NewGetAzureADApplicationSignInSummaryWithPeriodRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter, period *string)(*GetAzureADApplicationSignInSummaryWithPeriodRequestBuilder) {
    m := &GetAzureADApplicationSignInSummaryWithPeriodRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/reports/getAzureADApplicationSignInSummary(period='{period}'){?%24top,%24skip,%24search,%24filter,%24count,%24select,%24orderby}", pathParameters),
    }
    if period != nil {
        m.BaseRequestBuilder.PathParameters["period"] = *period
    }
    return m
}
// NewGetAzureADApplicationSignInSummaryWithPeriodRequestBuilder instantiates a new GetAzureADApplicationSignInSummaryWithPeriodRequestBuilder and sets the default values.
func NewGetAzureADApplicationSignInSummaryWithPeriodRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*GetAzureADApplicationSignInSummaryWithPeriodRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewGetAzureADApplicationSignInSummaryWithPeriodRequestBuilderInternal(urlParams, requestAdapter, nil)
}
// Get invoke function getAzureADApplicationSignInSummary
// Deprecated: This method is obsolete. Use GetAsGetAzureADApplicationSignInSummaryWithPeriodGetResponse instead.
func (m *GetAzureADApplicationSignInSummaryWithPeriodRequestBuilder) Get(ctx context.Context, requestConfiguration *GetAzureADApplicationSignInSummaryWithPeriodRequestBuilderGetRequestConfiguration)(GetAzureADApplicationSignInSummaryWithPeriodResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateGetAzureADApplicationSignInSummaryWithPeriodResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(GetAzureADApplicationSignInSummaryWithPeriodResponseable), nil
}
// GetAsGetAzureADApplicationSignInSummaryWithPeriodGetResponse invoke function getAzureADApplicationSignInSummary
func (m *GetAzureADApplicationSignInSummaryWithPeriodRequestBuilder) GetAsGetAzureADApplicationSignInSummaryWithPeriodGetResponse(ctx context.Context, requestConfiguration *GetAzureADApplicationSignInSummaryWithPeriodRequestBuilderGetRequestConfiguration)(GetAzureADApplicationSignInSummaryWithPeriodGetResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateGetAzureADApplicationSignInSummaryWithPeriodGetResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(GetAzureADApplicationSignInSummaryWithPeriodGetResponseable), nil
}
// ToGetRequestInformation invoke function getAzureADApplicationSignInSummary
func (m *GetAzureADApplicationSignInSummaryWithPeriodRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *GetAzureADApplicationSignInSummaryWithPeriodRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    if requestConfiguration != nil {
        if requestConfiguration.QueryParameters != nil {
            requestInfo.AddQueryParameters(*(requestConfiguration.QueryParameters))
        }
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.UrlTemplate = m.BaseRequestBuilder.UrlTemplate
    requestInfo.PathParameters = m.BaseRequestBuilder.PathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET
    requestInfo.Headers.TryAdd("Accept", "application/json;q=1")
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
func (m *GetAzureADApplicationSignInSummaryWithPeriodRequestBuilder) WithUrl(rawUrl string)(*GetAzureADApplicationSignInSummaryWithPeriodRequestBuilder) {
    return NewGetAzureADApplicationSignInSummaryWithPeriodRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
