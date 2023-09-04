package parsenip

import (
	"os"
	"reflect"
	"testing"
)

func TestParse(t *testing.T) {
	format := `
Name: {:s:Name}
Surname: {:s:Surname}
Age: {:d:Age}
Colors: {:a:Colors}
{:i}
Weight: {:f:Weight}
Scores: {:ad:Scores}
Grades: {:af:Grades}`

	target := `


Name: John
Surname: Wayne
Age: 30
Colors: red, blue, green
        
Foo: Boo
Goo: Coo

Weight: 75.5
Scores: 90, 80, 85
Grades: 3.6, 3.7, 4.0

Name: James
Surname: Johnson
Age: 37
Colors: green, blue, green
        
Roo: Loo
Hoo: Poo

Weight: 80.5
Scores: 50, 80, 85
Grades: 3.6, 3.7, 4.0

`

	expected := []map[string]interface{}{
		{
			"Name":    "John",
			"Surname": "Wayne",
			"Age":     30,
			"Colors":  []string{"red", "blue", "green"},
			"Weight":  75.5,
			"Scores":  []int{90, 80, 85},
			"Grades":  []float64{3.6, 3.7, 4.0},
		},
		{
			"Name":    "James",
			"Surname": "Johnson",
			"Age":     37,
			"Colors":  []string{"green", "blue", "green"},
			"Weight":  80.5,
			"Scores":  []int{50, 80, 85},
			"Grades":  []float64{3.6, 3.7, 4.0},
		},
	}

	result, err := Parse(format, target)
	if err != nil {
		t.Errorf("expected no error, got %v", err)
	}

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("expected %v, got %v", expected, result)
	}
}

func TestParseOnTestLog(t *testing.T) {
	format := `
    CA Name{:e}: {:s:CAName}
    Template Name{:e}: {:s:TemplateName}
{:i}
    msPKI-Certificate-Name-Flag{:e}: {:a:MSPKICertificateNameFlag}
    mspki-enrollment-flag{:e}: {:a:MSPKIEnrollmentFlag}
{:i}
    pkiextendedkeyusage{:e}: {:a:PKIExtendedKeyUsage}
    mspki-certificate-application-policy{:e}: {:a:MSPKICertificateApplicationPolicy}`

	dat, err := os.ReadFile("test_log")
	if err != nil {
		panic(err)
	}
	target := string(dat)

	expected := []map[string]interface{}{
		{
			"CAName":                            `SRV03.valhalla.corp\SRV03-Root-CA-01`,
			"TemplateName":                      `ESC1`,
			"MSPKICertificateNameFlag":          []string{`ENROLLEE_SUPPLIES_SUBJECT`},
			"MSPKIEnrollmentFlag":               []string{`INCLUDE_SYMMETRIC_ALGORITHMS`, `PUBLISH_TO_DS`},
			"PKIExtendedKeyUsage":               []string{`Client Authentication`, `Encrypting File System`, `Secure Email`},
			"MSPKICertificateApplicationPolicy": []string{`Client Authentication`, `Encrypting File System`, `Secure Email`},
		},
		{
			"CAName":                            `SRV03.valhalla.corp\SRV03-Root-CA-01`,
			"TemplateName":                      `ESC2`,
			"MSPKICertificateNameFlag":          []string{`SUBJECT_ALT_REQUIRE_UPN`, `SUBJECT_REQUIRE_DIRECTORY_PATH`},
			"MSPKIEnrollmentFlag":               []string{`INCLUDE_SYMMETRIC_ALGORITHMS`, `PUBLISH_TO_DS`, `AUTO_ENROLLMENT`},
			"PKIExtendedKeyUsage":               []string{`Any Purpose`},
			"MSPKICertificateApplicationPolicy": []string{`Any Purpose`},
		},
		{
			"CAName":                            `SRV03.valhalla.corp\SRV03-Root-CA-01`,
			"TemplateName":                      `ESC3`,
			"MSPKICertificateNameFlag":          []string{`SUBJECT_ALT_REQUIRE_UPN`, `SUBJECT_REQUIRE_DIRECTORY_PATH`},
			"MSPKIEnrollmentFlag":               []string{`INCLUDE_SYMMETRIC_ALGORITHMS`, `PUBLISH_TO_DS`, `AUTO_ENROLLMENT`},
			"PKIExtendedKeyUsage":               []string{`Certificate Request Agent`},
			"MSPKICertificateApplicationPolicy": []string{`Certificate Request Agent`},
		},
		{
			"CAName":                            `SRV03.valhalla.corp\SRV03-Root-CA-01`,
			"TemplateName":                      `ESC4`,
			"MSPKICertificateNameFlag":          []string{`SUBJECT_ALT_REQUIRE_UPN`, `SUBJECT_ALT_REQUIRE_EMAIL`, `SUBJECT_REQUIRE_EMAIL`, `SUBJECT_REQUIRE_DIRECTORY_PATH`},
			"MSPKIEnrollmentFlag":               []string{`INCLUDE_SYMMETRIC_ALGORITHMS`, `PUBLISH_TO_DS`, `AUTO_ENROLLMENT`},
			"PKIExtendedKeyUsage":               []string{`Client Authentication`, `Encrypting File System`, `Secure Email`},
			"MSPKICertificateApplicationPolicy": []string{`Client Authentication`, `Encrypting File System`, `Secure Email`},
		},
	}

	result, err := Parse(format, target)
	if err != nil {
		t.Errorf("expected no error, got %v", err)
	}

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("expected %v, got %v", expected, result)
	}
}