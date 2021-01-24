package main

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"net/http"
	"encoding/json"
	"io/ioutil"
)

func resourceVkStatus() *schema.Resource {
	return &schema.Resource{
		Create: resourceVkStatusCreate,
		Read:   resourceVkStatusRead,
		Update: resourceVkStatusUpdate,
		Delete: resourceVkStatusDelete,
		Exists: resourceVkStatusExists,
		Importer: &schema.ResourceImporter{
			State: resourceVkStatusImportState,
		},

		Schema: map[string]*schema.Schema{
			"status_text": &schema.Schema{
				Type:        schema.TypeString,
				Description: "Sets the status text",
				Required:     true,
			},
		},
	}
}

func resourceVkStatusExists(d *schema.ResourceData, meta interface{}) (b bool, e error) {
	return true, nil
}

func resourceVkStatusCreate(d *schema.ResourceData, meta interface{}) error {
	access_token := meta.(*Config).ACCESSToken
	URL := "https://api.vk.com/method/status.set?text="+d.Get("status_text").(string)
	URL = URL+"&access_token="+access_token+"&v=5.126"

	// set status
	resp, err := http.Get(URL)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	return nil
}

func resourceVkStatusRead(d *schema.ResourceData, meta interface{}) error {
  return nil
}

func resourceVkStatusUpdate(d *schema.ResourceData, meta interface{}) error {
	access_token := meta.(*Config).ACCESSToken
	URL := "https://api.vk.com/method/status.set?text="+d.Get("status_text").(string)
	URL = URL+"&access_token="+access_token+"&v=5.126"

	// set status
	resp, err := http.Get(URL)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	return nil
}

func resourceVkStatusDelete(d *schema.ResourceData, meta interface{}) error {
	access_token := meta.(*Config).ACCESSToken
	URL := "https://api.vk.com/method/status.set?text="
	URL = URL+"&access_token="+access_token+"&v=5.126"

	// set status
	resp, err := http.Get(URL)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	return nil
}

func resourceVkStatusImportState(d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
	access_token := meta.(*Config).ACCESSToken
	URL := "https://api.vk.com/method/status.get?access_token="+access_token+"&v=5.126"

	// get status
	resp, err := http.Get(URL)
	if err != nil {
		return []*schema.ResourceData{d},err
	}
	defer resp.Body.Close()
	responseData,err := ioutil.ReadAll(resp.Body)
  if err != nil {
    return []*schema.ResourceData{d},err
  }
	var result map[string]interface{}
	json.Unmarshal([]byte(responseData), &result)

	d.Set("status_text", result["response"].([]interface{})[0].(map[string]interface{})["text"])

	return []*schema.ResourceData{d}, nil
}
