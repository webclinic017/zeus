package hestia_client

import (
	"context"
	"fmt"

	"github.com/rs/zerolog/log"
	hestia_endpoints "github.com/zeus-fyi/zeus/pkg/hestia/client/endpoints"
	hestia_resp_types "github.com/zeus-fyi/zeus/pkg/hestia/client/resp_types"
)

func (h *Hestia) CreateIrisRoutes(ctx context.Context, rr any) (hestia_resp_types.Response, error) {
	h.PrintReqJson(rr)
	respJson := hestia_resp_types.Response{}
	resp, err := h.R().
		SetBody(rr).
		SetResult(&respJson).
		Post(hestia_endpoints.IrisCreateRoutesPath)
	if err != nil || resp.StatusCode() >= 400 {
		if err == nil {
			err = fmt.Errorf("non-OK status code: %d", resp.StatusCode())
		}
		log.Ctx(ctx).Err(err).Msg("Hestia: CreateIrisRoutes")
		return respJson, err
	}
	h.PrintRespJson(resp.Body())
	return respJson, err
}

func (h *Hestia) ReadIrisRoutes(ctx context.Context, rr any) (hestia_resp_types.Response, error) {
	h.PrintReqJson(rr)
	respJson := hestia_resp_types.Response{}
	resp, err := h.R().
		SetBody(rr).
		SetResult(&respJson).
		Post(hestia_endpoints.IrisReadRoutesPath)
	if err != nil || resp.StatusCode() >= 400 {
		if err == nil {
			err = fmt.Errorf("non-OK status code: %d", resp.StatusCode())
		}
		log.Ctx(ctx).Err(err).Msg("Hestia: ReadIrisRoutes")
		return respJson, err
	}
	h.PrintRespJson(resp.Body())
	return respJson, err
}

func (h *Hestia) UpdateIrisRoutes(ctx context.Context, rr any) (hestia_resp_types.Response, error) {
	h.PrintReqJson(rr)
	respJson := hestia_resp_types.Response{}
	resp, err := h.R().
		SetBody(rr).
		SetResult(&respJson).
		Post(hestia_endpoints.IrisUpdateRoutesPath)
	if err != nil || resp.StatusCode() >= 400 {
		if err == nil {
			err = fmt.Errorf("non-OK status code: %d", resp.StatusCode())
		}
		log.Ctx(ctx).Err(err).Msg("Hestia: IrisUpdateRoutesPath")
		return respJson, err
	}
	h.PrintRespJson(resp.Body())
	return respJson, err
}

func (h *Hestia) DeleteIrisRoutes(ctx context.Context, rr any) (hestia_resp_types.Response, error) {
	h.PrintReqJson(rr)
	respJson := hestia_resp_types.Response{}
	resp, err := h.R().
		SetBody(rr).
		SetResult(&respJson).
		Post(hestia_endpoints.IrisDeleteRoutesPath)
	if err != nil || resp.StatusCode() >= 400 {
		if err == nil {
			err = fmt.Errorf("non-OK status code: %d", resp.StatusCode())
		}
		log.Ctx(ctx).Err(err).Msg("Hestia: IrisDeleteRoutesPath")
		return respJson, err
	}
	h.PrintRespJson(resp.Body())
	return respJson, err
}