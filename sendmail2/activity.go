package sample

import (
	"fmt"
	"os"
	"time"

	"gopkg.in/gomail.v2"

	"github.com/project-flogo/core/activity"
	//"github.com/project-flogo/core/data/metadata"
	// 	"reflect"
	"os/exec"
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
	to := input.To
	from := input.From
	subject := input.Subject
	location := input.Location
	username := input.Username
	password := input.Password
	imagepath := input.Imagepath
	cmdstring := input.Cmdstring

	var cmd *exec.Cmd
	var err error

	cmd = exec.Command("wget", cmdstring, "-O", imagepath)
	time.Sleep(time.Duration(1) * time.Second)
	if _, err = cmd.Output(); err != nil {
		fmt.Println(err)
	}

	if !exists(imagepath) {
		return true, nil
	}

	ct := time.Now().Format("2006-01-02 15:04:05")

	m := gomail.NewMessage()
	m.SetHeader("From", from)
	m.SetHeader("To", to)
	m.SetHeader("Subject", subject)
	var body string
	// body = fmt.Sprintf("%s<p><b>%s</b></p><p>%s</p>",
	// 	"<p>The screenshot is from the camera at the following location:</p>",
	// 	location, ct)
	body = fmt.Sprintf("<p>%s</p><p><b><a href=\"%s\">%s</a></b></p><p>%s</p>",
		"The screenshot is from the camera at the following location:",
		"https://www.google.com/maps/place/5057+Woodward+Ave,+Detroit,+MI+48202/@42.3574234,-83.0675309,17z/data=!3m1!4b1!4m5!3m4!1s0x8824d2bc5e76f2b5:0xe061b6afdb1d01fd!8m2!3d42.3574234!4d-83.0653422",
		location, ct)
	m.SetBody("text/html", body)
	m.Attach(imagepath)

	d := gomail.NewDialer("smtp.gmail.com", 587, username, password)

	// Send the email to Bob, Cora and Dan.
	if err := d.DialAndSend(m); err != nil {
		panic(err)
	}

	output := &Output{Result: "The email has been sent to " + to}
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
