package crawl

import (
	"bytes"
	"github.com/PuerkitoBio/goquery"
	"golang.org/x/text/encoding/simplifiedchinese"
	"io/ioutil"
	"net/http"
	"shuTeacher/model"
)

func TeacherInfo(id string) (model.Teacher, error) {
	response, err := http.Get("http://jwc.shu.edu.cn:8080/jwc/tinfo/viewinfo1.jsp?tid=" + id)
	if err != nil {
		return model.Teacher{}, err
	}
	// f**k GBK
	body, _ := ioutil.ReadAll(response.Body)
	decodeBytes, _ := simplifiedchinese.GB18030.NewDecoder().Bytes(body)
	doc, err := goquery.NewDocumentFromReader(bytes.NewReader(decodeBytes))
	if err != nil {
		return model.Teacher{}, err
	}
	name := doc.Find(".chn tr:first-of-type td:nth-of-type(2)>strong").Text()
	result := model.Teacher{
		Id:   id,
		Name: name,
	}
	model.Save(result)
	return result, nil
}
