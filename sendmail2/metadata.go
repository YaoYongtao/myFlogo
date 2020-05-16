package sample

import "github.com/project-flogo/core/data/coerce"

type Input struct {
	to        string `md:"to,required"`
	from      string `md:"from,required"`
	subject   string `md:"subject,required"`
	location  string `md:"location,required"`
	username  string `md:"username,required"`
	password  string `md:"password,required"`
	imagepath string `md:"imagepath,required"`
}

func (r *Input) FromMap(values map[string]interface{}) error {
	strTo, _ := coerce.ToString(values["to"])
	r.to = strTo
	strFrom, _ := coerce.ToString(values["from"])
	r.from = strFrom
	strSubject, _ := coerce.ToString(values["subject"])
	r.subject = strSubject
	strLocation, _ := coerce.ToString(values["location"])
	r.location = strLocation
	strUsername, _ := coerce.ToString(values["username"])
	r.username = strUsername
	strPassword, _ := coerce.ToString(values["password"])
	r.password = strPassword
	strImagepath, _ := coerce.ToString(values["imagepath"])
	r.imagepath = strImagepath
	return nil
}

func (r *Input) ToMap() map[string]interface{} {
	return map[string]interface{}{
		"to": r.to,
		"from": r.from,
		"subject": r.subject,
		"location": r.location,
		"username": r.username,
		"password": r.password,
		"imagePath":r.imagepath
	}
}

type Output struct {
	result string `md:"result"`
}

func (o *Output) FromMap(values map[string]interface{}) error {
	strResult, _ := coerce.ToString(values["result"])
	o.result = strResult
	return nil
}

func (o *Output) ToMap() map[string]interface{} {
	return map[string]interface{}{
		"result": o.result
	}
}
