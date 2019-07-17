package authorize_test

import (
	"context"
	"fmt"

	"github.com/dgrijalva/jwt-go"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"google.golang.org/grpc/metadata"

	"github.com/smartatransit/fivepoints/pkg/authorize"
)

var _ = Describe("Authorizer", func() {
	Context("IsAuthorized", func() {
		var (
			client     authorize.Client
			authorized bool
			err        error
			ctx        context.Context
			secret     string
		)
		JustBeforeEach(func() {
			client = authorize.NewClient(secret)
			authorized, err = client.IsAuthorized(ctx)
		})
		BeforeEach(func() {
			secret = "test"
			// Create the Claims
			claims := make(jwt.MapClaims)
			claims["fvp.Do"] = true
			token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
			ss, _ := token.SignedString([]byte(secret))
			md := metadata.Pairs("authorization", fmt.Sprintf("Bearer %s", ss))
			ctx = metadata.NewIncomingContext(context.Background(), md)
		})
		When("not metadata is provided", func() {
			BeforeEach(func() {
				ctx = context.Background()
			})
			It("does return an error", func() {
				Expect(err).ToNot(BeNil())
			})
		})
		When("authorization is not provided", func() {
			BeforeEach(func() {
				md := metadata.Pairs()
				ctx = metadata.NewIncomingContext(context.Background(), md)
			})
			It("does return an error", func() {
				Expect(err).ToNot(BeNil())
			})
		})
		When("jwt uses wrong signing method", func() {
			BeforeEach(func() {
				claims := make(jwt.MapClaims)
				claims["fvp.Do"] = true
				token := jwt.NewWithClaims(jwt.SigningMethodRS384, claims)
				ss, _ := token.SignedString([]byte(secret))
				md := metadata.Pairs("authorization", fmt.Sprintf("Bearer %s", ss))
				ctx = metadata.NewIncomingContext(context.Background(), md)
			})
			It("does return an error", func() {
				Expect(err).ToNot(BeNil())
			})
		})
		When("jwt has invalid claim", func() {
			BeforeEach(func() {
				claims := make(jwt.MapClaims)
				token := jwt.NewWithClaims(jwt.SigningMethodHS384, claims)
				ss, _ := token.SignedString([]byte(secret))
				md := metadata.Pairs("authorization", fmt.Sprintf("Bearer %s", ss))
				ctx = metadata.NewIncomingContext(context.Background(), md)
			})
			It("does return an error", func() {
				Expect(err).ToNot(BeNil())
			})
		})
		When("a valid token is provided", func() {
			It("is authorized", func() {
				Expect(authorized).To(BeTrue())
			})
			It("does not return an error", func() {
				Expect(err).To(BeNil())
			})
		})
	})

})
