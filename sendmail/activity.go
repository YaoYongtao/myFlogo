package sendmail

import (
	"fmt"
	"time"

	"github.com/TIBCOSoftware/flogo-lib/core/activity"
	"gopkg.in/gomail.v2"
)

// MyActivity is a stub for your Activity implementation
type MyActivity struct {
	metadata *activity.Metadata
}

// NewActivity creates a new activity
func NewActivity(metadata *activity.Metadata) activity.Activity {
	return &MyActivity{metadata: metadata}
}

// Metadata implements activity.Activity.Metadata
func (a *MyActivity) Metadata() *activity.Metadata {
	return a.metadata
}

// Eval implements activity.Activity.Eval
func (a *MyActivity) Eval(context activity.Context) (done bool, err error) {

	// set inputs
	to := context.GetInput("to").(string)
	from := context.GetInput("from").(string)
	subject := context.GetInput("subject").(string)
	location := context.GetInput("location").(string)
	username := context.GetInput("username").(string)
	password := context.GetInput("password").(string)
	imagepath := context.GetInput("imagepath").(string)

	ct := time.Now().Format("2006-01-02 15:04:05")

	m := gomail.NewMessage()
	m.SetHeader("From", from)
	m.SetHeader("To", to)
	m.SetHeader("Subject", subject)
	var body string
	body = fmt.Sprintf("%s<p><b>%s</b></p><p>%s</p>",
		"<p>The screenshot is from the camera at the following location:</p>",
		location, ct)
	m.SetBody("text/html", body)
	m.Attach(imagepath)

	d := gomail.NewDialer("smtp.gmail.com", 587, username, password)

	// Send the email to Bob, Cora and Dan.
	if err := d.DialAndSend(m); err != nil {
		panic(err)
	}
	context.SetOutput("result", "The email has been sent to "+to)

	return true, nil
}
