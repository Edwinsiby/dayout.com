package Test

import (
	"main/Handlers"
	"net/http"
	"net/http/httptest"
	"net/url"
	"regexp"
	"strings"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/assert/v2"
	_ "github.com/stretchr/testify/assert"
)

// TestHelloName calls greetings.Hello with a name, checking
// for a valid return value.
func TestHelloName(t *testing.T) {
	name := "Gladys"
	want := regexp.MustCompile(`\b` + name + `\b`)
	msg, err := Handlers.Hello("Gladys")
	if !want.MatchString(msg) || err != nil {
		t.Fatalf(`Hello("Gladys") = %q, %v, want match for %#q, nil`, msg, err, want)
	}
}

// TestHelloEmpty calls greetings.Hello with an empty string,
// checking for an error.
func TestHelloEmpty(t *testing.T) {
	msg, err := Handlers.Hello("")
	if msg != "" || err == nil {
		t.Fatalf(`Hello("") = %q, %v, want "", error`, msg, err)
	}
}

func TestSignupPost(t *testing.T) {
	// Set up test environment
	router := gin.Default()
	router.POST("/signup", Handlers.SignupPost)

	form := url.Values{}
	form.Set("UserName", "JohnDoe")
	form.Set("Emailid", "john@example.com")
	form.Set("Password", "password123")

	request, _ := http.NewRequest("POST", "/signup", strings.NewReader(form.Encode()))
	request.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	recorder := httptest.NewRecorder()
	context, _ := gin.CreateTestContext(recorder)
	context.Request = request

	// Invoke the function being tested
	router.ServeHTTP(recorder, request)

	// Validate the results
	assert.Equal(t, http.StatusFound, recorder.Code)
	assert.Equal(t, "/login", recorder.Header().Get("Location"))
}
