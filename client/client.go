package main

import (
	"context"
	"log"
	"time"

	pb "github.com/afasola/payments/payments"
	"google.golang.org/grpc"
)

const (
	address = "localhost:20000"
)

func main() {
	// Set up a connection to the server.
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewPaymentsClient(conn)

	// Contact the server and print out its response.
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	options(ctx, c)
	deleteCard(ctx, c, "1234")
	deleteCard(ctx, c, "1234", "amex")
	options(ctx, c)
	deleteCard(ctx, c, "1234", "amex")
	deleteCard(ctx, c, "0608", "visa")
	deleteCard(ctx, c, "0777", "visa")
	options(ctx, c)
	checkout(ctx, c)

}

func options(ctx context.Context, c pb.PaymentsClient) {
	log.Printf("*** OPTIONS *** \n")
	if r, err := c.Options(ctx, &pb.OptionsRequest{
		Segment:  "prepay",
		Msisdn:   "606060",
		BillType: "BILL_PAYMENT",
		Username: "ciccio",
		Email:    "ciccio@mail.com"}); err != nil {
		log.Printf("Error: %s\n\n", err)
	} else {
		log.Printf("%s\n\n", r)
	}
}

func checkout(ctx context.Context, c pb.PaymentsClient) {
	log.Printf("*** CHECKOUT *** \n")
	if r, err := c.Checkout(ctx, &pb.CheckoutRequest{
		OrderDesc:            "Monthly Payment",
		SegmentType:          "SEGXYZ",
		ActionType:           "TopUp",
		Msisdn:               "606060",
		OrderAmount:          5550,
		ExtToken:             "SADDnlknlDSIDAS87SD(U£N93uqd",
		AdditionalProperties: "{cli: 210xxxxxxx, msisdn: 707070, payerId: N12}"}); err != nil {
		log.Printf("Error: %s\n\n", err)
	} else {
		log.Printf("%s\n\n", r)
	}
}

func deleteCard(ctx context.Context, c pb.PaymentsClient, params ...string) {
	log.Printf("*** DELETE CARD *** \n")
	if len(params) == 2 {
		if r, err := c.DeleteCard(ctx, &pb.DeleteCardRequest{PanLast4: params[0], Type: params[1]}); err != nil {
			log.Printf("Error while deleting (%s, %s): %s\n\n", params[0], params[1], err)
		} else {
			log.Printf("%s\n\n", r)
		}
	}
	if len(params) == 1 {
		if r, err := c.DeleteCard(ctx, &pb.DeleteCardRequest{PanLast4: params[0]}); err != nil {
			log.Printf("Error while deleting (%s, nil): %s\n\n", params[0], err)
		} else {
			log.Printf("%s\n\n", r)
		}
	}
}
