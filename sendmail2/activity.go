package sample

import (
	"fmt"
	"os"
	"time"

	"gopkg.in/gomail.v2"

	"github.com/project-flogo/core/activity"
	//"github.com/project-flogo/core/data/metadata"
	// 	"reflect"
)

var (
	activityMd = activity.ToMetadata(&Input{}, &Output{})
)

func init() {
	_ = activity.Register(&Activity{})
}

//New optional factory method, should be used if one activity instance per configuration is desired
func New(ctx activity.InitContext) (activity.Activity, error) {

	act := &Activity{} //add aSetting to instance

	return act, nil
}

// Activity is an sample Activity that can be used as a base to create a custom activity
type Activity struct {
}

// Metadata returns the activity's metadata
func (a *Activity) Metadata() *activity.Metadata {
	return activityMd
}

// Eval implements api.Activity.Eval - Logs the Message
func (a *Activity) Eval(ctx activity.Context) (done bool, err error) {
	input := &Input{}
	err = ctx.GetInputObject(input)
	if err != nil {
		return true, err
	}
	// fmt.Printf("Input serial: %s\n", input.Serial)
	to := input.to
	from := input.from
	subject := input.subject
	location := input.location
	username := input.username
	password := input.password
	imagepath := input.imagepath

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

	output := &Output{result: "The email has been sent to " + to}
	// output := &Output{Serial: `te[{:,"st`}
	err = ctx.SetOutputObject(output)
	if err != nil {
		return true, err
	}

	return true, nil
}

// determine if the file/folder of the given path exists
func exists(path string) bool {

	_, err := os.Stat(path)
	//os.Stat get the file information
	if err != nil {
		if os.IsExist(err) {
			return true
		}
		return false
	}
	return true
}
