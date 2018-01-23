// Package test provides test utilities.
package test

import (
	"os"
	"strings"
	"testing"

	"github.com/ghettovoice/gosip/log"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

// InitTestSuite initializes ginkgo test suite.
func InitTestSuite(t *testing.T, desc string) {
	// setup logger
	lvl := log.ErrorLevel
	forceColor := true
	for _, arg := range os.Args {
		if strings.HasPrefix(arg, "--test.v") || strings.HasPrefix(arg, "--ginkgo.v") {
			lvl = log.DebugLevel
		} else if strings.HasPrefix(arg, "--ginkgo.noColor") {
			forceColor = false
		}
	}
	log.SetLevel(lvl)
	log.SetFormatter(log.NewFormatter(true, forceColor))

	// setup Ginkgo
	RegisterFailHandler(Fail)
	RegisterTestingT(t)
	RunSpecs(t, desc)
}
