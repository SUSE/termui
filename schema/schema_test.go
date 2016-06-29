package schema_test

import (
	"bufio"
	"bytes"
	"io/ioutil"
	"testing"

	"github.com/hpcloud/termui/schema"

	term "github.com/cloudfoundry/cli/cf/terminal"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

// TestSchema tests schema package
func TestSchema(t *testing.T) {
	RegisterFailHandler(Fail)

	RunSpecs(t, "Schema")
}

var _ = Describe("Schema", func() {
	Describe("Test if schema asks correctly for a string", func() {
		Context("ask for string", func() {
			It("should be Insert string value for /stringtest [required]", func() {

				str := `{
   "type":"object",
   "properties":{
      "stringtest":{
         "type":"string"
      }
   },
   "required":[
      "stringtest"
   ]
}`

				in := bytes.Buffer{}
				out := bytes.Buffer{}
				writer := bufio.NewWriter(&out)

				ui := term.NewUI(
					&in,
					writer,
					term.NewTeePrinter(writer),
					nil,
				)

				schema.NewSchemaParser(ui).ParseSchema(str)
				writer.Flush()

				result, err := ioutil.ReadAll(&out)

				Expect(err).NotTo(HaveOccurred())
				Expect(result).To(ContainSubstring("Insert string value for /stringtest [required]"))
			})
		})
	})

	Describe("Test if schema asks correctly for a integer", func() {
		Context("ask for integer", func() {
			It("should be Insert integer value for /integertest [required]", func() {

				str := `{
   "type":"object",
   "properties":{
      "integertest":{
         "type":"integer"
      }
   },
   "required":[
      "integertest"
   ]
}`

				in := bytes.Buffer{}
				out := bytes.Buffer{}
				writer := bufio.NewWriter(&out)

				ui := term.NewUI(
					&in,
					writer,
					term.NewTeePrinter(writer),
					nil,
				)

				schema.NewSchemaParser(ui).ParseSchema(str)
				writer.Flush()

				result, err := ioutil.ReadAll(&out)

				Expect(err).NotTo(HaveOccurred())
				Expect(result).To(ContainSubstring("Insert integer value for /integertest [required]"))
			})
		})
	})

	Describe("Test if schema asks correctly for a number", func() {
		Context("ask for integer", func() {
			It("should be Insert numeric value for /numbertest [required]", func() {

				str := `{
   "type":"object",
   "properties":{
      "numbertest":{
         "type":"number"
      }
   },
   "required":[
      "numbertest"
   ]
}`

				in := bytes.Buffer{}
				out := bytes.Buffer{}
				writer := bufio.NewWriter(&out)

				ui := term.NewUI(
					&in,
					writer,
					term.NewTeePrinter(writer),
					nil,
				)

				schema.NewSchemaParser(ui).ParseSchema(str)
				writer.Flush()

				result, err := ioutil.ReadAll(&out)

				Expect(err).NotTo(HaveOccurred())
				Expect(result).To(ContainSubstring("Insert numeric value for /numbertest [required]"))
			})
		})
	})

	Describe("Test if schema asks correctly for a boolean", func() {
		Context("ask for boolean", func() {
			It("should be Insert boolean value for /booleantest [required]", func() {

				str := `{
   "type":"object",
   "properties":{
      "booleantest":{
         "type":"boolean"
      }
   },
   "required":[
      "booleantest"
   ]
}`

				in := bytes.Buffer{}
				out := bytes.Buffer{}
				writer := bufio.NewWriter(&out)

				ui := term.NewUI(
					&in,
					writer,
					term.NewTeePrinter(writer),
					nil,
				)

				schema.NewSchemaParser(ui).ParseSchema(str)
				writer.Flush()

				result, err := ioutil.ReadAll(&out)

				Expect(err).NotTo(HaveOccurred())
				Expect(result).To(ContainSubstring("Insert boolean value for /booleantest [required]"))
			})
		})
	})
})
