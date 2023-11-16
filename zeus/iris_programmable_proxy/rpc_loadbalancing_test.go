package iris_programmable_proxy

import (
	"context"
	"encoding/json"
	"fmt"
	"math/big"
	"sync"
	"time"

	"github.com/go-resty/resty/v2"
	web3_actions "github.com/zeus-fyi/gochain/web3/client"
	hestia_req_types "github.com/zeus-fyi/zeus/pkg/hestia/client/req_types"
	iris_programmable_proxy_v1_beta "github.com/zeus-fyi/zeus/zeus/iris_programmable_proxy/v1beta"
	resty_base "github.com/zeus-fyi/zeus/zeus/z_client/base"
)

func (t *IrisConfigTestSuite) TestRPCLoadBalancing() {
	routeGroup := "etherum-mainnet-sample"
	//path := fmt.Sprintf("https://iris.zeus.fyi/v1/router")
	path := fmt.Sprintf("http://localhost:8080/v1/router")

	web3a := web3_actions.NewWeb3ActionsClient(path)
	web3a.AddRoutingGroupHeader(routeGroup)
	web3a.AddBearerToken(t.IrisClient.Token)
	web3a.Dial()
	reqCount := 4
	defer web3a.Close()
	for i := 0; i < reqCount; i++ {
		start := time.Now()
		resp, err := web3a.C.BlockByNumber(context.Background(), nil)
		end := time.Now()
		fmt.Println("time taken: ", end.Sub(start).Milliseconds())
		t.NoError(err)
		t.NotNil(resp)
	}
}

func (t *IrisConfigTestSuite) TestDirectEndpoint() {
	t.Require().NotEmpty(t.Tc.QuickNodeEndpoint)
	web3a := web3_actions.NewWeb3ActionsClient(t.Tc.QuickNodeEndpoint)
	web3a.Dial()
	reqCount := 4
	fmt.Println("RPC Requests")

	defer web3a.Close()
	for i := 0; i < reqCount; i++ {
		start := time.Now()
		resp, err := web3a.C.BlockByNumber(context.Background(), nil)
		end := time.Now()
		fmt.Println("time taken: ", end.Sub(start).Milliseconds())
		t.NoError(err)
		t.NotNil(resp)
	}

	fmt.Println("POST Requests")
	r := resty.New()
	payload := `{
		"jsonrpc": "2.0",
		"method": "eth_getBlockByNumber",
		"params": ["latest", true],
		"id": 1
	}`
	for i := 0; i < reqCount; i++ {
		start := time.Now()
		resp, err := r.R().
			SetBody(payload).
			Post(t.Tc.QuickNodeEndpoint)
		t.Require().NoError(err)
		t.Require().NotNil(resp)
		end := time.Now()
		fmt.Println("time taken: ", end.Sub(start).Milliseconds())
	}

	fmt.Println("POST Requests + Decode")
	m := make(map[string]interface{})
	for i := 0; i < reqCount; i++ {
		start := time.Now()
		resp, err := r.R().
			SetBody(payload).
			SetResult(&m).
			Post(t.Tc.QuickNodeEndpoint)
		t.Require().NoError(err)
		t.Require().NotNil(resp)
		end := time.Now()
		fmt.Println("time taken: ", end.Sub(start).Milliseconds())
	}
	fmt.Println("Localhost Requests")
	routeGroup := "etherum-mainnet-sample"
	path := fmt.Sprintf("http://localhost:8080/v1/router")

	web3a = web3_actions.NewWeb3ActionsClient(path)
	web3a.AddRoutingGroupHeader(routeGroup)
	web3a.AddBearerToken(t.Tc.Bearer)
	web3a.Dial()
	defer web3a.Close()
	for i := 0; i < reqCount; i++ {
		start := time.Now()
		resp, err := web3a.C.BlockNumber(context.Background())
		t.NoError(err)
		end := time.Now()
		fmt.Println("time taken: ", end.Sub(start).Milliseconds())
		t.NotNil(resp)
	}
}

