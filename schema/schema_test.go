package schema_test

import (
	"bytes"
	"io"
	"os"
	"testing"

	"github.com/hpcloud/termui/schema"

	term "github.com/cloudfoundry/cli/cf/terminal"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

// tests for schema
func TestSchemaString(t *testing.T) {
	RegisterFailHandler(Fail)

	RunSpecs(t, "Schema")
}

var _ = Describe("Schema", func() {
	Describe("Test if schema asks correctly for a string", func() {
		Context("ask for string", func() {
			It("should be Insert string value for /stringtest [required]", func() {
				ui := term.NewUI(os.Stdin, term.NewTeePrinter())

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

				old := os.Stdout
				r, w, _ := os.Pipe()
				os.Stdout = w

				schemaParser := schema.NewSchemaParser(ui)
				schemaParser.ParseSchema(str)

				outC := make(chan string)

				go func() {
					var buf bytes.Buffer
					io.Copy(&buf, r)
					outC <- buf.String()
				}()

				w.Close()
				os.Stdout = old
				out := <-outC
				Expect(out).To(ContainSubstring("Insert string value for /stringtest [required]"))
			})
		})
	})

	Describe("Test if schema asks correctly for a integer", func() {
		Context("ask for integer", func() {
			It("should be Insert integer value for /integertest [required]", func() {
				ui := term.NewUI(os.Stdin, term.NewTeePrinter())

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

				old := os.Stdout
				r, w, _ := os.Pipe()
				os.Stdout = w

				schemaParser := schema.NewSchemaParser(ui)
				schemaParser.ParseSchema(str)

				outC := make(chan string)

				go func() {
					var buf bytes.Buffer
					io.Copy(&buf, r)
					outC <- buf.String()
				}()

				w.Close()
				os.Stdout = old
				out := <-outC
				Expect(out).To(ContainSubstring("Insert integer value for /integertest [required]"))
			})
		})
	})

	Describe("Test if schema asks correctly for a number", func() {
		Context("ask for integer", func() {
			It("should be Insert numeric value for /numbertest [required]", func() {
				ui := term.NewUI(os.Stdin, term.NewTeePrinter())

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

				old := os.Stdout
				r, w, _ := os.Pipe()
				os.Stdout = w

				schemaParser := schema.NewSchemaParser(ui)
				schemaParser.ParseSchema(str)

				outC := make(chan string)

				go func() {
					var buf bytes.Buffer
					io.Copy(&buf, r)
					outC <- buf.String()
				}()

				w.Close()
				os.Stdout = old
				out := <-outC
				Expect(out).To(ContainSubstring("Insert numeric value for /numbertest [required]"))
			})
		})
	})

	Describe("Test if schema asks correctly for a boolean", func() {
		Context("ask for boolean", func() {
			It("should be Insert boolean value for /booleantest [required]", func() {
				ui := term.NewUI(os.Stdin, term.NewTeePrinter())

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

				old := os.Stdout
				r, w, _ := os.Pipe()
				os.Stdout = w

				schemaParser := schema.NewSchemaParser(ui)
				schemaParser.ParseSchema(str)

				outC := make(chan string)

				go func() {
					var buf bytes.Buffer
					io.Copy(&buf, r)
					outC <- buf.String()
				}()

				w.Close()
				os.Stdout = old
				out := <-outC
				Expect(out).To(ContainSubstring("Insert boolean value for /booleantest [required]"))
			})
		})
	})
})
