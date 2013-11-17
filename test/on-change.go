import (
	. "github.com/azer/on-change"
	"io"
	"os"
)

Describe("OnChange", func (it It) {
	BeforeEach(func () {
		file, _ := os.Create("foobar.txt")
		defer file.Close()

		io.WriteString(file, "hello world")
	})

	AfterEach(func () {
		os.Remove("foobar.txt")
	})

	it("fires callback on any change in the given filename", func (expect Expect) {
		done := make(chan bool)

		OnChange("foobar.txt", func () {
			done <- true
		})

		go func () {
			file, _ := os.Create("foobar.txt")
			defer file.Close()

			io.WriteString(file, "edit")
		}()

		<-done
	})
})
