package sample

import "github.com/project-flogo/core/data/coerce"

type Input struct {
	To        string `md:"to,required"`
	From      string `md:"from,required"`
	Subject   string `md:"subject,required"`
	Location  string `md:"location,required"`
	Username  string `md:"username,required"`
	Password  string `md:"password,required"`
	Imagepath string `md:"imagepath,required"`
}

func (r *Input) FromMap(values map[string]interface{}) error {
	strTo, _ := coerce.ToString(values["to"])
	r.To = strTo
	strFrom, _ := coerce.ToString(values["from"])
	r.From = strFrom
	strSubject, _ := coerce.ToString(values["subject"])
	r.Subject = strSubject
	strLocation, _ := coerce.ToString(values["location"])
	r.Location = strLocation
	strUsername, _ := coerce.ToString(values["username"])
	r.Username = strUsername
	strPassword, _ := coerce.ToString(values["password"])
	r.Password = strPassword
	strImagepath, _ := coerce.ToString(values["imagepath"])
	r.Imagepath = strImagepath
	return nil
}

func (r *Input) ToMap() map[string]interface{} {
	return map[string]interface{}{
		"to":        r.To,
		"from":      r.From,
		"subject":   r.Subject,
		"location":  r.Location,
		"username":  r.Username,
		"password":  r.Password,
		"imagePath": r.Imagepath,
	}
}

type Output struct {
	Result string `md:"result"`
}

func (o *Output) FromMap(values map[string]interface{}) error {
	strResult, _ := coerce.ToString(values["result"])
	o.Result = strResult
	return nil
}

func (o *Output) ToMap() map[string]interface{} {
	return map[string]interface{}{
		"result": o.Result,
	}
}
