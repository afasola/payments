package main

import (
	"context"
	"errors"
	"log"
	"net"

	pb "github.com/afasola/payments/payments"
	"google.golang.org/grpc"
)

const (
	port = ":20000"
)

type server struct{}

var stubOptionsResponse *pb.OptionsResponse

func (s *server) Options(ctx context.Context, in *pb.OptionsRequest) (*pb.OptionsResponse, error) {

	log.Printf("Options: %s, %s, %s, %s, %s", in.Segment, in.Msisdn, in.BillType, in.Username, in.Email)
	if in.Segment == "" || in.Msisdn == "" || in.BillType == "" {
		return nil, errors.New("segment, msisdn and billType are mandatory")
	}
	return stubOptionsResponse, nil
}

func (s *server) DeleteCard(ctx context.Context, in *pb.DeleteCardRequest) (*pb.DeleteCardResponse, error) {

	log.Printf("Delete: %s, %s", in.PanLast4, in.Type)
	if in.PanLast4 == "" || in.Type == "" {
		return nil, errors.New("last 4 digits of the pan and the card type are mandatory")
	}
	if err := deleteCardFromTheStub(in.PanLast4, in.Type); err != nil {
		return nil, err
	}
	return &pb.DeleteCardResponse{DeletionResult: "OK"}, nil
}

func main() {
	initStub()
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterPaymentsServer(s, &server{})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %s", err)
	}
}

func initStub() {

	vaulted := true
	apo := []string{"PayPal", "CardLink", "One Click PayPal", "One Click CardLink"}
	cards := []*pb.Card{&pb.Card{PanLast4: "0608", Type: "visa", ExtToken: "DADSDdacadecsDFSASDASD"},
		&pb.Card{PanLast4: "0777", Type: "mastercard", ExtToken: "TKJDFKJBKDJBDFCAKJ"},
		&pb.Card{PanLast4: "1234", Type: "amex", ExtToken: "ateGSE4534tlgafgafga"}}
	to := []int32{5, 10, 20}

	stubOptionsResponse = &pb.OptionsResponse{Vaulted: vaulted, AvailablePaymentOptions: apo, Cards: cards, TopupOptions: to}
}

func deleteCardFromTheStub(digits, cardType string) error {

	var cards []*pb.Card

	for _, value := range stubOptionsResponse.Cards {
		if value.PanLast4 != digits && value.Type != cardType {
			cards = append(cards, value)
		}
	}
	if len(cards) < len(stubOptionsResponse.Cards) {
		stubOptionsResponse.Cards = cards
		log.Printf("Deleted")
		return nil
	}
	return errors.New("Not found")
}