func (t *IrisConfigTestSuite) TestParallelRPCLoadBalancing() {
	routeGroup := "ethereum-mainnet"
	path := fmt.Sprintf("https://iris.zeus.fyi/v1/router")

	reqCountInParallel := 100
	// Define a channel for controlling the number of concurrent requests
	sem := make(chan bool, reqCountInParallel)

	// Define an error channel to catch errors from goroutines
	errCh := make(chan error, reqCountInParallel)

	var wg sync.WaitGroup

	web3a := web3_actions.NewWeb3ActionsClient(path)
	web3a.AddRoutingGroupHeader(routeGroup)
	web3a.Headers[iris_programmable_proxy_v1_beta.LoadBalancingStrategy] = iris_programmable_proxy_v1_beta.Adaptive
	web3a.Headers[iris_programmable_proxy_v1_beta.AdaptiveLoadBalancingKey] = iris_programmable_proxy_v1_beta.JsonRpcAdaptiveMetrics
	web3a.AddBearerToken(t.BearerToken)
	web3a.Dial()
	defer web3a.Close()

	bn := 17500000
	offset := 0
	for i := 0; i < reqCountInParallel*100; i++ {
		// Acquire a semaphore
		sem <- true

		wg.Add(1)
		go func(offset int) {
			defer func() {
				// Release the semaphore
				<-sem
				wg.Done()
			}()

			resp, err := web3a.C.BlockByNumber(context.Background(), new(big.Int).SetInt64(int64(bn+offset)))
			if err != nil {
				errCh <- err
				return
			}
			t.NotNil(resp.Body())

			b, err := json.Marshal(*resp.Body())
			t.Nil(err)

			// uncomment to see full block body printed
			//fmt.Println(string(b))

			// resp body portion of the block ~125KB, total block size ~200-300KB
			fmt.Printf("Size of resp: %.2f KB\n", float64(len(string(b))/1024.0))
		}(offset)
	}
	offset++
	// Wait for all the goroutines to complete
	wg.Wait()
	close(errCh)

	// Check for any errors in the error channel
	for err := range errCh {
		t.NoError(err)
	}
}

func (t *IrisConfigTestSuite) TestGetLoadBalancing() {
	routeGroup := "olympus"
	path := fmt.Sprintf("https://iris.zeus.fyi/v1/router")
	//path = fmt.Sprintf("http://localhost:8080/v1/router")
	routeOne := "https://hestia.zeus.fyi/health"
	routeTwo := "https://iris.zeus.fyi/health"

	err := t.IrisClientProd.UpdateRoutingGroupEndpoints(ctx, hestia_req_types.IrisOrgGroupRoutesRequest{
		GroupName: routeGroup,
		Routes: []string{
			"https://hestia.zeus.fyi/health",
			"https://iris.zeus.fyi/health",
		},
	})
	t.Nil(err)
	r := resty_base.GetBaseRestyClient(path, t.IrisClientProd.Token)
	r.SetRoutingGroupHeader(routeGroup)
	reqCount := 4
	m := make(map[string]int)
	m[routeOne] = 0
	m[routeTwo] = 0
	for i := 0; i < reqCount; i++ {
		resp1, err1 := r.R().Get(path)
		t.NoError(err1)
		t.NotNil(resp1)

		selectedHeader := resp1.Header().Get("X-Selected-Route")
		t.NotEmpty(selectedHeader)
		fmt.Println(selectedHeader)

		m[selectedHeader]++
	}

	t.GreaterOrEqual(m[routeOne], 1)
	t.GreaterOrEqual(m[routeTwo], 1)

	rr := hestia_req_types.IrisOrgGroupRoutesRequest{
		Routes: []string{routeOne},
	}
	resp, err := t.IrisClientProd.DeleteRoutingEndpoints(ctx, rr)
	t.NoError(err)
	t.NotNil(resp)

	m[routeOne] = 0
	m[routeTwo] = 0

	// gives time for the routing group to update
	time.Sleep(5 * time.Second)
	for i := 0; i < reqCount; i++ {
		resp2, err2 := r.R().Get(path)
		t.NoError(err2)
		t.NotNil(resp)

		selectedHeader := resp2.Header().Get("X-Selected-Route")
		t.NotEmpty(selectedHeader)
		fmt.Println(selectedHeader)

		m[selectedHeader]++
	}
	t.Assert().Zero(m[routeOne])
	t.Assert().Equal(reqCount, m[routeTwo])
}
