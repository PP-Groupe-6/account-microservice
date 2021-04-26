package account_microservice

import (
	"context"
	"strings"

	"github.com/go-kit/kit/endpoint"
)

type AccountEndpoints struct {
	GetAmountEndpoint          endpoint.Endpoint
	GetUserInformationEndpoint endpoint.Endpoint
	AddEndpoint                endpoint.Endpoint
	UpdateEndpoint             endpoint.Endpoint
	DeleteEndpoint             endpoint.Endpoint
}

func MakeAccountEndpoints(s AccountService) AccountEndpoints {
	return AccountEndpoints{
		GetAmountEndpoint:          MakeGetAmountEndpoint(s),
		GetUserInformationEndpoint: MakeGetUserInformationEndpoint(s),
		AddEndpoint:                MakeAddEndpoint(s),
	}
}

type GetAmountRequest struct {
	ClientID string
}

type GetAmountResponse struct {
	AccountAmount float32 `json:"amount"`
}

func MakeGetAmountEndpoint(s AccountService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(GetAmountRequest)
		amount, err := s.GetAmountForID(ctx, req.ClientID)

		return GetAmountResponse{float32(amount)}, err
	}
}

type GetUserInformationRequest struct {
	ClientID string
}

type GetUserInformationResponse struct {
	FullName    string `json:"fullName"`
	MailAdress  string `json:"mailAdress"`
	PhoneNumber string `json:"phoneNumber"`
}

func MakeGetUserInformationEndpoint(s AccountService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(GetUserInformationRequest)
		account, err := s.GetAccountByID(ctx, req.ClientID)
		formatedName := account.Surname + " " + account.Name
		return GetUserInformationResponse{formatedName, account.MailAdress, account.PhoneNumber}, err
	}
}

type AddRequest struct {
	ClientID    string
	FullName    string
	PhoneNumber string
	MailAdress  string
}

type AddResponse struct {
	Added bool `json:"added"`
}

func MakeAddEndpoint(s AccountService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(AddRequest)
		sepName := strings.Split(req.FullName, " ")
		toAdd := Account{
			req.ClientID,
			sepName[0],
			sepName[1],
			req.PhoneNumber,
			req.MailAdress,
			0,
		}
		account, err := s.Add(ctx, toAdd)
		if (err == nil && account != Account{}) {
			return AddResponse{true}, nil
		} else {
			return AddResponse{false}, err
		}
	}
}
