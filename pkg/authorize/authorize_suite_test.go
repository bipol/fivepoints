package authorize_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestAuthorize(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Authorize Suite")
}
